<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# 🏗️ Create a Project
<!--endTocHeader-->

The recommended way to create a project is by invoking `zaruba please initProject`:

__Example:__

<!--startCode-->
```bash
mkdir -p examples/playground/use-cases/newProject
cd examples/playground/use-cases/newProject
zaruba please initProject

tree
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.195µs
         Current Time: 21:57:07
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject
💀    🚀 initProject          🚧 21:57:07.089 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject/.git/
💀    🚀 initProject          🚧 21:57:07.093 🎉🎉🎉
💀    🚀 initProject          🚧 21:57:07.093 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 111.087701ms
         Current Time: 21:57:07
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 312.40167ms
         Current Time: 21:57:07
zaruba please initProject  
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```````
</details>
<!--endCode-->

# Initial Project Structure

Once created, you will have two files:

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

Please visit [run task with custom environments documentation](../run-task/run-task-with-custom-environments.md) for more information.

## Run Task with Custom Input Values

```bash
zaruba please <task-name> [other-task-names...] [-v custom-values.yaml] [-v key=value]
```

Please visit [run task with custom values documentation](../run-task/run-task-with-custom-values.md) for more information.


## Synchronize Environments

```bash
zaruba please syncEnv
```

Please visit [syncrhonize environments documentation](./synchronize-environments.md) for more information.

<!--startTocSubTopic-->
<!--endTocSubTopic-->
