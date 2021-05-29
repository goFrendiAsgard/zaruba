# setupUbuntu
```
  TASK NAME     : setupUbuntu
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Setup ubuntu
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
  INPUTS        : setup.userPassword
                    DESCRIPTION : User password (Can be blank if you already run the task with "sudo")
                    PROMPT      : User password
                  setup.allowRoot
                    DESCRIPTION : Allow to install as root (using root home directory)
                    PROMPT      : Allow to install as root
                    OPTIONS     : [ yes, no ]
                    DEFAULT     : no
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will install essential packages for ubuntu/mint/debian. Root privilege is required."
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
                                           if [ $(whoami) = "root" ]
                                           then
                                             {{ if .IsFalse (.GetValue "setup.allowRoot") }}
                                               if echo "${1}" | grep -q "/root$"
                                               then
                                                   echo "{{ $d.Bold}}{{ $d.Red }}Your home directory seems to be '/root'. If this is intentional please set 'setup.allowRoot' to 'true'. Otherwise re-run this task with 'sudo -E'{{ $d.Normal}}"
                                                   exit 1
                                               fi
                                             {{ end }}
                                             . "${ZARUBA_HOME}/scripts/setup_ubuntu.sh"
                                           else
                                             {{ if .GetValue "setup.userPassword" }}
                                               echo "${ZARUBA_INPUT_SETUP_USERPASSWORD}" | sudo -E -S {{ .GetConfig "cmd" }} "${ZARUBA_HOME}/scripts/setup_ubuntu.sh"
                                             {{ else }}
                                               echo "{{ $d.Bold}}{{ $d.Red }}You need to set 'setup.userPassword' or run this task with 'sudo -E'{{ $d.Normal}}"
                                               exit 1
                                             {{ end }}
                                           fi
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
