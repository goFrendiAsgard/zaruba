<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# addFastAppEventHandler
<!--endTocHeader-->


## Information

File Location:

    ${ZARUBA_HOME}zaruba-tasks/make/fastAppEventHandler/task.addFastAppEventHandler.yaml

Should Sync Env:

    true

Type:

    simple


## Extends

- [makeApp](make-app.md)


## Dependencies

- [addFastApp](add-fast-app.md)
- [addFastAppModule](add-fast-app-module.md)
- [makeFastApp](make-fast-app.md)
- [makeFastAppRunner](make-fast-app-runner.md)
- [zrbIsProject](zrb-is-project.md)
- [zrbShowAdv](zrb-show-adv.md)


## Start

- `{{ .GetConfig "cmd" }}`
- `{{ .GetConfig "cmdArg" }}`
-
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


## Inputs


### Inputs.appDirectory

Description:

    Location of app

Prompt:

    Location of app

Secret:

    false


### Inputs.appEventName

Description:

    Event name (Required)


Prompt:

    Event name

Secret:

    false

Validation:

    ^[a-zA-Z0-9_\-\.]+$


### Inputs.appModuleName

Description:

    Module name (Required)

Prompt:

    Module name

Secret:

    false

Validation:

    ^[a-zA-Z0-9_]+$


## Configs


### Configs._adjustPermission

Value:

    if [ -f "${_ZRB_APP_DIRECTORY}/start.sh" ]
    then
      chmod 755 "${_ZRB_APP_DIRECTORY}/start.sh"
    fi



### Configs._finish


### Configs._generate

Value:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


### Configs._includeModuleIndex


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._integrate

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/fastAppEventHandler/bash/addEventHandler.sh"



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


### Configs._registerAppRunnerTasks


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

    {{ .ProjectName }}FastApp


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

    []


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1



<!--startTocSubtopic-->
<!--endTocSubtopic-->