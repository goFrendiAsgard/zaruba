<!--startTocHeader-->
[🏠](README.md)
# Authentication and Authorization
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

To authorize API end point, you can use any implementation of `AuthService` interface. You can define your own implementation of `AuthService`. But, for start, `ZtplAppDirectory` already has `TokenAuthService` and `NoAuthService`.

By default, `ZtplAppDirectory` use `TokenAuthService`. You can find the corresponding definition at `main.py`:

```python
# location: main.py
auth_service = TokenAuthService(rpc, oauth2_scheme)
```

`AuthService` has many way to define authorization of API endpoints:

- `auth_service.anyone(throw_error: bool = True)`. Everyone can access the API.
- `auth_service.is_user(throw_error: bool = True)`. Only authenticated users (user that has already login) can access the API.
- `auth_service.is_visitor(throw_error: bool = True)`. Only is_unauthenticated users (user that has not been login) can access the API.
- `auth_service.has_permission(permission: str, throw_error: bool = True)`. Only authorized users (user that has already login and has permission) can access the page.

Those methods will:
- return current user if the authorization success 
- throw an error if the authorization failed and `throw_error` is set to `True`.
- return None if the authorization failed and `throw_error` is set to `False`.

Let's see the following example:

```python
@app.get('/api/v1/books/', response_model=BookResult)
def find_books(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:book:read'))) -> BookResult:
    result = {}
    try:
        result = rpc.call('find_book', keyword, limit, offset)
    except:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
    return BookResult.parse_obj(result)
```

In this case, you want `GET /api/vi/books` to be accessible by any user that has `api:book:read` permission. Thus, you define the dependency as handler parameter: `current_user: Optional[User] = Depends(auth_service.has_permission('api:book:read'))`

# Authorize page

To authorize a page, you can use any implementation of `MenuService` interface. `ZtplAppDirectory` provide you with default implementation that you can modify later:

```python
menu_service = create_menu_service(rpc, auth_service)
```

Once you have a `MenuService`, you can start adding menu


```python
def register_book_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    menu_service.add_menu(name='library', title='Library', url='#', auth_type=AuthType.ANYONE)
    menu_service.add_menu(name='library:books', title='Books', url='/library/books', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:library:book', parent_name='library')
```

It also tell you about the menu structure:

```
- library (#)
    - library:books (/library/books)
```

Now to authorize a page, you can inject `menu_service.has_access('<menu-name>')`:

```python
def register_book_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

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

`menu_service.has_access` will return `MenuContext` that you can use to render jinja page template in `pages/modules/library/crud/books.html`.

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->