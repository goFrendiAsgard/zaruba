
# ZrbRunPythonScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunPythonScript.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Run python script
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


### Configs.cmdArg

Value:

    -c


### Configs.start

Value:

    print('hello world')


### Configs._finish


### Configs.afterStart


### Configs.cmd

Value:

    python


### Configs.finish


### Configs.setup


### Configs._setup


### Configs._start


### Configs.beforeStart


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1