#!/bin/sh

# USAGE:
#   /bin/sh demo_request.sh

EMAIL="${1}"
if [ -z "${EMAIL}" ]
then
    EMAIL="user$(shuf -i 0-10000 -n 1)@gmail.com"
fi

if [ -z "${DEMO_HTTP_PORT}" ]
then
    DEMO_HTTP_PORT=3000
fi

echo "Create new user"
curl --header "Content-Type: application/json" \
    --request POST \
    --data "{\"email\":\"${EMAIL}\"}" \
    "http://localhost:${DEMO_HTTP_PORT}/users/"
echo ""

echo "Get first user"
curl "http://localhost:${DEMO_HTTP_PORT}/users/1"
echo ""

echo "Update user"
curl --header "Content-Type: application/json" \
    --request PUT \
    --data "{\"email\":\"updated.${EMAIL}\"}" \
    "http://localhost:${DEMO_HTTP_PORT}/users/1"
echo ""

echo "Get all users"
curl "http://localhost:${DEMO_HTTP_PORT}/users/"
echo ""

echo "Delete user"
curl --header "Content-Type: application/json" \
    --request DELETE \
    "http://localhost:${DEMO_HTTP_PORT}/users/1"
echo ""

echo "Get all users"
curl "http://localhost:${DEMO_HTTP_PORT}/users/"
echo ""
