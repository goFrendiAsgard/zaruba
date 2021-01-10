# USAGE
# sh should_exist.sh <value> <error-message>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ ! -e "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi