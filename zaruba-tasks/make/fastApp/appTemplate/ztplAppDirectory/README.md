# ZtplAppDirectory

`ZtplAppDirectory` is a microservice-ready monolith application. It is built on top of [FastAPI](https://fastapi.tiangolo.com/), a modern, fast (high-performance) web framework for building APIs with Python 3.7+ based on standard Python-type hints.

`ZtplAppDirectory` is created by using [Zaruba](https://github.com/state-alchemists/zaruba).

To generate `ZtplAppDirectory`, you can use the following command:

```bash
zaruba please initProject # or cd to your zaruba project
zaruba please addFastApp appDirectory=ztplAppDirectory
```

Zaruba also allows you to add custom resources (like modules or API endpoints) into `ZtplAppDirectory`. For this to work, you need to follow a few conventions:

- Don't delete any line __preceded__ by this comment:

    `Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching`

- Don't delete any line __containing__ this comment:

    `Note: 💀 Don't delete this line; Zaruba uses it for pattern matching `

By using Zaruba, you can focus on features and business processes.

# 🏁 How to start

You can run `ztplAppDirectory` by invoking:

```bash
zaruba please startZtplAppDirectory
```

Alternatively, you can also invoke the following script:

```bash
cd ztplAppDirectory

# create virtual environment if not exist
if [ ! -d ./venv ]; then python -m venv ./venv; fi

# activate virtual environment
source venv/bin/activate

# install pip packages
pip install -r requirements.txt

# load environments
source template.env

# run the application
./start.sh
```

Once running, you can access `ZtplAppDirectory` by pointing your browser to [http://localhost:3000](http://localhost:3000).

You can log in to the system by using the following credentials:

- user: `root`
- password: `Alch3mist`

Make sure to change the `root` password by accessing [http://localhost:3000/auth/users](http://localhost:3000/auth/users).

# 🧪 How to test

Unit tests ensure that your software components work as intended. Thus, it is recommended to run the test every time you modify your code.

To run the `ztplAppDirectory` test, you can invoke the following command:

```bash
zaruba please testZtplAppDirectory
```

Alternatively, you can also run the following script:

```bash
cd ztplAppDirectory

# create virtual environment if not exist
if [ ! -d ./venv ]; then python -m venv ./venv; fi

# activate virtual environment
source venv/bin/activate

# install pip packages
pip install -r requirements.txt

# load environments
source template.env

# run pytest
pytest -rP -v --cov="$(pwd)" --cov-report html
```

# 🧩 How to add a new module

A module is a collection of code to handle a specific business domain.

Your application might consist of several modules that are independent of each other.

To add a new module, you can invoke the following command:

```bash
zaruba please addFastAppModule \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName
# or:
# zaruba please addFastAppModule -i
```

You can activate/deactivate a module by setting `ENABLE_YOUR_MODULE_NAME` to `1` or `0`.

For more detailed information, please visit [adding a new module section](_docs/adding-a-new-module/README.md)


# 📋 How to add a CRUD handler

CRUD (Create Read Update Delete) is common business logic.

A single CRUD handler contains several:

- API Route handlers
- UI Route handlers
- RPC handlers
- Service
- Repository

You can add a CRUD handler to your module by invoking the following command:

```bash
zaruba please addFastAppCrud \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appCrudEntity=yourCrudEntity \
    appCrudFields='["yourFirstField", "yourSecondField", "yourThirdField"]'
# or:
# zaruba please addFastAppCrud -i
```

Please note that Zaruba automatically adds the following fields:

- `id`
- `created_at`
- `created_by`
- `updated_at`
- `updated_by`

So, you no longer have to specify those fields.

For more detailed information, please visit [adding a CRUD handler section](_docs/adding-a-new-module/adding-a-crud-handler.md)


# 🍰 How to add a new column

Once you have a CRUD handler, you can add a new column by invoking the following command:

```bash
zaruba please addFastAppCrudField \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appCrudEntity=yourCrudEntity \
    appCrudField=yourNewColumn
# or
# zaruba please addFastAppCrudField -i
```

For example, you already have a `book` CRUD handler inside a `library` module, and you want to add a `synopsis` column:

```bash
zaruba please addFastAppCrudField \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudField=synopsis
```

For more detailed information, please visit [adding a new column section](_docs/adding-a-new-module/adding-a-new-column.md)


# 📄 How to add a page

A page is a human-readable response when a user hit a certain URL.

You can add a new page by invoking the following command:

```bash
zaruba please addFastAppPage \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appHttpMethod=get \
    appUrl=/your-end-point
# or:
# zaruba please addFastAppPage -i
```

For more detailed information, please visit [adding a page section](_docs/adding-a-new-module/adding-a-page.md)


# 🔗 How to add an API endpoint

API endpoint handles requests from UI pages or other systems.

You can add a new API endpoint by invoking the following command:

```bash
zaruba please addFastAppRouteHandler \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appHttpMethod=get \
    appUrl=/api/v1/your-end-point
# or:
# zaruba please addFastAppRouteHandler -i
```

For more detailed information, please visit [adding an API endpoint section](_docs/adding-a-new-module/adding-an-api-endpoint.md)


# 📢 How to add an event handler

Your software might consist of multiple applications working together.

Those applications need to talk to each other. Whenever an application needs to notify others about something, it will fire an event (e.g., an order created).

Other applications then need to listen to the event and do appropriate action (e.g., send a bill to the customer).

To listen to the event, you need an event handler.

The following command will help you add an event handler:

```bash
zaruba please addFastAppEventHandler \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appEventName=yourEventName
# or:
# zaruba please addFastAppEventHandler -i
```

For more detailed information, please visit [adding an event handler section](_docs/adding-a-new-module/adding-an-event-handler.md)


# 🤙 How to add an RPC handler

Sometimes when your application needs to get a reply from other applications.

For example, an application needs to know whether a user is authenticated or not. Thus it performs a `remote procedure call` (aka RPC). 

The authorization service then handles the RPC and sends a reply response to the application.

You can add an RPC handler by invoking the following command:

```bash
zaruba please addFastAppRpcHandler \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appRpcName=yourRPCFunctionName
```

For more detailed information, please visit [adding an RPC handler section](_docs/adding-a-new-module/adding-an-rpc-handler.md)


# ☑️ Prerequisites

- Python 3.9
- Pip
- Virtual env
- (Optional) [Zaruba](https://github.com/state-alchemists/zaruba)

# 📖 Documentation

Please visit [ZtplAppDirectory documentation](_docs/README.md) for more information.