
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


### Configs.finish


### Configs.start


### Configs._finish


### Configs.afterStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs._setup


### Configs._start


### Configs.beforeStart


### Configs.setup


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1