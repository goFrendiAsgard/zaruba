
# ZrbRunNodeJsScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunNodeJsScript.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Run Node.Js script
    Common configs:
      start : Start script



## Extends

* `zrbRunScript`


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


### Configs._finish


### Configs._setup


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    node


### Configs.cmdArg

Value:

    -p


### Configs.finish


### Configs.setup


### Configs.start

Value:

    console.log('hello world')


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1