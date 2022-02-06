set -e
echo "Preparing test command"

# test command
if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    _ZRB_APP_TEST_COMMAND="echo \"test ${_ZRB_APP_NAME}\""
fi

echo "Test command prepared"