if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    _ZRB_APP_TEST_COMMAND='pytest -rP -v --cov="$(pwd)" --cov-report html'
fi

if [ -z "${_ZRB_APP_PREPARE_COMMAND}" ]
then
    _ZRB_APP_PREPARE_COMMAND='echo "prepare command"'
fi


if [ -z "${_ZRB_APP_START_COMMAND}" ]
then
    if [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/start.sh" ]
    then
        _ZRB_APP_START_COMMAND="./start.sh"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/main.py" ]
    then
        _ZRB_APP_START_COMMAND="python main.py"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/manage.py" ]
    then
        _ZRB_APP_START_COMMAND="python manage.py"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/index.py" ]
    then
        _ZRB_APP_START_COMMAND="python index.py"
    fi
fi
