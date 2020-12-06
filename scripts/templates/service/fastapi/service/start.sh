if [ -z "${HTTP_PORT}" ]
then
    HTTP_PORT=3000
fi
pipenv run uvicorn main:app --reload --port=${HTTP_PORT}