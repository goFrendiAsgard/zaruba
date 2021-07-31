# makeFastApiModule
```
  TASK NAME     : makeFastApiModule
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/makeFastApiModule.zaruba.yaml
  DESCRIPTION   : Make FastAPI module
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
  CONFIG        : _setup                  : set -e
                                            alias zaruba=${ZARUBA_HOME}/zaruba
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
                                              if [! -d "./${SERVICE_NAME}" ]
                                              then
                                                echo "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME} directory should be exist{{ $d.Normal }}"
                                              fi
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}Creating Fast API module: ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                              PASCAL_SERVICE_NAME=$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")
                                              CAMEL_SERVICE_NAME=$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")
                                              PASCAL_MODULE_NAME=$({{ .Zaruba }} strToPascal "${MODULE_NAME}")
                                              CAMEL_MODULE_NAME=$({{ .Zaruba }} strToCamel "${MODULE_NAME}")
                                              SNAKE_MODULE_NAME=$({{ .Zaruba }} strToSnake "${MODULE_NAME}")
                                              REPLACEMENT_MAP=$({{ .Zaruba }} setMapElement "{}" \
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
                                              MAIN_LINES=$({{ .Zaruba }} mergeList "${IMPORT_MODULE_LINES}" "${MAIN_LINES}" "${LOAD_MODULE_LINES}")
                                              {{ .Zaruba }} writeLines "${CAMEL_SERVICE_NAME}/main.py" "${MAIN_LINES}"
                                            fi
                  createServiceScript     : {{- $d := .Decoration -}}
                                            if [ ! -d "./{{ .GetConfig "serviceName" }}" ]
                                            then
                                              SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                              SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                              PASCAL_SERVICE_NAME=$({{ .Zaruba }} strToPascal "${SERVICE_NAME}")
                                              CAMEL_SERVICE_NAME=$({{ .Zaruba }} strToCamel "${SERVICE_NAME}")
                                              REPLACEMENT_MAP=$({{ .Zaruba }} setMapElement "{}" \
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
                  finish                  : Blank
                  includeUtilScript       : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  moduleName              : {{ .GetValue "generatorFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiModule
                  serviceName             : {{ .GetValue "generatorFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            {{ .GetConfig "createModuleScript" }}
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API module created: ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                            echo "You probably need to check the following files:"
                                            echo "- ${SERVICE_NAME}/main.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/controller.py"
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/fastapi.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```