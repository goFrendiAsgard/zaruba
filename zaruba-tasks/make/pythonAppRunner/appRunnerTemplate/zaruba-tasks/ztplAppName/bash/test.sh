if [ -f "./Pipfile" ]
then
    pipenv run ztplAppTestCommand
else
    if [ -d "./venv" ]
    then
        source ./venv/bin/activate
    fi
    ztplAppTestCommand
fi