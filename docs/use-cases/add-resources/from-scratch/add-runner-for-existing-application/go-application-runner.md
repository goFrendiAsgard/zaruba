<!--startTocHeader-->
[🏠](../../../../README.md) > [👷🏽 Use Cases](../../../README.md) > [📦 Add Resources](../../README.md) > [✨ From Scratch](../README.md) > [🏃 Add Runner for Existing Application](README.md)
# Go Application Runner
<!--endTocHeader-->

To create application runner for your existing go applications, you can use [`addGoAppRunner`](../../core-tasks/addGoAppRunner.md)

# How to Add Go Application Runner


```bash
zaruba please addGoAppRunner \
  appDirectory=<directory-name> \             # Location of your application. Must be provided
  [appName=<app-name>] \                      # application name
  [appContainerName=<app-container-name>] \   # application's container name
  [appImageName=<app-image-name>] \           # application's image name
  [appEnvs=<app-envs>]                        # JSON map containing custom environments
  [appPorts=<app-ports>]                      # JSON list containing application's ports
  [appStartCommand=<start-command>]           # Command to run the app
```

# Go Application Runner Structure

```bash
❯ tree
.
├── configs.yaml
├── envs.yaml
├── index.yaml
├── inputs.yaml
├── tasks.container.yaml
└── tasks.yaml

0 directories, 6 files

```

# Use Case: Create Go Application Runner

Suppose you already have a project containing two directories:

* A third party service named `myDb`
* A directory named `myGoApp` containing static website resources.

Now you want to create an application runner for your static website.

## Install Prerequisites

If don't have `myDb` and `myGoApp`, then you should run these commands first:

```bash
zaruba please addMysql appDirectory=myDb
zaruba please makeSimpleGoApp appDirectory=myGoApp
```

## Create Go App Runner

To create app runner, you can run the following command:

```bash
zaruba please addGoAppRunner \
  appDirectory=./myGoApp \
  appName=myGoApp \
  appContainerName=myAppContainer \
  appImageName=my-app \
  appEnvs='{"HTTP_PORT":"3000"}' \
  appPorts='["3000"]' \
  appStartCommand='./start.sh'

zaruba task addDependencies runGoAppRunner runMyDb
zaruba task addDependencies runGoAppContainer runMyDbContainer
```


## Start Application Runner

Finally, To start your newly generated application runner you can invoke: 

```bash
# run natively
zaruba please startMyGoApp
```

or

```bash
# run as container
zaruba please startMyGoAppContainer
```

Zaruba will always start `myDb` first before starting `myGoApp`. This is expected since you define `myDb` as `myGoApp`'s application dependency.


<!--startTocSubTopic-->
<!--endTocSubTopic-->