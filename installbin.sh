#!/bin/bash

poetry install
poetry run python -m nuitka glee.py
if [ "$(uname)" == "Darwin" ]; then
    sudo mv glee.bin /usr/local/bin/glee
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    sudo mv glee.bin /usr/bin/glee
fi
