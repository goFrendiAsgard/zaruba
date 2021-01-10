# USAGE
# sh should_be_file.sh <value> <error-message>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ ! -f "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi