# core.prepareNodeJsApp
```
  TASK NAME     : core.prepareNodeJsApp
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.prepareNodeJsApp.zaruba.yaml
  DESCRIPTION   : Prepare NodeJs Application
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runNvmScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                  : set -e
                                            {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                            {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                                            {{ .Trim (.GetConfig "useNvmScript") "\n" }} 
                  _start                  : {{ .Trim (.GetConfig "removeNodeModulesScript") "\n" }} 
                                            {{ .Trim (.GetConfig "npmCleanCacheScript") "\n" }} 
                                            {{ .Trim (.GetConfig "npmInstallScript") "\n" }} 
                                            {{ .Trim (.GetConfig "npmRebuildScript") "\n" }} 
                                            {{ .Trim (.GetConfig "tsInstallScript") "\n" }} 
                                            {{ .Trim (.GetConfig "tsCompileScript") "\n" }} 
                  afterStart              : Blank
                  beforeStart             : Blank
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
                                            BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                            . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript       : . ${ZARUBA_HOME}/scripts/bash/util.sh
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
                  start                   : echo Prepare NodeJs App
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
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```