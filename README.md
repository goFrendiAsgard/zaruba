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
* 🤔 [FAQ](#faq)

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
export ZARUBA_DOCKER_HOST="0.0.0.0"
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

Our newly built project contains `zaruba.yaml`, `containers/rmq.zaruba.yaml`. When you run the project using `zaruba run` (more about this later), the configuration will be cascaded on-the-fly. This mean that you can have nested configurations in the project.

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
      scenario: test complete       # alpha-test match `scenario:test`, as well as `scenario:complete`
    env:
      ALPHA_HTTP_PORT: 4011
      TEST_RMQ_CONNECTION_STRING: amqp://${RMQ_USER}:${RMQ_PASS}@${rmq}:${RMQ_PORT}/
    start: go test ./...
    dependencies:
      - rmq

  alpha:
    type: service
    labels:
      scenario: default complete    # alpha match `scenario:default`, as well as `scenario:complete`
      domain: api-gateway           # as well as `domain:api-gateway`
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

You can run components by passing their names as arguments:

```sh
zaruba run rmq # Only run `rmq`
zaruba run alpha # Only run `alpha` and it's dependencies (i.e: rmq, alpha)
zaruba run alpha alpha-test # Run `alpha`, `alpha-test`, and their dependencies (i.e: rmq, alpha, alpha-test)
```

## 👨‍👩‍👧 Run by selectors

You can also run components by using selectors. You can use the following format as valid selector:

* `<component-name>`
* `<label-key>:<value>`

Zaruba will looks for any component matching any of the selector (i.e: `or` logic). 

```sh
zaruba run domain:api-gateway # Run any component match `domain:api-gateway` and their dependencies (i.e: rmq, alpha)
zaruba run domain:api-gateway scenario:test # Run any component match `domain:api-gateway` or `scenario:test` and their dependencies (i.e: rmq, alpha, alpha-test)
zaruba run alpha scenario:test # Run `alpha` or any component match `scenario:test` and their dependencies (i.e: rmq, alpha, alpha-test)
zaruba run scenario:complete # Run any component match `scenario:complete` and their dependencies (i.e: rmq, alpha, alpha-test)
```

> **Note:** Feel free to add as many labels as needed in order to help you run the components easily.

## 👨‍👩‍👧‍👧 Run with no parameter

When no argument given, Zaruba will try to use `scenario:default` selector. If there is no component matching `scenario:default`, zaruba will run all components based on their dependencies.

```sh
zaruba run
```

Zaruba will run all components matching `scenario:default` selector.


# <a name="monorepo-management"></a> 🐙 Using Zaruba as monorepo management tool

Zaruba is a powerful monorepo management tool. It can automatically copy `shared-libraries`, as well as `push to` / `pull from` multiple repositories at once. Thanks to `git subtree` to make this possible.

Please have a look at the following configurations:

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
 
  alpha:
    type: service
    labels:
      scenario: default
    origin: "git@github.com:state-alchemists/alpha-service.git" # When you perform `zaruba init`, alpha will be registered as new subtree. This value will be used as alpha's origin.
    env:
      ALPHA_HTTP_PORT: ${ALPHA_HTTP_PORT}
      DEFAULT_RMQ_CONNECTION_STRING: amqp://${RMQ_USER}:${RMQ_PASS}@${rmq}:${RMQ_PORT}/
    start: go build -o app && ./app
    readinessUrl: http://${alpha}:${ALPHA_HTTP_PORT}/readiness
    dependencies:
      - rmq

links:
  ../../libraries/go/transport: # When you perform `zaruba organize|init|pull|push|run`, this directory will be copied to `./transport`
    - ./transport
  ../../libraries/go/core: # When you perform `zaruba organize|init|pull|push|run`, this directory will be copied to `./core`
    - ./core
 
``` 

## 📁 Organize

Zaruba automatically organizes your project whenever you perform `zaruba init`, `zaruba push`, `zaruba pull`, or `zaruba run`. However, if you need to organize the project manually, you can perform this command:

```sh
zaruba organize
```

Organizing a project is basically copy-pasting `link sources` to their destinations.

## ☀️ Init

You need to perform `zaruba init` in order to register repo's subtree. You also need to run this comment whenever you clone a zaruba-based repository from your git provider (e.g: github/gitlab/bitbucket) or add `origin` configuration for your components.

```sh
zaruba init
```

## 🔼 Push

Like `git push`, this will push your project to your git provider (including all subtree components).

```sh
zaruba push
```

## 🔽 Pull

Like `git pull`, this will push your project from your git provider (including all subtree components).

```sh
zaruba pull
```

# <a name="faq"></a> 🤔 FAQ

## Why a 💀skull, and why you name this "Zaruba"?

Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is is the Madou Ring for Golden Knight Garo's duties as a Makai Knight [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba). Eventhough Garo need to rely on his own power and expertise (like every developer in this world), a little help from a good tool can make everything easier.

## How is Zaruba different from 🐳docker-compose?

Docker-compose is a great tool to compose your containers. However, in most of the time, you don't want to build images and create containers just to make sure that your application works. Zaruba also allows you to run command for preparing your development environments or running the test.

## Can Zaruba create 🐳docker-compose file for me?

Depends on the template. Right now we don't have one, but we plan to make it come true (along with PM2). Stay tune.

## How is Zaruba different from 💂‍♂️yeoman?

Yeoman is based on Javascript, Zaruba template can be written in any language. You can use Yeoman and Zaruba together in case of you need to.

## How is Zaruba different from 🧞‍♂️PM2?

Zaruba is focusing in development while PM2 is focusing in production. For development, you need to run and kill everything at once as well as seeing all the log chronologically in a single place.

## Who use 💀Zaruba?

[This guy](https://twitter.com/gofrendiasgard) use Zaruba to run 5 micro-services (written in golang, python, and nodejs) and see all the logs in a single panel. The services was written long before Zaruba exists. Now he don't have any reason to not run and test them as often as possible.

If you think Zaruba is useful and you want your name listed here, please shout to [this guy](https://twitter.com/gofrendiasgard) on twitter.

## I found a 🐞bug, I have feature request, I want to contribute, How should I start?

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).
