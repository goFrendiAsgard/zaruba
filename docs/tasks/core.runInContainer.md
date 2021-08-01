# core.runInContainer
```
  TASK NAME     : core.runInContainer
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.runInContainer.zaruba.yaml
  DESCRIPTION   : Run command from inside the container
                  Common config:
                    containerName : Name of the container
                    commands      : Command to be executed, separated by new line
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
                  _start            : {{ $this := . -}}
                                      {{ range $index, $command := .Split (.Trim (.GetConfig "commands") " \n") "\n" -}}
                                        {{ if ne $command "" -}}
                                          docker exec "{{ $this.GetConfig "containerName" }}" {{ $command }}
                                        {{ end -}}
                                      {{ end -}}
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  commands          : Blank
                  finish            : Blank
                  imagePrefix       : Blank
                  imageTag          : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup             : Blank
                  start             : echo "No script defined"
                  useImagePrefix    : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```