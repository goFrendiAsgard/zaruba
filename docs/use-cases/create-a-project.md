<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# Create a Project
<!--endTocHeader-->

The recommended way to create a project is by invoking `zaruba please initProject`:

```bash
❯ mkdir myproject

❯ cd myproject

❯ zaruba please initProject
💀 🔎 Job Starting...
         Elapsed Time: 1.2µs
         Current Time: 07:10:25
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/playground/myproject
💀    🚀 initProject          🚧 07:10:25.647 Initialized empty Git repository in /home/gofrendi/playground/myproject/.git/
💀    🚀 initProject          🚧 07:10:25.654 🎉🎉🎉
💀    🚀 initProject          🚧 07:10:25.654 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 311.8279ms
         Current Time: 07:10:25
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 512.9968ms
         Current Time: 07:10:25
```

# Initial Project Structure

Once created, you will have two files:

```
❯ tree
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```

* `default.values.yaml` is your default project value
* `index.zaruba.yaml` is the entry point of your zaruba script.

# Managing Your Project

## Run Tasks

```bash
zaruba please <task-name> [other-task-names...]
```

When you run many tasks at once, Zaruba will run your tasks in parallel. Please visit [core concept documentation](../core-concepts/README.md#run-many-tasks-in-parallel) for more information.

## Run Task with Custom Environments

```bash
zaruba please <task-name> [other-task-names...] [-e custom-environment.env] [-e KEY=value]
```

Please visit [core concept documentation](../core-concepts/README.md#set-task-environments) for more information.

## Run Task with Custom Input Values

```bash
zaruba please <task-name> [other-task-names...] [-v custom-values.yaml] [-v key=value]
```

Please visit [core concept documentation](../core-concepts/README.md#set-task-inputs) for more information.


## Synchronize Environments

```bash
zaruba please syncEnv
```

Please visit [core concept documentation](../core-concepts/project/task/task-envs/README.md#synchronize-tasks-environments) for more information.

# Next Step

Your next step is depending on your uses case:

* If you need to add third party services like MySQL, Redis, Kafka, etc, then you need to [add third party services](./add-third-party-service/README.md) to your project.
* If you need to create an application from scratch, then you can [generate new application](./generate-new-application/README.md).
* You can also add existing external application as [subrepo of your project](./add-subrepo.md). In this case you also need to [add runner for the existing application](./add-runner-for-existing-application/README.md).

<!--startTocSubTopic-->
<!--endTocSubTopic-->
