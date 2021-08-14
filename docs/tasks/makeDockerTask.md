# makeDockerTask
```
  TASK NAME     : makeDockerTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeDockerTask.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeDockerTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : generatorDockerImageName
                    DESCRIPTION : Docker image name (Required)
                    PROMPT      : Docker image name
                    VALIDATION  : ^[a-z0-9_]+$
                  generatorDockerContainerName
                    DESCRIPTION : Docker container name (Can be blank)
                    PROMPT      : Docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generatorServiceName
                    DESCRIPTION : Service name (Can be blank)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generatorServiceEnvs
                    DESCRIPTION : Service environments, JSON formated.
                                  E.g: {"HTTP_PORT" : "3000", "MODE" : writer"}
                                  
                                  Many applications rely on environment variables to configure their behavior.
                                  You might need to see service's documentation or open environment files (.env, template.env, etc) to see available options.
                                  
                                  If there is no documentation/environment files available, you probably need to run-through the code or ask the developer team.
                    PROMPT      : Service environments, JSON formated. E.g: {"HTTP_PORT" : "3000", "MODE" : "writer"}
                    DEFAULT     : {}
                    VALIDATION  : ^\{.*\}$
                  generatorServicePorts
                    DESCRIPTION : Service ports JSON formated.
                                  E.g: ["3001:3000", "8080" , "{{ .GetEnv \"HTTP_PORT\" }}"]
                    PROMPT      : Service ports, JSON formated. E.g: ["3001:3000", "8080", "{{ .GetEnv \"HTTP_PORT\"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
                  generatorTaskDependencies
                    DESCRIPTION : Task's dependencies, JSON formated.
                                  E.g: ["runMysql", "runRedis"]
                    PROMPT      : Task dependencies, JSON formated. E.g: ["runMysql", "runRedis"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
  CONFIG        : _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
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
                  setup             : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                      IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                      CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                      SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                      SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                      SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                      DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                      REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  start             : {{- $d := .Decoration -}}
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
                  templateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```