echo "Preparing check command"

# check command
if [ -z "${_ZRB_APP_CHECK_COMMAND}" ]
then
    _ZRB_APP_CHECK_COMMAND="echo \"check ${_ZRB_APP_NAME}\""
fi

echo "Check command prepared"