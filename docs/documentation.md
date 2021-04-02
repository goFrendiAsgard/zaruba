# Zaruba script structure

Zaruba script structure is as follow (take note that `UPPER_CASE:` indicate user-defined key):

```yaml
includes: # external scripts you want to load

- ${ZARUBA_HOME}/scripts/core.zaruba.yaml

inputs: # input variables for your tasks

  INPUT_NAME:
    description: DESCRIPTION OF INPUT
    default: DEFAULT VALUE
    secret: false

tasks: # task definition

  TASK_NAME:
    icon: 💚
    private: false
    description: DESCRIPTION OF TASK
    extend: PARENT_TASK_NAME
    location: TASK_WORKING_DIRECTORY
    saveLog: true
    dependencies:
    - PREREQUISITE_TASK_NAME
    inputs:
    - INPUT_NAME
    env:
      ENVIRONMENT_NAME:
        from: GLOBAL_ENVIRONMENT_NAME
        default: DEFAULT_VALUE
    config:
      CONFIG_NAME: CONFIG_VALUE
    lconfig:
      CONFIG_NAME:
      - CONFIG_VALUE_1
      - CONFIG_VALUE_2
    start: 
    - START_COMMAND
    - START_PARAM_1
    - START_PARAM_2
    check:
    - START_COMMAND
    - START_PARAM_1
    - START_PARAM_2
```

# Task Types

## Command

Command is a task that has `start` but doesn't have `check`. For example, the following task is a command:

```yaml
tasks:
  COMMAND_TASK:
    start: [bash, -c, echo "hello world"]
```

Zaruba will consider command task as done when `start` has been executed successfully.

## Service

Service is a task that has both `start` and `check`. For example, the following task is a service:

```yaml
tasks:
  SERVICE_TASK:
    start: [python, -m, http.server]
    check: [bash, -c, sleep 10]
```

Service is usually a long running process (e.g: a web service, a consumer, or anything). Zaruba will consider service task as done when `check` has been executed successfully.

Take note that `start` might still be running even though `check` is complete.

## Wrapper

Wrapper is a task with no `start` or `check`. Wrapper task typically has dependencies:

```yaml
tasks:
  WRAPPER_TASK:
    dependencies:
    - runFrontend
    - runBackend
```

A wrapper task is considered done when all it's dependencies are done.

# Go Template

Zaruba make use of go template to make tasks more configurable.

Please visit [config/taskdata.go](../config/taskdata.go) to see the implementation. Please take note that `taskdata` scope is isolated per task.

You can use go template in every:
* value in task's config.
* value in task's lconfig.
* default in task's env.
* element in task's start.
* element in task's check.

Some properties you can access are
* `{{ .Name }}` Name of current task.
* `{{ .ProjectName }}` Name of the project.
* `{{ .BasePath }}` Directory location of main zaruba script.
* `{{ .WorkPath }}` Working path of current task (i.e: task's location).
* `{{ .DirPath }}` Directory location of current task's definition script.
* `{{ .FileLocation }}` File location of current task's definition script.
* `{{ .Decoration }}` An object containing several output decoration. (e.g: `echo {{ .Decoration.blue }}this is blue text{{ .Decoration.Normal }}`). Please visit [monitor/decoration.go](../monitor/decoration.go) for more detail implementation.

## Available methods

To see available methods, you can either visit [config/taskdata.go](../config/taskdata.go) or see the documentation [here](go-template-available-methods.md)


# Common Base Tasks

To make things easier, zaruba already provide several base tasks.

## core.runShellScript

You can extend this task if you want to run shell script. Usage:

```yaml
tasks:
  MY_TASK:
    extends: core.runShellScript
    config:
      start: echo "hello world"
```

## core.startService

You can extend this task if you want to run long running process. Usage:

```yaml
  tasks:
    MY_SERVICE:
      extend: core.startService
      location: ./my-service
      lconfig:
        ports: [3000]
      config:
        start: go run .
```

By default, zaruba will check every ports define in `lconfig.ports` before considering service as running.

## core.startDockerContainer

You can extend this task if you want to run container. Usage:

```yaml
  tasks:
    runContainerRegistry:
      extend: core.startDockerContainer
      config:
        useImagePrefix: false
        imageName: registry
        imageTag: 2
        containerName: containerRegistry
        localhost: host.docker.internal
        port::9999: 5000
```

By default, zaruba will check every ports define as `config.port`'s subkeys before considering service as running.