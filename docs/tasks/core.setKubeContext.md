# core.setKubeContext
```
  TASK NAME     : core.setKubeContext
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.setKubeContext.zaruba.yaml
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
  CONFIG        : _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  kubeContext       : {{ .GetValue "kubeContext" }}
                  setup             : Blank
                  start             : if [ "$(kubectl config current-context)" != "{{ .GetConfig "kubeContext" }}" ]
                                      then
                                        kubectl config set-context "{{ .GetConfig "kubeContext" }}"
                                      fi
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```