![zaruba-logo](arts/zaruba-250.png)
# 💀 Zaruba 

Zaruba is a declarative Task Runner Framework. It helps you to define and orchestrate tasks. 

In zaruba, a task might extends or depends on each other. This make developing tasks with zaruba is more flexible than creating bunch of shell scripts.

To speed up your development, several tasks has been already included and can be extended accordingly. You can even build a full-fledge FastAPI application and having it deployed to your kubernetes cluster by performing this (no coding required):

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

> 💡 __TIPS:__ Execute task with `--interactive` flag is probably a good idea if you don't want to memorize the parameters.

# 👨‍💻 Installation

## Using docker

Using docker is probably the quickest way to setup zaruba, especially if you need to use zaruba in your CI/CD.

For more information about zaruba's docker image, please visit [dockerhub](https://hub.docker.com/repository/docker/stalchmst/zaruba).

> **⚠️NOTE** There will be some limitations if you run zaruba container using docker-desktop for mac/windows. For example, docker-desktop doesn't support host networking, so that you need to expose the ports manually (e.g: `docker run -d --name zaruba -p 8200-8300:8200-8300 -v "$(pwd):/project" stalchmst/zaruba:latest`)

## From source

Installing from source is the best way to setup zaruba for day-to-day use. Currently we don't have any plan to create `apt` or platform-specific packages for zaruba. If you are using windows, you need to install `wsl` in order to get started.

In order to install zaruba from source, you need to have some prerequisites software:

* `go 1.13` or newer (To install `go` quickly you can visit it's [official website](https://golang.org/doc/install))
* `wget` or `curl`
* `git`

> **💡HINT** Ubuntu user (including ubuntu-wsl) can simply invoke `sudo apt-get install golang wget curl git` to install all prerequisites.

After having the prerequisites installed you can then install zaruba by using `curl` or `wget`:

```sh
# Install zaruba by using curl
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"

# Install zaruba by using wget
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

# 📜 Simple Use Case

As software engineer you need to work with a lot of things. Some of those things are probably under your control, but in most cases you also need to deal with something beyond your control.

Probably you need to do something with with legacy-code, distributed monolith, and other monstrousities. In some cases you need to run some of those legacy-codes in your computer just to figure out how things work before adding some features here and there.

Sure, refactoring is probably a good idea, but not for now. Not when you don't know how things suppose to work in those hyped-driven-without-unit-tests things.

To begin with, suppose you have clone two services named `bahemoth` and `leviathan`. Bahemoth was written by using Node.Js version 10, while leviathan was written by using Node.Js 14. Leviathan is somehow depends on bahemoth. You cannot run leviathan before bahemoth is running. Also, bahemoth is actually contains of two services, `bahemoth-master` and `bahemoth-worker`. But don't worry, `bahemoth-master` and `bahemoth-client` share the same code-base.

Okay, to recap you have this:

```
# The codebase

bahemoth-code (Node.Js v10)  ---runAs--> bahemoth-master
                                |
                                |------> bahemoth-worker

leviathan-code (Node.Js v14) --runAs--> leviathan


# Service dependencies

bahemoth-master ---neededBy--> bahemoth-worker
                   |
                   |---------> leviathan
```

# 🗺️ Roadmap


## Doing

* Refactor
* Technical Documentation

## To do

* Third party script (i.e: Load script from github repository)
* NLP (i.e: running tasks by using natural language)
* OSX Setup

# 🐞 Bug, Feature Request and Contribution

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).


# 🎉 Fun Fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
