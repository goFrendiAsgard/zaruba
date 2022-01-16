[⬅️](../README.md)

# Core concept

Zaruba is a task runner and CLI Utilities.

It has several builtin [core-tasks](../core-tasks/README.md) that you can execute from everywhere, as well as several builtin [utilities](../utilities/README.md).

# Run a Task

To execute a `task`, you can invoke:

```bash
zaruba please <task-name>
```

Beside the builtin [core-tasks](../core-tasks/README.md), you can also create a [project](project/README.md) and define [project-specific tasks](../core-tasks/project/task/README.md).

To execute any [project-specifict task](../core-tasks/project/task/README.md) you need to be in the project directory first:

```bash
cd <project-name>
zaruba please <task-name>
```

# Run Multiple Tasks

Zaruba allows you to run several tasks in parallel. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

# Set Task Environments

While running tasks, you can also set environments for those tasks by performing:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

Additionaly, you can also load environments from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

By default, zaruba will always load `.env` first before loading any user-defined environment.

Please note that the order is matter, the first one will always be overridden by the later.

# Set Task Inputs

Some tasks might dependes on [task inputs](core-concepts/task/task-inputs.md). To set input values for your tasks, you can invoke:

```bash
zaruba please <task-name> -v <first-key=first-value> -e <second-key=second-value>
```

or

```bash
zaruba please <task-name> -e <first-file.value.yaml> -e <second-file.value.yaml>
```

By default, zaruba will always load `default.values.yaml` first before loading any user-defined values.

Please note that the order is matter, the first one will always be overridden by the later.

# Run Tasks Interactively

To run task interactively you can invoke:

```bash
zaruba please <task-name> -i
zaruba please <first-task-name> <second-task-name> -i
```

or even

```
zaruba plaese -i
```

so that you can choose the task name later.

When you run a task interactively, Zaruba will also ask you to set environments and input values.

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