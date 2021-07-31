# makeServiceTask
```
  TASK NAME     : makeServiceTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeServiceTask.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeServiceTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : generatorServiceLocation
                    DESCRIPTION : Service location, relative to this directory
                    PROMPT      : Service location
                    VALIDATION  : ^.+$
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
                  generatorServiceStartCommand
                    DESCRIPTION : Command to start the service (Required)
                    PROMPT      : Start command
                    VALIDATION  : ^.+$
                  generatorServicePorts
                    DESCRIPTION : Service ports JSON formated.
                                  E.g: ["3001:3000", "8080" , "{{ .GetEnv \"HTTP_PORT\" }}"]
                    PROMPT      : Service ports, JSON formated. E.g: ["3001:3000", "8080", "{{ .GetEnv \"HTTP_PORT\"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
                  generatorServiceDockerImageName
                    DESCRIPTION : Service's docker image name (Can be blank)
                    PROMPT      : Service's docker image name
                    VALIDATION  : ^[a-z0-9_]*$
                  generatorServiceDockerContainerName
                    DESCRIPTION : Service's docker container name (Can be blank)
                    PROMPT      : Service's docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generatorTaskDependencies
                    DESCRIPTION : Task's dependencies, JSON formated.
                                  E.g: ["runMysql", "runRedis"]
                    PROMPT      : Task dependencies, JSON formated. E.g: ["runMysql", "runRedis"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
  CONFIG        : _setup               : set -e
                                         {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start               : Blank
                  afterStart           : Blank
                  beforeStart          : Blank
                  cmd                  : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg               : -c
                  containerName        : {{ .GetValue "generatorServiceDockerContainerName" }}
                  dependencies         : {{ .GetValue "generatorTaskDependencies" }}
                  finish               : Blank
                  imageName            : {{ .GetValue "generatorServiceDockerImageName" }}
                  includeUtilScript    : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  serviceEnvs          : {{ .GetValue "generatorServiceEnvs" }}
                  serviceLocation      : {{ .GetValue "generatorServiceLocation" }}
                  serviceName          : {{ .GetValue "generatorServiceName" }}
                  servicePorts         : {{ .GetValue "generatorServicePorts" }}
                  serviceRunnerVersion : Blank
                  serviceStartCommand  : {{ .GetValue "generatorServiceStartCommand" }}
                  setup                : Blank
                  start                : {{ .GetConfig "start.declareCommonVariables" -}}
                                         {{- $d := .Decoration -}}
                                         . ${ZARUBA_HOME}/scripts/bash/util.sh
                                         TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                         SERVICE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceLocation") }}
                                         SERVICE_START_COMMAND={{ .EscapeShellArg (.GetConfig "serviceStartCommand") }}
                                         SERVICE_RUNNER_VERSION={{ .EscapeShellArg (.GetConfig "serviceRunnerVersion") }}
                                         
                                         DEFAULT_SERVICE_NAME="$({{ .Zaruba }} getServiceName "${SERVICE_LOCATION}")"
                                         SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                         SERVICE_NAME="$(get_value_or_default "${SERVICE_NAME}" "${DEFAULT_SERVICE_NAME}")"
                                         
                                         DEFAULT_IMAGE_NAME="$({{ .Zaruba }} strToKebab "${SERVICE_NAME}")"
                                         IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                         IMAGE_NAME="$(get_value_or_default "${IMAGE_NAME}" "${DEFAULT_IMAGE_NAME}")"
                                         
                                         DEFAULT_CONTAINER_NAME="$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")"
                                         CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                         CONTAINER_NAME="$(get_value_or_default "${CONTAINER_NAME}" "${DEFAULT_CONTAINER_NAME}")"
                                         
                                         PASCAL_SERVICE_NAME="$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")"
                                         KEBAB_SERVICE_NAME="$({{ .Zaruba }} strToKebab "${SERVICE_NAME}")"
                                         SNAKE_SERVICE_NAME="$({{ .Zaruba }} strToSnake "${SERVICE_NAME}")"
                                         UPPER_SNAKE_SERVICE_NAME="$({{ .Zaruba }} strToUpper "${SNAKE_SERVICE_NAME}")"
                                         
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
                                         
                                         DESTINATION="./zaruba-task"
                                         TASK_FILE_NAME="${DESTINATION}/${SERVICE_NAME}.zaruba.yaml"
                                         if [ -f "${TASK_FILE_NAME}" ]
                                         then
                                           echo "{{ $d.Red }}{{ $d.Bold }}file already exist: ${TASK_FILE_NAME}{{ $d.Normal }}"
                                           exit 1
                                         fi
                                         
                                         DEFAULT_PORT_LIST="$({{ .Zaruba }} getPortConfig "${SERVICE_LOCATION}")"
                                         DEFAULT_PORT_CONFIG="$({{ .Zaruba }} join "${DEFAULT_PORT_LIST}" "\n")"
                                         
                                         REPLACEMENT_MAP=$({{ .Zaruba }} setMapElement "{}" \
                                           "zarubaImageName" "${IMAGE_NAME}" \
                                           "zarubaContainerName" "${CONTAINER_NAME}" \
                                           "zarubaServiceName" "${SERVICE_NAME}" \
                                           "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
                                           "zaruba-service-name" "${KEBAB_SERVICE_NAME}" \
                                           "ZARUBA_SERVICE_NAME" "${UPPER_SNAKE_SERVICE_NAME}" \
                                           "zarubaStartCommand" "${SERVICE_START_COMMAND}" \
                                           "zarubaServiceLocation" "$({{ .Zaruba }} getRelativePath "${DESTINATION}" "${SERVICE_LOCATION}")" \
                                           "zarubaRunnerVersion" "${SERVICE_RUNNER_VERSION}" \
                                           "zarubaDefaultPortConfig" "${DEFAULT_PORT_CONFIG}" \
                                         ) 
                                         
                                         {{ .Zaruba }} generate "${TEMPLATE_LOCATION}" "${DESTINATION}" "${REPLACEMENT_MAP}"
                                         
                                         . ${ZARUBA_HOME}/scripts/bash/register_task_file.sh
                                         register_task_file "${TASK_FILE_NAME}" "${SERVICE_NAME}"
                                         
                                         {{ .Zaruba }} addTaskDependency ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${DEPENDENCIES}"
                                         {{ .Zaruba }} setTaskEnv ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${SERVICE_ENVS}"
                                         
                                                 
                                         if [ "$({{ .Zaruba }} getListLength "${SERVICE_PORTS}")" -gt 0 ]
                                         then
                                           PORT_CONFIG_VALUE="$({{ .Zaruba }} join "${SERVICE_PORTS}" )"
                                           PORT_CONFIG="$({{ .Zaruba }} setMapElement "{}" "ports" "$PORT_CONFIG_VALUE" )"
                                           {{ .Zaruba }} setTaskConfig ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${PORT_CONFIG}"
                                         fi
                                         
                                         echo 🎉🎉🎉
                                         echo "{{ $d.Bold }}{{ $d.Yellow }}Service task created{{ $d.Normal }}"
                  templateLocation     : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```