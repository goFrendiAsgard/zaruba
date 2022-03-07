<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🏗️ Project](../README.md)
# 🔨 Task
<!--endTocHeader-->


Tasks are the most important component of your zaruba scripts. It defines what you can do with your project resources.

To run a task, you can invoke:

```bash
zaruba please <task-name>
```

Form more information about running a task, please visit [this document](../../../run-task/README.md)

# Task Types

Generally there are two types of tasks:

* [simple command](simple-command.md): Once completed, the task will be ended.
* [long running process](long-running-process.md): Once completed, the task will keep on running (e.g: web server or daemon).

# Task Behavior

A task might also has several [dependencies](define-task-dependies.md). A task with dependencies will never started unless all its dependencies completed. For example, before running a database migration, the database server should already been started.

Futhermore, you can [extend task](extend-task.md) and define custom [environments](task-envs/README.md), [configs](task-configs/README.md), and [inputs](task-inputs.md).

# Task Anatomy

Please see [task anatomy documentation](task-anatomy.md) to learn more about task anatomy.

<!--startTocSubTopic-->
# Sub-topics
* [Task Anatomy](task-anatomy.md)
* [Simple Command](simple-command.md)
* [Long Running Process](long-running-process.md)
* [Task Inputs](task-inputs.md)
* [Task Configs](task-configs/README.md)
  * [Shared Configs](task-configs/shared-configs.md)
* [Task Envs](task-envs/README.md)
  * [Shared Envs](task-envs/shared-envs.md)
* [Extend task](extend-task.md)
* [Define task dependencies](define-task-dependencies.md)
<!--endTocSubTopic-->
