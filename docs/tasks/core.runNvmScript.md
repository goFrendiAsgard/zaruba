# core.runNvmScript
```
  TASK NAME     : core.runNvmScript
  LOCATION      : /scripts/tasks/core.runNvmScript.zaruba.yaml
  DESCRIPTION   : Run shell script under nvm
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runShellScript ]
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
  CONFIG        : _finish                 : Blank
                  _setup                  : set -e
                                            {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                                            {{ .Util.Str.Trim (.GetConfig "useNvmScript") "\n" }} 
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
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
                  removeNodeModules       : false
                  removeNodeModulesScript : {{ if .IsTrue (.GetConfig "removeNodeModules") -}}
                                              rm -Rf node_modules
                                            {{ end -}}
                  setup                   : Blank
                  start                   : echo hello world
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