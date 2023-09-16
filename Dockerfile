FROM arm64v8/python:3.10-slim

RUN mkdir /app 
WORKDIR /app


# Install essential packages, including CA certificates
RUN apt-get update && \
    apt-get install -y python3 python3-pip curl ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -sSL https://install.python-poetry.org | python3 -
ENV PATH="${PATH}:/root/.local/bin"

# Install additional packages
RUN apt-get update -y \
    && apt-get install --no-install-recommends -y \
        patchelf \
        ccache \
        clang \
        libfuse-dev 

# Create a working directory

# Copy the pyproject.toml and poetry.lock to the container
COPY pyproject.toml poetry.lock /app/

RUN poetry config virtualenvs.create false && poetry install --no-interaction --no-ansi

# Copy the rest of the application code
COPY . /app/

#architecture
RUN uname -m 

# Set the entry point to start the application
RUN python3 -m nuitka \
        --onefile \
        --follow-imports  \
        hello.py 
