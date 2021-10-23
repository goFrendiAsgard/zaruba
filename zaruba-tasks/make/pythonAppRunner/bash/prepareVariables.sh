if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    if [ -f "${_ZRB_APP_DIRECTORY}/Pipfile" ]
    then
        _ZRB_APP_TEST_COMMAND='pipenv run pytest -rP -v --cov="$(pwd)" --cov-report html'
    else
        _ZRB_APP_TEST_COMMAND='pytest -rP -v --cov="$(pwd)" --cov-report html'
    fi
fi

if [ -z "${_ZRB_APP_PREPARE_COMMAND}" ]
then
    if [ -f "${_ZRB_APP_DIRECTORY}/Pipfile" ]
    then
        _ZRB_APP_PREPARE_COMMAND='pipenv install'
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
        else
            _ZRB_APP_START_COMMAND="./start.sh"
        fi
    fi
fi
