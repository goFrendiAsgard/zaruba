#!/bin/sh

# USAGE
# sh update_link.sh <src> <destination>

_SRC="${1}"
_DST="${2}"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set -e
if [ -e "${_DST}" ]
then
    chmod 777 -R "${_DST}"
    rm -Rf "${_DST}"
fi
cp -r "${_SRC}" "${_DST}"
chmod 555 -R "${_DST}"
set "${_OLD_STATE}"
