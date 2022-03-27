<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🐍 addPythonAppRunner
<!--endTocHeader-->

[1m[33m## Information[0m

[1m[34mFile Location[0m:

    ~/.zaruba/zaruba-tasks/make/pythonAppRunner/task.addPythonAppRunner.yaml

[1m[34mShould Sync Env[0m:

    false

[1m[34mType[0m:

    wrapper


[1m[33m## Dependencies[0m

* [makePythonAppRunner](make-python-app-runner.md)
* [zrbIsProject](zrb-is-project.md)
* [zrbShowAdv](zrb-show-adv.md)


[1m[33m## Inputs[0m


[1m[33m### Inputs.appCheckCommand[0m

[1m[34mDescription[0m:

    Command to check app

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appContainerName[0m

[1m[34mDescription[0m:

    Application container name

[1m[34mPrompt[0m:

    Application container name

[1m[34mSecret[0m:

    false

[1m[34mValidation[0m:

    ^[a-zA-Z0-9_]*$


[1m[33m### Inputs.appDependencies[0m

[1m[34mDescription[0m:

    Application dependencies

[1m[34mPrompt[0m:

    Application dependencies

[1m[34mDefault Value[0m:

    []

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appDirectory[0m

[1m[34mDescription[0m:

    Location of app

[1m[34mPrompt[0m:

    Location of app

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appEnvs[0m

[1m[34mDescription[0m:

    Application envs

[1m[34mPrompt[0m:

    Application envs

[1m[34mDefault Value[0m:

    {}

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appImageName[0m

[1m[34mDescription[0m:

    App's image name

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appMigrateCommand[0m

[1m[34mDescription[0m:

    Command to do migration

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appName[0m

[1m[34mDescription[0m:

    Name of the app

[1m[34mPrompt[0m:

    Name of the app

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appPorts[0m

[1m[34mDescription[0m:

    Application ports

[1m[34mDefault Value[0m:

    []

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appStartCommand[0m

[1m[34mDescription[0m:

    Command to start app

[1m[34mSecret[0m:

    false


[1m[33m### Inputs.appTestCommand[0m

[1m[34mDescription[0m:

    Command to test app

[1m[34mSecret[0m:

    false