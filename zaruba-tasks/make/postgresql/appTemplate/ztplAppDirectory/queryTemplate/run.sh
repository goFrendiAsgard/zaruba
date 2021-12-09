set -e
WORKING_DIRECTORY="$(realpath "$(dirname "${0}")")"
echo "Working Directory: ${WORKING_DIRECTORY}"
echo "SQL: $(cat "${WORKING_DIRECTORY}/query.sql")"
PGPASSWORD="${POSTGRESQL_PASSWORD}" psql --username=${POSTGRESQL_USERNAME} --no-password --file="${WORKING_DIRECTORY}/query.sql"