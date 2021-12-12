if [ -f "./Pipfile" ]
then
    echo "${_BOLD}${_YELLOW}Migrate using pipenv${_NORMAL}"
    pipenv run ztplAppMigrateCommand
else
    if [ -d "./venv" ]
    then
        echo "${_BOLD}${_YELLOW}Activate venv${_NORMAL}"
        source ./venv/bin/activate
    fi
    echo "${_BOLD}${_YELLOW}Migrate${_NORMAL}"
    ztplAppMigrateCommand
fi
echo "${_BOLD}${_YELLOW}Migration complete${_NORMAL}"