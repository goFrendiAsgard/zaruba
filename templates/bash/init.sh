ZARUBA_INIT_SCRIPT_LOADED=1

# load configurations from .profile
if [ -f "${HOME}/.profile" ]
then
    . "${HOME}/.profile"
fi

# load zaruba
PATH="${PATH}:${HOME}/.zaruba"

# load local's bin
if [ -d "${HOME}/.local/bin" ]
then
    PATH="${PATH}:${HOME}/.local/bin"
fi
