# core.runScript
```
  TASK NAME     : core.runScript
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.runScript.zaruba.yaml
  DESCRIPTION   : Run script.
                  Common config:
                    cmd           : Executable shell name
                    cmdArg        : Executable shell argument
                    setup         : Setup script
                    beforeStart   : Before start script
                    start         : Start script
                    afterStart    : After start script
                    finish        : Finish script
  TASK TYPE     : Command Task
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup      : Blank
                  _start      : Blank
                  afterStart  : Blank
                  beforeStart : Blank
                  cmd         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg      : -c
                  finish      : Blank
                  setup       : Blank
                  start       : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```