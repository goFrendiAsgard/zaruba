# showLog
```
  TASK NAME     : showLog
  LOCATION      : /zaruba-tasks/chore/log/task.showLog.yaml
  DESCRIPTION   : Show log for all/particular tasks using regex
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
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
  INPUTS        : keyword
                    DESCRIPTION : Keyword
                    PROMPT      : Keyword
  CONFIG        : _finish          : Blank
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : . ${ZARUBA_HOME}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  keyword          : {{ if .GetValue "keyword" }}{{ .GetValue "keyword" }}{{ else }}.*{{ end }}
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
                                     if [ ! -f "log.zaruba.csv" ]
                                     then
                                       echo "{{ $d.Bold }}{{ $d.Red }}Log is not exist{{ $d.Normal }}"
                                       exit 1
                                     fi
                                     ""${ZARUBA_HOME}/zaruba"" project showLog "{{ .GetWorkPath "log.zaruba.csv" }}" "{{ .GetConfig "keyword"}}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```