
# ZrbRunShellScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunShellScript.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Run shell script
    Common configs:
      start : Start script




## Extends

* `zrbRunScript`


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





### Configs.start

Value:

    echo hello world



### Configs.strictMode

Value:

    true



### Configs._finish

Value:





### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}




### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs.cmdArg

Value:

    -c



### Configs.includeShellUtil

Value:

    true



### Configs.setup

Value:





### Configs._start

Value:





### Configs.afterStart

Value:





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