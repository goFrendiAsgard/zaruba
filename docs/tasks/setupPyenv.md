# setupPyenv
```
  TASK NAME     : setupPyenv
  LOCATION      : /home/gofrendi/zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Install pyenv and pipenv
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
  INPUTS        : setup.homeDir
                    DESCRIPTION : Home directory (Can be blank)
                    PROMPT      : Home directory
                  setup.pythonVersion
                    DESCRIPTION : Python version to be installed when install pyenv
                    PROMPT      : Python version
                    OPTIONS     : [ 3.7, 3.8, 3.9 ]
                    DEFAULT     : 3.8.6
                    VALIDATION  : ^.+$
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will install pyenv and pipenv in your home directory. Root privilege is not required"
                                           echo "If this command doesn't run successfully, please open an issue on https://github.com/state-alcemists/zaruba."
                                           echo "Please also specify your OS version."
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
                  start                  : {{ $d := .Decoration -}}
                                           {{ if .GetValue "setup.homeDir" }}HOME="{{ .GetValue "setup.homeDir" }}"{{ end }}
                                           if [ "$(is_command_exist pyenv --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Pyenv was already installed{{ $d.Normal }}"
                                           else
                                             rm -Rf "${HOME}/.pyenv"
                                             echo "🐍 {{ $d.Bold }}{{ $d.Yellow }}Install pyenv{{ $d.Normal }}"
                                             curl -L https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | sh
                                             TEMPLATE_CONTENT="$(cat "${ZARUBA_HOME}/scripts/templates/shell/pyenv.sh")"
                                             append_if_exist "${TEMPLATE_CONTENT}" "${BOOTSTRAP_SCRIPT}"
                                             . "${BOOTSTRAP_SCRIPT}"
                                             echo "🐍 {{ $d.Bold }}{{ $d.Yellow }}Install python {{ .GetValue "setup.pythonVersion" }}{{ $d.Normal }}"
                                             pyenv install {{ .GetValue "setup.pythonVersion" }}
                                             pyenv global {{ .GetValue "setup.pythonVersion" }}
                                             . "${BOOTSTRAP_SCRIPT}"
                                           fi
                                           if [ "$(is_command_exist python --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Python was already installed{{ $d.Normal }}"
                                           else
                                             echo "🐍 {{ $d.Bold }}{{ $d.Yellow }}Install python {{ .GetValue "setup.pythonVersion" }}{{ $d.Normal }}"
                                             pyenv install {{ .GetValue "setup.pythonVersion" }}
                                             pyenv global {{ .GetValue "setup.pythonVersion" }}
                                           fi
                                           if [ "$(is_command_exist pipenv --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Pipenv was already installed{{ $d.Normal }}"
                                           else
                                             echo "🐍 {{ $d.Bold }}{{ $d.Yellow }}Install pipenv{{ $d.Normal }}"
                                             pip install pipenv
                                             if [ -d "${HOME}/.pipenv/shims" ]
                                             then
                                               chmod 755 "${HOME}/.pipenv/shims"
                                             fi
                                           fi
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Complete !!!{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
