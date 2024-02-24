#!/bin/bash

go build -o glee.bin   
if [ "$(uname)" == "Darwin" ]; then
    sudo mv glee.bin /usr/local/bin/glee
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    sudo mv glee.bin /usr/bin/glee
fi
