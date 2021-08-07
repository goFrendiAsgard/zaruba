# load zaruba
PATH="${PATH}:${HOME}/.zaruba"

# load local's bin
if [ -d "${HOME}/.local/bin" ]
then
    PATH="${PATH}:${HOME}/.local/bin"
fi