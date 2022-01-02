[⬆️](../README.md)

# Core concept

## Zaruba scripts

Zaruba scripts are written in [YAML](https://yaml.org/) and [go template](https://pkg.go.dev/text/template). The scripts define how your tasks should be executed and how they depend to/extend other tasks/resources.

We will learn more about the anatomy of Zaruba script [here](README.md#anatomy-of-zaruba-scripts).

## Invoking task

To run any zaruba task, you can invoke:

```sh
zaruba please <taskName> [input=value...]
# or
zaruba please <taskName> -i
# or
zaruba please
```

When you invoke `zaruba please` with any parameter, Zaruba will load the scripts from `index.zaruba.yaml` in your current directory.

Any directory with `index.zaruba.yaml` inside is considered a `zaruba project`. To run any task defined in your `zaruba project` you should be in that directory.

If you are not currently in a `zaruba project` directory, you can still execute every tasks defined in `preloaded script`.

Since we didn't create any zaruba project yet, let's try to execute `update` task by invoking `zaruba please update`. This task is defined at `preloaded script`, thus you can invoke it from anywhere.

```
gofrendi@sanctuary [16:17:15] [~]
-> % zaruba please update
💀 🔎 Job Starting...
         Elapsed Time: 1.1µs
         Current Time: 16:17:19
💀 🏁 Run 🔄 'update' command on /home/gofrendi
💀    🚀 update               🔄 16:17:19.471 🔽 Pull zaruba
💀 🔥 🚀 update               🔄 16:17:19.548 Already on 'master'
💀    🚀 update               🔄 16:17:19.548 Your branch is up to date with 'origin/master'.
💀 🔥 🚀 update               🔄 16:17:22.364 From github.com:state-alchemists/zaruba
💀 🔥 🚀 update               🔄 16:17:22.364  * branch              master     -> FETCH_HEAD
💀    🚀 update               🔄 16:17:22.64  Already up to date.
💀    🚀 update               🔄 16:17:22.641 🚧 Compile zaruba
💀    🚀 update               🔄 16:17:22.935 🎉🎉🎉
💀    🚀 update               🔄 16:17:22.935 Zaruba ready!!!
💀    🚀 update               🔄 16:17:22.938 v0.9.0-alpha-1-130-gaced1b33
💀 🎉 Successfully running 🔄 'update' command
💀 🔎 Job Running...
         Elapsed Time: 3.7539664s
         Current Time: 16:17:23
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.955685s
         Current Time: 16:17:23
```

To see list of available tasks in `preloaded script`, you can visit [this list](../tasks/README.md)

## Preloaded scripts

By default, Zaruba always load preloaded scripts from `~/.zaruba/core.zaruba.yaml`. The tasks defined in that special files can be invoked from anywhere.

You can add more pre-loaded scripts by overriding `ZARUBA_SCRIPTS` variable.

Let's try this trick by creating a YAML file containing a task definition:

```
gofrendi@sanctuary [16:22:49] [~]
-> % mkdir -p ~/playground/figlet
gofrendi@sanctuary [16:24:31] [~]
-> % cat > ~/playground/figlet/example.yaml << EOF
heredoc> tasks:
heredoc>   sayHello:
heredoc>     start: [figlet, hello]
heredoc> EOF
```

> 💡 __TIPS:__ If you don't have `figlet` installed in your computer, you can simply replace it with `echo`. I.e: `start: [echo, hello]`.


> 💡 __PRO TIPS:__ Install `figlet`, `cowsay`, and `lolcat` to add fun to your terminal.

After the task has been defined, you can add it's path to `ZARUBA_SCRIPTS`. 

```
gofrendi@sanctuary [16:28:33] [~]
-> % export ZARUBA_SCRIPTS="${ZARUBA_SCRIPTS}:${HOME}/playground/figlet/example.yaml"
```

> ⚠️  Take note that somehow `~` is not working because the symbol is parsed by shell (see [this issue](https://github.com/golang/go/issues/15827)). Thus, you need to use `${HOME}` instead.

This changes allow you to invoke `sayHello` from anywhere, even if you are not in a `zaruba project`.

```
gofrendi@sanctuary [16:29:07] [~]
-> % zaruba please sayHello
💀 🔎 Job Starting...
         Elapsed Time: 1.7µs
         Current Time: 16:29:13
💀 🏁 Run 🍏 'sayHello' command on /home/gofrendi
💀    🚀 sayHello             🍏 16:29:13.612  _          _ _
💀    🚀 sayHello             🍏 16:29:13.612 | |__   ___| | | ___
💀    🚀 sayHello             🍏 16:29:13.613 | '_ \ / _ \ | |/ _ \
💀    🚀 sayHello             🍏 16:29:13.613 | | | |  __/ | | (_) |
💀    🚀 sayHello             🍏 16:29:13.613 |_| |_|\___|_|_|\___/
💀    🚀 sayHello             🍏 16:29:13.613
💀 🎉 Successfully running 🍏 'sayHello' command
💀 🔎 Job Running...
         Elapsed Time: 103.9409ms
         Current Time: 16:29:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 215.3865ms
         Current Time: 16:29:13
```

You can make this changes permanent by adding the environment variables to your `~/.bashrc` or `~/.zshrc` (depends on your shell).

## Configuring zaruba by using environment variables

There are several environment variables you can use to control Zaruba's behavior:

* `ZARUBA_HOME` Location of your Zaruba installation directory. Default to `${HOME}/.zaruba`.
* `ZARUBA_BIN` Location of your Zaruba executable binary. Default to `${HOME}/.zaruba/zaruba`.
* `ZARUBA_SHELL` The shell zaruba used to execute shell scripts. Default to `bash`.

## Anatomy of zaruba scripts

On it's top level, Zaruba scripts only contains few keywords:

* `includes`: Reference to other zaruba scripts you want to load along with the current one.
* `inputs`: Input parameters that you can use interactively.
* `tasks`: Task definitions.
* `configs`: Task configurations that can be shared among tasks.
* `envs`: Environment definitions that can be shared among tasks.

You will learn about each of those keywords in each [subtopic](README.md#subtopics).


> 💡 __NOTE:__ Zaruba scripts are meant to be edited, not created from scratch. In most cases, you will find generator for most of your use cases.

## Example of zaruba scripts

In case of you want to see how those keywords work in action and how a project typically structured, you can continue to read this section. Otherwise, you can skip this section entirely and jumps into the [subtopics](README.md#subtopics).

Let's say you have a zaruba project containing `index.zaruba.yaml`, `provoFastApi`, and `zaruba-tasks` directory.

```
.
├── index.zaruba.yaml
├── provoFastApi
│   ├── Dockerfile
│   ├── __pycache__
│   ├── auth
│   ├── blog
│   ├── database.db
│   ├── helpers
│   ├── library
│   ├── main.py
│   ├── repos
│   ├── requirements.txt
│   ├── schemas
│   ├── start.sh
│   ├── template.env
│   └── venv
└── zaruba-tasks
    └── provoFastApi
        ├── bash
        │   ├── migrate.sh
        │   ├── prepare.sh
        │   ├── start.sh
        │   └── test.sh
        ├── configs.yaml
        ├── envs.yaml
        ├── index.yaml
        ├── inputs.yaml
        ├── tasks.container.yaml
        └── tasks.yaml    
```

Let's dive into some of those files and directories

### index.zaruba.yaml

The `index.zaruba.yaml` contains a simple script:

```yaml
includes:
  - ./zaruba-tasks/provoFastApi/index.yaml
```

Typically, you will find `index.zaruba.yaml` includes other scripts or define some simple tasks. This allows you to make your project to be structured modularly.

### provoFastApi

This directory contains the code for your fastAPI service. In some cases, you can also put your service outside of your project directory. For example if you handle multiple repositories.

### zaruba-tasks

In most zaruba project, there is a directory named `zaruba-tasks`. This directory contains any domain-specific scripts thats going to be loaded in `index.zaruba.yaml`.

Inside `zaruba-tasks` directory, you might find several drectories containing scripts related to a very specific domain. In this case you can see `zaruba-tasks/provoFastApi` directory.

Now let's jumps into some files in this `zaruba-tasks/provoFatApi`.

> 💡 __NOTE:__ You might wonder why don't we put those scripts inside our outer `provoFastApi` folder along with the source code. There are 2 reasons for this. __1)__ First, the source code should be independent from any irelevant resources. People might want to use the source code without Zaruba's runner. __2)__ The second reason is because a single source code can be deployed into different services using feature flag. Zaruba tasks is about how to run those services. You might want different names for your services eventhough they share the same codebase.

### zaruba-tasks/provoFastApi/index.yaml

```yaml
includes:
  - ./configs.yaml
  - ./envs.yaml
  - ./inputs.yaml
  - ./tasks.yaml
  - ./tasks.container.yaml
```

This script is pretty self-explanatory. It loads scripts from other files: `configs.yaml`, `envs.yaml`, `inputs.yaml`, `tasks.yaml`, and `tasks.container.yaml`.

Once loaded, any resources in those files can interact to each other. This is a bit different compared to `include` keyword in `C` or `import` in `Python`.

Includes make your scripts more manageable by allowing you to separate your scripts into several parts based on your preference.

### zaruba-tasks/provoFastApi/configs.yaml

```yaml
configs:

  provoFastApi:
    runInLocal: '{{ .GetValue "runProvoFastApiInLocal" }}'
    ports: |
      {{ .GetEnv "APP_HTTP_PORT" }}

  provoFastApiContainer:
    useImagePrefix: true
    imageName: provo-fast-api
    containerName: provoFastApi
    localhost: host.docker.internal
    checkCommand: |
      echo "check provoFastApi"
    volumes: |


  startProvoFastApi:
    start: |
      . "{{ .GetProjectPath "zaruba-tasks/provoFastApi/bash/start.sh" }}"
    check: |
      echo "check provoFastApi"

  prepareProvoFastApi:
    start: |
      . "{{ .GetProjectPath "zaruba-tasks/provoFastApi/bash/prepare.sh" }}"

  testProvoFastApi:
    start: |
      . "{{ .GetProjectPath "zaruba-tasks/provoFastApi/bash/test.sh" }}"

  migrateProvoFastApi:
    start: |
      echo "migrate provoFastApi"
```

Some tasks might share configurations with each others. By define the configurations outside of your task, you can reduce redundancy.

In this script, you can find 6 configurations, namely: `provoFastApi`, `provoFastApiContainer`, `startProvoFastApi`, `prepareProvoFastApi`, `testProvoFastApi`, and `migrateProvoFastApi`.

To use those configurations in your task, you should y$use `configRef` or `configRefs` property. More about this later.

Each configuration contains simple map. The value has to be string but can be multi-line.

You might also notice several strings inside double curly-braces. Those are go-template. We will discuss about it in the [go template subtocpic](using-go-template.md).

### zaruba-tasks/provoFastApi/envs.yaml

```yaml
envs:

  provoFastApi:
    APP_ACCESS_TOKEN_ALGORITHM:
      default: HS256
      from: PROVO_FAST_API_APP_ACCESS_TOKEN_ALGORITHM
    APP_ACCESS_TOKEN_EXPIRE_MINUTES:
      default: "30"
      from: PROVO_FAST_API_APP_ACCESS_TOKEN_EXPIRE_MINUTES
    APP_ACCESS_TOKEN_SECRET_KEY:
      default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7
      from: PROVO_FAST_API_APP_ACCESS_TOKEN_SECRET_KEY
    APP_ENABLE_EVENT_HANDLER:
      default: "1"
      from: PROVO_FAST_API_APP_ENABLE_EVENT_HANDLER
    APP_ENABLE_ROUTE_HANDLER:
      default: "1"
      from: PROVO_FAST_API_APP_ENABLE_ROUTE_HANDLER
    APP_ENABLE_RPC_HANDLER:
      default: "1"
      from: PROVO_FAST_API_APP_ENABLE_RPC_HANDLER
```

Next you have environments definition. In this script, we just define a single environment named `provoFastApi`. Just like configurations, environments can also be shared among tasks.

Each environment is a map containing environment variable's name as it's key, and another map as it's value.

Every value contains 2 keys, `default` and `from`.

You might wonder, why don't we use simple map just like the configuration.

Imagine you run two applications (i.e: `alpa` and `beta`) that depend on the same environment variable. You want the variable for each application hold different values.

By having `from` directive, you can simply have these global environments:

```sh
ALPHA_HTTP_PORT=3000
BETA_HTTP_PORT=5000
```

and load them into two different environments with the same name:

```yaml
envs:

  alpha:
    HTTP_PORT:
      from: ALPHA_HTTP_PORT     
      default: 8080

  beta:
    HTTP_PORT:
      from: BETA_HTTP_PORT     
      default: 8081
```

Now, `alpha` and `beta` can run on different ports. If the global environments don't exist, the default values are going to be used (`8080` for `alpha`'s `HTTP_PORT`, and `8081` for `beta`'s `HTTP_PORT`).

Just like `configurations`, you can  use `envRef` and `envRefs` property to share environment among tasks.

### zaruba-tasks/provoFastApi/inputs.yaml

```yaml
inputs:

  runProvoFastApiInLocal:
    default: yes
    options:
      - yes
      - no
    allowCustom: false
    description: Whether to run provoFastApi locally
    prompt: Run provoFastApi locally?
```

Beside configurations and environments, you can also share inputs. among your tasks. If you run the task interactively, Zaruba will show you prompt dialog based on your tasks and inputs.

To use inputs in your tasks, you should use `inputs` property.

### zaruba-tasks/provoFastApi/tasks.yaml

```yaml
tasks:

  startProvoFastApi:
    icon: ⚡
    extend: zrbStartApp
    location: ../../provoFastApi
    inputs:
      - runProvoFastApiInLocal
    dependencies:
      - prepareProvoFastApi
    configRefs:
      - startProvoFastApi
      - provoFastApi
    envRef: provoFastApi

  prepareProvoFastApi:
    icon: 🔧
    extend: zrbRunShellScript
    location: ../../provoFastApi
    configRefs:
      - prepareProvoFastApi
      - provoFastApi
    envRef: provoFastApi

  testProvoFastApi:
    icon: ✅
    extend: zrbRunShellScript
    location: ../../provoFastApi
    dependencies:
      - prepareProvoFastApi
    configRefs:
      - testProvoFastApi
      - provoFastApi
    envRef: provoFastApi

  migrateProvoFastApi:
    icon: 🦆
    extend: zrbRunShellScript
    location: ../../provoFastApi
    dependencies:
      - prepareProvoFastApi
    configRefs:
      - migrateProvoFastApi
      - provoFastApi
    envRef: provoFastApi
```

This script show you how tasks really looks like. Let's take a look on `startProvoFastApi`.

First of all `startProvoFastApi` is extended from `zrbStartApp`. You can find more about `zrbStartApp` [here](../tasks/zrbStartApp.md), but for now let's say `zrbStartApp` is a special task that contains all configuration to run a long running process. Not only run the task, it will also make sure that all ports are accessible.

`startProvoFastApi` will run on `../../provoFastApi` directory (relative to the `tasks.yaml`). It has an input named `runProvoFastApiInLocal`. That means a prompt will be shown up when you try to run the task in interactive mode.

Beside extending `zrbStartApp`, `startProvoFastApi` also has a dependency to `prepareProvoFastApi`. That's mean that whenever you execute `startProvoFastApi`, Zaruba will automatically execute and wait for `prepareProvoFastApi` first.

`startProvoFastApi` uses two configurations, namely `startProvoFastApi` and `provoFastApi`. Those configurations was already defined in `configs.yaml`.

Finally, this task also uses `provoFastApi` environment that has been delcared in `envs.yaml`.

### Executing the script

We will discuss everything comprehensively in the subtopics.

But before we dwelve any further, let's see the task in action:

```
gofrendi@sanctuary [18:26:07] [~/zaruba/playground] [master *]
-> % zaruba please startProvoFastApi
💀 🔎 Job Starting...
         Elapsed Time: 32.8µs
         Current Time: 18:26:24
💀 🏁 Run 🔧 'prepareProvoFastApi' command on /home/gofrendi/zaruba/playground/provoFastApi
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/playground
💀    🚀 updateProjectLinks   🔗 18:26:24.605 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 18:26:24.605 Links updated
💀    🚀 prepareProvoFastApi  🔧 18:26:24.61  Activate venv
💀    🚀 prepareProvoFastApi  🔧 18:26:24.611 Install dependencies
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀    🚀 prepareProvoFastApi  🔧 18:26:25.102 Requirement already satisfied: aiofiles==0.7.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (0.7.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.104 Requirement already satisfied: asgiref==3.4.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.4.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.108 Requirement already satisfied: bcrypt==3.2.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 3)) (3.2.0)
...
💀    🚀 prepareProvoFastApi  🔧 18:26:25.358 Requirement already satisfied: pyparsing!=3.0.5,>=2.0.2 in ./venv/lib/python3.8/site-packages (from packaging->pytest==6.2.5->-r requirements.txt (line 17)) (3.0.6)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.361 Requirement already satisfied: tomli; extra == "toml" in ./venv/lib/python3.8/site-packages (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 18)) (2.0.0)
💀 🔥 🚀 prepareProvoFastApi  🔧 18:26:25.384 WARNING: You are using pip version 19.2.3, however version 21.3.1 is available.
💀 🔥 🚀 prepareProvoFastApi  🔧 18:26:25.384 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareProvoFastApi  🔧 18:26:25.408 Prepare
💀    🚀 prepareProvoFastApi  🔧 18:26:25.408 prepare command
💀    🚀 prepareProvoFastApi  🔧 18:26:25.408 Preparation complete
💀 🎉 Successfully running 🔧 'prepareProvoFastApi' command
💀 🏁 Run ⚡ 'startProvoFastApi' service on /home/gofrendi/zaruba/playground/provoFastApi
💀 🏁 Check ⚡ 'startProvoFastApi' readiness on /home/gofrendi/zaruba/playground/provoFastApi
💀    🔎 startProvoFastApi    ⚡ 18:26:25.727 📜 Waiting for port '3000'
💀    🚀 startProvoFastApi    ⚡ 18:26:25.727 Activate venv
💀    🚀 startProvoFastApi    ⚡ 18:26:25.727 Start
💀    🚀 startProvoFastApi    ⚡ 18:26:26.268 2022-01-01 18:26:26,268 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startProvoFastApi    ⚡ 18:26:26.268 2022-01-01 18:26:26,268 INFO sqlalchemy.engine.Engine PRAGMA main.table_info("book")
...
💀    🚀 startProvoFastApi    ⚡ 18:26:26.293 Register library route handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.293 Register library event handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.293 Handle RPC for library.Book
💀    🚀 startProvoFastApi    ⚡ 18:26:26.293 Register library RPC handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.294 Register blog route handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.294 Register blog event handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.294 Register blog RPC handler
💀 🔥 🚀 startProvoFastApi    ⚡ 18:26:26.294 INFO:     Started server process [7496]
💀 🔥 🚀 startProvoFastApi    ⚡ 18:26:26.295 INFO:     Waiting for application startup.
💀 🔥 🚀 startProvoFastApi    ⚡ 18:26:26.295 INFO:     Application startup complete.
💀 🔥 🚀 startProvoFastApi    ⚡ 18:26:26.295 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startProvoFastApi    ⚡ 18:26:26.737 📜 Port '3000' is ready
💀    🔎 startProvoFastApi    ⚡ 18:26:26.737 check provoFastApi
💀    🔎 startProvoFastApi    ⚡ 18:26:26.738 🎉🎉🎉
💀    🔎 startProvoFastApi    ⚡ 18:26:26.738 📜 Task 'startProvoFastApi' is ready
💀 🎉 Successfully running ⚡ 'startProvoFastApi' readiness check
💀 🔎 Job Running...
         Elapsed Time: 2.4560141s
         Current Time: 18:26:26
         Active Process:
           * (PID=7113) ⚡ 'startProvoFastApi' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
```

Nice, our fastAPI application is running. Note that since `startProvoFastApi` depends on `prepareProvoFastApi`, the later will always be executed first.

![](./images/provoFastApi.png)

> 🎉 __FUN FACT:__ Do you know that the zaruba scripts you see in this example, as well as the FastAPI application are generated by Zaruba? They are actually part of Zaruba's integration testing. To generate them, you can simply run `make test` from inside your `${ZARUBA_HOME}` directory.

# Subtopics

* [Tasks](tasks.md)
* [Configs](configs.md)
* [Envs](envs.md)
* [Inputs](inputs.md)
* [Includes](includes.md)
* [Using go template](using-go-template.md)
