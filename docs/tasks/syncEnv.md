# syncEnv
```
  TASK NAME     : syncEnv
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/syncEnv.zaruba.yaml
  DESCRIPTION   : Update environment of every task in the current project
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.isProject ]
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
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : {{ $d := .Decoration -}}
                                      "{{ .ZarubaBin }}" project syncEnv "./main.zaruba.yaml"
                                      "{{ .ZarubaBin }}" project syncEnvFiles "./main.zaruba.yaml"
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Environment updated{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```