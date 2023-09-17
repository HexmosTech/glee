#!/bin/bash


CONFIG_URL="https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml"

DEST_DIR=""

if [ "$(uname)" == "Darwin" ]; then
    DEST_DIR="/usr/local/bin/glee"
    GLEE_URL="https://github.com/HexmosTech/glee/releases/latest/download/glee_mac.bin"
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    DEST_DIR="/usr/bin/glee"
    GLEE_URL="https://github.com/HexmosTech/glee/releases/latest/download/glee_linux.bin"
else
    echo "Unsupported operating system. Please install glee manually from"
    echo "https://github.com/HexmosTech/glee/releases/latest"
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


echo "Downloading configuration file..."
wget -O "$HOME/.glee.toml" $CONFIG_URL

echo "Installation completed successfully!"
echo "Add the Ghost and S3 Configuration in $HOME/.glee.toml file"
