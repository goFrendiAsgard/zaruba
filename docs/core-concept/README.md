[⬆️](../README.md)

# Core concept

## Zaruba scripts

Zaruba scripts are written in [YAML](https://yaml.org/) and [go template](https://pkg.go.dev/text/template). The scripts define how your tasks should be executed and how they depend to/extend other tasks/resources.

## Invoking task

When you invoke `zaruba please`, Zaruba will load the scripts from `index.zaruba.yaml` file in your current directory.
Please take note that any directory containing `index.zaruba.yaml` is called a `zaruba project`.

In order to run any tasks defined in your zaruba project, you should execute Zaruba from inside the project directory.

Several tasks are defined in `preloaded script`. That means the tasks can be executed from anywhere.

The simplest way to run your a Zaruba task is by invoking `zaruba please <task-name>`. Since we didn't create any zaruba project yet, let's try to execute `update` task by invoking `zaruba please update`.

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

`update` is defined in a preloaded script, thus you can invoke it from anywhere.

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

After the task has been defined, you can add it's path to `ZARUBA_SCRIPTS`. Take note that somehow `~` is not working because the symbol is parsed by shell (see [this issue](https://github.com/golang/go/issues/15827)). Thus, you need to use `${HOME}` instead.

```
gofrendi@sanctuary [16:28:33] [~]
-> % export ZARUBA_SCRIPTS="${ZARUBA_SCRIPTS}:${HOME}/playground/figlet/example.yaml"
```

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

## Configuration using environment variables

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

You will learn about each of those keywords in each subtopics.


> 💡 __NOTE:__ Zaruba scripts are meant to be edited, not created from scratch. In most cases, you will find generator for most of your use cases.

For now, I will show you how a zaruba scripts looks like, so that you can get some idea about it before jumping into the subtopics:

### index.yaml

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

### configs.yaml

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

You might also notice several strings inside double curly-braces. Those are go-template. We will discuss about it in the subtocpic.

### envs.yaml

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

### inputs.yaml

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

### tasks.yaml

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
💀    🚀 prepareProvoFastApi  🔧 18:26:25.112 Requirement already satisfied: certifi==2021.5.30 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 4)) (2021.5.30)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.113 Requirement already satisfied: charset-normalizer==2.0.6 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 5)) (2.0.6)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.115 Requirement already satisfied: click==8.0.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 6)) (8.0.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.118 Requirement already satisfied: cryptography==36.0.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 7)) (36.0.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.132 Requirement already satisfied: fastapi==0.68.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 8)) (0.68.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.173 Requirement already satisfied: greenlet==1.1.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 9)) (1.1.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.175 Requirement already satisfied: h11==0.12.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 10)) (0.12.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.177 Requirement already satisfied: idna==3.2 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 11)) (3.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.179 Requirement already satisfied: jsons==1.5.1 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 12)) (1.5.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.181 Requirement already satisfied: kafka-python==2.0.2 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 13)) (2.0.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.184 Requirement already satisfied: passlib==1.7.4 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 14)) (1.7.4)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.189 Requirement already satisfied: pika==1.2.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 15)) (1.2.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.193 Requirement already satisfied: pydantic==1.8.2 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 16)) (1.8.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.199 Requirement already satisfied: pytest==6.2.5 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 17)) (6.2.5)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.209 Requirement already satisfied: pytest-cov==3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 18)) (3.0.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.216 Requirement already satisfied: python-jose==3.3.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 19)) (3.3.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.221 Requirement already satisfied: python-multipart==0.0.5 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 20)) (0.0.5)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.223 Requirement already satisfied: requests==2.26.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 21)) (2.26.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.234 Requirement already satisfied: sqlalchemy==1.4.23 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 22)) (1.4.23)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.273 Requirement already satisfied: starlette==0.14.2 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 23)) (0.14.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.279 Requirement already satisfied: typing-extensions==3.10.0.2 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 24)) (3.10.0.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.281 Requirement already satisfied: typish==1.9.3 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 25)) (1.9.3)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.288 Requirement already satisfied: urllib3==1.26.6 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 26)) (1.26.6)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.296 Requirement already satisfied: uuid==1.30 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 27)) (1.30)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.297 Requirement already satisfied: uvicorn==0.15.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 28)) (0.15.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.308 Requirement already satisfied: six>=1.4.1 in ./venv/lib/python3.8/site-packages (from bcrypt==3.2.0->-r requirements.txt (line 3)) (1.16.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.309 Requirement already satisfied: cffi>=1.1 in ./venv/lib/python3.8/site-packages (from bcrypt==3.2.0->-r requirements.txt (line 3)) (1.15.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.311 Requirement already satisfied: iniconfig in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (1.1.1)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.312 Requirement already satisfied: attrs>=19.2.0 in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (21.4.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.335 Requirement already satisfied: py>=1.8.2 in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (1.11.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.337 Requirement already satisfied: pluggy<2.0,>=0.12 in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (1.0.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.341 Requirement already satisfied: toml in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (0.10.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.343 Requirement already satisfied: packaging in ./venv/lib/python3.8/site-packages (from pytest==6.2.5->-r requirements.txt (line 17)) (21.3)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.346 Requirement already satisfied: coverage[toml]>=5.2.1 in ./venv/lib/python3.8/site-packages (from pytest-cov==3.0.0->-r requirements.txt (line 18)) (6.2)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.349 Requirement already satisfied: pyasn1 in ./venv/lib/python3.8/site-packages (from python-jose==3.3.0->-r requirements.txt (line 19)) (0.4.8)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.35  Requirement already satisfied: rsa in ./venv/lib/python3.8/site-packages (from python-jose==3.3.0->-r requirements.txt (line 19)) (4.8)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.352 Requirement already satisfied: ecdsa!=0.15 in ./venv/lib/python3.8/site-packages (from python-jose==3.3.0->-r requirements.txt (line 19)) (0.17.0)
💀    🚀 prepareProvoFastApi  🔧 18:26:25.356 Requirement already satisfied: pycparser in ./venv/lib/python3.8/site-packages (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 3)) (2.21)
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
💀    🚀 startProvoFastApi    ⚡ 18:26:26.268 2022-01-01 18:26:26,268 INFO sqlalchemy.engine.Engine [raw sql] ()
💀    🚀 startProvoFastApi    ⚡ 18:26:26.269 2022-01-01 18:26:26,269 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startProvoFastApi    ⚡ 18:26:26.269 2022-01-01 18:26:26,269 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startProvoFastApi    ⚡ 18:26:26.27  2022-01-01 18:26:26,269 INFO sqlalchemy.engine.Engine PRAGMA main.table_info("user")
💀    🚀 startProvoFastApi    ⚡ 18:26:26.27  2022-01-01 18:26:26,270 INFO sqlalchemy.engine.Engine [raw sql] ()
💀    🚀 startProvoFastApi    ⚡ 18:26:26.27  2022-01-01 18:26:26,270 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startProvoFastApi    ⚡ 18:26:26.272 2022-01-01 18:26:26,272 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startProvoFastApi    ⚡ 18:26:26.274 2022-01-01 18:26:26,274 INFO sqlalchemy.engine.Engine SELECT user.id AS user_id, user.username AS user_username, user.email AS user_email, user.json_permissions AS user_json_permissions, user.active AS user_active, user.hashed_password AS user_hashed_password, user.full_name AS user_full_name, user.created_at AS user_created_at, user.updated_at AS user_updated_at
💀    🚀 startProvoFastApi    ⚡ 18:26:26.274 FROM user
💀    🚀 startProvoFastApi    ⚡ 18:26:26.274 WHERE user.username = ?
💀    🚀 startProvoFastApi    ⚡ 18:26:26.274  LIMIT ? OFFSET ?
💀    🚀 startProvoFastApi    ⚡ 18:26:26.274 2022-01-01 18:26:26,274 INFO sqlalchemy.engine.Engine [generated in 0.00013s] ('root', 1, 0)
💀    🚀 startProvoFastApi    ⚡ 18:26:26.275 2022-01-01 18:26:26,275 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startProvoFastApi    ⚡ 18:26:26.276 Register app shutdown handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.287 Handle HTTP routes for auth.User
💀    🚀 startProvoFastApi    ⚡ 18:26:26.287 Register auth route handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.287 Register auth event handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.287 Handle RPC for auth.User
💀    🚀 startProvoFastApi    ⚡ 18:26:26.287 Register auth RPC handler
💀    🚀 startProvoFastApi    ⚡ 18:26:26.293 Handle HTTP routes for library.Book
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

Nice, our fastAPI application is running. BTW, did I told you that the zaruba scripts and the fastAPI application are also generated by Zaruba?

![](./images/provoFastApi.png)

## Subtopics

* [Tasks](./tasks.md)
* [Configs](./configs.md)
* [Envs](./envs.md)
* [Inputs](inputs.md)
* [Includes](includes.md)
* Go templates
