<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

Suppose you want to build two applications:

* A 🐍 `CRUD API application` that __depends__ on 🐬 `MySQL server`.
* A simple 🐹 `Go webserver` that has no dependencies.

`Go webserver` has no dependencies. Thus, you should be able to run it __in parallel__ with the `CRUD API application` and the `MySQL server`.

Meanwhile, the `CRUD API application` __depends__ on the `MySQL server`. Thus, you cannot run the `CRUD API application` unless the `MySQL server` is ready.

![Application dependencies](images/from-zero-to-cloud-app-dependencies.png)

Zaruba allows you to build, run, and deploy everything using simple commands (no coding required 😉).

> 💡 __TIPS:__ You can execute tasks with `-i` or `--interactive` flag (i.e: `zaruba please addFastApiCrud -i`).

[![asciicast](https://asciinema.org/a/bZ7kA443kSV40lPiVxPysuAWE.svg)](https://asciinema.org/a/bZ7kA443kSV40lPiVxPysuAWE)


# ✨ Creating Project and Applications

```bash
# ✨ Init project
mkdir myProject
cd myProject
zaruba please initProject

# Add 🐬 MySQL container
zaruba please addMysql appDirectory=myDb

# Add 🐍 CRUD API Application.
zaruba please addFastApiCrud \
  appDirectory=myPythonApp \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["myDb"]' \
  appEnvs='{"APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

# Add 🐹 Go web app.
zaruba please addSimpleGoApp appDirectory=myGoApp appEnvs='{"APP_HTTP_PORT":"3001"}'
```

# 🏃 Run Applications

```bash
# Start 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL container.
# To run this command, you need:
# - go 1.13 or newer
# - python 3.8
# - docker
zaruba please start
# Ctrl+c to stop
```

<details>
<summary>Test applications</summary>

```bash
# Checking 🐹 Go web app.
curl http://localhost:3001

# 🔑 Login to 🐍 CRUD API application
LOGIN_RESPONSE="$(curl --location --request POST 'localhost:3000/token/' \
--form 'username="root"' \
--form 'password="Alch3mist"')"

echo "${LOGIN_RESPONSE}"

# 🔑 Extract token
ACCESS_TOKEN=$(zaruba map get "${LOGIN_RESPONSE}" "access_token")
echo "Token: ${ACCESS_TOKEN}"

# ✨ Insert new book
curl --location --request POST 'localhost:3000/books/' \
--header "Authorization: Bearer ${ACCESS_TOKEN}" \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "The Alchemist",
    "author": "Paulo Coelho",
    "synopsis": "Magical story of Santiago, an Andalusian shepherd boy who yearns to travel in search of a worldly treasure as extravagant as any ever found."
}'

# 📖 Get books
curl --location --request GET 'localhost:3000/books/' \
--header "Authorization: Bearer ${ACCESS_TOKEN}" \
--header 'Content-Type: application/json'
```

</details>

<details>
<summary>Run applications individually</summary>

```bash
# Only start 🐹 Go web server.
zaruba please startMyGoApp
# Ctrl+c to stop

# Only start 🐹 Go web server and 🐍 CRUD API application.
# Please note that MySQL container is automatically started
# since CRUD API application depends on it.
zaruba please startMyGoApp startMyPythonApp
# Ctrl+c to stop
```
</details>

# 🐳 Run Applications as Containers

```bash
# Start 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL as containers.
# To run this command, you need:
# - docker
zaruba please startContainers
zaruba please stopContainers
```

<details>
<summary>Run applications individually</summary>

```bash
# Only start 🐹 Go web server.
zaruba please startMyGoAppContainer
zaruba please stopContainers

# Only start 🐹 Go web server and 🐍 CRUD API application.
# Please note that MySQL container is automatically started
# since CRUD API application depends on it.
zaruba please startMyGoAppContainer startMyPythonAppContainer
# To stop containers (Note: Ctrl+C won't kill the containers):
zaruba please stopContainers
```
</details>

# ☁️ Deploy Applications

```bash
# Deploy 🐹 Go web server, 🐍 CRUD API application, and 🐬 MySQL to kubernetes cluster
# To run this command, you need:
# - kubectl
# - helm
# - pulumi
# - cloud provider or a computer that can run kubernetes locally (we use docker-desktop in this example)
zaruba please buildImages # or `zaruba please pushImages`
zaruba please addAppHelmDeployment appDirectory=myPythonApp
zaruba please addAppHelmDeployment appDirectory=myGoApp
zaruba please addAppHelmDeployment appDirectory=myDb
zaruba please syncEnv
zaruba please deploy kubeContext=docker-desktop
zaruba please destroy kubeContext=docker-desktop
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->