# makeAirflowServiceTask
```
  TASK NAME     : makeAirflowServiceTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeAirflowServiceTask.zaruba.yaml
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
  INPUTS        : serviceName
                    DESCRIPTION : Service name (Can be blank)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  serviceContainerName
                    DESCRIPTION : Service's docker container name (Can be blank)
                    PROMPT      : Service's docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  serviceEnvs
                    DESCRIPTION : Service environments, JSON formated.
                                  E.g: {"HTTP_PORT" : "3000", "MODE" : writer"}
                                  
                                  Many applications rely on environment variables to configure their behavior.
                                  You might need to see service's documentation or open environment files (.env, template.env, etc) to see available options.
                                  
                                  If there is no documentation/environment files available, you probably need to run-through the code or ask the developer team.
                    PROMPT      : Service environments, JSON formated. E.g: {"HTTP_PORT" : "3000", "MODE" : "writer"}
                    DEFAULT     : {}
                    VALIDATION  : ^\{.*\}$
                  taskDependencies
                    DESCRIPTION : Task's dependencies, JSON formated.
                                  E.g: ["runMysql", "runRedis"]
                    PROMPT      : Task dependencies, JSON formated. E.g: ["runMysql", "runRedis"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
                  redisServiceName
                    DESCRIPTION : Redis service name (Required)
                    PROMPT      : Redis service name
                    DEFAULT     : redis
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  postgreServiceName
                    DESCRIPTION : Postgre service name (Required)
                    PROMPT      : Postgre service name
                    DEFAULT     : postgre
                    VALIDATION  : ^[a-zA-Z0-9_]+$
  CONFIG        : _setup                       : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                                 IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                                 CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                                 SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                                 SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                                 SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                                 DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                                 REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  _start                       : {{- $d := .Decoration -}}
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
                  afterStart                   : Blank
                  asllowInexistServiceLocation : true
                  beforeStart                  : . "${ZARUBA_HOME}/bash/generateDockerTask.sh"
                                                 REDIS_SERVICE={{ .EscapeShellArg (.GetConfig "redisServiceName") }}
                                                 REDIS_TASK="run$("{{ .ZarubaBin }}" str toPascal "${REDIS_SERVICE}")"
                                                 POSTGRE_SERVICE={{ .EscapeShellArg (.GetConfig "postgreServiceName") }}
                                                 POSTGRE_TASK="run$("{{ .ZarubaBin }}" str toPascal "${POSTGRE_SERVICE}")"
                                                 REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaRedisTask" "${REDIS_TASK}" )"
                                                 REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaRedisService" "${REDIS_SERVICE}" )"
                                                 REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaPostgreTask" "${POSTGRE_TASK}" )"
                                                 REPLACEMENT_MAP="$("{{ .ZarubaBin }}" map set "${REPLACEMENT_MAP}" "zarubaPostgreService" "${POSTGRE_SERVICE}" )"
                                                 if [ "$("{{ .ZarubaBin }}" task isExist ./main.zaruba.yaml "${REDIS_TASK}")" = 0 ]
                                                 then
                                                   echo "create redis task: ${REDIS_TASK}"
                                                   generateDockerTask \
                                                     "${ZARUBA_HOME}/templates/task/docker/redis" "" "${REDIS_SERVICE}" \
                                                     "" "[]" "{}" "[]" "{}"
                                                 fi
                                                 if [ "$("{{ .ZarubaBin }}" task isExist ./main.zaruba.yaml "${POSTGRE_TASK}")" = 0 ]
                                                 then
                                                   echo "create postgre task: ${POSTGRE_TASK}"
                                                   generateDockerTask \
                                                     "${ZARUBA_HOME}/templates/task/docker/postgre" "" "${POSTGRE_SERVICE}" \
                                                     "" "[]" "{}" "[]" "{}" "1"
                                                 fi
                  cmd                          : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                       : -c
                  containerName                : {{ .GetValue "dockerContainerName" }}
                  dependencies                 : {{ .GetValue "taskDependencies" }}
                  finish                       : Blank
                  generatorFunctionName        : generateDockerTask
                  generatorScriptLocation      : ${ZARUBA_HOME}/bash/generateDockerTask.sh
                  imageName                    : {{ .GetValue "dockerImageName" }}
                  includeUtilScript            : . ${ZARUBA_HOME}/bash/util.sh
                  postgreServiceName           : {{ .GetValue "postgreServiceName" }}
                  redisServiceName             : {{ .GetValue "redisServiceName" }}
                  registerRunner               : true
                  replacementMap               : {}
                  serviceEnvs                  : {{ .GetValue "serviceEnvs" }}
                  serviceName                  : {{ .GetValue "serviceName" }}
                  servicePorts                 : {{ .GetValue "servicePorts" }}
                  setup                        : Blank
                  start                        : Blank
                  templateLocation             : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/airflow
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```