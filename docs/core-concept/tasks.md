[⬆️](./README.md)

# Tasks

Tasks are the core of your zaruba tasks. A task define what Zaruba can do and how to do it.

Let's start by creating a very simple script.

```
gofrendi@sanctuary [17:07:47] [~/playground/example]
-> % cat > index.zaruba.yaml << EOF
heredoc> tasks:
heredoc>   sayHello:
heredoc>     start: [figlet, hello]
heredoc> EOF
```

and executing it:

```
gofrendi@sanctuary [17:12:33] [~/playground/example]
-> % zaruba please sayHello
💀 🔎 Job Starting...
         Elapsed Time: 1.3µs
         Current Time: 17:12:49
💀 🏁 Run 🍏 'sayHello' command on /home/gofrendi/playground/example
💀    🚀 sayHello             🍏 17:12:49.475  _          _ _
💀    🚀 sayHello             🍏 17:12:49.475 | |__   ___| | | ___
💀    🚀 sayHello             🍏 17:12:49.475 | '_ \ / _ \ | |/ _ \
💀    🚀 sayHello             🍏 17:12:49.475 | | | |  __/ | | (_) |
💀    🚀 sayHello             🍏 17:12:49.475 |_| |_|\___|_|_|\___/
💀    🚀 sayHello             🍏 17:12:49.475
💀 🎉 Successfully running 🍏 'sayHello' command
💀 🔎 Job Running...
         Elapsed Time: 106.3051ms
         Current Time: 17:12:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 216.7833ms
         Current Time: 17:12:49
```

Perfect.

Now let's see how this really works:

```yaml
tasks:
  sayHello:
    start: [figlet, hello]
```

The script contains a single task named `sayHello` that has `start` property. This `start` property is an array containing two elements: `figlet`, and `hello`. The first element of this array should refer to a binnary or executable path, while the others elements should contains parameters to run the executable.

At the beginning, you will find this approach looks too complicated, but please bear with me until the end to see how this actually gives you advantage later.

Now let's dive into the anatomy of a task:

## Task anatomy

A task might contains several properties. We will learn the purpose of each properties later, but for now here is a list of those properties:

* `icon`: Emoji representing a task. You can put any character as task's icon, but there is no better way than copy-pasting your favorit emoji from [emojipedia](https://emojipedia.org/).
* `location`: The location path you want your tasks to be run from. If you don't specify task's location, then Zaruba will use your current location.
* `description`: The description of your task. Can be multiline.
* `extend` or `extends`: Your task's parent name(s). These properties cannot be used simultaneously.
    * `extend`: Your task's parent name
    * `extends`: List of your tasks's parent names.
* `timeout`: The duration Zaruba should wait before a task is considered as failing. Timeout contains a possitive number or zero and followed by any of this suffix: "ns", "us" (or "µs"), "ms", "s", "m", "h".
* `private`: Boolean value to represent whether your task is private or not. Private tasks are interactively inaccessible. Usually private tasks act as template to be extended by other tasks.
* `inputs`: List of input names you want to associate with your task.
* `dependencies`: Task dependencies. Zaruba will make sure that all dependencies of a task are completed before starting one.
* `envRef` or `envRefs`: Your task's environment reference(s). These properties cannot be used simultaneously.
    * `envRef`: Name of environment you want to use for your task.
    * `envRefs`: List of environments you want to use for your task.
* `env`: Task environment. Any environment you declare in this property will override anything defined in your `envRef` or `envRefs`.
* `configRef` or `configRefs`: Your task's configuration reference(s). These properties cannot be used simultaneously.
    * `configRef`: Name of configuration you want to use for your task.
    * `configRefs`: List of configuration you want to use for your task.
* `config`: Task configuration. Any configuration you declare in this property will override `configRef` or `configRefs`.
* `start`: Task's start command.
* `check`: Command to check whether start command can be considered "running" or not. This property is used to run a long running service.

Let's see an example:

```yaml
tasks:

  taskName:
    icon: 🎻
    location: ./task-location
    description: task's description
    extend: parentTaskName # use "extends" for multiple values
    timeout: 1h
    private: false
    inputs: [] # list of input's name
    dependencies: [] # tasks's dependencies
    envRef: envRefName # use "envRefs" for multiple values
    envs:
      SOME_ENV:
        default: defaultValue
        from: GLOBAL_ENV
    configRef: configRefName # use "configRefs" for multiple values
    configs:
      configName: configValue
    start: [figlet, hello] # start command
    check: [] # check command
```

## Running a simple command vs starting a long running service

To make things easier, let's first agree on two terminology: `simple command`, and `long running service`.

### Simple command

Simple command is something you run and considered `done` once the process has been ended.

For example, `python -c "print('hello')"` is a command:

```
gofrendi@sanctuary [09:45:16] [~/playground/example]
-> % python -c "print('hello')"
hello
gofrendi@sanctuary [10:27:21] [~/playground/example]
-> %
```

We can see that once the process has been ended, the command are done. When you compile your Go/Typescript/Java application you are definitely running a command (even if you don't really open a terminal).

### Long running service

Simple command is pretty simple and straightforward. But long running service on the other hand, has a very different nature.

A long running service might keep running forever until it is killed. It considered `ready` when it serve what it intended to.

Web server and database server are definitely long running services. Those services might run in the background automatically, thus less obviously visible by end user. But invisible is not equal to inexistance. In fact, you can find long running service everywhere. Arguably, your OS is big long running service. It is always there, waiting for your inputs or external events, and act accordingly.

Let's try running a static web service by invoking `python -m http.server 8080`.

```
gofrendi@sanctuary [10:31:47] [~/playground/example]
-> % python -m http.server 8080
Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
```

Now you can see that the process doesn't immediately quit once it is started. It will wait any incoming http request to port 8080, giving a response, and wait again until the end of time (or until you press ctrl + C).

Okay coool, now how do you make sure that a service is ready?

You can make sure a service is ready by giving it a request and observe it's response. In our case, you can verify that the service is ready by openning a browser and visit `http://localhost:8080`.

Making sure that a service is ready can be tricky since `running` doesn't mean `ready`. To make it more complicated, even if a service is considered as `ready`, it doesn't always be in that state forever. Runtime error might occurred, and your service might stop serving eventhough it is still `running`. 

Orchestration system like kubernetes overcome this problem by periodically sending request to your services using `liveness` and `readiness` probe. But that's going to be another topic of discussion.

## Checking readiness of long running service

In the previous section you have see that handle a service and make sure it is already running can be a bit challenging.

Under the hood, Zaruba make sure that your service is ready by running two commands simultaneously. The first command is responsible to run the service, while the other one is responsible to check it's readiness. 

Let's see how this work in the low level.

First of all, you will need two terminals in the same computer. You can also use tmux/screen if you are familiar with those tools.

__Starter__

In your first terminal, you can spawn this command `sleep 10 && python -m http.server 8080`: 

```
gofrendi@sanctuary [11:07:54] [~/playground/example]
-> % sleep 10 && python -m http.server 8080
Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
```

This command ask the computer to wait for 10 seconds before starting the web server. We use this to emulate real world cases. Some server might even take more than a minute to be ready. Also, those 10 seconds give you enough time to open the second terminal and invoke the service checker (if you are really that quick).

__Checker__

Our service checker contains a single loop to check whether `localhost:8080` is up and serving. In order to start the checker, you can invoke this in your second terminal `until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"`:

```
gofrendi@sanctuary [11:07:56] [~/playground/example]
-> % until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"
not ready
not ready
not ready
not ready
ready
gofrendi@sanctuary [11:32:44] [~/playground/example]
-> %
```

Great, now you can make sure that your service is really `ready` before deal with it any further.

> 💡 __TIPS:__  if you find the service is already started before you are able to start the checker, please feel free to change the sleep duration of the server (e.g: `sleep 30 && python -m http.server 8080`, will make the computer wait for 30 seconds before starting the server)

To see how our starter and checker works, let's take a look on this diagram:

![](images/starter-and-checker.png)


## Starting a long running service with Zaruba

Unless you are a [starcraft](https://starcraft2.com/en-us/) pro player, probably running multiple terminals and tmux panels is not a very good idea.

You might also want to run the server in the background or make a docker container for this simple use case. But let's not do that.

We will use Zaruba instead.

### Low level approach

First you declare this script in your `index.zaruba.yaml`

```yaml
tasks:

  startServer:
    start: [bash, -c, 'sleep 10 && python -m http.server 8080']
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"']
```

then, you can invoke `zaruba please startServer`.

```
gofrendi@sanctuary [12:11:35] [~/playground/example]
-> % zaruba please startServer
💀 🔎 Job Starting...
         Elapsed Time: 1.3µs
         Current Time: 12:11:54
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/playground/example
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/playground/example
💀    🔎 startServer          🍏 12:11:56.908 not ready
💀    🔎 startServer          🍏 12:11:58.91  not ready
💀    🔎 startServer          🍏 12:12:00.912 not ready
💀    🔎 startServer          🍏 12:12:02.92  not ready
💀    🔎 startServer          🍏 12:12:04.927 not ready
💀    🔎 startServer          🍏 12:12:06.932 not ready
💀    🔎 startServer          🍏 12:12:06.936 ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 12.1374875s
         Current Time: 12:12:07
         Active Process:
           * (PID=16029) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
```

As you can see, this is what `check` property actually is for. It tells Zaruba how to check your service readiness. Any task with `check` property set will be considered as long running service.
    
> 💡 __TIPS:__  You might wonder why the server log doesn't show up unless you terminate it with `ctrl + c`. This is happened because of python buffering mechanism. To turn off this feature, you can set `PYTHONUNBUFFERED` to `1`. (i.e: by using this as start command, `start: [bash, -c, 'sleep 10 && export PYTHONUNBUFFERED=1 && python -m http.server 8080']`)


### Higher level approach

The previous approach looks good. But in most cases, you will most likely see this instead:

```yaml
tasks:

  startServer:
    extend: zrbStartApp
    configs:
      start: sleep 10 && python -m http.server 8080
      ports: 8080
```

Let's try to modify your `index.zaruba.yaml` and invoke `zaruba please startServer`.

```
gofrendi@sanctuary [12:21:19] [~/playground/example]
-> % zaruba please startServer
💀 🔎 Job Starting...
         Elapsed Time: 1.8µs
         Current Time: 12:21:28
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/playground/example
💀    🚀 updateProjectLinks   🔗 12:21:28.719 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 12:21:28.719 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/playground/example
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/playground/example
💀    🔎 startServer          🍏 12:21:29.015 📜 Waiting for port '8080'
💀    🚀 startServer          🍏 12:21:39.186 Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
💀    🔎 startServer          🍏 12:21:40.089 📜 Port '8080' is ready
💀    🔎 startServer          🍏 12:21:40.089 🎉🎉🎉
💀    🔎 startServer          🍏 12:21:40.089 📜 Task 'startServer' is ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 11.7249222s
         Current Time: 12:21:40
         Active Process:
           * (PID=16854) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
```

This code is easier to write since you no longer need to write the checker's loop.

You might also notice that in this example, we don't have any `start` and `check` property. Instead, we have `extend` and `configs` property. We will learn about those properties later.

## Extending a task