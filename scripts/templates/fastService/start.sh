if [ -z "${ZARUBA_ENV_PREFIX_HTTP_PORT}" ]
then
    ZARUBA_ENV_PREFIX_HTTP_PORT=3000
fi
pipenv run uvicorn main:app --host=0.0.0.0 --port=${ZARUBA_ENV_PREFIX_HTTP_PORT}