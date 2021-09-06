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