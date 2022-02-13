<!--startTocHeader-->
[🏠](../README.md)
# 🧠 Core Concepts
<!--endTocHeader-->

Zaruba is a task runner and CLI Utilities.

There are several builtin [core-tasks](../core-tasks/README.md) and [utilities](../utilities/README.md) that you can use right away. But you can also define a [project](./project/README.md) containing custom [task](./project/task/README.md) definitions.

# Run a Task

To execute builtin [core-tasks](../core-tasks/README.md), you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

But, to execute any [project](./project/README.md) specific tasks, you need to be in the project directory first:

```bash
cd <project-name>
zaruba please <task-name>
```

# Run Many Tasks in Parallel

Zaruba allows you to run several tasks in parallel. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

# Set Task Environments

You can set `environments` for your tasks by performing:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

You can also load those `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

If you don't define any environment, Zaruba will load `.env` as default value.

> ⚠️ __WARNING:__ Parameter order matters, if you set an environment twice, Zaruba will only use the __last__ one.

# Set Task Inputs

Some tasks might dependes on [task inputs](./project/task/task-inputs.md). To set input values for your tasks, you can invoke:

```bash
zaruba please <task-name> -v <first-key=first-value> -v <second-key=second-value>
```

or shorter:


```bash
zaruba please <task-name> <first-key=first-value> <second-key=second-value>
```

or even use a `value files`.

```bash
zaruba please <task-name> -v <first-file.value.yaml> -v <second-file.value.yaml>
```

If you don't define any value, Zaruba will load `default.values.yaml` as default value.

> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ one.


# Run Tasks in Interactive Mode

When you run tasks in interactive mode, Zaruba will ask you to fill out [task inputs](./project/task/task-inputs.md) and `environments`. This is useful in case you cannot remember your `task inputs`/`environments`.

To run a task in interactive mode you can invoke:

```bash
zaruba please <task-name> -i
```

or

```bash
zaruba please <first-task-name> <second-task-name> -i
```

or

```
zaruba please -i
```

When you run a task in interactive mode, Zaruba will also ask you to set environments and input values.

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