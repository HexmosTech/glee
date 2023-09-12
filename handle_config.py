import requests
import sys


def get_toml_file(config_path):
    print(f"The configuration file at {config_path} was not found.")
    config_response = input(
        "Would you like me to create the configuration file? (yes/no): "
    )
    if config_response == "yes" or config_response == "y":
        url = "https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml"
        response = requests.get(url)
        if response.status_code == 200:
            with open(config_path, "wb") as file:
                file.write(response.content)
            print(f"Created the configuration file in {config_path}")
            msg = f"Include the Ghost and AWS S3 configurations in the file located at {config_path}"
            print(msg)

        else:
            print(f"Failed to create the configuration file")


def ghost_crediential_not_found(config_path):
    msg = f"Include the Ghost configurations in the file located at {config_path}"
    print(msg)
    sys.exit(0)
