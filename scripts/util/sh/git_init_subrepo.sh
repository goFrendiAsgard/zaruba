#!/bin/sh

# USAGE
# sh git_init_subrepo.sh <name> <prefix> <url> <branch>


_NAME="${1}"
_PREFIX="${2}"
_URL="${3}"
_BRANCH="${4}"

set +e
git remote add "${_NAME}" "${_URL}"
git subtree add --prefix="${_PREFIX}" "${_NAME}" "${_BRANCH}"
git fetch "${_NAME}" "${_BRANCH}"
git pull "${_NAME}" "${_BRANCH}"
