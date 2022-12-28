<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🧬 Project Anatomy
<!--endTocHeader-->

A project usually contains:
* __zaruba scripts__ containing task definitions.
* __application resources__ to run your tasks (e.g., Source code, Dockerfile, etc).

# Zaruba Scripts

Any directory containing `index.zaruba.yaml` is a valid project.

Alongside `index.zaruba.yaml`, a project might contain other Zaruba scripts as well. Those Zaruba scripts are written in [YAML](https://yaml.org) and [go template](https://pkg.go.dev/text/template). 

You can define several `keys` in a single Zaruba script:

* [includes](project-includes.md)
* [inputs](project-inputs.md)
* [configs](project-configs.md)
* [envs](project-envs.md)
* [tasks](../task/README.md)

All `keys` are optional. If you want to see how each key related to each other, you can look at this example:


```yaml
# includes: path of other zaruba scripts
includes:
  - ./path/to/other-script.yaml


# project inputs: definition of task inputs you can share among your tasks
inputs:

  inputName:
    default: defaultValue
    description: input description
    options: [option1, option2]
    prompt: inputPrompt
    allowCustom: true
    secret: false


# project configs: definition of task configurations you can share among your tasks
configs:

  configName:
    sharedConfig: sharedValue


# project envs: definition of task environments you can share among your tasks
envs:

  envName:
    SHARED_ENV:
      from: SHARED_GLOBAL_ENV
      default: defaultSharedValue


# tasks: your task definitions
tasks:

  taskName:
    icon: ✨                        # icon of your task
    description: task description
    extend: ''                      # parent task name
    location: './some-directory'    # task execution location
    private: false                  # whether the task is private or not
    timeout: 5m
    dependencies: []                # task's upstreams
    inputs:                         # task's inputs
      -inputName
    start:                          # command to start simple-command/long running service
      - bash
      - '-c' 
      - python -m http.server 8080
    maxStartRetry: 3                # max start retries (0 indicate infinite retries)
    startRetryDelay: 10s            # delay interval between each start retry
    check:                          # command to check the completeness of a long-running process
      - bash, 
      - '-c' 
      - nc -z localhost 8080
    maxCheckRetry: 0                # max check retries (0 indicate infinite retries)
    checkRetryDelay: 1s             # delay interval between each check retry
    configs:                        # task's configurations
      someConfig: someValue
    envs:                           # task's environments
      SOME_ENV:
        from: SOME_GLOBAL_ENV
        default: defaultValue
    configRef: configName           # shared project configs used by this task
    envRef: envName                 # shared project envs used by this task
    autoTerminate: false            # whether this task should be autoterminated or not
    syncEnv: true                   # whether you can sync task env or not
    syncEnvLocation: ''             # location of environment files. If not set, Zaruba will use `location` instead
    saveLog: true                   # wether to save log or not
```

# Application Resources

To run a task, you might need other resources like source code, Dockerfile, static files, etc.

> 💡 __TIPS:__ Instead of creating application resources and Zaruba script manually, you can use [built-in generator](../../use-cases/add-resources/README.md).


<!--startTocSubtopic--><!--endTocSubtopic-->