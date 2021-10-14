WORKING_DIRECTORY="$(dirname ${0})"
for MIGRATION_SCRIPT in $(ls "${WORKING_DIRECTORY}/migrations")
do
    mysql --user=root --password="${MYSQL_ROOT_PASSWORD}" < "${WORKING_DIRECTORY}/migrations/${MIGRATION_SCRIPT}"
done