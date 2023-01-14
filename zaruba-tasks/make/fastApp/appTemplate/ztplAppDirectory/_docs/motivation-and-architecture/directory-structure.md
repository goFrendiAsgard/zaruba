<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

# `_docs` directory

Containing the documentation of `ztplAppDirectory`. 

# `_alembic` directory

This directory contains database migration. We use [alembic](https://alembic.sqlalchemy.org/en/latest/) for database migration.

To generate migration file, you should import your SQLALchemy model into `_alembic/env.py`:

```python
# file: _alembic/env.py
from module.library.book.repo.dbBookRepo import DBBookRepo
from module.auth import DBRoleRepo
from module.auth import DBUserRepo
```

Once the model has been imported, you can run the following commands:

```bash
export MIGRATION_SQLALCHEMY_DATABASE_URL=sqlite://migration.db
./create-migration.sh
```

or, if you use Zaruba, you can do this intead:

```bash
zaruba please createZtplAppNameMigration
```

Some important files/directories in `alembic` directory are:

- `_alembic/env.py`
- `_alembic/script.py.mako`
- `_alembic/versions` directory


## `_alembic/env.py`

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



## `_alembic/script.py.mako`

This file is a [mako template](https://www.makotemplates.org/) to generate new migration versions in `alembic/versions` directory.

## `_alembic/versions` directory

This directory contains database migration versions.

There are a few things you need to consider:

- Every migration file contains `revision` and `down_revision`. If you need to delete/add migration version manually, please make sure that the `revision` and `down_version` are in sync with other versions.
- As defined in `alembic/script.py.mako`, there is a `run_migration` function that you can modify to accomodate feature flag.

```python

def upgrade() -> None:
    # migration code
    pass


def downgrade() -> None:
    # migration code
    pass
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
    create_cred_token_url_path, create_oauth_cred_token_url_path, renew_cred_token_url,
    # auth
    root_initial_email, root_initial_fullname, root_initial_password, 
    root_initial_phone_number, root_username, root_permission, cred_token_algorithm,
    cred_token_expire, cred_token_secret_key,
    # activity
    activity_events
)
```

Feel free to add your own configuration to this directory.

There are several important files in `config` directory:

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

Feel free to modify this file and add custom logics as neeeded.

```python
def create_app(mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates) -> FastAPI:
    app = FastAPI(title=site_name)
    # middlewares, urls, and error handling
    # > TODO: add your custom logics here
    return app
```

## `config/page_template_factory.py`

Containing definition of `create_page_template` function. This function is responsible for

- creating `Jinja2Templates` instance
- injecting some values to the template

Feel free to inject global values as needed.

```python
def create_page_template() -> Jinja2Templates:
    templates = Jinja2Templates(directory=page_dir)
    templates.env.globals['site_name'] = site_name
    templates.env.globals['tagline'] = tagline
    templates.env.globals['footer'] = footer
    templates.env.globals['backend_address'] = backend_address
    templates.env.globals['public_url_path'] = public_url_path
    templates.env.globals['renew_cred_token_url'] = renew_cred_token_url
    templates.env.globals['renew_cred_token_interval'] = renew_cred_token_interval
    templates.env.globals['vue'] = escape_template
    templates.env.globals['getenv'] = os.getenv
    return templates
```


## `config/messagebus_factory.py`

Containing definition of `create_message_bus` function. This function is responsible for creating messagebus.

By default you can choose to use `kafka`, `kafkaAvro`, `rmq`, or `local` messagebus based on `mb_type` parameter.


```python
def create_message_bus(mb_type: str, activity_events: List[str]) -> AppMessageBus:
    if mb_type == 'rmq':
        return AppMessageBus(RMQMessageBus(rmq_connection_parameters, rmq_event_map), activity_events)
    if mb_type == 'kafka':
        return AppMessageBus(KafkaMessageBus(kafka_connection_parameters, kafka_event_map), activity_events)
    if mb_type == 'kafkaAvro':
        return AppMessageBus(KafkaAvroMessageBus(kafka_avro_connection_parameters, kafka_avro_event_map), activity_events)
    return AppMessageBus(LocalMessageBus(), activity_events)
```

Feel free to add more messagebus implementations.

## `config/rpc_factory.py`

Containing definition of `create_rpc` function. This function is responsible for creating rpc.

By default you can choose to use `rmq` or `local` RPC based on `rpc_type` parameter.

```python
def create_rpc(rpc_type: str) -> AppRPC:
    if rpc_type == 'rmq':
        return AppRPC(RMQRPC(rmq_connection_parameters, rmq_event_map))
    return AppRPC(LocalRPC())
```

Feel free to add more RPC implementations.


## `config/url.py`

This file contains URL setting.

The settings are taken from environment variables as follows:

```python
create_oauth_cred_token_url_path: str = os.getenv('APP_CREATE_OAUTH_CRED_TOKEN_URL', '/api/v1/create-oauth-access-token/')
create_cred_token_url_path: str = os.getenv('APP_CREATE_CRED_TOKEN_URL', '/api/v1/create-access-token/')
renew_cred_token_url: str = os.getenv('APP_RENEW_CRED_TOKEN_URL', '/api/v1/refresh-access-token/')
public_url_path: str = os.getenv('APP_PUBLIC_URL_PATH', '/public')

backend_address: str = os.getenv('APP_BACKEND_ADDRESS', 'http://localhost:{}'.format(http_port))
```


## `config/feature_flag.py`

This file contains some feature flag definitions.

The settings are taken from environment variables as follows:

```python
enable_auth_module: bool = os.getenv('APP_ENABLE_AUTH_MODULE', '1') != '0'
enable_cms_module: bool = os.getenv('APP_ENABLE_CMS_MODULE', '1') != '0'
enable_log_module: bool = os.getenv('APP_ENABLE_LOG_MODULE', '1') != '0'
enable_route_handler: bool = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_ui: bool = os.getenv('APP_ENABLE_UI', '1') != '0'
enable_api: bool = os.getenv('APP_ENABLE_API', '1') != '0'
enable_error_page: bool = os.getenv('APP_ENABLE_ERROR_PAGE', '1') != '0'
enable_event_handler: bool = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler: bool = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'
```

# `core` directory

This file contains core components/services definitions.

Some services like `AuthService` and `MenuService` are used for authentication. Thus, they can be called from other modules.

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

This file contains definition of `MenuService`. This service is responsible for:

- add page menu
- provide menu structure
- authorize page

Most likely, you will not change any implementation of `MenuService`.

To use `MenuService`, you can add the following code into your `UI Route handlers`:

```python
def register_book_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # Telling the menuService to add a new menu named `library:books`.
    # This menu is child of `library` menu, and it should point to `/library/books` URL.
    # To access the menu, a user should have `ui:library:book` permissions.
    menu_service.add_menu(name='library:books', title='Books', url='/library/books', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:library:book', parent_name='library')

    # Handle the `/library/books` route.
    # Make sure only user who can access `library:books` menu can access the URL.
    @app.get('/library/books', response_class=HTMLResponse)
    async def manage_book(request: Request, context: MenuContext = Depends(menu_service.has_access('library:books'))):
        '''
        Serving user interface for managing book.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/library/crud/books.html',
            'request': request, 
            'context': context
        }, status_code=200)
```

Multiple `auth_type` options are available:

- `AuthType.ANYONE`: __Anyone__ can access the menu.
- `AuthType.VISITOR`: Only unauthenticated user (i.e., __has not logged in__) can access the menu.
- `AuthType.USER`: Only authenticated user (i.e., __logged in__ to the system) can access the menu.
- `AuthType.HAS_PERMISSION`: Only authenticated user who has `permission_name` can access the menu.

## `core/page/page_template_exception.py`

This file contains definition of `PageTemplateException`.

Most likely, you will not change any implementation of `PageTemplateException`.


## `core/security/middleware/user_fetcher.py`

This file contains `UserFetcher` interface.

```python
class UserFetcher(abc.ABC):

    @abc.abstractmethod
    def get_user_fetcher(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Return a function to fetch current user based on HTTP request.
        To be used with fastapi.Depends
        '''
        pass
```

`UserFetcher` interface is responsible to fetch user.
Together with `AuthRule`, you can use `UserFetcher` to create an `AuthService`.

## `core/security/middleware/default_user_fetcher.py`

This file contains the defailt mplementation of `UserFetcher`.

## `core/security/rule/auth_rule`

This file contains `AuthRule` interface.

```python
class AuthRule(abc.ABC):

    @abc.abstractmethod
    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        '''
        Return boolean, representing whether current_user should pass the authentication/authorization or not.
        '''
        pass
```

`AuthRule` interface is responsible to determine authentication rule.
Together with `UserFetcher`, you can use `AuthRule` to create an `AuthService`.

## `core/security/rule/default_auth_rule`

This file contains the definition of default `AuthRule`.

## `core/security/rule/no_auth_rule`

This file contains the definition of no-authentication `AuthRule`.

When using this rule, all user are treated as `authenticated` and `authorized`.

## `core/security/service/auth_service.py`

This file contains the definition of `AuthService`.

To create an `AuthService`, you need two dependencies:

- `UserFetcher`
- `AuthRule`

To create an `AuthService`, you can use the following code:

```python
oauth2_scheme = OAuth2PasswordBearer(tokenUrl = create_oauth_cred_token_url_path, auto_error = False)
auth_rule = DefaultAuthRule(rpc)
user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
auth_service = AuthService(auth_rule, user_fetcher, root_permission)
```

You can use `AuthService` to check API endpoint authorization as follow:

```python
def register_book_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/books/', response_model=BookResult)
    async def find_books(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:book:read'))) -> BookResult:
        '''
        Serving API to find books by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_book', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return BookResult.parse_obj(result)
```

`AuthService` has several methods that you can use to authorize your API endpoints:

- `auth_service.anyone()`: __Anyone__ should be able to access the endpoint
- `auth_service.is_visitor()`: Only they who __has not logged__ in can access the endpoint
- `auth_service.is_user()`: Only they who __has logged in__ can access the endpoint
- `auth_service.has_permission(permission_name: str)`: Only they who has logged in and has `permission_name` can access the endpoint.

## `core/sesion/session_route.py`

This file contains API endpoints and session pages:

The provided API endpoints are:

- Create access token (`/api/v1/create-access-token` and `/api/v1/create-oauth-access-token`)
- Renew access token (`api/v1/renew-access-token`)

The provided session pages are:

- Login (`/account/login`)
- Logout (`/account/logout`)

## `core/sesion/session_rpc.py`

This file contains RPC handler to handle session related requests.

## `core/sesion/session_service.py`

This file contains of `SessionService` definition.

To create a `SessionService`, you need a `TokenService` and a `UserService`. For example, check on the following code:

```python
user_repo = DBUserRepo(engine=engine, create_all=db_create_all)
role_service = RoleService(mb, rpc, role_repo)
user_service = DefaultUserService(mb, rpc, user_repo, role_service, root_permission=root_permission)
token_service = JWTTokenService(
    user_service = user_service,
    cred_token_secret_key = cred_token_secret_key,
    cred_token_algorithm = cred_token_algorithm,
    cred_token_expire = cred_token_expire
)
session_service = SessionService(user_service, token_service)
```

## `core/token/token_service.py`

This file contains of `TokenService` interface definition.

`TokenService` is responsible to generate/validate authentication token.

```python
class TokenService(abc.ABC):

    @abc.abstractmethod
    def create_cred_token(self, user: User, current_user: Optional[User]) -> str:
        pass

    @abc.abstractmethod
    def get_user_by_token(self, token: str, current_user: Optional[User]) -> Optional[User]:
        pass
```

## `core/token/jwt_token_service.py`

This file contains of `JWTTokenService` implementation. You can learn more about JWT Token from [wikipedia](https://en.wikipedia.org/wiki/JSON_Web_Token).

# `helper` directory

Containing some functions/class that will be used in your application.

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

You probably won't change any implementation of the helpers.

# `module` directory

This directory contains modules, including `auth`, `cms`, and many others.

Typically, a module contains the following files/directories:

- `module/<module-name>/event.py`
- `module/<module-name>/route.py`
- `module/<module-name>/rpc.py`

If your module handle CRUD or other business process, it might also contains `module/<module-name>/<entity-name>` directory.

## `module/<module-name>/event.py`

This file contains the definition of `register_<module_name>_event_handler` function.

Event handlers handle any event related to the module. For example, let's look at the following code:

```python
def register_library_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @mb.handle('new_visitor')
    def handle_new_visitor(data: Mapping[str, Any]):
        print('handle new visitor with data: {}'.format(data))

    print('Register library event handler', file=sys.stderr)
```

To trigger the event, you can use `Messagebus` instance as follow:

```python
mb = LocalMessagebus()
mb.publish('new_visitor', {'name': 'Luke Skywalker', 'address': 'Galaxy far far away'})
```

## `module/<module-name>/route.py`

This file contains the definition of `register_<module_name>_api_route_handler` and `register_<module_name>_ui_route_handler` function.

Let's look at the following example:

```python
# handle API endpoints
def register_library_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/books/', response_model=BookResult)
    async def find_books(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:book:read'))) -> BookResult:
        '''
        Serving API to find books by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_book', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return BookResult.parse_obj(result)


# serve pages
def register_library_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):
    menu_service.add_menu(name='library', title='Library', url='#', auth_type=AuthType.ANYONE)

    menu_service.add_menu(name='library:books', title='Books', url='/library/books', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:library:book', parent_name='library')

    @app.get('/library/books', response_class=HTMLResponse)
    async def manage_book(request: Request, context: MenuContext = Depends(menu_service.has_access('library:books'))):
        '''
        Serving user interface for managing book.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/library/crud/books.html',
            'request': request, 
            'context': context
        }, status_code=200)
```

## `module/<module-name>/rpc.py`

This file contains the definition of `register_<module_name>_rpc_handler`.

Let's look at the foolowing code:

```python
def register_library_rpc_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, book_service: BookService):

    @rpc.handle('isBookAvailable')
    def is_book_available(parameter: str) -> bool:
        print('handle RPC call isBookAvailable with parameter: {}'.format(parameter))
        return True
```

To call the RPC, you can use `RPC` instance as follow:

```python
rpc = LocalRPC()
is_available = rpc.call('is_book_available', {'title': 'Doraemon'})
print(is_available)
```


## `module/<module-name>/<entity-name>/<entity-name>_route.py`

This file contains API/UI route handler for your entity.

## `module/<module-name>/<entity-name>/<entity-name>_rpc.py`

This file contains RPC handler for your entity.

## `module/<module-name>/<entity-name>/<entity-name>_service.py`

This file contains service definition of your entity.

A service contains your business logic.

## `module/<module-name>/<entity-name>/repo/<entity-name>_repo.py`

This file contains repository interface of your entity.
## `module/<module-name>/<entity-name>/repo/db_<entity-name>_repo.py`

This file contains db repository implementation of your entity.

# `pages` directory

Containing Jinja templates for serving UI.

You can breakdown a page into several `partial` components. For example:

## default_page.html

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        {% include 'default-partials/meta.html' %}
        {% include 'default-partials/include-css.html' %}
        <link rel="icon" type="image/x-icon" href="{{ public_url_path}}/favicon/favicon.ico">
        {%if context.current_menu %}<title>{{ context.current_menu.title }}</title>{% endif %}
    </head>
    <body>
        {% include 'default-partials/navigation.html' %}
        {% include 'default-partials/include-js.html' %}

        <div class="container">
            {% include 'default-partials/header.html' %}
            {%if context.current_menu %}<h1>{{ context.current_menu.title }}</h1>{% endif %}
            {% if content_path is defined %}
                {% include content_path %}
            {% elif content is defined %}
                {{ content }}
            {% endif %}
            {% include 'default-partials/footer.html' %}
        </div>
    </body>
</html>
```

## default_partials/meta.html

```html
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
```

# `public` directory

This directory contains public resources (e.g., css/js/images, etc).

# `repo` directory

This directory contains `Declarative Base` definition.

```python
from sqlalchemy.ext.declarative import declarative_base
Base = declarative_base()
```

This declarative base is crucial to make a DB repository.

# `schema` directory

This directory contains all schemas/DTO definition.

Since the data might move from one module to another, it is important to put all schemas here.

# `transport` directory

This directory contains the definition of `AppMessageBus` and `AppRPC`.

Currently, we define `publish_activity` method in `AppMessageBus`, so that we can re-use the logic without changing the `MessgeBus` implementation itself.

```python
class AppMessageBus(MessageBus):
    '''
    MesssageBus with special methods to support app use case.
    Feel free to add methods as necessary
    '''

    def __init__(self, mb: MessageBus, activity_events: List[str] = []):
        self.mb = mb
        self.activity_events = activity_events


    def handle(self, event_name: str) -> Callable[..., Any]:
        return self.mb.handle(event_name)


    def publish(self, event_name: str, message: Any) -> Any:
        return self.mb.publish(event_name, message)


    def shutdown(self) -> Any:
        return self.mb.shutdown()


    def get_error_count(self) -> int:
        return self.mb.get_error_count()


    def is_failing(self) -> bool:
        return self.mb.is_failing()


    def broadcast(self, event_names: List[str], message: Any):
        for event_name in event_names:
            self.mb.publish(event_name, message)


    def publish_activity(self, activity_data: ActivityData):
        self.mb.publish('new_activity', activity_data.dict())
        self.broadcast(self.activity_events, activity_data.dict())
```

Feel free to extend this to match your business processes.

<!--startTocSubtopic-->
<!--endTocSubtopic-->
