FROM python:3.7
RUN mkdir /app 

WORKDIR /app


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
