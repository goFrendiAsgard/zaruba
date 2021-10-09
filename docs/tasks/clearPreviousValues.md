# clearPreviousValues
```
  TASK NAME     : clearPreviousValues
  LOCATION      : /zaruba-tasks/chore/value/task.clearPreviousValues.yaml
  DESCRIPTION   : Clear log
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
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : . ${ZARUBA_HOME}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     rm -Rf .previous.values.yaml
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Previous values removed{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```