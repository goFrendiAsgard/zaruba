<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🪄 zrbGenerateAndRun
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/generateAndRun/task.zrbGenerateAndRun.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Generate script and run it.
    Common configs:
      templateLocation        : Location of the template.
      generatedScriptLocation : Location of the generated script.
      runGeneratedScript      : Script to run generated script.
      _prepareVariables       : Script to initiate additional environment variables.
      _prepareReplacementMap  : Script to modify _ZRB_REPLACEMENT_MAP.
      _validate               : Script to validate configurations.
    Replacements:
      ZTPL_ENV_[.+]           : Environment of current task
      ${[.+]}                 : Environment of current task
      $[.+]                   : Environment of current task
      ztplCfg[.+]             : Configuration of current task



## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


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
    echo "${_YELLOW}🏁 Generated Script${_NORMAL}"
    echo "${ZARUBA_CONFIG_RUN_GENERATED_SCRIPT}"
    echo "${_YELLOW}🏁 Run Generated Script${_NORMAL}"
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



### Configs._validate


### Configs._validateTemplateLocation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/validateTemplateLocation.sh"


### Configs.afterStart

Value:

    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"



### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.generatedScriptLocation

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


### Configs.runGeneratedScript

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}/run.sh


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