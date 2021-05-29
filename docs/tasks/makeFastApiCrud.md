# makeFastApiCrud
```
  TASK NAME     : makeFastApiCrud
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.generator.zaruba.yaml
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
  INPUTS        : generator.fastApi.service.name
                    DESCRIPTION : Service name (Required)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.module.name
                    DESCRIPTION : Module name (Required)
                    PROMPT      : Module name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.crud.entity
                    DESCRIPTION : Entity name (Required)
                                  Usually plural word (e.g: books, articles)
                    PROMPT      : Entity name
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.crud.fields
                    DESCRIPTION : Field names, comma separated.
                                  E.g: name,address
                                  The following fields are included by default:
                                  - id
                                  - created_at
                                  - updated_at
                    PROMPT      : Field names
                    VALIDATION  : ^[a-zA-Z0-9_,]*$
  CONFIG        : _setup                  : set -e
                                            {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                            {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                  : Blank
                  afterStart              : Blank
                  beforeStart             : Blank
                  cmd                     : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                  : -c
                  createModuleScript      : {{- $d := .Decoration -}}
                                            {{ .GetConfig "createServiceScript" }}
                                            if [ ! -d "./{{ .GetConfig "serviceName" }}/{{ .GetConfig "moduleName" }}" ]
                                            then
                                              MODULE_TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "moduleTemplateLocation") }}
                                              SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                              MODULE_NAME={{ .SingleQuoteShellValue (.GetConfig "moduleName") }}
                                              should_be_dir "./${SERVICE_NAME}" "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME} directory should be exist{{ $d.Normal }}"
                                              create_fast_module "template_location=${MODULE_TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "module_name=${MODULE_NAME}"
                                            fi
                  createServiceScript     : {{- $d := .Decoration -}}
                                            if [ ! -d "./{{ .GetConfig "serviceName" }}" ]
                                            then
                                              SERVICE_TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "serviceTemplateLocation") }}
                                              SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                              create_fast_service "template_location=${SERVICE_TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}"
                                              if [ -f "./main.zaruba.yaml" ]
                                              then
                                                if [ ! -d "./shared-libs/python/helpers" ]
                                                then
                                                  mkdir -p "./shared-libs/python/helpers"
                                                  cp -rnT "./${SERVICE_NAME}/helpers" "./shared-libs/python/helpers"
                                                fi
                                                add_link "shared-libs/python/helpers" "${SERVICE_NAME}/helpers"
                                                link_resource "shared-libs/python/helpers" "${SERVICE_NAME}/helpers"
                                              fi
                                            fi
                  entityName              : {{ .GetValue "generator.fastApi.crud.entity" }}
                  fieldNames              : {{ .GetValue "generator.fastApi.crud.fields" }}
                  finish                  : Blank
                  includeBootstrapScript  : if [ -f "${HOME}/.profile" ]
                                            then
                                                . "${HOME}/.profile"
                                            fi
                                            if [ -f "${HOME}/.bashrc" ]
                                            then
                                                . "${HOME}/.bashrc"
                                            fi
                                            BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                            . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript       : . "${ZARUBA_HOME}/scripts/util.sh"
                  moduleName              : {{ .GetValue "generator.fastApi.module.name" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiModule
                  playBellScript          : echo $'\a'
                  serviceName             : {{ .GetValue "generator.fastApi.service.name" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            {{ .GetConfig "createModuleScript" }}
                                            TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "templateLocation") }}
                                            SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .SingleQuoteShellValue (.GetConfig "moduleName") }}
                                            ENTITY_NAME={{ .SingleQuoteShellValue (.GetConfig "entityName") }}
                                            FIELD_NAMES={{ .SingleQuoteShellValue (.GetConfig "fieldNames") }}
                                            should_be_dir "./${SERVICE_NAME}/${MODULE_NAME}" "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME} directory should be exist{{ $d.Normal }}"
                                            create_fast_crud "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "module_name=${MODULE_NAME}" "entity_name=${ENTITY_NAME}" "field_names=${FIELD_NAMES}"
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
                  templateLocation        : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiCrud
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
