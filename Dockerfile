FROM debian:bullseye-slim as base

# Install essential packages, including CA certificates
RUN apt-get update && \
    apt-get install -y python3 python3-pip curl ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install additional packages
RUN apt-get update -y \
    && apt-get install --no-install-recommends -y \
        patchelf \
        ccache \
        clang \
        libfuse-dev 

# Create a working directory
RUN mkdir /app 
WORKDIR /app

# Copy the pyproject.toml and poetry.lock to the container
COPY pyproject.toml poetry.lock /app/

# Set the Python path and install Poetry
ENV PYTHONPATH=${PYTHONPATH}:${PWD} 
RUN pip3 install poetry

# Configure Poetry and install dependencies
RUN poetry config virtualenvs.create false
RUN poetry install 

# Copy the rest of the application code
COPY . /app/

# Set the entry point to start the application
RUN python3 -m nuitka \
        --onefile \
        --follow-imports  \
        --include-package=pygments  \
        glee.py 
