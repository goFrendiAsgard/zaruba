# setProjectValue
```
  TASK NAME     : setProjectValue
  LOCATION      : /zaruba-tasks/chore/value/task.setProjectValue.yaml
  DESCRIPTION   : Set project value.
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunShellScript ]
  DEPENDENCIES  : [ zrbIsProject ]
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
  INPUTS        : variableName
                    DESCRIPTION : Variable name (Required)
                    PROMPT      : Name
                    VALIDATION  : ^.+$
                  variableValue
                    DESCRIPTION : Variable value (Required)
                    PROMPT      : Value
                    VALIDATION  : ^.+$
  CONFIG        : _finish          : Blank
                  _initShell       : {{ if .IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
                                     {{ if .IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}
                  _setup           : {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : true
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "{{ .GetConfig "variableName" }}" "{{ .GetConfig "variableValue" }}"
                                     echo 🎉🎉🎉
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Kwarg ${KEY} : ${VALUE} has been set{{ $d.Normal }}"
                  strictMode       : true
                  variableName     : {{ .GetValue "variableName" }}
                  variableValue    : {{ .GetValue "variableValue" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```