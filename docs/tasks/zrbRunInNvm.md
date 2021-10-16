# zrbRunInNvm
```
  TASK NAME     : zrbRunInNvm
  LOCATION      : /zaruba-tasks/_base/run/inNvm/task.zrbRunInNvm.yaml
  DESCRIPTION   : Run shell script under nvm
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
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtilScript") "\n" }}
                                     {{ .Util.Str.Trim (.GetConfig "useNvmScript") "\n" }} 
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  nodeVersion      : node
                  setup            : Blank
                  start            : Blank
                  useNvmScript     : if [ "$(isCommandExist nvm)" = 1 ]
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