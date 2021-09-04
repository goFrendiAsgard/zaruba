# makeFastApiRoute
```
  TASK NAME     : makeFastApiRoute
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeFastApiRoute.zaruba.yaml
  DESCRIPTION   : Make FastAPI route handler
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.showAdv ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : newFastApiServiceName
                    DESCRIPTION : Service name (Required)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  newFastApiModuleName
                    DESCRIPTION : Module name (Required)
                    PROMPT      : Module name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  newFastApiHttpMethod
                    DESCRIPTION : HTTP Method (Required)
                    PROMPT      : HTTP Method
                    OPTIONS     : [ get, post, put, delete ]
                    DEFAULT     : get
                    VALIDATION  : ^[a-z]+$
                  newFastApiUrl
                    DESCRIPTION : URL to be handled (Required)
                    PROMPT      : URL to be handled
                    VALIDATION  : ^[a-zA-Z0-9_\-/\{\}]+$
  CONFIG        : _setup                  : set -e
                                            {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  finish                  : Blank
                  httpMethod              : {{ .GetValue "newFastApiHttpMethod" }}
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  moduleName              : {{ .GetValue "newFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiModule
                  serviceName             : {{ .GetValue "newFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            MODULE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "moduleTemplateLocation") }}
                                            SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                            TASK_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                            URL={{ .EscapeShellArg (.GetConfig "url") }}
                                            HTTP_METHOD={{ .EscapeShellArg (.GetConfig "httpMethod") }}
                                            . ${ZARUBA_HOME}/bash/generate_fast_api_route.sh
                                            generate_fast_api_route \
                                              "${MODULE_TEMPLATE_LOCATION}" \
                                              "${SERVICE_TEMPLATE_LOCATION}" \
                                              "${TASK_TEMPLATE_LOCATION}" \
                                              "${SERVICE_NAME}" \
                                              "${MODULE_NAME}" \
                                              "${URL}" \
                                              "${HTTP_METHOD}" 
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API Route handler created: ${HTTP_METHOD} ${URL} on ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                            echo "You probably need to check the following files:"
                                            echo "- ${SERVICE_NAME}/main.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/controller.py"
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/fastapi
                  url                     : {{ .GetValue "newFastApiUrl" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```