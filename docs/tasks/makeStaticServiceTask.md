# makeStaticServiceTask
```
  TASK NAME     : makeStaticServiceTask
  LOCATION      : /scripts/tasks/makeStaticServiceTask.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeServiceTask ]
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
  INPUTS        : serviceLocation
                    DESCRIPTION : Service location, relative to this directory
                    PROMPT      : Service location
                    VALIDATION  : ^.+$
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
                  servicePorts
                    DESCRIPTION : Service ports JSON formated.
                                  E.g: ["3001:3000", "8080" , "{{ .GetEnv \"HTTP_PORT\" }}"]
                    PROMPT      : Service ports, JSON formated. E.g: ["3001:3000", "8080", "{{ .GetEnv \"HTTP_PORT\"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
                  taskDependencies
                    DESCRIPTION : Task's dependencies, JSON formated.
                                  E.g: ["runMysql", "runRedis"]
                    PROMPT      : Task dependencies, JSON formated. E.g: ["runMysql", "runRedis"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
                  serviceImageName
                    DESCRIPTION : Service's docker image name (Can be blank)
                    PROMPT      : Service's docker image name
                    VALIDATION  : ^[a-z0-9_]*$
                  serviceContainerName
                    DESCRIPTION : Service's docker container name (Can be blank)
                    PROMPT      : Service's docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
  CONFIG        : _finish                     : {{- $d := .Decoration -}}
                                                "{{ .ZarubaBin }}" project syncEnv "./main.zaruba.yaml"
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Service task for ${SERVICE_NAME} has been created{{ $d.Normal }}"
                  _setup                      : set -e
                                                {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                                                . "${ZARUBA_HOME}/bash/generatorUtil.sh"
                                                {{ if .IsTrue (.GetConfig "allowInexistServiceLocation") -}}
                                                mkdir -p "{{ .GetConfig "serviceLocation" }}"
                                                {{ end -}}
                                                TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "templateLocation") }}
                                                SERVICE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "serviceLocation") }}
                                                SERVICE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "serviceName") }}
                                                IMAGE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "imageName") }}
                                                CONTAINER_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "containerName") }}
                                                SERVICE_START_COMMAND={{ .Util.Str.EscapeShellArg (.GetConfig "serviceStartCommand") }}
                                                SERVICE_RUNNER_VERSION={{ .Util.Str.EscapeShellArg (.GetConfig "serviceRunnerVersion") }}
                                                SERVICE_PORTS={{ .Util.Str.EscapeShellArg (.GetConfig "servicePorts") }}
                                                SERVICE_ENVS={{ .Util.Str.EscapeShellArg (.GetConfig "serviceEnvs") }}
                                                DEPENDENCIES={{ .Util.Str.EscapeShellArg (.GetConfig "dependencies") }}
                                                REPLACEMENT_MAP={{ .Util.Str.EscapeShellArg (.GetConfig "replacementMap") }}
                                                # ensure SERVICE_NAME is not empty
                                                SERVICE_NAME="$(getAppName "${SERVICE_NAME}" "${SERVICE_LOCATION}")"
                                                # ensure IMAGE_NAME is not empty
                                                IMAGE_NAME="$(getServiceImageName "${IMAGE_NAME}" "${SERVICE_NAME}")"
                                                # ensure CONTAINER_NAME is not empty
                                                CONTAINER_NAME="$(getServiceContainerName "${CONTAINER_NAME}" "${SERVICE_NAME}")"
                  _start                      : __PWD="$(pwd)"
                                                . "{{ .GetConfig "generatorScriptLocation" }}"
                                                {{ .GetConfig "generatorFunctionName" }} \
                                                {{ .GetConfig "generatorFunctionArgs" }}
                                                cd "${__PWD}"
                  afterStart                  : Blank
                  allowInexistServiceLocation : false
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  containerName               : {{ .GetValue "serviceContainerName" }}
                  dependencies                : {{ .GetValue "taskDependencies" }}
                  finish                      : Blank
                  generatorFunctionArgs       : "${TEMPLATE_LOCATION}" \
                                                "${SERVICE_LOCATION}" \
                                                "${SERVICE_NAME}" \
                                                "${IMAGE_NAME}" \
                                                "${CONTAINER_NAME}" \
                                                "${SERVICE_START_COMMAND}" \
                                                "${SERVICE_RUNENR_VERSION}" \
                                                "${SERVICE_PORTS}" \
                                                "${SERVICE_ENVS}" \
                                                "${DEPENDENCIES}" \
                                                "${REPLACEMENT_MAP}" \
                                                "{{ if .IsFalse (.GetConfig "registerRunner") }}0{{ else }}1{{ end }}"
                  generatorFunctionName       : generateServiceTask
                  generatorScriptLocation     : ${ZARUBA_HOME}/bash/generateServiceTask.sh
                  imageName                   : {{ .GetValue "serviceImageName" }}
                  includeUtilScript           : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner              : true
                  replacementMap              : {}
                  serviceEnvs                 : {{ .GetValue "serviceEnvs" }}
                  serviceLocation             : {{ .GetValue "serviceLocation" }}
                  serviceName                 : {{ .GetValue "serviceName" }}
                  servicePorts                : {{ .GetValue "servicePorts" }}
                  serviceRunnerVersion        : Blank
                  serviceStartCommand         : {{ .GetValue "startCommand" }}
                  setup                       : Blank
                  start                       : Blank
                  templateLocation            : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/static
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```