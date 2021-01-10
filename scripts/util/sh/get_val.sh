# USAGE
# /bin/sh get_val.sh <val> <default-val>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ -z ${1} ]
then
    echo ${2}
else
    echo ${1}
fi