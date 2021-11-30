if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
then
    echo "${_RED}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.${_NORMAL}"
    exit 1
fi