if [ -z "${_ZRB_APP_TEST_COMMAND}" ]
then
    _ZRB_APP_TEST_COMMAND='pipenv run pytest -rP -v --cov="$(pwd)" --cov-report html'
fi