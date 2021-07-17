# makeFastApiRpcHandler
```
  TASK NAME     : makeFastApiRpcHandler
  LOCATION      : /home/gofrendi/zaruba/scripts/task.makeFastApiRpcHandler.zaruba.yaml
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
                  generatorFastApiRpcName
                    DESCRIPTION : RPC name (Required)
                    PROMPT      : RPC name
                    VALIDATION  : ^[a-zA-Z0-9_\-\.]+$
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
                                              create_fast_module "template_location=${MODULE_TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "module_name=${MODULE_NAME}"
                                            fi
                  createServiceScript     : {{- $d := .Decoration -}}
                                            if [ ! -d "./{{ .GetConfig "serviceName" }}" ]
                                            then
                                              SERVICE_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceTemplateLocation") }}
                                              SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}Creating Fast API Service: ${SERVICE_NAME}{{ $d.Normal }}"
                                              create_fast_service "template_location=${SERVICE_TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}"
                                              chmod 755 "${SERVICE_NAME}/start.sh"
                                              if [ -f "./main.zaruba.yaml" ]
                                              then
                                                if [ ! -d "./shared-libs/python/helpers" ]
                                                then
                                                  echo "{{ $d.Bold }}{{ $d.Yellow }}Creating shared-lib{{ $d.Normal }}"
                                                  mkdir -p "./shared-libs/python/helpers"
                                                  cp -rnT "./${SERVICE_NAME}/helpers" "./shared-libs/python/helpers"
                                                fi
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Creating shared-lib link for ${SERVICE_NAME}{{ $d.Normal }}"
                                                "${ZARUBA_HOME}/zaruba" setProjectValue "{{ .GetWorkPath "default.values.yaml" }}" "link::${SERVICE_NAME}/helpers" "shared-libs/python/helpers"
                                                link_resource "shared-libs/python/helpers" "${SERVICE_NAME}/helpers"
                                                {{ if .IsTrue (.GetConfig "createTask") -}}
                                                TASK_TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "taskTemplateLocation") }}
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Creating service task for ${SERVICE_NAME}{{ $d.Normal }}"
                                                create_service_task "template_location=${TASK_TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "image_name=" "container_name=" "location=${SERVICE_NAME}" "start_command=./start.sh" "ports=" "envs=" "dependencies=" "runner_version="
                                                {{ end -}}
                                              fi
                                            fi
                  createTask              : {{ .GetValue "generatorFastApiCreateTask" }}
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
                  moduleName              : {{ .GetValue "generatorFastApiModuleName" }}
                  moduleTemplateLocation  : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiModule
                  playBellScript          : echo $'\a'
                  rpcName                 : {{ .GetValue "generatorFastApiRpcName" }}
                  serviceName             : {{ .GetValue "generatorFastApiServiceName" }}
                  serviceTemplateLocation : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiService
                  setup                   : Blank
                  start                   : {{- $d := .Decoration -}}
                                            {{ .GetConfig "createModuleScript" }}
                                            TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                            SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                            MODULE_NAME={{ .EscapeShellArg (.GetConfig "moduleName") }}
                                            RPC_NAME={{ .EscapeShellArg (.GetConfig "rpcName") }}
                                            should_be_dir "./${SERVICE_NAME}/${MODULE_NAME}" "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME}/${MODULE_NAME} directory should be exist{{ $d.Normal }}"
                                            create_fast_rpc_handler "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "module_name=${MODULE_NAME}" "rpc_name=${RPC_NAME}"
                                            echo 🎉🎉🎉
                                            echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API RPC handler created: ${RPC_NAME} on ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                                            echo "You probably need to check the following files:"
                                            echo "- ${SERVICE_NAME}/main.py"
                                            echo "- ${SERVICE_NAME}/${MODULE_NAME}/controller.py"
                  taskTemplateLocation    : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/fastapi.zaruba.yaml
                  templateLocation        : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiModule
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
