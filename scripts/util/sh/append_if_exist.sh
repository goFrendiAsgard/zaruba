#/bin/sh

# USAGE
# /bin/sh append_if_exist.sh <some string> <some file>

if [ -f "${2}" ]
then
    echo "${1}" >> "${2}"
fi