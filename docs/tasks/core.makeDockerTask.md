# core.makeDockerTask
```
  TASK NAME     : core.makeDockerTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.makeDockerTask.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.showAdv, core.isProject ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                  : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                            IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                            CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                            SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                            DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                            REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  _start                  : {{- $d := .Decoration -}}
                                            . "{{ .GetConfig "generatorScriptLocation" }}"
                                            {{ .GetConfig "generatorFunctionName" }} \
                                              "${TEMPLATE_LOCATION}" \
                                              "${IMAGE_NAME}" \
                                              "${CONTAINER_NAME}" \
                                              "${SERVICE_NAME}" \
                                              "${SERVICE_PORTS}" \
                                              "${SERVICE_ENVS}" \
                                              "${DEPENDENCIES}" \
                                              "${REPLACEMENT_MAP}" \
                                              "{{ if .IsFalse (.GetConfig "registerRunner") }}0{{ else }}1{{ end }}"
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task created{{ $d.Normal }}"
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  containerName           : {{ .GetValue "dockerContainerName" }}
                  dependencies            : {{ .GetValue "taskDependencies" }}
                  finish                  : Blank
                  generatorFunctionName   : generateDockerTask
                  generatorScriptLocation : ${ZARUBA_HOME}/bash/generateDockerTask.sh
                  imageName               : {{ .GetValue "dockerImageName" }}
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner          : true
                  replacementMap          : {}
                  serviceEnvs             : {{ .GetValue "serviceEnvs" }}
                  serviceName             : {{ .GetValue "serviceName" }}
                  servicePorts            : {{ .GetValue "servicePorts" }}
                  setup                   : Blank
                  start                   : Blank
                  templateLocation        : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```