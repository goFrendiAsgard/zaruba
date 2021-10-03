if [ -z "${HTTP_PORT}" ]
then
    HTTP_PORT=8080
fi
zaruba serve . "${HTTP_PORT}"