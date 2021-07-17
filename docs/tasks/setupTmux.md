# setupTmux
```
  TASK NAME     : setupTmux
  LOCATION      : /home/gofrendi/zaruba/scripts/task.setupTmux.zaruba.yaml
  DESCRIPTION   : setup tmux (tmux should be already installed before running this task)
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
  INPUTS        : setupHomeDir
                    DESCRIPTION : Home directory (Can be blank)
                    PROMPT      : Home directory
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will setup tmux. Tmux should be installed first"
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
                                           {{ if .GetValue "setupHomeDir" }}HOME="{{ .GetValue "setupHomeDir" }}"{{ end }}
                                           if [ "$(is_command_exist tmux -V)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Tmux was already installed{{ $d.Normal }}"
                                           else
                                             echo "{{ $d.Bold }}{{ $d.Red }}Tmux was not installed. Please install tmux first{{ $d.Normal }}"
                                             exit 1
                                           fi
                                           if [ -d "${HOME}/.tmux/plugins/tpm" ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Tmux plugin was already installed{{ $d.Normal }}"
                                           else
                                             git clone https://github.com/tmux-plugins/tpm "${HOME}/.tmux/plugins/tpm"
                                           fi
                                           if [ -f "${HOME}/.tmux.conf" ]
                                           then
                                             mv "${HOME}/.tmux.conf" "${HOME}/.tmux.bak.conf"
                                           fi
                                           cp "${ZARUBA_HOME}/scripts/templates/tmux.conf" "${HOME}/.tmux.conf"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Complete !!!{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
