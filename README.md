![zaruba-logo](arts/zaruba-250.png)

> ⚠️ Things might change pretty fast and pretty often before we reach `v.1.0.0`. Please open issue if you find any problem using Zaruba.

# 💀 Zaruba 

Zaruba is a [task](docs/core-concepts/task/README.md) runner and [CLI utility](docs/utilities/README.md). It helps you `write`, `generate`, and `run` your tasks with ease.

## ❓ Problem

Developing/debugging/deploying applications can be challenging. First, you need to run many tasks __in parallel__. Then, you also need to execute those tasks in a __particular order__.

Some tasks might __depend on each other__ or __share similar behavior__. 

For example:

* You cannot start a web application unless the database server is ready. This means that the web application is __depending__ on the database server.

* You might have several Typescript applications in your project. And to start those applications, you need to perform `npm install && tsc && npm start`. This means that your Typescript applications __share similar behavior__.

There should be a way to declare and run your tasks with ease.

## 💡 Solution

Creating __clear instructions/checklists__ might help in most cases. For example, suppose your tasks are serial. Then, you can turn them into an excellent [shell script](https://www.shellscript.sh/first.html), which is all you need.

But, if your workflow is more complicated, you need a __better automation tool__ like Zaruba.

Zaruba allows you to __simplify your workflow__ by letting you:

* Create configurable tasks (i,e., using `configs`, `inputs`, or `envs`).
* Define task dependencies (i,e., using `dependencies`).
* Re-use and share configurations/behaviors (i,e., using `extend`, `configRef`, or `envRef`).
* Run tasks in parallels.
* Generate new tasks.

There are some [built-in tasks](docs/core-tasks/README.md) to achieve those goals. You can run `zaruba please` to see the list of available tasks.

## 🔍 Example

Please see the [end-to-end tutorials](docs/use-cases/from-zero-to-cloud.md) to see how you can use Zaruba in real life.


# 👨‍💻 Installation

<details>
<summary><bold>TL;DR</bold></summary>

```bash
sudo apt-get install golang wget curl git
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
zaruba install ubuntuEssentials
zaruba install docker
zaruba install kubectl
zaruba install helm
zaruba install pulumi
```

Visit the [getting started section](#-getting-started).

</details>

## 📖 From Source

Installing from source is the best way to set up Zaruba for day-to-day use.

We don't plan to create [APT](https://en.wikipedia.org/wiki/APT_(software)) or platform-specific packages for Zaruba. If you are using windows, you need to install [WSL](https://docs.microsoft.com/en-us/windows/wsl/install) to get started.

### 🧅 Prerequisites

Before installing Zaruba from the source, you need to install some prerequisites:

* `go 1.13` or newer (To install `go`, you can visit its [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> __💡 HINT__ If you are using Ubuntu, you can install all prerequisites by invoking: `sudo apt-get install golang wget curl git`.

### 🥗 Installing From Source

To install Zaruba using __curl__, you can do the following:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

To install Zaruba using __wget__, you can do the following:

 ```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

## 🐳 Using Docker

Using docker is the quickest way to install Zaruba, especially if you need to use Zaruba in your CI/CD.

To create and run a Zaruba container on a __🐧Linux__ host, you can do the following:

```bash
docker run -d --name zaruba --network host \
  -v "$(pwd):/project" \
  -e "ZARUBA_HOST_DOCKER_INTERNAL=172.17.0.1" \
  -e "DOCKER_HOST=tcp://172.17.0.1:2375" \
  stalchmst/zaruba:latest
```

To create and run a Zaruba container on a __🪟 Windows__/__🍎 Mac__ host, you can do the following:

```bash
docker run -d --name zaruba \
  -p 8500-8700:8500-8700 \
  -v "$(pwd):/project" \
  stalchmst/zaruba:latest
```

For more information about Zaruba's docker image, please visit [docker hub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> __⚠️ NOTE__ There will be some limitations if you run Zaruba container in `docker-desktop`. For example, docker-desktop doesn't support host networking. Thus, you need to expose the ports by yourself. (e.g., `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

# 📜 Getting Started

To get started, you can:

* [📖 Browse the documentation](docs/README.md),
* [❇️ Follow end to end tutorials](docs/use-cases/from-zero-to-cloud.md), and
* [🧠 Learn the core concept](docs/core-concepts/README.md)

# ➕ Extra Prerequisites

Some tasks need `docker`, `kubectl`, `helm`, and `pulumi` installed. You can invoke the following command to install those extra prerequisites: 

```bash
zaruba install <extra-prerequisite>
```

To see whether you need to install extra prerequisites or not, you can use this guide:

* You need [docker](https://www.docker.com/get-started) to build, pull or push images. You also need docker if you want to run your applications as containers.
* You need [kubectl](https://kubernetes.io/docs/home/#learn-how-to-use-kubernetes) to access your Kubernetes cluster.
* You need [helm](https://helm.sh/) and [pulumi](https://www.pulumi.com/) to deploy your applications into a Kubernetes cluster.
* You need [tocer](https://github.com/state-alchemists/tocer) to scaffold Zaruba's documentation.
* You need [pyenv](https://github.com/pyenv/pyenv) to run many `Python` versions.
* You need [nvm](https://github.com/nvm-sh/nvm) to run many `Node.Js` versions.

To install all extra prerequisites, please perform:

```bash
zaruba install docker
zaruba install kubectl
zaruba install helm
zaruba install pulumi
zaruba install pyenv
zaruba install nvm
```

# 🐞 Bug Report, Feature Request, and Contribution

You can always open an [issue](https://github.com/state-alchemists/zaruba/issues) or a [pull request](https://github.com/state-alchemists/zaruba/pulls).

When opening a pull request, please write down:

* Zaruba version you used.
* Your expectation/goal.
* Things you have tried to achieve the goal.
* The result you get.

> __💡 HINT__ You can get the zaruba version invoking: `zaruba version`.

# ☑️ Testing

To perform the test, you need to have:

* docker desktop
* kubectl
* helm
* pulumi
* go 1.13
* make

Once you meet all the prerequisites, you can perform:

```bash
make test
```

# ☕ Donation

[Paypal](https://paypal.me/gofrendi?country.x=ID&locale.x=en_US)

# 🎉 Fun fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba) is a Madougu which supports bearers of the Garo Armor. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)

![Madou Ring Zaruba on Kouga's Hand](arts/madou-ring-zaruba.jpg)
