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
                  generatorFastApiCreateTask
                    DESCRIPTION : Create service task if not exist.
                    PROMPT      : Create service task if not exist
                    OPTIONS     : [ yes, no ]
                    DEFAULT     : no
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
                                            alias zaruba=${ZARUBA_HOME}/zaruba
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
                                              MODULE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "moduleTemplateLocation") }}
                                              SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                              MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                              should_be_dir "./${SERVICE_NAME}" "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME} directory should be exist{{ $d.Normal }}"
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}Creating Fast API module: ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                              PASCAL_SERVICE_NAME=$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")
                                              CAMEL_SERVICE_NAME=$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")
                                              PASCAL_MODULE_NAME=$({{ .Zaruba }} strToPascal "${MODULE_NAME}")
                                              CAMEL_MODULE_NAME=$({{ .Zaruba }} strToCamel "${MODULE_NAME}")
                                              SNAKE_MODULE_NAME=$({{ .Zaruba }} strToSnake "${MODULE_NAME}")
                                              REPLACEMENT_MAP=$({{ .Zaruba }} mapSet "{}" \
                                                "zarubaServiceName" "${CAMEL_SERVICE_NAME}" \
                                                "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
                                                "zarubaModuleName" "${CAMEL_MODULE_NAME}" \
                                                "ZarubaModuleName" "${PASCAL_MODULE_NAME}" \
                                                "zaruba_module_name" "${SNAKE_MODULE_NAME}" \
                                              )
                                              {{ .Zaruba }} generate "${MODULE_TEMPLATE_LOCATION}/zarubaServiceName" "${CAMEL_SERVICE_NAME}" "${REPLACEMENT_MAP}"
                                              # get main.py lines
                                              MAIN_LINES=$({{ .Zaruba }} readLines "${CAMEL_SERVICE_NAME}/main.py")
                                              # import module
                                              IMPORT_MODULE_PARTIAL=$(cat "${MODULE_TEMPLATE_LOCATION}/partials/import_module.py")
                                              IMPORT_MODULE_PARTIAL=$({{ .Zaruba }} strReplace "${IMPORT_MODULE_PARTIAL}" "${REPLACEMENT_MAP}")
                                              IMPORT_MODULE_LINES=$({{ .Zaruba }} split "${IMPORT_MODULE_PARTIAL}")
                                              # load module
                                              LOAD_MODULE_PARTIAL=$(cat "${MODULE_TEMPLATE_LOCATION}/partials/load_module.py")
                                              LOAD_MODULE_PARTIAL=$({{ .Zaruba }} strReplace "${LOAD_MODULE_PARTIAL}" "${REPLACEMENT_MAP}")
                                              LOAD_MODULE_LINES=$({{ .Zaruba }} split "${LOAD_MODULE_PARTIAL}")
                                              # update main.py
                                              MAIN_LINES=$({{ .Zaruba }} listMerge "${IMPORT_MODULE_LINES}" "${MAIN_LINES}" "${LOAD_MODULE_LINES}")
                                              {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/main.py" "${MAIN_LINES}"
                                            fi
                  createServiceScript     : {{- $d := .Decoration -}}
                                            if [ ! -d "./{{ .GetConfig "serviceName" }}" ]
                                            then
                                              SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                              SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                              PASCAL_SERVICE_NAME=$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")
                                              CAMEL_SERVICE_NAME=$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")
                                              REPLACEMENT_MAP=$({{ .Zaruba }} mapSet "{}" \
                                                "zarubaServiceName" "${CAMEL_SERVICE_NAME}" \
                                                "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
                                              )
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}Creating Fast API Service: ${SERVICE_NAME}{{ $d.Normal }}"
                                              {{ .Zaruba }} generate "${SERVICE_TEMPLATE_LOCATION}" . "${REPLACEMENT_MAP}"
                                              chmod 755 "${CAMEL_SERVICE_NAME}/start.sh"
                                              if [ -f "./main.zaruba.yaml" ]
                                              then
                                                if [ ! -d "./shared-libs/python/helpers" ]
                                                then
                                                  echo "{{ $d.Bold }}{{ $d.Yellow }}Creating shared-lib{{ $d.Normal }}"
                                                  mkdir -p "./shared-libs/python/helpers"
                                                  cp -rnT "./${SERVICE_NAME}/helpers" "./shared-libs/python/helpers"
                                                fi
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Creating shared-lib link for ${SERVICE_NAME}{{ $d.Normal }}"
                                                {{ .Zaruba }} setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "link::${SERVICE_NAME}/helpers" "shared-libs/python/helpers"
                                                link_resource "shared-libs/python/helpers" "${SERVICE_NAME}/helpers"
                                                {{ if .IsTrue (.GetConfig "createTask") -}}
                                                TASK_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Creating service task for ${SERVICE_NAME}{{ $d.Normal }}"
                                                {{ .Zaruba }} makeFastApiServiceTask generatorServiceName=${SERVICE_NAME}
                                                {{ end -}}
                                              fi
                                            fi
                  createTask              : {{ .GetValue "generatorFastApiCreateTask" }}
                  entityName              : {{ .GetValue "generatorFastApiCrudEntity" }}
                  fieldNames              : {{ .GetValue "generatorFastApiCrudFields" }}
                  finish                  : Blank
                  includeBootstrapScript  : if [ -f "${HOME}/.profile" ]
                                            then
                                                . "${HOME}/.profile"
                                            fi
                                            if [ -f "${HOME}/.bashrc" ]
                                            then
                                                . "${HOME}/.bashrc"
                                            fi
                                            BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                            . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript       : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  moduleName              : {{ .GetValue "generatorFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiModule
                  serviceName             : {{ .GetValue "generatorFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            {{ .GetConfig "createModuleScript" }}
                                            TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            CAMEL_SERVICE_NAME=$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")
                                            PASCAL_SERVICE_NAME=$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")
                                            MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                            CAMEL_MODULE_NAME=$({{ .Zaruba }} strToCamel "${MODULE_NAME}")
                                            PASCAL_MODULE_NAME=$({{ .Zaruba }} strToPascal "${MODULE_NAME}")
                                            SNAKE_MODULE_NAME=$({{ .Zaruba }} strToSnake "${MODULE_NAME}")
                                            ENTITY_NAME={{ .EscapeShellArg (.GetConfig "entityName") }}
                                            CAMEL_ENTITY_NAME=$({{ .Zaruba }} strToCamel "${ENTITY_NAME}")
                                            PASCAL_ENTITY_NAME=$({{ .Zaruba }} strToPascal "${ENTITY_NAME}")
                                            SNAKE_ENTITY_NAME=$({{ .Zaruba }} strToSnake "${ENTITY_NAME}")
                                            FIELD_NAMES={{ .EscapeShellArg (.GetConfig "fieldNames") }}
                                            
                                            REPLACEMENT_MAP=$({{ .Zaruba }} mapSet "{}" \
                                              "zarubaServiceName" "${CAMEL_SERVICE_NAME}" \
                                              "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
                                              "zarubaModuleName" "${CAMEL_MODULE_NAME}" \
                                              "ZarubaModuleName" "${PASCAL_MODULE_NAME}" \
                                              "zaruba_module_name" "${SNAKE_MODULE_NAME}" \
                                              "zarubaEntityName" "${CAMEL_ENTITY_NAME}" \
                                              "ZarubaEntityName" "${PASCAL_ENTITY_NAME}" \
                                              "zaruba_entity_name" "${SNAKE_ENTITY_NAME}" \
                                            )
                                            {{ .Zaruba }} generate "${TEMPLATE_LOCATION}/zarubaServiceName" "${CAMEL_SERVICE_NAME}" "${REPLACEMENT_MAP}"
                                            
                                            
                                            MAIN_LINES="$({{ .Zaruba }} readLines "${CAMEL_SERVICE_NAME}/main.py" )"
                                            
                                            # import repo
                                            IMPORT_REPO_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/import_repo.py")"
                                            IMPORT_REPO_PARTIAL="$({{ .Zaruba }} strReplace "${IMPORT_REPO_PARTIAL}" "${REPLACEMENT_MAP}" )"
                                            MAIN_LINES="$({{ .Zaruba }} insertLineBeforeIndex "${MAIN_LINES}" 0 "${IMPORT_REPO_PARTIAL}")"
                                            
                                            # init repo on main.py
                                            INIT_REPO_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/init_repo.py")"
                                            INIT_REPO_PARTIAL="$({{ .Zaruba }} strReplace "${INIT_REPO_PARTIAL}" "${REPLACEMENT_MAP}" )"
                                            ENGINE_DECLARATION_PATTERN="$({{ .Zaruba }} listAppend "[]" "^\s*engine[\s]*=.*$")"
                                            ENGINE_DECLARATION_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${MAIN_LINES}" "${ENGINE_DECLARATION_PATTERN}")"
                                            MAIN_LINES="$({{ .Zaruba }} insertLineAfterIndex "${MAIN_LINES}" "${ENGINE_DECLARATION_LINE_INDEX}" "${INIT_REPO_PARTIAL}")"
                                            
                                            # event controller call
                                            EVENT_CONTROLLER_CALL_PATTERN="$({{ .Zaruba }} listAppend "[]" "^(\s*)${SNAKE_MODULE_NAME}_event_controller\((.*)\)(.*)$")"
                                            EVENT_CONTROLLER_CALL_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${MAIN_LINES}" "${EVENT_CONTROLLER_CALL_PATTERN}")"
                                            EVENT_CONTROLLER_CALL_SUBMATCH="$({{ .Zaruba }} getLineSubmatch "${MAIN_LINES}" "${EVENT_CONTROLLER_CALL_PATTERN}")"
                                            INDENTATION="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_CALL_SUBMATCH}" 1)"
                                            PARAMETERS="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_CALL_SUBMATCH}" 2)"
                                            SUFFIX="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_CALL_SUBMATCH}" 3)"
                                            NEW_EVENT_CONTROLLER_CALL="${INDENTATION}${SNAKE_MODULE_NAME}_event_controller(${PARAMETERS}, ${SNAKE_ENTITY_NAME}_repo)${SUFFIX}"
                                            MAIN_LINES="$({{ .Zaruba }} replaceLineAtIndex "${MAIN_LINES}" "${EVENT_CONTROLLER_CALL_LINE_INDEX}" "${NEW_EVENT_CONTROLLER_CALL}")"
                                            
                                            {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/main.py" "${MAIN_LINES}"
                                            
                                            
                                            CONTROLLER_LINES="$({{ .Zaruba }} readLines "${CAMEL_SERVICE_NAME}/${CAMEL_MODULE_NAME}/controller.py" )"
                                            
                                            # import to controller
                                            IMPORT_TO_CONTROLLER_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/import_to_controller.py")"
                                            IMPORT_TO_CONTROLLER_PARTIAL="$({{ .Zaruba }} strReplace "${IMPORT_TO_CONTROLLER_PARTIAL}" "${REPLACEMENT_MAP}" )"
                                            CONTROLLER_LINES="$({{ .Zaruba }} insertLineBeforeIndex "${CONTROLLER_LINES}" 0 "${IMPORT_TO_CONTROLLER_PARTIAL}")"
                                            
                                            # handle route on controller.py
                                            CONTROLLER_HANDLE_ROUTE_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/controller_handle_route.py")"
                                            CONTROLLER_HANDLE_ROUTE_PARTIAL="$({{ .Zaruba }} strReplace "${CONTROLLER_HANDLE_ROUTE_PARTIAL}" "${REPLACEMENT_MAP}" )"
                                            CONTROLLER_HANDLE_ROUTE_PARTIAL="$({{ .Zaruba }} strIndent "${CONTROLLER_HANDLE_ROUTE_PARTIAL}" "    " )"
                                            ROUTE_CONTROLLER_PATTERN="$({{ .Zaruba }} listAppend "[]" "^\s*def route_controller\(.*\):.*$")"
                                            ROUTE_CONTROLLER_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${CONTROLLER_LINES}" "${ROUTE_CONTROLLER_PATTERN}")"
                                            CONTROLLER_LINES="$({{ .Zaruba }} insertLineAfterIndex "${CONTROLLER_LINES}" "${ROUTE_CONTROLLER_LINE_INDEX}" "${CONTROLLER_HANDLE_ROUTE_PARTIAL}")"
                                            
                                            # handle event on controller.py
                                            CONTROLLER_HANDLE_EVENT_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/controller_handle_event.py")"
                                            CONTROLLER_HANDLE_EVENT_PARTIAL="$({{ .Zaruba }} strReplace "${CONTROLLER_HANDLE_EVENT_PARTIAL}" "${REPLACEMENT_MAP}" )"
                                            CONTROLLER_HANDLE_EVENT_PARTIAL="$({{ .Zaruba }} strIndent "${CONTROLLER_HANDLE_EVENT_PARTIAL}" "    " )"
                                            EVENT_CONTROLLER_PATTERN="$({{ .Zaruba }} listAppend "[]" "^(\s*)def event_controller\((.*)\):(.*)$")"
                                            EVENT_CONTROLLER_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${CONTROLLER_LINES}" "${EVENT_CONTROLLER_PATTERN}")"
                                            CONTROLLER_LINES="$({{ .Zaruba }} insertLineAfterIndex "${CONTROLLER_LINES}" "${EVENT_CONTROLLER_LINE_INDEX}" "${CONTROLLER_HANDLE_EVENT_PARTIAL}")"
                                            
                                            EVENT_CONTROLLER_SUBMATCH="$({{ .Zaruba }} getLineSubmatch "${CONTROLLER_LINES}" "${EVENT_CONTROLLER_PATTERN}")"
                                            INDENTATION="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_SUBMATCH}" 1)"
                                            PARAMETERS="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_SUBMATCH}" 2)"
                                            SUFFIX="$({{ .Zaruba }} listGet "${EVENT_CONTROLLER_SUBMATCH}" 3)"
                                            NEW_EVENT_CONTROLLER="${INDENTATION}def event_controller(${PARAMETERS}, ${SNAKE_ENTITY_NAME}_repo: ${PASCAL_ENTITY_NAME}Repo):${SUFFIX}"
                                            CONTROLLER_LINES="$({{ .Zaruba }} replaceLineAtIndex "${CONTROLLER_LINES}" "${EVENT_CONTROLLER_LINE_INDEX}" "${NEW_EVENT_CONTROLLER}")"
                                            
                                            {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/${CAMEL_MODULE_NAME}/controller.py" "${CONTROLLER_LINES}"
                                            
                                            
                                            # per field
                                            SCHEMA_LINES="$({{ .Zaruba }} readLines "${CAMEL_SERVICE_NAME}/schemas/${CAMEL_ENTITY_NAME}.py")"
                                            REPO_LINES="$({{ .Zaruba }} readLines "${CAMEL_SERVICE_NAME}/repos/db${PASCAL_ENTITY_NAME}.py")"
                                            
                                            FIELD_COUNT="$({{ .Zaruba}} listLength "${FIELD_NAMES}")"
                                            MAX_FIELD_INDEX="$((${FIELD_COUNT}-1))"
                                            for FIELD_INDEX in $(seq "${MAX_FIELD_INDEX}" -1 0)
                                            do
                                              FIELD_NAME="$({{ .Zaruba }} listGet "${FIELD_NAMES}" "${FIELD_INDEX}")"
                                              SNAKE_FIELD_NAME="$({{ .Zaruba }} strToSnake "${FIELD_NAME}")"
                                            
                                              REPLACEMENT_MAP="$({{ .Zaruba }} mapSet "{}" \
                                                "zaruba_entity_name" "${SNAKE_ENTITY_NAME}" \
                                                "zaruba_field_name" "${SNAKE_FIELD_NAME}" \
                                              )"
                                            
                                              # schema field declaration
                                              SCHEMA_FIELD_DECLARATION_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/schema_field_declaration.py")"
                                              SCHEMA_FIELD_DECLARATION_PARTIAL="$({{ .Zaruba }} strReplace "${SCHEMA_FIELD_DECLARATION_PARTIAL}" "${REPLACEMENT_MAP}")"
                                              SCHEMA_FIELD_DECLARATION_PARTIAL="$({{ .Zaruba }} strIndent "${SCHEMA_FIELD_DECLARATION_PARTIAL}" "    ")"
                                            
                                              SCHEMA_FIELD_DECLARATION_PATTERN="$({{ .Zaruba }} listAppend "[]" \
                                                "^\s*class\s*${PASCAL_ENTITY_NAME}Data\s*\(.*$"
                                              )"
                                              SCHEMA_FIELD_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${SCHEMA_LINES}" "${SCHEMA_FIELD_DECLARATION_PATTERN}")"
                                            
                                              SCHEMA_LINES="$({{ .Zaruba }} insertLineAfterIndex "${SCHEMA_LINES}" "${SCHEMA_FIELD_LINE_INDEX}" "${SCHEMA_FIELD_DECLARATION_PARTIAL}")"
                                            
                                              # repo field declaration
                                              REPO_FIELD_DECLARATION_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/repo_field_declaration.py")"
                                              REPO_FIELD_DECLARATION_PARTIAL="$({{ .Zaruba }} strReplace "${REPO_FIELD_DECLARATION_PARTIAL}" "${REPLACEMENT_MAP}")"
                                              REPO_FIELD_DECLARATION_PARTIAL="$({{ .Zaruba }} strIndent "${REPO_FIELD_DECLARATION_PARTIAL}" "    ")"
                                            
                                              REPO_FIELD_DECLARATION_PATTERN="$({{ .Zaruba }} listAppend "[]" \
                                                "^\s*class\s*DB${PASCAL_ENTITY_NAME}Entity\s*\(.*$" \
                                                "^\s*__tablename__.*$" \
                                              )"
                                              REPO_FIELD_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${REPO_LINES}" "${REPO_FIELD_DECLARATION_PATTERN}")"
                                            
                                              REPO_LINES="$({{ .Zaruba }} insertLineAfterIndex "${REPO_LINES}" "${REPO_FIELD_LINE_INDEX}" "${REPO_FIELD_DECLARATION_PARTIAL}")"
                                            
                                              # repo field insert
                                              REPO_FIELD_INSERT_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/repo_field_insert.py")"
                                              REPO_FIELD_INSERT_PARTIAL="$({{ .Zaruba }} strReplace "${REPO_FIELD_INSERT_PARTIAL}" "${REPLACEMENT_MAP}")"
                                              REPO_FIELD_INSERT_PARTIAL="$({{ .Zaruba }} strIndent "${REPO_FIELD_INSERT_PARTIAL}" "$({{ .Zaruba }} strRepeat "    " 4)")"
                                            
                                              REPO_FIELD_INSERT_PATTERN="$({{ .Zaruba }} listAppend "[]" \
                                                "^\s*class\s*DB${PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
                                                "^\s*def\s*insert\s*\(.*$" \
                                                "^\s*db_entity\s*=.*$" \
                                              )"
                                              REPO_FIELD_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${REPO_LINES}" "${REPO_FIELD_INSERT_PATTERN}")"
                                            
                                              REPO_LINES="$({{ .Zaruba }} insertLineAfterIndex "${REPO_LINES}" "${REPO_FIELD_LINE_INDEX}" "${REPO_FIELD_INSERT_PARTIAL}")"
                                            
                                              # repo field update
                                              REPO_FIELD_UPDATE_PARTIAL="$(cat "${TEMPLATE_LOCATION}/partials/repo_field_update.py")"
                                              REPO_FIELD_UPDATE_PARTIAL="$({{ .Zaruba }} strReplace "${REPO_FIELD_UPDATE_PARTIAL}" "${REPLACEMENT_MAP}")"
                                              REPO_FIELD_UPDATE_PARTIAL="$({{ .Zaruba }} strIndent "${REPO_FIELD_UPDATE_PARTIAL}" "$({{ .Zaruba }} strRepeat "    " 3)")"
                                            
                                              REPO_FIELD_UPDATE_PATTERN="$({{ .Zaruba }} listAppend "[]" \
                                                "^\s*class\s*DB${PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
                                                "^\s*def\s*update\s*\(.*$" \
                                                "^\s*db_entity\.updated_at\s*=.*$" \
                                              )"
                                              REPO_FIELD_LINE_INDEX="$({{ .Zaruba }} getLineIndex "${REPO_LINES}" "${REPO_FIELD_UPDATE_PATTERN}")"
                                            
                                              REPO_LINES="$({{ .Zaruba }} insertLineAfterIndex "${REPO_LINES}" "${REPO_FIELD_LINE_INDEX}" "${REPO_FIELD_UPDATE_PARTIAL}")"
                                              
                                            done
                                            
                                            {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/schemas/${CAMEL_ENTITY_NAME}.py" "${SCHEMA_LINES}"
                                            {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/repos/db${PASCAL_ENTITY_NAME}.py" "${REPO_LINES}"
                                            
                                            
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
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/fastapi.zaruba.yaml
                  templateLocation        : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiCrud
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```