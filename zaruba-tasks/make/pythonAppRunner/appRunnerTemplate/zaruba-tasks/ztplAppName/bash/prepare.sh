if [ -f "./Pipfile" ]
then
    pipenv install
    pipenv run ztplAppPrepareCommand
else
    if [ ! -d "./venv" ]
    then
        python -m venv ./venv
    fi
    source ./venv/bin/activate
    if [ -f "requirements.txt" ]
    then
        pip install -r requirements.txt
    fi
    ztplAppPrepareCommand
fi