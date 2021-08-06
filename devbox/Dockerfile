FROM ubuntu:20.04

# Prepare environments
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

# Install ubuntu packages
RUN apt-get update --fix-missing
RUN apt-get upgrade -y

# Install common tools
RUN apt-get install -y git ncat curl wget inetutils-tools docker.io

RUN apt-get clean