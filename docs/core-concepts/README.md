<!--startTocHeader-->
[🏠](../README.md)
# 🧠 Core Concepts
<!--endTocHeader-->

Zaruba is a task runner and CLI Utilities.

There are several builtin [core-tasks](../core-tasks/README.md) and [utilities](../utilities/README.md) that you can use right away. But you can also define a [project](./project/README.md) containing custom [task](./project/task/README.md) definitions.


# Explain a Task

To explain a task, you can add `-x` flag:

```bash
zaruba please <task-name> -x
```

# Configure Zaruba by using Environment Variables

There are several environment variables you can use to control Zaruba's behavior:

* `ZARUBA_HOME` Location of your Zaruba installation directory. Default to `${HOME}/.zaruba`.
* `ZARUBA_BIN` Location of your Zaruba executable binary. Default to `${HOME}/.zaruba/zaruba`.
* `ZARUBA_SHELL` The shell zaruba used to execute shell scripts. Default to `bash`.
* `ZARUBA_SCRIPTS` List of zaruba scripts that are going to be available from everywhere. Use colon (`:`) to separate the scripts.

# Understanding Project

Please visit [project documentation](./project/README.md) to learn more about project, or visit [here](../use-cases/create-a-project.md) if you want to see how to create a project.

<!--startTocSubTopic-->
# Sub-topics
* [🏗️ Project](project/README.md)
  * [🧬 Project Anatomy](project/project-anatomy.md)
  * [🧳 Includes](project/includes.md)
  * [🔤 Project Inputs](project/project-inputs.md)
  * [⚙️ Project Configs](project/project-configs.md)
  * [🏝️ Project Envs](project/project-envs.md)
  * [🔨 Task](project/task/README.md)
    * [Task Anatomy](project/task/task-anatomy.md)
    * [Simple Command](project/task/simple-command.md)
    * [Long Running Process](project/task/long-running-process.md)
    * [Task Inputs](project/task/task-inputs.md)
    * [Task Configs](project/task/task-configs/README.md)
      * [Shared Configs](project/task/task-configs/shared-configs.md)
    * [Task Envs](project/task/task-envs/README.md)
      * [Shared Envs](project/task/task-envs/shared-envs.md)
    * [Extend task](project/task/extend-task.md)
    * [Define task dependencies](project/task/define-task-dependencies.md)
  * [🐹 Using Go Template](project/using-go-template.md)
<!--endTocSubTopic-->