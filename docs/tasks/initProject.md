
# InitProject

File Location:

    /zaruba-tasks/chore/initProject/task.initProject.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Initiate empty zaruba project.



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


### Configs.afterStart

Value:


### Configs.beforeStart

Value:


### Configs.finish

Value:


### Configs.includeShellUtil

Value:

    true


### Configs.start

Value:

    {{ $d := .Decoration -}}
    if [ -f "index.zaruba.yaml" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}$(pwd) is a zaruba project.{{ $d.Normal }}"
      exit 1
    fi
    git init
    cp -rT "{{ .ZarubaHome }}/zaruba-tasks/chore/initProject/template/" .
    touch .env
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Project created{{ $d.Normal }}"



### Configs._finish

Value:


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.setup

Value:


### Configs.strictMode

Value:

    true


### Configs._start

Value:


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1