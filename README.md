![zaruba-logo](arts/zaruba-250.png)

> ⚠️ Things might change pretty fast and pretty often before we reach `v.1.0.0`. Please open issue if you find any problem using Zaruba.

# 💀 Zaruba 

Zaruba is a declarative Task Runner Framework. It helps you to define and orchestrate tasks in a fun way. 

Zaruba allows you to declare configurable task that extends/depends on each other. You will find it better than writing a bunch of spaghetti shell scripts.

You will find several pre-defined tasks. Some of them are useful to speed up your development by providing scaffolding.

You can even build a full-fledge FastAPI application and have it deployed to your Kubernetes cluster by performing this (no coding required 😉):

```sh
# Init project
mkdir myProject
cd myProject
zaruba please initProject

# Create FastAPI app with functional book CRUD
zaruba please addFastApiCrud \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudFields='["title","author","synopsis"]'

# Run the service locally 
# To run this command, you need:
# - pyenv
zaruba please start
# Ctrl+c to stop

# Run the service as docker container
# To run this command, you need:
# - docker
zaruba please startContainers
zaruba please stopContainers

# Deploy the service to the kubernetes cluster
# To run this command, you need:
# - kubectl
# - helm
# - pulumi
# - cloud provider or a computer that can run kubernetes locally
zaruba please pushContainers
zaruba please addAppDeployment appDirectory=myApp
zaruba please syncEnv
zaruba please deploy kubeContext=docker-desktop
```

> 💡 __TIPS:__ Execute tasks with `-i` or `--interactive` flag is probably a good idea if you don't want to memorize the parameters. Otherwise, you can also type `zaruba please` to select available tasks.

# 👨‍💻 Installation

## Using docker

Using docker is probably the quickest way to set up Zaruba, especially if you need to use Zaruba in your CI/CD.

For more information about Zaruba's docker image, please visit [dockerhub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> **⚠️NOTE** There will be some limitations if you run Zaruba container using docker-desktop for mac/windows. For example, docker-desktop doesn't support host networking, so that you need to expose the ports manually (e.g: `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

## From source

Installing from source is the best way to setup Zaruba for day-to-day use. Currently we don't have any plan to create `apt` or platform-specific packages for Zaruba. If you are using windows, you need to install `wsl` in order to get started.

In order to install Zaruba from source, you need to have some prerequisites software:

* `go 1.13` or newer (To install `go` quickly you can visit it's [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> **💡HINT** Ubuntu user (including ubuntu-wsl) can simply invoke `sudo apt-get install golang wget curl git` to install all prerequisites.

After having the prerequisites installed you can then install Zaruba by using `curl` or `wget`:

```sh
# Install zaruba by using curl
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"

# Install zaruba by using wget
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

# 📜 Getting Started

Before getting started, it is recommended to have `docker`, `kubectl`, `helm`, `pyenv`, `pipenv`, and `nvm` installed. To install those prerequisites, please visit their websites:

* [docker](https://www.docker.com/get-started) is needed to build, pull or push image. You also need docker to run your services as container.
* [kubectl](https://kubernetes.io/docs/home/#learn-how-to-use-kubernetes) is needed to access your kubernetes cluster.
* [helm](https://helm.sh/) is needed to deploy your services.

You should also able to install those third party packages by running zaruba's third party installer:

```sh
zaruba install docker
zaruba install kubectl
zaruba install helm
```

Now let's get started by [creating a project](docs/creating-a-project.md)


# 🗺️ Roadmap


## Doing

* Technical Documentation
* Third party script (i.e: Install script from github repository)

## To do

* UI (i.e: web server)
* NLP (i.e: running tasks by using natural language)
* OSX Setup

# 🐞 Bug, Feature Request and Contribution

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).

# ☑️ Testing

To perform test, you need to have:

* docker desktop
* kubectl
* helm
* pyenv and pipenv
* go 1.13

Once the prerequisites met, you can perform:

```
make test
```

# 🎉 Fun Fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
