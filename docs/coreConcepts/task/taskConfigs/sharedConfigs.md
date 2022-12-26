<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🔨 Task](../README.md) > [⚙️ Task Configs](README.md)
# Shared Configs
<!--endTocHeader-->

Some of your taks might have overlapping configurations. In this case, you can make your tasks refer to the same [project config](../../project/project-configs.md).

To do this, you need to use `configRef` or `configRefs` property:

```yaml
configs:

  app:
    ports: 8080

  container:
    imageName: app-image
    containerName: appContainer

tasks:

  startApp:
    extend: zrbStartApp
    configRef: app   # this task can use every configuration in `app` project config.
    configs:
      start: python -m http.server ${ZARUBA_CONFIG_PORT}
  
  startAppContainer:
    extend: zrbStartDockerContainer
    configRefs:      # this task can use every configuration in `app` and `container` project config.
      - app
      - container
```

In the example, we have tw project configs:

* `app` (used by `startApp` and `startAppContainer`)
* `container` (used by `startAppContainer`)

# How it Works

When you refer a `project config` into your task, Zaruba will merge the configurations on run time. Our previous example will be rendered as:

```yaml
tasks:

  startApp:
    extend: zrbStartApp
    configs:
      start: python -m http.server ${ZARUBA_CONFIG_PORT}
      ports: 8080
  
  startAppContainer:
    extend: zrbStartDockerContainer
    configs:
      ports: 8080
      imageName: app-image
      containerName: appContainer
```

# ⚠️ Behavior

* You cannot use `configRef` and `configRefs` in a single task. Use `configRef` if you want to refer to a single project config. Otherwise, use `configRefs`.

* `configRefs` order matters. If your `configRefs` contains the same configuration, Zaruba will use the first one.

* task's `configs` will always override `configRef` and `configRefs`.


<!--startTocSubtopic-->
<!--endTocSubtopic-->