echo "Preparing migrate command"

# test command
if [ -z "${_ZRB_APP_MIGRATE_COMMAND}" ]
then
    _ZRB_APP_MIGRATE_COMMAND="echo \"migrate ${_ZRB_APP_NAME}\""
fi

echo "Migrate command prepared"