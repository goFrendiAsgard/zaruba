# core.setupPyUtil
```
  TASK NAME     : core.setupPyUtil
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Setup zaruba's python util.
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  finish                 : Blank
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  playBellScript         : echo $'\a'
                  setup                  : Blank
                  start                  : {{ if eq (.GetValue "setup.initPyUtil") "true" }}
                                             {{ $d := .Decoration -}}
                                             if [ -z "$(pipenv --version)" ]
                                             then
                                                 echo "{{ $d.Bold }}{{ $d.Red }}Pipenv is not installed{{ $d.Normal }}"
                                                 echo "Please perform:"
                                                 echo "* 'zaruba please setupPyenv' (recommended) or"
                                                 echo "* 'pip install pipenv' (if you don't want to install pyenv)"
                                                 exit 1
                                             fi
                                             export PIPENV_IGNORE_VIRTUALENVS=1
                                             export PIPENV_DONT_LOAD_ENV=1
                                             export PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/util/python/Pipfile"
                                             pipenv install
                                           {{ end }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
