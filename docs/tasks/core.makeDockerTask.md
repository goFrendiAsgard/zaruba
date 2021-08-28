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
  CONFIG        : _setup            : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                      IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                      CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                      SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                      SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                      SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                      DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                      REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  _start            : {{- $d := .Decoration -}}
                                      . "${ZARUBA_HOME}/bash/generate_docker_task.sh"
                                      generate_docker_task \
                                        "${TEMPLATE_LOCATION}" \
                                        "${IMAGE_NAME}" \
                                        "${CONTAINER_NAME}" \
                                        "${SERVICE_NAME}" \
                                        "${SERVICE_PORTS}" \
                                        "${SERVICE_ENVS}" \
                                        "${DEPENDENCIES}" \
                                        "${REPLACEMENT_MAP}"
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task created{{ $d.Normal }}"
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  containerName     : {{ .GetValue "generatorDockerContainerName" }}
                  dependencies      : {{ .GetValue "generatorTaskDependencies" }}
                  finish            : Blank
                  imageName         : {{ .GetValue "generatorDockerImageName" }}
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  replacementMap    : {}
                  serviceEnvs       : {{ .GetValue "generatorServiceEnvs" }}
                  serviceName       : {{ .GetValue "generatorServiceName" }}
                  servicePorts      : {{ .GetValue "generatorServicePorts" }}
                  setup             : Blank
                  start             : Blank
                  templateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```