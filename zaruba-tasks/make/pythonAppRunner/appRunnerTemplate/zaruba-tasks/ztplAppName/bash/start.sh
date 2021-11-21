if [ -f "./Pipfile" ]
then
    pipenv run ztplAppStartCommand
else
    if [ -d "./venv" ]
    then
        source ./venv/bin/activate
    fi
    ztplAppStartCommand
fi