# Fast API Message Bus

Message bus/message broker is third party service that can help your services to communicate asynchronously. Some common message bus are kafka, nats, and rabbitmq. Sure, there are many others out there. Redis for example, can also be used as message broker.

Currently zaruba's Fast API service support two kind of message bus:

* rabbitmq
* local

## About local message bus

Local message bus is basically a simple map of event and function. So when you fire up an event, a corresponding function will be called immediately. Since local and rabbitmq message bus share the same interface, you can simply swap them as needed.

Local message bus will be useful when you just getting started. At the very first stage of your development, it is suggested that you go with monolith first. You can then split your monolith into microservices when it is really necessary. Thus, the best thing you can do is make your code to be as decoupled as possible.

## Dependency Injection

Dependency injection is a good solution to decouple your code. To put it simple, dependency injection allows you to inject components into your code instead of including components from your code.

Here is an ilustration to show how dependency injection work:

```python
class Pistol:
    def shoot(self):
        print('bang bang')

class Crossbow:
    def shoot(self):
        print('pew')


# a class without dependency injection
class Soldier1:
    def attack(self):
        weapon = Pistol()
        weapon.shoot()


# a class with constructor-based dependency injection
class Soldier2:
    def __init__(self, weapon):
        self.weapon = weapon
    
    def attack(self):
        self.weapon.shoot()
```

You can see that with `Soldier1`, you cannot change the weapon without modifying it's internal code. This make `Soldier1` entangled to `Pistol`.

Now when you see `Soldier2`, you can put anything with `shoot` method as it's weapon.

```python
# this is a soldier with gun
soldier_with_gun = Soldier2(Pistol())
soldier_with_gun.shoot()

# this is a soldier with crossbow
soldier_with_crossbow = Soldier2(Crossbow())
soldier_with_crossbow.shoot()
```

Furthermore you can even make something like this:

```python
def get_weapon(weapon_name):
    if weapon_name == 'pistol':
        return Pistol()
    if weapon_name == 'crossbow':
        return Crossbow()
    return SomeDefaultWeapon() # or throw error

# now you can change the weapon by changing the environment variable
weapon = get_weapon(os.getenv('WEAPON'))
soldier = Soldier2(weapon)
```

Pretty neat, isn't it? By using dependency injection you can make your code more configurable.

Let's say some time in the future you want to introduce a new type of weapon like laser gun or something, you can simply use it as long as the new weapon has `shoot` method.

Even if your new weapon doesn't have a `shoot` method, you can always make an adapter so that your soldier can use the weapon seemlessly.

Most programming language has `interface` to support dependency inject. But since Python doesn't support `interface`, you can use [Abstract base class](https://docs.python.org/3/library/abc.html) instead.

## Local messagebus + dependency injection

Let's take a look again at your Fast API service's `main.py`.

```python
from fastapi import FastAPI
from sqlalchemy import create_engine
from helpers.transport import MessageBus, RMQMessageBus, RMQEventMap, LocalMessageBus

import os

def create_message_bus(mb_type: str) -> MessageBus:
    if mb_type == 'rmq':
        rmq_host = os.getenv('MY_SERVICE_RABBITMQ_HOST', 'localhost')
        rmq_user = os.getenv('MY_SERVICE_RABBITMQ_USER', 'root')
        rmq_pass = os.getenv('MY_SERVICE_RABBITMQ_PASS', 'toor')
        rmq_vhost = os.getenv('MY_SERVICE_RABBITMQ_VHOST', '/')
        rmq_event_map = RMQEventMap({})
        return RMQMessageBus(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalMessageBus()

db_url = os.getenv('MY_SERVICE_SQLALCHEMY_DATABASE_URL', 'sqlite://')
mb_type = os.getenv('MY_SERVICE_MESSAGE_BUS_TYPE', 'local')
enable_route = os.getenv('MY_SERVICE_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event = os.getenv('MY_SERVICE_ENABLE_EVENT_HANDLER', '1') != '0'

engine = create_engine(db_url, echo=True)
app = FastAPI()
mb = create_message_bus(mb_type)
```

Notice that we have `create_message_bus` function that will return local or rabbitmq message bus depends on your environment variable. As long as you don't make your modules entangled to each other, separating your Fast API monolith into microservices should not require too much work.

## Trade off of dependency injection

Dependency injection solve a lot of cases, and it can be used everywhere. But as anything in tech, there is no silver bullet. Dependency injection might even brings more problem to your code, especially if you never need to swap the component.

Since dependency injection usually use `interface`, you should put extra effort to know which component is actually being used right now. Thus, if used brutally, dependency injection can make your code less traceable.