if [ -f "./Pipfile" ]
then
    echo "${_BOLD}${_YELLOW}Test in pipenv${_NORMAL}"
    pipenv run ztplAppTestCommand
else
    if [ -d "./venv" ]
    then
        echo "${_BOLD}${_YELLOW}Activate venv${_NORMAL}"
        source ./venv/bin/activate
    fi
    echo "${_BOLD}${_YELLOW}Test${_NORMAL}"
    ztplAppTestCommand
fi
echo "${_BOLD}${_YELLOW}Test complete${_NORMAL}"
