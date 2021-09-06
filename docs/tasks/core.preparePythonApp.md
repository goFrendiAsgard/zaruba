# core.preparePythonApp
```
  TASK NAME     : core.preparePythonApp
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.preparePythonApp.zaruba.yaml
  DESCRIPTION   : Prepare Python App
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runShellScript ]
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
  CONFIG        : _finish     : Blank
                  _setup      : Blank
                  _start      : pipenv install
                  afterStart  : Blank
                  beforeStart : Blank
                  cmd         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg      : -c
                  finish      : Blank
                  setup       : Blank
                  start       : echo Prepare Python App
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```