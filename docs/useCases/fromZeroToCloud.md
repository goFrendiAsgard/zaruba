<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

This end-to-end tutorial shows you how you can use Zaruba to develop/deploy your application.

At the end of this tutorial, you will have:

- A working [microservices-ready monolith](https://github.com/state-alchemists/amalgam/blob/master/myProject/myApp/_docs/motivation-and-architecture/README.md).
- A single command to run everything on your 🖥️ local computer.
- A single command to run everything on your local computer as 🐳 containers.
- A single command to deploy everything on your ☸️ kubernetes cluster.

To see the end result, you can visit: [amalgam repository](https://github.com/state-alchemists/amalgam).

To follow this tutorial, you need to have:

- `Zaruba` installed. Please visit this [guide](../../README.md#-installation) to install Zaruba quickly.
- [Python](https://www.python.org/) with [venv](https://docs.python.org/3/library/venv.html) installed.
- [Docker desktop](https://www.docker.com/products/docker-desktop/) installed, with [kubernetes enabled](https://www.techrepublic.com/article/how-to-install-docker-desktop-kubernetes-support/).

# Start a project

```bash
echo "👷 Make Project a zaruba project"
cd myProject
zaruba please initProject
```

A [project](../core-concepts/project/README.md) is a container for your tasks and resources. In order to create task/applications with Zaruba, you need to create a project first.

# Create an application

```bash
echo "👷 Add App application + it's runner"
zaruba please addFastAppCrud \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudFields='["title", "author"]'
```

Once your project has been created, you can start adding application to your project. In this example, we want to:

- Create an application named `myApp`.
- Create a module named `library` inside `myApp`.
- Create a CRUD entity named `books` inside `library` module.
- Deefine that `books` entity has two fields, `title` and `author`.

Note that your application has several [feature-flags](https://github.com/state-alchemists/amalgam/blob/master/myProject/myApp/_docs/motivation-and-architecture/feature-flags.md) that you can set/unset to change it's behavior. It also has a [layered architecture](https://github.com/state-alchemists/amalgam/blob/master/myProject/myApp/_docs/motivation-and-architecture/interface-and-layers.md) to keep the [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns).

The application is created on top of [FastAPI](https://fastapi.tiangolo.com/) framework and written in [Python](https://www.python.org/).

![](../../zaruba-tasks/make/fastApp/appTemplate/ztplAppDirectory/_docs/motivation-and-architecture/images/fastApp-monolith.png)

## Add new field

Now, let's say you want to add `synopsis` field to your `books` entity. You can do that by invoking the following command:

```bash
echo "👷 Add synopsis field to Book Entity"
zaruba please addFastAppCrudField \
    appDirectory=myApp \
    appModuleName=library \
    appCrudEntity=books \
    appCrudField=synopsis
```

## Add homepage

Furthermore, you can also add a page in your application by invoking the following command:

```bash
echo "👷 Add homepage"
zaruba please addFastAppPage \
    appDirectory=myApp \
    appModuleName=library \
    appHttpMethod=get \
    appUrl=/
```

The page is written in [HTML](https://en.wikipedia.org/wiki/HTML) and [jinja](https://jinja.palletsprojects.com/en/3.1.x/).

You can always visit `<application>/_jinja_templates/<module>/<url>.html` to do some necessary editing.

When rendering your page, the application will automatically add menu bar, bootstrap, vue.js, axios, etc.

## Generate migration

```bash
echo "👷 Generate myAppMigration"
zaruba please createMyAppMigration
```

Last but not least, you can add database migration using [alembic](https://alembic.sqlalchemy.org/en/latest/). You need to generate migration whenever you add new entity/fields.

# Run your application as a monolith

Once your application is ready, you can start it by invoking:

```bash
echo "👷 Run myApp"
zaruba please startApp
# ctrl + c
```

# Run your application as a container

If you want to run your application as a [container](https://www.docker.com/resources/what-container/), you can invoke this instead:

```bash
echo "👷 Run myApp as container"
zaruba please startAppContainer
# ctrl + c
# zaruba please stopContainers
```

Running your application in container is a good step if you want to run your application in [kubernetes](https://kubernetes.io/) or any orchestration system.

# Turn your application into microservices

Once your business grow bigger, you probably needs to scale your use cases differently. In this case, scalability is very important, so that you can even tolerate [network unrelaibility](https://particular.net/blog/the-network-is-reliable).

As an analogy, a person can single handledly manage his/her/their business. But once the business grow, he/she/they need to hire more people, manage payroll, and handle communication. It is more work to do, but it is necessary.

Since your application has [feature-flags](https://github.com/state-alchemists/amalgam/blob/master/myProject/myApp/_docs/motivation-and-architecture/feature-flags.md) and [layered architecture](https://github.com/state-alchemists/amalgam/blob/master/myProject/myApp/_docs/motivation-and-architecture/interface-and-layers.md), you can easily deploy multiple instances of your application. Each of them has specific purposes, and together they work as a team.

![](../../zaruba-tasks/make/fastApp/appTemplate/ztplAppDirectory/_docs/motivation-and-architecture/images/fastApp-microservices.png)

## Add Mysql

First, you need two instances of [MySQL](https://www.mysql.com/). Each to serve different services (i.e., `authSvc` and `libSvc`):

```bash
echo "👷 Add mysql for authSvc"
zaruba please addMysql \
    appDirectory=myAuthSvcDb \
    appPorts='["3307:3306"]'
zaruba task setConfig startMyAuthSvcDbContainer afterCheck 'sleep 10'

echo "👷 Add mysql for libSvc"
zaruba please addMysql \
    appDirectory=myLibSvcDb \
    appPorts='["3308:3306"]'
zaruba task setConfig startMyLibSvcDbContainer afterCheck 'sleep 10'
```

## Add Rabbitmq

For communication between your services, we can use [Rabbitmq](https://www.rabbitmq.com/):

```bash
echo "👷 Add rabbitmq"
zaruba please addRabbitmq appDirectory=myRabbitmq
zaruba task setConfig startMyRabbitmqContainer afterCheck 'sleep 15'
```

## Add frontend

Now you need to add a `frontend` service to your system. It is still the same application you previously created. You just run it with different configuration and feature flag.

Your frontend service needs to talk to `authSvc` service, so you need to start `rabbitmq` first before starting your `frontend` service.

Also, we don't want any specific database migration, so we override `prepareMyFrontend.start` an `migrateMyFrontend.start`.

```bash
echo "👷 Add frontend"
zaruba please makeFastAppRunner \
    appDirectory=myApp \
    appName=myFrontend

zaruba task setEnv startMyFrontend APP_HTTP_PORT 3001
zaruba task setEnv startMyFrontend APP_ENABLE_API 0
zaruba task setEnv startMyFrontend APP_ENABLE_RPC_HANDLER 0
zaruba task setEnv startMyFrontend APP_ENABLE_EVENT_HANDLER 0
zaruba task setEnv startMyFrontend APP_RPC_TYPE rmq
zaruba task setEnv startMyFrontend APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv startMyFrontend APP_RABBITMQ_HOST localhost
zaruba task setEnv startMyFrontend APP_RABBITMQ_USER root
zaruba task setEnv startMyFrontend APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv startMyFrontend APP_RABBITMQ_VHOST /
zaruba task setEnv startMyFrontend APP_UI_BACKEND_URL http://localhost:3002
zaruba task setEnv startMyFrontend APP_SEED_ROOT_USER 0

zaruba task addDependencies prepareMyFrontend prepareMyApp
zaruba task setConfig prepareMyFrontend start 'echo "Done"'
zaruba task addDependencies migrateMyFrontend migrateMyApp
zaruba task setConfig migrateMyFrontend start 'echo "Done"'

zaruba task addDependencies startMyFrontend startMyRabbitmq
zaruba task addDependencies startMyFrontendContainer startMyRabbitmqContainer
```

## Add backend

Your `backend` service will respond to any external HTTP requests, including the ones from your `frontend`.

It needs to talk to both `authSvc` and `libSvc`, so you should start `rabbitmq` first before running your `backend`.

```bash
echo "👷 Add backend"
zaruba please makeFastAppRunner \
    appDirectory=myApp \
    appName=myBackend

zaruba task setEnv startMyBackend APP_HTTP_PORT 3002
zaruba task setEnv startMyBackend APP_ENABLE_UI 0
zaruba task setEnv startMyBackend APP_ENABLE_RPC_HANDLER 0
zaruba task setEnv startMyBackend APP_ENABLE_EVENT_HANDLER 0
zaruba task setEnv startMyBackend APP_RPC_TYPE rmq
zaruba task setEnv startMyBackend APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv startMyBackend APP_RABBITMQ_HOST localhost
zaruba task setEnv startMyBackend APP_RABBITMQ_USER root
zaruba task setEnv startMyBackend APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv startMyBackend APP_RABBITMQ_VHOST /
zaruba task setEnv startMyBackend APP_SEED_ROOT_USER 0

zaruba task addDependencies prepareMyBackend prepareMyApp
zaruba task setConfig prepareMyBackend start 'echo "Done"'
zaruba task addDependencies migrateMyBackend migrateMyApp
zaruba task setConfig migrateMyBackend start 'echo "Done"'

zaruba task addDependencies startMyBackend startMyRabbitmq
zaruba task addDependencies startMyBackendContainer startMyRabbitmqContainer
```

## Add authSvc

`AuthSvc` handle authentication/authorization. It depends on `authSvcDb` and `rabbitmq`.

```bash
echo "👷 Add authSvc"
zaruba please makeFastAppRunner \
    appDirectory=myApp \
    appName=myAuthSvc

zaruba task setEnv startMyAuthSvc APP_HTTP_PORT 3003
zaruba task setEnv startMyAuthSvc APP_ENABLE_ROUTE_HANDLER 0
zaruba task setEnv startMyAuthSvc APP_ENABLE_UI 0
zaruba task setEnv startMyAuthSvc APP_ENABLE_API 0
zaruba task setEnv startMyAuthSvc APP_ENABLE_LIBRARY_MODULE 0
zaruba task setEnv startMyAuthSvc APP_RPC_TYPE rmq
zaruba task setEnv startMyAuthSvc APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv startMyAuthSvc APP_RABBITMQ_HOST localhost
zaruba task setEnv startMyAuthSvc APP_RABBITMQ_USER root
zaruba task setEnv startMyAuthSvc APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv startMyAuthSvc APP_RABBITMQ_VHOST /
zaruba task setEnv startMyAuthSvc APP_SQLALCHEMY_DATABASE_URL 'mysql+pymysql://root:Alch3mist@localhost:3307/sample?charset=utf8mb4'

zaruba task addDependencies prepareMyAuthSvc prepareMyApp
zaruba task setConfig prepareMyAuthSvc start 'echo "Done"'

zaruba task addDependencies startMyAuthSvc startMyRabbitmq
zaruba task addDependencies startMyAuthSvcContainer startMyRabbitmqContainer
zaruba task addDependencies migrateMyAuthSvc startMyAuthSvcDb
zaruba task addDependencies startMyAuthSvc startMyAuthSvcDb
zaruba task addDependencies startMyAuthSvcContainer startMyAuthSvcDbContainer
```

## Add libSvc

`LibSvc` handle your business use case (in this case, just book CRUD). It depends on `LibSvcDb` and `rabbitmq`.

```bash
echo "👷 Add libSvc"
zaruba please makeFastAppRunner \
    appDirectory=myApp \
    appName=myLibSvc

zaruba task setEnv startMyLibSvc APP_HTTP_PORT 3004
zaruba task setEnv startMyLibSvc APP_ENABLE_ROUTE_HANDLER 0
zaruba task setEnv startMyLibSvc APP_ENABLE_UI 0
zaruba task setEnv startMyLibSvc APP_ENABLE_API 0
zaruba task setEnv startMyLibSvc APP_ENABLE_AUTH_MODULE 0
zaruba task setEnv startMyLibSvc APP_RPC_TYPE rmq
zaruba task setEnv startMyLibSvc APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv startMyLibSvc APP_RABBITMQ_HOST localhost
zaruba task setEnv startMyLibSvc APP_RABBITMQ_USER root
zaruba task setEnv startMyLibSvc APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv startMyLibSvc APP_RABBITMQ_VHOST /
zaruba task setEnv startMyLibSvc APP_SQLALCHEMY_DATABASE_URL 'mysql+pymysql://root:Alch3mist@localhost:3308/sample?charset=utf8mb4'
zaruba task setEnv startMyLibSvc APP_SEED_ROOT_USER 0

zaruba task addDependencies prepareMyLibSvc prepareMyApp
zaruba task setConfig prepareMyLibSvc start 'echo "Done"'

zaruba task addDependencies startMyLibSvc startMyRabbitmq
zaruba task addDependencies migrateMyLibSvc startMyLibSvcDb
zaruba task addDependencies startMyLibSvcContainer startMyRabbitmqContainer
zaruba task addDependencies startMyLibSvc startMyLibSvcDb
zaruba task addDependencies startMyLibSvcContainer startMyLibSvcDbContainer
```

## Add microservices runner

Once everything is set, you can create a runner for your microservices. You can also perform `zaruba please syncEnv` and edit your `.env` to fit your need.

```bash
echo "👷 Add microservice runner"
zaruba project addTask startMyMicroservices
zaruba project addTask startMyMicroservicesContainers
zaruba task addDependencies startMyMicroservices startMyFrontend
zaruba task addDependencies startMyMicroservicesContainers startMyFrontendContainer
zaruba task addDependencies startMyMicroservices startMyBackend
zaruba task addDependencies startMyMicroservicesContainers startMyBackendContainer
zaruba task addDependencies startMyMicroservices startMyAuthSvc
zaruba task addDependencies startMyMicroservicesContainers startMyAuthSvcContainer
zaruba task addDependencies startMyMicroservices startMyLibSvc
zaruba task addDependencies startMyMicroservicesContainers startMyLibSvcContainer

echo "👷 Synchronize environment"
zaruba please syncEnv
```

# Run your application as microservices

To run your microservices, you can invoke:

```bash
echo "👷 Run myMicroservices"
zaruba please startMyMicroservices
# ctrl + c
```

# Run your application as multiple containers

To run your microservices as containers, you can invoke:

```bash
echo "👷 Run myMicroservices as containers"
zaruba please startMyMicroservicesContainers
# zaruba please stopContainers
```

# Create kubernetes deployments

There are [many ways to run your microservices](https://semaphoreci.com/blog/deploy-microservices). At some point you will find that managing every instance of your microservices is challenging. In that case, you probably need an orchestrator like [kubernetes](https://kubernetes.io/).

Zaruba can help you to create kubernetes deployments for your applications.

Under the hood, Zaruba will make [helm charts](https://helm.sh/docs/topics/charts/) and [pulumi](https://www.pulumi.com/) scripts.

Most cloud provider offer kubernetes services, so you can deploy your applications in [AWS](https://aws.amazon.com/eks/), [GCP](https://cloud.google.com/kubernetes-engine/), [Azure](https://docs.microsoft.com/en-us/azure/aks/), [Okteto](https://www.okteto.com/), etc. But right now, we will use [docker desktop](https://www.docker.com/products/docker-desktop/).

```bash
echo "👷 Set project values"
zaruba project setValue defaultKubeContext docker-desktop
zaruba project setValue pulumiUseLocalBackend yes

echo "👷 Add auth svc db deployment"
zaruba please addMysqlHelmDeployment deploymentDirectory=myAuthSvcDbDeployment
zaruba task setEnv deployMyAuthSvcDbDeployment FULLNAMEOVERRIDE "auth-svc-db"
zaruba task setEnv deployMyAuthSvcDbDeployment AUTH_DATABASE "auth"

echo "👷 Add lib svc db deployment"
zaruba please addMysqlHelmDeployment deploymentDirectory=myLibSvcDbDeployment
zaruba task setEnv deployMyLibSvcDbDeployment FULLNAMEOVERRIDE "lib-svc-db"
zaruba task setEnv deployMyLibSvcDbDeployment AUTH_DATABASE "lib"

echo "👷 Add rabbitmq deployment"
zaruba please addRabbitmqHelmDeployment deploymentDirectory=myRabbitmqDeployment
zaruba task setEnv deployMyRabbitmqDeployment FULLNAMEOVERRIDE "rabbitmq"

echo "👷 Add auth svc deployment"

zaruba please addAppHelmDeployment \
    appDirectory=myApp \
    deploymentDirectory=myAuthSvcDeployment \
    appPorts='["3000"]'

zaruba task setEnv deployMyAuthSvcDeployment LIVENESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyAuthSvcDeployment READINESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyAuthSvcDeployment SERVICE_ENABLED False
zaruba task setEnv deployMyAuthSvcDeployment FULLNAME_OVERRIDE auth-svc
zaruba task setEnv prepareMyAuthSvcDeployment APP_HTTP_PORT 3000
zaruba task setEnv prepareMyAuthSvcDeployment APP_ENABLE_ROUTE_HANDLER 0
zaruba task setEnv prepareMyAuthSvcDeployment APP_ENABLE_UI 0
zaruba task setEnv prepareMyAuthSvcDeployment APP_ENABLE_API 0
zaruba task setEnv prepareMyAuthSvcDeployment APP_ENABLE_LIB_MODULE 0
zaruba task setEnv prepareMyAuthSvcDeployment APP_RPC_TYPE rmq
zaruba task setEnv prepareMyAuthSvcDeployment APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv prepareMyAuthSvcDeployment APP_RABBITMQ_HOST rabbitmq
zaruba task setEnv prepareMyAuthSvcDeployment APP_RABBITMQ_USER root
zaruba task setEnv prepareMyAuthSvcDeployment APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv prepareMyAuthSvcDeployment APP_RABBITMQ_VHOST /
zaruba task setEnv prepareMyAuthSvcDeployment APP_SQLALCHEMY_DATABASE_URL "mysql+pymysql://root:Alch3mist@auth-svc-db/auth?charset=utf8mb4"
zaruba task setEnv prepareMyAuthSvcDeployment APP_RABBITMQ_HOST "rabbitmq"

echo "👷 Add lib svc deployment"

zaruba please addAppHelmDeployment \
    appDirectory=myApp \
    deploymentDirectory=myLibSvcDeployment\
    appPorts='["3000"]'

zaruba task setEnv deployMyLibSvcDeployment LIVENESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyLibSvcDeployment READINESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyLibSvcDeployment SERVICE_ENABLED False
zaruba task setEnv deployMyLibSvcDeployment FULLNAME_OVERRIDE lib-svc
zaruba task setEnv prepareMyLibSvcDeployment APP_HTTP_PORT 3000
zaruba task setEnv prepareMyLibSvcDeployment APP_ENABLE_ROUTE_HANDLER 0
zaruba task setEnv prepareMyLibSvcDeployment APP_ENABLE_UI 0
zaruba task setEnv prepareMyLibSvcDeployment APP_ENABLE_API 0
zaruba task setEnv prepareMyLibSvcDeployment APP_ENABLE_AUTH_MODULE 0
zaruba task setEnv prepareMyLibSvcDeployment APP_RPC_TYPE rmq
zaruba task setEnv prepareMyLibSvcDeployment APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv prepareMyLibSvcDeployment APP_RABBITMQ_HOST rabbitmq
zaruba task setEnv prepareMyLibSvcDeployment APP_RABBITMQ_USER root
zaruba task setEnv prepareMyLibSvcDeployment APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv prepareMyLibSvcDeployment APP_RABBITMQ_VHOST /
zaruba task setEnv prepareMyLibSvcDeployment APP_SQLALCHEMY_DATABASE_URL "mysql+pymysql://root:Alch3mist@lib-svc-db/lib?charset=utf8mb4"

echo "👷 Add frontend deployment"

zaruba please addAppHelmDeployment \
    appDirectory=myApp \
    deploymentDirectory=myFrontendDeployment \
    appPorts='["3001"]'

zaruba task setEnv deployMyFrontendDeployment LIVENESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyFrontendDeployment READINESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyFrontendDeployment SERVICE_TYPE LoadBalancer
zaruba task setEnv deployMyFrontendDeployment SERVICE_ENABLED True
zaruba task setEnv deployMyFrontendDeployment FULLNAME_OVERRIDE frontend
zaruba task setEnv prepareMyFrontendDeployment APP_HTTP_PORT 3001
zaruba task setEnv prepareMyFrontendDeployment APP_ENABLE_API 0
zaruba task setEnv prepareMyFrontendDeployment APP_ENABLE_RPC_HANDLER 0
zaruba task setEnv prepareMyFrontendDeployment APP_ENABLE_EVENT_HANDLER 0
zaruba task setEnv prepareMyFrontendDeployment APP_RPC_TYPE rmq
zaruba task setEnv prepareMyFrontendDeployment APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv prepareMyFrontendDeployment APP_RABBITMQ_HOST rabbitmq
zaruba task setEnv prepareMyFrontendDeployment APP_RABBITMQ_USER root
zaruba task setEnv prepareMyFrontendDeployment APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv prepareMyFrontendDeployment APP_RABBITMQ_VHOST /
zaruba task setEnv prepareMyFrontendDeployment APP_UI_BACKEND_URL http://localhost:3002
zaruba task setEnv prepareMyFrontendDeployment APP_SEED_ROOT_USER 0

echo "👷 Add backend deployment"
zaruba please addAppHelmDeployment \
    appDirectory=myApp \
    deploymentDirectory=myBackendDeployment \
    appPorts='["3002"]'

zaruba task setEnv deployMyBackendDeployment LIVENESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyBackendDeployment READINESS_PROBE_HTTP_GET_PATH /readiness
zaruba task setEnv deployMyBackendDeployment SERVICE_TYPE LoadBalancer
zaruba task setEnv deployMyBackendDeployment SERVICE_ENABLED True
zaruba task setEnv deployMyBackendDeployment FULLNAME_OVERRIDE backend
zaruba task setEnv prepareMyBackendDeployment APP_HTTP_PORT 3002
zaruba task setEnv prepareMyBackendDeployment APP_ENABLE_UI 0
zaruba task setEnv prepareMyBackendDeployment APP_ENABLE_RPC_HANDLER 0
zaruba task setEnv prepareMyBackendDeployment APP_ENABLE_EVENT_HANDLER 0
zaruba task setEnv prepareMyBackendDeployment APP_RPC_TYPE rmq
zaruba task setEnv prepareMyBackendDeployment APP_MESSAGE_BUS_TYPE rmq
zaruba task setEnv prepareMyBackendDeployment APP_RABBITMQ_HOST rabbitmq
zaruba task setEnv prepareMyBackendDeployment APP_RABBITMQ_USER root
zaruba task setEnv prepareMyBackendDeployment APP_RABBITMQ_PASS Alch3mist
zaruba task setEnv prepareMyBackendDeployment APP_RABBITMQ_VHOST /
zaruba task setEnv prepareMyBackendDeployment APP_SEED_ROOT_USER 0
```

# Deploy your application at kubernetes

Once your deployment is ready, you can sync your environment, build images, and start deploying.
To even make it better, you can run this in your CI/CD.

```bash
echo "👷 Synchronize environment"
zaruba please syncEnv

echo "👷 Build image"
zaruba please buildImages

echo "👷 Prepare deployments"
zaruba please prepareDeployments

echo "👷 Deploy"
zaruba please deploy
```

You can do the exact same steps whenever you just add features/fix bugs in your systems.

![](../../zaruba-tasks/make/fastApp/appTemplate/ztplAppDirectory/_docs/motivation-and-architecture/images/on-kubernetes.png)

# Destroy kubernetes deployments

That's all with kubernetes. To remove everything, you can do:

```bash
echo "👷 Prepare deployments"
zaruba please prepareDeployments
```

<!--startTocSubtopic-->

<!--endTocSubtopic-->