<!--startTocHeader-->
[🏠](../../../../README.md) > [🧠 Core Concept](../../../README.md) > [🏗️ Project](../../README.md) > [Task](../README.md) > [Task Configs](README.md)
# Shared Configs
<!--endTocHeader-->

Your tasks might refer to several [project config](../../project-configs.md). To do this, you can use `configRef` or `configRefs` property:

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
      start: python -m http.server {{ .GetConfig "port" }}
  
  startAppContainer:
    extend: zrbStartDockerContainer
    configRefs:      # this task can use every configuration in `app` and `container` project config.
      - app
      - container
```

Both `startApp` and `startAppContainer` shared `app` project config. Additionally, `startAppContainer` also use `container` project config.


Behind the scene, Zaruba will embed refered project configs into tasks:

```yaml
tasks:

  startApp:
    extend: zrbStartApp
    configs:
      start: python -m http.server {{ .GetConfig "port" }}
      ports: 8080
  
  startAppContainer:
    extend: zrbStartDockerContainer
    configs:
      ports: 8080
      imageName: app-image
      containerName: appContainer
```

# Behavior

> ⚠️ __WARNING:__ `configrefs` order matters, if your `configRefs` contains the same configuration, Zaruba will use the __first__ one.

* You cannot use `configRef` and `configRefs` simultaneously in a single task. Use `configRef` if you want to refer to a single project config. Otherwise, use `configRefs`.

* `configRefs` order matters, if your `configRefs` contains the same configuration, Zaruba will use the first one.

* `configs` will always override `configRef` and `configRefs`.


<!--startTocSubtopic-->

<!--endTocSubtopic-->