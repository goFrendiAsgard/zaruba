. ${ZARUBA_HOME}/bash/util.sh

show_version() {
    _GIT_VERSION=$(git describe --tags)
    if [ ! -z "${_GIT_VERSION}" ]
    then
        echo ${_GIT_VERSION}
    elif [ -f "./.version" ]
    then
        cat ./.version
    else
        echo "dev"
    fi
}