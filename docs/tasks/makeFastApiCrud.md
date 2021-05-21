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
                    DESCRIPTION : Service name
                    PROMPT      : Service name (Required)
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.module.name
                    DESCRIPTION : Module name
                    PROMPT      : Module name (Required)
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.crud.entity
                    DESCRIPTION : Entity name.
                    PROMPT      : Entity (Required)
                    VALIDATION  : ^[a-zA-Z0-9_]+$
                  generator.fastApi.crud.fields
                    DESCRIPTION : Field names.
                    PROMPT      : Field names (comma separated)
                    VALIDATION  : ^[a-zA-Z0-9_,]*$
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  entityName             : {{ .GetValue "generator.fastApi.crud.entity" }}
                  fieldNames             : {{ .GetValue "generator.fastApi.crud.fields" }}
                  finish                 : Blank
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  moduleName             : {{ .GetValue "generator.fastApi.module.name" }}
                  playBellScript         : echo $'\a'
                  serviceName            : {{ .GetValue "generator.fastApi.service.name" }}
                  setup                  : Blank
                  start                  : {{- $d := .Decoration -}}
                                           TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "templateLocation") }}
                                           SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                           MODULE_NAME={{ .SingleQuoteShellValue (.GetConfig "moduleName") }}
                                           ENTITY_NAME={{ .SingleQuoteShellValue (.GetConfig "entityName") }}
                                           FIELD_NAMES={{ .SingleQuoteShellValue (.GetConfig "fieldNames") }}
                                           should_be_dir "./${SERVICE_NAME}/${MODULE_NAME}" "{{ $d.Bold }}{{ $d.Red }}${SERVICE_NAME} directory should be exist{{ $d.Normal }}"
                                           create_fast_crud "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "module_name=${MODULE_NAME}" "entity_name=${ENTITY_NAME}" "field_names=${FIELD_NAMES}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API module created: ${SERVICE_NAME}/${MODULE_NAME}{{ $d.Normal }}"
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiCrud
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
