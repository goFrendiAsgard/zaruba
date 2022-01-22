[⬅️ Table of Content](../../README.md)

# Includes

Zaruba scripts, like any other scripts/source code can grow into humongous spaghetti code in no time. Before it happens, you need to act as quickly as possible by dividing your scripts based on the domains and functionalities.

To do that, you need `includes`.

# Behavior of Includes

Take note that `includes` in Zaruba works differently from python/C. Every scripts included in your `index.zaruba.yaml` (directly or indirectly) will be able to access each other's resources.

For example, suppose you have the following structure:

```
.
├── index.zaruba.yaml
├── configs.yaml
└── tasks.yaml
```

Suppose `index.zaruba.yaml` includes both `configs.yaml` and `tasks.yaml` like this:

```yaml
# file: index.zaruba.yaml
includes:
  - configs.yaml
  - tasks.yaml
```

If your `configs.yaml` contains a configuration named `myConfig`:

```yaml
# file: configs.yaml
configs:
  myConfig:
    sacredNumber: 73
```

then you will be able to access `myConfig` from inside `tasks.yaml` like this:

```yaml
# file: tasks.yaml
tasks:
  myTask:
    extend: zrbRunShellScript
    configRef: myConfig # this refer to project config defined in configs.yaml
    config:
      start: echo '{{ .GetConfig "sacredNumber" }}'
```

# Project Directory Structure

Although you can structure your project as you like, typically a sane zaruba project looks like this:

```
.
├── index.zaruba.yaml          # script entry point
├── zaruba-tasks
│   ├── application            # collection of scripts to manage application
|   |   ├── index.yaml         # application's entry point
|   |   ├── configs.yaml       # application's shared config
|   |   ├── inputs.yaml        # application's shared input
|   |   └── tasks.yaml         # application's tasks
│   └── otherApplication
|       ├── index.yaml
|       ├── configs.yaml
|       ├── inputs.yaml
|       └── tasks.yaml
├── application                # application's source code
└── otherApplication
```

An `index.zaruba.yaml` should only contains `includes` and wrapper `tasks` like this:

```yaml
# file: index.zaruba.yaml
includes:
  - zaruba-tasks/application/index.yaml
  - zaruba-tasks/otherApplication/index.yaml

tasks:

  start:
    dependencies:
      - startApplication
      - startOtherApplication

  startContainers:
    dependencies:
      - startApplicationContainer
      - startOtherApplicationContainer
```

Meanwhile, application's `index.yaml` should includes `configs.yaml`, `inputs.yaml`, and `tasks.yaml`:

```yaml
# file: application/index.yaml
includes:
  - configs.yaml
  - inputs.yaml
  - tasks.yaml
```

By using this directory structure, your project will be more maintainable. Also, you will be able to manage your application tasks/resources independent to each other.
