# makeRedisDockerTask
```
  TASK NAME     : makeRedisDockerTask
  LOCATION      : /scripts/tasks/makeRedisDockerTask.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeDockerTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  INPUTS        : dockerContainerName
                    DESCRIPTION : Docker container name (Can be blank)
                    PROMPT      : Docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  serviceName
                    DESCRIPTION : Service name (Can be blank)
                    PROMPT      : Service name
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
  CONFIG        : _finish                 : {{- $d := .Decoration -}}
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task for ${SERVICE_NAME} has been created{{ $d.Normal }}"
                  _setup                  : set -e
                                            {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                                            . "${ZARUBA_HOME}/bash/generatorUtil.sh"
                                            TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "templateLocation") }}
                                            IMAGE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "imageName") }}
                                            CONTAINER_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "containerName") }}
                                            SERVICE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "serviceName") }}
                                            SERVICE_PORTS={{ .Util.Str.EscapeShellArg (.GetConfig "servicePorts") }}
                                            SERVICE_ENVS={{ .Util.Str.EscapeShellArg (.GetConfig "serviceEnvs") }}
                                            DEPENDENCIES={{ .Util.Str.EscapeShellArg (.GetConfig "dependencies") }}
                                            REPLACEMENT_MAP={{ .Util.Str.EscapeShellArg (.GetConfig "replacementMap") }}
                                            # ensure CONTAINER_NAME is not empty
                                            CONTAINER_NAME="$(getDockerContainerName "${CONTAINER_NAME}" ${IMAGE_NAME} ${_TEMPLATE_LOCATION})"
                                            # ensure SERVICE_NAME is not empty
                                            SERVICE_NAME="$(getDockerServiceName "${SERVICE_NAME}" "${CONTAINER_NAME}")"
                  _start                  : __PWD="$(pwd)"
                                            . "{{ .GetConfig "generatorScriptLocation" }}"
                                            {{ .GetConfig "generatorFunctionName" }} \
                                            {{ .GetConfig "generatorFunctionArgs" }}
                                            cd "${__PWD}"
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  containerName           : {{ .GetValue "dockerContainerName" }}
                  dependencies            : {{ .GetValue "taskDependencies" }}
                  finish                  : Blank
                  generatorFunctionArgs   : "${TEMPLATE_LOCATION}" \
                                            "${IMAGE_NAME}" \
                                            "${CONTAINER_NAME}" \
                                            "${SERVICE_NAME}" \
                                            "${SERVICE_PORTS}" \
                                            "${SERVICE_ENVS}" \
                                            "${DEPENDENCIES}" \
                                            "${REPLACEMENT_MAP}" \
                                            "{{ if .IsFalse (.GetConfig "registerRunner") }}0{{ else }}1{{ end }}"
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
                  start                   : . "${ZARUBA_HOME}/bash/setProjectValueUnlessExist.sh"
                                            setProjectValueUnlessExist redisServiceName "${SERVICE_NAME}"
                  templateLocation        : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/redis
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```