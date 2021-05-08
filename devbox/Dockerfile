FROM ubuntu:20.04

# Prepare environments
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

# Install ubuntu packages
RUN apt-get update --fix-missing
RUN apt-get upgrade -y

# Install common tools
RUN apt-get install -y git golang ncat curl wget inetutils-tools tmux neovim zsh

# install toys
RUN apt-get install -y cowsay figlet lolcat

# Install pyenv dependencies
RUN apt-get install -y make build-essential libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev xz-utils tk-dev libxml2-dev libxmlsec1-dev libffi-dev liblzma-dev

# install pyenv
RUN curl -L https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | sh
RUN pyenv install 3.8.6
RUN pyenv global 3.8.6
RUN pip install pipenv

# install nvm
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | sh
RUN npm install -g node-gyp typescript

RUN apt-get clean