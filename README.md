![zaruba-logo](arts/zaruba-250.png)

> ⚠️ Things might change pretty fast and pretty often before we reach `v.1.0.0`. Please open issue if you find any problem using Zaruba.

# 💀 Zaruba 

Zaruba is a [task](docs/core-concepts/project/task/README.md) runner and [CLI utilities](docs/utilities/README.md). It helps you to `write`, `generate`, `orchestrate`, and `run` tasks quickly.

## ❓ Problem

While developing/debugging applications, you might need to run several tasks __in parallel__ or __in particular order__.

Some tasks might __depends__ on several pre-requisited, while some others probably __share similar behavior__. For example:

* A web application can only be started once the database server is running. In this case, web application __depends__ on database server.
* To start TypeScript applications, you need to perform `npm install` and `tsc`. In this case, TypeScript applications __share similar behavior__.

When you manually run tasks in parallel or sequentially, you might do some mistakes and you need to re-do everything from scratch.

## 💡 Solution

Creating __clear instructions/checklist__ might help. If your tasks are simple and sequential, you can turn your checklist into a nice [shell script](https://www.shellscript.sh/first.html), and that's all you need.

However, when your workflow become more complicated, you need a __better automation tool__ like Zaruba.

Zaruba allows you to __simplify your workflow__ by let you:

* Create configurable tasks (i.e: using `configs`, `inputs`, or `envs`).
* Define tasks dependencies (i.e: using `extend`).
* Re-use and share configurations/behviors (i.e: using `extend`, `configRef`, or `envRef`).
* Run tasks in parallels.
* Generate new tasks.

There are several [built-in tasks](docs/core-tasks/README.md) specially crafted to achieve those goals. To see list of available tasks, you can run `zaruba please`.

## 🔍 Example

Suppose you want to build two applications:

* A 🐍 `CRUD API application` that __depends__ on 🐬 `MySQL`.
* A simple 🐹 `Go web server` that has no dependencies.

Since `Go web server` has no dependencies, you should be able to run it __in parallel__ with `CRUD API application` and `MySQL server`.

On the other hand, `CRUD API application` __depends__ on `MySQL server`. Thus, you cannot run `CRUD API application` without running `MySQL server`.

Zaruba allows you to build, run, and deploy everything using simple commands (no coding required 😉).

> 💡 __TIPS:__ You can execute tasks with `-i` or `--interactive` flag (i.e: `zaruba please addFastApiCrud -i`).

### ✨ Creating Project and Applications

```bash
# ✨ Init project
mkdir myProject
cd myProject
zaruba please initProject

# Add 🐬 MySQL container
zaruba please addMysql appDirectory=myDb

# Add 🐍 CRUD API Application.
zaruba please addFastApiCrud \
  appDirectory=myPythonApp \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["myDb"]' \
  appEnvs='{"APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

# Add 🐹 Go web app.
zaruba please addSimpleGoApp appDirectory=myGoApp appEnvs='{"APP_HTTP_PORT":"3001"}'
```

### 🏃 Run Applications

```bash
# Start 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL container.
# To run this command, you need:
# - go 1.13 or newer
# - python 3.8
# - docker
zaruba please start
# Ctrl+c to stop
```

<details>
<summary>You can also run applications individually</summary>

```bash
# Only start 🐹 Go web server.
zaruba please startMyGoApp
# Ctrl+c to stop

# Only start 🐹 Go web server and 🐍 CRUD API application.
# Please note that MySQL container is automatically started
# since CRUD API application depends on it.
zaruba please startMyGoApp startMyPythonApp
# Ctrl+c to stop
```
</details>

### 🐳 Run Applications as Containers

```bash
# Start 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL as containers.
# To run this command, you need:
# - docker
zaruba please startContainers
zaruba please stopContainers
```

<details>
<summary>You can also run applications individually</summary>

```bash
# Only start 🐹 Go web server.
zaruba please startMyGoAppContainer
zaruba please stopContainers

# Only start 🐹 Go web server and 🐍 CRUD API application.
# Please note that MySQL container is automatically started
# since CRUD API application depends on it.
zaruba please startMyGoAppContainer startMyPythonAppContainer
zaruba please stopContainers
```
</details>

### ☁️ Deploy Applications

```bash
# Deploy 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL to kubernetes cluster
# To run this command, you need:
# - kubectl
# - helm
# - pulumi
# - cloud provider or a computer that can run kubernetes locally (we use docker-desktop in this example)
zaruba please buildImages # or `zaruba please pushImages`
zaruba please addAppKubeDeployment appDirectory=myPythonApp
zaruba please addAppKubeDeployment appDirectory=myGoApp
zaruba please addAppKubeDeployment appDirectory=myDb
zaruba please syncEnv
zaruba please deploy kubeContext=docker-desktop
zaruba please destroy kubeContext=docker-desktop
```

# 👨‍💻 Installation

## 📖 From Source

Installing from source is the best way to set up Zaruba for day-to-day use.

Currently, we don't have any plan to create `apt` or platform-specific packages for Zaruba. If you are using windows, you need to install `wsl` in order to get started.

### 🧅 Prerequisites

Before installing Zaruba from the source, you need to install some prerequisites software:

* `go 1.13` or newer (To install `go` quickly you can visit its [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> **💡HINT** Ubuntu user (including ubuntu-wsl) can simply invoke `sudo apt-get install golang wget curl git` to install all prerequisites.

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

To create and run zaruba container on __🐧linux__ host, you can do:

```bash
docker run -d --name zaruba --network host \
  -v "$(pwd):/project" \
  -e "ZARUBA_HOST_DOCKER_INTERNAL=172.17.0.1" \
  -e "DOCKER_HOST=tcp://172.17.0.1:2375" \
  stalchmst/zaruba:latest
```

To create and run zaruba container on __🪟 windows__/__🍎 mac__ host, you can do:

```bash
docker run -d --name zaruba \
  -p 8500-8700:8500-8700 \
  -v "$(pwd):/project" \
  stalchmst/zaruba:latest
```

For more information about Zaruba's docker image, please visit [dockerhub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> **⚠️ NOTE** There will be some limitations if you run Zaruba container using `docker-desktop` for mac/windows. For example, docker-desktop doesn't support host networking, so that you need to expose the ports manually (e.g: `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

# 📜 Getting started

To get started, you can start:

* [📖 browsing the documention](docs/README.md)
* [🧙‍♂️ learning the core concept](docs/core-concepts/README.md), or 
* [🪄 creating a project](docs/use-cases/create-a-project.md)

But before doing that, you probably need to install additional prerequisites.

## ➕ Additional prerequisites

Before getting started, it is recommended to have `docker`, `kubectl`, `helm`, and `pulumi` installed. To install those prerequisites, you can visit their websites or simply invoke `zaruba install`.

To see whether you need to install those pre-requisites or not, you can use this guide:

* [docker](https://www.docker.com/get-started) is needed to build, pull or push images. You also need docker if you want to run your application as a container.
* [kubectl](https://kubernetes.io/docs/home/#learn-how-to-use-kubernetes) is needed to access your kubernetes cluster.
* [helm](https://helm.sh/) and [pulumi](https://www.pulumi.com/) is needed to deploy your application in kubernetes cluster.

You should also be able to install those third party packages by running zaruba's third party installer:

```bash
zaruba install docker
zaruba install kubectl
zaruba install helm
zaruba install pulumi
```

# 🐞 Bug, feature request and contribution

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).

If you open an issue, please make sure to let us know:

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

Once the prerequisites are met, you can perform:

```bash
make test
```

# 🎉 Fun fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)

![Madou Ring Zaruba on Kouga's Hand](arts/madou-ring-zaruba.jpg)