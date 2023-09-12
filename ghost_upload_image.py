import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder
from bs4 import BeautifulSoup
import json
import os
import toml
import sys
from handle_config import get_toml_file, ghost_crediential_not_found


def upload_to_ghost(token, image, hash_name, blog_image_list):
    try:
        config_path = os.path.join(os.path.expanduser("~"), ".glee.toml")
        config = toml.load(config_path)
    except:
        get_toml_file(config_path)
        sys.exit(0)

    GHOST_VERSION = config["ghost-configuration"]["GHOST_VERSION"]

    if GHOST_VERSION == "":
        ghost_crediential_not_found(config_path)

    if GHOST_VERSION == "v5":
        POSTS_API_BASE = (
            f"{config['ghost-configuration']['GHOST_URL']}/api/admin/images/upload/"
        )
    else:
        POSTS_API_BASE = f"{config['ghost-configuration']['GHOST_URL']}/api/{GHOST_VERSION}/admin/images/upload/"

    headers = {"Authorization": "Ghost {}".format(token)}

    try:
        for name in blog_image_list:
            image_name = name.split("/")[-1]
            hash_value = hash_name.split(".")[0]
            if hash_value in image_name:
                print(f"The image already exists and is being reused.")
                return name

        mulit_encoder = MultipartEncoder(
            fields={
                "file": (hash_name, open(image, "rb"), "image/png"),
                "ref": hash_name,
            }
        )
        boundary_value = mulit_encoder.boundary_value
        response = {}
        response = requests.post(
            POSTS_API_BASE,
            headers={
                "Authorization": headers["Authorization"],
                "Content-Type": f"multipart/form-data; boundary={boundary_value}",
            },
            data=mulit_encoder,
        )
        result = response.json()

        image_link = result["images"][0]["url"]
        print("The image has been uploaded into Ghost.")
        return image_link
    except Exception as e:
        pass


def get_images_from_post(post):
    parsed_post = json.loads(post)
    image_list = []
    soup = BeautifulSoup(parsed_post["cards"][0][1]["html"], features="html.parser")
    for img in soup.find_all("img"):
        image_list.append(img["src"])
    return image_list
