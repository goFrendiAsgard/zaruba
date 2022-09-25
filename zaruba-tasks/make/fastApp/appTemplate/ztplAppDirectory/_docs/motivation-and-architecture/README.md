<!--startTocHeader-->
[🏠](../README.md)
# Motivation and architecture
<!--endTocHeader-->

Building a system is like building a town. First, you start with some basic requirements. Over time, your requirements grow and so does your software.

A monolith architecture is your best option if you are getting started. Monolith is easy to maintain/develop. Even though it won't scale very well.

You can think about microservices when you need to scale some features independently.

It is always better to __start with a monolith__ and __refactor to microservices__ later.

Refactoring is not always easy. Sometimes it is even impossible. There are many stories of companies who failed to do so because their code was not modular/testable. Their growth is now hindered by their own technology.

Some companies invest too much in Kubernetes or any other expensive technologies. They built a great microservices architecture, yet failed to maintain/develop it properly.

So we need a middle ground. We need a monolith application that is ready to be deployed as microservices.

Let's see the following scenarios, and see what we can do in each phase:

- __You are starting a business.__ Your IT team probably consists of a CTO and a couple of software engineers. In that case, you start with a monolith application. Make sure your monolith is modular so that you can split it into microservices later. If you have database repositories in your modules, be sure to use them in the same modules. This will make your modules decoupled from other modules.

- __Your software handles a lot of requests and has to be always available.__ Now you might consider using better hardware. Or you can split your monolith application into microservices. As your monolith application is modular, this process should not be very painful. You might also need to consider deploying your application into a Kubernetes cluster. Your services will be able to talk to each other using `Messagebus` or `RPC call`.

- __You need different technologies/programming languages for your services.__ Python is a good general-purpose programming language. But it is not your best choice when you need to take concurrency into account. In this case, you can write some of your services in other programming languages. Make sure the new service will handle all necessary events and RPC calls.

`ZtplAppDirectory` is a microservices-ready monolith. It has `LocalMessageBus` and `LocalRPC` for inter-module communication. Later, when your modules turn into microservices, you can switch into `RMQMessageBus` or `RMQRPC`. You can achieve this by changing the configuration, without touching the code.

# Microservices vs Monolith

In 2016, [DHH](https://twitter.com/dhh) wrote an article titled [Majestic Monolith](https://m.signalvnoise.com/the-majestic-monolith/).

Since big tech companies use microservices architecture, people are curious about this. They start to adopt the architecture without understanding the drawbacks. In the article, DHH argued that not all companies need microservices architecture.

Let's see how microservices and a monolith are different from each other.

## Microservices: The good and the bad

![image of individual zords from mighty morphin power ranger](images/individual-zords.jpg)

Microservices architecture is good because:

- It is easy to scale up/down particular services.
- Users can still access the system even though some services are down.
- Services can be developed/deployed independently from each other.

Microservices architecture is bad because:

- There is a lot of network communication.
- Either you build a correct one, or just create a [distributed monolith](https://www.techtarget.com/searchapparchitecture/tip/The-distributed-monolith-What-it-is-and-how-to-escape-it).
- Deploying microservices can be challenging.

## Monolith: The good and the bad

![image of megazord from mighty morphin power ranger](images/megazord.jpg)

Monolith architecture is good because:

- It is easy to run, debug, or deploy.

Monolith architecture is bad because:

- People tend to cross the domain boundary because they can.
- When it is down, the entire system is down.
- When you need to scale up/down, everything is scaled up/down

# Microservices-ready monolith


Many companies failed to refactor their monolith application into microservices. Other companies are investing too much in Kubernetes or other expensive technologies. They are paying too much for something they didn't need.

We can improve the situation by creating a modular monolith application. This modular monolith should be deployable as microservices later.

You can think of the application as a super-soldier with much training. Alone, this super-soldier is already capable of many things. This super soldier can shoot, fight, drive a helicopter, and even hack a satellite.

![image of a single super soldier](images/super-soldier-monolith.jpg)

But when you do a bigger mission, you cannot rely on a single super-soldier. A team consisting of multiple soldiers with different specializations will do better.

So you deploy multiple super-soldiers and assign specific roles/equipment for each of them.

![image of multiple super soldiers](images/super-soldier-microservices.jpg)

## Microservices-ready monolith: The good and the bad

Microservices-ready monolith is good because:

- It can be deployed and executed easily as a monolith.
- It has a single codebase.
- It can be deployed as microservices using [feature flags](feature-flags.md).
- Everyone is aware of code changes.

It is bad because:

- Everyone is aware of any code changes (Code-level authorization is nearly impossible).
- Without proper conventions, people tend to cross domain/module boundaries.

> ⚠️ __Warning:__ Never cross domain boundary even if you can. Assume every module will live in different servers/pods and accessing different database server. Otherwise, your application will ends up as big spaghetti-code monolith that can't go anywhere.

# Example

## Run as monolith

Suppose you have two modules in `ZtplAppDirectory`:

- Library
- Auth

These module serves different domains.

Someday you will need to scale them independently. But for now, let's start with a monolith:

![Monolith mode](images/fastApp-monolith.png)

At the beginning, you want to enable every feature flags. You can do this by setting all feature-flags environments into `1`.

You also want to use `LocalMessageBus` and `LocalRPC`, so you set `APP_MESSAGE_BUS_TYPE` and `APP_RPC_TYPE` into `local`. By setting your messagebus and RPC into local, the intermodule communication can be performed internally without any network overhead. This is a good thing because [network is not reliable](https://particular.net/blog/the-network-is-reliable)

All good, now you can run everything locally.

## Run as microservices

Now you want to run `ZtplAppDirectory` as microservices. You don't need to modify the source code at all. Instead, you just need to deploy the application with different feature flags (configurations).

![Microservices mode](images/fastApp-microservices.png)

In this example, we want to have:

- __Frontend.__ This service only serve the UI. It won't serve any API request, but it handle the UI of both `auth` and `library` module.
- __Backend.__ This service only serve API request from `Frontend` or other external applications, and pass them into respective services.
- __Auth Service.__ This service handle any event/RPC call for `auth` module. It also handle every database operation related to `auth` module.
- __Library Service.__ This service handle any event/RPC call for `library` module. It also handle every database operation related to `library` module.

# Next

The idea of a microservices-ready monolith is only possible because of [feature flags](feature-flags.md) and [layered architecture](interface-and-layers.md). You can learn more about the technical details in the sub-topics.

If you want to see how we organize our code, you can visit [directory structure](directory-structure.md).

<!--startTocSubTopic-->
# Sub-topics
* [Directory structure](directory-structure.md)
* [Feature flags](feature-flags.md)
* [Interface and layers](interface-and-layers.md)
* [Connecting components](connecting-components.md)
<!--endTocSubTopic-->