set -e
echo "Validate app directory"

if [ -z "${_ZRB_APP_DIRECTORY}" ]
then
    echo "${_RED}Invalid _ZRB_APP_DIRECTORY: ${_ZRB_APP_DIRECTORY}${_NORMAL}"
    exit 1
fi

echo "Done validating app directory"