![zaruba-logo](arts/zaruba-250.png)

> ⚠️ Things might change pretty fast and pretty often before we reach `v.1.0.0`. Please open issue if you find any problem using Zaruba.

# 💀 Zaruba 

Zaruba is a [task](docs/core-concepts/task/README.md) runner and [CLI utility](docs/utilities/README.md). It helps you to `write`, `generate`, and `run` your tasks easily.

## ❓ Problem

Developing/debugging/deploying applications can be challenging. You constantly need to run many tasks __in parallel__. Some of those tasks also need to be executed in a __particular order__.

Some tasks might __depend on each__ or __share similar behavior__. 

For example:

* You cannot start a web application unless the database server is ready. This means that the web application is __depending__ on the database server.

* You might have several Typescript applications in your project. And to start those applications, you need to perform `npm install && tsc && npm start`. This means that your Typescript applications __share similar behavior__.

There should be a way to declare and run your tasks accurately and easily.

## 💡 Solution

Creating __clear instructions/checklists__ might help in most cases. Assuming your tasks are sequential, you can turn it into a nice [shell script](https://www.shellscript.sh/first.html), and that's all you need.

But, if your workflow is more complicated, then you need a __better automation tool__ like Zaruba.

Zaruba allows you to __simplify your workflow__ by letting you:

* Create configurable tasks (i.e: using `configs`, `inputs`, or `envs`).
* Define task dependencies (i.e: using `dependencies`).
* Re-use and share configurations/behaviors (i.e: using `extend`, `configRef`, or `envRef`).
* Run tasks in parallels.
* Generate new tasks.

There are some [built-in tasks](docs/core-tasks/README.md) to achieve those goals. You can run `zaruba please` to see the list of available tasks.

## 🔍 Example

Please see the [end-to-end tutorial](docs/use-cases/from-zero-to-cloud.md) to see how you can use Zaruba in real life.


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

We don't have any plan to create `apt` or platform-specific packages for Zaruba. If you are using windows, you need to install `wsl` to get started.

### 🧅 Prerequisites

Before installing Zaruba from the source, you need to install some prerequisites software:

* `go 1.13` or newer (To install `go` quickly you can visit its [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> __💡HINT__ If you are using Ubuntu, you can invoke `sudo apt-get install golang wget curl git` to install all prerequisites.

### 🥗 Installing From Source

To install Zaruba using __curl__, you can do:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

To install Zaruba using __wget__, you can do:

 ```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

## 🐳 Using Docker

Using docker is the quickest way to set up Zaruba, especially if you need to use Zaruba in your CI/CD.

To create and run Zaruba container on __🐧linux__ host, you can do:

```bash
docker run -d --name zaruba --network host \
  -v "$(pwd):/project" \
  -e "ZARUBA_HOST_DOCKER_INTERNAL=172.17.0.1" \
  -e "DOCKER_HOST=tcp://172.17.0.1:2375" \
  stalchmst/zaruba:latest
```

To create and run Zaruba container on __🪟 windows__/__🍎 mac__ host, you can do:

```bash
docker run -d --name zaruba \
  -p 8500-8700:8500-8700 \
  -v "$(pwd):/project" \
  stalchmst/zaruba:latest
```

For more information about Zaruba's docker image, please visit [dockerhub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> __⚠️ NOTE__ There will be some limitations if you run Zaruba container in `docker-desktop`. For example, docker-desktop doesn't support host networking. Thus you need to expose the ports by yourself. (e.g: `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

# 📜 Getting Started

To get started, you can:

* [📖 Browse the documentation](docs/README.md),
* [❇️ Follow end to end tutorial](docs/use-cases/from-zero-to-cloud.md), and
* [🧠 Learn the core concept](docs/core-concepts/README.md)

# ➕ Extra Prerequisites

Some tasks need `docker`, `kubectl`, `helm`, and `pulumi` installed. To install those extra prerequisites, you can simply invoke `zaruba install <extra-prerequisite>`.

To see whether you need to install extra prerequisites or not, you can use this guide:

* You need [docker](https://www.docker.com/get-started) to build, pull or push images. You also need docker if you want to run your application as a container.
* You need [kubectl](https://kubernetes.io/docs/home/#learn-how-to-use-kubernetes) to access your kubernetes cluster.
* You need [helm](https://helm.sh/) and [pulumi](https://www.pulumi.com/) to deploy your application in kubernetes cluster.
* You need [tocer](https://github.com/state-alchemists/tocer) to scaffold Zaruba's documentation.
* You need [pyenv](https://github.com/pyenv/pyenv) to run many Python versions.
* You need [nvm](https://github.com/nvm-sh/nvm) to run many Node.Js versions.

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

You can always open [an issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).

When opening a pull request, please write down:

* The version of Zaruba you are using. You can run `zaruba version` to get the version.
* Your expectation/goal.
* What you have tried.
* The result you get.

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

# 🎉 Fun fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)

![Madou Ring Zaruba on Kouga's Hand](arts/madou-ring-zaruba.jpg)
