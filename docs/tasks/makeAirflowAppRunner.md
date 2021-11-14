
# MakeAirflowAppRunner

`File Location`:

    /zaruba-tasks/make/airflow/task.makeAirflowAppRunner.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:





## Extends

* `makeDockerAppRunner`


## Dependencies

* `makeAirflowApp`
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





### Inputs.airflowWebPorts

`Default Value`:

    ["8080:8080"]


`Description`:

    Airflow's web port


`Prompt`:




`Secret`:

    false


`Validation`:




`Options`:





### Inputs.airflowPostgreSqlPorts

`Default Value`:

    ["5433:5432"]


`Description`:

    Airflow's postgresql port


`Prompt`:




`Secret`:

    false


`Validation`:




`Options`:





### Inputs.airflowRedisPorts

`Default Value`:

    ["6380:6379"]


`Description`:

    Airflow's redis port


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





## Configs

`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_integrate`:

    {{ .GetConfig "_registerIndex" }}
    {{ .GetConfig "_registerAppRunnerTasks" }}



`_validateAppContainerVolumes`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CONTAINER_VOLUMES}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_CONTAINER_VOLUMES: ${_ZRB_APP_CONTAINER_VOLUMES}{{ $d.Normal }}"
      exit 1
    fi



`appContainerName`:

    {{ .GetValue "appContainerName" }}


`appPostgreSqlPorts`:

    {{ .GetValue "airflowPostgreSqlPorts" }}


`_prepareBaseTestCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"


`_registerIndex`:

    {{ if .GetConfig "_taskIndexPath" -}}
    "{{ .ZarubaBin }}" project include "${_ZRB_PROJECT_FILE_NAME}" "{{ .GetConfig "_taskIndexPath" }}"
    {{ end -}}



`_validate`:




`appEnvs`:

    {{ .GetValue "appEnvs" }}


`appBuildImageCommand`:

    {{ .GetValue "appBuildImageCommand" }}


`strictMode`:

    true


`appEventName`:

    {{ .GetValue "appEventName" }}


`appStartCommand`:

    {{ .GetValue "appStartCommand" }}


`_finish`:




`_prepareBaseCheckCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"


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



`_validateAppPorts`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_PORTS}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_PORTS: ${_ZRB_APP_PORTS}{{ $d.Normal }}"
      exit 1
    fi



`appCrudFields`:

    {{ .GetValue "appCrudFields" }}


`appDependencies`:

    {{ .GetValue "appDependencies" }}


`appPrepareCommand`:

    {{ .GetValue "appPrepareCommand" }}


`defaultAppBaseImageName`:




`_skipCreationPath`:

    zaruba-tasks/${_ZRB_APP_NAME}


`appBaseImageName`:

    {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}


`appCheckCommand`:

    {{ .GetValue "appCheckCommand" }}


`appContainerVolumes`:

    {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}


`appEnvPrefix`:

    {{ .GetValue "appEnvPrefix" }}


`appImageName`:

    {{ .GetValue "appImageName" }}


`defaultAppContainerVolumes`:

    [
      "../dags:/opt/bitnami/airflow/dags"
    ]



`_prepareBaseVariables`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareVariables.sh"


`_prepareVariables`:

    _ZRB_APP_POSTGRESQL_PORTS='{{ .GetConfig "appPostgreSqlPorts" }}'
    _ZRB_APP_REDIS_PORTS='{{ .GetConfig "appRedisPorts" }}'
    _ZRB_APP_WEB_PORTS='{{ .GetConfig "appWebPorts" }}'
    . "{{ .ZarubaHome }}/zaruba-tasks/make/airflow/bash/prepareVariables.sh"



`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`_prepareReplacementMap`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/airflow/bash/setReplacementMap.sh"


`_validateAppCrudFields`:

    {{ $d := .Decoration -}}
    if [ "$("{{ .ZarubaBin }}" list validate "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_CRUD_FIELDS: ${_ZRB_APP_CRUD_FIELDS}{{ $d.Normal }}"
      exit 1
    fi



`_validateAppDirectory`:

    {{ $d := .Decoration -}}
    if [ -z "${_ZRB_APP_DIRECTORY}" ]
    then
      echo "{{ $d.Red }}Invalid _ZRB_APP_DIRECTORY: ${_ZRB_APP_DIRECTORY}{{ $d.Normal }}"
      exit 1
    fi



`appTestCommand`:

    {{ .GetValue "appTestCommand" }}


`_prepareBasePrepareCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"


`appDirectory`:

    {{ if .GetValue "appDirectory" }}{{ .GetValue "appDirectory" }}{{ else }}{{ .GetConfig "defaultAppDirectory" }}{{ end }}


`defaultAppPorts`:

    []


`deploymentName`:

    {{ .GetValue "deploymentName" }}


`appModuleName`:

    {{ .GetValue "appModuleName" }}


`appWebPorts`:

    {{ .GetValue "airflowWebPorts" }}


`setup`:




`_taskIndexPath`:

    ./zaruba-tasks/${_ZRB_APP_NAME}/index.yaml


`appName`:

    {{ .GetValue "appName" }}


`appRedisPorts`:

    {{ .GetValue "airflowRedisPorts" }}


`cmdArg`:

    -c


`defaultAppDirectory`:

    {{ .ProjectName }}Airflow


`_generateBase`:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


`appIcon`:

    🎐


`appPushImageCommand`:

    {{ .GetValue "appPushImageCommand" }}


`defaultDeploymentDirectory`:

    {{ if .GetConfig "defaultAppDirectory" }}{{ .GetConfig "defaultAppDirectory" }}Deployment{{ end }}


`finish`:




`start`:




`_prepareBaseReplacementMap`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/setReplacementMap.sh"


`appRunnerVersion`:

    {{ .GetValue "appRunnerVersion" }}


`appUrl`:

    {{ .GetValue "appUrl" }}


`deploymentDirectory`:

    {{ if .GetValue "deploymentDirectory" }}{{ .GetValue "deploymentDirectory" }}{{ else if .GetConfig "appDirectory" }}{{ .GetConfig "appDirectory" }}Deployment{{ else }}{{ .GetConfig "defaultDeploymentDirectory" }}{{ end }}


`_generate`:

    {{ .GetConfig "_generateBase" }}


`afterStart`:

    {{ $d := .Decoration -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"


`appStartContainerCommand`:

    {{ .GetValue "appStartContainerCommand" }}


`includeShellUtil`:

    true


`_registerDeploymentTasks`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/task/bash/registerDeploymentTasks.sh" "${_ZRB_PROJECT_FILE_NAME}" "${_ZRB_DEPLOYMENT_NAME}"


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



`appCrudEntity`:

    {{ .GetValue "appCrudEntity" }}


`templateLocations`:

    [
      "{{ .ZarubaHome }}/zaruba-tasks/make/appRunner/_base/template",
      "{{ .ZarubaHome }}/zaruba-tasks/make/appRunner/docker/template",
      "{{ .ZarubaHome }}/zaruba-tasks/make/airflow/appRunnerTemplate"
    ]



`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`appHttpMethod`:

    {{ .GetValue "appHttpMethod" }}


`_prepareBaseStartCommand`:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"


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



`appPorts`:

    {{ if ne (.GetValue "appPorts") "[]" }}{{ .GetValue "appPorts" }}{{ else }}{{ .GetConfig "defaultAppPorts" }}{{ end }}


`appRpcName`:

    {{ .GetValue "appRpcName" }}



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1