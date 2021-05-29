# Running Tasks

There are several ways to run tasks:

* Interactively: `zaruba please <task-name-1> <task-name-2> ... -i`
* With Parameters `zaruba please <task-name> [param-1=value-1] [param-2-value-2] ...`

To see all available tasks and run/explain them, you can simply type `zaruba please`.

# Running generated service/docker task

If your task was created by using zaruba's generator (i.e: [generated docker task](creating-docker-task.md) or [generated service task](creating-service-task.md)), then your task is typically has this naming: `run-<service-name>`.

Every generated service task also accompanied with several auto-generated tasks:

* `run<Service-name>Container`: Run the service as docker container. Require `Dockerfile` in your service directory
* `stop<Service-name>Container`: Stop container
* `remove<Service-name>Container`: Stop and remove container
* `build<service-name>Image`: Build image. Require `Dockerfile` in your service directory.
* `push<service-name>Image`: Push image to image repository.

Similarly, every generated docker task is accompanied with several auto-generated tasks:

* `stop<Service-name>Container`: Stop container
* `remove<Service-name>Container`: Stop and remove container

Aside from the auto-generated tasks, Zaruba will also add several common tasks to your project:

* `run`: Run all services (not containerized) + all docker container
* `runContainer`: Run all services (containerized) + all docker container. This require `Dockerfile` in every service directory.
* `stopContainer`: Stop all containers.
* `removeContainer`: Stop and remove all containers.
* `buildImage`: Build all service images. This also require `Dockerfile` in every service directory.
* `pushImage`: Push all service images to image repository.

Plese take note that containers are not going to be stopped when you press `ctrl+c`.

# Example

```sh
# run task interactively.
zaruba please run -i
zaruba please runContainer -i

# run tasks in parallel
zaruba please runService1 runService2

# run task interactively and kill immediately once it is complated.
zaruba please runContainer -i -t


# run task interactively, wait for 5s after it is complated and kill it.
zaruba please runContainer -i -t -w 5s

# explain the task
zaruba please runContainer -x
```

# Involved tasks

* [core.startService](tasks/core.startService.md) (extended)
* [core.startDockerContainer](tasks/core.startDockerContainer.md) (extended)
* [core.stopDockerContainer](tasks/core.stopDockerContainer.md) (extended)
* [core.buildDockerImage](tasks/core.buildDockerImage.md) (extended)
* [core.pushDockerImage](tasks/core.pushDockerImage.md) (extended)

# What's next

* [Creating docker task](creating-docker-task.md)
* [Creating service task](creating-service-task.md)
* [Creating fastAPI service](creating-fast-api-service.md)
* [Understanding the concept](concept.md)