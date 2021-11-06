if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    if [ -f "${_ZRB_APP_DIRECTORY}/Pipfile" ]
    then
        _ZRB_APP_TEST_COMMAND='pipenv run pytest -rP -v --cov="$(pwd)" --cov-report html'
    elif [ -f "${_ZRB_APP_DIRECTORY}/venv/bin/activate" ]
    then
        _ZRB_APP_TEST_COMMAND='source ./venv/bin/activate && pytest -rP -v --cov="$(pwd)" --cov-report html'
    else
        _ZRB_APP_TEST_COMMAND='pytest -rP -v --cov="$(pwd)" --cov-report html'
    fi
fi

if [ -z "${_ZRB_APP_PREPARE_COMMAND}" ]
then
    if [ -f "${_ZRB_APP_DIRECTORY}/Pipfile" ]
    then
        _ZRB_APP_PREPARE_COMMAND='pipenv install'
    elif [ -f "${_ZRB_APP_DIRECTORY}/venv/bin/activate" ]
    then
        _ZRB_APP_PREPARE_COMMAND='source ./venv/bin/activate && pip install -r requirements.txt'
    elif [ -f "requirements.txt" ]
    then
        _ZRB_APP_PREPARE_COMMAND='pip install -r requirements.txt'
    else
        _ZRB_APP_PREPARE_COMMAND='echo "prepare command"'
    fi
fi


if [ -z "${_ZRB_APP_START_COMMAND}" ]
then
    if [ -f "${_ZRB_APP_DIRECTORY}/start.sh" ]
    then
        if [ -f "${_ZRB_APP_DIRECTORY}/Pipfile" ]
        then
            _ZRB_APP_START_COMMAND="pipenv run bash ./start.sh"
        elif [ -f "${_ZRB_APP_DIRECTORY}/venv/bin/activate" ]
        then
            _ZRB_APP_START_COMMAND="source ./venv/bin/activate && ./start.sh"
        else
            _ZRB_APP_START_COMMAND="./start.sh"
        fi
    fi
fi
