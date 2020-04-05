# 💀 Zaruba 

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him."

Zaruba is agnostic generator and service runner. Zaruba treat your project like a big monorepo application containing multiple micro-services. However, zaruba also allows you to publish the micro-services as individual repo, independent from each other.

To see why you probably need monorepo, please look at this [awesome slide](https://www.symfony-news.com/slide/details/a-monorepo-vs-manyrepos).

Despite of the advantages, managing monorepo without a proper tool might brings you a nightmare. Zaruba is specifically created to make monorepo management easier.

# 👨‍💻 Installation

To install zaruba in your local computer, you have to install `golang 1.13`. Next, you can install zaruba by using `curl`:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

or `wget`:

```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

If you are using windows, it is recommended to use WSL.

# 🥗 Use Case

Let's say you have two services communicating via [NATS messaging system](https://nats.io/).

The first service is named `Gateway`. It listen to client's HTTP request, send a message to NATS, and waiting for another message before sendir HTTP response to client.

The second service is named `Greeter`. It listen to NATS and send a message as response.

Both services depended on `greeting-lib`. As the library contains a small amount of code, we don't want to publish is as `npm package`.

We want to publish `Gateway` and `Greeter` on separate repo, so that we can have different CI/CD and access privilege. However, we want to develop the application as a single monorepo because it is going to be easier to maintain.

Finally, we want our engineers to run all services in their local computer, isolated from each others.

![](readme-resource/usecase.png)

Zaruba was created to solve those problems. Let's see how zaruba helps you

## Creating Remote Respository

First of all you need to create 3 repositories. One for monorepo, and two for our services. In this tutorial, I will use github.

![](readme-resource/create-remote-repo.png)

The repository addresses in this tutorial are as follow:

* Monorepo: https://github.com/goFrendiAsgard/zaruba-project
* Gateway Service: https://github.com/goFrendiAsgard/zaruba-gateway
* Greeter Service: https://github.com/goFrendiAsgard/zaruba-greeter

Please make sure you use the different addresses by make your own repos.

![](readme-resource/remote-repos.png)

## Initiating Project

To initiate the project, you need to invoke:

```sh
mkdir myproject
cd myproject
zaruba init
git remote add origin git@github.com:goFrendiAsgard/zaruba-project.git # Adjust this one with your repo address
```

## Coding

Before writing some codes into our project, let's start with making some directories:

```
.
├── libraries
│   └── greeting-lib
│       └── greeting.js
└── services
    ├── gateway
    │   ├── package.json  # going to be generated when you perform `npm init`
    │   └── start.js
    └── greeter
        ├── package.json  # going to be generated when you perform `npm init`
        └── start.js
```

You can create all necessary directories by performing:

```sh
mkdir -p libraries/greeting-lib
mkdir -p services/gateway
mkdir -p services/greeter
```

### Greeting-Lib

Greeting-lib is a shared library that will be used by our services. It only contains a single file `libraries/greeting-lib/greeting.js`:

```js
module.exports = {
    greet: (name) => name ? `Hello ${name}` : "",
    defaultGreeting: "Hello world"
}
```

### Gateway

We use Node.Js to write gateway service. In order to prepare the package you need to change your working-directory to `services/gateway` and perform:

```sh
npm init
npm install --save nats
```

Finally, you have to create `services/gateway/start.js`:

```js
const http = require("http");
const NATS = require("nats");
const greeting = require("./greeting-lib/greeting")

const nats = NATS.connect();

const port = process.env.port || 80;
const subscribeEvent = process.env.sendMessageEvent || "bar";
const publishEvent = process.env.getMessageEvent || "foo";

//create a server object:
console.log(`Serve HTTP on port ${port}`);
http.createServer(function (req, res) {
    nats.subscribe(subscribeEvent,
        (receivedMessage) => {
            console.log(`Receive request from url: ${req.url}`);
            receivedMessage = receivedMessage || greeting.defaultGreeting
            res.end(receivedMessage);
        },
        (error) => {
            console.error(error);
            res.end("Internal server error");
        }
    );
    nats.publish(publishEvent, req.url);
}).listen(port);
```

### Greeter

Just like we made `Gateway`, we will make `Greeter` with a similar manner. First, you need to invoke:

```sh
npm init
npm install --save nats
```

After initiating npm package, we can make `services/greeter/start.js`:

```js
const NATS = require("nats");
const greeting = require("./greeting-lib/greeting")

const nats = NATS.connect();

const subscribeEvent = process.env.getMessageEvent || "foo";
const publishEvent = process.env.sendMessageEvent || "bar";

console.log(`Listen to ${subscribeEvent} event`);
nats.subscribe(subscribeEvent, (receivedMessage) => {
    console.log(`Received a message: ${receivedMessage}`);
    receivedMessage = receivedMessage.split("/").join(" ").trim();
    const publishedMessage = greeting.greet(receivedMessage);
    console.log(`Publish to ${publishEvent} event: ${publishedMessage}`);
    nats.publish(publishEvent, publishedMessage);
});
```

## Defining Project Configuration

After defining the project, you should create `zaruba.config.yaml`. This configuration define how to run the service locally and how to 

```yaml
environments:

  general: # General environments are shared among services
    natsUrl: nats://nats.io:4222 # Connection to nats
    getMessageEvent: foo   # gateway publish this event while greeter listen to it
    sendMessageEvent: bar  # gateway listen to this event while greeter send message to it

  services: # Service environments are localized to particular services
    gateway:
      port: 3000


components:

  greeting-lib: # greeting-lib is a shared library, used by gateway and greeter
    type: library
    location: "./libraries/greeting-lib"

  gateway: # gateway service
    type: service
    origin: git@github.com:goFrendiAsgard/zaruba-gateway.git # Adjust this one with your repo address
    branch: master
    location: "./services/gateway"
    start: npm install && node start

  greeter: # greeter service
    type: service
    origin: git@github.com:goFrendiAsgard/zaruba-greeter.git # Adjust this one with your repo address
    branch: master
    location: "./services/greeter"
    start: npm install && node start

  nats: # nats docker container for running the application locally
    type: container
    run: docker run --name nats -p 4222:4222 -p 6222:6222 -p 8222:8222 -d nats
    containerName: nats

links:
  ./libraries/greeting-lib: # greeting-lib should be copied to greeter and gateway
    - ./services/greeter/greeting-lib
    - ./services/gateway/greeting-lib
  
executions: # Execution order
  - nats
  - greeter
  - gateway
```

Once you have create the configuration, you need to re-init the project in order to register the subtrees

```sh
zaruba init
```
![](readme-resource/zaruba-init.png)

## Run The Project

Everything has been set. Now let's try to run the project in our local computer. Please note that you also need docker in order to install and run NATS.

To run the project, you should perform:

```sh
zaruba run
```

![](readme-resource/zaruba-run.png)

As you see, zaruba put the services logs in a single screen. This help us to debug the project as a whole.

Just to make sure, you can try send some HTTP request to the gateway.

## Publishing Changes to Service Repo

After making sure everything behave as expected, you can push your project by using:

```sh
zaruba push
```
![](readme-resource/zaruba-push.png)

Now your mono-repo, as well as your services repo are updated.

![](readme-resource/sub-repo-updated.png)

__Note:__ If you don't want to publish your work to services repos, you can invoke `git push -u origin HEAD` instead.

## Fetching Changes from Service Repo

As you publish your service-repo, someone might contribute to it.

![](readme-resource/change-on-sub-repo.png)

It is important to be able to retrieve the changes into the monorepo. In this case, you can perform:

```sh
zaruba pull
```

![](readme-resource/zaruba-pull.png)

# ✍🏻 Available Actions

```
Usage:
  zaruba <action> [...args] [flags]
  zaruba [command]

Available Commands:
  create           Create Component
  help             Help about any command
  init             Init a project
  install-template Install template
  organize         Organize a project
  pull             Pull from subtrees
  push             Push from subtrees
  run              Run project
  version          Print the version number of Zaruba

Flags:
  -h, --help   help for zaruba

Use "zaruba [command] --help" for more information about a command.
```

# 👴 Concepts

## Template

Templates are component's blueprint. It can be a node.js package, python program, or even bunch of shell scripts.

A template should contains at least two files:

* `install-template.zaruba`: Any executable script, containing set of commands to be executed after user install the template. You might find some `npm init` or `pip install` here.
* `create-component.zaruba`: Any executable script, containing set of commands to be executed when user creates new component based on current template. `create-component` should at least a single argument containing project directory.

To install a template, you can perform:

```sh
zaruba install-template https://github.com/someUser/someTemplate.git
```

## Project

Project is a directory containing set of components. A project might also be a component on it's own.

## Component

Component can be anything from a project, a shared library, or a single service.

A component is usually based on specific template, but user can also create their own components from scratch. Also, a component should contains at least a single file:

* `link`: Any executable script, containing set of commands to be executed when user perform `zaruba organize` or `zaruba watch`.

Optionally, a component can also has `organize-project.zaruba` or any other shell script for custom command.

To create a new component, you can perform:

```sh
zaruba create someTemplate
```

## pre-action

pre-action is will be executed before an action is executed. You can make pre-action by simply create an executable file with `pre` prefix.

## post-action

post-action is will be executed after an action is executed. You can make post-action by simply create an executable file with `post` prefix.

# Commmands

## install-template

```sh
zaruba install-template <template-git-url> [folder-name]
```

This one basically run `git clone <template-gir-url>` and executing `install-template`.

While running `install-template`, current working directory is set to `[folder-name]`. However, if `[folder-name]` is not specified, zaruba will use `<template-git-url>`'s repository name as `[folder-name]`.

## create

```sh
zaruba create <template> [project-dir [...args]]
```

This will run template's `create <project-dir> [...args]`. Typically, it should create new component based on `<template>`. It is assumed that current working directory is pointing to `<template>`.

## organize

```sh
zaruba organize [project-dir [...args]]
```

This command will do the following actions:

* Copy sources to their respective destinations.
* Recursively look for and run `organize-project.zaruba` in every sub-directory of `<project-dir>` and execute it. This command support pre-action (`pre-organize-project.zaruba`) and post-action (`post-organize-project.zaruba`).

# 👩🏻‍🏫 Configuration

## Environment Variable

* `ZARUBA_TEMPLATE_DIR`: Zaruba's template directory, default to `<zaruba-parent-dir>/templates`
* `ZARUBA_SHELL`: Default to `/bin/bash`
* `ZARUBA_SHELL_ARG`: Default to `-c`

# 🙇🏽 Testing

Copy `template.env` to `.env` and invoke:

```sh
source .env
make test # or make test-verbose
```

# Clone (Only for State-alchemists crews)

```
git clone git@github.com:state-alchemists/zaruba.git
cd zaruba
make install
zaruba init
```