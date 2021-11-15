
# Update

File Location:

    /zaruba-tasks/chore/task.update.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Update zaruba to the latest version.



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


### Configs.strictMode

Value:

    true


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.finish


### Configs.includeShellUtil

Value:

    true


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.setup


### Configs.start

Value:

    {{ $d := .Decoration -}}
    cd "{{ .ZarubaHome }}"
    echo "🔽 {{ $d.Bold }}{{ $d.Yellow }}Pull zaruba{{ $d.Normal }}"
    git checkout master
    git pull origin master
    echo "🚧 {{ $d.Bold }}{{ $d.Yellow }}Compile zaruba{{ $d.Normal }}"
    go build
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Zaruba ready!!!{{ $d.Normal }}"
    echo "{{ $d.Bold }}{{ $d.Yellow }}$(getVersion){{ $d.Normal }}"


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1