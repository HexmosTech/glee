
#!/bin/bash

GLEE_URL="https://github.com/HexmosTech/glee/releases/latest/download/glee.bin"
CONFIG_URL="https://raw.githubusercontent.com/HexmosTech/glee/main/config.toml"

DEST_DIR=""
CONFIG_DEST_DIR="$HOME/glee"

if [ "$(uname)" == "Darwin" ]; then
    DEST_DIR="/usr/local/bin/glee"
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    DEST_DIR="/usr/bin/glee"
else
    echo "Unsupported operating system. Please install glee manually."
    exit 1
fi

echo "Downloading glee.bin..."
wget -O glee.bin $GLEE_URL

# Check if the download was successful
if [ $? -ne 0 ]; then
    echo "Failed to download glee.bin. Please check your internet connection and try again."
    exit 1
fi

echo "Moving glee.bin to $DEST_DIR..."
sudo mv glee.bin $DEST_DIR

if [ $? -ne 0 ]; then
    echo "Failed to move glee.bin to $DEST_DIR. Please ensure you have sudo privileges."
    exit 1
fi

sudo chmod +x $DEST_DIR

mkdir -p "$CONFIG_DEST_DIR"

echo "Downloading config.toml..."
wget -O "$CONFIG_DEST_DIR/config.toml" $CONFIG_URL

echo "Installation completed successfully!"
echo "Add the Ghost and S3 configuration in glee/config.toml file"
