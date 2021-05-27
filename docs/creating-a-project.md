# Creating a Project

A Zaruba project is a git repository containing `main.zaruba.yaml`. 

To initialize a project, you can invoke `zaruba please initProject`

## Example

```sh
mkdir newProject
cd newProject
zaruba please initProject
```

## Some terminologies

* `Project`: A container for tasks, environments, and configurations.
    - `Task`: A task definition. A project might contain a lot of tasks that depend/extend to each other. You can execute a task interactively by invoking `zaruba please invoke <taskName> -i`.
        - `Command Task`: A task that will be terminated upon completion.
        - `Service Task`: A task that might still be running after completion (i.e: long-running process). Typically this `service task` starts a service/docker container.
            - `Docker Task`: A kind of service task that starts a docker container.
        - `Wrapper Task`: A task that does nothing except wait for other task's completion.
* `Service`: Any long-running application is considered a service (E.g: web server, database server, etc)
* `Legacy code`: Code that you hate so much but you need to deal with. Probably written by someone who has resign 2 years ago.
* `Monorepo`: A git repository contains all services and the entire universe.
* `Multirepo`: Multiple git repositories that suppose to work and compatible with each others (but probably does not).

## Involved tasks

* [initProject](tasks/initProject.md)

## What's next


* [Creating Fast API service](creating-fast-api-service.md)
* [Creating docker task](creating-docker-task.md)
* [Creating service task](creating-service-task.md)
* [Understanding the concepts](concept.md)
* [Working with legacy code](working-with-legacy-code.md)
* [Starting a monorepo](starting-a-monorepo.md)
* [Working with multirepo](working-with-multirepo.md)