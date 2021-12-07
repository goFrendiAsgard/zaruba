if [ "${SHOULD_REMOVE_NODE_MODULES}" = 1 ]
then
    echo "Remove node_modules"
    rm -Rf node_modules
fi
if [ "${SHOULD_CLEAN_NPM_CACHE}" = 1 ]
then
    echo "Clean npm cache"
    npm cache clean --force
fi
if [ ! -d "node_modules" ]
then
    echo "Install npm packages"
    npm install --include=dev
fi
if [ "${SHOULD_REBUILD_NPM}" = 1 ]
then
    echo "Rebuild npm packages"
    npm rebuild
fi
if [ "${SHOULD_INSTALL_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ] || [ "$(isCommandExist tsc)" = 1 ]
    then
        echo "Typescript is already installed"
    else
        echo "Install typescript"
        npm install -g typescript${TS_VERSION}
    fi
fi
if [ "${SHOULD_COMPILE_TS}" = 1 ]
then
    if [ -f "./node_modules/.bin/tsc" ]
    then
        echo "Compile using local tsc"
        ./node_modules/.bin/tsc
    else
        echo "Compile using global tsc"
        tsc
    fi
fi
