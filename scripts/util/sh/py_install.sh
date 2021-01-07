#!/bin/sh

# USAGE
# sh py_install.sh <py-util> [arg1]... [arg9]

_SH_UTIL=$(dirname $(realpath $0))
_PY_UTIL=$(realpath "${_SH_UTIL}/../python")

export PIPENV_IGNORE_VIRTUALENVS=1
export PIPENV_DONT_LOAD_ENV=1
export PIPENV_PIPFILE="${_PY_UTIL}/Pipfile"
pipenv sync