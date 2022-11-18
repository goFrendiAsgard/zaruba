<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

# `_docs` directory

Containing the documentation of `ztplAppDirectory`. 

# `alembic` directory

This directory contains database migration. We use [alembic](https://alembic.sqlalchemy.org/en/latest/) for database migration.

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

Some important files/directories in `alembic` directory are:

- `alembic/env.py`
- `alembic/script.py.mako`
- `alembic/versions` directory


## `alembic/env.py`

Alembic use this file to generate/run migration.

Whenever you create a new [SQLALchemy model](https://docs.sqlalchemy.org/en/14/orm/quickstart.html#declare-models), you need to import the model into this file.

```python
# file: alembic/env.py
from repo import Base
from module.auth import DBRoleEntity, DBUserEntity
from module.log.activity import DBActivityEntity
from module.cms.content import DBContentEntity, DBContentAttributeEntity
from module.cms.content_type import DBContentTypeEntity
```

For more information about `env.py`, please refer to [alembic documentation](https://alembic.sqlalchemy.org/en/latest/tutorial.html#the-migration-environment).



## `alembic/script.py.mako`

This file is a [mako template](https://www.makotemplates.org/) to generate new migration versions in `alembic/versions` directory.

`ZtplAppDirectory` add `run_migration` function to accomodate feature flag.

```python
def run_migration() -> bool:
    return os.getenv('MIGRATION_RUN_ALL', '0') != '0'


def upgrade() -> None:
    if not run_migration():
        return None
    ${upgrades if upgrades else "pass"}


def downgrade() -> None:
    if not run_migration():
        return None
    ${downgrades if downgrades else "pass"}
```

When you generate a migration, Alembic will first run existing migrations in migration database. Once the migration has been done, it will compare existing models with the migration database. If there is any discrepancy, Alembic will make a new migration version to fill the gap.

Thus, when generating a migration, you need to set `MIGRATION_RUN_ALL` to `1`.


## `alembic/versions` directory

This directory contains database migration versions.

There are a few things you need to consider:

- Every migration file contains `revision` and `down_revision`. If you need to delete/add migration version manually, please make sure that the `revision` and `down_version` are in sync with other versions.
- As defined in `alembic/script.py.mako`, there is a `run_migration` version that you can modify to accomodate feature flag.

To accomodate feature flag, you can add `or` logic to manipulate `run_migration` return value. For example, you want the migration version to be ignored when `APP_ENABLE_LOG_MODULE` is set to `0`:

```python
def run_migration() -> bool:
    return os.getenv('MIGRATION_RUN_ALL', '0') != '0' or os.getenv('APP_ENABLE_LOG_MODULE', '1') != '0'


def upgrade() -> None:
    if not run_migration():
        return None
    # migration code


def downgrade() -> None:
    if not run_migration():
        return None
    # migration code
```


# `avro` directory

Containing Avro schema for `KafkaAvro` message bus. You can read more about avro schema in [this article](https://www.tutorialspoint.com/avro/avro_schemas.htm).

# `config` directory

Containing configurations for `ZtplAppDirectory`.

Any values from the configuration should be imported into `main.py` instead of being imported directly in the module, e.g.,

```python
# file location: main.py
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


There are several important files in this directory:

- `config/app_factory.py`
- `config/page_template_factory.py`
- `config/messagebus_factory.py`
- `configs/rpc_factory.py`
- `config/url.py`
- `config/feature_flag.py`


## `config/app_factory.py`

Containing definition of `create_app` function. This function is responsible for:

- creating `FastAPI` instance
- apply middlewares (e.g., CORSMiddleware)
- handle `readiness` URL
- handle `shutdown` event
- serve `public` directory
- serve home page
- handle UI for error


## `config/page_template_factory.py`

Containing definition of `create_page_template` function. This function is responsible for

- creating `Jinja2Templates` instance
- injecting some values to the template


## `config/messagebus_factory.py`

Containing definition of `create_message_bus` function. This function is responsible for creating messagebus based on `mb_type` parameter.


## `config/rpc_factory.py`

Containing definition of `create_rpc` function. This function is responsible for creating rpc based on `rpc_type` parameter.


## `config/url.py`

This file contains URL setting.


## `config/feature_flag.py`

This file contains some feature flag definitions.


# `core` directory

Containing core components/services.

Some important directories are:

- `core/menu/menu_service.py`
- `core/page/page_template_exception.py`
- `core/security/middleware/user_fetcher.py`
- `core/security/middleware/default_user_fetcher.py`
- `core/security/rule/auth_rule`
- `core/security/rule/default_auth_rule`
- `core/security/rule/no_auth_rule`
- `core/security/service/auth_service.py`
- `core/sesion/session_route.py`
- `core/sesion/session_rpc.py`
- `core/sesion/session_service.py`
- `core/token/token_service.py`
- `core/token/jwt_token_service.py`

## `core/menu/menu_service.py`

Containing definition of `MenuService`. This service is responsible for:

- add page menu
- provide menu structure
- authorize page

## `core/page/page_template_exception.py`

Containing definition of `PageTemplateException`. This exception should be used in UI.


## `core/security/middleware/user_fetcher.py`
## `core/security/middleware/default_user_fetcher.py`
## `core/security/rule/auth_rule`
## `core/security/rule/default_auth_rule`
## `core/security/rule/no_auth_rule`

## `core/security/service/auth_service.py`

## `core/sesion/session_route.py`
## `core/sesion/session_rpc.py`
## `core/sesion/session_service.py`
## `core/token/token_service.py`
## `core/token/jwt_token_service.py`



# `helper` directory

Containing some functions/factories that will be used in your application.

Unlike configs, helpers should be stateless and has less dependencies. Usually containing function or class declaration.

Some important files are

- `helper/transport/messagebus.py`
- `helper/transport/local_messagebus.py`
- `helper/transport/kafka_messagebus.py`
- `helper/transport/kafka_avro_messagebus.py`
- `helper/transport/rmq_messagebus.py`
- `helper/transport/rpc.py`
- `helper/transport/local_rpc.py`
- `helper/transport/rmq_rpc.py`

# `module` directory

Containing modules, including `auth` module, `cms` module, and any other custom modules.

Typically, a module contains the following files/directories:

## `module/<module-name>/event.py`

Module event handler. Containing definition of the following functions:

- `register_<module-name>_event_handler`

## `module/<module-name>/route.py`

Module route handler. Containing definition of the following functions:

- `register_<module>_api_route`
- `register_<module>_ui_route`

## `module/<module-name>/rpc.py`

Module RPC handler. Containing definition of the following functions:

- `register_<module>_rpc_handler`


## `module/<module-name>/<entity-name>/<entity-name>_route.py`
## `module/<module-name>/<entity-name>/<entity-name>_rpc.py`
## `module/<module-name>/<entity-name>/<entity-name>_service.py`
## `module/<module-name>/<entity-name>/repo/<entity-name>_repo.py`
## `module/<module-name>/<entity-name>/repo/db_<entity-name>_repo.py`

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