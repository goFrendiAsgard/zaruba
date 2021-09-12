# Creating Fast API Route

After creating FastAPI service and module, you might want to handle HTTP request. HTTP is probably the most common/well-known communication protocol in the internet.

In order to serve HTTP request, and send necessary response, you need to provide:

* HTTP URL (i.e: `/`, `/api`, etc)
* HTTP method (get/post/put/delete)

The generated HTTP URL handler (or route) is located at your module's controller.py.
Thus, you have to make sure that you already have a [Fast API service](creating-fast-api-service.md) and [Fast API module](creating-fast-api-module.md).

Here is how your controller looks like once you generate a route handler:

```python
from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus

import traceback

class Controller():

    def __init__(self, app: FastAPI, mb: MessageBus, enable_route: bool, enable_event: bool):
        self.app = app
        self.mb = mb
        self.enable_route = enable_route
        self.enable_event = enable_event


    def start(self):
        if self.enable_event:
            self.handle()
        if self.enable_route:
            self.handle_route()
    

    def handle(self):
        print('Handle events for myModule')
    

    def handle_route(self):

        @self.app.get('/hello')
        def get_hello():
            try:
                return 'OK'
            except HTTPException as error:
                raise error
            except Exception as error:
                print(traceback.format_exc()) 
                raise HTTPException(status_code=500, detail='Internal Server Error')

        print('Handle routes for myModule')
```

As you guess, the function to handle the route is located inside `handle_route` method. Whenever your Fast API service receive `GET /hello` HTTP request, it will give `OK` response.

To create Fast API route, you can perform: `zaruba please makeFastApiRoute -i` and follow the on screen instruction.

# Example

```sh
# run interactively
zaruba please makeFastApiRoute -i

# run with parameter
zaruba please makeFastApiRoute generator.fastApi.service.name=myService generator.fastApi.module.name=myModule generator.fastApi.httpMethod=get generator.fastApi.url='/hello'
```


# Involved tasks

* [makeFastApiRoute](tasks/makeFastApiRoute.md)


# What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API event handler](creating-fast-api-event-handler.md)
* [Creating Fast API RPC handler](creating-fast-api-rpc-handler.md)
* [Creating Fast API Crud](creating-fast-api-crud.md)