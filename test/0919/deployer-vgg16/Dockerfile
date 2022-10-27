# ghcr.io/freckie/deployer-vgg16:0.4.0

# Use the official lightweight Python image.
# https://hub.docker.com/_/python
FROM ghcr.io/freckie/shmtorch:0.4.0

# Allow statements and log messages to immediately appear in the Knative logs
ENV PYTHONUNBUFFERED True

# Copy local code to the container image.
ENV APP_HOME /app
WORKDIR $APP_HOME
COPY . ./

# Install production dependencies.
RUN pip install torchvision

CMD [ "python", "deploy.py" ]