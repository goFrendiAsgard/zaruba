# zrbIsValidSubrepos
```
  TASK NAME     : zrbIsValidSubrepos
  LOCATION      : /zaruba-tasks/_base/validation/task.zrbIsValidSubrepos.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunShellScript ]
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
                                     {{ $names := .GetSubValueKeys "subrepo" -}}
                                     {{ $this := . -}}
                                     {{ range $index, $name := $names -}}
                                       PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
                                       URL="{{ $this.GetValue "subrepo" $name "url" }}"
                                       NAME="{{ $name }}"
                                       if [ -z "${URL}" ]
                                       then
                                         echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have url{{ $d.Normal }}"
                                         exit 1
                                       fi
                                       if [ -z "${PREFIX}" ]
                                       then
                                         echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have prefix{{ $d.Normal }}"
                                         exit 1
                                       fi
                                     {{ end }}
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}All Subrepos are valid{{ $d.Normal }}"
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```