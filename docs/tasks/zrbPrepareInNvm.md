# zrbPrepareInNvm
```
  TASK NAME     : zrbPrepareInNvm
  LOCATION      : /zaruba-tasks/_base/run/inNvm/task.zrbPrepareInNvm.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunInNvm ]
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
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Util.Str.Trim (.GetConfig "includeShellUtilScript") "\n" }}
                                      {{ .Util.Str.Trim (.GetConfig "useNvmScript") "\n" }} 
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  compileTypeScript : false
                  finish            : Blank
                  includeShellUtil  : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  installTsc        : false
                  nodeVersion       : node
                  npmCleanCache     : false
                  npmRebuild        : false
                  removeNodeModules : false
                  setup             : Blank
                  start             : {{ if .IsTrue (.GetConfig "removeNodeModules") -}}
                                        rm -Rf node_modules
                                      {{ end -}}
                                      {{ if .IsTrue (.GetConfig "npmCleanCache") -}}
                                        npm cache clean --force
                                      {{ end -}}
                                      if [ ! -d "node_modules" ]
                                      then
                                        npm install
                                      fi
                                      {{ if .IsTrue (.GetConfig "npmRebuild") -}}
                                        npm rebuild
                                      {{ end -}}
                                      {{ if .IsTrue (.GetConfig "installTsc") -}}
                                        if [ -f "./node_modules/.bin/tsc" ] || [ "$(isCommandExist tsc)" = 1 ]
                                        then
                                          echo "Typescript is already installed"
                                        else
                                          npm install -g typescript{{ if .GetConfig "typeScriptVersion" }}@{{ .GetConfig "typeScriptVersion" }}{{ end }}
                                        fi
                                      {{ end -}}
                                      {{ if .IsTrue (.GetConfig "compileTypeScript") -}}
                                        if [ -f "./node_modules/.bin/tsc" ]
                                        then
                                          ./node_modules/.bin/tsc
                                        else
                                          tsc
                                        fi
                                      {{ end -}} 
                  typeScriptVersion : Blank
                  useNvmScript      : if [ "$(isCommandExist nvm)" = 1 ]
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