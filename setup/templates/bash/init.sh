ZARUBA_INIT_SCRIPT_LOADED=1

# https://stackoverflow.com/a/66695181
export DOCKER_BUILDKIT=0

# load configurations from .profile
if [ -f "${HOME}/.profile" ]
then
    . "${HOME}/.profile"
fi

# load zaruba
PATH="${PATH}:${HOME}/.zaruba"
alias zrb=zaruba

# load local's bin
if [ -d "${HOME}/.local/bin" ]
then
    PATH="${PATH}:${HOME}/.local/bin"
fi
