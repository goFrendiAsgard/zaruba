# addProjectLink
```
  TASK NAME     : addProjectLink
  LOCATION      : ${ZARUBA_HOME}/scripts/task.addProjectLink.zaruba.yaml
  DESCRIPTION   : Add link.
                  TIPS: To update links, you should perform `zaruba please updateProjectLinks`
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
  INPUTS        : linkFrom
                    DESCRIPTION : Link source (Required)
                    PROMPT      : Source
                    VALIDATION  : ^.+$
                  linkTo
                    DESCRIPTION : Link destination (Required)
                    PROMPT      : Destination
                    VALIDATION  : ^.+$
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
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  linkFrom               : {{ .GetValue "linkFrom" }}
                  linkTo                 : {{ .GetValue "linkTo" }}
                  setup                  : Blank
                  start                  : {{ $d := .Decoration -}}
                                           "${ZARUBA_HOME}/zaruba" setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "link::{{ .GetConfig "linkTo" }}" "{{ .GetConfig "linkFrom" }}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Link ${SOURCE} -> ${DESTINATION} has been added{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```