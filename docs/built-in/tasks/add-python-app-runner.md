<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# addPythonAppRunner
<!--endTocHeader-->


## Information

File Location:

    /home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/task.addPythonAppRunner.yaml

Should Sync Env:

    false

Type:

    wrapper


## Dependencies

- `makePythonAppRunner`
- `zrbIsProject`
- `zrbShowAdv`


## Inputs


### Inputs.appCheckCommand

Description:

    Command to check app

Secret:

    false


### Inputs.appContainerName

Description:

    Application container name

Prompt:

    Application container name

Secret:

    false

Validation:

    ^[a-zA-Z0-9_]*$


### Inputs.appDirectory

Description:

    Location of app

Prompt:

    Location of app

Secret:

    false


### Inputs.appEnvs

Description:

    Application envs

Prompt:

    Application envs

Default Value:

    {}

Secret:

    false


### Inputs.appImageName

Description:

    App's image name

Secret:

    false


### Inputs.appMigrateCommand

Description:

    Command to do migration

Secret:

    false


### Inputs.appName

Description:

    Name of the app

Prompt:

    Name of the app

Secret:

    false


### Inputs.appPorts

Description:

    Application ports

Default Value:

    []

Secret:

    false


### Inputs.appStartCommand

Description:

    Command to start app

Secret:

    false


### Inputs.appTestCommand

Description:

    Command to test app

Secret:

    false



<!--startTocSubtopic-->

<!--endTocSubtopic-->