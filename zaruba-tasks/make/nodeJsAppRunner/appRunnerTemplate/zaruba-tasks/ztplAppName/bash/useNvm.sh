if [ "$(isCommandExist nvm)" = 1 ]
then
    if [ "$(isCommandError nvm ls "${NODE_VERSION}" )" ]
    then
        echo "${_BOLD}${_YELLOW}Install node ${NODE_VERSION}${_NORMAL}"
        nvm install "${NODE_VERSION}"
    else
        echo "${_BOLD}${_YELLOW}Use node ${NODE_VERSION}${_NORMAL}"
        nvm use "${NODE_VERSION}"
    fi
fi