# core.removeDockerContainer
```
  TASK NAME     : core.removeDockerContainer
  LOCATION      : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
  DESCRIPTION   : Remove docker container.
                  Common config:
                    containerName : Container's name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.setupPyUtil, updateProjectLinks ]
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
                  containerName          : Blank
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
                  start                  : {{ $d := .Decoration -}}
                                           CONTAINER="{{ if .GetConfig "containerName" }}{{ .GetConfig "containerName" }}{{ else }}$("${ZARUBA_HOME}/zaruba" getDefaultServiceName "$(pwd)"){{ end }}"
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Stop docker container ${CONTAINER}{{ $d.Normal }}"
                                           stop_container "${CONTAINER}" 
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Remove docker container ${CONTAINER}{{ $d.Normal }}"
                                           remove_container "${CONTAINER}" 
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Docker container ${CONTAINER} removed{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
