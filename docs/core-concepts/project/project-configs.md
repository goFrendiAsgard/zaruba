<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# ⚙️ Project Configs
<!--endTocHeader-->

A task might have some key-value [configurations](../task/task-configs/README.md).

Since a project might contains several tasks, you will find that some tasks might [share some of their configurations](../task/task-configs/shared-configs.md).

This is where project configs come into play.

Defining a project config is pretty straightforward. Here is the anatomy of a project config:

```yaml
configs:

  configName:                 # project config name
    someKey: someValue        # key-value
    otherKey: otherValue
```

To use project config in your task, please refer to [task's shared configs](../task/task-configs/shared-configs.md).


<!--startTocSubTopic-->
<!--endTocSubTopic-->