# USAGE
# sh should_be_dir.sh <value> <error-message>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ ! -d "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi