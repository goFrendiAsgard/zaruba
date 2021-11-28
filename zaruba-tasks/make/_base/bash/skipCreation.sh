if [ ! -z "${_ZRB_CFG_SKIP_CREATION_PATH}" ]
then
    if [ -x "${_ZRB_CFG_SKIP_CREATION_PATH}" ]
    then
        echo "${_YELLOW}[SKIP] ${_ZRB_CFG_SKIP_CREATION_PATH} already exist.${_NORMAL}"
        exit 0
    fi
fi