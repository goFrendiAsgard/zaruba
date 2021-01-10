# USAGE
# /bin/sh append_if_exist.sh <some string> <some file>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ -f "${2}" ]
then
    echo "" >> "${2}"
    echo "${1}" >> "${2}"
fi