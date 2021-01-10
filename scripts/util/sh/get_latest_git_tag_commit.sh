# USAGE
# /bin/sh get_latest_git_tag_commit.sh

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

git rev-parse --verify "$(git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags)"^{commit}