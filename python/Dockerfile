FROM python:3.10-slim-buster as build

WORKDIR /app

# Copy entrypoint and make executable
COPY ./docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh

# Copy manifest
COPY ./requirements.txt ./

# Install dependencies
RUN python -m pip install -r requirements.txt

# Copy source code
COPY ./*.py ./

# Run as non-root
RUN groupadd -r python-users && useradd -r -g python-users python-user
RUN chown python-user:python-users -R /app
USER python-user

ENTRYPOINT ["/app/docker-entrypoint.sh"]
