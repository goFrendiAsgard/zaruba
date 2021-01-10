# USAGE
# sh link_resource.sh <src> <destination>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

_SRC="${1}"
_DST="${2}"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set -e
echo "${Yellow}Link \"${_SRC}\" into \"${_DST}\"${Normal}"
if [ -e "${_DST}" ]
then
    chmod 777 -R "${_DST}"
    rm -Rf "${_DST}"
fi
cp -r "${_SRC}" "${_DST}"
chmod 555 -R "${_DST}"
set "${_OLD_STATE}"
