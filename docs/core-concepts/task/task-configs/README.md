<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🔨 Task](../README.md)
# ⚙️ Task Configs
<!--endTocHeader-->

# Why You Need It

Suppose you have a task to run a http server:

```yaml
tasks:

  startServer:
    start:
      - bash
      - '-c'
      - python -m http.server 8080
    check:
      - bash
      - '-c'
      - |
        until nc -z localhost 8080
        do 
            sleep 2 && echo "not ready"
        done
        echo "ready"
```

As you can see, `startServer.start` and `startServer.check` uses the same port. This mean that you cannot edit `startServer.start` and `startServer.check` independently. You must always make sure that they use the same port.

# Using Config

`startServer.start` and `startServer.check` should always refer to the same port. Thus, it is better if you can take this port configuration somewhere.

To do this, you can add `startServer.configs.port`.

```yaml
tasks:

  startServer:
    configs:
      port: 8080
    start:
      - bash
      - '-c'
      - 'sleep 10 && python -m http.server ${ZARUBA_CONFIG_PORT}'
    check:
      - bash
      - '-c'
      - |
          until nc -z localhost ${ZARUBA_CONFIG_PORT}
          do 
            sleep 2 && echo "not ready"
          done
          echo "ready"
```

Now, you can use  `${ZARUBA_CONFIG_PORT}` to refer to `configs.port`.

# Accessing Task Configuration

There are two ways to access task configuration:

* `${ZARUBA_CONFIG_<CONFIG_NAME>}` variable.
* `{{ .GetConfig "configName" }}` [go template](../../go-template.md).

Let's take this a bit further:

```yaml
tasks:

  startServer:
    configs:
      port: 8080
      start: 'sleep 10 && python -m http.server ${ZARUBA_CONFIG_PORT}'
      check: | 
        until nc -z localhost ${ZARUBA_CONFIG_PORT}
        do 
          sleep 2 && echo "not ready"
        done
        echo "ready"
    start: [bash, -c, "${ZARUBA_CONFIG_START}"]
    check: [bash, -c, "${ZARUBA_CONFIG_CHECK}"]
```

Nice. now you even make `start` and `check` command configurable.

# Shared Config

In some cases, you might need to share your configuration among tasks. To do this, you need [project config](../../project/project-configs.md).

Please have a look at [shared config documentation](./shared-configs.md) for more information.


<!--startTocSubtopic-->
# Sub-topics
* [Shared Configs](shared-configs.md)
<!--endTocSubtopic-->