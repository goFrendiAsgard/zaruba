[⬅️ Table of Content](../../../README.md)

# Task

Tasks are the most important component of your zaruba scripts. It defines what you can do with your project resources.

To run a task, you can invoke:

```bash
zaruba please <task-name>
```

Generally there are two types of tasks:

* [simple command](simple-command.md)
* [long running process](long-running-process.md)

A task might also [depends](define-task-dependies.md) on other tasks. So, before you run a particular task, zaruba will run it's dependencies first.

Futhermore, you can [extend task](extend-task.md), define task's [environment](task-envs/README.md) and [configs](task-configs/README.md), and use [input](task-inputs.md).