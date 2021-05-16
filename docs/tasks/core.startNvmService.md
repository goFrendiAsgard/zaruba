# core.startNvmService
```
      TASK NAME    : core.startNvmService
      LOCATION     : /home/gofrendi/.zaruba/scripts/core.service.zaruba.yaml
      DESCRIPTION  : Start service and check it's readiness.
                     Common config:
                       setup       : Script to be executed before start service or check service readiness.
                       start       : Script to start the service (e.g: python -m http.server 9000).
                       beforeStart : Script to be executed before start service.
                       afterStart  : Script to be executed after start service.
                       beforeCheck : Script to be executed before check service readiness.
                       afterCheck  : Script to be executed before check service readiness.
                       finish      : Script to be executed after start service or check service readiness.
                     Common lconfig:
                       ports: Port to be checked to confirm service readiness (e.g: [9000])
      TASK TYPE    : Service Task
      PARENT TASKS : [ core.startService ]
      START        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{- $d := .Decoration -}}
                       {{ if .IsFalse (.GetConfig "runLocally") -}}
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
      CHECK        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{- $d := .Decoration -}}
                       {{ if .IsFalse (.GetConfig "runLocally") -}}
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
      CONFIG       :   _setup                  : set -e
                                                 {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                 {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                                                 {{ .Trim (.GetConfig "useNvmScript") "\n" }} 
                       _start                  : Blank
                       afterCheck              : Blank
                       afterStart              : Blank
                       beforeCheck             : Blank
                       beforeStart             : Blank
                       check                   : {{- $d := .Decoration -}}
                                                 {{ range $index, $port := .GetLConfig "ports" -}}
                                                   echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                                   wait_port "localhost" "{{ $port }}"
                                                   echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                                 {{ end -}}
                       cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                       cmdArg                  : -c
                       compileTypeScript       : false
                       finish                  : Blank
                       includeBootstrapScript  : if [ -f "${HOME}/.profile" ]
                                                 then
                                                     . "${HOME}/.profile"
                                                 fi
                                                 if [ -f "${HOME}/.bashrc" ]
                                                 then
                                                     . "${HOME}/.bashrc"
                                                 fi
                                                 BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                                 . "${BOOTSTRAP_SCRIPT}"
                       includeUtilScript       : . ${ZARUBA_HOME}/scripts/util.sh
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
                       playBellScript          : echo $'\a'
                       removeNodeModules       : false
                       removeNodeModulesScript : {{ if .IsTrue (.GetConfig "removeNodeModules") -}}
                                                   rm -Rf node_modules
                                                 {{ end -}}
                       runLocally              : true
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
                                                   if [ -f "./node_modules/.bin/tsc" ] || [ "$(is_command_exist tsc)" = 1 ]
                                                   then
                                                     echo "Typescript is already installed"
                                                   else
                                                     npm install -g typescript{{ if .GetConfig "typeScriptVersion" }}@{{ .GetConfig "typeScriptVersion" }}{{ end }}
                                                   fi
                                                 {{ end -}}
                       typeScriptVersion       : Blank
                       useNvmScript            : if [ "$(is_command_exist nvm)" = 1 ]
                                                 then
                                                   if [ "$(is_command_error nvm ls "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}" )" ]
                                                   then
                                                     nvm install "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}"
                                                   else
                                                     nvm use "{{ if .GetConfig "nodeVersion" }}{{ .GetConfig "nodeVersion" }}{{ else }}node{{ end }}"
                                                   fi
                                                 fi
      LCONFIG      :   ports : []
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
