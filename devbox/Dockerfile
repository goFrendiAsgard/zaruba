FROM ubuntu:20.04

# Prepare environments
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

# Install ubuntu packages
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y ncat make wget curl git golang python3 python3-pip python-is-python3 nodejs npm

# Install NVM
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.0/install.sh | bash
RUN NVM_DIR="$HOME/.nvm"
RUN \. "$NVM_DIR/nvm.sh"
RUN nvm install node
RUN npm install -g typescript
RUN npm install -g node-gyp

# Install pipenv
RUN pip3 install pipenv

# Clean up package
RUN apt-get clean