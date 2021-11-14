
# ClearLog

File Location:

    /zaruba-tasks/chore/log/task.clearLog.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Clear log




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


### Configs.finish

Value:





### Configs.includeShellUtil

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





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



### Configs.setup

Value:





### Configs.start

Value:

    {{ $d := .Decoration -}}
    rm -Rf log.zaruba.csv
    echo "{{ $d.Bold }}{{ $d.Yellow }}Log removed{{ $d.Normal }}"



### Configs.strictMode

Value:

    true



### Configs._finish

Value:





### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs._start

Value:





### Configs.cmdArg

Value:

    -c



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1