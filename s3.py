import boto3
import toml

# Load the TOML file
config = toml.load("config.toml")


ACCESS_KEY_ID = config['aws-s3-configuration']['ACCESS_KEY_ID']
SECRET_ACCESS_KEY = config['aws-s3-configuration']['SECRET_ACCESS_KEY']
BUCKET = config['aws-s3-configuration']['BUCKET']

session = boto3.Session(aws_access_key_id=ACCESS_KEY_ID, aws_secret_access_key=SECRET_ACCESS_KEY)
s3 = session.resource('s3')

def upload_to_s3(local_filepath, s3_filepath):
    return s3.Bucket(BUCKET).upload_file(local_filepath, s3_filepath)