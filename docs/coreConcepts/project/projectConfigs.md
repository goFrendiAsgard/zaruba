<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# ⚙️ Project Configs
<!--endTocHeader-->

You can control a task's behavior by setting its [configuration](../task/task-configs/README.md).

A project might contain several tasks. And you will find that some of those tasks might [share configurations](../task/task-configs/shared-configs.md).

This is where project configs come into play.

Defining a project config is pretty straightforward. Here is the anatomy of a project config:

```yaml
configs:

  projectConfigName:          # project config name
    someKey: someValue        # key-value
    otherKey: otherValue
```

To use project config in your task, please refer to [task's shared configs](../task/task-configs/shared-configs.md).


<!--startTocSubtopic-->

<!--endTocSubtopic-->