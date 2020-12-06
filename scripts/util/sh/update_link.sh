#!/bin/sh

# USAGE
# sh update_link.sh <src> <destination>

_SRC="${1}"
_DST="${2}"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set -e
rm -Rf "${_DST}"
cp -r "${_SRC}" "${_DST}"
chmod 555 "${_DST}"
set "${_OLD_STATE}"
