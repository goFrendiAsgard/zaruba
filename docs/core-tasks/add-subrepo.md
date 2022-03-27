<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🥂 addSubrepo
<!--endTocHeader-->

[1m[33m## Information[0m

[1m[34mFile Location[0m:

    ~/.zaruba/zaruba-tasks/chore/subrepo/task.addSubrepo.yaml

[1m[34mShould Sync Env[0m:

    true

[1m[34mType[0m:

    command

[1m[34mDescription[0m:

    Add subrepository.
    TIPS: To init added subrepositories, you should perform `zaruba please initSubrepos`



[1m[33m## Extends[0m

* [zrbRunShellScript](zrb-run-shell-script.md)


[1m[33m## Dependencies[0m

* [zrbIsProject](zrb-is-project.md)


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


[1m[33m### Inputs.subrepoName[0m

[1m[34mDescription[0m:

    Subrepo name (Can be blank)

[1m[34mPrompt[0m:

    Subrepo name

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.subrepoPrefix[0m

[1m[34mDescription[0m:

    Subrepo directory name (Can be blank)

[1m[34mPrompt[0m:

    Subrepo directory name

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.subrepoUrl[0m

[1m[34mDescription[0m:

    Subrepo url (Required)

[1m[34mPrompt[0m:

    Subrepo url

[1m[34mSecret[0m:

    false

[1m[34mValidation[0m:

    ^.+$


[1m[33m## Configs[0m


[1m[33m### Configs._finish[0m


[1m[33m### Configs._initShell[0m

[1m[34mValue[0m:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



[1m[33m### Configs._setup[0m

[1m[34mValue[0m:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


[1m[33m### Configs._start[0m


[1m[33m### Configs.afterStart[0m


[1m[33m### Configs.beforeStart[0m


[1m[33m### Configs.cmd[0m

[1m[34mValue[0m:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


[1m[33m### Configs.cmdArg[0m

[1m[34mValue[0m:

    -c


[1m[33m### Configs.finish[0m


[1m[33m### Configs.setup[0m


[1m[33m### Configs.shouldInitConfigMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitEnvMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitUtil[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.start[0m

[1m[34mValue[0m:

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



[1m[33m### Configs.strictMode[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.subrepoName[0m

[1m[34mValue[0m:

    {{ .GetValue "subrepoName" }}


[1m[33m### Configs.subrepoPrefix[0m

[1m[34mValue[0m:

    {{ .GetValue "subrepoPrefix" }}


[1m[33m### Configs.subrepoUrl[0m

[1m[34mValue[0m:

    {{ .GetValue "subrepoUrl" }}


[1m[33m## Envs[0m


[1m[33m### Envs.PYTHONUNBUFFERED[0m

[1m[34mFrom[0m:

    PYTHONUNBUFFERED

[1m[34mDefault[0m:

    1