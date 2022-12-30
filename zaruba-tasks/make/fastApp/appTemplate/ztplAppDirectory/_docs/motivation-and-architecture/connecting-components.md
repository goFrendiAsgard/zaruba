<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Connecting components
<!--endTocHeader-->

You can connect every layer components to each other.

`ZtplAppDirectory` rely heavily on [dependency injection mechanism](https://en.wikipedia.org/wiki/Dependency_injection). With this mechanism, you can choose which layer component you want to use for particular use cases.

For example, when you deploy your application as a monolith, you want to use internal communication protocol. Thus you choose `LocalMessagebus` and `LocalRPC`.

However, when you deploy your application as microservices, you should use external communication protocol like Kafka or RabbitMq.

You can see how layers are connecting to each other in `main.py`.

In general, there are two ways to connect layer components:

- by passing the component as function parameter
- by passing the component to object's constructor parameter.

# Passing component as function parameter

Layers like `route handler`, `rpc handler`, and `event handler` are defined as functions. You can pass dependency componets into those layers as function parameters.

For example, you have `register_auth_event_handler` with the following definition:

```python
def register_auth_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    print('Register auth event handler', file=sys.stderr)
```

To call `register_auth_event_handler`, you need `AppMessageBus`, `AppRPC`, and `AuthService`. In that case, you can pass those components as function parameters:

```python
mb = LocalMessageBus()
rpc = LocalRPC()
auth_service = AuthService()
register_auth_event_handler(mb, rpc, auth_service)
```

# Passing component as object constructor parameter

Layers like `service` and `repo` are defined as objects. You can pass dependency components into those layers as object constructor parameter.

For example, you have a `SessionService` with the following definition:

```python
class SessionService():

    def __init__(self, user_service: UserService, token_service: TokenService) ->
        self.user_service = user_service
        self.token_service = token_service
    
    # and other detail implementation
```

To create a `SessionService`, you need `user_service` and `token_service`. Thus, you can pass those components as `SessionService`'s constructor parameter.

```python
user_service = UserService()
token_service = TokenService()
session_service = SessionService(user_service, token_service)
```

# Next

That was the basic mechanism of `ZtplAppDirectory`. You might want to check about [module](../creating-new-module/README.md), [authentication/authorization](../authentication-authorization.md), or [user interface](../user-interface/README.md).

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->