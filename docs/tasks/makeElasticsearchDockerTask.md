# makeElasticsearchDockerTask
```
  TASK NAME     : makeElasticsearchDockerTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeElasticsearchDockerTask.zaruba.yaml
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
  INPUTS        : generatorDockerContainerName
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
                  generatorTaskDependencies
                    DESCRIPTION : Task's dependencies, JSON formated.
                                  E.g: ["runMysql", "runRedis"]
                    PROMPT      : Task dependencies, JSON formated. E.g: ["runMysql", "runRedis"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
  CONFIG        : _setup                       : set -e
                                                 alias zaruba=${ZARUBA_HOME}/zaruba
                                                 {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                 {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                       : Blank
                  afterStart                   : Blank
                  beforeStart                  : Blank
                  cmd                          : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                       : -c
                  containerName                : {{ .GetValue "generatorDockerContainerName" }}
                  dependencies                 : {{ .GetValue "generatorTaskDependencies" }}
                  finish                       : Blank
                  imageName                    : {{ .GetValue "generatorDockerImageName" }}
                  includeBootstrapScript       : if [ -f "${HOME}/.profile" ]
                                                 then
                                                     . "${HOME}/.profile"
                                                 fi
                                                 if [ -f "${HOME}/.bashrc" ]
                                                 then
                                                     . "${HOME}/.bashrc"
                                                 fi
                                                 BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                                 . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript            : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  serviceEnvs                  : {{ .GetValue "generatorServiceEnvs" }}
                  serviceName                  : {{ .GetValue "generatorServiceName" }}
                  servicePorts                 : {{ .GetValue "generatorServicePorts" }}
                  setup                        : Blank
                  start                        : {{- $d := .Decoration -}}
                                                 {{ .GetConfig "start.declareVariables" }}
                                                 {{ .GetConfig "start.createReplacementMap" }}
                                                 {{ .Zaruba }} generate "${TEMPLATE_LOCATION}" "${DESTINATION}" "${REPLACEMENT_MAP}"
                                                 {{ .GetConfig "start.linkToProject" }}
                                                 {{ .Zaruba }} addTaskDependency ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${DEPENDENCIES}"
                                                 {{ .Zaruba }} setTaskEnv ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${SERVICE_ENVS}"
                                                 if [ "$({{ .Zaruba }} listLength "${SERVICE_PORTS}")" -gt 0 ]
                                                 then
                                                   PORT_CONFIG_VALUE="$({{ .Zaruba }} join "${SERVICE_PORTS}" )"
                                                   PORT_CONFIG="$({{ .Zaruba }} mapSet "{}" "ports" "$PORT_CONFIG_VALUE" )"
                                                   {{ .Zaruba }} setTaskConfig ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${PORT_CONFIG}"
                                                 fi
                                                 echo 🎉🎉🎉
                                                 echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task created{{ $d.Normal }}"
                  start.createReplacementMap   : REPLACEMENT_MAP=$({{ .Zaruba }} mapSet "{}" \
                                                   "zarubaImageName" "${IMAGE_NAME}" \
                                                   "zarubaContainerName" "${CONTAINER_NAME}" \
                                                   "zarubaServiceName" "${SERVICE_NAME}" \
                                                   "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
                                                 )
                  start.declareCommonVariables : {{- $d := .Decoration -}}
                                                 TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                                 DESTINATION="./zaruba-tasks"
                                                 SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                                 if [ "$({{ .Zaruba}} isValidMap "$SERVICE_ENVS")" -eq 0 ]
                                                 then
                                                   echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_ENVS} is not a valid map{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                                                 SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                                 if [ "$({{ .Zaruba}} isValidList "$SERVICE_PORTS")" -eq 0 ]
                                                 then
                                                   echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_PORTS} is not a valid port{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                                                 DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                                 if [ "$({{ .Zaruba}} isValidList "$DEPENDENCIES")" -eq 0 ]
                                                 then
                                                   echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_PORTS} is not a valid port{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  start.declareVariables       : {{ .GetConfig "start.declareCommonVariables" -}}
                                                 {{- $d := .Decoration -}}
                                                 IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                                 CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                                 CONTAINER_NAME=$({{ .Zaruba }} getValue "${CONTAINER_NAME}" \
                                                   "$({{ .Zaruba }} strToCamel "${IMAGE_NAME}")" \
                                                 )
                                                 SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                                 SERVICE_NAME=$({{ .Zaruba }} getValue "${SERVICE_NAME}" \
                                                   "${CONTAINER_NAME}" \
                                                 )
                                                 PASCAL_SERVICE_NAME="$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")"
                                                 KEBAB_SERVICE_NAME="$({{ .Zaruba }} strToKebab "${SERVICE_NAME}")"
                                                 TASK_FILE_NAME="${DESTINATION}/${SERVICE_NAME}.zaruba.yaml"
                                                 if [ -f "${TASK_FILE_NAME}" ]
                                                 then
                                                   echo "{{ $d.Red }}{{ $d.Bold }}file already exist: ${TASK_FILE_NAME}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  start.linkToProject          : {{ .Zaruba }} includeFileToProject "./main.zaruba.yaml" "${TASK_FILE_NAME}"
                                                 {{ .Zaruba }} syncProjectEnvFiles "./main.zaruba.yaml"
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "build${PASCAL_SERVICE_NAME}Image")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "buildImage"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "buildImage" "[\"build${PASCAL_SERVICE_NAME}Image\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "pull${PASCAL_SERVICE_NAME}Image")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "pullImage"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "pullImage" "[\"pull${PASCAL_SERVICE_NAME}Image\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "push${PASCAL_SERVICE_NAME}Image")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "pushImage"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "pushImage" "[\"push${PASCAL_SERVICE_NAME}Image\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "run${PASCAL_SERVICE_NAME}")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "run"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "run" "[\"run${PASCAL_SERVICE_NAME}\"]"
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "runContainer"
                                                   if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "run${PASCAL_SERVICE_NAME}Container")" -eq 1 ]
                                                   then
                                                     {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "runContainer" "[\"run${PASCAL_SERVICE_NAME}Container\"]"
                                                   else
                                                     {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "runContainer" "[\"run${PASCAL_SERVICE_NAME}\"]"
                                                   fi
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "stop${PASCAL_SERVICE_NAME}Container")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "stopContainer"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "stopContainer" "[\"stop${PASCAL_SERVICE_NAME}Container\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "remove${PASCAL_SERVICE_NAME}Container")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "removeContainer"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "removeContainer" "[\"remove${PASCAL_SERVICE_NAME}Container\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "helmInstall${PASCAL_SERVICE_NAME}")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "helmInstall"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "helmInstall" "[\"helmInstall${PASCAL_SERVICE_NAME}\"]"
                                                 fi
                                                 
                                                 if [ "$({{ .Zaruba }} isTaskExist "./main.zaruba.yaml" "helmUninstall${PASCAL_SERVICE_NAME}")" -eq 1 ]
                                                 then
                                                   {{ .Zaruba }} createTaskIfNotExist "./main.zaruba.yaml" "helmUninstall"
                                                   {{ .Zaruba }} addTaskDependency "./main.zaruba.yaml" "helmUninstall" "[\"helmUninstall${PASCAL_SERVICE_NAME}\"]"
                                                 fi
                  templateLocation             : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/docker/elasticsearch
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```