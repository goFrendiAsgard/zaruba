#!/bin/sh

# USAGE
# /bin/sh get_latest_git_tag_commit.sh

git rev-parse --verify "$(git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags)"^{commit}