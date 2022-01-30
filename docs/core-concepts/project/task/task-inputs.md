[⬅️ Table of Content](../../../README.md)

# Task Inputs

There are two ways to configure how a task should be executed. The first one is using `envs` property. The other one is by using `inputs`.

If your application/service can be configured by using environment variable, it is always better to use `envs` property. Otherwise, you might find `inputs` is probably better.

Let's revisit our previous example:

```yaml
tasks:

  startServer:
    extend: zrbStartApp
    configs:
      httpPort: '{{ .GetEnv "HTTP_PORT" }}'
      start: 'sleep 10 && python -m http.server {{ .GetConfig "httpPort" }}'
      ports: '{{ .GetConfig "httpPort" }}'
    envs:
      HTTP_PORT:
        from: SERVER_HTTP_PORT
        default: 8080
```

Now if you want to make the delay configurable, you can surely use `inputs` property. But firstly, you have to declare the `inputs` first. For more information about `inputs`, you can visit [project inputs document](../project-inputs.md) later.

```yaml
inputs:
  
  serverDelay:
    prompt: Server delay
    options: [5, 10, 20]

tasks:

  startServer:
    extend: zrbStartApp
    inputs:
      - serverDelay
    configs:
      delay: '{{ .GetValue "serverDelay" }}'
      httpPort: '{{ .GetEnv "HTTP_PORT" }}'
      start: |
        sleep {{ .GetConfig "delay" }}
        python -m http.server {{ .GetConfig "httpPort" }}
      ports: '{{ .GetConfig "httpPort" }}'
    envs:
      HTTP_PORT:
        from: SERVER_HTTP_PORT
        default: 8080
```

Now you can run the task by invoking `zaruba please startServer serverDelay=5`:

```
❯ zaruba please startServer serverDelay=5
💀 🔎 Job Starting...
         Elapsed Time: 1.3µs
         Current Time: 06:56:40
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/playground/example
💀    🚀 updateProjectLinks   🔗 06:56:40.696 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:56:40.696 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/playground/example
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/playground/example
💀    🔎 startServer          🍏 06:56:41.003 📜 Waiting for port '3000'
💀    🚀 startServer          🍏 06:56:46.156 Serving HTTP on 0.0.0.0 port 3000 (http://0.0.0.0:3000/) ...
💀    🔎 startServer          🍏 06:56:47.02  📜 Port '3000' is ready
💀    🔎 startServer          🍏 06:56:47.02  🎉🎉🎉
💀    🔎 startServer          🍏 06:56:47.021 📜 Task 'startServer' is ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 6.6353767s
         Current Time: 06:56:47
         Active Process:
           * (PID=12643) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
```

Notice that the task is started at `06.56.41`, but the server is started at `06.56.46`. Because there is a 5 seconds delay.

Moreover, you can also set `serverDelay` interactively by invoking `zaruba please startServer -i`:

```
❯ zaruba please startServer -i
💀 Load additional value file
✔ 🏁 No
💀 Load additional env
✔ 🏁 No
💀 1 of 1) serverDelay
Search: █
? Server delay:
    Blank
    5
  ▸ 10
    20
    Let me type it!
```

Once you fill up the value, the server will run as expected.

```
❯ zaruba please startServer -i
💀 Load additional value file
✔ 🏁 No
💀 Load additional env
✔ 🏁 No
💀 1 of 1) serverDelay
✔ 10
💀 🔎 Job Starting...
         Elapsed Time: 1.6µs
         Current Time: 07:00:22
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/playground/example
💀    🚀 updateProjectLinks   🔗 07:00:22.274 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:00:22.274 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/playground/example
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/playground/example
💀    🔎 startServer          🍏 07:00:22.551 📜 Waiting for port '3000'
💀    🚀 startServer          🍏 07:00:32.696 Serving HTTP on 0.0.0.0 port 3000 (http://0.0.0.0:3000/) ...
💀    🔎 startServer          🍏 07:00:33.583 📜 Port '3000' is ready
💀    🔎 startServer          🍏 07:00:33.583 🎉🎉🎉
💀    🔎 startServer          🍏 07:00:33.583 📜 Task 'startServer' is ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 11.6126199s
         Current Time: 07:00:33
         Active Process:
           * (PID=13409) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
```