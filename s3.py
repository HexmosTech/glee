import boto3


from handle_config import view_toml_file


def upload_to_s3(local_filepath, s3_filepath, logging):
    config, config_path = view_toml_file(logging)

    ACCESS_KEY_ID = config["aws-s3-configuration"]["ACCESS_KEY_ID"]
    # if ACCESS_KEY_ID == "":
    #     ghost_crediential_not_found(config_path)
    SECRET_ACCESS_KEY = config["aws-s3-configuration"]["SECRET_ACCESS_KEY"]
    BUCKET_NAME = config["aws-s3-configuration"]["BUCKET_NAME"]

    session = boto3.Session(
        aws_access_key_id=ACCESS_KEY_ID, aws_secret_access_key=SECRET_ACCESS_KEY
    )
    s3 = session.resource("s3")

    return s3.Bucket(BUCKET_NAME).upload_file(local_filepath, s3_filepath)
