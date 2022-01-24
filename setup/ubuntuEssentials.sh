set -e

echo "📡 Updating respository."
sudo apt-get update
sudo apt-get upgrade -y

echo "🔽 Installing packages."
sudo apt-get install -y build-essential python3-distutils libssl-dev zlib1g-dev libbz2-dev libreadline-dev \
  libsqlite3-dev libpq-dev python3-dev llvm libncurses5-dev libncursesw5-dev xz-utils tk-dev libffi-dev \
  liblzma-dev python-openssl bison libblas-dev liblapack-dev gfortran rustc \
  fd-find ripgrep wget curl git ncat make tmux zsh neovim cowsay figlet lolcat

# Determine whether docker is already installed or not
DOCKER_INSTALLED=1
(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e
docker version >> /dev/null
if [ "$?" = 127 ]
then
    DOCKER_INSTALLED=0
else
    echo "👏 Docker was already installed."
fi
set "${_OLD_STATE}"

# Install docker if it is not intalled
if [ "${DOCKER_INSTALLED}" = 0 ]
then
    echo "🐳 Installing docker."
    sudo apt-get install -y docker.io
fi

echo "🎉🎉🎉"
echo "Setup complete."
