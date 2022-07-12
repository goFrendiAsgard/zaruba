if [ -f "./Pipfile" ]
then
    echo "${_BOLD}${_YELLOW}Run in pipenv${_NORMAL}"
    pipenv run . $@
elif [ -d "./venv" ]
then
    echo "${_BOLD}${_YELLOW}Activate venv${_NORMAL}"
    source ./venv/bin/activate
    echo "${_BOLD}${_YELLOW}Run in venv${_NORMAL}"
    . $@
else
    _NO_VENV=1
    . $@
fi