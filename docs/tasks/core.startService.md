# core.startService
```
  TASK NAME     : core.startService
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.startService.zaruba.yaml
  DESCRIPTION   : Start service and check it's readiness.
                  Common config:
                    setup       : Script to be executed before start service or check service readiness.
                    start       : Script to start the service (e.g: python -m http.server 9000).
                    beforeStart : Script to be executed before start service.
                    afterStart  : Script to be executed after start service.
                    beforeCheck : Script to be executed before check service readiness.
                    afterCheck  : Script to be executed before check service readiness.
                    finish      : Script to be executed after start service or check service readiness.
                    RunInLocal  : Run service locally or not.
                    ports       : Port to be checked to confirm service readiness, separated by new line.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
                      sleep infinity
                    {{ end -}}
                    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
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
  CONFIG        : RunInLocal        : true
                  _finish           : Blank
                  _setup            : set -e
                                      {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterCheck        : Blank
                  afterStart        : Blank
                  beforeCheck       : Blank
                  beforeStart       : Blank
                  check             : {{- $d := .Decoration -}}
                                      {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
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
                  ports             : Blank
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```