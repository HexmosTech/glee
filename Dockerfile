FROM debian:bullseye-slim as base

RUN mkdir /app 

RUN apt-get update && \
    apt-get install -y python3 python3-pip curl && \
    apt-get clean && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*


WORKDIR /app

RUN apt-get update -y \
    && apt-get install --no-install-recommends -y \
        patchelf \
        ccache \
        clang \
        libfuse-dev 

# Copy the pyproject.toml and poetry.lock to the container
COPY pyproject.toml poetry.lock /app/

ENV PYTHONPATH=${PYTHONPATH}:${PWD} 
RUN pip3 install poetry
RUN poetry config virtualenvs.create false
RUN poetry install --no-dev



# Copy the rest of the application code
COPY . /app/

# Set the entry point to start the application
RUN python3 -m nuitka \
        --onefile \
        --follow-imports  \
        --include-package=pygments  \
        glee.py 
