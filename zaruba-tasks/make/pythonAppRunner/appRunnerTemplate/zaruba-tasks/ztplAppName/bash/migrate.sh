if [ -f "./Pipfile" ]
then
    pipenv run ztplAppMigrateCommand
else
    if [ -d "./venv" ]
    then
        source ./venv/bin/activate
    fi
    ztplAppMigrateCommand
fi