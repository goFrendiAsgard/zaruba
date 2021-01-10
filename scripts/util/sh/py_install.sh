# USAGE
# sh py_install.sh <py-util> [arg1]... [arg9]

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

if [ -z "$(pipenv --version)" ]
then
    echo "${Bold}${Red}Pipenv is not installed${Normal}"
    echo "Please perform:"
    echo "* 'zaruba please setupPyenv' (recommended) or"
    echo "* 'pip install pipenv' (if you don't want to install pyenv)"
    exit 1
fi

export PIPENV_IGNORE_VIRTUALENVS=1
export PIPENV_DONT_LOAD_ENV=1
export PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/util/python/Pipfile"
pipenv sync
unset PIPENV_IGNORE_VIRTUALENVS
unset PIPENV_DONT_LOAD_ENV
unset PIPENV_PIPFILE