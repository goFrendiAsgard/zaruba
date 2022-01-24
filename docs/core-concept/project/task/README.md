[⬅️ Table of Content](../../../README.md)

# Task

Tasks are the most important component of your zaruba scripts. It defines what you can do with your project resources.

To run a task, you can invoke:

```bash
zaruba please <task-name>
```

You can also:

* [run many task in parallel](../../README.md#run-many-tasks-in-parallel)
* [run task in interactive mode](../../README.md#run-tasks-in-interactive-mode) or
* ask Zaruba to [explain a tasks](../../README.md#explain-a-task)

# Task Types

Generally there are two types of tasks:

* [simple command](simple-command.md): Once completed, the task will be ended.
* [long running process](long-running-process.md): Once completed, the task will keep on running (e.g: web server or daemon).

# Task Behavior

A task might also has several [dependencies](define-task-dependies.md). A task with dependencies will never started unless all it's dependencies completed. For example, before running a database migration, the database server should already been started.

Futhermore, you can [extend task](extend-task.md) and define custom [environments](task-envs/README.md), [configs](task-configs/README.md), and [inputs](task-inputs.md).

# Task Anatomy

Please see [task anatomy documentation](task-anatomy.md) to learn more about task anatomy.