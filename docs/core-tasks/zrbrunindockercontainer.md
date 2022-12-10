<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🐳 zrbRunInDockerContainer
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/run/inDockerContainer/task.zrbRunInDockerContainer.yaml

Should Sync Env:

    true

Type:

    simple

Description:

    Run command in a docker container.
    Common configs:
      containerName  : Name of the container.
      containerShell : Shell to run script, default to sh.
      containerUser  : Container's user to run the command.
      remoteCommand  : Command to be executed.
      script         : Script to be executed (Can be multi line).



## Extends

* [zrbGenerateAndRun](zrb-generate-and-run.md)


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


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"


### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"


### Configs._prepareReplacementMap


### Configs._prepareVariables


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
    _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
    _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
    _ZRB_REPLACEMENT_MAP='{}'
    __ZRB_PWD=$(pwd)
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Validate${_NORMAL}"
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Generate${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Template Location:${_NORMAL} ${_FAINT}${_ZRB_TEMPLATE_LOCATION}${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Generated Script Location:${_NORMAL} ${_FAINT}${_ZRB_GENERATED_SCRIPT_LOCATION}${_NORMAL}"
    _PRINTED_REPLACEMENT_MAP="$("{{ .ZarubaBin }}" json print "${_ZRB_REPLACEMENT_MAP}" --pretty=false)"
    _STYLED_PRINTED_REPLACEMENT_MAP="${_FAINT}${_PRINTED_REPLACEMENT_MAP}${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Replacement Map:${_NORMAL} ${_STYLED_PRINTED_REPLACEMENT_MAP}"
    mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_START_ICON} Generated Script${_NORMAL}"
    echo "${ZARUBA_CONFIG_RUN_GENERATED_SCRIPT}"
    echo "${_YELLOW}${_START_ICON} Run Generated Script${_NORMAL}"
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



### Configs._validate


### Configs._validateTemplateLocation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/validateTemplateLocation.sh"


### Configs.afterStart

Value:

    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"



### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.containerName


### Configs.containerShell

Value:

    sh


### Configs.containerUser


### Configs.finish


### Configs.generatedScriptLocation

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


### Configs.remoteCommand

Value:

    {{ .GetConfig "containerShell" }} "{{ .GetConfig "remoteScriptLocation" }}/run.sh"


### Configs.remoteScriptLocation

Value:

    _{{ .Name }}.script.{{ .UUID }}


### Configs.runGeneratedScript

Value:

    _ZRB_CONTAINER_NAME="{{ .GetConfig "containerName" }}"
    _ZRB_REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Make ${_ZRB_GENERATED_SCRIPT_LOCATION} executable${_NORMAL}"
    chmod -R 755 "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Copy from ${_ZRB_GENERATED_SCRIPT_LOCATION} at host to ${_ZRB_REMOTE_SCRIPT_LOCATION} at container ${_ZRB_CONTAINER_NAME}${_NORMAL}"
    docker cp "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_CONTAINER_NAME}:${_ZRB_REMOTE_SCRIPT_LOCATION}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Execute remote command${_NORMAL}"
    docker exec {{ if .GetConfig "containerUser" }}-u {{ .GetConfig "containerUser" }}{{ end }} "${_ZRB_CONTAINER_NAME}" {{ .GetConfig "remoteCommand" }}
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Remove ${_ZRB_REMOTE_SCRIPT_LOCATION} at container ${_ZRB_CONTAINER_NAME}${_NORMAL}"
    docker exec -u 0 "${_ZRB_CONTAINER_NAME}" rm -Rf "${_ZRB_REMOTE_SCRIPT_LOCATION}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Remove ${_ZRB_GENERATED_SCRIPT_LOCATION}${_NORMAL}"
    rm -Rf "${_ZRB_GENERATED_SCRIPT_LOCATION}"


### Configs.script

Value:

    {{ .GetValue "script" }}


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


### Configs.sql

Value:

    {{ .GetValue "sql" }}


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.templateLocation

Value:

    {{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/template


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1