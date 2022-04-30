<!--startTocHeader-->
[🏠](../../../../README.md) > [👷🏽 Use Cases](../../../README.md) > [Add Resources](../../README.md) > [✨ From Scratch](../README.md)
# 🏃 Add Runner for Existing Application
<!--endTocHeader-->


There are several builtin tasks to create application runner for your existing applications. One of them is [`addAppRunner`](../../core-tasks/addAppRunner.md)

Please note that `addAppRunner` will make you a generic application runner. This will work in most cases.

However, if your application was written in Go, NodeJs, or Python, you better visit these docs:
* [Add runner for existing Go application](./go-application-runner.md)
* [Add runner for existing Node.Js application](./node-js-application-runner.md)
* [Add runner for existing Python application](./python-application-runner.md)


# How to Add Generic Application Runner

To add a generic application runner, you can perform:

```bash
zaruba please addAppRunner \
  appDirectory=<directory-name> \             # Location of your application. Must be provided
  [appName=<app-name>] \                      # application name
  [appContainerName=<app-container-name>] \   # application's container name
  [appImageName=<app-image-name>] \           # application's image name
  [appDependencies=<app-dependencies>] \      # JSON list containing names of other applications
  [appEnvs=<app-envs>]                        # JSON map containing custom environments
  [appPorts=<app-ports>]                      # JSON list containing application's ports
  [appStartCommand=<start-command>]           # Command to run the app
```


# Generic Application Runner Structure

Once created, you will find a directory named `zaruba-tasks/<app-name>`.

This directory should contains Zaruba scripts to run your application:

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

# Use Case: Create App Runner for Static Web

Suppose you already have a project containing two directories:

* A third party service named `myDb`
* A directory named `myApp` containing static website resources.

Now you want to create an application runner for your static website.

## Install Prerequisites

If don't have `myDb` and `myApp`, then you should run these commands first:

```bash
zaruba please addMysql appDirectory=myDb
zaruba please addSubrepo subrepoUrl="https://github.com/state-alchemists/fibonacci-clock" subrepoPrefix="myApp" 
zaruba please pullSubrepos
```

## Create App Runner

To create app runner, you can run the following command:

```bash
zaruba please addAppRunner \
  appDirectory=./myApp \
  appName=myApp \
  appContainerName=myAppContainer \
  appImageName=my-app \
  appDependencies='["myDb"]' \
  appEnvs='{"HTTP_PORT":"3000"}' \
  appPorts='["3000"]' \
  appStartCommand='./start.sh'
```


## Start Application Runner

Finally, To start your newly generated application runner you can invoke: 

```bash
# run natively
zaruba please startMyApp
```

or

```bash
# run as container
zaruba please startMyAppContainer
```

Zaruba will always start `myDb` first before starting `myApp`. This is expected since you define `myDb` as `myApp`'s application dependency.


<!--startTocSubTopic-->
# Sub-topics
* [Go Application Runner](go-application-runner.md)
* [NodeJs Application Runner](node-js-application-runner.md)
* [Python Application Runner](python-application-runner.md)
<!--endTocSubTopic-->