# serveHttp
```
  TASK NAME     : serveHttp
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/serveHttp.zaruba.yaml
  DESCRIPTION   : Run static web server from your working directory.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.startService ]
  START         : - {{ .GetEnv "ZARUBA_HOME" }}/zaruba
                  - util
                  - serve
                  - .
                  - {{ index (.Split (.Trim (.GetConfig "ports") "\n ") "\n") 0 }}
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
                      exit 0
                    {{ end -}}
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Trim (.GetConfig "_check") "\n " }}
                    {{ .Trim (.GetConfig "check") "\n " }}
                    {{ .Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    {{ .Trim (.GetConfig "_finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  INPUTS        : serverHttpPort
                    DESCRIPTION : HTTP port to serve static files
                    PROMPT      : HTTP port
                    OPTIONS     : [ 8080, 8000, 3000, 5000 ]
                    DEFAULT     : 8080
                    VALIDATION  : ^[0-9]+$
  CONFIG        : RunInLocal        : true
                  _finish           : Blank
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterCheck        : Blank
                  afterStart        : Blank
                  beforeCheck       : Blank
                  beforeStart       : Blank
                  check             : {{- $d := .Decoration -}}
                                      {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                        {{ if ne $port "" -}}
                                          echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                          waitPort "localhost" {{ $port }}
                                          echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                        {{ end -}}
                                      {{ end -}}
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  ports             : {{ .GetValue "serverHttpPort" }}
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```