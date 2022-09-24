if [ -z "${APP_HTTP_PORT}" ]
then
    APP_HTTP_PORT=3000
fi

uvicorn main:app --host=0.0.0.0 --port=${APP_HTTP_PORT} --workers 3