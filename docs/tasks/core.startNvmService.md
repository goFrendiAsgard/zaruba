# core.startNvmService
```
  TASK NAME     : core.startNvmService
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.startNvmService.zaruba.yaml
  DESCRIPTION   : Start service and check it's readiness.
                  Common config:
                    setup       : Script to be executed before start service or check service readiness.
                    start       : Script to start the service (e.g: python -m http.server 9000).
                    beforeStart : Script to be executed before start service.
                    afterStart  : Script to be executed after start service.
                    beforeCheck : Script to be executed before check service readiness.
                    afterCheck  : Script to be executed before check service readiness.
                    finish      : Script to be executed after start service or check service readiness.
                    ports       : Port to be checked to confirm service readiness, separated by new line.
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
  CONFIG        : RunInLocal              : true
                  _finish                 : Blank
                  _setup                  : set -e
                                            {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                                            {{ .Util.Str.Trim (.GetConfig "useNvmScript") "\n" }} 
                  _start                  : Blank
                  afterCheck              : Blank
                  afterStart              : Blank
                  beforeCheck             : Blank
                  beforeStart             : Blank
                  check                   : {{- $d := .Decoration -}}
                                            {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                              {{ if ne $port "" -}}
                                                echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                                waitPort "localhost" {{ $port }}
                                                echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                              {{ end -}}
                                            {{ end -}}
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  compileTypeScript       : false
                  finish                  : Blank
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  installTypeScript       : false
                  nodeVersion             : node
                  npmCleanCache           : false
                  npmCleanCacheScript     : {{ if .IsTrue (.GetConfig "npmCleanCache") -}}
                                              npm cache clean --force
                                            {{ end -}}
                  npmInstallScript        : if [ ! -d "node_modules" ]
                                            then
                                              npm install
                                            fi
                  npmRebuild              : false
                  npmRebuildScript        : {{ if .IsTrue (.GetConfig "npmRebuild") -}}
                                              npm rebuild
                                            {{ end -}}
                  ports                   : Blank
                  removeNodeModules       : false
                  removeNodeModulesScript : {{ if .IsTrue (.GetConfig "removeNodeModules") -}}
                                              rm -Rf node_modules
                                            {{ end -}}
                  setup                   : Blank
                  start                   : Blank
                  tsCompileScript         : {{ if .IsTrue (.GetConfig "compileTypeScript") -}}
                                              if [ -f "./node_modules/.bin/tsc" ]
                                              then
                                                ./node_modules/.bin/tsc
                                              else
                                                tsc
                                              fi
                                            {{ end -}}
                  tsInstallScript         : {{ if .IsTrue (.GetConfig "installTypeScript") -}}
                                              if [ -f "./node_modules/.bin/tsc" ] || [ "$(isCommandExist tsc)" = 1 ]
                                              then
                                                echo "Typescript is already installed"
                                              else
                                                npm install -g typescript{{ if .GetConfig "typeScriptVersion" }}@{{ .GetConfig "typeScriptVersion" }}{{ end }}
                                              fi
                                            {{ end -}}
                  typeScriptVersion       : Blank
                  useNvmScript            : if [ "$(isCommandExist nvm)" = 1 ]
                                            then
                                              if [ "$(isCommandError nvm ls "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}" )" ]
                                              then
                                                nvm install "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}"
                                              else
                                                nvm use "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}"
                                              fi
                                            fi
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```