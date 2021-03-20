FROM ubuntu:20.04

# Prepare environments
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV DEBIAN_FRONTEND=noninteractive

# Install ubuntu packages
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y git golang build-essential python3-distutils libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev python-openssl git ncat make tmux zsh neovim cowsay figlet lolcat inetutils-tools

RUN apt-get clean