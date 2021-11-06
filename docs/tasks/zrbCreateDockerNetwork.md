# zrbCreateDockerNetwork
```
  TASK NAME     : zrbCreateDockerNetwork
  LOCATION      : /zaruba-tasks/_base/dockerChore/task.zrbCreateDockerNetwork.yaml
  DESCRIPTION   : Create docker network.
                  Common configs:
                    network : Network name
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
                  network          : {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     set -e
                                     if [ "$(inspectDocker network ".Name" "{{ .GetConfig "network" }}")" = "{{ .GetConfig "network" }}" ]
                                     then
                                       echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Network '{{ .GetConfig "network" }}' is already exist{{ $d.Normal }}"
                                     else
                                       echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Creating network '{{ .GetConfig "network" }}'{{ $d.Normal }}"
                                       docker network create "{{ .GetConfig "network" }}"
                                     fi
                  strictMode       : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```