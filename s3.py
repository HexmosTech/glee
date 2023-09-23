import boto3
from handle_config import view_toml_file
import sys
import mimetypes


def get_mime_type(file_path):
    mime_type, _ = mimetypes.guess_type(file_path)
    return mime_type


def upload_to_s3(local_filepath, s3_filepath, logging):
    config, config_path = view_toml_file(logging)

    ACCESS_KEY_ID = config["aws-s3-configuration"]["ACCESS_KEY_ID"]
    SECRET_ACCESS_KEY = config["aws-s3-configuration"]["SECRET_ACCESS_KEY"]
    BUCKET_NAME = config["aws-s3-configuration"]["BUCKET_NAME"]
    S3_BASE_URL = config["aws-s3-configuration"]["S3_BASE_URL"]
    image_link = f"{S3_BASE_URL}{s3_filepath}"
    mime_type = get_mime_type(local_filepath)

    try:
        session = boto3.Session(
            aws_access_key_id=ACCESS_KEY_ID, aws_secret_access_key=SECRET_ACCESS_KEY
        )
        s3 = session.resource("s3")

        s3.Object(BUCKET_NAME, s3_filepath).put(
            Body=open(local_filepath, "rb"), ContentType=mime_type
        )

        image_link = f"{S3_BASE_URL}{s3_filepath}"
        print(f"Uploaded {local_filepath} to {image_link}")
        return image_link
    except Exception as e:
        sys.exit(f"Error: {e}")
