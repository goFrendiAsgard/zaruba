if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    _ZRB_APP_TEST_COMMAND='go test -v ./... --race -count=1'
fi