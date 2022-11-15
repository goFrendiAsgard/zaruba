<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

# _docs directory

Containing the documentation of `ztplAppDirectory`. 

# alembic directory

Containing the database migration. We use [alembic](https://alembic.sqlalchemy.org/en/latest/) for database migration.

To generate migration file, you should import your SQLALchemy model into `alembic/env.py`:

```python
# file: alembic/env.py
from modules.library.book.repos.dbBookRepo import DBBookRepo
from modules.auth import DBRoleRepo
from modules.auth import DBUserRepo
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

# avro directory

Containing Avro schema for `KafkaAvro` message bus. You can read more about avro schema in [this article](https://www.tutorialspoint.com/avro/avro_schemas.htm).

# configs directory

Containing configurations for `ZtplAppDirectory`.

There are several important files in this directory:

- `configs/appFactory.py`: Containing logics to create FastApi application. The application also handle readiness route, error handling, etc.
- `configs/pageTemplateFactory.py`: Containing page (Jinja templates) configuration.
- `configs/messagebusFactory.py` and `configs/rpcFactory.py`: Containing a function to create messagebus/RPC.
- `configs/url.py`: Contaiing URL/path settings.
- `configs/featureFlag.py`: Containing some feature flags.

Any values from the configuration should be imported into `main.py` instead of being imported directly in the module, e.g.,

```python
from configs import (
    # feature flags
    enable_api, enable_auth_module, enable_event_handler, enable_route_handler, enable_rpc_handler,
    enable_ui, seed_root_user,
    # factories
    create_app, create_menu_service, create_message_bus, create_rpc, create_page_template,
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

# core directory

Containing core components/services.

Some important components are:

- `authService`, `authRule`, and `userFetcher` (located inside `core/security`)
- `sessionService` (located inside `core/session`)
- `menuService` (located inside `core/menu`)
- `tokenService` (located inside `core/token`)

# helpers directory

Containing some functions/factories that will be used in your application.

Unlike configs, helpers should be stateless and has less dependencies. Usually containing function or class declaration.

Some important components are:

- `RPC`, `LocalRPC`, `RMQRPC`
- `MessageBus`, `LocalMessageBus`, `RMQMessageBus`, `KafkaMessageBus`, `KafkaAvroMessageBus`

# modules directory

Containing modules, including `auth` module and any other custom modules.

# pages directory

Containing Jinja templates.

# public directory

Containing public assets (css/js/images, etc).

# repos directory

Containing base repo declaration.

# schemas directory

Containing schemas/DTO.


# transport directory

Containing definition of `AppMessageBus` and `AppRPC`.