# setupSdkman
```
  TASK NAME     : setupSdkman
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Install SDKMan for Java/Scala development.
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
                    DESCRIPTION : Home directory
                    PROMPT      : Home directory (Can be blank)
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will install sdkman, java, and scala in your home directory. Root privilege is not required"
                                           echo "If this command doesn't run successfully, please open an issue on https://github.com/state-alcemists/zaruba."
                                           echo "Please also specify your OS version."
                  cmd                    : /bin/bash
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
                                           if [ "$(is_command_exist sdk version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Sdkman was already installed{{ $d.Normal }}"
                                           else
                                             echo "☕ {{ $d.Bold }}{{ $d.Yellow }}Install sdkman{{ $d.Normal }}"
                                             curl -s "https://get.sdkman.io" | bash
                                             TEMPLATE_CONTENT="$(cat "${ZARUBA_HOME}/scripts/templates/shell/sdkman.sh")"
                                             append_if_exist "${TEMPLATE_CONTENT}" "${BOOTSTRAP_SCRIPT}"
                                             . "${BOOTSTRAP_SCRIPT}"
                                             echo "☕{{ $d.Bold }}{{ $d.Yellow }}Install java{{ $d.Normal }}"
                                             sdk install java
                                             echo "☕ {{ $d.Bold }}{{ $d.Yellow }}Install scala{{ $d.Normal }}"
                                             sdk install scala
                                           fi 
                                           if [ "$(is_command_exist java -version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Java was already installed{{ $d.Normal }}"
                                           else
                                             echo "☕{{ $d.Bold }}{{ $d.Yellow }}Install java{{ $d.Normal }}"
                                             sdk install java
                                           fi
                                           if [ "$(is_command_exist scala -version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Scala was already installed{{ $d.Normal }}"
                                           else
                                             echo "☕{{ $d.Bold }}{{ $d.Yellow }}Install scala{{ $d.Normal }}"
                                             sdk install scala
                                           fi
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Complete !!!{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
