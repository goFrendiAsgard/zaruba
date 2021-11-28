
# AddFastApiRouteHandler

File Location:

    /zaruba-tasks/make/fastApiRouteHandler/task.addFastApiRouteHandler.yaml

Should Sync Env:

    true

Type:

    command


## Extends

* `makeApp`


## Dependencies

* `addFastApi`
* `addFastApiModule`
* `makeFastApiApp`
* `makeFastApiAppRunner`
* `zrbIsProject`
* `zrbShowAdv`


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


## Inputs


### Inputs.appDirectory

Description:

    Location of app

Prompt:

    Location of app

Secret:

    false


### Inputs.appHttpMethod

Description:

    HTTP Method (Required)

Prompt:

    HTTP Method

Default Value:

    get

Secret:

    false

Validation:

    ^[a-z]+$

Options:

    get; post; put; delete


### Inputs.appModuleName

Description:

    Module name (Required)

Prompt:

    Module name

Secret:

    false

Validation:

    ^[a-zA-Z0-9_]+$


### Inputs.appUrl

Description:

    URL to be handled (Required)


Prompt:

    URL to be handled

Secret:

    false

Validation:

    ^[a-zA-Z0-9_\-/\{\}]+$


## Configs


### Configs._finish


### Configs._generate

Value:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.SingleQuote (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.SingleQuote (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._integrate

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/fastApiRouteHandler/bash/addRouteHandler.sh"



### Configs._prepareBaseCheckCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"


### Configs._prepareBasePrepareCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"


### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareReplacementMap.sh"


### Configs._prepareBaseStartCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"


### Configs._prepareBaseTestCommand

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"


### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareVariables.sh"


### Configs._prepareReplacementMap


### Configs._prepareVariables


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._skipCreation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/skipCreation.sh"


### Configs._start

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/util.sh"
    _ZRB_REPLACEMENT_MAP='{}'
    _ZRB_PROJECT_FILE_NAME='./index.zaruba.yaml'
    _ZRB_TEMPLATE_LOCATIONS='{{ .GetConfig "templateLocations" }}'
    _ZRB_APP_BASE_IMAGE_NAME='{{ .GetConfig "appBaseImageName" }}'
    _ZRB_APP_BUILD_IMAGE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appBuildImageCommand") "\n " }}'
    _ZRB_APP_CHECK_COMMAND='{{ .Util.Str.Trim (.GetConfig "appCheckCommand") "\n " }}'
    _ZRB_APP_CONTAINER_NAME='{{ .GetConfig "appContainerName" }}'
    _ZRB_APP_CONTAINER_VOLUMES='{{ .GetConfig "appContainerVolumes" }}'
    _ZRB_APP_DEPENDENCIES='{{ .GetConfig "appDependencies" }}'
    _ZRB_APP_DIRECTORY='{{ .GetConfig "appDirectory" }}'
    _ZRB_APP_ENV_PREFIX='{{ .GetConfig "appEnvPrefix" }}'
    _ZRB_APP_ENVS='{{ .GetConfig "appEnvs" }}'
    _ZRB_DEPLOYMENT_DIRECTORY='{{ .GetConfig "deploymentDirectory" }}'
    _ZRB_DEPLOYMENT_NAME='{{ .GetConfig "deploymentName" }}'
    _ZRB_APP_ICON='{{ .GetConfig "appIcon" }}'
    _ZRB_APP_IMAGE_NAME='{{ .GetConfig "appImageName" }}'
    _ZRB_APP_NAME='{{ .GetConfig "appName" }}'
    _ZRB_APP_PORTS='{{ .GetConfig "appPorts" }}'
    _ZRB_APP_PREPARE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appPrepareCommand") "\n " }}'
    _ZRB_APP_PUSH_IMAGE_COMMAND='{{ .Util.Str.Trim (.GetConfig "appPushImageCommand") "\n " }}'
    _ZRB_APP_RUNNER_VERSION='{{ .GetConfig "appRunnerVersion" }}'
    _ZRB_APP_START_COMMAND='{{ .Util.Str.Trim (.GetConfig "appStartCommand") "\n " }}'
    _ZRB_APP_START_CONTAINER_COMMAND='{{ .Util.Str.Trim (.GetConfig "appStartContainerCommand") "\n " }}'
    _ZRB_APP_TEST_COMMAND='{{ .Util.Str.Trim (.GetConfig "appTestCommand") "\n " }}'
    _ZRB_APP_CRUD_ENTITY='{{ .GetConfig "appCrudEntity" }}'
    _ZRB_APP_CRUD_FIELDS='{{ .GetConfig "appCrudFields" }}'
    _ZRB_APP_EVENT_NAME='{{ .GetConfig "appEventName" }}'
    _ZRB_APP_HTTP_METHOD='{{ .GetConfig "appHttpMethod" }}'
    _ZRB_APP_MODULE_NAME='{{ .GetConfig "appModuleName" }}'
    _ZRB_APP_RPC_NAME='{{ .GetConfig "appRpcName" }}'
    _ZRB_APP_URL='{{ .GetConfig "appUrl" }}'
    __ZRB_PWD=$(pwd)
    echo "${_YELLOW}🧰 Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseStartCommand" }}
    {{ .GetConfig "_prepareBasePrepareCommand" }}
    {{ .GetConfig "_prepareBaseTestCommand" }}
    {{ .GetConfig "_prepareBaseCheckCommand" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}✅ Validate${_NORMAL}"
    {{ .GetConfig "_validateAppDirectory" }}
    {{ .GetConfig "_validateAppContainerVolumes" }}
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validateAppPorts" }}
    {{ .GetConfig "_validateAppCrudFields" }}
    {{ .GetConfig "_skipCreation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🚧 Generate${_NORMAL}"
    echo "${_YELLOW}🚧 Template Location:${_NORMAL} ${_ZRB_TEMPLATE_LOCATIONS}"
    echo "${_YELLOW}🚧 Replacement Map:${_NORMAL} ${_ZRB_REPLACEMENT_MAP}"
    {{ .GetConfig "_generate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🔩 Integrate${_NORMAL}"
    {{ .GetConfig "_integrate" }}
    cd "${__ZRB_PWD}"



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

    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"


### Configs.appBaseImageName

Value:

    {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}


### Configs.appBuildImageCommand

Value:

    {{ .GetValue "appBuildImageCommand" }}


### Configs.appCheckCommand

Value:

    {{ .GetValue "appCheckCommand" }}


### Configs.appContainerName

Value:

    {{ .GetValue "appContainerName" }}


### Configs.appContainerVolumes

Value:

    {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}


### Configs.appCrudEntity

Value:

    {{ .GetValue "appCrudEntity" }}


### Configs.appCrudFields

Value:

    {{ .GetValue "appCrudFields" }}


### Configs.appDependencies

Value:

    {{ .GetValue "appDependencies" }}


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

    {{ .GetValue "appPrepareCommand" }}


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

    {{ .GetValue "appStartCommand" }}


### Configs.appStartContainerCommand

Value:

    {{ .GetValue "appStartContainerCommand" }}


### Configs.appTestCommand

Value:

    {{ .GetValue "appTestCommand" }}


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


### Configs.defaultAppContainerVolumes

Value:

    []


### Configs.defaultAppDirectory

Value:

    {{ .ProjectName }}FastApi


### Configs.defaultAppPorts

Value:

    []


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


### Configs.skipCreationPath


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.templateLocations

Value:

    []


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1