# addSubrepo
```
  TASK NAME     : addSubrepo
  LOCATION      : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
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
                    DESCRIPTION : Subrepo url (Required)
                    PROMPT      : Subrepo url
                    VALIDATION  : ^.+$
                  subrepo.prefix
                    DESCRIPTION : Subrepo directory name (Can be blank)
                    PROMPT      : Subrepo directory name
                  subrepo.name
                    DESCRIPTION : Subrepo name (Can be blank)
                    PROMPT      : Subrepo name
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
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
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  playBellScript         : echo $'\a'
                  setup                  : Blank
                  start                  : set -e
                                           {{ $d := .Decoration -}}
                                           URL="{{ .GetValue "subrepo.url" }}"
                                           should_not_be_empty "${URL}" "{{ $d.Bold }}{{ $d.Red }}subrepo.url is not defined{{ $d.Normal }}"
                                           {{ if .GetValue "subrepo.prefix" }}
                                             PREFIX="{{ .GetValue "subrepo.prefix" }}"
                                           {{ else }}
                                             {{ $urlSegment := .Split (.GetConfig "subrepoUrl") "/" -}}
                                             {{ $urlSegmentLastIndex := .Subtract (len $urlSegment) 1 -}}
                                             {{ $urlLastSegment := index $urlSegment $urlSegmentLastIndex -}}
                                             {{ $prefix := index (.Split $urlLastSegment ".") 0 -}}
                                             PREFIX="{{ $prefix }}"
                                           {{ end }}
                                           NAME="{{ if .GetValue "subrepo.name" }}{{ .GetValue "subrepo.name" }}{{ else }}${PREFIX}{{ end }}"
                                           zaruba setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::prefix" "${PREFIX}"
                                           zaruba setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::url" "${URL}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepo ${NAME} has been added{{ $d.Normal }}"
                  subrepoName            : {{ .GetValue "subrepo.name" }}
                  subrepoPrefix          : {{ .GetValue "subrepo.prefix" }}
                  subrepoUrl             : {{ .GetValue "subrepo.url" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
