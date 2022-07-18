<!--startTocHeader-->
[🏠](README.md)
# Authentication + Authorization
<!--endTocHeader-->

Authentication is about validating who the user is, while authorization is about validating what a user can do.

When you login into the system, the system authenticate your identity. Verifying that you are a registered user.

Once you are authenticated, the system then determine what you can/cannot do. For example, a user with valid permission should be able to insert new data.

In `ZtplAppDirectory`, authentication and authorization are tightly related to `users`, `roles`, and `permissions`.

# Users, roles, and permissions

Every authenticated user is associated with a `user` entity.

A user might have multiple `permissions` and `roles`.

A `role` is a collection of `permissions`. So, instead of assigning the same set of permissions to multiple users, you can create a single role and assign that role to the users.

User's role and permissions are stored as JSON array.

# Authorize API end point

To authorize API end point, you can use any implementation of `AuthService` interface. You can define your own implementation of `AuthService`. But, for start, `ZtplAppDirectory` already has `TokenOAuth2AuthService` and `NoAuthService`.

By default, `ZtplAppDirectory` use `TokenOAuth2AuthService`. You can find the corresponding definition at `main.py`:

```python
# location: main.py
auth_service = TokenOAuth2AuthService(rpc, oauth2_scheme)
```

`AuthService` has many way to define authorization of API endpoints:

- `auth_service.everyone(throw_error: bool = True)`. Everyone can access the API.
- `auth_service.is_authenticated(throw_error: bool = True)`. Only authenticated users (user that has already login) can access the API.
- `auth_service.is_unauthenticated(throw_error: bool = True)`. Only is_unauthenticated users (user that has not been login) can access the API.
- `auth_service.is_authorized(permission: str, throw_error: bool = True)`. Only authorized users (user that has already login and has permission) can access the page.

Those methods will:
- return current user if the authorization success 
- throw an error if the authorization failed and `throw_error` is set to `True`.
- return None if the authorization failed and `throw_error` is set to `False`.

Let's see the following example:

```python
# location: 
# - <module>/route.py
# - <module>/<entity>Route.py
@app.get('/api/v1/books/', response_model=BookResult)
def find_books(keyword: str='', limit: int=100, offset: int=0, current_user:  User = Depends(auth_service.is_authorized('api:book:read'))) -> BookResult:
    result = {}
    try:
        result = rpc.call('find_book', keyword, limit, offset)
    except:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
    return BookResult.parse_obj(result)
```

In this case, you want `GET /api/vi/books` to be accessible by any user that has `api:book:read` permission. Thus, you define the dependency as handler parameter: `current_user:  User = Depends(auth_service.is_authorized('api:book:read'))`

# Authorize page

To authorize a page, you can use any implementation of `MenuService` interface. `ZtplAppDirectory` provide you with default implementation that you can modify later:

```python
# location: main.py
menu_service = create_menu_service(rpc, auth_service)
```

Inside `create_menu_service` function you can see how the menu is being aranged:

```python
# location: helpers/app/createMenuService.py
def create_menu_service(rpc: RPC, auth_service: AuthService) -> MenuService:
    menu_service = DefaultMenuService(rpc, auth_service)
    menu_service.add_menu(name='account', title='Account', url='#', auth_type=AuthType.EVERYONE)
    menu_service.add_menu(name='account/login', title='Log in', url='/account/login', auth_type=AuthType.UNAUTHENTICATED, parent_name='account')
    menu_service.add_menu(name='account/logout', title='Log out', url='/account/logout', auth_type=AuthType.AUTHENTICATED, parent_name='account')
    menu_service.add_menu(name='auth', title='Security', url='#', auth_type=AuthType.EVERYONE)
    menu_service.add_menu(name='auth/roles', title='Roles', url='/auth/roles', auth_type=AuthType.AUTHORIZED, permission_name='ui:auth:role', parent_name='auth')
    menu_service.add_menu(name='auth/users', title='Users', url='/auth/users', auth_type=AuthType.AUTHORIZED, permission_name='ui:auth:user', parent_name='auth')
    return menu_service
```

First, it tell you that it will use `DefaultMenuService`.

It also tell you about the menu structure:

```
account (Account)
  - account/login (Log in)
  - account/logout (Log out)
auth (Security)
  - auth/roles (Roles)
  - auth/users (Users)
```

Be sure to register your page to `create_menu_service` whenever you add a new page.

To authorize a page, you can inject `menu_service.authenticate('<menu-name>')`:

```python
@app.get('/', response_class=HTMLResponse)
async def get_(request: Request, context: MenuContext = Depends(menu_service.authenticate('library:/'))) -> HTMLResponse:
    '''
    Handle (get) /
    '''
    try:
        return templates.TemplateResponse('default_page.html', context={
            'request': request,
            'context': context,
            'content_path': 'library/.html'
        }, status_code=200)
    except:
        print(traceback.format_exc()) 
        return templates.TemplateResponse('default_error.html', context={
            'request': request,
            'status_code': 500,
            'detail': 'Internal server error'
        }, status_code=500)
```

`menu_service.authenticate` will return `MenuContext` that you can use to render jinja templates.

<!--startTocSubTopic-->
<!--endTocSubTopic-->