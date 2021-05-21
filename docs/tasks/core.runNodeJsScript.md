# core.runNodeJsScript
```
  TASK NAME     : core.runNodeJsScript
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.run.zaruba.yaml
  DESCRIPTION   : Run Node.Js script
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
  CONFIG        : _setup      : Blank
                  _start      : Blank
                  afterStart  : Blank
                  beforeStart : Blank
                  cmd         : node
                  cmdArg      : -p
                  finish      : Blank
                  setup       : Blank
                  start       : console.log('hello world')
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
