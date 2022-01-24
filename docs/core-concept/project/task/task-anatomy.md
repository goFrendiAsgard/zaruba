[⬅️ Table of Content](../../../README.md)

# Task Anatomy

A task contains several keys and values. Hare are some possible keys/values for a task:

```yaml
tasks:

  taskName:
    icon: ✨                        # icon of your task
    description: task description
    extend: ''                      # other task name extended by this task. for multiple extend, use `extends` instead (but no, don't use it)
    location: './some-directory'    # directory location where your task should run on
    private: false                  # if true, the task is inteded to be extended instead of run directly
    timeout: 5m
    dependencies: []                # task's upstreams
    inputs:                         # task's inputs
      -inputName
    start: [bash, -c, 'python -m http.server 8080'] # command to start simple-command/long running process
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"'] # command to check readiness of long-running process
    configs:                        # task's configurations
      someConfig: someValue
    envs:                           # task's environments
      SOME_ENV:
        from: SOME_GLOBAL_ENV
        default: defaultValue
    configRef: configName           # shared project configs used by this task
    envRef: envName                 # shared project envs used by this task
    autoTerminate: false            # whether this task should be autoterminated or not
    syncEnv: true                   # whether the environments should be synchronized when running `zaruba please syncEnv` or not
    syncEnvLocation: ''             # location of environment file's directory. If not set, `location` will be used
    saveLog: true                   # wether to save log or not
```

Please note that some keys cannot be used simultaneously in a single task:

* `extend` and `extends`
* `configRef` and `configRefs`
* `envRef` and `envRefs`

With that in mind, let's dive into each keys.

# Icon

This is the icon representing a task. If task's icon is not define, Zaruba will automatically generate one on runtime.

You can put anything as task's icon. Preferably if it takes 2 character width in monospaced fonts.

For some interesing icon you can use, please visit [emojipedia](https://emojipedia.org/).

# Description

Description can be either single line or multi line. It should describe what a task do and how to configure/run it.

# Extend

Your task might extending other task. For example, consider the following example:

```yaml
tasks:
  
  print:
    configs:
      str: anything
    start: [bash, -c, 'echo {{ .GetConfig "str" }}']
  

  printHello:
    extend: print
    configs:
      str: hello
```

To see how this works, please visit [extend task documentation](./extend-task.md)

Please take note that you cannot use `extend` and `extends` in a single task.

# Extends

Ocassionally you might need to extend multiple tasks at once. Generally, this is not a good idea, since you will be exposed to [diamond problem](https://www.makeuseof.com/what-is-diamond-problem-in-cpp/).

A better approach in this situation is by using [shared config](./task-configs/shared-configs.md) and [shared-envs](./task-configs/shared-envs.md).

Being said so, you can use `extends` for quick workaround:

```yaml
tasks:

  firstParentTask:
    configs:
      foo: bar
  
  secondParentTask:
    configs:
      smap: egg
  
  childTask:
    extends:
      - firstParentTask
      - secondParentTask
    start: [bash, -c 'echo {{ .GetConfig "foo" }} {{ .GetConfig "spam" }}']
```

>  ⚠️ __WARNING__ Seriusly, don't use this unless you a very good reason. This mechanism was created before `configRef` and `envRef` created. This property is merely here for historical purpose, quick workaround (that need to be fixe later), and lure the 👿 demon.

# Location

Location of your task relative to your zaruba script location.

For example, consider this directory structure:

```
.
├── index.zaruba.yaml          # script entry point
├── zaruba-tasks
│   └── application
|       └── tasks.yaml         # Your task definition is here
└── application                # And you want your task to run here
```

Suppose your task is defined at `./zaruba-tasks/application/tasks.yaml` and you want to run the task inside `./application`, directory, then you need to define task's location as follow:

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

If a task's `private` property is true, you won't be able to run it in interactive mode.

A private task is intended to be extended by other tasks.

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

Some tasks might need a lot of time to get started. By default, Zaruba will kill stop running a task when it takse more than 10 minutes.

But you can make it shorter or longer:

```yaml
tasks:

  longTask: # zaruba will wait for 1 hour before consider kill this task
    timeout: 1h
    start: [bash, -c, 'slepp 3500 && echo done']
```

# Dependencies

Some tasks should not be executed unless it's dependencies are completed. For example, you cannot start a Typescript application without invoking `tsc`.

Let's see on this example:

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

Inputs keyword allows a task to use [project inputs](../project-inputs.md).

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
      start: echo 'Hello {{ .GetValue "name" }}'
```

# Start

`Start` is a low level property allows you to define what a task should do.

In most cases, you don't need to define `start` property:

* If your task run [simple command](./simple-command.md), you can create a task extending [zrbRunShellScript](../../../core-tasks/zrbRunShellScript.md)
* If your task run [long running process](./long-running-process.md), you can create a task extending [zrbStartApp](../../../core-tasks/zrbStartApp.md) or [zrbStartDockerContainer](../../../core-tasks/zrbStartDockerContainer.md)


```yaml
tasks:

  startServer:
    start: [bash, -c, 'echo "hello world"']
```


# Check


`Check` is a low level property allows you to define [long running process](./long-running-process.md) readiness.


In mose cases, you don't need to define `check` property. Instead, you can create a task extending [zrbStartApp](../../../core-tasks/zrbStartApp.md) or [zrbStartDockerContainer](../../../core-tasks/zrbStartDockerContainer.md)

```yaml
tasks:

  startServer:
    start: [bash, -c, 'python -m http.server 8080']
    check: [bash, -c, 'until nc -z localhost 8080; do sleep 2 && echo "not ready"; done && echo "ready"']
```

# Configs

Task configurations. Please check [task config document](./task-configs/README.md) for more information.

```yaml
tasks:

  sayHello:
    extend: zrbRunShellScript
    configs:
      name: 'world'
      start: echo 'Hello {{ .GetConfig "name" }}'
```

# Envs

Task environments. Please check [task env document](./task-envs/README.md) for more information.

```yaml
tasks:

  sayHello:
    extend: zrbRunShellScript
    configs:
      start: echo 'Hello {{ .GetEnv "NAME" }}'
    envs:
     NAME:
       from: HELLO_NAME 
       default: world
```

# ConfigRef

You can use `configRef` property to utilize [project config](../project-configs.md) in your task.

```yaml
configs:

  helloConfig:
    name: world

tasks:

  sayHello:
    extend: zrbRunShellScript
    configRef: helloConfig
    configs:
      start: echo 'Hello {{ .GetConfig "name" }}'
```

# ConfigRefs

You can use `configRefs` property to utilize multiple [project configs](../project-configs.md) in your task.


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
      start: echo 'Hello {{ .GetConfig "name" }}'
```

# EnvRef

You can use `envRef` property to utilize [project env](../project-envs.md) in your task.

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

You can use `envRefs` property to utilize multiple [project envs](../project-envs.md) in your task.

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

Whether a task should be autoterminated or not.

If `autoTerminate` is set to true then:

* If your task is a [simple command](./simple-command.md), it will quit immediately after completed.
* If your task is a [long running service](./long-running-process.md). It will quite immediately after ready.

You can also force autotermination and define waiting time by ivoking:

```bash
zaruba please <task-name> -t -w 10s
```

# SyncEnv

Whether task environment should be synced by `syncEnvLocation`/`location` or not.

If `syncEnv` is set to true, the task's environment will be synchronized whenever you invoke:

 ```bash
 zaruba please syncEnv
 ```

# SyncEnvLocation

Directory location of task's environment file. Used along with `syncEnv` property.

If `syncEnvLocation` is not set, then `location` property will be used instead.

# SaveLog

Whether task log should be saved or not. You can see the saved logs by invoking: 

```
zaruba please showLog [taskName]
```