![zaruba-logo](arts/zaruba-250.png)
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

# Create FastAPI service, module, and book CRUD
zaruba please makeFastApiService generator.fastApi.service.name=myService
zaruba please makeFastApiModule generator.fastApi.service.name=myService generator.fastApi.module.name=myModule
zaruba please makeFastApiCrud generator.fastApi.service.name=myService generator.fastApi.module.name=myModule generator.fastApi.crud.entity=book generator.fastApi.crud.fields=title,author,synopsis

# Create task to start zaruba
zaruba please makeServiceTask generator.service.location=myService generator.service.type=fastapi

# Run service locally
zaruba please run

# Run service (this time containerized)
zaruba please runContainer

# Create helm chart
zaruba please makeHelmCharts

# Create helm deployment for myService
zaruba please makeHelmDeployment generator.service.name=myService

# Apply helm deployments to docker-desktop kubernetes cluster
zaruba please helmApply kube.context=docker-desktop
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

Before getting started, it is recommended to have `pyenv` and `nvm` installed. To install those prerequisites, you can perform:

```
zaruba please setupPyenv
zaruba please setupNvm
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


# 🎉 Fun Fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
