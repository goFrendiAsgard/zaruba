# Creating Fast API Module

Fast API module is not part of Fast API itself. Instead it is an opinionated approach to manage your code.

Zaruba's need Fast API module to generate route, event handler, RPC handler, and CRUD.

Module should be isolated from each others. Furthermore, if a module need anything like database connection or message bus, they should be injected via module controller's constructor.

In that sense, our `main.py` only contains component constructors as well as how those components interacting with each other. Let's take a look on a `main.py` example:

```python
from fastapi import FastAPI
from sqlalchemy import create_engine
from helpers.transport import MessageBus, RMQMessageBus, RMQEventMap, LocalMessageBus
from myModule.controller import Controller as MyModuleController
from repos.dbBook import DBBookRepo


import os

def create_message_bus(mb_type: str) -> MessageBus:
    if mb_type == 'rmq':
        rmq_host = os.getenv('MY_SERVICE_RABBITMQ_HOST', 'localhost')
        rmq_user = os.getenv('MY_SERVICE_RABBITMQ_USER', 'root')
        rmq_pass = os.getenv('MY_SERVICE_RABBITMQ_PASS', 'toor')
        rmq_vhost = os.getenv('MY_SERVICE_RABBITMQ_VHOST', '/')
        rmq_event_map = RMQEventMap({})
        return RMQMessageBus(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalMessageBus()

db_url = os.getenv('MY_SERVICE_SQLALCHEMY_DATABASE_URL', 'sqlite://')
mb_type = os.getenv('MY_SERVICE_MESSAGE_BUS_TYPE', 'local')
enable_route = os.getenv('MY_SERVICE_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event = os.getenv('MY_SERVICE_ENABLE_EVENT_HANDLER', '1') != '0'

engine = create_engine(db_url, echo=True)
app = FastAPI()
mb = create_message_bus(mb_type)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()

book_repo = DBBookRepo(engine=engine, create_all=True)
my_module_controller = MyModuleController(app=app, mb=mb, enable_route=enable_route, enable_event=enable_event, book_repo=book_repo)
my_module_controller.start()
```

Module initialization is defined at the very bottom of the code:

```python
my_module_controller = MyModuleController(app=app, mb=mb, enable_route=enable_route, enable_event=enable_event, book_repo=book_repo)
my_module_controller.start()
```

First you make `my_module_controller`. In order to do that, you need several parameters (`app`, `mb`, `enable_route`, `enable_event`, and `book_repo`). After everything has been set, you then start `my_module_controller` by calling it's `start` method.

Zaruba add module as python package inside Fast API service. It also automatically import the generated module into your `main.py`. Thus, before creating Fast API module, you should make sure that you already have [Fast API Service](creating-fast-api-service.md).

To create Fast API module, you can perform: `zaruba please makeFastApiModule -i` and follow the on screen instruction.

## Example

```sh
# run interactively
zaruba please makeFastApiModule -i

# run with parameters
zaruba please makeFastApiModule generator.fastApi.service.name=myService generator.fastApi.module.name=myModule
```

## Involved tasks

* [makeFastApiModule](tasks/makeFastApiModule.md)

## What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API route](creating-fast-api-route.md)
* [Creating Fast API event handler](creating-fast-api-event-handler.md)
* [Creating Fast API RPC handler](creating-fast-api-rpc-handler.md)
* [Creating Fast API Crud](creating-fast-api-crud.md)