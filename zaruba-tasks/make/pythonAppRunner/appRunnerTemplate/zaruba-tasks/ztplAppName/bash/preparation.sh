echo "${_BOLD}${_YELLOW}Start preparation${_NORMAL}"
if [ ! -d "./venv" ]
then
    echo "${_BOLD}${_YELLOW}Creating venv${_NORMAL}"
    python -m venv ./venv
fi
if [ -f "requirements.txt" ]
then
    echo "${_BOLD}${_YELLOW}Install dependencies${_NORMAL}"
    pip install -r requirements.txt
fi
echo "prepare command"
echo "${_BOLD}${_YELLOW}Preparation completed${_NORMAL}"