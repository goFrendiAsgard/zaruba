<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md)
# 🔨 Task
<!--endTocHeader-->


Tasks are the most important component of your zaruba project. It defines what you can do with application resources.

To run a task, you can invoke:

```bash
zaruba please <task-name>
```

Form more information about running a task, please visit [run task section](../../run-task/README.md)

# Task Types

Generally there are two types of tasks:

* [simple command](simple-command.md): Ended once completed (e.g: `npm install`, or `docker build`).
* [long running service](long-running-service.md): Keep running once completed. (e.g: `docker start`)

# Task Behavior

## Dependencies

A task might has several [dependencies](define-task-dependies.md). A task with dependencies will never started unless all its dependencies completed. For example, you cannot run a __web server__ before a __database server__ is started. In that case, you can say that __startWebServer__ depends on __startDatabaseServer__.

## Extends

A task might [extend](extend-task.md) other task. For example, running a redis container can be defined as running a container, with redis image. Thus, you can say that __startRedisContainer__ extend __startDockerContainer__.

## Share Environments, Configs, and Inputs

A task might share [environments](task-envs/shared-envs), [configs](task-configs/shared-configs.md), and [inputs](task-inputs.md).

<!--startTocSubTopic-->
# Sub-topics
* [🧬 Task Anatomy](task-anatomy.md)
* [🥛 Simple Command](simple-command.md)
* [🍹 Long Running Service](long-running-service.md)
* [🔤 Task Inputs](task-inputs.md)
* [⚙️ Task Configs](task-configs/README.md)
  * [Shared Configs](task-configs/shared-configs.md)
* [🏝️ Task Envs](task-envs/README.md)
  * [Shared Envs](task-envs/shared-envs.md)
* [🧒 Extend task](extend-task.md)
* [🍲 Define task dependencies](define-task-dependencies.md)
<!--endTocSubTopic-->
