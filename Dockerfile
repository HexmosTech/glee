FROM debian:bullseye-slim as base

# Install essential packages
RUN apt-get update && \
    apt-get install -y python3 python3-pip curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install Poetry
RUN curl -sSL https://install.python-poetry.org | python3 -

ENV PATH="${PATH}:/root/.poetry/bin"

# Set the working directory
WORKDIR /app

# Copy the pyproject.toml and poetry.lock to the container
COPY pyproject.toml poetry.lock /app/

# Install dependencies using Poetry
RUN poetry config virtualenvs.create true && \
    poetry install --no-interaction --no-ansi

# Copy the rest of the application code
COPY . /app/

# Set the entry point to start the application
RUN python3 -m nuitka \
        --onefile \
        --follow-imports  \
        --include-package=pygments  \
        glee.py 
