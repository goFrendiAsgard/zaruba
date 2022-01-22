[⬅️ Table of Content](../../README.md)

# Project Configs

Your tasks might have several key-value [configurations](./task/task-configs/README.md).

Furthermore, tasks might [share some of their configurations](./tasks/task-configs/shared-configs.md) with other tasks. In order to do that, you will need a project config.

Defining a project config is pretty straightforward. Here is the anatomy of project config:

```yaml
configs:

  configName:
    someKey: someValue
    otherKey: otherValue
```

To use project configuration in your tasks, you can either use `configRef` or `configRefs` property. Please see [task's shared configs](./task/task-configs/shared-configs.md) for more information.
