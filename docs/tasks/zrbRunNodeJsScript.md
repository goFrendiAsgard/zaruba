# zrbRunNodeJsScript
```
  TASK NAME     : zrbRunNodeJsScript
  LOCATION      : /zaruba-tasks/_base/run/task.zrbRunNodeJsScript.yaml
  DESCRIPTION   : Run Node.Js script
                  Common config:
                    start : Start script
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunScript ]
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
                  cmd         : node
                  cmdArg      : -p
                  finish      : Blank
                  setup       : Blank
                  start       : console.log('hello world')
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```