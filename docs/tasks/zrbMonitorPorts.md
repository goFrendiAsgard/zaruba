# zrbMonitorPorts
```
  TASK NAME     : zrbMonitorPorts
  LOCATION      : /scripts/tasks/zrbMonitorPorts.zaruba.yaml
  DESCRIPTION   : Throw error when any port is inactive
                  Common config:
                    ports : Port to be checked to confirm service readiness, separated by new line.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ zrbStartApp ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runInLocal") -}}
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
  CONFIG        : runInLocal        : true
                  _check            : {{ $d := .Decoration -}}
                                      echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Port monitoring started for: ${PORTS}{{ $d.Normal }}"
                  _finish           : Blank
                  _setup            : {{ .GetConfig "_setupPorts" }}
                  _setupPorts       : PORTS=""
                                      {{ range $index, $hostPort := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports" "\n ") "\n") -}}
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
                  interval          : 1
                  ports             : 8080
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```