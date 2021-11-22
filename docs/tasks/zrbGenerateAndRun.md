
# ZrbGenerateAndRun

File Location:

    /zaruba-tasks/_base/generateAndRun/task.zrbGenerateAndRun.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Generate script and run it


## Extends

* `zrbRunShellScript`


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


### Configs.beforeStart


### Configs.cmdArg

Value:

    -c


### Configs.setup


### Configs.start


### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"


### Configs._prepareVariables


### Configs.generatedScriptLocation

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


### Configs.runGeneratedScript

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}/run.sh


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._validate


### Configs._validateTemplateLocation

Value:

    if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
    then
      echo "${_RED}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.${_NORMAL}"
      exit 1
    fi



### Configs.finish


### Configs.includeShellUtil

Value:

    true


### Configs.strictMode

Value:

    true


### Configs.templateLocation

Value:

    {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template


### Configs._start

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
    _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
    _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
    _ZRB_TASK_NAME="{{ .Name }}"
    _ZRB_REPLACEMENT_MAP='{}'
    _ZRB_SCRIPT='{{ .GetConfig "script" }}'
    _ZRB_SQL='{{ .GetConfig "sql" }}'
    _ZRB_IMAGE_NAME="{{ .GetDockerImageName }}"
    _ZRB_IMAGE_TAG="{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}"
    _ZRB_ENVS='{{ .Util.Json.FromStringDict .GetEnvs }}'
    __ZRB_PWD=$(pwd)
    echo "${_YELLOW}🧰 Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}✅ Validate${_NORMAL}"
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🚧 Generate${_NORMAL}"
    echo "${_YELLOW}🚧 Template Location:${_NORMAL} ${_ZRB_TEMPLATE_LOCATION}"
    echo "${_YELLOW}🚧 Generated Script Location:${_NORMAL} ${_ZRB_GENERATED_SCRIPT_LOCATION}"
    echo "${_YELLOW}🚧 Replacement Map:${_NORMAL} ${_ZRB_REPLACEMENT_MAP}"
    mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🏁 Run Script${_NORMAL}"
    echo '{{ .GetConfig "runGeneratedScript" }}'
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"


### Configs._prepareReplacementMap


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.afterStart

Value:

    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"



### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.script

Value:

    {{ .GetValue "script" }}


### Configs.sql

Value:

    {{ .GetValue "sql" }}


### Configs._finish


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1