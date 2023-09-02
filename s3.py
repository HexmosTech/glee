import boto3
import toml
import os
import sys

# Load the TOML file
from handle_config import get_toml_file,crediential_not_found

try:
    config_path = os.path.join(os.path.expanduser("~"), ".glee.toml")
    config = toml.load(config_path)
except:
    get_toml_file(config_path)
    sys.exit(0)


ACCESS_KEY_ID = config["aws-s3-configuration"]["ACCESS_KEY_ID"]
if ACCESS_KEY_ID == "":
    crediential_not_found(config_path)
SECRET_ACCESS_KEY = config["aws-s3-configuration"]["SECRET_ACCESS_KEY"]
BUCKET = config["aws-s3-configuration"]["BUCKET"]

session = boto3.Session(
    aws_access_key_id=ACCESS_KEY_ID, aws_secret_access_key=SECRET_ACCESS_KEY
)
s3 = session.resource("s3")


def upload_to_s3(local_filepath, s3_filepath):
    return s3.Bucket(BUCKET).upload_file(local_filepath, s3_filepath)
