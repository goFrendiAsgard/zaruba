<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🔨 Task](../README.md)
# 🏝️ Task Envs
<!--endTocHeader-->


At some degree, you might need to configure your application by using environment variables. Zaruba allows you to do this by utilizing `envs` property.

Let's have a look at this example:

```yaml
tasks:

  startServer:
    extend: zrbStartApp
    configs:
      httpPort: 8080
      start: 'sleep 10 && python -m http.server {{ .GetConfig "httpPort" }}'
      ports: '{{ .GetConfig "httpPort" }}'
```

You might wonder how to override the port without editing the task over and over again.

There are two approach to solve this. The first one is by using `envs`, while the other one is by using `inputs`.

Let's try the first approach by adding `envs` property:


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

Now you have an environment variable named `HTTP_PORT`. By default its value is `8080`, but you can override it by using global environment variable `SERVER_HTTP_PORT`

Let's set `SERVER_HTTP_PORT` to `3000` and start the server:

<!--startCode-->
```bash
cd examples/core-concepts/task/task-envs
export SERVER_HTTP_PORT=3000
zaruba please startServer -t -w 1s
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.139µs
         Current Time: 20:09:08
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/core-concepts/task/task-envs
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/zaruba/docs/examples/core-concepts/task/task-envs
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/zaruba/docs/examples/core-concepts/task/task-envs
💀    🔎 startServer          🍏 🔎 Waiting for port '3000'
💀    🚀 startServer          🍏 Serving HTTP on 0.0.0.0 port 3000 (http://0.0.0.0:3000/) ...
💀    🔎 startServer          🍏 🔎 Port '3000' is ready
💀    🔎 startServer          🍏 🎉🎉🎉
💀    🔎 startServer          🍏 📜 Task 'startServer' is ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 11.243469682s
         Current Time: 20:09:19
         Active Process:
           * (PID=24849) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🍏 'startServer' service (PID=24849)
💀    🚀 startServer          🍏 
💀    🚀 startServer          🍏 Keyboard interrupt received, exiting.
💀    🚀 startServer          🍏 🎉🎉🎉
💀    🚀 startServer          🍏 📜 Task 'startServer' is started
💀 🔎 Job Ended...
         Elapsed Time: 13.347960128s
         Current Time: 20:09:22
💀 🔥 🍏 'startServer' service exited without any error message
zaruba please startServer   -t -w 1s
```````
</details>
<!--endCode-->


The server is now running on port `3000`.


> 💡 __TIPS:__  Configuring application/services using environment variables is a very common practice. If you are building an application/service, please make sure it is configurable.

# Using Environment in Your Application

Most modern (and ancient) programming language allows you to read from environment variable.

A common practice among developers when running application locally is by putting the environments in an `.env` file, and load it before run the application:

```bash
source .env
python main.py
```

Here are some resources that you might find useful:

* [Read environment in go](https://pkg.go.dev/os#Getenv)
* [Read environment in nodejs](https://nodejs.org/api/process.html#processenv)
* [Read environment in python](https://docs.python.org/3/library/os.html#os.getenv)

# Synchronize Task's Environment

You can ask Zaruba to parse environment files in task's `syncEnvLocation`/`location` and update task/project envs.

If `syncEnv` is set to true, the task's environment will be synchronized whenever you invoke:

 ```bash
 zaruba please syncEnv
 ```

# Shared Envs

Furthermore you can also take out the environments and put it as [project env](../../project/project-envs.md) so that you can share it with other tasks.

To see how to do this, please have a look at [shared env documentation](./shared-envs.md).


<!--startTocSubTopic-->
# Sub-topics
* [Shared Envs](shared-envs.md)
<!--endTocSubTopic-->