<!--startTocHeader-->
[🏠](../README.md)
# Adding a new module
<!--endTocHeader-->

You can use [Zaruba](https://github.com/state-alchemists/zaruba) to add a new module, or you can write a new module from scratch.

# Using Zaruba

To create a new module using Zaruba, you can invoke the following code:

```bash
zaruba please addFastAppModule \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName
# or:
# zaruba please addFastAppModule -i
```

# From scratch

You can add a new module by creating a new directory containing `__init__.py`.
A module usually consists of `API/UI route handler`, `event handler`, and `RPC handler`.

Let's take a look at the following shell commands:

```bash
mkdir module/library
touch module/library/__init__.py
touch module/library/event.py
touch module/library/route.py
touch module/library/rpc.py
```

## Add handlers

Now let's add a few handlers.


### Event handler

```python
# location: module/library/event.py
from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import traceback
import sys

def register_library_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):
    # > TODO: add event handler here
    print('Register library event handler', file=sys.stderr)
```

### Route handler

```python
# location: module/library/route.py
from typing import Mapping, List, Any, Optional
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_library_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):
    # > TODO: add API endpoint handler here
    print('Register library api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
def register_library_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):
    menu_service.add_menu(name='library', title='Library', url='#', auth_type=AuthType.ANYONE)
    # > TODO: add menu and page handler here
    print('Register library UI route handler', file=sys.stderr)
```


### RPC handler

```python
# location: module/library/rpc.py
from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import traceback
import sys

# Note: 🤖 Don't delete the following line, Zaruba use it for pattern matching
def register_library_rpc_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):
    # > TODO: add RPC handler here
    print('Register library RPC handler', file=sys.stderr)
```

## Register module

Once you have your module created, you need to add it to `main.py`.

First, you need to import all the handlers, then you need to initiate the module based on the feature flag.

To make things easier, you can modify your `library/__init__.py`

```python
# lcoation: module/library/__init__.py
from module.library.route import register_library_api_route, register_library_ui_route
from module.library.event import register_library_event_handler
from module.library.rpc import register_library_rpc_handler
```

Now you can import everything at once.

```python
# location: main.py
from module.library import (
    register_library_api_route, register_library_ui_route, register_library_event_handler, register_library_rpc_handler
)

# ... rest of the code
```

Finally, you initiate the handlers based on the feature flag:

```python
# location: main.py
from module.library import (
    register_library_api_route, register_library_ui_route, register_library_event_handler, register_library_rpc_handler
)

# ... rest of the code


################################################
# -- 🧩 Library module
################################################
enable_library_module = os.getenv('APP_ENABLE_LIBRARY_MODULE', '1') != '0'
if enable_library_module:
    # API route
    if enable_route_handler and enable_api:
        register_library_api_route(app, mb, rpc, auth_service)
    # UI route
    if enable_route_handler and enable_ui:
        register_library_ui_route(app, mb, rpc, menu_service, page_template)
    # handle event
    if enable_event_handler:
        register_library_event_handler(mb, rpc, auth_service)
    # serve RPC
    if enable_rpc_handler:
        register_library_rpc_handler(mb, rpc, auth_service)
```

Done.

You can set `APP_ENABLE_LIBRARY_MODULE` to `0` or `1` to disable/enable the module.

<!--startTocSubtopic-->
# Subtopics
- [Adding a CRUD handler](adding-a-crud-handler.md)
- [Adding a new column](adding-a-new-column.md)
- [Adding an API endpoint](adding-an-api-endpoint.md)
- [Adding a page](adding-a-page.md)
- [Adding an event handler](adding-an-event-handler.md)
- [Adding an RPC handler](adding-an-rpc-handler.md)
<!--endTocSubtopic-->