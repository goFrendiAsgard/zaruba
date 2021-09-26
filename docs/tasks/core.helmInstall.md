# core.helmInstall
```
  TASK NAME     : core.helmInstall
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmInstall.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.helmRepoUpdate, core.setKubeContext ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
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