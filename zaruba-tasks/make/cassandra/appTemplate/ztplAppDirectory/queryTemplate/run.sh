WORKING_DIRECTORY="$(realpath "$(dirname "${0}")")"
echo "Working Directory: ${WORKING_DIRECTORY}"
cqlsh -u "${CASSANDRA_USER}" -p "${CASSANDRA_PASSWORD}" -f "${WORKING_DIRECTORY}/query.sql"