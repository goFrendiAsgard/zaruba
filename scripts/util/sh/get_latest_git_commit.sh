# USAGE
# /bin/sh get_latest_git_commit.sh

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

git rev-parse --verify HEAD

set "${_OLD_STATE}"