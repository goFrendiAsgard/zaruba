if [ "$("${ZARUBA_BIN}" list validate "${_ZRB_APP_PORTS}")" = 0 ]
then
    echo "${_RED}Invalid _ZRB_APP_PORTS: ${_ZRB_APP_PORTS}${_NORMAL}"
    exit 1
fi