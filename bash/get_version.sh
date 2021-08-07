. ${ZARUBA_HOME}/bash/util.sh

get_version() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    _GIT_VERSION=$(git describe --tags --always 2>/dev/null)
    if [ ! -z "${_GIT_VERSION}" ]
    then
        echo ${_GIT_VERSION}
    elif [ -f "./.version" ]
    then
        cat ./.version
    else
        echo "dev"
    fi
    set "${_OLD_STATE}"
}