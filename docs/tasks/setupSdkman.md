# setupSdkman
```
  TASK NAME     : setupSdkman
  LOCATION      : ${ZARUBA_HOME}/scripts/task.setupSdkman.zaruba.yaml
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
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup                  : Blank
                  start                  : {{ $d := .Decoration -}}
                                           {{ if .GetValue "setupHomeDir" }}HOME="{{ .GetValue "setupHomeDir" }}"{{ end }}
                                           if [ "$(is_command_exist sdk version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Sdkman was already installed{{ $d.Normal }}"
                                           else
                                             rm -Rf "${HOME}/.sdkman"
                                             echo "☕ {{ $d.Bold }}{{ $d.Yellow }}Install sdkman{{ $d.Normal }}"
                                             curl -s "https://get.sdkman.io" | bash
                                             TEMPLATE_CONTENT="$(cat "${ZARUBA_HOME}/scripts/templates/shell/sdkman.sh")"
                                             echo "" >> "${BOOTSTRAP_SCRIPT}"
                                             echo "${TEMPLATE_CONTENT}" >> "${BOOTSTRAP_SCRIPT}"
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