<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🔨 Task](../README.md) > [Task Envs](README.md)
# Shared Envs
<!--endTocHeader-->


Your tasks might refer to several [project env](../../project-envs.md). To do this, you can use `envRef` or `envRefs` property:

```yaml
envs:

  app:
    HTTP_PORT:
      from: APP_HTTP_PORT
      default: 8080

  docker:
    DOCKER_CERT_PATH:
      from: DOCKER_CERT_PATH
      default: ./certificate.pem

tasks:

  startApp:
    extend: zrbStartApp
    envRef: app   # this task can use every environment in `app` project env.
    configs:
      start: python -m http.server {{ .GetEnv "HTTP_PORT" }}
      ports: '{{ .GetEnv "HTTP_PORT" }}'
  
  startAppContainer:
    extend: zrbStartDockerContainer
    envRefs:      # this task can use every environment in `app` and `docker` project env.
      - app
      - docker
    configs:
      imageName: app-image
      containerName: appContainer
      ports: '{{ .GetEnv "HTTP_PORT" }}'
```

Both `startApp` and `startAppContainer` shared `app` project env. Additionally, `startAppContainer` also use `container` project env.


Behind the scene, Zaruba will embed refered project envs into tasks:

```yaml
tasks:

  startApp:
    extend: zrbStartApp
    envs:
      HTTP_PORT:
        from: APP_HTTP_PORT
        default: 8080
    configs:
      start: python -m http.server {{ .GetEnv "port" }}
      ports: '{{ .GetEnv "HTTP_PORT" }}'
  
  startAppContainer:
    extend: zrbStartDockerContainer
    envs:
      HTTP_PORT:
        from: APP_HTTP_PORT
        default: 8080
      DOCKER_CERT_PATH:
        from: DOCKER_CERT_PATH
        default: ./certificate.pem
    configs:
      imageName: app-image
      containerName: appContainer
      ports: '{{ .GetEnv "HTTP_PORT" }}'
```

# Behavior

> ⚠️ __WARNING:__ `envrefs` order matters, if your `envRefs` contains the same environment, Zaruba will use the __first__ one.

* You cannot use `envRef` and `envRefs` simultaneously in a single task. Use `envRef` if you want to refer to a single project env. Otherwise, use `envRefs`.

* `envRefs` order matters, if your `envRefs` contains the same environment, Zaruba will use the first one.

* `envs` will always override `envRef` and `envRefs`.


<!--startTocSubTopic-->
<!--endTocSubTopic-->