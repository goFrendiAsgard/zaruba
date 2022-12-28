<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🏝️ Project Envs
<!--endTocHeader-->


To run a task, you might need to set some [environment variables](../task/task-envs/README.md).

A project might contain several tasks. And you will find that some of those tasks might [share environment variables](../task/task-envs/shared-envs.md)

This is where project envs come into play.

Defining a project env is pretty straightforward. Here is the anatomy of a project env:

```yaml
envs:

  projectEnvName:                   # project env name
    SOME_VARIABLE:                  # local environment variable
      from: GLOBAL_VARIABLE         # global environment linked to the SOME_VARIABLE
      default: default value        # default value for SOME_VARIABLE
```

A single env consists of two keys:

* `from`: Global environment variable linked to current env.
* `default`: Default value for current env.

To use project env in your task, please refer to [task's shared envs](../task/task-envs/shared-envs.md).


<!--startTocSubtopic--><!--endTocSubtopic-->