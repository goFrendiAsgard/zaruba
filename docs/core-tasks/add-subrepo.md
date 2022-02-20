<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🥂 addSubrepo
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/chore/subrepo/task.addSubrepo.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Add subrepository.
    TIPS: To init added subrepositories, you should perform `zaruba please initSubrepos`



## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


## Dependencies

* [zrbIsProject](zrb-is-project.md)


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


### Inputs.subrepoName

Description:

    Subrepo name (Can be blank)

Prompt:

    Subrepo name

Secret:

    false


### Inputs.subrepoPrefix

Description:

    Subrepo directory name (Can be blank)

Prompt:

    Subrepo directory name

Secret:

    false


### Inputs.subrepoUrl

Description:

    Subrepo url (Required)

Prompt:

    Subrepo url

Secret:

    false

Validation:

    ^.+$


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



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


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitConfigVariables

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start

Value:

    URL="{{ .GetValue "subrepoUrl" }}"
    if [ -z "${URL}" ]
    then
      echo "${_BOLD}${_RED}subrepoUrl is not defined${_NORMAL}"
      exit 1
    fi
    {{ if .GetValue "subrepoPrefix" }}
      PREFIX="{{ .GetValue "subrepoPrefix" }}"
    {{ else }}
      {{ $urlSegment := .Util.Str.Split (.GetConfig "subrepoUrl") "/" -}}
      {{ $urlSegmentLastIndex := .Subtract (len $urlSegment) 1 -}}
      {{ $urlLastSegment := index $urlSegment $urlSegmentLastIndex -}}
      {{ $prefix := index (.Util.Str.Split $urlLastSegment ".") 0 -}}
      PREFIX="{{ $prefix }}"
    {{ end }}
    NAME="{{ if .GetValue "subrepoName" }}{{ .GetValue "subrepoName" }}{{ else }}${PREFIX}{{ end }}"
    "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::prefix" "${PREFIX}"
    "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::url" "${URL}"
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Subrepo ${NAME} has been added${_NORMAL}"



### Configs.strictMode

Value:

    true


### Configs.subrepoName

Value:

    {{ .GetValue "subrepoName" }}


### Configs.subrepoPrefix

Value:

    {{ .GetValue "subrepoPrefix" }}


### Configs.subrepoUrl

Value:

    {{ .GetValue "subrepoUrl" }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1