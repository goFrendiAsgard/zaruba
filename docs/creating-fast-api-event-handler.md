# Creating Fast API Event Handler

At some point, your service might need to send event (and message) without waiting for any reply. This asynchronous communication is commonly used in microservices since it let your services to be decoupled from each other.

This kind of communication pattern usually require third party message bus such as kafka, nats, or rabbitmq.

You can read more about supported message bus [here](fast-api-message-bus.md). It is strongly recommended that you read it first before continuing any further.

Also You need to make sure that you already have [Fast Api service](creating-fast-api-service.md) and [Fast Api module](creating-fast-api-module.md)

To make Fast API event handler, you can invoke `zaruba please makeFastApiEventHandler -i`

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

        @self.mb.handle_event('myEvent')
        def handle_event_my_event(message: Mapping[str, Any]):
            print('handle event myEvent with message: {}'.format(message))

        print('Handle events for myModule')
    

    def handle_route(self):
        print('Handle routes for myModule')

```

To trigger an event, you can call `mb.publish('event_name', message)`

## Example

```sh
# run interactively
zaruba please makeFastApiEventHandler -i

# run with paramter
zaruba please makeFastApiEventHandler generator.fastApi.service.name=myService generator.fastApi.module.name=myModule generator.fastApi.event.name=myEvent
```

## Involved tasks

* [makeFastApiEventHandler](tasks/makeFastApiEventHandler.md)


## What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API route](creating-fast-api-route.md)
* [Creating Fast API RPC handler](creating-fast-api-rpc-handler.md)
* [Creating Fast API Crud](creating-fast-api-crud.md)