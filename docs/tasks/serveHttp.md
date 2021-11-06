# serveHttp
```
  TASK NAME     : serveHttp
  LOCATION      : /zaruba-tasks/chore/serveHttp/task.serveHttp.yaml
  DESCRIPTION   : Run static web server from your working directory.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ zrbStartApp ]
  START         : - {{ .GetEnv "ZARUBA_HOME" }}/zaruba
                  - serve
                  - .
                  - {{ index (.Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n") 0 }}
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
                      exit 0
                    {{ end -}}
                    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_check") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "check") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  INPUTS        : serverHttpPort
                    DESCRIPTION : HTTP port to serve static files
                    PROMPT      : HTTP port
                    OPTIONS     : [ 8080, 8000, 3000, 5000 ]
                    DEFAULT     : 8080
                    VALIDATION  : ^[0-9]+$
  CONFIG        : _finish          : Blank
                  _initShell       : {{ if .IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
                                     {{ if .IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}
                  _setup           : {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}
                  _start           : Blank
                  afterCheck       : Blank
                  afterStart       : Blank
                  beforeCheck      : Blank
                  beforeStart      : Blank
                  check            : {{- $d := .Decoration -}}
                                     {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                       {{ if ne $port "" -}}
                                         echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                         waitPort "localhost" {{ $port }}
                                         echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                       {{ end -}}
                                     {{ end -}}
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : true
                  ports            : {{ .GetValue "serverHttpPort" }}
                  runInLocal       : true
                  setup            : Blank
                  start            : Blank
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```