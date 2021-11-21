
# AddSubrepo

File Location:

    /zaruba-tasks/chore/subrepo/task.addSubrepo.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Add subrepository.
    TIPS: To init added subrepositories, you should perform `zaruba please initSubrepos`



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


### Inputs.subrepoUrl

Description:

    Subrepo url (Required)

Prompt:

    Subrepo url

Secret:

    false

Validation:

    ^.+$


### Inputs.subrepoPrefix

Description:

    Subrepo directory name (Can be blank)

Prompt:

    Subrepo directory name

Secret:

    false


### Inputs.subrepoName

Description:

    Subrepo name (Can be blank)

Prompt:

    Subrepo name

Secret:

    false


## Configs


### Configs._finish


### Configs.strictMode

Value:

    true


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._start


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.start

Value:

    set -e
    {{ $d := .Decoration -}}
    URL="{{ .GetValue "subrepoUrl" }}"
    if [ -z "${URL}" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}subrepoUrl is not defined{{ $d.Normal }}"
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
    echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepo ${NAME} has been added{{ $d.Normal }}"



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.afterStart


### Configs.includeShellUtil

Value:

    true


### Configs.subrepoPrefix

Value:

    {{ .GetValue "subrepoPrefix" }}


### Configs.subrepoUrl

Value:

    {{ .GetValue "subrepoUrl" }}


### Configs.beforeStart


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.setup


### Configs.subrepoName

Value:

    {{ .GetValue "subrepoName" }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1