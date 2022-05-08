<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

# A Use Case

Suppose you want to build a simple book catalogue system.

In your first iteration, you want to deploy your book catalogue as a web application. But in the future, you also want to build a mobile app version as well.

Furthermore, you also want to some relevant information in your website. For example, you want to show company profile, office location, etc.

Thus, you decide to split up your system into three components:

* 🐍 `Book Catalogue API`
* 🐸 `Static web server`
* 🐬 `MySQL server`.

![Application components](images/from-zero-to-cloud-architecture.png)

# Discover Dependencies

Your 🐸 `Static web server` might not only serve book catalogue. It also show company profile and other information. Thus, you want your 🐸 `Static web server` to be independent from other components.

In the other hand, your 🐍 `Book Catalogue API` is pretty unusable once the 🐬 `MySQL server` is down. In this case, you can say that your `Book Catalogue API` __depends on__ `MySQL Server`.

![Component dependencies](images/from-zero-to-cloud-dependencies.png)

# Create a Project

```bash
mkdir -p examples/playground/myEndToEndDemo
cd examples/playground/myEndToEndDemo
zaruba please initProject
```

# Add MySQL

```bash
cd examples/playground/myEndToEndDemo
zaruba please addMysql appDirectory=myDb
```

# Add Book Catalogue API

```bash
cd examples/playground/myEndToEndDemo
zaruba please addFastApiCrud \
  appDirectory=myBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["myDb"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'
```

# Add Static Web Server

```bash
cd examples/playground/myEndToEndDemo
zaruba please addNginx \
  appDirectory=myFrontend \
  appPorts='["80:80"]' \
  appEnvs='{"API_HOST":"localhost:3000"}'
```


# Create Front Page

```bash
cd examples/playground/myEndToEndDemo
```


# Run Project

```bash
cd examples/playground/myEndToEndDemo

zaruba please start -t -w 3s
# or
# zaruba please start
# and press ctrl + c once you finish evaluating
```


# Run Project as Containers


```bash
cd examples/playground/myEndToEndDemo

zaruba please startContainers -t -w 3s
# or:
# zaruba please startContainers

zaruba please stopContainers
```


# Push Images


```bash
cd examples/playground/myEndToEndDemo
# zaruba please setProjectValue \
#   variableName=defaultImagePrefix \
#   variableValue=gofrendi
# zaruba please pushImages imagePrefix=gofrendi
```

# Create Deployments


```bash
cd examples/playground/myEndToEndDemo
zaruba please addAppHelmDeployment appDirectory=myDb
zaruba please addAppHelmDeployment appDirectory=myBackend
zaruba please addAppHelmDeployment appDirectory=myFrontend
zaruba please syncEnv
```


# Deploy to Kubernetes

```bash
cd examples/playground/myEndToEndDemo
# zaruba please setProjectValue \
#   variableName=defaultKubeContext \
#   variableValue=docker-desktop
zaruba please deploy kubeContext=docker-desktop
zaruba please destroy kubeContext=docker-desktop
```

