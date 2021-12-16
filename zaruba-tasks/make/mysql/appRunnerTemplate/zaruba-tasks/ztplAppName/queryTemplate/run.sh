set -e
WORKING_DIRECTORY="$(realpath "$(dirname "${0}")")"
echo "Working Directory: ${WORKING_DIRECTORY}"
echo "SQL: $(cat "${WORKING_DIRECTORY}/query.sql")"
mysql --user=root --password="${MYSQL_ROOT_PASSWORD}" < "${WORKING_DIRECTORY}/query.sql"