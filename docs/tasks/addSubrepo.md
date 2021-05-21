# addSubrepo
```
  TASK NAME     : addSubrepo
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.zaruba.yaml
  DESCRIPTION   : Add subrepository.
                  TIPS: To init added subrepositories, you should perform `zaruba please initSubrepos`
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.isProject, core.setupPyUtil ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : subrepo.url
                    DESCRIPTION : Subrepo url
                    PROMPT      : Subrepo url (Required)
                    VALIDATION  : ^.+$
                  subrepo.prefix
                    DESCRIPTION : Subrepo directory name
                    PROMPT      : Subrepo directory name (Can be blank)
                  subrepo.name
                    DESCRIPTION : Subrepo name
                    PROMPT      : Subrepo name (Can be blank)
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  finish                 : Blank
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/util.sh
                  playBellScript         : echo $'\a'
                  setup                  : Blank
                  start                  : set -e
                                           {{ $d := .Decoration -}}
                                           URL="{{ .GetValue "subrepo.url" }}"
                                           should_not_be_empty "${URL}" "{{ $d.Bold }}{{ $d.Red }}subrepo.url is not defined{{ $d.Normal }}"
                                           {{ if .GetValue "subrepo.prefix" }}
                                             PREFIX="{{ .GetValue "subrepo.prefix" }}"
                                           {{ else }}
                                             PREFIX=$(get_segment "${URL}" "/" "-1")
                                             PREFIX=$(get_segment "${PREFIX}" "." "0")
                                           {{ end }}
                                           NAME="{{ if .GetValue "subrepo.name" }}{{ .GetValue "subrepo.name" }}{{ else }}${PREFIX}{{ end }}"
                                           set_project_value "subrepo::${NAME}::prefix" "${PREFIX}"
                                           set_project_value "subrepo::${NAME}::url" "${URL}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepo ${NAME} has been added{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
