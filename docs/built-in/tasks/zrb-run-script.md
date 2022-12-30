<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# zrbRunScript
<!--endTocHeader-->


## Information

File Location:

    ${ZARUBA_HOME}zaruba-tasks/_base/run/task.zrbRunScript.yaml

Should Sync Env:

    true

Type:

    simple

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

- `{{ .GetConfig "cmd" }}`
- `{{ .GetConfig "cmdArg" }}`
-
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


### Configs._finish


### Configs._setup


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.setup


### Configs.start


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1



# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->