# Creating a Project

A zaruba project is a git repository containing at least one file: `main.zaruba.yaml`.

In order to initialize a project, you can invoke `zaruba please initProject`

## Example

```
mkdir newProject
cd newProject
zaruba please initProject
```

## Involved Tasks

* [initProject](tasks/initProject.md)

## What's next


* [Creating Fast API service](creating-fast-api-service.md): Creating a Fast API service/app by using zaruba is very easy and enjoyable. If you don't have any existing code to begin with or if you want to create a brand new Fast API service, this is going to be the right step.
* [Creating docker task](creating-docker-task.md): So you are working with docker? Running third party application as docker container is a good idea. If you want to add MySQL/Cassandra/RabbitMq/Redis container to your project, this is going to be the right step.
* [Creating service task](creating-service-task.md): Sometime you might need to run a service/app without containerize it. If running non-containerized service is what you want, this is going to be the right step.
* [Creating task manually](understanding-task.md): Do you want to understand zaruba script in detail so that you can make your own task without using any generator? Then this is going to be the right step.
* [Working with legacy code](working-with-legacy-code.md)