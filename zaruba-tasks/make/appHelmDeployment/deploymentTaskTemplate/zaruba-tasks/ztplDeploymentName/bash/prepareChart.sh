if [ -d "./chart" ]
then
    echo "${_YELLOW}🚧 Prepare chart dependencies.${_NORMAL}"
    cd ./chart
    helm dependency upgrade
    cd ..
fi
echo "${_YELLOW}🚧 Chart prepared.${_NORMAL}"