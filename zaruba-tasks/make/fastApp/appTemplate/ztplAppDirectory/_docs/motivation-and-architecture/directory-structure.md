<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

# _docs directory

Containing the documentation of `ztplAppDirectory`. 

# alembic directory

Containing the database migration. We use [alembic](https://alembic.sqlalchemy.org/en/latest/) for database migration.

You can generate migration file by importing your SQLALchemy model into `alembic/env.py`.

For example:

```python
from modules.library.book.repos.dbBookRepo import DBBookRepo
from modules.auth import DBRoleRepo
from modules.auth import DBUserRepo
```

To generate the migration, you can perform:

```bash
export IS_GENERATING_MIGRATION=1
export MIGRATION_SQLALCHEMY_DATABASE_URL=sqlite://migration.db
export MIGRATION_RUN_ALL=1
alembic upgrade head
alembic revision --autogenerate -m '<your-migration-name>'
```

or, if you use Zaruba:

```bash
zaruba please createZtplAppNameMigration
```

# avro directory

Containing Avro schema for `KafkaAvro` message bus.

# configs directory

Containing configurations for `ZtplAppDirectory`.

There are several important files in this directory:

- `configs/appFactory.py`: Containing logics to create FastApi application. The application also handle readiness route, error handling, etc.
- `configs/menuServiceFactory.py`: Containing menu structure of `ZtplAppDirectory`.
- `configs/pageTemplateFactory.py`: Containing page (Jinja templates) configuration.
- `configs/messagebusFactory.py` and `configs/rpcFactory.py`: Containing a function to create messagebus/RPC.
- `configs/url.py`: Contaiing URL/path settings.
- `configs/featureFlag.py`: Containing some feature flags.

Any values from the configuration should be imported into `main.py` instead of being imported directly in the module.

# core directory

Containing core components/services.

# helpers directory

Containing some functions/factories that will be used in your application.

Unlike configs, helpers should be stateless and has less dependencies. Usually containing function or class declaration.

# modules directory

Containing modules, including `auth` and any other custom modules.

# pages directory

Containing Jinja templates.

# public directory

Containing public assets (css/js/images, etc).

# repos directory

Containing base repo declaration.

# schemas directory

Containing schemas/DTO.