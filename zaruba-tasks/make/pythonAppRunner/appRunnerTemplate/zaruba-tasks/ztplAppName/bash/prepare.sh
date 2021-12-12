if [ -f "./Pipfile" ]
then
    echo "${_BOLD}${_YELLOW}Install dependencies using pipenv${_NORMAL}"
    pipenv install
    echo "${_BOLD}${_YELLOW}Prepare using pipenv${_NORMAL}"
    pipenv run ztplAppPrepareCommand
else
    if [ ! -d "./venv" ]
    then
        echo "${_BOLD}${_YELLOW}Activate venv${_NORMAL}"
        python -m venv ./venv
    fi
    source ./venv/bin/activate
    if [ -f "requirements.txt" ]
    then
        echo "${_BOLD}${_YELLOW}Install dependencies${_NORMAL}"
        pip install -r requirements.txt
    fi
    echo "${_BOLD}${_YELLOW}Prepare${_NORMAL}"
    ztplAppPrepareCommand
fi
echo "${_BOLD}${_YELLOW}Preparation complete${_NORMAL}"
