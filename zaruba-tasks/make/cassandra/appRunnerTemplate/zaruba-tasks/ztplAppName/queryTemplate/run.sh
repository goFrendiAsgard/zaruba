set -e
WORKING_DIRECTORY="$(realpath "$(dirname "${0}")")"
echo "Working Directory: ${WORKING_DIRECTORY}"
echo "SQL: $(cat "${WORKING_DIRECTORY}/query.sql")"
cqlsh -u "${CASSANDRA_USER}" -p "${CASSANDRA_PASSWORD}" -f "${WORKING_DIRECTORY}/query.sql"