
# Update

File Location:

    /zaruba-tasks/chore/task.update.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Update zaruba to the latest version.




## Extends

* `zrbRunShellScript`


## Dependencies




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


## Configs


### Configs.cmdArg

Value:

    -c



### Configs.finish

Value:





### Configs.includeShellUtil

Value:

    true



### Configs.setup

Value:





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



### Configs._finish

Value:





### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs._start

Value:





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



### Configs.strictMode

Value:

    true



### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}




### Configs.afterStart

Value:





### Configs.beforeStart

Value:





## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1