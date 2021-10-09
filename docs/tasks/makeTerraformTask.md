# makeTerraformTask
```
  TASK NAME     : makeTerraformTask
  LOCATION      : /zaruba-tasks/make/_task/terraform/task.makeTerraformTask.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbMakeTask ]
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
  INPUTS        : appDirectory
                    DESCRIPTION : Location of app
                    PROMPT      : Location of app
                    VALIDATION  : ^[a-zA-Z0-9_]+$
  CONFIG        : _finish                        : Blank
                  _integrate                     : {{ .GetConfig "_registerModule" }}
                                                   {{ .GetConfig "_registerTasks" }}
                  _prepare                       : {{ .GetConfig "_prepareVariables" }}
                                                   {{ .GetConfig "_prepareReplacementMap" }}
                  _prepareReplacementMap         : . "${ZARUBA_HOME}/zaruba-tasks/make/_base/bash/setReplacementMap.sh"
                  _prepareVariables              : . "${ZARUBA_HOME}/zaruba-tasks/make/_base/bash/prepareVariables.sh"
                  _registerModule                : . "${ZARUBA_HOME}/zaruba-tasks/make/_task/bash/registerModule.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_MODULE_FILE_NAME}" "${_ZRB_APP_NAME}"
                  _registerTasks                 : . "${ZARUBA_HOME}/zaruba-tasks/make/_task/bash/registerTask.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_MODULE_FILE_NAME}" "${_ZRB_APP_NAME}"
                  _setDefaultAppContainerVolumes : if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
                                                   then
                                                     _ZRB_APP_CONTAINER_VOLUMES='{{ .GetConfig "defaultAppContainerVolumes" }}'
                                                   fi
                  _setDefaultAppPorts            : if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_APP_PORTS}")" = 0 ]
                                                   then
                                                     _ZRB_APP_PORTS='{{ .GetConfig "defaultAppPorts" }}'
                                                   fi
                  _setup                         : set -e
                                                   {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                         : . "${ZARUBA_HOME}/zaruba-tasks/make/_base/bash/util.sh"
                                                   _ZRB_PROJECT_FILE_NAME='./index.zaruba.yaml'
                                                   _ZRB_TEMPLATE_LOCATIONS='{{ .GetConfig "templateLocations" }}'
                                                   {{ .GetConfig "_validateTemplateLocation" }}
                                                   _ZRB_APP_BUILD_IMAGE_COMMAND='{{ .GetConfig "appBuildImageCommand" }}'
                                                   _ZRB_APP_CHECK_COMMAND='{{ .GetConfig "appCheckCommand" }}'
                                                   _ZRB_APP_CONTAINER_NAME='{{ .GetConfig "appContainerName" }}'
                                                   _ZRB_APP_CONTAINER_VOLUMES='{{ .GetConfig "appContainerVolumes" }}'
                                                   {{ .GetConfig "_setDefaultAppContainerVolumes" }}
                                                   {{ .GetConfig "_validateAppContainerVolumes" }}
                                                   _ZRB_APP_DEPENDENCIES='{{ .GetConfig "appDependencies" }}'
                                                   _ZRB_APP_DIRECTORY='{{ .GetConfig "appDirectory" }}'
                                                   _ZRB_APP_ENV_PREFIX='{{ .GetConfig "appEnvPrefix" }}'
                                                   _ZRB_APP_ENVS='{{ .GetConfig "appEnvs" }}'
                                                   _ZRB_APP_HELM_CHART_URL='{{ .GetConfig "appHelmChartUrl" }}'
                                                   _ZRB_APP_IMAGE_NAME='{{ .GetConfig "appImageName" }}'
                                                   _ZRB_APP_NAME='{{ .GetConfig "appName" }}'
                                                   _ZRB_APP_PORTS='{{ .GetConfig "appPorts" }}'
                                                   {{ .GetConfig "_setDefaultAppPorts" }}
                                                   {{ .GetConfig "_validateAppPorts" }}
                                                   _ZRB_APP_PREPARE_COMMAND='{{ .GetConfig "appPrepareCommand" }}'
                                                   _ZRB_APP_PUSH_IMAGE_COMMAND='{{ .GetConfig "appPushImageCommand" }}'
                                                   _ZRB_APP_RUNNER_VERSION='{{ .GetConfig "appRunnerVersion" }}'
                                                   _ZRB_APP_START_COMMAND='{{ .GetConfig "appStartCommand" }}'
                                                   _ZRB_APP_START_CONTAINER_COMMAND='{{ .GetConfig "appStartContainerCommand" }}'
                                                   _ZRB_APP_TEST_COMMAND='{{ .GetConfig "appTestCommand" }}'
                                                   _ZRB_APP_CRUD_ENTITY='{{ .GetConfig "appCrudEntity" }}'
                                                   _ZRB_APP_CRUD_FIELDS='{{ .GetConfig "appCrudFields" }}'
                                                   _ZRB_APP_EVENT_NAME='{{ .GetConfig "appEventName" }}'
                                                   _ZRB_APP_HTTP_METHOD='{{ .GetConfig "appHttpMethod" }}'
                                                   _ZRB_APP_MODULE_NAME='{{ .GetConfig "appModuleName" }}'
                                                   _ZRB_APP_RPC_NAME='{{ .GetConfig "appRpcName" }}'
                                                   _ZRB_APP_URL='{{ .GetConfig "appUrl" }}'
                                                   _ZRB_REPLACEMENT_MAP='{}'
                                                   __ZRB_PWD=$(pwd)
                                                   {{ .GetConfig "_prepare" }}
                                                   cd "${__ZRB_PWD}"
                                                   _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"
                                                   {{ .GetConfig "_integrate" }}
                                                   cd "${__ZRB_PWD}"
                  _validateAppContainerVolumes   : {{ $d := .Decoration -}}
                                                   if [ "$("${ZARUBA_HOME}/zaruba" list validate "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
                                                   then
                                                     echo "{{ $d.Bold }}{{ $d.Red }}Invalid _ZRB_APP_CONTAINER_VOLUMES: ${_ZRB_APP_CONTAINER_VOLUMES}{{ $d.Normal }}"
                                                     exit 1
                                                   fi
                  _validateAppPorts              : {{ $d := .Decoration -}}
                                                   if [ "$("${ZARUBA_HOME}/zaruba" list validate "${_ZRB_APP_PORTS}")" = 0 ]
                                                   then
                                                     echo "{{ $d.Bold }}{{ $d.Red }}Invalid _ZRB_APP_PORTS: ${_ZRB_APP_PORTS}{{ $d.Normal }}"
                                                     exit 1
                                                   fi
                  _validateTemplateLocation      : {{ $d := .Decoration -}}
                                                   if [ "$("${ZARUBA_HOME}/zaruba" list validate "${_ZRB_TEMPLATE_LOCATIONS}")" = 0 ]
                                                   then
                                                     echo "{{ $d.Bold }}{{ $d.Red }}Invalid _ZRB_TEMPLATE_LOCATIONS: ${_ZRB_TEMPLATE_LOCATIONS}{{ $d.Normal }}"
                                                     exit 1
                                                   fi
                  afterStart                     : Blank
                  appBuildImageCommand           : {{ .GetValue "appBuildImageCommand" }}
                  appCheckCommand                : {{ .GetValue "appCheckCommand" }}
                  appContainerName               : {{ .GetValue "appContainerName" }}
                  appContainerVolumes            : {{ .GetValue "appContainerVolumes" }}
                  appCrudEntity                  : {{ .GetValue "appCrudEntity" }}
                  appCrudFields                  : {{ .GetValue "appCrudFields" }}
                  appDependencies                : {{ .GetValue "appDependencies" }}
                  appDirectory                   : {{ .GetValue "appDirectory" }}
                  appEnvPrefix                   : {{ .GetValue "appEnvPrefix" }}
                  appEnvs                        : {{ .GetValue "appEnvs" }}
                  appEventName                   : {{ .GetValue "appEventName" }}
                  appHelmChartUrl                : {{ .GetValue "appHelmChartUrl" }}
                  appHttpMethod                  : {{ .GetValue "appHttpMethod" }}
                  appImageName                   : {{ .GetValue "appImageName" }}
                  appModuleName                  : {{ .GetValue "appModuleName" }}
                  appName                        : {{ .GetValue "appName" }}
                  appPorts                       : {{ .GetValue "appPorts" }}
                  appPrepareCommand              : {{ .GetValue "appPrepareCommand" }}
                  appPushImageCommand            : {{ .GetValue "appPushImageCommand" }}
                  appRpcName                     : {{ .GetValue "appRpcName" }}
                  appRunnerVersion               : {{ .GetValue "appRunnerVersion" }}
                  appStartCommand                : {{ .GetValue "appStartCommand" }}
                  appStartContainerCommand       : {{ .GetValue "appStartContainerCommand" }}
                  appTestCommand                 : {{ .GetValue "appTestCommand" }}
                  appUrl                         : {{ .GetValue "appUrl" }}
                  beforeStart                    : Blank
                  cmd                            : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                         : -c
                  defaultAppContainerVolumes     : []
                  defaultAppPorts                : []
                  finish                         : Blank
                  includeShellUtil               : . ${ZARUBA_HOME}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup                          : Blank
                  start                          : Blank
                  templateLocations              : ["${ZARUBA_HOME}/make/_task/terraform/template"]
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```