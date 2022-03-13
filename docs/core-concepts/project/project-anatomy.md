<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🧬 Project Anatomy
<!--endTocHeader-->

A project usually contains:
* __zaruba scripts__ containing task definitions.
* __application resources__ to run your tasks (e.g: Source code, Dockerfile, etc).

# Zaruba Scripts

A project should at least contains a single zaruba script named `index.zaruba.yaml` in it's top level directory.

Zaruba scripts are written in [YAML](https://yaml.org) and [go template](https://pkg.go.dev/text/template). 

Here is an example of a complete zaruba script:


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
    extend: ''                      # other task name extended by this task. for multiple extend, use `extends` instead (but no, don't use it)
    location: './some-directory'    # directory location where your task should run on
    private: false                  # if true, the task is inteded to be extended instead of run directly
    timeout: 5m
    dependencies: []                # task's upstreams
    inputs:                         # task's inputs
      -inputName
    start: [bash, -c, 'python -m http.server 8080'] # command to start simple-command/long running service
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"'] # command to check readiness of long-running process
    configs:                        # task's configurations
      someConfig: someValue
    envs:                           # task's environments
      SOME_ENV:
        from: SOME_GLOBAL_ENV
        default: defaultValue
    configRef: configName           # shared project configs used by this task
    envRef: envName                 # shared project envs used by this task
    autoTerminate: false            # whether this task should be autoterminated or not
    syncEnv: true                   # whether the environments should be synchronized when running `zaruba please syncEnv` or not
    syncEnvLocation: ''             # location of environment file's directory. If not set, `location` will be used
    saveLog: true                   # wether to save log or not
```

For more information/detail behaviors of each property, please visit their respective documentations:

* [includes](./includes.md)
* [configs](./project-configs.md)
* [envs](./project-envs.md)
* [tasks](../task/README.md)

Zaruba is capable to generate tasks and application resources. You can generate [tasks for existing applications](../../use-cases/add-runner-for-existing-application/README.md) or a [new application along with its task runner](../../use-cases/generate-new-application/README.md).

# Application Resources

Aside from zaruba script, you also need other resources like source code, Dockerfile, static files, etc. We name those resources as `application resources`.

Zaruba can help you generate some of those resources. For example, you can generate:

* [new applications](../../use-cases/generate-new-application/README.md) 
* [deployments](../../use-cases/generate-new-application/add-application-deployment.md), or 
* [third party services](../../use-cases/add-third-party-service/README.md)



<!--startTocSubTopic-->
<!--endTocSubTopic-->