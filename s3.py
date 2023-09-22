import boto3


from handle_config import view_toml_file
import sys


def upload_to_s3(local_filepath, s3_filepath, logging):
    config, config_path = view_toml_file(logging)

    ACCESS_KEY_ID = config["aws-s3-configuration"]["ACCESS_KEY_ID"]
    SECRET_ACCESS_KEY = config["aws-s3-configuration"]["SECRET_ACCESS_KEY"]
    BUCKET_NAME = config["aws-s3-configuration"]["BUCKET_NAME"]
    S3_BASE_URL = config["aws-s3-configuration"]["S3_BASE_URL"]
    image_link = f"{S3_BASE_URL}{s3_filepath}"
    try:
        session = boto3.Session(
            aws_access_key_id=ACCESS_KEY_ID, aws_secret_access_key=SECRET_ACCESS_KEY
        )
        s3 = session.resource("s3")
        bucket = s3.Bucket(BUCKET_NAME)

        result = bucket.upload_file(local_filepath, s3_filepath)

        bucket.Object(s3_filepath).put(ContentType="image/jpeg")

        image_link = f"{S3_BASE_URL}{s3_filepath}"
        return image_link
    except Exception as e:
        sys.exit(f"Error: {e}")
