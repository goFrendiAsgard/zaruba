[⬅️ Table of Content](../../../../README.md)

# Task Configs

Suppose you have a task to run a server:

```yaml
tasks:

  startServer:
    start: [bash, -c, 'python -m http.server 8080']
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"']
```

Let's first break this down into something more readable:

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

As you might realize, `startServer.start` and `startServer.check` should use the same port.

Unfortunately, there is a great chance that someone edit the the script without aware of this fact.

Let's improve the task by making it more configurable.

First we should make a configuration using `configs` property, and use a builtin go template function `{{ .GetConfig "someProperty" }}`. You can learn more about Zaruba's go template [here](../../using-go-template.md).

Let's edit your script a little bit:

```yaml
tasks:

  startServer:
    config:
      port: 8080
    start:
      - bash
      - '-c'
      - 'sleep 10 && python -m http.server {{ .GetConfig "port" }}'
    check:
      - bash
      - '-c'
      - |
          until nc -z localhost {{ .GetConfig "port" }}
          do 
            sleep 2 && echo "not ready"
          done
          echo "ready"
```

Perfect, now anyone can edit `startServer.config.port` without screwing everything.

Let's take this a bit further:

```yaml
tasks:

  startServer:
    config:
      port: 8080
      start: 'sleep 10 && python -m http.server {{ .GetConfig "port" }}'
      check: | 
        until nc -z localhost {{ .GetConfig "port" }}
        do 
          sleep 2 && echo "not ready"
        done
        echo "ready"
    start: [bash, -c, '{{ .GetConfig "start" }}']
    check: [bash, -c, '{{ .GetConfig "check" }}']
```

Nice. so now you even make `start` and `check` command configurable.

# Shared Config

Furthermore, you can take out this configuration and put it as [project config](../../project-configs.md) so that you can share the configurations with other tasks.

To see how to do this, please have a look at [shared config documentation](./shared-configs.md).
