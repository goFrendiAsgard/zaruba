<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🔨 Task](README.md)
# 🧬 Task Anatomy
<!--endTocHeader-->


A task contains properties. Here are some possible properties:

```yaml
tasks:

  taskName:
    icon: ✨                        # icon of your task
    description: task description
    extend: ''                      # parent task name, you can also use `extends` for many inheritance
    location: './some-directory'    # task execution location
    private: false                  # whether the task is private or not
    timeout: 5m
    dependencies: []                # task's upstreams
    inputs:                         # task's inputs
      -inputName
    start:                          # command to start simple-command/long running service
      - bash
      - '-c' 
      - python -m http.server 8080
    check:                          # command to check the completeness of a long-running process
      - bash, 
      - '-c' 
      - until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready" 
    configs:                        # task's configurations
      someConfig: someValue
    envs:                           # task's environments
      SOME_ENV:
        from: SOME_GLOBAL_ENV
        default: defaultValue
    configRef: configName           # shared project configs used by this task
    envRef: envName                 # shared project envs used by this task
    autoTerminate: false            # whether this task should be auto terminated or not
    syncEnv: true                   # whether you can sync task env or not
    syncEnvLocation: ''             # location of environment files. If not set, Zaruba will use `location` instead
    saveLog: true                   # wether to save log or not
```

Please note that some properties are __exclusive to each other__. That means you cannot use the following properties together:

* `extend` and `extends`
* `configRef` and `configRefs`
* `envRef` and `envRefs`

With that in mind, let's dive into each key.

# Icon

This is the icon representing a task. If the task's icon is not defined, Zaruba will generate one on runtime.

You can put anything as a task icon. It is preferable to use an icon with 2 monospaced-character widths.

For some inspiration, please visit [emojipedia](https://emojipedia.org/).

# Description

Description can be either single-line or multi-line. It should describe what a task does and how to configure/run it.

# Extend

Your task might be extending another task. For example, consider the following example:

```yaml
tasks:
  
  print:
    configs:
      str: anything
    start: [bash, -c, 'echo "${ZARUBA_CONFIG_STR}"']
  

  printHello:
    extend: print
    configs:
      str: hello
```

To see how this works, please visit [this documentation](./extend-task.md)

Please take note that you cannot use `extend` and `extends` in a single task.

# Extends

In some cases, you might need to extend many tasks at once. Generally, this is not a good idea, since it makes your tasks prone to [diamond-problem](https://www.makeuseof.com/what-is-diamond-problem-in-cpp/).

A better approach in this situation is by using [shared config](./task-configs/shared-configs.md) and [shared-envs](./task-configs/shared-envs.md).

Being said so, you can use `extends` for a quick workaround:

```yaml
tasks:

  firstParentTask:
    configs:
      foo: bar
  
  secondParentTask:
    configs:
      spam: egg
  
  childTask:
    extends:
      - firstParentTask
      - secondParentTask
    start: [bash, -c 'echo "${ZARUBA_CONFIG_FOO} ${ZARUBA_CONFIG_SPAM}"']
```

>  ⚠️ __WARNING:__ `extends` property is here for historical purpose and a quick workaround.

# Location

Location of your task relative to your Zaruba script location.

For example, consider this directory structure:

```
.
├── index.zaruba.yaml          # script entry point
├── zaruba-tasks
│   └── application
|       └── tasks.yaml         # Your task definition is here
└── application                # And you want your task to run here
```

Suppose you define a task at `./zaruba-tasks/application/tasks.yaml`. To run that task from the `./application` directory, you need to define its location as follows:

```yaml
# file: ./zaruba-tasks/application/tasks.yaml
tasks:

  startApplication:
    extend: zrbStartService
    location: ../../application         # task location relative to this script
    configs:
      start: python -m http.server 8080
      ports: 8080
```

# Private

When you set a task's private property to `true`, that task becomes a private task.

A private task is not accessible from the interactive mode.

To use a private task, you need to create another task that extends from it.

```yaml
tasks:

  privateTask: # this task is private
    private: true
    start: [bash, -c, echo hello]

  publicTask: # this task is not private
    private: false
    extend: privateTask
```

# Timeout

Some tasks might need a lot of time to get started. By default, Zaruba will kill any task that takes more than 10 minutes to be ready.

But you can make it shorter or longer:

```yaml
tasks:

  longTask: # zaruba will wait for 1 hour before consider kill this task
    timeout: 1h
    start: [bash, -c, 'slepp 3500 && echo done']
```

# Dependencies

Some tasks might have some dependencies. For example, you cannot start a Typescript application without invoking `tsc` first.

You can use `dependencies` property to define this behavior.

Let's look at the following example:

```yaml
tasks:

  installDependencies:
    extend: zrbRunShellScript
    configs:
      start: npm install
  
  compileTypescript:
    extend: zrbRunShellScript
    dependencies:
      - installDependencies
    configs:
      start: tsc
  
  startApplication:
    extend: zrbRunShellScript
    dependencies:
      - compileTypescript
    configs:
      start: npm start
```

For more information about task dependencies, please visit [this documentation](extending-task-dependencies.md)

# Inputs

Inputs keyword allows a task to use [project inputs](../project/project-inputs.md).

```yaml
inputs:

  name:
    default: world
    prompt: your name

tasks:

  sayHello:
    extend: zrbRunShellScript
    inputs:
      - name
    configs:
      start: echo "Hello $ZARUBA_INPUT_NAME"
```

# Start

`Start` is a low-level property that allows you to define what a task should do. Please look at the following example:

```yaml
tasks:

  startServer:
    start: [bash, -c, 'echo "hello world"']
```

In most cases, you don't need to set `start` property at all. Instead, you can make a task extending any of the following:

* [zrbRunShellScript](../../core-tasks/zrb-runShell-script.md)
* [zrbStartApp](../../core-tasks/zrb-start-app.md)
* [zrbStartDockerContainer](../../core-tasks/zrb-start-docker-container.md)
* [zrbStartDockerCompose](../../core-tasks/zrb-start-docker-compose.md)

The tasks we mentioned above already have `configs.start` property, so you can use them as follow:

```yaml
tasks:

  startServer:
    extend: zrbRunShellScript
    configs:
      start: echo "hello world"
```


# Check


`Check` is a low-level property that allows you to define [long-running process](./long-running-service.md) completeness. Please look at the following example:


```yaml
tasks:

  startServer:
    start: [bash, -c, 'python -m http.server 8080']
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"']
```

In most cases, you don't need to set `check` property at all. Instead, you can make a task extending any of the following:

* [zrbStartApp](../../core-tasks/zrb-start-app.md)
* [zrbStartDockerContainer](../../core-tasks/zrb-start-docker-container.md)
* [zrbStartDockerCompose](../../core-tasks/zrb-start-docker-container.md)


The tasks we mentioned above already have `configs.check` property, so you can use them as follow:

```yaml
tasks:

  startServer:
    extend: zrbStartApp
    configs:
      start: python -m http.server 8080
      check: until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"
```



# Configs

Task configurations. Please check [this document](./task-configs/README.md) for more information.

```yaml
tasks:

  sayHello:
    extend: zrbRunShellScript
    configs:
      name: 'world'
      start: echo "Hello ${ZARUBA_CONFIG_NAME}"
```

# Envs

Task environments. Please check [this document](./task-envs/README.md) for more information.

```yaml
tasks:

  sayHello:
    extend: zrbRunShellScript
    configs:
      start: echo "Hello ${NAME}"
    envs:
     NAME:
       from: HELLO_NAME 
       default: world
```

# ConfigRef

You can use `configRef` property to use [project config](../project/project-configs.md) in your task.

```yaml
configs:

  helloConfig:
    name: world

tasks:

  sayHello:
    extend: zrbRunShellScript
    configRef: helloConfig
    configs:
      start: echo "Hello ${ZARUBA_CONFIG_NAME}"
```

# ConfigRefs

You can use `configRefs` property to use many [project configs](../project/project-configs.md) in your task.


```yaml
configs:

  helloConfig:
    name: world
  
  otherConfig:
    key: value

tasks:

  sayHello:
    extend: zrbRunShellScript
    configRefs:
      - helloConfig
      - otherConfig
    configs:
      start: echo "Hello ${ZARUBA_CONFIG_NAME}"
```

# EnvRef

You can use `envRef` property to use [project env](../project/project-envs.md) in your task.

```yaml
envs:

  helloEnv:
    NAME:
      from: HELLO_NAME
      default: world

tasks:

  sayHello:
    extend: zrbRunShellScript
    envRef: helloEnv
    configs:
      start: echo 'Hello {{ .GetEnv "NAME" }}'
```


# EnvRefs

You can use `envRefs` property to use many [project envs](../project/project-envs.md) in your task.

```yaml
envs:

  helloEnv:
    NAME:
      from: HELLO_NAME
      default: world
  
  otherEnv:
    SOME_ENV:
      from:  SOME_OTHER_ENV
      default: default

tasks:

  sayHello:
    extend: zrbRunShellScript
    envRefs:
      - helloEnv
      - otherEnv
    configs:
      start: echo 'Hello {{ .GetEnv "NAME" }}'
```

# AutoTerminate

Whether a task should be auto terminated or not.

If `autoTerminate` is set to true then:

* For [simple command](./simple-command.md): It will quit immediately after completed.
* For [long-running service](./long-running-service.md): It will quit immediately after ready.

You can also force auto termination and define waiting time by invoking:

```bash
zaruba please <task-name> -t -w 10s
```

__Example:__

<!--startCode-->
```bash
# Start a webserver. After ready, wait for 2 seconds, and stop.
zaruba please serveHttp -t -w 2s
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.575µs
         Current Time: 11:40:58
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🌐 'serveHttp' service on /home/gofrendi/zaruba/docs
💀 🏁 Check 🌐 'serveHttp' readiness on /home/gofrendi/zaruba/docs
💀    🔎 serveHttp            🌐 🔎 Waiting for port '8080'
💀    🚀 serveHttp            🌐 Serving /home/gofrendi/zaruba/docs on HTTP port 8080
💀    🚀 serveHttp            🌐 You can open http://localhost:8080
💀    🔎 serveHttp            🌐 🔎 Port '8080' is ready
💀    🔎 serveHttp            🌐 🎉🎉🎉
💀    🔎 serveHttp            🌐 📜 Task 'serveHttp' is ready
💀 🎉 Successfully running 🌐 'serveHttp' readiness check
💀 🔎 Job Running...
         Elapsed Time: 223.939332ms
         Current Time: 11:40:58
         Active Process:
           * (PID=2174) 🌐 'serveHttp' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🌐 'serveHttp' service (PID=2174)
💀 🔥 🌐 'serveHttp' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 3.428390712s
         Current Time: 11:41:01
zaruba please serveHttp   -t -w 2s
```````
</details>
<!--endCode-->

# SyncEnv

When you set `syncEnv` property to true, Zaruba will be able to update your task environments. You need to set `syncEnvLocation` or `location` to define the environment files location.

To update your task environment, you can invoke:

 ```bash
 zaruba please syncEnv
 ```

# SyncEnvLocation

Directory location of task's environment file. Used along with `syncEnv` property.

If `syncEnvLocation` property is not set, Zaruba will use `location` property instead.

# SaveLog

Whether Zaruba should save the task log into the log file or not.

To see the task logs you can invoke: 

```
zaruba please showLog [taskName]
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->