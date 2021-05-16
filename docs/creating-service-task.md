# Creating Service Task

To create a service task, you can perform: `zaruba please makeServiceTask -i`.

Zaruba also provide several preset docker-tasks, namely:

* FastAPI
* Go
* NodeJs
* Python


## Example

```sh
# run interactively
zaruba please makeServiceTask -i
zaruba please makeGoServiceTask -i
zaruba please makeFastApiServiceTask -i

# run with parameters
zaruba please makeServiceTask generator.service.location=./some-directory/myService generator.service.name=myService generator.service.envs="MY_ENV=MY_VALUE" ports=3000 generator.service.docker.image.name=myService generator.service.docker.container.name=myServiceContainer
```

## Involved Tasks:

* [makeServiceTask](tasks/makeServiceTask.md)
* [makeFastApiServiceTask](tasks/makeFastApiServiceTask.md)
* [makeGoServiceTask](tasks/makeGoServiceTask.md)
* [makeNodeJsServiceTask](tasks/makeNodeJsServiceTask.md)
* [makePythonServiceTask](tasks/makePythonServiceTask.md)


## What's next

* Running tasks
* [Creating docker task](creating-docker-task.md): So you are working with docker? Running third party application as docker container is a good idea. If you want to add MySQL/Cassandra/RabbitMq/Redis container to your project, you need to go here.
* [Creating task manually](understanding-task.md): Do you want to understand zaruba script in detail so that you can make your own task without using any generator? Then this is going to be the right step.
