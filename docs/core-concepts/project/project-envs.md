<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# Project Envs
<!--endTocHeader-->


Your tasks might need [environment variables](./task/task-envs/README.md).

Furthermore, tasks might [share some of their environments](./tasks/task-envs/shared-envs.md) with other tasks. To do that, you will need a project env.

Defining a project env is pretty straightforward. Here is the anatomy of a project env:

```yaml
envs:

  envName:
    SOME_VARIABLE:                  # Your env name
      from: GLOBAL_VARIABLE         # Global environment linked to the env name
      default: default value        # Default value for env name
```

A single env consists of two keys:

* `from`: Global environment variable linked to current env.
* `default`: Default value for current env. Default value will be used if `from` variable is either undefined or doesn't exist.

You can reuse a project env by including it in your task. To do that you need to use `envRef` or `envRefs` property.

To use project env in your task, please see [task's shared envs](./task/task-envs/shared-envs.md).


<!--startTocSubTopic-->
<!--endTocSubTopic-->