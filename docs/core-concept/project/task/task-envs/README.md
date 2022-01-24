[⬅️ Table of Content](../../../../README.md)

# Task Envs

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

Now you have an environment variable named `HTTP_PORT`. By default it's value is `8080`, but you can override it by using global environment variable `SERVER_HTTP_PORT`

Let's set `SERVER_HTTP_PORT` to `3000` and start the server:

```
gofrendi@sanctuary [15:35:50] [~/playground/example]
-> % export SERVER_HTTP_PORT=3000
gofrendi@sanctuary [15:36:04] [~/playground/example]
-> % zaruba please startServer
💀 🔎 Job Starting...
         Elapsed Time: 59.3µs
         Current Time: 15:37:15
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/playground/example
💀    🚀 updateProjectLinks   🔗 15:37:16.144 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 15:37:16.144 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/playground/example
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/playground/example
💀    🔎 startServer          🍏 15:37:16.44  📜 Waiting for port '3000'
```

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

# Shared Envs

Furthermore you can also take out the environments and put it as [project env](../../project-envs.md) so that you can share it with other tasks.

To see how to do this, please have a look at [shared env documentation](./shared-envs.md).