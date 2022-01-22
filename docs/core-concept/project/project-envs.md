[⬅️ Table of Content](../../README.md)

# Project Envs


Your tasks might need [environment variables](./task/task-envs/README.md).

Furthermore, tasks might [share some of their environments](./tasks/task-envs/shared-envs.md) with other tasks. To do that, you will need a project env.

Defining a project env is pretty straightforward. Here is the anatomy of a project env:

```yaml
envs:

  envName:
    SOME_VARIABLE:
      from: GLOBAL_VARIABLE
      default: default value
```

You can reuse a project env by including it in your task. To do that you need to use `envRef` or `envRefs` property. Please see [task's shared envs](./task/task-configs/shared-envs.md) for more information.