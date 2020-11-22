#!/bin/sh
pipenv run uvicorn main:app --reload --port=${SEED_HTTP_PORT}