# 💀 Zaruba 

```
      _____
     /     \        My name is 💀Zaruba.
    | () () |       I came to be when Garo came to be.
     \  ^  /        And I am forever with him.
      |||||

Zaruba is an agnostic scaffolding tool, service runner, as well as monorepo management tool.
Zaruba will help you build project faster.

Usage:
  zaruba <action> [...args] [flags]
  zaruba [command]

Available Commands:
  create            Create component.
  help              Help about any command
  init              Init project.
  install-template  Install template.
  organize          Organize project.
  pull              Pull project.
  push              Push project.
  remove-containers Remove all containers.
  run               Run project.
  stop-containers   Stop all containers.
  version           Zaruba's version.

Flags:
  -h, --help   help for zaruba

```

**What Zaruba is**

Zaruba is an agnostic 🔨scaffolding tool, ✈️service runner, as well as 🐙monorepo management tool. Zaruba will help you build and test your project faster.

> **Fun fact:** Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)

**Who Zaruba is for**

Do you have to run several tmux panels in order to test your app? Do you need to copy-pasting codes from one app to another one during development phase? Do you want to publish some parts of your monorepo as open-source projects?

If you answer "yes" for any question above, then Zaruba is for you❤️.

**How to use Zaruba**

* 👨‍💻 [Installing Zaruba](#install)
* 🔨 [Using Zaruba as scaffolding tool](#scaffolding)
* ✈️ [Using Zaruba as service runner](#service-runner)
* 🐙 [Using Zaruba as monorepo management tool](#monorepo-management)

# <a name="install"></a> 👨‍💻 Installing Zaruba

In order to install and use zaruba, you need to set couple of things:

* 💻WSL/Linux/Mac
* 🐙Git
* 🐳Docker
* 🐹Golang
* 🔽Curl or Wget

To install Zaruba using `curl`, you should invoke:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

To install Zaruba using `wget`, you should invoke:

```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

## ⚙️ Configuring Zaruba

> **Note:** You can safely skip this part unless something goes wrong with your installation

Zaruba injects several environment variables to your `.bashrc` and `.zshrc`. You can find Zaruba's environment definition at `~/.zaruba/zaruba.env`:

```bash
if [ -z ${HOME} ]
then
    HOME=~
fi
export PATH="$(go env GOPATH)/bin:${PATH}"
export ZARUBA_SHELL="/bin/bash"
export ZARUBA_SHELL_ARG="-c"
export ZARUBA_TEMPLATE_DIR="${HOME}/.zaruba/templates"
```

* `ZARUBA_TEMPLATE_DIR` is the location of your templates. You can install template by performing `zaruba install-template <git-url>`
* `ZARUBA_SHELL` define shell used by zaruba (for running command). It doesn't have to be the same as your default one. For example, you might use `zsh` as your default terminal and `bash` as zaruba's terminal.
* `ZARUBA_SHELL_ARG` default to `-c`


# <a name="scaffolding"></a> 🔨 Using Zaruba as scaffolding tool

Currently Zaruba comes with at least 4 pre-installed templates:

* 🌳project
* 🦎nodejs-service
* 🐹go-service
* 🐍python-service

Those templates are somehow opinionated. We take some structure from [NestJS](https://nestjs.com/), but we strip down all the magics. We also write an easy-to-read README in the template, so you can see how everything works and decide whether you can go with it or you prefer to build your own templates.

**🥗 Creating project**

First of all, let's build our first project using this command:

```sh
zaruba create project myproject # create a project named "myproject"
cd myproject
```

Our newly built project contains `zaruba.config.yaml`, `default_rmq/zaruba.config.yaml`. When you run the project using `zaruba run` (more about this later), the configuration will be cascaded on-the-fly. This mean that you can have nested configurations in the project.

**🥬 Creating service**

After successfully build your project, you might want to add a service or two. Let's try to make one based on `go-service` template:

```sh
# assuming you are on `myproject` directory
zaruba create go-service . myservice # create a go-service in current project named `myservice`
```

By convention, service templates usually take two arguments. The first one is project-location (in our case it is `.`), and the second one is service-name (in our case it is `myservice`).

Your newly built service can be found at `services/myservice`. It should contains `main.go`, a `Dockerfile` (won't be used for now), and couples of code.

**🍻 Run project**

To run the project, you can simply invoke:

```sh
# assuming you are on `myproject` directory
zaruba run

# you can also run the test by invoking:
zaruba run scenario:test
```

The service already has `readiness` URL, as well as `pub-sub` and `RPC` mechanism. Try to send several requests to `myservice`:

```sh
curl localhost:3011/readiness # check readiness.

curl localhost:3011/hello # check "hello" example.
curl localhost:3011/hello/Kouga

curl localhost:3011/hello-rpc/Kaoru # invoking myservice.helloRPC RPC call.

curl localhost:3011/hello-pub/Ryuga # publish myservice.helloEvent event.
curl localhost:3011/hello-pub/Rian
curl localhost:3011/hello-all # get every name catched by myservice.helloEvent subscriber.
```


## 📦 Installing custom template

In case of you need more templates, you can always add one to your arsenal. You can install custom template from github (or any other git hosting) by invoking:

```sh
zaruba install-template <repository-url> [costum-template-name]
```

For example, if you want to install `project-template`, you can perform:

```sh
zaruba install-template https://github.com/state-alchemists/zaruba-project-template project
```

## 🧬 Creating custom template

A template should at least consists of two executable files:

* `install-template.zaruba`: This file will be executed once after you install the template.
* `create-component.zaruba`: This file will be executed everytime you perform `zaruba create <template> [project-dir [args...]]`.

By default, zaruba come with several pre-defined templates. Feel free to explore `~/.zaruba/templates` to see them and to create your own. If you think your template is going to be useful for everyone else, please consider put publish them in github.

# <a name="service-runner"></a> ✈️ Using Zaruba as service runner

The simplest command to run your components is by running `zaruba run` inside your project directory. However, you can also use selector to run particular services/container/commands.

To understand about selector, let's look on this configuration:

```yaml
name: testRun
env: 
  RMQ_USER: root
  RMQ_PASS: toor
  RMQ_PORT: 5672
  RMQ_API_PORT: 15672
  ALPHA_HTTP_PORT: 3011

components: 

  rmq:
    symbol: 🐇 
    type: container
    image: rabbitmq:3-management
    env:
      RABBITMQ_DEFAULT_USER: ${RMQ_USER}
      RABBITMQ_DEFAULT_PASS: ${RMQ_PASS}
    ports: 
        ${RMQ_PORT}: 5672
        ${RMQ_API_PORT}: 15672
    readinessUrl: http://localhost:15672 
 
  alpha-test:
    type: command
    labels:
      scenario: test complete
    env:
      ALPHA_HTTP_PORT: 4011
      TEST_RMQ_CONNECTION_STRING: amqp://${RMQ_USER}:${RMQ_PASS}@${rmq}:${RMQ_PORT}/
    start: go test ./...
    dependencies:
      - rmq

  alpha:
    type: service
    labels:
      scenario: default complete
      domain: api-gateway
    env:
      ALPHA_HTTP_PORT: ${ALPHA_HTTP_PORT}
      DEFAULT_RMQ_CONNECTION_STRING: amqp://${RMQ_USER}:${RMQ_PASS}@${rmq}:${RMQ_PORT}/
    start: go build -o app && ./app
    readinessUrl: http://${alpha}:${ALPHA_HTTP_PORT}/readiness
    dependencies:
      - rmq
```

We have three components in the configuration:

* **rmq** is a rabbitmq container with no `labels`.
* **alpha-test** is a command with `scenario label`. It depends on `rmq` and it matches two scenarios: `scenario:test`, and `scenario:complete`. 
* **alpha** is a service with `scenario label` and `domain label`. This component depends on `rmq`. It also matches two scenarios and one domain: `scenario:default`, `scenario:complete` and `domain:api-gateway`.

There are some possibilities to run the components. Let's dive into them:

## 🙎‍♀️ Run by component's names

To run specific components, you can pass component's names as arguments. You can put as many component's names as necessary.

```sh
zaruba run rmq # Only run rmq
zaruba run alpha # Only run alpha after running all alpha's dependencies
zaruba run alpha alpha-test # Run alpha and alpha-test after running their dependencies
```

## 👨‍👩‍👧 Run by selectors

## 👨‍👩‍👧‍👧 Run with no parameter

```sh
zaruba run
```

Zaruba will run all components matching `scenario:default` selector.


# <a name="monorepo-management"></a> 🐙 Using Zaruba as monorepo management tool

つづく