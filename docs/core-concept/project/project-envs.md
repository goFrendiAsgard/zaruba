[⬅️ Table of Content](../../README.md)

# Project Envs


Your tasks might need [environment variables](./task/task-envs/README.md).

Furthermore, tasks might [share some of their environments](./tasks/task-envs/shared-envs.md) with other tasks. In order to do that, you will need a project env.

Defining a project env is pretty straightforward. Here is the anatomy of project env:

```yaml
envs:

  envName:
    SOME_VARIABLE:
      from: GLOBAL_VARIABLE
      default: default value
```

To use project environment in your tasks, you can either use `envRef` or `envRefs` property. Please see [task's shared envs](./task/task-envs/shared-envs.md) for more information.
