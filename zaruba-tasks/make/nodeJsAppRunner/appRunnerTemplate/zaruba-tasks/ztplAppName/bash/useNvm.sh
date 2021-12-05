if [ "$(isCommandExist nvm)" = 1 ]
then
    if [ "$(isCommandError nvm ls "${NODE_VERSION}" )" ]
    then
        nvm install "${NODE_VERSION}"
    else
        nvm use "${NODE_VERSION}"
    fi
fi