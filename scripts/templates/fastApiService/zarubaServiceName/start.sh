if [ -z "${ZARUBA_SERVICE_NAME_HTTP_PORT}" ]
then
    ZARUBA_SERVICE_NAME_HTTP_PORT=3000
fi

if [ -n "${ZARUBA_SERVICE_NAME_SSL_CERT_FILE}" ] && [ -n "$ZARUBA_SERVICE_NAME_SSL_KEY_FILE" ]
then
    pipenv run uvicorn main:app --host=0.0.0.0 --ssl-keyfile "${ZARUBA_SERVICE_NAME_SSL_KEY_FILE}" --ssl-certfile "${ZARUBA_SERVICE_NAME_SSL_CERT_FILE}" --port=${ZARUBA_SERVICE_NAME_HTTP_PORT}
else
    pipenv run uvicorn main:app --host=0.0.0.0 --port=${ZARUBA_SERVICE_NAME_HTTP_PORT}
fi