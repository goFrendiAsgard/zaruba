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
bash: line 1: cd: examples/playground/use-cases: No such file or directory
💀 🔎 Job Starting...
         Elapsed Time: 1.133µs
         Current Time: 21:48:49
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/newproject
💀    🚀 initProject          🚧 21:48:49.91  Initialized empty Git repository in /home/gofrendi/zaruba/docs/newproject/.git/
💀    🚀 initProject          🚧 21:48:49.914 🎉🎉🎉
💀    🚀 initProject          🚧 21:48:49.914 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 108.83157ms
         Current Time: 21:48:50
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 310.019525ms
         Current Time: 21:48:50
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
