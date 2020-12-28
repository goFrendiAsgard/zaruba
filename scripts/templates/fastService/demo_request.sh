#!/bin/sh

# USAGE:
#   /bin/sh demo_request.sh

EMAIL="${1}"
if [ ! -e "${EMAIL}" ]
then
    EMAIL="user$(shuf -i 0-10000 -n 1)@gmail.com"
fi

PASSWORD="${2}"
if [ ! -e "${PASSWORD}" ]
then
    PASSWORD="password$(shuf -i 0-10000 -n 1)"
fi

if [ ! -e "${DEMO_HTTP_PORT}" ]
then
    DEMO_HTTP_PORT=3000
fi

echo "Create new user"
curl --header "Content-Type: application/json" \
    --request POST \
    --data "{\"email\":\"${EMAIL}\",\"password\":\"${PASSWORD}\"}" \
    "http://localhost:${DEMO_HTTP_PORT}/users/"

echo "Get first user"
curl "http://localhost:${DEMO_HTTP_PORT}/users/1"

echo "Get all users"
curl "http://localhost:${DEMO_HTTP_PORT}/users/"