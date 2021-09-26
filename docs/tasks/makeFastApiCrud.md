# makeFastApiCrud
```
  TASK NAME     : makeFastApiCrud
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeFastApiCrud.zaruba.yaml
  DESCRIPTION   : Make FastAPI crud
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.showAdv ]
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
  INPUTS        : fastApiServiceName
                    DESCRIPTION : Service name (Required)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  fastApiModuleName
                    DESCRIPTION : Module name (Required)
                    PROMPT      : Module name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  fastApiCrudEntity
                    DESCRIPTION : Entity name (Required)
                                  Usually plural word (e.g: books, articles)
                    PROMPT      : Entity name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  fastApiCrudFields
                    DESCRIPTION : Field names, JSON formated.
                                  E.g: ["name", "address"]
                    PROMPT      : Field names, JSON formated. E.g: ["name", "address"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
  CONFIG        : _finish                 : Blank
                  _setup                  : set -e
                                            {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  crudTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiCrud
                  entityName              : {{ .GetValue "fastApiCrudEntity" }}
                  fieldNames              : {{ .GetValue "fastApiCrudFields" }}
                  finish                  : Blank
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  moduleName              : {{ .GetValue "fastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiModule
                  serviceName             : {{ .GetValue "fastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            CRUD_TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "crudTemplateLocation") }}
                                            MODULE_TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "moduleTemplateLocation") }}
                                            SERVICE_TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                            TASK_TEMPLATE_LOCATION={{ .Util.Str.EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                            SERVICE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "moduleName") }}
                                            ENTITY_NAME={{ .Util.Str.EscapeShellArg (.GetConfig "entityName") }}
                                            FIELD_NAMES={{ .Util.Str.EscapeShellArg (.GetConfig "fieldNames") }}
                                            . ${ZARUBA_HOME}/bash/generateFastApiCrudHandler.sh
                                            generateFastApiCrudHandler \
                                              "${CRUD_TEMPLATE_LOCATION}" \
                                              "${MODULE_TEMPLATE_LOCATION}" \
                                              "${SERVICE_TEMPLATE_LOCATION}" \
                                              "${TASK_TEMPLATE_LOCATION}" \
                                              "${SERVICE_NAME}" \
                                              "${MODULE_NAME}" \
                                              "${ENTITY_NAME}" \
                                              "${FIELD_NAMES}"
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API module created: ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                            echo "You probably need to check the following files:"
                                            echo "- ${SERVICE_NAME}/main.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/controller.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/handle<Entity>Event.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/handle<Entity>Route.py"
                                            echo "- ${SERVICE_NAME}/repos/<entity>.py"
                                            echo "- ${SERVICE_NAME}/repos/db<Entity>.py"
                                            echo "- ${SERVICE_NAME}/schemas/<entity>.py"
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/fastapi
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```