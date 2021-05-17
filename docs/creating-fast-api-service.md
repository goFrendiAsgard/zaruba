# Creating Fast API Service

FastAPI is a modern, fast (high-performance), web framework for building APIs with Python 3.6+ based on standard Python type hints. You can visit [Fast API official website](https://fastapi.tiangolo.com/) for more information.

When you make a Fast API service, zaruba will create a file named `main.py`. You can thing this `main.py` as dependency injection container without any magic. Thus, you need to initialize every object you are going to need in this file.

But of course, in most cases, you can just use zaruba's generator so that you don't need to touch the file at all. You can even make a fully working CRUD API without touching your text editor.

To create a Fast API service, you can perform: `zaruba please makeFastApiService -i` and follow the on screen instruction.

## Example

```sh
# run interactively
zaruba please makeFastApiService -i

# run with parameters
zaruba please makeFastApiService generator.fastApi.service.name=myService
```

## Involved tasks

* [makeFastApiService](tasks/makeFastApiService.md)


## What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API module](creating-fast-api-module.md)
* [Creating Fast API route](creating-fast-api-route.md)
* [Creating Fast API event handler](creating-fast-api-event-handler.md)
* [Creating Fast API RPC handler](creating-fast-api-rpc-handler.md)
* Creating Fast API CRUD