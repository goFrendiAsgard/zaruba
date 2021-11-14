
# ZrbRunPythonScript

File Location:

    /zaruba-tasks/_base/run/task.zrbRunPythonScript.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Run python script
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


### Configs._setup

Value:





### Configs._start

Value:





### Configs.beforeStart

Value:





### Configs.cmd

Value:

    python



### Configs.cmdArg

Value:

    -c



### Configs._finish

Value:





### Configs.afterStart

Value:





### Configs.finish

Value:





### Configs.setup

Value:





### Configs.start

Value:

    print('hello world')



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1