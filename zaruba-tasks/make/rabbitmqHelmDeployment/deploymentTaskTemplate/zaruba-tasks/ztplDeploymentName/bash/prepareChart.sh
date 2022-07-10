if [ -d "./chart" ]
then
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Prepare chart dependencies.${_NORMAL}"
    cd ./chart
    helm dependency upgrade
    cd ..
fi
echo "${_YELLOW}${_CONSTRUCTION_ICON} Chart prepared.${_NORMAL}"