<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md)
# 🔨 Task
<!--endTocHeader-->


Tasks are the most important component of your Zaruba project. It defines what you can do with application resources.

To run a task, you can invoke:

```bash
zaruba please <task-name>
```

For more information about running a task, please visit [this guide](../../run-task/README.md).

# Task Types

There are two types of tasks:

* [simple command](simple-command.md): Ended once completed (e.g., `npm install`, or `docker build`).
* [long-running service](long-running-service.md): Keep running once completed. (e.g., `docker start`)

# Task Behavior

## Dependencies

A task might have several [dependencies](define-task-dependies.md). A task with dependencies cannot be started unless all its dependencies are completed.

For example, you cannot run a `web server` before running the `database server`. In that case, you can say that `startWebServer` __depends on__ `startDatabaseServer`.

This is how you define the relation between `startDatabaseServer` and `startWebServer`:

```yaml
tasks:

  startDatabaseServer: {}

  startWebServer:
    dependencies:
      - startDatabaseServer
```

Please note that __you don't have to run the task's dependencies manually__. Zaruba will always run the dependencies for you.

For more information about task dependencies, please visit [this document](define-task-dependies.md).

## Extends

A task might [extend](extend-task.md) other task. Tasks that extend the same parent task are sharing the same behavior.

When you run a `redis/mysql container`, you are basically running a `docker container`. Aside from some specific configurations, the two tasks are having the same behavior.

In that case you can have `startRedisContainer` and `startMysqlContainer` sharing the same parent:

```yaml
tasks:

  startDockerContainer: {}

  startRedisContainer:
    extend: startDockerContainer
    configs:
      imageName: redis:latest
      containerName: redis

  startMysqlContainer:
    extend: startMysqlContainer
    configs:
      imageName: mysql:latest
      containerName: mysql
```

For more information about extending a task, please visit [this document](extend-task.md).


## Shared Environments, Configs, and Inputs

Some tasks might share [environments](task-envs/shared-envs.md), [configs](task-configs/shared-configs.md), and [inputs](task-inputs.md).

For example, when you run `npm install` and `npm start`, you have to share the same node version.

```yaml
envs:

  nvmEnv:
    NODE_VERSION:
      from: NVM_NODE_VERSION
      default: 14

configs:

  nvmConfig:
    init: nvm use ${NODE_VERSION}

tasks:

  npmInstall:
    envRef: nvmEnv
    configRef: nvmConfig
    configs:
      start: |
        ${ZARUBA_CONFIG_INIT} # this refer to `configs.nvmConfig.init`
        npm install

  npmStart:
    envRef: nvmEnv
    configRef: nvmConfig
    configs:
      start: |
        ${ZARUBA_CONFIG_INIT} # this refer to `configs.nvmConfig.init`
        npm start
```

For more information about sharing task resources, please visit the following documents:

* [Shared inputs](task-inputs.md)
* [Shared configs](task-configs/shared-configs.md)
* [Shared envs](task-envs/shared-envs.md)

<!--startTocSubtopic-->
- [🧬 Task Anatomy](taskAnatomy.md)
- [🥛 Simple Command](simpleCommand.md)
- [🍹 Long Running Service](longRunningService.md)
- [⚙️ Task Configs](taskConfigs/README.md)
  - [Shared Configs](taskConfigs/sharedConfigs.md)
- [🏝️ Task Envs](taskEnvs/README.md)
  - [Shared Envs](taskEnvs/sharedEnvs.md)
- [🔤 Task Inputs](taskInputs.md)
- [🧒 Extend task](extendTask.md)
- [🍲 Define task dependencies](defineTaskDependencies.md)
<!--endTocSubtopic-->