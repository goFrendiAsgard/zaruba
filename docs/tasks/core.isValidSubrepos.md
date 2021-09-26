# core.isValidSubrepos
```
  TASK NAME     : core.isValidSubrepos
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.isValidSubrepos.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
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
                                      {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : {{ $d := .Decoration -}}
                                      {{ $names := .GetSubValueKeys "subrepo" -}}
                                      {{ $this := . -}}
                                      {{ range $index, $name := $names -}}
                                        PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
                                        URL="{{ $this.GetValue "subrepo" $name "url" }}"
                                        NAME="{{ $name }}"
                                        if [ -z "${URL}" ]
                                        then
                                          echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have url{{ $d.Normal }}"
                                          exit 1
                                        fi
                                        if [ -z "${PREFIX}" ]
                                        then
                                          echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have prefix{{ $d.Normal }}"
                                          exit 1
                                        fi
                                      {{ end }}
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}All Subrepos are valid{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```