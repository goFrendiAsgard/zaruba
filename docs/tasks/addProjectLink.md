# addProjectLink
```
  TASK NAME     : addProjectLink
  LOCATION      : /zaruba-tasks/chore/link/task.addProjectLink.yaml
  DESCRIPTION   : Add link.
                  TIPS: To update links, you should perform `zaruba please updateProjectLinks`
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
  INPUTS        : linkFrom
                    DESCRIPTION : Link source (Required)
                    PROMPT      : Source
                    VALIDATION  : ^.+$
                  linkTo
                    DESCRIPTION : Link destination (Required)
                    PROMPT      : Destination
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
                  linkFrom         : {{ .GetValue "linkFrom" }}
                  linkTo           : {{ .GetValue "linkTo" }}
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "link::{{ .GetConfig "linkTo" }}" "{{ .GetConfig "linkFrom" }}"
                                     echo 🎉🎉🎉
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Link ${SOURCE} -> ${DESTINATION} has been added{{ $d.Normal }}"
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```