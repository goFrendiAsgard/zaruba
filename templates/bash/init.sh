for FILE in "${HOME}/.bash_profile" "${HOME}/.bash_login" "${HOME}/.profile" "${HOME}/.bashrc"
do
    if [ -f "${FILE}" ]
    then
        . "${FILE}"
    fi
done

# load zaruba
PATH="${PATH}:${HOME}/.zaruba"

# load local's bin
if [ -d "${HOME}/.local/bin" ]
then
    PATH="${PATH}:${HOME}/.local/bin"
fi