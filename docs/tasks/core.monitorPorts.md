# core.monitorPorts
```
  TASK NAME     : core.monitorPorts
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.monitorPorts.zaruba.yaml
  DESCRIPTION   : Throw error when any port is inactive
                  Common config:
                    ports : Port to be checked to confirm service readiness, separated by new line.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.startService ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
                      sleep infinity
                    {{ end -}}
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
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
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Trim (.GetConfig "_check") "\n " }}
                    {{ .Trim (.GetConfig "check") "\n " }}
                    {{ .Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  CONFIG        : RunInLocal        : true
                  _check            : {{ $d := .Decoration -}}
                                      echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Port monitoring started for: ${PORTS}{{ $d.Normal }}"
                  _setup            : {{ .GetConfig "_setupPorts" }}
                  _setupPorts       : PORTS=""
                                      {{ range $index, $hostPort := .Split (.Trim (.GetConfig "ports" "\n ") "\n") -}}
                                        {{ if ne $port "" -}}
                                          PORTS="${PORTS} {{ $port }}"
                                        {{ end -}}
                                      {{ end -}}
                  _start            : {{ $d := .Decoration -}}
                                      while true
                                      do
                                        {{ .GetConfig "_startCheckPorts" }}
                                        sleep {{ .GetConfig "interval" }}
                                      done
                  _startCheckPorts  : {{ $d := .Decoration -}}
                                      for PORT in ${PORTS}
                                      do
                                        if nc -z "localhost" "${PORT}"
                                        then
                                          continue
                                        fi
                                        echo "🔎 {{ $d.Bold }}{{ $d.Red }}Port '${PORT}' is not listening{{ $d.Normal }}"
                                        exit 1
                                      done
                  afterCheck        : Blank
                  afterStart        : Blank
                  beforeCheck       : Blank
                  beforeStart       : Blank
                  check             : {{- $d := .Decoration -}}
                                      {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                        {{ if ne $port "" -}}
                                          echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                          wait_port "localhost" {{ $port }}
                                          echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                        {{ end -}}
                                      {{ end -}}
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  interval          : 1
                  ports             : 8080
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```