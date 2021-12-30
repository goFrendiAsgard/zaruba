echo "Preparing prepare command"

# prepare command
if [ -z "${_ZRB_APP_PREPARE_COMMAND}" ]
then
    _ZRB_APP_PREPARE_COMMAND="echo \"prepare ${_ZRB_APP_NAME}\""
fi

echo "Prepare command prepared"