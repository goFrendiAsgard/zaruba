echo "${_BOLD}${_YELLOW}Start preparation${_NORMAL}"
if [ "${_NO_VENV}" = "1" ] && [ ! -d "./venv" ]
then
    echo "${_BOLD}${_YELLOW}Creating venv${_NORMAL}"
    python -m venv ./venv
    ./venv/bin/python -m pip install --upgrade pip
fi
if [ -f "requirements.txt" ]
then
    echo "${_BOLD}${_YELLOW}Install dependencies${_NORMAL}"
    pip install -r requirements.txt
fi
echo "prepare command"
echo "${_BOLD}${_YELLOW}Preparation completed${_NORMAL}"