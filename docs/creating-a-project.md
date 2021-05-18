# Creating a Project

A zaruba project is a git repository containing at least one file: `main.zaruba.yaml`.

In order to initialize a project, you can invoke `zaruba please initProject`

## Example

```sh
mkdir newProject
cd newProject
zaruba please initProject
```

## Some terminologies

* `Project`: A container for tasks, environments, configurations, and probably services.
* `Task`: Declarative definition about how to run something. A project might contains a lot of task that depends/extend to each other. A task can be executed manually by invoking `zaruba please invoke <taskName> -i`.
* `Docker task`: Docker task is a task that run/start a docker container in your computer.
* `Service task`: Service task is a task that run app/service in your local computer.
* `Service`: Any long-running application is considered as service (e.g: web server, database server, etc)
* `Legacy code`: Code that you hate so much but you need to deal with. Usually written by someone who has resign 2 years ago.
* `Monorepo`: A git repository contains all services and the entire universe.
* `Multirepo`: Multiple git repositories that suppose to work and compatible with each others (but usually does not).

## Involved tasks

* [initProject](tasks/initProject.md)

## What's next


* [Creating Fast API service](creating-fast-api-service.md)
* [Creating docker task](creating-docker-task.md)
* [Creating service task](creating-service-task.md)
* [Understanding task](understanding-task.md)
* [Working with legacy code](working-with-legacy-code.md)
* [Starting a monorepo](starting-a-monorepo.md)
* [Working with multirepo](working-with-multirepo.md)