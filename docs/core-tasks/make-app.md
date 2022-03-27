<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 📜 makeApp
<!--endTocHeader-->

[1m[33m## Information[0m

[1m[34mFile Location[0m:

    ~/.zaruba/zaruba-tasks/make/app/_base/task.makeApp.yaml

[1m[34mShould Sync Env[0m:

    true

[1m[34mType[0m:

    command


[1m[33m## Extends[0m

* [zrbMake](zrb-make.md)


[1m[33m## Dependencies[0m

* [zrbShowAdv](zrb-show-adv.md)


[1m[33m## Start[0m

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


[1m[33m## Inputs[0m


[1m[33m### Inputs.appDirectory[0m

[1m[34mDescription[0m:

    Location of app

[1m[34mPrompt[0m:

    Location of app

[1m[34mSecret[0m:

    false


[1m[33m## Configs[0m


[1m[33m### Configs._adjustPermission[0m

[1m[34mValue[0m:

    if [ -f "${_ZRB_APP_DIRECTORY}/start.sh" ]
    then
      chmod 755 "${_ZRB_APP_DIRECTORY}/start.sh"
    fi



[1m[33m### Configs._finish[0m


[1m[33m### Configs._generate[0m

[1m[34mValue[0m:

    _generate "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_REPLACEMENT_MAP}"


[1m[33m### Configs._includeModuleIndex[0m


[1m[33m### Configs._initShell[0m

[1m[34mValue[0m:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



[1m[33m### Configs._integrate[0m


[1m[33m### Configs._prepare[0m


[1m[33m### Configs._prepareBaseCheckCommand[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareCheckCommand.sh"


[1m[33m### Configs._prepareBaseMigrateCommand[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareMigrateCommand.sh"


[1m[33m### Configs._prepareBasePrepareCommand[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/preparePrepareCommand.sh"


[1m[33m### Configs._prepareBaseReplacementMap[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareBaseReplacementMap.sh"


[1m[33m### Configs._prepareBaseStartCommand[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareStartCommand.sh"


[1m[33m### Configs._prepareBaseTestCommand[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareTestCommand.sh"


[1m[33m### Configs._prepareBaseVariables[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/prepareBaseVariables.sh"


[1m[33m### Configs._prepareReplacementMap[0m


[1m[33m### Configs._prepareVariables[0m


[1m[33m### Configs._registerAppDependencies[0m


[1m[33m### Configs._registerAppDeploymentTasks[0m


[1m[33m### Configs._registerAppRunnerTasks[0m


[1m[33m### Configs._setup[0m

[1m[34mValue[0m:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


[1m[33m### Configs._skipCreation[0m

[1m[34mValue[0m:

    _skipIfExist "{{ .GetConfig "_skipCreationPath" }}"


[1m[33m### Configs._skipCreationPath[0m

[1m[34mValue[0m:

    ${_ZRB_APP_DIRECTORY}


[1m[33m### Configs._start[0m

[1m[34mValue[0m:

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
    echo "${_YELLOW}🧰 Prepare${_NORMAL}"
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
    {{ .GetConfig "_includeModuleIndex" }}
    {{ .GetConfig "_registerAppRunnerTasks" }}
    {{ .GetConfig "_registerAppDeploymentTasks" }}
    {{ .GetConfig "_registerAppDependencies" }}
    {{ .GetConfig "_integrate" }}
    {{ .GetConfig "_adjustPermission" }}
    cd "${__ZRB_PWD}"



[1m[33m### Configs._taskIndexPath[0m


[1m[33m### Configs._validate[0m


[1m[33m### Configs._validateAppContainerVolumes[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppContainerVolumes.sh"


[1m[33m### Configs._validateAppCrudFields[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppCrudFields.sh"


[1m[33m### Configs._validateAppDirectory[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppDirectory.sh"


[1m[33m### Configs._validateAppPorts[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateAppPorts.sh"


[1m[33m### Configs._validateTemplateLocation[0m

[1m[34mValue[0m:

    . "{{ .ZarubaHome }}/zaruba-tasks/make/_base/bash/validateTemplateLocation.sh"


[1m[33m### Configs.afterStart[0m

[1m[34mValue[0m:

    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"


[1m[33m### Configs.appBaseImageName[0m

[1m[34mValue[0m:

    {{ if .GetValue "appBaseImageName" }}{{ .GetValue "appBaseImageName" }}{{ else }}{{ .GetConfig "defaultAppBaseImageName" }}{{ end }}


[1m[33m### Configs.appBuildImageCommand[0m

[1m[34mValue[0m:

    {{ .GetValue "appBuildImageCommand" }}


[1m[33m### Configs.appCheckCommand[0m

[1m[34mValue[0m:

    {{ if ne (.GetValue "appCheckCommand") "" }}{{ .GetValue "appCheckCommand" }}{{ else }}{{ .GetConfig "defaultAppCheckCommand" }}{{ end }}


[1m[33m### Configs.appContainerName[0m

[1m[34mValue[0m:

    {{ .GetValue "appContainerName" }}


[1m[33m### Configs.appContainerVolumes[0m

[1m[34mValue[0m:

    {{ if ne (.GetValue "appContainerVolumes") "[]" }}{{ .GetValue "appContainerVolumes" }}{{ else }}{{ .GetConfig "defaultAppContainerVolumes" }}{{ end }}


[1m[33m### Configs.appCrudEntity[0m

[1m[34mValue[0m:

    {{ .GetValue "appCrudEntity" }}


[1m[33m### Configs.appCrudFields[0m

[1m[34mValue[0m:

    {{ .GetValue "appCrudFields" }}


[1m[33m### Configs.appDependencies[0m

[1m[34mValue[0m:

    {{ .GetValue "appDependencies" }}


[1m[33m### Configs.appDirectory[0m

[1m[34mValue[0m:

    {{ if .GetValue "appDirectory" }}{{ .GetValue "appDirectory" }}{{ else }}{{ .GetConfig "defaultAppDirectory" }}{{ end }}


[1m[33m### Configs.appEnvPrefix[0m

[1m[34mValue[0m:

    {{ .GetValue "appEnvPrefix" }}


[1m[33m### Configs.appEnvs[0m

[1m[34mValue[0m:

    {{ .GetValue "appEnvs" }}


[1m[33m### Configs.appEventName[0m

[1m[34mValue[0m:

    {{ .GetValue "appEventName" }}


[1m[33m### Configs.appHttpMethod[0m

[1m[34mValue[0m:

    {{ .GetValue "appHttpMethod" }}


[1m[33m### Configs.appIcon[0m


[1m[33m### Configs.appImageName[0m

[1m[34mValue[0m:

    {{ .GetValue "appImageName" }}


[1m[33m### Configs.appMigrateCommand[0m

[1m[34mValue[0m:

    {{ .GetValue "appMigrateCommand" }}


[1m[33m### Configs.appModuleName[0m

[1m[34mValue[0m:

    {{ .GetValue "appModuleName" }}


[1m[33m### Configs.appName[0m

[1m[34mValue[0m:

    {{ .GetValue "appName" }}


[1m[33m### Configs.appPorts[0m

[1m[34mValue[0m:

    {{ if ne (.GetValue "appPorts") "[]" }}{{ .GetValue "appPorts" }}{{ else }}{{ .GetConfig "defaultAppPorts" }}{{ end }}


[1m[33m### Configs.appPrepareCommand[0m

[1m[34mValue[0m:

    {{ .GetValue "appPrepareCommand" }}


[1m[33m### Configs.appPushImageCommand[0m

[1m[34mValue[0m:

    {{ .GetValue "appPushImageCommand" }}


[1m[33m### Configs.appRpcName[0m

[1m[34mValue[0m:

    {{ .GetValue "appRpcName" }}


[1m[33m### Configs.appRunnerVersion[0m

[1m[34mValue[0m:

    {{ .GetValue "appRunnerVersion" }}


[1m[33m### Configs.appStartCommand[0m

[1m[34mValue[0m:

    {{ if ne (.GetValue "appStartCommand") "" }}{{ .GetValue "appStartCommand" }}{{ else }}{{ .GetConfig "defaultAppStartCommand" }}{{ end }}


[1m[33m### Configs.appStartContainerCommand[0m

[1m[34mValue[0m:

    {{ if ne (.GetValue "appStartContainerCommand") "" }}{{ .GetValue "appStartContainerCommand" }}{{ else }}{{ .GetConfig "defaultAppStartContainerCommand" }}{{ end }}


[1m[33m### Configs.appTestCommand[0m

[1m[34mValue[0m:

    {{ .GetValue "appTestCommand" }}


[1m[33m### Configs.appUrl[0m

[1m[34mValue[0m:

    {{ .GetValue "appUrl" }}


[1m[33m### Configs.beforeStart[0m


[1m[33m### Configs.cmd[0m

[1m[34mValue[0m:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


[1m[33m### Configs.cmdArg[0m

[1m[34mValue[0m:

    -c


[1m[33m### Configs.defaultAppBaseImageName[0m


[1m[33m### Configs.defaultAppCheckCommand[0m


[1m[33m### Configs.defaultAppContainerVolumes[0m

[1m[34mValue[0m:

    []


[1m[33m### Configs.defaultAppDirectory[0m

[1m[34mValue[0m:

    {{ .GeneratedRandomName }}


[1m[33m### Configs.defaultAppPorts[0m

[1m[34mValue[0m:

    []


[1m[33m### Configs.defaultAppStartCommand[0m


[1m[33m### Configs.defaultAppStartContainerCommand[0m


[1m[33m### Configs.defaultDeploymentDirectory[0m

[1m[34mValue[0m:

    {{ if .GetConfig "defaultAppDirectory" }}{{ .GetConfig "defaultAppDirectory" }}Deployment{{ end }}


[1m[33m### Configs.deploymentDirectory[0m

[1m[34mValue[0m:

    {{ if .GetValue "deploymentDirectory" }}{{ .GetValue "deploymentDirectory" }}{{ else if .GetConfig "appDirectory" }}{{ .GetConfig "appDirectory" }}Deployment{{ else }}{{ .GetConfig "defaultDeploymentDirectory" }}{{ end }}


[1m[33m### Configs.deploymentName[0m

[1m[34mValue[0m:

    {{ .GetValue "deploymentName" }}


[1m[33m### Configs.finish[0m


[1m[33m### Configs.setup[0m


[1m[33m### Configs.shouldInitConfigMapVariable[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.shouldInitConfigVariables[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.shouldInitEnvMapVariable[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.shouldInitUtil[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.start[0m


[1m[33m### Configs.strictMode[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.taskName[0m

[1m[34mValue[0m:

    {{ .GetValue "taskName" }}


[1m[33m### Configs.templateLocations[0m

[1m[34mValue[0m:

    [
      "{{ .ZarubaHome }}/zaruba-tasks/make/app/_base/template"
    ]


[1m[33m## Envs[0m


[1m[33m### Envs.PYTHONUNBUFFERED[0m

[1m[34mFrom[0m:

    PYTHONUNBUFFERED

[1m[34mDefault[0m:

    1