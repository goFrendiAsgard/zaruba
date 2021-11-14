
# ZrbRunScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunScript.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Run script.
    Common configs:
      cmd           : Executable shell name
      cmdArg        : Executable shell argument
      setup         : Setup script
      beforeStart   : Before start script
      start         : Start script
      afterStart    : After start script
      finish        : Finish script




## Extends




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


### Configs.afterStart

Value:





### Configs.beforeStart

Value:





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



### Configs.cmdArg

Value:

    -c



### Configs.finish

Value:





### Configs.setup

Value:





### Configs.start

Value:





### Configs._setup

Value:





### Configs._start

Value:





### Configs._finish

Value:





## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1