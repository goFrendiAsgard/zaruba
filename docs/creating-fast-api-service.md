# Creating Fast API Service

FastAPI is a modern, fast (high-performance), web framework for building APIs with Python 3.6+ based on standard Python type hints. You can visit [Fast API official website](https://fastapi.tiangolo.com/) for more information.

To create a Fast API service, you can perform: `zaruba please makeFastApiService -i` and follow on screen instruction.

## Example

```sh
# run interactively
zaruba please makeFastApiService -i

# run with parameters
zaruba please makeFastApiService generator.fastApi.service.name=myService
```

## Involved Tasks

* [makeFastApiService](tasks/makeFastApiService.md)




## What's next

* Creating Fast API service task
* Creating Fast API module
* Creating Fast API route
* Creating Fast API event handler
* Creating Fast API RPC handler