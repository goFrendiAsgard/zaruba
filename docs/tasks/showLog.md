# showLog
```
  TASK NAME     : showLog
  LOCATION      : /zaruba-tasks/chore/log/task.showLog.yaml
  DESCRIPTION   : Show log for all/particular tasks using regex
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
  INPUTS        : keyword
                    DESCRIPTION : Keyword
                    PROMPT      : Keyword
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
                  keyword          : {{ if .GetValue "keyword" }}{{ .GetValue "keyword" }}{{ else }}.*{{ end }}
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     if [ ! -f "log.zaruba.csv" ]
                                     then
                                       echo "{{ $d.Bold }}{{ $d.Red }}Log is not exist{{ $d.Normal }}"
                                       exit 1
                                     fi
                                     "{{ .ZarubaBin }}" project showLog "{{ .GetWorkPath "log.zaruba.csv" }}" "{{ .GetConfig "keyword"}}"
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```