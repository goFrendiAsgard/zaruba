#!/bin/sh

# USAGE:
#   /bin/sh demo_start.sh

if [ -z "${DEMO_HTTP_PORT}" ]
then
    DEMO_HTTP_PORT=3000
fi

rm -f database.db
pipenv run uvicorn demo_main:app --reload --port=${DEMO_HTTP_PORT}