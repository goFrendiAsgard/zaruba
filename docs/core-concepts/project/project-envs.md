<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🏝️ Project Envs
<!--endTocHeader-->


A task might need [environment variables](../task/task-envs/README.md).

Since a project might contains several tasks, you will find that some tasks might [share some their environment variables](../task/task-envs/shared-envs.md)

This is where project envs come into play.

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

To use project env in your task, please see [task's shared envs](../task/task-envs/shared-envs.md).


<!--startTocSubTopic-->
<!--endTocSubTopic-->