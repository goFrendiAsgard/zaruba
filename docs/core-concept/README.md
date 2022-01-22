[⬅️ Table of Content](../README.md)

# 🧠 Core Concept

Zaruba is a task runner and CLI Utilities.

Beside several builtin [core-tasks](../core-tasks/README.md) and [utilities](../utilities/README.md), you can also define a [project](./project/README.md) containing your own resources and [task](./project/task/README.md) definitions.

# Run a Task

To execute builtin [core-tasks](../core-tasks/README.md), you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

However, to execute any [project](./project/README.md) specific tasks, you need to be in the project directory first:

```bash
cd <project-name>
zaruba please <task-name>
```

# Run Multiple Tasks in Parallel

Zaruba allows you to run several tasks in parallel. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

# Set Task Environments

You can manually set `environments` for your tasks by performing:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

Additionaly, you can also load those `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

By default, zaruba will always load `.env` first before loading any user-defined environment.

Please note that the order is matter, the first one will always be overridden by the later.

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

By default, zaruba will always load `default.values.yaml` first before loading any user-defined values.

Please note that the order is matter, the first one will always be overridden by the later.

# Run Tasks Interactively

When you run tasks interactively, Zaruba will ask you to fill out [task inputs](./project/task/task-inputs.md) and `environments`. This is useful in case you cannot remember your `task inputs`/`environments`.

In order to run task interactively you can invoke:

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