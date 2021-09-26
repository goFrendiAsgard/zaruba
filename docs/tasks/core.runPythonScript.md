# core.runPythonScript
```
  TASK NAME     : core.runPythonScript
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.runPythonScript.zaruba.yaml
  DESCRIPTION   : Run python script
                  Common config:
                    start : Start script
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runScript ]
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
                  cmd         : python
                  cmdArg      : -c
                  finish      : Blank
                  setup       : Blank
                  start       : print('hello world')
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```