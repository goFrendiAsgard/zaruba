# updateProjectLinks
```
  TASK NAME     : updateProjectLinks
  LOCATION      : /zaruba-tasks/chore/link/task.updateProjectLinks.yaml
  DESCRIPTION   : Update "links" in your project. Very useful if you have multiple apps sharing some parts of code
                  USAGE:
                    zaruba please updateProjectLinks
                    zaruba please updateProjectLinks "link::fibo/css=common-css"
                    zaruba please updateProjectLinks "link::app/css=common-css"
                  ARGUMENTS
                    link::<destination> : Location of the shared code
                  TIPS:
                    It is recommended to put `link` arguments in `default.values.yaml`.
                    In order to do that, you can invoke `zaruba please addProjectLink <linkFrom=source-location> <linkTo=destination-location>`
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
                  _initShell       : {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
                                     {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}
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
                                     {{ $this := . -}}
                                     {{ $destinations := .GetSubValueKeys "link" -}}
                                     {{ range $index, $destination := $destinations -}}
                                       {{ $source := $this.GetValue "link" $destination -}}
                                       {{ $absSource := $this.GetWorkPath $source -}}
                                       {{ $absDestination := $this.GetWorkPath $destination -}}
                                       linkResource "{{ $absSource }}" "{{ $absDestination }}"
                                     {{ end -}}
                                     echo 🎉🎉🎉
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Links updated{{ $d.Normal }}"
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```