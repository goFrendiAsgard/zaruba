if [ -z "${SERVICE_HTTP_PORT}" ]
then
    SERVICE_HTTP_PORT=3000
fi
pipenv run uvicorn main:app --reload --port=${SERVICE_HTTP_PORT}