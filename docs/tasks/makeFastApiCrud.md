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
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : generatorFastApiServiceName
                    DESCRIPTION : Service name (Required)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generatorFastApiModuleName
                    DESCRIPTION : Module name (Required)
                    PROMPT      : Module name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generatorFastApiCrudEntity
                    DESCRIPTION : Entity name (Required)
                                  Usually plural word (e.g: books, articles)
                    PROMPT      : Entity name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generatorFastApiCrudFields
                    DESCRIPTION : Field names, JSON formated.
                                  E.g: ["name", "address"]
                    PROMPT      : Field names, JSON formated. E.g: ["name", "address"]
                    DEFAULT     : []
                    VALIDATION  : ^\[.*\]$
  CONFIG        : _setup                  : set -e
                                            {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  crudTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiCrud
                  entityName              : {{ .GetValue "generatorFastApiCrudEntity" }}
                  fieldNames              : {{ .GetValue "generatorFastApiCrudFields" }}
                  finish                  : Blank
                  includeUtilScript       : . ${ZARUBA_HOME}/bash/util.sh
                  moduleName              : {{ .GetValue "generatorFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiModule
                  serviceName             : {{ .GetValue "generatorFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            CRUD_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "crudTemplateLocation") }}
                                            MODULE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "moduleTemplateLocation") }}
                                            SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                            TASK_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                            ENTITY_NAME={{ .EscapeShellArg (.GetConfig "entityName") }}
                                            FIELD_NAMES={{ .EscapeShellArg (.GetConfig "fieldNames") }}
                                            . ${ZARUBA_HOME}/bash/generate_fast_api_crud_handler.sh
                                            generate_fast_api_crud_handler \
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