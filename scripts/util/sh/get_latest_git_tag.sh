# USAGE
# /bin/sh get_latest_git_tag.sh

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags