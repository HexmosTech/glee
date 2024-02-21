import requests
import sys
import os
import toml


def get_toml_file(config_path, logging):
    logging.error(f"The configuration file at {config_path} was not found.")
    config_response = input(
        "Would you like me to create the configuration file? (yes/no): "
    )
    if config_response == "yes" or config_response == "y":
        url = "https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml"
        response = requests.get(url)
        if response.status_code == 200:
            with open(config_path, "wb") as file:
                file.write(response.content)
            logging.info(f"Created the configuration file in {config_path}")
            msg = (
                f"Include the Ghost configurations in the file located at {config_path}"
            )
            logging.info(msg)

        else:
            logging.error(f"Failed to create the configuration file")


def ghost_crediential_not_found(config_path, logging):
    msg = f"Include the Ghost configurations in the file located at {config_path}"
    logging.info(msg)
    sys.exit(0)


def view_toml_file(logging):
    try:
        config_path = os.path.join(os.path.expanduser("~"), ".glee.toml")
        config = toml.load(config_path)
        return config, config_path
    except:
        get_toml_file(config_path, logging)
        sys.exit(0)


def check_configurations_exist(logging):
    try:
        config, config_path = view_toml_file(logging)
        ghost_config = config["ghost-configuration"]
        image_config = config["image-configuration"]
        aws_config = config["aws-s3-configuration"]

        if ghost_config["ADMIN_API_KEY"] == "":
            sys.exit(
                f"Error: Include the Ghost ADMIN_API_KEY in the file located at {config_path}"
            )

        if ghost_config["GHOST_URL"] == "":
            sys.exit(
                f"Error: Include the GHOST_URL in the file located at {config_path}"
            )

        if ghost_config["GHOST_VERSION"] == "":
            sys.exit(
                f"Error: Include the GHOST_VERSION in the file located at {config_path}"
            )

        if image_config["IMAGE_BACKEND"] not in ["ghost", "s3"]:
            sys.exit(
                f"Error: Inproper option for IMAGE_BACKEND , the option should be 'ghost' or 's3' in the file located at {config_path}"
            )
        if image_config["IMAGE_BACKEND"] == "s3":
            if aws_config["ACCESS_KEY_ID"] == "":
                sys.exit(
                    f"Error: Include AWS S3 ACCESS_KEY_ID in the file located at {config_path}"
                )
            if aws_config["SECRET_ACCESS_KEY"] == "":
                sys.exit(
                    f"Error: Include AWS S3 SECRET_ACCESS_KEY in the file located at {config_path}"
                )
            if aws_config["BUCKET_NAME"] == "":
                sys.exit(
                    f"Error: Include AWS S3 BUCKET_NAME in the file located at {config_path}"
                )
            if aws_config["S3_BASE_URL"] == "":
                sys.exit(
                    f"Error: Include AWS S3 S3_BASE_URL in the file located at {config_path}"
                )

    except Exception as e:
        sys.exit(f"An error occurred: {e}")


def print_configuration(logging):
    try:
        config, config_path = view_toml_file(logging)
        print(f"configuration path: {config_path}")
        print("-------------------")
        print(toml.dumps(config))
        sys.exit(0)
    except:
        sys.exit(0)
