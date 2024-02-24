#!/bin/bash


CONFIG_URL="https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml"

DEST_DIR=""

if [ "$(uname)" == "Darwin" ]; then
    DEST_DIR="/usr/local/bin/glee"
    GLEE_URL="https://github.com/HexmosTech/glee/releases/v1.1.12/download/glee_mac.bin"
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    DEST_DIR="/usr/bin/glee"
    GLEE_URL="https://github.com/HexmosTech/glee/releases/v1.1.12/download/glee_linux.bin"
else
    echo "Unsupported operating system. Please install glee manually from"
    echo "https://github.com/HexmosTech/glee/releases/v1.1.12"
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

config_file="$HOME/.glee.toml"

if [ ! -f "$config_file" ]; then
    echo "Downloading configuration file..."
    wget -O "$config_file" "$CONFIG_URL"
    echo "Installation completed successfully!"
    echo "Add the Ghost Configuration in $config_file file for using glee."
else
    echo "Update completed successfully!"
    echo "Reusing the configurations from the $config_file file."
    
fi


