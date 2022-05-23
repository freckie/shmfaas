# Base image for shmfaas function
# ghcr.io/freckie/shmtorch:3.10-slim

# Use the official lightweight Python image.
FROM python:3.10-slim

LABEL maintainer="freckie@frec.kr"

# Install production dependencies
RUN pip install torch==1.11.0 numpy==1.22.4 shmtorch
