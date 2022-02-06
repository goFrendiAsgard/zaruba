set -e
echo "Validate app crud fields"

if [ "$("${ZARUBA_BIN}" list validate "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
then
    echo "${_RED}Invalid _ZRB_APP_CRUD_FIELDS: ${_ZRB_APP_CRUD_FIELDS}${_NORMAL}"
    exit 1
fi

echo "Done validating app crud fields"
