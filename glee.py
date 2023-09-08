import jwt
import arrow
import requests
import argparse
import frontmatter
import markdown
from markdown.extensions.toc import TocExtension
from markdown.extensions.fenced_code import FencedCodeExtension
from markdown.extensions.codehilite import CodeHiliteExtension
from markdown.extensions.tables import TableExtension
import sys
from styles import style, sidebar_toc_head, sidebar_toc_footer
from images import ImgExtExtension
from hasher import sha256sum
from s3 import upload_to_s3
from bs4 import BeautifulSoup
import os
import shutil
import toml
from datetime import datetime as date
from handle_config import get_toml_file, crediential_not_found


# Load the TOML file
try:
    config_path = os.path.join(os.path.expanduser("~"), ".glee.toml")
    config = toml.load(config_path)
except:
    get_toml_file(config_path)
    sys.exit(0)


GHOST_VERSION = config["ghost-configuration"]["GHOST_VERSION"]

if GHOST_VERSION == "":
    crediential_not_found(config_path)


if GHOST_VERSION == "v5":
    POSTS_API_BASE = f"{config['ghost-configuration']['GHOST_URL']}/api/admin/posts/"
else:
    POSTS_API_BASE = (
        f"{config['ghost-configuration']['GHOST_URL']}/api/{GHOST_VERSION}/admin/posts/"
    )

S3_BASE_URL = config["aws-s3-configuration"]["S3_BASE_URL"]
if S3_BASE_URL == "":
    crediential_not_found(config_path)

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
    if key == "":
        crediential_not_found(config_path)
    id, secret = key.split(":")
    if GHOST_VERSION == "v5":
        aud_value = "/admin/"
    else:
        aud_value = f"/{GHOST_VERSION}/admin/"
    iat = int(date.now().timestamp())

    h = {"iat": iat, "exp": iat + 5 * 60, "aud": aud_value}
    token = jwt.encode(h, bytes.fromhex(secret), algorithm="HS256", headers={"kid": id})
    return token


def get_post_id(slug, headers):
    url = f"{POSTS_API_BASE}slug/{slug}/"
    r = requests.get(url, headers=headers)
    if r.ok:
        j = r.json()
        return (j["posts"][0]["id"], j["posts"][0]["updated_at"])

    else:
        return (None, None)


def make_request(token, body, slug):
    headers = {"Authorization": "Ghost {}".format(token)}

    pid, updated_at = get_post_id(slug, headers)

    if not pid:
        url = f"{POSTS_API_BASE}?source=html"
        r = requests.post(url, json=body, headers=headers)
        # print(r.json())
        print("Created new post")
    else:
        body["posts"][0]["updated_at"] = updated_at

        url = f"{POSTS_API_BASE}{pid}?source=html"

        r = requests.put(url, json=body, headers=headers)
        # print(r.json())
        print("Updated existing post based on slug")

    return


def upload_images():
    local_to_s3 = {}
    for i in mdlib.images:
        if i.startswith("http://") or i.startswith("https://"):
            iext = i.split(".")[-1]
            tp = "/tmp/img." + iext
            response = requests.get(i, stream=True)
            with open(tp, "wb") as out_file:
                shutil.copyfileobj(response.raw, out_file)
            del response
            s3fname_base = sha256sum(tp)
        else:
            s3fname_base = sha256sum(i)
        _, file_extension = os.path.splitext(i)
        s3name = s3fname_base + file_extension
        if i.startswith("http://") or i.startswith("https://"):
            upload_to_s3(tp, s3name)
        else:
            upload_to_s3(i, s3name)
        local_to_s3[i] = f"{S3_BASE_URL}{s3name}"
    print("Uploaded images to s3")
    return local_to_s3


def replace_image_links(post, img_map):
    soup = BeautifulSoup(post["html"], features="html.parser")
    for img in soup.find_all("img"):
        img["src"] = img_map[img["src"]]
    result = str(soup)
    post["html"] = result


def upload_feature_image(meta):
    try:
        i = meta["feature_image"]
        _, file_extension = os.path.splitext(i)
        s3fname_base = sha256sum(i)
        s3name = s3fname_base + file_extension
        upload_to_s3(i, s3name)
        meta["feature_image"] = f"{S3_BASE_URL}{s3name}"
        print("Uploaded feature image")
    except Exception as e:
        pass


def post_to_ghost(meta, md):
    if not "slug" in meta:
        print(
            "ERROR: Include a URL friendly slug field in your markdown file and retry! This is required to support updates"
        )
        return

    if meta["sidebar_toc"]:
        meta["codeinjection_head"] = style + sidebar_toc_head
        meta["codeinjection_foot"] = sidebar_toc_footer
    else:
        meta["codeinjection_head"] = style
    meta["html"] = to_html(md)
    upload_feature_image(meta)
    local_to_s3 = upload_images()
    replace_image_links(meta, local_to_s3)
    post_obj = meta
    token = get_jwt()
    body = {"posts": [post_obj]}
    return make_request(token, body, meta["slug"])


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print("Usage: glee.py <path_to_markdown_file>")
        # print("Trying to read 'sample_post.md'")
        # post = frontmatter.load("sample_post.md")
        sys.exit(0)
    elif len(sys.argv) == 2:
        post = frontmatter.load(sys.argv[1])
    else:
        print("Usage: glee.py <path_to_markdown_file>")
        sys.exit(0)
    # print(post.metadata)
    # print(post.content)
    post_to_ghost(post.metadata, post.content)
