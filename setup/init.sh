set -e

SETUP_PATH="$(dirname $(realpath "${0}"))"
PROMPT='
0. Exit
1. Setup ubuntu
2. Setup tmux
3. Install kubectl
4. Install helm
5. Install nvm
6. Install pyenv
7. Install sdkman
8. Install gvm

💀 Type your choice, carefully:
'

until [ "${CHOICE}" = "0" ]
do
    echo "${PROMPT}"
    read CHOICE

    case "${CHOICE}" in
    "1")
        . "${SETUP_PATH}/setup_ubuntu.sh"
        ;;
    "2")
        . "${SETUP_PATH}/setup_tmux.sh"
        ;;
    "3")
        . "${SETUP_PATH}/install_kubectl.sh"
        ;;
    "4")
        . "${SETUP_PATH}/install_helm.sh"
        ;;
    "5")
        . "${SETUP_PATH}/install_nvm.sh"
        ;;
    "6")
        . "${SETUP_PATH}/install_pyenv.sh"
        ;;
    "7")
        . "${SETUP_PATH}/install_sdkman.sh"
        ;;
    "8")
        . "${SETUP_PATH}/install_gvm.sh"
        ;;
    esac
done
