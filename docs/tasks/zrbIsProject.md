
# ZrbIsProject

File Location:

    /zaruba-tasks/_base/validation/task.zrbIsProject.yaml

Should Sync Env:

    true

Type:

    command


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


### Configs._start

Value:


### Configs.afterStart

Value:


### Configs.cmdArg

Value:

    -c


### Configs.finish

Value:


### Configs.start

Value:

    {{ $d := .Decoration -}}
    if [ ! -f "index.zaruba.yaml" ]
    then
      "{{ $d.Bold }}{{ $d.Red }}$(pwd) is not a zaruba project.{{ $d.Normal }}"
    fi
    echo "{{ $d.Bold }}{{ $d.Yellow }}Current directory is a valid zaruba project{{ $d.Normal }}"



### Configs.includeShellUtil

Value:

    true


### Configs.setup

Value:


### Configs._finish

Value:


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.beforeStart

Value:


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1