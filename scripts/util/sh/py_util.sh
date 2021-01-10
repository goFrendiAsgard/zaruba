# USAGE
# sh py_util.sh <py-util> [arg1]... [arg9]

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

export PIPENV_IGNORE_VIRTUALENVS=1
export PIPENV_DONT_LOAD_ENV=1
export PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/util/python/Pipfile"
pipenv run python "${ZARUBA_HOME}/scripts/util/python/${1}.py" "${2}" "${3}" "${4}" "${5}" "${6}" "${7}" "${8}" "${9}" "${10}"
unset PIPENV_IGNORE_VIRTUALENVS
unset PIPENV_DONT_LOAD_ENV
unset PIPENV_PIPFILE