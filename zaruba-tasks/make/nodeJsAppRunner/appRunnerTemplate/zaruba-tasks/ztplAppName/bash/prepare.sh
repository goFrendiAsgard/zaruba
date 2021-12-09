if [ "${SHOULD_REMOVE_NODE_MODULES}" = 1 ]
then
    echo "${_BOLD}${_YELLOW}Remove node_modules${_NORMAL}"
    rm -Rf node_modules
fi
if [ "${SHOULD_CLEAN_NPM_CACHE}" = 1 ]
then
    echo "${_BOLD}${_YELLOW}Clean npm cache${_NORMAL}"
    npm cache clean --force
fi
if [ ! -d "node_modules" ]
then
    echo "${_BOLD}${_YELLOW}Install npm packages${_NORMAL}"
    npm install --include=dev
fi
if [ "${SHOULD_REBUILD_NPM}" = 1 ]
then
    echo "${_BOLD}${_YELLOW}Rebuild npm packages${_NORMAL}"
    npm rebuild
fi
if [ "${SHOULD_INSTALL_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ] || [ "$(isCommandExist tsc)" = 1 ]
    then
        echo "${_BOLD}${_YELLOW}Typescript is already installed${_NORMAL}"
    else
        echo "${_BOLD}${_YELLOW}Install typescript${_NORMAL}"
        npm install -g typescript${TS_VERSION}
    fi
fi
if [ "${SHOULD_COMPILE_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ]
    then
        echo "${_BOLD}${_YELLOW}Compile using local tsc${_NORMAL}"
        ./node_modules/.bin/tsc
    else
        echo "${_BOLD}${_YELLOW}Compile using global tsc${_NORMAL}"
        tsc
    fi
fi