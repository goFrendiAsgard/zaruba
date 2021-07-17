echo "📡 Update respository."
apt-get update
apt-get upgrade -y
echo "🔽 Install packages."
apt-get install -y build-essential python3-distutils libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev python-openssl git ncat make tmux zsh neovim cowsay figlet lolcat
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
if [ "${DOCKER_INSTALLED}" = 0 ]
then
    echo "🐳 Installing docker."
    apt-get install -y docker.io
fi
echo "🎉🎉🎉"
echo "Setup complete."
