# USAGE
# /bin/sh inject_bootstrap.sh <bashrc-path>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ -f "${1}" ]
then
    if cat "${1}" | grep -Fqe "/scripts/bootstrap.sh"
    then
        echo -e "${Faint}Bootstrap script ${ZARUBA_HOME}/scripts/bootstrap.sh is already loaded in ${1}${Normal}"
    else
        echo "" >> "${1}"
        echo "# Load zaruba's bootstrap" >> "${1}"
        echo "if [ -x "${ZARUBA_HOME}/scripts/bootstrap.sh" ]" >> "${1}"
        echo 'then' >> "${1}"
        echo "    . ${ZARUBA_HOME}/scripts/bootstrap.sh" >> "${1}"
        echo 'fi' >> "${1}"
        echo "" >> "${1}"
    fi
fi