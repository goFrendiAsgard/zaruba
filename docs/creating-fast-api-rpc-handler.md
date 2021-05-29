# Creating Fast API RPC Handler

Sometime your service needs to talk to each other. In this case, you can use RPC (Remote procedure call).

RPC might need/don't need message bus, but currently zaruba's Fast API service only support RPC with message bus.

You can read more about supported message bus [here](fast-api-message-bus.md). It is strongly recommended that you read it first before continuing any further.

Also You need to make sure that you already have [Fast Api service](creating-fast-api-service.md) and [Fast Api module](creating-fast-api-module.md)

To make Fast API RPC handler, you can invoke `zaruba please makeFastApiRPCHandler -i`

Here is how your controller looks like after you generate an event handler:


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
            self.handle_event()
        if self.enable_route:
            self.handle_route()
    

    def handle_event(self):

        @self.mb.handle_rpc('myRpc')
        def handle_rpc_my_rpc(parameter: str) -> str:
            print('handle RPC call myRpc with parameter: {}'.format(parameter))
            return parameter

        print('Handle events for myModule')
    

    def handle_route(self):
        print('Handle routes for myModule')

```

To call an RPC, you can call `mb.call_rpc('rpc_name', *parameters)`. The function should return a result depending on what you return in the corresponding RPC handler.


# Example

```sh
# run interactively
zaruba please makeFastApiRpcHandler -i

# run with paramter
zaruba please makeFastApiRpcHandler generator.fastApi.service.name=myService generator.fastApi.module.name=myModule generator.fastApi.rpc.name=myRpc
```

# Involved tasks

* [makeFastApiRpcHandler](tasks/makeFastApiRpcHandler.md)


# What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API route](creating-fast-api-route.md)
* [Creating Fast API event handler](creating-fast-api-event-handler.md)
* [Creating Fast API Crud](creating-fast-api-crud.md)