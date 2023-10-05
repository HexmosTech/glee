import jwt
import requests
import argparse
import frontmatter
import markdown
from markdown.extensions.toc import TocExtension
from markdown.extensions.fenced_code import FencedCodeExtension
from markdown.extensions.codehilite import CodeHiliteExtension
from markdown.extensions.tables import TableExtension

from styles import  default_style,sidebar_toc_head, sidebar_toc_footer
from code_theme import select_codehilite_theme
from images import ImgExtExtension
from hasher import sha256sum
from s3 import upload_to_s3
from bs4 import BeautifulSoup
import os
import shutil

from datetime import datetime as date
from handle_config import (
    view_toml_file,
    check_configurations_exist,
    print_configuration,
)
from ghost_upload_image import upload_to_ghost, get_images_from_post
import logging
import sys


parser = argparse.ArgumentParser(description="Publish Markdown Files to Ghost Blog")
parser.add_argument(
    "--config", action="store_true", help="Show glee configuration file"
)
parser.add_argument(
    "markdown_file", nargs="?", type=str, help="<path_to_markdown_file>"
)
parser.add_argument("--debug", action="store_true", help="Enable debug mode")

args = parser.parse_args()
if args.config:
    print_configuration(logging)

if not args.markdown_file:
    print("Usage: glee <path_to_markdown_file>")
    sys.exit(1)


log_format = "%(levelname)s:%(message)s"

logging.basicConfig(
    level=logging.DEBUG if args.debug else logging.INFO, format=log_format
)


config, config_path = view_toml_file(logging)

check_configurations_exist(logging)

GHOST_VERSION = config["ghost-configuration"]["GHOST_VERSION"]
IMAGE_BACKEND = config["image-configuration"]["IMAGE_BACKEND"]
GHOST_URL = config["ghost-configuration"]["GHOST_URL"]

if GHOST_VERSION == "v5":
    POSTS_API_BASE = f"{GHOST_URL}/api/admin/posts/"
else:
    POSTS_API_BASE = f"{GHOST_URL}/api/{GHOST_VERSION}/admin/posts/"


mdlib = markdown.Markdown(
    extensions=[
        TocExtension(),
        FencedCodeExtension(),
        CodeHiliteExtension(),
        ImgExtExtension(),
        TableExtension(),
    ]
)


def to_html(md):
    start = "<!--kg-card-begin: html-->"
    end = "<!--kg-card-end: html-->"
    html = mdlib.convert(md)
    return start + html + end


def get_jwt():
    key = config["ghost-configuration"]["ADMIN_API_KEY"]
    try:
        id, secret = key.split(":")
        if GHOST_VERSION == "v5":
            aud_value = "/admin/"
        else:
            aud_value = f"/{GHOST_VERSION}/admin/"
        iat = int(date.now().timestamp())

        h = {"iat": iat, "exp": iat + 5 * 60, "aud": aud_value}
        token = jwt.encode(
            h, bytes.fromhex(secret), algorithm="HS256", headers={"kid": id}
        )
        return token
    except Exception as e:
        logging.error(
            f"Unable to generate the JWT token. Please check the Admin API key.({e})"
        )
        sys.exit(1)


def get_post_id(slug, headers):
    try:
        url = f"{POSTS_API_BASE}slug/{slug}/"
        r = requests.get(url, headers=headers)
        if r.ok:
            j = r.json()
            return (
                j["posts"][0]["id"],
                j["posts"][0]["updated_at"],
                j["posts"][0]["mobiledoc"],
                j["posts"][0]["feature_image"],
            )
        else:
            if r.status_code == 404:
                return None, None, None, None
            sys.exit(f"Unable to communicate with the Ghost Admin API:{r.json()}")
    except Exception as e:
        sys.exit(
            f"Error:Unable to communicate with the Ghost Admin API. Please verify your Ghost configurations: {e}"
        )


def make_request(headers, body, pid, updated_at):
    try:
        if not pid:
            url = f"{POSTS_API_BASE}?source=html"
            r = requests.post(url, json=body, headers=headers)
            if r.status_code == 200:
                preview_link = r.json()["posts"][0]["url"]

                logging.info("Created new post")
                logging.info(f"Blog preview link: {preview_link}")
            else:
                raise Exception(r.json())
        else:
            body["posts"][0]["updated_at"] = updated_at

            url = f"{POSTS_API_BASE}{pid}?source=html"

            r = requests.put(url, json=body, headers=headers)
            if r.status_code == 200:
                preview_link = r.json()["posts"][0]["url"]
                logging.info(f"Updated existing post based on slug")
                logging.info(f"Blog preview link: {preview_link}")
            else:
                raise Exception(r.json())
    except Exception as e:
        logging.error(f"Error:{e}")

    return


def get_temp_dir():
    if os.name == "nt":  # Check if the operating system is Windows
        return os.environ["TEMP"]
    else:
        return "/tmp/"


def image_to_hash(image):
    tp = ""
    if image.startswith("http://") or image.startswith("https://"):
        iext = image.split(".")[-1]
        temp_dir = get_temp_dir()
        tp = os.path.join(temp_dir, "img." + iext)
        response = requests.get(image, stream=True)
        with open(tp, "wb") as out_file:
            shutil.copyfileobj(response.raw, out_file)
        del response
        hash_value = sha256sum(tp)
    else:
        hash_value = sha256sum(image)
    _, file_extension = os.path.splitext(image)
    image_name = hash_value + file_extension
    return image_name, tp


def upload_images(token, html_data, IMAGE_BACKEND):
    # add two options s3 or ghost upload
    try:
        uploaded_images = {}
        if IMAGE_BACKEND == "ghost":
            blog_image_list = get_images_from_post(html_data)

        for image in mdlib.images:
            hash_value, image_data = image_to_hash(image)

            if image.startswith("http://") or image.startswith("https://"):
                if IMAGE_BACKEND == "s3":
                    image_link = upload_to_s3(image_data, hash_value, logging)
                else:
                    # image comparison here
                    image_link = upload_to_ghost(
                        token, image_data, hash_value, blog_image_list, logging
                    )

            else:
                if IMAGE_BACKEND == "s3":
                    image_link = upload_to_s3(image, hash_value, logging)
                else:
                    image_link = upload_to_ghost(
                        token, image, hash_value, blog_image_list, logging
                    )

            uploaded_images[image] = image_link

        logging.info("Uploaded images")
        return uploaded_images
    except:
        sys.exit("Error: Add content")


def replace_image_links(post, img_map):
    soup = BeautifulSoup(post["html"], features="html.parser")
    for img in soup.find_all("img"):
        img["src"] = img_map[img["src"]]
    result = str(soup)
    post["html"] = result


def upload_feature_image(meta, token, feature_image):
    try:
        i = meta["feature_image"]
        if i:
            hash_value = sha256sum(i)
            _, file_extension = os.path.splitext(i)
            image_name = hash_value + file_extension

            if IMAGE_BACKEND == "s3":
                meta["feature_image"] = upload_to_s3(i, image_name, logging)
            else:
                if feature_image is not None:
                    feature_img_list = [feature_image]
                else:
                    feature_img_list = []
                image_link = upload_to_ghost(
                    token, i, image_name, feature_img_list, logging
                )
                meta["feature_image"] = image_link

            logging.info("Uploaded feature image")
    except Exception as e:
        logging.error("Error in feature image uploading", e)

        pass


def post_to_ghost(meta, md):
    if not "slug" in meta:
        logging.error(
            "ERROR: Include a URL friendly slug field in your markdown file and retry! This is required to support updates"
        )

        return

    meta = add_blog_configurations(meta)
    meta["html"] = to_html(md)
    token = get_jwt()

    headers = {"Authorization": "Ghost {}".format(token)}
    pid, updated_at, html_data, feature_image = get_post_id(meta["slug"], headers)
    if "feature_image" in meta:
        upload_feature_image(meta, token, feature_image)
    else:
        meta["feature_image"] = ""
    uploaded_images = upload_images(token, html_data, IMAGE_BACKEND)

    replace_image_links(meta, uploaded_images)
    post_obj = meta

    body = {"posts": [post_obj]}
    return make_request(headers, body, pid, updated_at)


def add_blog_configurations(meta):
    try:
        global_sidebar_toc = config.get("blog-configuration", {}).get("SIDEBAR_TOC")
        global_featured = config.get("blog-configuration", {}).get("FEATURED")
        global_status = config.get("blog-configuration", {}).get("STATUS")
        global_theme = config.get("blog-configuration", {}).get("CODE-HILITE-THEME")

        side_bar_toc = meta.get("sidebar_toc", global_sidebar_toc)
        code_theme = meta.get("code_hilite_theme",global_theme)
        meta["featured"] = meta.get("featured", global_featured)
        meta["status"] = meta.get("status", global_status)
        if meta["status"] is None or meta["featured"] is None:
            raise Exception("required featured and status")
        
        # theme manipulation from here
        theme = select_codehilite_theme(code_theme)
        
        if side_bar_toc:
            meta["codeinjection_head"] =  f"""<style>{default_style+theme}</style>""" + sidebar_toc_head
            meta["codeinjection_foot"] = sidebar_toc_footer
        else:
            
            meta["codeinjection_head"] = f"""<style>{default_style+theme}</style>"""

        return meta

    except Exception as e:
        sys.exit(f'''Error: {e}''')


if __name__ == "__main__":
    post = frontmatter.load(args.markdown_file)
    post_to_ghost(post.metadata, post.content)
