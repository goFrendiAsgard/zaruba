
# SetProjectValue

File Location:

    /zaruba-tasks/chore/value/task.setProjectValue.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Set project value.


## Extends

* `zrbRunShellScript`


## Dependencies

* `zrbIsProject`


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


### Inputs.variableName

Description:

    Variable name (Required)

Prompt:

    Name

Secret:

    false

Validation:

    ^.+$


### Inputs.variableValue

Description:

    Variable value (Required)

Prompt:

    Value

Secret:

    false

Validation:

    ^.+$


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "initShellConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "initShellConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.SingleQuote (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "initShellEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.SingleQuote (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.includeShellUtil

Value:

    true


### Configs.initShellConfigMapVariable

Value:

    false


### Configs.initShellConfigVariables

Value:

    false


### Configs.initShellEnvMapVariable

Value:

    false


### Configs.setup


### Configs.start

Value:

    "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "{{ .GetConfig "variableName" }}" "{{ .GetConfig "variableValue" }}"
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Kwarg ${KEY} : ${VALUE} has been set${_NORMAL}"



### Configs.strictMode

Value:

    true


### Configs.variableName

Value:

    {{ .GetValue "variableName" }}


### Configs.variableValue

Value:

    {{ .GetValue "variableValue" }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1