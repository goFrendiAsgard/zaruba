# core.helmInstall
```
  TASK NAME     : core.helmInstall
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmInstall.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.helmRepoUpdate, core.setKubeContext ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    {{ .Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  chart             : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  releaseName       : Blank
                  setup             : Blank
                  start             : if [ "$(isCommandError helm status "{{ .GetConfig "releaseName" }}")" -eq 1 ]
                                      then
                                        helm install "{{ .GetConfig "releaseName" }}" "{{ .GetConfig "chart" }}" -f "{{ .GetConfig "valueFile" }}"
                                      else
                                        helm upgrade "{{ .GetConfig "releaseName" }}" "{{ .GetConfig "chart" }}" -f "{{ .GetConfig "valueFile" }}"
                                      fi
                  valueFile         : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```