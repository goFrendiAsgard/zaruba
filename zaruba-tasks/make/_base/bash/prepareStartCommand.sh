set -e
echo "Preparing start command"

# start command
if [ -z "${_ZRB_APP_START_COMMAND}" ]
then
    if [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/start.sh" ]
    then
        _ZRB_APP_START_COMMAND="./start.sh"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/main.go" ]
    then
        _ZRB_APP_START_COMMAND="go run main.go"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/package.json" ]
    then
        _ZRB_APP_START_COMMAND="npm start"
    elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/main.py" ]
    then
        if [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/Pipfile" ]
        then
            _ZRB_APP_START_COMMAND="pipenv run python main.py"
        elif [ -f "${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/venv/bin/activate" ]
        then
            _ZRB_APP_START_COMMAND="source ./venv/bin/activate && python main.py"
        else
            _ZRB_APP_START_COMMAND="python main.py"
        fi
    else
        _ZRB_APP_START_COMMAND="echo \"Replace this with command to start ${_ZRB_APP_NAME}\" && exit 1"
    fi
fi

echo "Start command prepared"