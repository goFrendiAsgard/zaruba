<!--startTocHeader-->
[🏠](../README.md)
# Motivation and architecture
<!--endTocHeader-->

Building software is like building a city. You __cannot plan everything right from the start__.

When you build software, you usually __start with some basic requirements__. But your needs will grow over time, and so will your software.

<p align="center">
    <img src="./images/night-city.jpg" width="500px" alt="Night city from cyberpunk 2077" />
    <p align="center">A city grows over time, new buildings emerged, old ones abandoned and destroyed.</p>
</p>

If you are getting started with a project, monolith architecture is your safe bet. Monolith is easy to maintain/develop. Even though it won't scale very well.

You can start considering microservices when you need to scale some features independently.

It is usually better to __start with a monolith__ and __migrate to microservices__ later.

Yet, migration to microservices is not always easy. There are many stories of companies who failed to do so because their code was __not modular/testable__. Their technology turns into a liability that hinders their business growth.

Some companies __invest too much__ in Kubernetes or other expensive technologies. They built a great microservices architecture yet failed to maintain it properly. Or even worse, the technology becomes their cost center.

So we need a middle ground. We need a __monolith application__ that is ready to be __deployed as microservices anytime__.

<p align="center">
    <img src="./images/palmon.gif" width="500px" alt="Animation of Digimon evolution" />
    <p align="center">A good system can be scaled up/down at any time. Like a digimon</p>
</p>


We design `ZtplAppDirectory` with this consideration. With `ZtplAppDirectory` you will have a straightforward development experience like building a simple monolith. But you always know that `ZtplAppDirectory`` is ready for microservices.

Let's see the following scenarios, and see what we can do in each phase:

- __You are starting a business.__ Your IT team probably consists of a CTO and a couple of software engineers. In that case, you start with a monolith application. Make sure your monolith is modular so that you can split it into microservices later. If you have database repositories in your modules, be sure to only call them from the same modules. This will make your modules decoupled from other modules.

- __Your software handles a lot of requests and has to be always available.__ Now you might consider using better hardware. Or you can split your application into microservices. As your application is modular, this process should not be very painful. You might also need to consider deploying your application into a Kubernetes cluster. Your services will be able to talk to each other using `Messagebus` or `RPC call`.

- __You need different technologies/programming languages for your services.__ Python is a good general-purpose programming language. But it is not your best choice when you need to take concurrency into account. In this case, you can write some of your services in other programming languages. Make sure the new service will handle all necessary `events` and `RPC calls`.

Now let's dive into the architecture; so that you can get a better picture of why `ztplAppDirectory` is probably a good solution to your use case.

# Microservices vs Monolith

In 2016, [DHH](https://twitter.com/dhh) wrote an article titled [Majestic Monolith](https://m.signalvnoise.com/the-majestic-monolith/).

Since big tech companies like Google or Netflix use microservices, people are getting curious about this. They start to adopt the architecture without understanding the drawbacks. In the article, DHH argued that not all companies need microservices architecture.

Let's see how microservices and a monolith are different from each other.

## Microservices: The good and the bad


<p align="center">
    <img src="./images/individual-zords.jpg" width="500px" alt="Individual Zords from Mighty Morphin Power Rangers" />
    <p align="center">Individual Zords from Mighty Morphin Power Rangers</p>
</p>


Microservices architecture is good because:

- It is easy to scale up/down only particular services.
- Users can still access the system even though some services are down.
- Services can be developed/deployed independently from each other.

Microservices architecture is bad because:

- There is a lot of network communication.
- Either you build a correct one, or just create a [distributed monolith](https://www.techtarget.com/searchapparchitecture/tip/The-distributed-monolith-What-it-is-and-how-to-escape-it).
- Deploying microservices can be challenging.

## Monolith: The good and the bad

<p align="center">
    <img src="./images/megazord.jpg" width="500px" alt="A Megazord from Mighty Morphin Power Rangers" />
    <p align="center">A Megazord from Mighty Morphin Power Rangers</p>
</p>

Monolith architecture is good because:

- It is easy to run, debug, or deploy.

Monolith architecture is bad because:

- People tend to cross the domain boundary because they can.
- When it is down, the entire system is down.
- When you need to scale up/down, everything is scaled up/down

# Microservices-ready monolith (aka Modular Monolith)


Many companies failed to refactor their monolith application into microservices, thus failing to support their business growth.

Other companies are investing too much in Kubernetes or other expensive technologies. They are investing too much for something they didn't need. Technology become their cost center, and they have very few resources to grow their businesses.

We can improve the situation by creating a __modular monolith application__. At the very beginning of your business, you should start with a cheap monolith app. But later, you should be able to split your application into microservices.

You can think of the __modular monolith application__ as a __super-soldier__ with much training and pieces of equipment. Alone, this super-soldier is already capable of many things. This super soldier can shoot, fight, drive a helicopter, and even hack a satellite.

<p align="center">
    <img src="./images/super-soldier-monolith.jpg" width="500px" alt="A super soldier with all equipments" />
    <p align="center">A single super monolith with every feature activated</p>
</p>

But when you do a bigger mission, you cannot rely on a single super-soldier. A team consisting of multiple soldiers with less equipment and different specializations will do better.

<p align="center">
    <img src="./images/super-soldier-microservices.jpg" width="500px" alt="Multiple soldiers with a single feature for each soldier" />
    <p align="center">Multiple microservices with a very specific task</p>
</p>

## Microservices-ready monolith: The good and the bad

Microservices-ready monolith is good because:

- It can be deployed and executed easily as a monolith.
- It has a single codebase.
- It can be deployed as microservices using [feature flags](feature-flags.md).
- Everyone is aware of code changes.

It is bad because:

- Everyone is aware of any code changes (Code-level authorization is nearly impossible).
- Without proper conventions, people tend to cross their domain/boundaries.

> ⚠️ __Warning:__ Never cross domain/boundaries even if you can. Assume __every module__ will __live in different servers/pods__ and __accessing different databases__. Otherwise, your application will ends up as big spaghetti-code monolith that can't go anywhere.

# Example

## Run as monolith

Suppose you have two modules in `ZtplAppDirectory`:

- Library
- Auth

Those modules serve different domains.

Let's start to deploy your application as a monolith:

<p align="center">
    <img src="images/fastApp-monolith.png" alt="App diagram, monolith mode" />
    <p align="center">Single instance of App with every feature enabled</p>
</p>

In the beginning, you want to enable every feature flag. You can do this by setting all feature-flags environments into `1`.

You also want to use `LocalMessageBus` and `LocalRPC`, so you set `APP_MESSAGE_BUS_TYPE` and `APP_RPC_TYPE` into `local`. By setting your message bus and RPC into local, the inter-module communication can be performed internally without any network overhead. This is a good thing because [network is not reliable](https://particular.net/blog/the-network-is-reliable)

All good, now you can run everything locally.

## Run as microservices

Now you want to run `ZtplAppDirectory` as microservices.

The good news is: You don't need to modify the source code at all. Instead, you just need to deploy the application with different feature flags (configurations).

<p align="center">
    <img src="images/fastApp-microservices.png" alt="App diagram, microservices mode" />
    <p align="center">Multiple instances of App with a different set of features for every instance</p>
</p>


In this example, we want to have:

- __Frontend.__ This service only serves the UI. It won't serve any API request, but it handles the UI of both `auth` and `library` modules.
- __Backend.__ This service only serves API request from `Frontend` or other external applications, and pass them into respective services.
- __Auth Service.__ This service handles any event/RPC call for `auth` module. It also handles every database operation related to `auth` module.
- __Library Service.__ This service handles any event/RPC call for `library` module. It also handles every database operation related to the `library` module.

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