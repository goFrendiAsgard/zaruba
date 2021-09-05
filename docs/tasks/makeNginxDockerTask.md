# makeNginxDockerTask
```
  TASK NAME     : makeNginxDockerTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeNginxDockerTask.zaruba.yaml
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
  CONFIG        : _generatorFunctionName   : generateDockerTask
                  _generatorScriptLocation : ${ZARUBA_HOME}/bash/generateDockerTask.sh
                  _setup                   : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                             IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                             CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                             SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                             SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                             SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                             DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                             REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                  _start                   : {{- $d := .Decoration -}}
                                             . "{{ .GetConfig "_generatorScriptLocation" }}"
                                             {{ .GetConfig "_generatorFunctionName" }} \
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
                  afterStart               : Blank
                  beforeStart              : Blank
                  cmd                      : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                   : -c
                  containerName            : {{ .GetValue "dockerContainerName" }}
                  dependencies             : {{ .GetValue "taskDependencies" }}
                  finish                   : Blank
                  imageName                : {{ .GetValue "dockerImageName" }}
                  includeUtilScript        : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner           : true
                  replacementMap           : {}
                  serviceEnvs              : {{ .GetValue "serviceEnvs" }}
                  serviceName              : {{ .GetValue "serviceName" }}
                  servicePorts             : {{ .GetValue "servicePorts" }}
                  setup                    : Blank
                  start                    : Blank
                  templateLocation         : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/docker/nginx
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```