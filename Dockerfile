FROM debian:bullseye-slim as base

# Install Python 3 and Poetry
RUN apt-get update && apt-get install -y python3 python3-pip poetry

# Create a virtual environment and activate it
ENV POETRY_VIRTUALENVS_IN_PROJECT=true
RUN python3 -m venv /venv

# Install the dependencies defined in pyproject.toml
WORKDIR /app
COPY pyproject.toml poetry.lock ./

RUN poetry install



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
