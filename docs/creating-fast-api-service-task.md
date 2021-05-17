# Creating Fast API Service Task

To create a Fast API service task, you can perform: `zaruba please makeFastApiServiceTask -i`.

## Example

```sh
# run interactively
zaruba please makeFastApiServiceTask -i

# run with parameters
zaruba please makeFastApiServiceTask generator.service.location=./some-directory/myService generator.service.name=myService generator.service.envs="MY_ENV=MY_VALUE" ports=3000 generator.service.docker.image.name=myService generator.service.docker.container.name=myServiceContainer
```

## Involved tasks:

* [makeFastApiServiceTask](tasks/makeFastApiServiceTask.md)

## What's next:

TODO