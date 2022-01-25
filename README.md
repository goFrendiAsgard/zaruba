![zaruba-logo](arts/zaruba-250.png)

> ⚠️ Things might change pretty fast and pretty often before we reach `v.1.0.0`. Please open issue if you find any problem using Zaruba.

# 💀 Zaruba 

Zaruba is a [task](docs/core-concept/project/task/README.md) runner and [CLI utilities](docs/utilities/README.md). It helps you to `write`, `generate`, and `orchestrate` tasks quickly.

## Problem

While developing your applications, you might find yourself opening several `tmux` panels and running some commands in parallel. You might also find that some tasks could only be executed once their dependencies are executed. For example, a web application can only be started after the database server is running. Not only complicated, this also lead to human errors.

## Solution

Zaruba exists to solve those problems by allowing you to define configurable tasks that can extend/depend on each other. Furthermore, dependency tasks might run in parallel. This will greatly reduce human error and save your time.

Some [built-in tasks](docs/core-tasks/README.md) are also available. Ready to be used/extended to match your use case. You can run `zaruba please` to see the list of available tasks.

## Example

You can build a full-fledge FastAPI application and have it deployed to your Kubernetes cluster by performing these commands (no coding required 😉):

```bash
# ✨ Init project
mkdir myProject
cd myProject
zaruba please initProject

# Add 🐬 MySQL container
zaruba please addMysql appDirectory=myDb

# Add ⚡ FastAPI app with book CRUD API.
zaruba please addFastApiCrud \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudFields='["title","author","synopsis"]' \
    appEnvs='{"APP_SQLALCHEMY_DATABASE_URL": "mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' \
    appDependencies='startMyDbContainer'

# Run ⚡ FastAPI app + 🐬 MySQL container
# To run this command, you need:
# - python 3.8
# - docker
zaruba please start
# Ctrl+c to stop

# Run ⚡ FastAPI app + 🐬 MySQL (both as containers)
# Run FastAPI app as docker container
# To run this command, you need:
# - docker
zaruba please startContainers
zaruba please stopContainers

# (Optional) Deploy ⚡ FastAPI app to the kubernetes cluster
# To run this command, you need:
# - kubectl
# - helm
# - pulumi
# - cloud provider or a computer that can run kubernetes locally (we use docker-desktop in this example)
zaruba please pushContainers
zaruba please addAppDeployment appDirectory=myApp
zaruba please syncEnv
zaruba please deploy kubeContext=docker-desktop
```

> 💡 __TIPS:__ You can execute tasks with `-i` or `--interactive` flag (i.e: `zaruba please addFastApiCrud -i`).


# 👨‍💻 Installation

## Using docker

Using docker is the quickest way to set up Zaruba, especially if you need to use Zaruba in your CI/CD.

For more information about Zaruba's docker image, please visit [dockerhub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> **⚠️ NOTE** There will be some limitations if you run Zaruba container using `docker-desktop` for mac/windows. For example, docker-desktop doesn't support host networking, so that you need to expose the ports manually (e.g: `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

## From source

Installing from source is the best way to set up Zaruba for day-to-day use. Currently, we don't have any plan to create `apt` or platform-specific packages for Zaruba. If you are using windows, you need to install `wsl` in order to get started.

In order to install Zaruba from the source, you need to have some prerequisites software:

* `go 1.13` or newer (To install `go` quickly you can visit its [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> **💡HINT** Ubuntu user (including ubuntu-wsl) can simply invoke `sudo apt-get install golang wget curl git` to install all prerequisites.

After having the prerequisites installed you can then install Zaruba by using `curl`:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

 or `wget`:

 ```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

# 📜 Getting started

## Additional prerequisites

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

## Let's get started

Finally, let's get started by:

* [📖 Skimming documentation's TOC](docs/README.md)
* [🧙‍♂️ Learning the core concept](docs/core-concept/README.md), or 
* [🪄 Creating a project](docs/use-cases/creating-a-project.md)

# 🐞 Bug, feature request and contribution

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).

Whenever you open an issue, please make sure to let us know:

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

Once the prerequisites are met, you can perform:

```bash
make test
```

# 🎉 Fun fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
