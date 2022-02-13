<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# Project Configs
<!--endTocHeader-->

Your tasks might have several key-value [configurations](./task/task-configs/README.md).

Furthermore, tasks might [share some of their configurations](./tasks/task-configs/shared-configs.md) with other tasks. To do that, you will need a project config.

Defining a project config is pretty straightforward. Here is the anatomy of a project config:

```yaml
configs:

  configName:                 # project config name
    someKey: someValue        # key-value
    otherKey: otherValue
```

You can reuse a project config by including it in your task. To do that you need to use `configRef` or `configRefs` property.

To use project config in your task, please see [task's shared configs](./task/task-configs/shared-configs.md).


<!--startTocSubTopic-->
<!--endTocSubTopic-->