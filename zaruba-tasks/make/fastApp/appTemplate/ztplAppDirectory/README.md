# ZtplAppDirectory

`ZtplAppDirectory` is a microservice-ready monolith application. It is built on top of [FastAPI](https://fastapi.tiangolo.com/), a modern, fast (high-performance), web framework for building APIs with Python 3.7+ based on standard Python type hints

`ZtplAppDirectory` was generated with the following command using [Zaruba](https://github.com/state-alchemists/zaruba):

```bash
zaruba please initProject # or cd to your zaruba project
zaruba please addFastApp appDirectory=ztplAppDirectory
```

## 💀 Note

Zaruba uses pattern matching (aka regex) so that it can add new routes/modules/entities/fields to `ZtplAppDirectory`.

To make sure the feature works as intended, please don't delete/edit any line __preceeded__ by this comment:

```
Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
```

or any line __containing__ this comment:

```
Note: 💀 Don't delete this line, Zaruba use it for pattern matching 
```

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

# 🧪 How to test

You can run `ztplAppDirectory` unit-test by invoking:

```bash
zaruba please testZtplAppDirectory
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

# run pytest
pytest -rP -v --cov="$(pwd)" --cov-report html
```

# 🧩 How to add a new module

A module is a collection of code to handle a specific business domain.

To add a new module, you can invoke the following command:

```bash
zaruba please addFastAppModule \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName
# or:
# zaruba please addFastAppModule -i
```

You can activate/deactivate a module by setting `ENABLE_YOUR_MODULE_NAME` to `1` or `0`

# 📋 How to add a CRUD handler

CRUD (Create Read Update Delete) is common business logic.

A single CRUD handler contains several:

- API Route handlers
- Frontend Page
- RPC handlers
- Repository

You can add a CRUD handler by invoking the following command:

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

For example, you already have a `book` CRUD handler inside a `library` module, and you want to add `synopsis` column:

```bash
zaruba please addFastAppCrudField \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudField=synopsis
```

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

# 🔗 How to add an API endpoint

When other systems/apps need to talk to your application, you should provide an API endpoint.

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

## 📢 How to add an event handler

When you have multiple apps in your system, some of your apps might fire an event to be handled by other apps.

You can add an event handler by invoking the following command:

```bash
zaruba please addFastAppEventHandler \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appEventName=yourEventName
# or:
# zaruba please addFastAppEventHandler -i
```

# 🤙 How to add an RPC handler

When you have multiple apps in your system, they need to talk to each other and get appropriate responses. This is called RPC (Remote Procedure Call).

You can add an RPC handler by invoking the following command:

```bash
zaruba please addFastAppRpcHandler \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appRpcName=yourRPCFunctionName
```

# ☑️ Prerequisites

- Python 3.9
- Pip
- Virtual env
- (Optional) [Zaruba](https://github.com/state-alchemists/zaruba)

# 📖 Documentation

Please visit ZtplAppDirectory documentation [here](_docs/README.md).