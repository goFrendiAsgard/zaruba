# makeFastApiRpcHandler
```
  TASK NAME     : makeFastApiRpcHandler
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeFastApiRpcHandler.zaruba.yaml
  DESCRIPTION   : Make FastAPI Route
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
                  newFastApiRpcName
                    DESCRIPTION : RPC name (Required)
                    PROMPT      : RPC name
                    VALIDATION  : ^[a-zA-Z0-9_\-\.]+$
  CONFIG        : _setup                  : set -e
                                            {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  finish                  : Blank
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  moduleName              : {{ .GetValue "newFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiModule
                  rpcName                 : {{ .GetValue "newFastApiRpcName" }}
                  serviceName             : {{ .GetValue "newFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            MODULE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "moduleTemplateLocation") }}
                                            SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                            TASK_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                            RPC_NAME={{ .EscapeShellArg (.GetConfig "rpcName") }}
                                            . ${ZARUBA_HOME}/bash/generate_fast_api_rpc_handler.sh
                                            generate_fast_api_rpc_handler \
                                              "${MODULE_TEMPLATE_LOCATION}" \
                                              "${SERVICE_TEMPLATE_LOCATION}" \
                                              "${TASK_TEMPLATE_LOCATION}" \
                                              "${SERVICE_NAME}" \
                                              "${MODULE_NAME}" \
                                              "${RPC_NAME}"
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API RPC handler created: ${RPC_NAME} on ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                            echo "You probably need to check the following files:"
                                            echo "- ${SERVICE_NAME}/main.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/controller.py"
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/fastapi
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```