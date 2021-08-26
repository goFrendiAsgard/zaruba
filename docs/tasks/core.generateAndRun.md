# core.generateAndRun
```
  TASK NAME     : core.generateAndRun
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.generateAndRun.zaruba.yaml
  DESCRIPTION   : Generate script and run it
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
  CONFIG        : _executeScript    : Blank
                  _generateScript   : {{ $err := .WriteFile (.GetConfig "_localScriptFile") (.GetConfig "_template") -}}
                  _localScriptFile  : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  _remoteScriptFile : /{{ .Name }}.tmp.sh
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ .GetConfig "_generateScript" }}
                                      {{ .GetConfig "_uploadScript" }}
                                      {{ .GetConfig "_executeScript" }}
                  _template         : Blank
                  _uploadScript     : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```