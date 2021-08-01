# showLog
```
  TASK NAME     : showLog
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/showLog.zaruba.yaml
  DESCRIPTION   : Show log for all/particular tasks using regex
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : logKeyword
                    DESCRIPTION : Task regex pattern
                    PROMPT      : Task regex pattern
  CONFIG        : _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  logKeyword        : {{ if .GetValue "logKeyword" }}{{ .GetValue "logKeyword" }}{{ else }}.*{{ end }}
                  setup             : Blank
                  start             : {{ $d := .Decoration -}}
                                      if [ ! -f "log.zaruba.csv" ]
                                      then
                                        echo "{{ $d.Bold }}{{ $d.Red }}Log is not exist{{ $d.Normal }}"
                                        exit 1
                                      fi
                                      "{{ .ZarubaBin }}" showLog "{{ .GetWorkPath "log.zaruba.csv" }}" "{{ .GetConfig "logKeyword"}}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```