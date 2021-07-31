# initProject
```
  TASK NAME     : initProject
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/initProject.zaruba.yaml
  DESCRIPTION   : Initiate empty zaruba project.
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
                                      alias zaruba=${ZARUBA_HOME}/zaruba
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup             : Blank
                  start             : {{ $d := .Decoration -}}
                                      if [ -f "main.zaruba.yaml" ]
                                      then
                                        echo "{{ $d.Bold }}{{ $d.Red }}$(pwd) is a zaruba project.{{ $d.Normal }}"
                                        exit 1
                                      fi
                                      git init
                                      cp -rT "${ZARUBA_HOME}/scripts/templates/project/" .
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Project created{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```