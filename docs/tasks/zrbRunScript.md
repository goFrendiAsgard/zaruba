
# ZrbRunScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunScript.yaml

Should Sync Env:

    true

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


### Configs.setup

Value:


### Configs._setup

Value:


### Configs._start

Value:


### Configs.beforeStart

Value:


### Configs.cmdArg

Value:

    -c


### Configs.finish

Value:


### Configs._finish

Value:


### Configs.afterStart

Value:


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.start

Value:


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1