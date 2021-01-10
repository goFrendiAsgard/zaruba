# USAGE
# /bin/sh check_setup_ubuntu.sh <home>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ ! -w /var/lib/apt/lists ]
then
    echo "${Bold}${Red}You don't have root privilege${Normal}"
    echo "${Bold}${Red}You might want to re-run the command using root privilege:${Normal}"
    echo "  ${Yellow}sudo -E zaruba please setupUbuntu${Normal}"
    exit 1
fi
if echo "${1}" | grep -q "/root$"
then
    echo "${Bold}${Red}Your HOME is seems to be root's:${Normal}"
    echo "  ${Faint}${1}${Normal}"
    echo "${Bold}${Red}You might want to re-run the command using 'sudo -E' option:${Normal}"
    echo "  ${Yellow}sudo -E zaruba please setupUbuntu${Normal}"
    echo "${Bold}${Red}Do you really want to continue with current home directory? (Y/n)${Normal}"
    read _INPUT
    if [ "${_INPUT}" != "Y" ] && [ "${_INPUT}" != "y" ]
    then
        echo "👏 Wise choice. Setup canceled"
        exit 1
    fi
    echo "👏 Brave choice. Let's continue"
fi