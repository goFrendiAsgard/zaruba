# core.wrapper
```
  TASK NAME     : core.wrapper
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.wrapper.zaruba.yaml
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
  CONFIG        : _setup         : Blank
                  _start         : Blank
                  afterStart     : Blank
                  beforeStart    : Blank
                  cmd            : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg         : -c
                  finish         : Blank
                  playBellScript : echo $'\a'
                  setup          : Blank
                  start          : {{ .GetConfig "playBellScript" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```