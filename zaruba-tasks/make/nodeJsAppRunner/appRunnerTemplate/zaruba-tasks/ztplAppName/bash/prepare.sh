if [ "${SHOULD_REMOVE_NODE_MODULES}" = 1 ]
then
    npm -Rf node_modules
fi
if [ "${SHOULD_CLEAN_NPM_CACHE}" = 1 ]
then
    npm cache clean --force
fi
if [ ! -d "node_modules" ]
then
    npm install
fi
if [ "${SHOULD_REBUILD_NPM}" = 1 ]
then
    npm rebuild
fi
if [ "${SHOULD_INSTALL_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ] || [ "$(isCommandExist tsc)" = 1 ]
    then
        echo "Typescript is already installed"
    else
        npm install -g typescript${TS_VERSION}
    fi
fi
if [ "${SHOULD_COMPILE_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ]
    then
        ./node_modules/.bin/tsc
    else
        tsc
    fi
fi
