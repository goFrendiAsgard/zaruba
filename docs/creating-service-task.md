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

## Involved tasks:

* [makeServiceTask](tasks/makeServiceTask.md)
* [makeFastApiServiceTask](tasks/makeFastApiServiceTask.md)
* [makeGoServiceTask](tasks/makeGoServiceTask.md)
* [makeNodeJsServiceTask](tasks/makeNodeJsServiceTask.md)
* [makePythonServiceTask](tasks/makePythonServiceTask.md)


## What's next

* [Running tasks](running-task.md)
* [Creating docker task](creating-docker-task.md)
* [Creating Fast API service](creating-fast-api-service.md)
* [Understanding task](understanding-task.md)