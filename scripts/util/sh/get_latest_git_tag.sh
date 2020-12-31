#!/bin/sh

# USAGE
# /bin/sh get_latest_git_tag.sh

git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags