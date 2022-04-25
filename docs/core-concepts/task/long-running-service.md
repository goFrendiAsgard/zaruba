<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🔨 Task](README.md)
# 🍹 Long Running Service
<!--endTocHeader-->

Long running services are considered completed once they are ready. Once completed, a long running service might run in indefinitely until it is killed.

Web servers and database servers are example of long running services.

__Example:__

Now let's try running a static web service by invoking `python -m http.server 8080`.

```bash
python -m http.server 8080
```

<details>
<summary>Output</summary>

```````
Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
```````
</details>

Once ready, the http server will wait for any incoming request at port, giving response, and wait again until you press ctrl + C.


# Process Readiness

You can make sure a service is ready by giving it a request and observe its response. In our case, you can verify that the service is ready by openning a browser and visit `http://localhost:8080`.

Making sure that a service is ready can be tricky since `running` doesn't mean `ready`.

Orchestration system like kubernetes overcome this problem by periodically sending request to your services using `liveness` and `readiness` [probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

# Start Process with Readiness Checker

In the previous section you have see that handling a service and making sure it is already running can be a bit challenging.

Under the hood, Zaruba make sure that your service is ready by running two commands simultaneously. The first command is responsible to run the service, while the other one is responsible to check its readiness. 

Let's see how this work on the low level.

First of all, you will need two terminals in the same computer. You can also use tmux/screen if you are familiar with those tools.

## Starter

In your first terminal, you can spawn this command `sleep 10 && python -m http.server 8080`: 

```
❯ sleep 10 && python -m http.server 8080
Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
```

This command ask the computer to wait for 10 seconds before starting the web server. We use this to emulate real world cases. Some server might even take more than a minute to be ready. Also, those 10 seconds give you enough time to open the second terminal and invoke the service checker (if you are really that quick).

## Checker

Our service checker contains a single loop to check whether `localhost:8080` is up and serving. In order to start the checker, you can invoke this in your second terminal `until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"`:

```
❯ until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"
not ready
not ready
not ready
not ready
ready
```

Great, now you can make sure that your service is really `ready` before deal with it any further.

> 💡 __TIPS:__  if you find the service is already started before you are able to start the checker, please feel free to change the sleep duration of the server (e.g: `sleep 30 && python -m http.server 8080`, will make the computer wait for 30 seconds before starting the server)

To see how our starter and checker works, let's take a look on this diagram:

![](images/starter-and-checker.png)


# Starting Long Running Service with Zaruba

Unless you are a [starcraft](https://starcraft2.com/en-us/) pro player, probably running multiple terminals and tmux panels is not a very good idea.

You might also want to run the server in the background or make a docker container for this simple use case. But let's not do that.

We will use Zaruba instead.

## Lower Level Approach

First you declare this script in your `index.zaruba.yaml`

```yaml
tasks:

  startServer:
    start: [bash, -c, 'sleep 10 && python -m http.server 8080']
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"']
```

then, you can invoke `zaruba please startServer`.


__Example:__

<!--startCode-->
```bash
cd examples/core-concepts/task/long-running-service/low-level
zaruba please startServer -t -w 1s
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.236µs
         Current Time: 08:25:50
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/zaruba/docs/examples/core-concepts/task/long-running-service/low-level
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/zaruba/docs/examples/core-concepts/task/long-running-service/low-level
💀    🔎 startServer          🍏 08:25:52.057 not ready
💀    🔎 startServer          🍏 08:25:54.059 not ready
💀    🔎 startServer          🍏 08:25:56.061 not ready
💀    🔎 startServer          🍏 08:25:58.063 not ready
💀    🔎 startServer          🍏 08:26:00.065 not ready
💀    🔎 startServer          🍏 08:26:02.067 not ready
💀    🔎 startServer          🍏 08:26:02.068 ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 12.115832811s
         Current Time: 08:26:02
         Active Process:
           * (PID=12017) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🍏 'startServer' service (PID=12017)
💀    🚀 startServer          🍏 08:26:03.673 Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
💀    🚀 startServer          🍏 08:26:03.673 
💀    🚀 startServer          🍏 08:26:03.673 Keyboard interrupt received, exiting.
💀 🔎 Job Ended...
         Elapsed Time: 14.2198797s
         Current Time: 08:26:04
💀 🔥 🍏 'startServer' service exited without any error message
zaruba please startServer   -t -w 1s
```````
</details>
<!--endCode-->


Good. This is what `check` property actually is for. It tells Zaruba how to check your service readiness. 

Any task with `start` and `check` property will be considered as `long running service`, while every tasks without `check` property are considered as `simple command`.

Please also take note that sometime a task might have `check` property eventhough it is not explicitly written. This is especially true if you [extend/inherit](extend-task.md) your task from another task.
    
> 💡 __TIPS:__  You might wonder why the server log doesn't show up unless you terminate it with `ctrl + c`. This is happened because of python buffering mechanism. To turn off this feature, you can set `PYTHONUNBUFFERED` to `1`. (i.e: by using this as start command, `start: [bash, -c, 'sleep 10 && export PYTHONUNBUFFERED=1 && python -m http.server 8080']`)


## Higher Level Approach

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

__Example:__

<!--startCode-->
```bash
cd examples/core-concepts/task/long-running-service/high-level
zaruba please startServer -t -w 1s
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.732µs
         Current Time: 08:26:04
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/core-concepts/task/long-running-service/high-level
💀    🚀 updateProjectLinks   🔗 08:26:04.495 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:26:04.496 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🍏 'startServer' service on /home/gofrendi/zaruba/docs/examples/core-concepts/task/long-running-service/high-level
💀 🏁 Check 🍏 'startServer' readiness on /home/gofrendi/zaruba/docs/examples/core-concepts/task/long-running-service/high-level
💀    🔎 startServer          🍏 08:26:04.6   🔎 Waiting for port '8080'
💀    🚀 startServer          🍏 08:26:14.699 Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
💀    🔎 startServer          🍏 08:26:15.624 🔎 Port '8080' is ready
💀    🔎 startServer          🍏 08:26:15.624 🎉🎉🎉
💀    🔎 startServer          🍏 08:26:15.624 📜 Task 'startServer' is ready
💀 🎉 Successfully running 🍏 'startServer' readiness check
💀 🔎 Job Running...
         Elapsed Time: 11.233390359s
         Current Time: 08:26:15
         Active Process:
           * (PID=12093) 🍏 'startServer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🍏 'startServer' service (PID=12093)
💀    🚀 startServer          🍏 08:26:17.227 
💀    🚀 startServer          🍏 08:26:17.227 Keyboard interrupt received, exiting.
💀    🚀 startServer          🍏 08:26:17.234 🎉🎉🎉
💀    🚀 startServer          🍏 08:26:17.234 📜 Task 'startServer' is started
💀 🔎 Job Ended...
         Elapsed Time: 13.33740883s
         Current Time: 08:26:17
💀 🔥 🍏 'startServer' service exited without any error message
zaruba please startServer   -t -w 1s
```````
</details>
<!--endCode-->


This code is easier to write since you no longer need to write the checker's loop.

You might also notice that in this example, we don't have any `start` and `check` property. Instead, we have [extend](./extend-task.md) and [configs](./task-configs/README.md) property.

Here are some of the tasks you can extend when you want to start long running service:

* [zrbStartApp](../../core-tasks/zrb-start-app.md): Lowest level, general use case
* [zrbStartDockerContainer](../../core-tasks/zrb-start-docker-container.md): Start a docker container

<!--startTocSubTopic-->
<!--endTocSubTopic-->
