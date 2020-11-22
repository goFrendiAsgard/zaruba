#!/bin/sh

# USAGE
# /bin/sh git_save.sh <message>

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

git add . -A
git commit -m "💀 ${1}"

set "${_OLD_STATE}"
