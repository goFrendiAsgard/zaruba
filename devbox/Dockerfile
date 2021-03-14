FROM ubuntu:20.04

# Prepare environments
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

# Install ubuntu packages
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y git golang --fix-missing

# Clean up package
RUN apt-get clean