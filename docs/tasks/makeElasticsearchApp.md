# makeElasticsearchApp
```
  TASK NAME     : makeElasticsearchApp
  LOCATION      : /zaruba-tasks/make/elasticsearch/task.makeElasticsearchApp.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ makeApp ]
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
                    VALIDATION  : ^[a-zA-Z0-9_]*$
  CONFIG        : _finish                      : Blank
                  _generate                    : {{ .GetConfig "_generateBase" }}
                  _generateBase                : _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"
                  _integrate                   : if [ -f "${_ZRB_APP_DIRECTORY}/start.sh" ]
                                                 then
                                                   chmod 755 "${_ZRB_APP_DIRECTORY}/start.sh"
                                                 fi
                  _prepareBase                 : {{ .GetConfig "_prepareBaseVariables" }}
                                                 {{ .GetConfig "_prepareVariables" }}
                                                 {{ .GetConfig "_prepareBaseStartCommand" }}
                                                 {{ .GetConfig "_prepareBasePrepareCommand" }}
                                                 {{ .GetConfig "_prepareBaseTestCommand" }}
                                                 {{ .GetConfig "_prepareBaseCheckCommand" }}
                                                 {{ .GetConfig "_prepareBaseReplacementMap" }}
                                                 {{ .GetConfig "_prepareReplacementMap" }}
                  _prepareBaseCheckCommand     : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"
                  _prepareBasePrepareCommand   : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"
                  _prepareBaseReplacementMap   : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/setReplacementMap.sh"
                  _prepareBaseStartCommand     : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"
                  _prepareBaseTestCommand      : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"
                  _prepareBaseVariables        : . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareVariables.sh"
                  _prepareReplacementMap       : Blank
                  _prepareVariables            : Blank
                  _setup                       : set -e
                                                 {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                       : {{ $d := .Decoration -}}
                                                 . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/util.sh"
                                                 _ZRB_PROJECT_FILE_NAME='./index.zaruba.yaml'
                                                 _ZRB_TEMPLATE_LOCATIONS='{{ .GetConfig "templateLocations" }}'
                                                 _ZRB_APP_BASE_IMAGE_NAME='{{ .GetConfig "appBaseImageName" }}'
                                                 _ZRB_APP_BUILD_IMAGE_COMMAND='{{ .GetConfig "appBuildImageCommand" }}'
                                                 _ZRB_APP_CHECK_COMMAND='{{ .GetConfig "appCheckCommand" }}'
                                                 _ZRB_APP_CONTAINER_NAME='{{ .GetConfig "appContainerName" }}'
                                                 _ZRB_APP_CONTAINER_VOLUMES='{{ .GetConfig "appContainerVolumes" }}'
                                                 _ZRB_APP_DEPENDENCIES='{{ .GetConfig "appDependencies" }}'
                                                 _ZRB_APP_DIRECTORY='{{ .GetConfig "appDirectory" }}'
                                                 _ZRB_APP_ENV_PREFIX='{{ .GetConfig "appEnvPrefix" }}'
                                                 _ZRB_APP_ENVS='{{ .GetConfig "appEnvs" }}'
                                                 _ZRB_APP_HELM_DIRECTORY='{{ .GetConfig "appHelmDirectory" }}'
                                                 _ZRB_APP_HELM_RELEASE_NAME='{{ .GetConfig "appHelmReleaseName" }}'
                                                 _ZRB_APP_ICON='{{ .GetConfig "appIcon" }}'
                                                 _ZRB_APP_IMAGE_NAME='{{ .GetConfig "appImageName" }}'
                                                 _ZRB_APP_NAME='{{ .GetConfig "appName" }}'
                                                 _ZRB_APP_PORTS='{{ .GetConfig "appPorts" }}'
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
                                                 echo "{{ $d.Yellow }}🧰 Prepare{{ $d.Normal }}"
                                                 {{ .GetConfig "_prepareBase" }}
                                                 cd "${__ZRB_PWD}"
                                                 echo "{{ $d.Yellow }}✅ Validate{{ $d.Normal }}"
                                                 {{ .GetConfig "_validateAppDirectory" }}
                                                 {{ .GetConfig "_validateAppContainerVolumes" }}
                                                 {{ .GetConfig "_validateTemplateLocation" }}
                                                 {{ .GetConfig "_validateAppPorts" }}
                                                 {{ .GetConfig "_validateAppCrudFields" }}
                                                 {{ .GetConfig "_validate" }}
                                                 cd "${__ZRB_PWD}"
                                                 echo "{{ $d.Yellow }}🚧 Generate{{ $d.Normal }}"
                                                 echo "{{ $d.Yellow }}🚧 Template Location:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATIONS}"
                                                 echo "{{ $d.Yellow }}🚧 Replacement Map:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
                                                 {{ .GetConfig "_generate" }}
                                                 cd "${__ZRB_PWD}"
                                                 echo "{{ $d.Yellow }}🔩 Integrate{{ $d.Normal }}"
                                                 {{ .GetConfig "_integrate" }}
                                                 cd "${__ZRB_PWD}"
                  _validate                    : {{ $d := .Decoration -}}
                                                 if [ -d "${_ZRB_APP_DIRECTORY}" ]
                                                 then
                                                   echo "{{ $d.Yellow }}[SKIP] Directory ${_ZRB_APP_DIRECTORY} already exist.{{ $d.Normal }}"
                                                   exit 0
                                                 fi
                  _validateAppContainerVolumes : {{ $d := .Decoration -}}
                                                 if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
                                                 then
                                                   echo "{{ $d.Red }}Invalid _ZRB_APP_CONTAINER_VOLUMES: ${_ZRB_APP_CONTAINER_VOLUMES}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _validateAppCrudFields       : {{ $d := .Decoration -}}
                                                 if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
                                                 then
                                                   echo "{{ $d.Red }}Invalid _ZRB_APP_CRUD_FIELDS: ${_ZRB_APP_CRUD_FIELDS}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _validateAppDirectory        : {{ $d := .Decoration -}}
                                                 if [ -z "${_ZRB_APP_DIRECTORY}" ]
                                                 then
                                                   echo "{{ $d.Red }}Invalid _ZRB_APP_DIRECTORY: ${_ZRB_APP_CONTAINER_DIRECTORY}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _validateAppPorts            : {{ $d := .Decoration -}}
                                                 if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_PORTS}")" = 0 ]
                                                 then
                                                   echo "{{ $d.Red }}Invalid _ZRB_APP_PORTS: ${_ZRB_APP_PORTS}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _validateTemplateLocation    : {{ $d := .Decoration -}}
                                                 if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_TEMPLATE_LOCATIONS}")" = 0 ]
                                                 then
                                                   echo "{{ $d.Red }}Invalid _ZRB_TEMPLATE_LOCATIONS: ${_ZRB_TEMPLATE_LOCATIONS}{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                                                 for _ZRB_TEMPLATE_LOCATION_INDEX in $("{{ .ZarubaBin }}" list rangeIndex "${_ZRB_TEMPLATE_LOCATIONS}")
                                                 do
                                                   _ZRB_TEMPLATE_LOCATION="$("{{ .ZarubaBin }}" list get "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_TEMPLATE_LOCATION_INDEX}")"
                                                   if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
                                                   then
                                                     echo "{{ $d.Red }}{{ $d.Bold }}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.{{ $d.Normal }}"
                                                     exit 1
                                                   fi
                                                 done
                  afterStart                   : {{ $d := .Decoration -}}
                                                 echo 🎉🎉🎉
                                                 echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"
                  appBaseImageName             : {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}
                  appBuildImageCommand         : {{ .GetValue "appBuildImageCommand" }}
                  appCheckCommand              : {{ .GetValue "appCheckCommand" }}
                  appContainerName             : {{ .GetValue "appContainerName" }}
                  appContainerVolumes          : {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}
                  appCrudEntity                : {{ .GetValue "appCrudEntity" }}
                  appCrudFields                : {{ .GetValue "appCrudFields" }}
                  appDependencies              : {{ .GetValue "appDependencies" }}
                  appDirectory                 : {{ if .GetValue "appDirectory" }}{{ .GetValue "appDirectory" }}{{ else }}{{ .GetConfig "defaultAppDirectory" }}{{ end }}
                  appEnvPrefix                 : {{ .GetValue "appEnvPrefix" }}
                  appEnvs                      : {{ .GetValue "appEnvs" }}
                  appEventName                 : {{ .GetValue "appEventName" }}
                  appHelmDirectory             : {{ if .GetValue "appHelmDirectory" }}{{ .GetValue "appHelmDirectory" }}{{ else if .GetConfig "appDirectory" }}{{ .GetConfig "appDirectory" }}Helm{{ else }}{{ .GetConfig "defaultAppHelmDirectory" }}{{ end }}
                  appHelmReleaseName           : {{ .GetValue "appHelmReleaseName" }}
                  appHttpMethod                : {{ .GetValue "appHttpMethod" }}
                  appIcon                      : Blank
                  appImageName                 : {{ .GetValue "appImageName" }}
                  appModuleName                : {{ .GetValue "appModuleName" }}
                  appName                      : {{ .GetValue "appName" }}
                  appPorts                     : {{ if ne (.GetValue "appPorts") "[]" }}{{ .GetValue "appPorts" }}{{ else }}{{ .GetConfig "defaultAppPorts" }}{{ end }}
                  appPrepareCommand            : {{ .GetValue "appPrepareCommand" }}
                  appPushImageCommand          : {{ .GetValue "appPushImageCommand" }}
                  appRpcName                   : {{ .GetValue "appRpcName" }}
                  appRunnerVersion             : {{ .GetValue "appRunnerVersion" }}
                  appStartCommand              : {{ .GetValue "appStartCommand" }}
                  appStartContainerCommand     : {{ .GetValue "appStartContainerCommand" }}
                  appTestCommand               : {{ .GetValue "appTestCommand" }}
                  appUrl                       : {{ .GetValue "appUrl" }}
                  beforeStart                  : Blank
                  cmd                          : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                       : -c
                  defaultAppBaseImageName      : Blank
                  defaultAppContainerVolumes   : []
                  defaultAppDirectory          : {{ .ProjectName }}Cassandra
                  defaultAppHelmDirectory      : {{ if .GetConfig "defaultAppDirectory" }}{{ .GetConfig "defaultAppDirectory" }}Helm{{ end }}
                  defaultAppPorts              : []
                  finish                       : Blank
                  includeShellUtil             : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup                        : Blank
                  start                        : Blank
                  templateLocations            : [
                                                   "{{ .ZarubaHome }}/zaruba-tasks/make/elasticsearch/appTemplate"
                                                 ]
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```