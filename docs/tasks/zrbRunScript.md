# zrbRunScript
```
  TASK NAME     : zrbRunScript
  LOCATION      : /zaruba-tasks/_base/run/task.zrbRunScript.yaml
  DESCRIPTION   : Run script.
                  Common configs:
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
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish     : Blank
                  _setup      : Blank
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