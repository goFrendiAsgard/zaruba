# USAGE
# /bin/sh get_latest_git_tag.sh

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags

set "${_OLD_STATE}"