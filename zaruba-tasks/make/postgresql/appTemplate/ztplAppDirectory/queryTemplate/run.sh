WORKING_DIRECTORY="$(realpath "$(dirname "${0}")")"
echo "Working Directory: ${WORKING_DIRECTORY}"
psql --user=${POSTGRESQL_USERNAME} -w --file="${WORKING_DIRECTORY}/query.sql"