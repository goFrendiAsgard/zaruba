<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

# `_docs` directory

Containing the documentation of `ztplAppDirectory`. 

# `alembic` directory

Containing the database migration. We use [alembic](https://alembic.sqlalchemy.org/en/latest/) for database migration.

To generate migration file, you should import your SQLALchemy model into `alembic/env.py`:

```python
# file: alembic/env.py
from module.library.book.repo.dbBookRepo import DBBookRepo
from module.auth import DBRoleRepo
from module.auth import DBUserRepo
```

Once the model has been imported, you can run the following commands:

```bash
export IS_GENERATING_MIGRATION=1
export MIGRATION_SQLALCHEMY_DATABASE_URL=sqlite://migration.db
export MIGRATION_RUN_ALL=1
alembic upgrade head
alembic revision --autogenerate -m '<your-migration-name>'
```

or, if you use Zaruba, you can do this intead:

```bash
zaruba please createZtplAppNameMigration
```

# `avro` directory

Containing Avro schema for `KafkaAvro` message bus. You can read more about avro schema in [this article](https://www.tutorialspoint.com/avro/avro_schemas.htm).

# `config` directory

Containing configurations for `ZtplAppDirectory`.

There are several important files in this directory:

- `configs/app_factory.py`: Containing logics to create FastApi application. The application also handle readiness route, error handling, etc.
- `configs/page_template_factory.py`: Containing page (Jinja templates) configuration.
- `configs/messagebus_factory.py` and `configs/rpc_factory.py`: Containing a function to create `AppMessagebus`/`AppRpc`.
- `configs/url.py`: Containing URL/path settings.
- `configs/feature_flag.py`: Containing some feature flags.

Any values from the configuration should be imported into `main.py` instead of being imported directly in the module, e.g.,

```python
from config import (
    # feature flags
    enable_api, enable_auth_module, enable_event_handler, enable_route_handler, enable_rpc_handler,
    enable_ui, seed_root_user,
    # factories
    create_app, create_message_bus, create_rpc, create_page_template,
    # db
    db_create_all, db_url,
    # messagebus + rpc
    message_bus_type, rpc_type,
    # url
    create_access_token_url_path, create_oauth_access_token_url_path, renew_access_token_url_path,
    # auth
    root_initial_email, root_initial_fullname, root_initial_password, 
    root_initial_phone_number, root_username, root_permission, access_token_algorithm,
    access_token_expire, access_token_secret_key,
    # activity
    activity_events
)
```

Feel free to add your own configuration to this directory.

# `core` directory

Containing core components/services.

Some important components are:

- `auth_service`, `auth_rule`, and `user_fetcher` (located inside `core/security`)
- `session_service` (located inside `core/session`)
- `menu_service` (located inside `core/menu`)
- `token_service` (located inside `core/token`)

# `helper` directory

Containing some functions/factories that will be used in your application.

Unlike configs, helpers should be stateless and has less dependencies. Usually containing function or class declaration.

Some important components are:

- `RPC`, `LocalRPC`, `RMQRPC`
- `MessageBus`, `LocalMessageBus`, `RMQMessageBus`, `KafkaMessageBus`, `KafkaAvroMessageBus`

# `module` directory

Containing modules, including `auth` module, `cms` module, and any other custom modules.

Typically, a module contains the following files/directories:

## `event.py`

Module event handler. Containing definition of the following functions:

- `register_<module>_event_handler`

## `route.py`

Module route handler. Containing definition of the following functions:

- `register_<module>_api_route`
- `register_<module>_ui_route`

## `rpc.py`

Module RPC handler. Containing definition of the following functions:

- `register_<module>_rpc_handler`


## `<entity>/<entity>_route.py`
## `<entity>/<entity>_rpc.py`
## `<entity>/<entity>_service.py`
## `<entity>/repo/<entity>_repo.py`
## `<entity>/repo/db_<entity>_repo.py`

# `pages` directory

Containing Jinja templates for serving UI

# `public` directory

Containing public resources (css/js/images, etc).

# `repo` directory

Containing base repo declaration.

# `schema` directory

Containing schemas/DTO.

# `transport` directory

Containing definition of `AppMessageBus` and `AppRPC`.

Feel free to extend this to match your business processes.