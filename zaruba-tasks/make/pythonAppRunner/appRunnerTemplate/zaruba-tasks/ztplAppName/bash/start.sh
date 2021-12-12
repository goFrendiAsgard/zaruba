if [ -f "./Pipfile" ]
then
    echo "${_BOLD}${_YELLOW}Start in pipenv${_NORMAL}"
    pipenv run ztplAppStartCommand
else
    if [ -d "./venv" ]
    then
        echo "${_BOLD}${_YELLOW}Activate venv${_NORMAL}"
        source ./venv/bin/activate
    fi
    echo "${_BOLD}${_YELLOW}Start${_NORMAL}"
    ztplAppStartCommand
fi