# makeAirflowDockerTask
```
  TASK NAME     : makeAirflowDockerTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeAirflowDockerTask.zaruba.yaml
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
                  generatorRedisTask
                    DESCRIPTION : Redis task (Required)
                    PROMPT      : Redis task
                    DEFAULT     : runRedis
                    VALIDATION  : ^run[a-zA-Z0-9_]+$
                  generatorPostgreTask
                    DESCRIPTION : Postgre task (Required)
                    PROMPT      : Postgre task
                    DEFAULT     : runPostgre
                    VALIDATION  : ^run[a-zA-Z0-9_]+$
  CONFIG        : _setup               : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                         IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                         CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                         SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                         SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                         SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                         DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                         REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  _start               : {{- $d := .Decoration -}}
                                         . "${ZARUBA_HOME}/bash/generate_docker_task.sh"
                                         generate_docker_task \
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
                  afterStart           : Blank
                  beforeStart          : . "${ZARUBA_HOME}/bash/generate_docker_task.sh"
                                         . "${ZARUBA_HOME}/bash/get_service_name_by_task_name.sh"
                                         REDIS_TASK={{ .EscapeShellArg (.GetConfig "generatorRedisTask") }}
                                         REDIS_SERVICE="$(get_service_name_by_task_name "${REDIS_TASK}")"
                                         POSTGRE_TASK={{ .EscapeShellArg (.GetConfig "generatorPostgreTask") }}
                                         POSTGRE_SERVICE="$(get_service_name_by_task_name "${POSTGRE_TASK}")"
                                         REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaRedisTask" "${REDIS_TASK}" )"
                                         REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaRedisService" "${REDIS_SERVICE}" )"
                                         REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaPostgreTask" "${POSTGRE_TASK}" )"
                                         REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaPostgreService" "${POSTGRE_SERVICE}" )"
                                         if [ "$("{{ .ZarubaBin }}" task isExist ./main.zaruba.yaml "${REDIS_SERVICE}")" = 0 ]
                                         then
                                           echo "create redis task: ${REDIS_SERVICE}"
                                           generate_docker_task \
                                             "${ZARUBA_HOME}/templates/task/docker/redis" "" "${REDIS_SERVICE}" \
                                             "" "[]" "{}" "[]" "{}"
                                         fi
                                         if [ "$("{{ .ZarubaBin }}" task isExist ./main.zaruba.yaml "${POSTGRE_SERVICE}")" = 0 ]
                                         then
                                           echo "create postgre task: ${POSTGRE_SERVICE}"
                                           generate_docker_task \
                                             "${ZARUBA_HOME}/templates/task/docker/postgre" "" "${POSTGRE_SERVICE}" \
                                             "" "[]" "{}" "[]" "{}" "1"
                                         fi
                  cmd                  : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg               : -c
                  containerName        : {{ .GetValue "generatorDockerContainerName" }}
                  dependencies         : {{ .GetValue "generatorTaskDependencies" }}
                  finish               : Blank
                  generatorPostgreTask : {{ .GetValue "generatorPostgreTask" }}
                  generatorRedisTask   : {{ .GetValue "generatorRedisTask" }}
                  imageName            : {{ .GetValue "generatorDockerImageName" }}
                  includeUtilScript    : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner       : true
                  replacementMap       : {}
                  serviceEnvs          : {{ .GetValue "generatorServiceEnvs" }}
                  serviceName          : {{ .GetValue "generatorServiceName" }}
                  servicePorts         : {{ .GetValue "generatorServicePorts" }}
                  setup                : Blank
                  start                : Blank
                  templateLocation     : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/airflow
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```