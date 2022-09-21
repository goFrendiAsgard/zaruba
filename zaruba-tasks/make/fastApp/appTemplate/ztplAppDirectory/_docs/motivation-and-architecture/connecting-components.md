<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Connecting components
<!--endTocHeader-->

Every layer should be able to connect to each other.

You can choose which layer component you want to use for particular use cases. In `ZtplAppDirectory`, you can see how layers are connecting to each other in `main.py`.

There are two ways to connect layer components:

- by passing the component as function parameter
- by passing the component to object's constructor parameter.

# Passing component as function parameter

Layers like `route handler`, `rpc handler`, and `event handler` are defined as functions. You can pass dependency componets into those layers as function parameters.

For example, to create and register `auth_route_handler`, you need `app`, `mb`, `rpc`, `auth_service`, `menu_service`, etc. In that case, you can pass those components as function parameters:

```python
register_auth_route_handler(app, mb, rpc, auth_service, menu_service, page_template, enable_ui, enable_api, create_oauth_acess_token_path, create_acess_token_path, renew_access_token_url_path)
```

# Passing component as object constructor parameter

Layers like `service` and `repo` are defined as objects. You can pass dependency components into those layers as object constructor parameter.

For example, to create `session_service`, you need `user_service` and `token_service`. Thus, you can pass those components as `SessionService`'s constructor parameter.

```python
session_service = SessionService(user_service, token_service)
```

# Next

That was the basic mechanism of `ZtplAppDirectory`. You might want to check about [module](../creating-new-module/README.md), [authentication/authorization](../authentication-authorization.md), or [user interface](../user-interface/README.md).

<!--startTocSubTopic-->
<!--endTocSubTopic-->