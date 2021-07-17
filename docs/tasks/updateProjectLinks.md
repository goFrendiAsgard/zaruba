# updateProjectLinks
```
  TASK NAME     : updateProjectLinks
  LOCATION      : /home/gofrendi/zaruba/scripts/task.updateProjectLinks.zaruba.yaml
  DESCRIPTION   : Update "links" in your project. Very useful if you have multiple apps sharing some parts of code
                  USAGE:
                    zaruba please updateProjectLinks
                    zaruba please updateProjectLinks "link::fibo/css=common-css"
                    zaruba please updateProjectLinks "link::app/css=common-css"
                  ARGUMENTS
                    link::<destination> : Location of the shared code
                  TIPS:
                    It is recommended to put `link` arguments in `default.values.yaml`.
                    In order to do that, you can invoke `zaruba please addProjectLink <linkFrom=source-location> <linkTo=destination-location>`
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
                  setup                  : Blank
                  start                  : {{ $d := .Decoration -}}
                                           {{ $this := . -}}
                                           {{ $workPath := .WorkPath }}
                                           {{ $destinations := .GetSubValueKeys "link" -}}
                                           {{ range $index, $destination := $destinations -}}
                                             {{ $source := $this.GetValue "link" $destination -}}
                                             {{ $absSource := $this.GetWorkPath $source -}}
                                             {{ $absDestination := $this.GetWorkPath $destination -}}
                                             link_resource "{{ $absSource }}" "{{ $absDestination }}"
                                           {{ end -}}
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Links updated{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
