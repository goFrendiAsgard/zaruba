# setProjectValue
```
  TASK NAME     : setProjectValue
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/setProjectValue.zaruba.yaml
  DESCRIPTION   : Set project value.
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.isProject ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : variableName
                    DESCRIPTION : Variable name (Required)
                    PROMPT      : Name
                    VALIDATION  : ^.+$
                  variableValue
                    DESCRIPTION : Variable value (Required)
                    PROMPT      : Value
                    VALIDATION  : ^.+$
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
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
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup                  : Blank
                  start                  : {{ $d := .Decoration -}}
                                           {{ .Zaruba }} setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "{{ .GetConfig "variableName" }}" "{{ .GetConfig "variableValue" }}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Kwarg ${KEY} : ${VALUE} has been set{{ $d.Normal }}"
                  variableName           : {{ .GetValue "variableName" }}
                  variableValue          : {{ .GetValue "variableValue" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```