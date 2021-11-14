
# MakeNodeJsAppRunner

`File Location`:

    /zaruba-tasks/make/nodeJsAppRunner/task.makeNodeJsAppRunner.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:





## Extends

* `makeNativeAppRunner`


## Dependencies

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


## Check




## Inputs


### Inputs.appName

`Default Value`:




`Description`:

    Name of the app


`Prompt`:

    Name of the app


`Secret`:

    false


`Validation`:




`Options`:





### Inputs.appEnvs

`Default Value`:

    {}


`Description`:

    Application envs


`Prompt`:

    Application envs


`Secret`:

    false


`Validation`:




`Options`:





### Inputs.appPorts

`Default Value`:

    []


`Description`:

    Application ports


`Prompt`:




`Secret`:

    false


`Validation`:




`Options`:





### Inputs.appImageName

`Default Value`:




`Description`:

    App's image name


`Prompt`:




`Secret`:

    false


`Validation`:




`Options`:





### Inputs.appContainerName

`Default Value`:




`Description`:

    Application container name


`Prompt`:

    Application container name


`Secret`:

    false


`Validation`:

    ^[a-zA-Z0-9_]*$


`Options`:





### Inputs.appDirectory

`Default Value`:




`Description`:

    Location of app


`Prompt`:

    Location of app


`Secret`:

    false


`Validation`:

    ^[a-zA-Z0-9_]*$


`Options`:





### Inputs.appDependencies

`Default Value`:

    []


`Description`:

    Application dependencies


`Prompt`:

    Application dependencies


`Secret`:

    false


`Validation`:




`Options`:





## Configs

`_finish`:




`_prepareBaseStartCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_taskIndexPath`:

    ./zaruba-tasks/${_ZRB_APP_NAME}/index.yaml


`_validateAppPorts`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_PORTS}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_PORTS: ${_ZRB_APP_PORTS}{{ $d.Normal }}"
      exit 1
    fi



`appPorts`:

    {{ if ne (.GetValue "appPorts") "[]" }}{{ .GetValue "appPorts" }}{{ else }}{{ .GetConfig "defaultAppPorts" }}{{ end }}


`appRpcName`:

    {{ .GetValue "appRpcName" }}


`defaultAppDirectory`:




`setup`:




`start`:




`strictMode`:

    true


`_generate`:

    {{ .GetConfig "_generateBase" }}


`_prepareBaseReplacementMap`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/setReplacementMap.sh"


`_validateAppContainerVolumes`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_CONTAINER_VOLUMES: ${_ZRB_APP_CONTAINER_VOLUMES}{{ $d.Normal }}"
      exit 1
    fi



`appBuildImageCommand`:

    {{ .GetValue "appBuildImageCommand" }}


`appContainerName`:

    {{ .GetValue "appContainerName" }}


`appEventName`:

    {{ .GetValue "appEventName" }}


`appModuleName`:

    {{ .GetValue "appModuleName" }}


`cmdArg`:

    -c


`_integrate`:

    {{ .GetConfig "_registerIndex" }}
    {{ .GetConfig "_registerAppRunnerTasks" }}



`_prepareBaseCheckCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"


`_skipCreationPath`:

    zaruba-tasks/${_ZRB_APP_NAME}


`_validateAppCrudFields`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_CRUD_FIELDS: ${_ZRB_APP_CRUD_FIELDS}{{ $d.Normal }}"
      exit 1
    fi



`appBaseImageName`:

    {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}


`appHttpMethod`:

    {{ .GetValue "appHttpMethod" }}


`appPrepareCommand`:

    {{ .GetValue "appPrepareCommand" }}


`appStartCommand`:

    {{ .GetValue "appStartCommand" }}


`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`deploymentName`:

    {{ .GetValue "deploymentName" }}


`templateLocations`:

    [
      "{{ .ZarubaHome }}/zaruba-tasks/make/appRunner/_base/template",
      "{{ .ZarubaHome }}/zaruba-tasks/make/appRunner/native/template",
      "{{ .ZarubaHome }}/zaruba-tasks/make/nodeJsAppRunner/appRunnerTemplate"
    ]



`_prepareBasePrepareCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"


`_registerAppRunnerTasks`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/task/bash/registerAppRunnerTasks.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_APP_NAME}"


`_skipCreation`:

    {{ $d := .Decoration -}}
    {{ if .GetConfig "_skipCreationPath" -}}
    if [ -x "{{ .GetConfig "_skipCreationPath" }}" ]
    then
      echo "{{ $d.Yellow }}[SKIP] {{ .GetConfig "_skipCreationPath" }} already exist.{{ $d.Normal }}"
      exit 0
    fi
    {{ end -}}



`appDirectory`:

    {{ if .GetValue "appDirectory" }}{{ .GetValue "appDirectory" }}{{ else }}{{ .GetConfig "defaultAppDirectory" }}{{ end }}


`appRunnerVersion`:

    {{ .GetValue "appRunnerVersion" }}


`_prepareBaseTestCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"


`afterStart`:

    {{ $d := .Decoration -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"


`appTestCommand`:

    {{ .GetValue "appTestCommand" }}


`deploymentDirectory`:

    {{ if .GetValue "deploymentDirectory" }}{{ .GetValue "deploymentDirectory" }}{{ else if .GetConfig "appDirectory" }}{{ .GetConfig "appDirectory" }}Deployment{{ else }}{{ .GetConfig "defaultDeploymentDirectory" }}{{ end }}


`_start`:

    {{ $d := .Decoration -}}
    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/util.sh"
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
    _ZRB_REPLACEMENT_MAP='{}'
    __ZRB_PWD=$(pwd)
    echo "{{ $d.Yellow }}🧰 Prepare{{ $d.Normal }}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseStartCommand" }}
    {{ .GetConfig "_prepareBasePrepareCommand" }}
    {{ .GetConfig "_prepareBaseTestCommand" }}
    {{ .GetConfig "_prepareBaseCheckCommand" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "{{ $d.Yellow }}✅ Validate{{ $d.Normal }}"
    {{ .GetConfig "_validateAppDirectory" }}
    {{ .GetConfig "_validateAppContainerVolumes" }}
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validateAppPorts" }}
    {{ .GetConfig "_validateAppCrudFields" }}
    {{ .GetConfig "_skipCreation" }}
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



`_registerIndex`:

    {{ if .GetConfig "_taskIndexPath" -}}
    "{{ .ZarubaBin }}" project include "${_ZRB_PROJECT_FILE_NAME}" "{{ .GetConfig "_taskIndexPath" }}"
    {{ end -}}



`appCheckCommand`:

    {{ .GetValue "appCheckCommand" }}


`appContainerVolumes`:

    {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}


`appCrudFields`:

    {{ .GetValue "appCrudFields" }}


`appDependencies`:

    {{ .GetValue "appDependencies" }}


`appImageName`:

    {{ .GetValue "appImageName" }}


`defaultDeploymentDirectory`:

    {{ if .GetConfig "defaultAppDirectory" }}{{ .GetConfig "defaultAppDirectory" }}Deployment{{ end }}


`includeShellUtil`:

    true


`appStartContainerCommand`:

    {{ .GetValue "appStartContainerCommand" }}


`appName`:

    {{ .GetValue "appName" }}


`defaultAppPorts`:

    []


`_prepareReplacementMap`:




`_validateAppDirectory`:

    {{ $d := .Decoration -}}
    if [ -z "${_ZRB_APP_DIRECTORY}" ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_DIRECTORY: ${_ZRB_APP_DIRECTORY}{{ $d.Normal }}"
      exit 1
    fi



`appPushImageCommand`:

    {{ .GetValue "appPushImageCommand" }}


`beforeStart`:




`defaultAppContainerVolumes`:

    []


`_prepareVariables`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/nodeJsAppRunner/bash/prepareVariables.sh"



`_registerDeploymentTasks`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/task/bash/registerDeploymentTasks.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_DEPLOYMENT_NAME}"


`appIcon`:

    🐸


`_validate`:




`defaultAppBaseImageName`:




`defaultNodeJsAppPorts`:

    [
      "3000"
    ]



`_generateBase`:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


`_prepareBaseVariables`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareVariables.sh"


`_validateTemplateLocation`:

    {{ $d := .Decoration -}}
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



`appEnvPrefix`:

    {{ .GetValue "appEnvPrefix" }}


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`appCrudEntity`:

    {{ .GetValue "appCrudEntity" }}


`appEnvs`:

    {{ .GetValue "appEnvs" }}


`appUrl`:

    {{ .GetValue "appUrl" }}


`finish`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1