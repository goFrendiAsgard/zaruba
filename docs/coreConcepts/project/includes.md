<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🧳 Includes
<!--endTocHeader-->

Over time your scripts and source code tend to grow bigger. At some point, you won't be able to keep everything in a single file.

Usually, people will divide scripts based on their domains and functionalities.

By using `includes` keyword, you can also do this with your Zaruba scripts as well.

# Behavior of Includes

> __TL;DR:__ You should include everything in `index.zaruba.yaml`

Unlike in Python or C, `includes` in Zaruba works in a different way.

Every script you include in `index.zaruba.yaml` will be accessible from each other.

For example, let's say you split your Zaruba scripts in three different files:

```
.
├── index.zaruba.yaml
├── configs.yaml
└── tasks.yaml
```

__index.zaruba.yaml__

```yaml
# file: index.zaruba.yaml
includes:
  - configs.yaml
  - tasks.yaml
```

__configs.yaml__

```yaml
# file: configs.yaml
configs:
  myConfig:
    sacredNumber: 73
```

__tasks.yaml__

```yaml
# file: tasks.yaml
tasks:
  myTask:
    extend: zrbRunShellScript
    configRef: myConfig # this refer to project config defined in configs.yaml
    config:
      start: echo "${ZARUBA_CONFIG_SACRED_NUMBER}"
```

You can see that `tasks.yaml` doesn't explicitly includes `configs.yaml`. But, you can still access `myConfig` from inside `myTask`.

This is possible since you already includes `configs.yaml` and `tasks.yaml` in `index.zaruba.yaml`

# Convention

By convention, you should arrange your scripts as follow:

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

An `index.zaruba.yaml` should only contain `includes` and wrapper `tasks`. Please look at the following example:

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

Application's `index.yaml` should includes `configs.yaml`, `inputs.yaml`, and `tasks.yaml`:

```yaml
# file: application/index.yaml
includes:
  - configs.yaml
  - inputs.yaml
  - tasks.yaml
```

By following this convention, you will make your project more predictable and manageable.


<!--startTocSubtopic-->

<!--endTocSubtopic-->