<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🪄 zrbMake
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/make/_base/task.zrbMake.yaml

Should Sync Env:

    true

Type:

    simple

Description:

    Make piece of code based on template and replacement map.
    Common configs:
      templateLocations      : JSON array, location of the templates.
      _prepareVariables      : Script to initiate additional environment variables.
      _prepareReplacementMap : Script to modify _ZRB_REPLACEMENT_MAP.
      _validate              : Script to validate configurations.
      _integrate             : Script to integrate the newly generated code to the system.
    Replacements:
      ZTPL_ENV_[.+]                       : Environment of current task
      ztplCfg[.+]                         : Configuration of current task
      [\t ]*ztplAppBuildImageCommand      : Command to build container image
      [\t ]*ztplAppCheckCommand           : Command to check app's readiness
      [\t ]*ztplAppMigrateCommand         : Command to migrate
      [\t ]*ztplAppPrepareCommand         : Command to prepare app
      [\t ]*ztplAppPushImageCommand       : Command to push app's container image
      [\t ]*ztplAppStartCommand           : Command to start app
      [\t ]*ztplAppStartContainerCommand  : Command to start app as container
      [\t ]*ztplAppTestCommand            : Command to test app
      [\t ]*ztplAppYamlContainerVolumes   : Task config value, representing Container volume
      [\t ]*ztplAppYamlEnvs               : Task config value, representing app's environment
      [\t ]*ztplAppYamlPorts              : Task config value, representing app's ports
      ztplAppContainerName                : App's container name
      ztplAppContainerVolumes             : App's container volumes (JSON list)
      ztpl_app_crud_entity                : App's crud entity (snake case)
      ztplAppCrudEntity                   : App's crud entity (camel case)
      ztpl-app-crud-entity                : App's crud entity (kebab case)
      ZtplAppCrudEntity                   : App's curd entity (pascal case)
      ztpl_app_crud_field                 : App's crud field (snake case)
      ztplAppCrudField                    : App's crud field (camel case)
      ztpl-app-crud-field                 : App's crud field (kebab case)
      ZtplAppCrudField                    : App's curd field (pascal case)
      ztplAppCrudFields                   : App's crud fields (JSON list)
      ztpl_app_directory                  : App's directory (snake case)
      ztplAppDirectory                    : App's directory (camel case)
      ztpl-app-directory                  : App's directory (kebab case)
      ZtplAppDirectory                    : App's directory (pascal case)
      ZTPL_APP_ENV_PREFIX                 : App's environment prefix
      ztplAppEnvs                         : App's Environments (JSON map)
      ztpl_app_event_name                 : App's event name (snake case)
      ztplAppEventName                    : App's event name (camel case)  
      ztpl-app-event-name                 : App's event name (kebab case)
      ZtplAppEventName                    : App's event name (pascal case)
      ztplAppHttpMethod                   : App's HTTP method (i,e., get, post, put, delete)
      ztplAppIcon                         : App's icon
      ztpl-app-image-name                 : App's image name
      ztpl_app_module_name                : App's module name (snake case)
      ztplAppModuleName                   : App's module name (camel case)
      ztpl-app-module-name                : App's module name (kebab case)
      ZtplAppModuleName                   : App's module name (pascal case)
      ztpl_app_name                       : App's name (snake case)
      ztplAppName                         : App's name (camel case)
      ztpl-app-name                       : App's name (kebab case)
      ZtplAppName                         : App's name (pascal case)
      ztplAppPorts                        : App's ports (JSON list)
      ztpl_app_rpc_name                   : App's RPC name (snake case)
      ztplAppRpcName                      : App's RPC name (camel case)
      ztpl-app-rpc-name                   : App's RPC name (kebab case)
      ZtplAppRpcName                      : App's RPC name (pascal case)
      ztplAppRunnerVersion                : App's runner version (e.g., node, lts, 14.0, etc)
      ztplAppTaskLocation                 : App's task location
      ztpl_app_url                        : App's url (snake case)
      ztplAppUrl                          : App's url (camel case)
      ztpl-app-url                        : App's url (kebab case)
      ztpl-normalized-app-url             : App's url (kebab case and normalized, i.e., turn to / if empty)
      ZtplAppUrl                          : App's url (pascal case)
      ZtplAppUrlTitle                     : App's url title
      ztpl_deployment_directory           : App's deployment directory (snake case)
      ztplDeploymentDirectory             : App's deployment directory (camel case)
      ztpl-deployment-directory           : App's deployment directory (kebab case)
      ZtplDeploymentDirectory             : App's deployment directory (pascal case)
      ztpl_deployment_name                : App's deployment name (snake case)
      ztplDeploymentName                  : App's deployment name (camel case)
      ztpl-deployment-name                : App's deployment name (kebab case)
      ZtplDeploymentName                  : App's deployment name (pascal case)
      ztplDeploymentTaskLocation          : App's deployment directory relative to task's location
    You can see the detail at ~/.zaruba/zaruba-tasks/make/_base/bash/prepareBaseReplacementMap.sh



## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


## Dependencies

* [zrbShowAdv](zrb-show-adv.md)


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}

    ```


## Configs


### Configs._adjustPermission


### Configs._finish


### Configs._generate

Value:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


### Configs._includeModuleIndex

Value:

    {{ if .GetConfig "_taskIndexPath" -}}
    "{{ .ZarubaBin }}" project include "{{ .GetConfig "_taskIndexPath" }}" "${_ZRB_PROJECT_FILE_NAME}" 
    {{ end -}}



### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._integrate


### Configs._prepare


### Configs._prepareBaseCheckCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"


### Configs._prepareBaseMigrateCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareMigrateCommand.sh"


### Configs._prepareBasePrepareCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"


### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareBaseReplacementMap.sh"


### Configs._prepareBaseStartCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"


### Configs._prepareBaseTestCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"


### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareBaseVariables.sh"


### Configs._prepareReplacementMap


### Configs._prepareVariables


### Configs._registerAppDeploymentTasks

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/registerDeploymentTasks.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_DEPLOYMENT_NAME}"


### Configs._registerAppRunnerTasks

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/registerAppRunnerTasks.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_APP_NAME}"


### Configs._setProjectValue


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._skipCreation

Value:

    _skipIfPathExist "{{ .GetConfig "_skipCreationPath" }}"


### Configs._skipCreationPath


### Configs._start

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/util.sh"
    _ZRB_REPLACEMENT_MAP='{}'
    _ZRB_PROJECT_FILE_NAME='./index.zaruba.yaml'
    _ZRB_APP_BUILD_IMAGE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appBuildImageCommand") "\n " }}'
    _ZRB_APP_CHECK_COMMAND='{{ .Util.Str.Trim (.GetConfig "appCheckCommand") "\n " }}'
    _ZRB_APP_PREPARE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appPrepareCommand") "\n " }}'
    _ZRB_APP_PUSH_IMAGE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appPushImageCommand") "\n " }}'
    _ZRB_APP_START_COMMAND='{{ .Util.Str.Trim (.GetConfig "appStartCommand") "\n " }}'
    _ZRB_APP_START_CONTAINER_COMMAND='{{ .Util.Str.Trim (.GetConfig "appStartContainerCommand") "\n " }}'
    _ZRB_APP_TEST_COMMAND='{{ .Util.Str.Trim (.GetConfig "appTestCommand") "\n " }}'
    _ZRB_APP_MIGRATE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appMigrateCommand") "\n " }}'
    __ZRB_PWD=$(pwd)
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseStartCommand" }}
    {{ .GetConfig "_prepareBasePrepareCommand" }}
    {{ .GetConfig "_prepareBaseTestCommand" }}
    {{ .GetConfig "_prepareBaseMigrateCommand" }}
    {{ .GetConfig "_prepareBaseCheckCommand" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    {{ .GetConfig "_prepare" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Validate${_NORMAL}"
    {{ .GetConfig "_validateAppDirectory" }}
    {{ .GetConfig "_validateAppContainerVolumes" }}
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validateAppPorts" }}
    {{ .GetConfig "_validateAppCrudFields" }}
    {{ .GetConfig "_skipCreation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Generate${_NORMAL}"
    _PRINTED_TEMPLATE_LOCATIONS="$("{{ .ZarubaBin }}" json print "${_ZRB_TEMPLATE_LOCATIONS}" --pretty=false)"
    _STYLED_PRINTED_TEMPLATE_LOCATIONS="${_FAINT}${_PRINTED_TEMPLATE_LOCATIONS}${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Template Location:${_NORMAL} ${_STYLED_PRINTED_TEMPLATE_LOCATIONS}"
    _PRINTED_REPLACEMENT_MAP="$("{{ .ZarubaBin }}" json print "${_ZRB_REPLACEMENT_MAP}" --pretty=false)"
    _STYLED_PRINTED_REPLACEMENT_MAP="${_FAINT}${_PRINTED_REPLACEMENT_MAP}${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Replacement Map:${_NORMAL} ${_STYLED_PRINTED_REPLACEMENT_MAP}"
    {{ .GetConfig "_generate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Integrate${_NORMAL}"
    {{ .GetConfig "_includeModuleIndex" }}
    {{ .GetConfig "_registerAppRunnerTasks" }}
    {{ .GetConfig "_registerAppDeploymentTasks" }}
    {{ .GetConfig "_integrate" }}
    {{ .GetConfig "_adjustPermission" }}
    {{ .GetConfig "_setProjectValue" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}Synchronize task environments${_NORMAL}"
    "{{ .ZarubaBin }}" project syncEnv "./index.zaruba.yaml"
    echo "${_YELLOW}Synchronize project's environment files${_NORMAL}"
    "{{ .ZarubaBin }}" project syncEnvFiles "./index.zaruba.yaml"



### Configs._taskIndexPath


### Configs._validate


### Configs._validateAppContainerVolumes

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppContainerVolumes.sh"


### Configs._validateAppCrudFields

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppCrudFields.sh"


### Configs._validateAppDirectory

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppDirectory.sh"


### Configs._validateAppPorts

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppPorts.sh"


### Configs._validateTemplateLocation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateTemplateLocation.sh"


### Configs.afterStart

Value:

    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"


### Configs.appBaseImageName

Value:

    {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}


### Configs.appBuildImageCommand

Value:

    {{ .GetValue "appBuildImageCommand" }}


### Configs.appCheckCommand

Value:

    {{ if ne (.GetValue "appCheckCommand") "" }}{{ .GetValue "appCheckCommand" }}{{ else }}{{ .GetConfig "defaultAppCheckCommand" }}{{ end }}


### Configs.appContainerName

Value:

    {{ .GetValue "appContainerName" }}


### Configs.appContainerVolumes

Value:

    {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}


### Configs.appCrudEntity

Value:

    {{ .GetValue "appCrudEntity" }}


### Configs.appCrudField

Value:

    {{ .GetValue "appCrudField" }}


### Configs.appCrudFields

Value:

    {{ .GetValue "appCrudFields" }}


### Configs.appDirectory

Value:

    {{ if .GetValue "appDirectory" }}{{ .GetValue "appDirectory" }}{{ else }}{{ .GetConfig "defaultAppDirectory" }}{{ end }}


### Configs.appEnvPrefix

Value:

    {{ .GetValue "appEnvPrefix" }}


### Configs.appEnvs

Value:

    {{ .GetValue "appEnvs" }}


### Configs.appEventName

Value:

    {{ .GetValue "appEventName" }}


### Configs.appHttpMethod

Value:

    {{ .GetValue "appHttpMethod" }}


### Configs.appIcon


### Configs.appImageName

Value:

    {{ .GetValue "appImageName" }}


### Configs.appMigrateCommand

Value:

    {{ if ne (.GetValue "appMigrateCommand") "" }}{{ .GetValue "appMigrateCommand" }}{{ else }}{{ .GetConfig "defaultAppMigrateCommand" }}{{ end }}


### Configs.appModuleName

Value:

    {{ .GetValue "appModuleName" }}


### Configs.appName

Value:

    {{ .GetValue "appName" }}


### Configs.appPorts

Value:

    {{ if ne (.GetValue "appPorts") "[]" }}{{ .GetValue "appPorts" }}{{ else }}{{ .GetConfig "defaultAppPorts" }}{{ end }}


### Configs.appPrepareCommand

Value:

    {{ if ne (.GetValue "appPrepareCommand") "" }}{{ .GetValue "appPrepareCommand" }}{{ else }}{{ .GetConfig "defaultAppPrepareCommand" }}{{ end }}


### Configs.appPushImageCommand

Value:

    {{ .GetValue "appPushImageCommand" }}


### Configs.appRpcName

Value:

    {{ .GetValue "appRpcName" }}


### Configs.appRunnerVersion

Value:

    {{ .GetValue "appRunnerVersion" }}


### Configs.appStartCommand

Value:

    {{ if ne (.GetValue "appStartCommand") "" }}{{ .GetValue "appStartCommand" }}{{ else }}{{ .GetConfig "defaultAppStartCommand" }}{{ end }}


### Configs.appStartContainerCommand

Value:

    {{ if ne (.GetValue "appStartContainerCommand") "" }}{{ .GetValue "appStartContainerCommand" }}{{ else }}{{ .GetConfig "defaultAppStartContainerCommand" }}{{ end }}


### Configs.appTestCommand

Value:

    {{ if ne (.GetValue "appTestCommand") "" }}{{ .GetValue "appTestCommand" }}{{ else }}{{ .GetConfig "defaultAppTestCommand" }}{{ end }}


### Configs.appUrl

Value:

    {{ .GetValue "appUrl" }}


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.defaultAppBaseImageName


### Configs.defaultAppCheckCommand


### Configs.defaultAppContainerVolumes

Value:

    []


### Configs.defaultAppDirectory

Value:

    {{ .GeneratedRandomName }}


### Configs.defaultAppMigrateCommand


### Configs.defaultAppPorts

Value:

    []


### Configs.defaultAppPrepareCommand


### Configs.defaultAppStartCommand


### Configs.defaultAppStartContainerCommand


### Configs.defaultAppTestCommand


### Configs.defaultDeploymentDirectory

Value:

    {{ if .GetConfig "defaultAppDirectory" }}{{ .GetConfig "defaultAppDirectory" }}Deployment{{ end }}


### Configs.deploymentDirectory

Value:

    {{ if .GetValue "deploymentDirectory" }}{{ .GetValue "deploymentDirectory" }}{{ else if .GetConfig "appDirectory" }}{{ .GetConfig "appDirectory" }}Deployment{{ else }}{{ .GetConfig "defaultDeploymentDirectory" }}{{ end }}


### Configs.deploymentName

Value:

    {{ .GetValue "deploymentName" }}


### Configs.finish


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    true


### Configs.shouldInitConfigVariables

Value:

    true


### Configs.shouldInitEnvMapVariable

Value:

    true


### Configs.shouldInitUtil

Value:

    true


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.taskName

Value:

    {{ .GetValue "taskName" }}


### Configs.templateLocations

Value:

    {{ .GetValue "templateLocations" }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1