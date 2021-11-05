# zrbStopDockerContainer
```
  TASK NAME     : zrbStopDockerContainer
  LOCATION      : /zaruba-tasks/_base/dockerChore/task.zrbStopDockerContainer.yaml
  DESCRIPTION   : Stop docker container.
                  Common configs:
                    containerName : Container's name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
  DEPENDENCIES  : [ updateProjectLinks ]
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
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  containerName    : Blank
                  finish           : Blank
                  includeShellUtil : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     CONTAINER="{{ if .GetConfig "containerName" }}{{ .GetConfig "containerName" }}{{ else }}$("{{ .ZarubaBin }}" path getAppName "$(pwd)"){{ end }}"
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Stop docker container ${CONTAINER}{{ $d.Normal }}"
                                     stopContainer "${CONTAINER}" 
                                     echo 🎉🎉🎉
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Docker container ${CONTAINER} stopped{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```