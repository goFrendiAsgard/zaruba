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
                    {{ .Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ .GetConfig "generateScript" }}
                                      {{ .GetConfig "uploadScript" }}
                                      {{ .GetConfig "runScript" }}
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  generateScript    : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  localScriptFile   : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  remoteScriptFile  : /{{ .Name }}.tmp.sh
                  runScript         : Blank
                  scriptTemplate    : Blank
                  setup             : Blank
                  start             : Blank
                  uploadScript      : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```