# pullSubrepos
```
  TASK NAME     : pullSubrepos
  LOCATION      : /scripts/tasks/pullSubrepos.zaruba.yaml
  DESCRIPTION   : Pull subrepositories.
                  ARGUMENTS:
                    subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
                    subrepo::<name>::url      : Remote url of the subrepo
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
  DEPENDENCIES  : [ initSubrepos ]
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
                  start             : set -e
                                      {{ $d := .Decoration -}}
                                      {{ $names := .GetSubValueKeys "subrepo" -}}
                                      {{ $this := . -}}
                                      ORIGINS=$("{{ .ZarubaBin }}" str split "$(git remote)")
                                      BRANCH="{{ if .GetValue "defaultBranch" }}{{ .GetValue "defaultBranch" }}{{ else }}main{{ end }}"
                                      {{ range $index, $name := $names -}}
                                        PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
                                        URL="{{ $this.GetValue "subrepo" $name "url" }}"
                                        NAME="{{ $name }}"
                                        ORIGIN_EXISTS=$("{{ $this.ZarubaBin }}" list contain "${ORIGINS}" "${NAME}")
                                        if [ $ORIGIN_EXISTS = 1 ]
                                        then
                                          gitSave "Save works before pull"
                                          git subtree pull --prefix="${PREFIX}" "${NAME}" "${BRANCH}"
                                        fi
                                      {{ end -}}
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepos pulled{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```