if [ ! -d "./venv" ]
then
    echo "${_YELLOW}🚧 Create virtual environment.${_NORMAL}"
    python -m venv ./venv
fi
source ./venv/bin/activate
echo "${_YELLOW}🚧 Install pip packages.${_NORMAL}"
pip install -r requirements.txt