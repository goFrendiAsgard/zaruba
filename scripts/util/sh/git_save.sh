#!/bin/sh

# USAGE
# /bin/sh git_save.sh <message>

(echo $- | grep -Eq ^.*e.*$) && OLD_STATE=-e || OLD_STATE=+e
set +e
git add . -A
git commit -m "${1}"
set "$OLD_STATE"
