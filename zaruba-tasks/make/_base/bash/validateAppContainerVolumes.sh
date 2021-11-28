if [ "$("${ZARUBA_BIN}" list validate "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
then
    echo "${_RED}Invalid _ZRB_APP_CONTAINER_VOLUMES: ${_ZRB_APP_CONTAINER_VOLUMES}${_NORMAL}"
    exit 1
fi