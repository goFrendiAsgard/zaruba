# zrbGenerateAndRun
```
  TASK NAME     : zrbGenerateAndRun
  LOCATION      : /zaruba-tasks/_base/run/task.zrbGenerateAndRun.yaml
  DESCRIPTION   : Generate script and run it
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
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
  CONFIG        : _finish          : Blank
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : {{ .GetConfig "generateScript" }}
                                     {{ .GetConfig "uploadScript" }}
                                     {{ .GetConfig "runScript" }}
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  generateScript   : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeShellUtil : . ${ZARUBA_HOME}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  localScriptFile  : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  remoteScriptFile : /{{ .Name }}.tmp.sh
                  runScript        : Blank
                  scriptTemplate   : Blank
                  setup            : Blank
                  start            : Blank
                  uploadScript     : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```