<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

This end-to-end tutorial shows you how you can use Zaruba to develop/deploy your application.

At the end of this tutorial, you will have:

* A working  🐍 backend + 🐸 frontend application.
* A single command to run everything on your 🖥️ local computer.
* A single command to run everything on your local computer as 🐳 containers.
* A single command to deploy everything on your ☸️ kubernetes cluster.

# A Use Case

Suppose you want to build a simple book catalog system.

You want to deploy your book catalog as a web application in your first iteration. But in the future, you also want to build a mobile app version of your web.

Furthermore, you also want to show some relevant information on your website. For example, you want to show your company profile, office location, etc.

Thus, you decide to split up your system into three components:

* 🐍 `Book Catalog API`
* 🐸 `Static Web Server`
* 🐬 `MySQL server`.

![Application components](images/from-zero-to-cloud-architecture.png)

🐍 `Book Catalog API` handles your `business logic`. Users and other third-party applications can talk to it using `API Requests/Responses`.

🐬 `MySQL server` handles your `data storage`. It needs to be accessible from 🐍 `Book Catalog API`, but it doesn't need to be accessible outside the system. To store/fetch data, users should send requests to `Book Catalog API`.

The 🐸 `Static Web Server` handles your `user interface`. The user interface helps users to send `API requests` and fetch responses.

# Discover Dependencies

Your 🐸 `Static Web Server` not only serves book catalog. It also shows your company profile and other information. You want `Static Web Server` to keep running, even if 🐍 `Book Catalog API` is down. Thus, your `Static Web Server` should be independent of other components.

But, your 🐍 `Book Catalog API` is unusable once the 🐬 `MySQL server` is down. Your business logic always involves data storage. Thus, your `Book Catalog API` __depends on__ `MySQL Server`.

![Component dependencies](images/from-zero-to-cloud-dependencies.png)

# Create a Project

To start working with Zaruba, you need to create a [project](../core-concepts/project/README.md).

```bash
mkdir -p myProject
cd myProject
zaruba please initProject
```

Inside this project, you can add some [task](../core-concepts/task/README.md) definitions and other resources.

You can find several wrapper tasks under `./index.zaruba.yaml`:

* `start`: Start applications.
* `startContainers`: Start applications as containers.
* `stopContainers`: Stop application containers.
* `buildImages`: Build application images.
* `pushImages`: Push application images to the image registry.
* `prepareDeployments`: Prepare application deployments.
* `deploy`: Deploy applications.
* `destroy`: Destroy application deployments.

To make those tasks usable, you need to add new applications.

For example, once you already have a 🐬 `MySQL server`, a 🐍 `Book Catalog API`, and a 🐸 `Static Web Server` in your project. You can start them all by invoking: 

```bash
zaruba please start
```

# Add MySQL

Let's add 🐬 `MySQL server` to your project:

```bash
zaruba please addMysql \
  appDirectory=demoDb
```

This command does several things at once:

* Create MySQL related resources under the `./demoDb` directory.
* Create scripts to manage `demoDb` under the `./zaruba-tasks/demoDb` directory.
* Register the tasks into `./index.zaruba.yaml`.

You can run `demoDb` by invoking:

```bash
zaruba please startDemoDb
# or
# zaruba please startDemoDbContainer
```

By default, your `demoDb` will run on port `3306`. You can change this (and other configurations) by editing `./.env`:

```bash
DEMO_DB_MYSQL_DATABASE="sample"
DEMO_DB_MYSQL_PASSWORD="mysql"
DEMO_DB_MYSQL_ROOT_PASSWORD="Alch3mist"
DEMO_DB_MYSQL_USER="mysql"
```

You can run MySQL CLI from inside the `demoDb` container by invoking:

```bash
docker exec -it demoDb mysql -u root -pAlchemist
```

# Add Book Catalog API

Now, let's add 🐍 `Book Catalog API` to your project.

```bash
zaruba please addFastAppCrud \
  appDirectory=demoBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

zaruba task addDependencies runDemoBackend runDemoDb
zaruba task addDependencies runDemoBackendContainer runDemoDbContainer

# or
# zaruba please addFastAppCrud -i
```

This command does several things at once:

* Create a FastAPI CRUD application under the `./demoBackend` directory.
  * Declare that the CRUD application should handle the `books` data.
  * Declare that `books` CRUD handler should be located under `./demoBackend/library`.
  * Declare that `books` entity has 3 fields: `title`, `author`, `synopsis`.
* Create scripts to manage `demoBackend` under `./zaruba-tasks/demoBackend` directory.
* Register the tasks into `./index.zaruba.yaml`.
* Declare that Zaruba needs to start `demoDb` first before running `demoBackend`. 
* Override several environment variables for `demoBackend`:
  * `APP_HTTP_PORT`: `3000`
  * `APP_SQLALCHEMY_DATABASE_URL`: `mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4`

## Create Migration

```bash
zaruba please createDemoBackendMigration
```

# Run Backend

You can run `demoBackend` by invoking:

```bash
zaruba please startDemoBackend
# or
# zaruba please startDemoBackendContainer
```

By default, your `demoBackend` will run on port `3000`. You can change this (and other configurations) by editing `./.env`:

```bash
DEMO_BACKEND_APP_HTTP_PORT=3000
DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL="root@innistrad.com"
DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME="root"
DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD="Alch3mist"
DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER=621234567890
DEMO_BACKEND_APP_ROOT_PERMISSION="root"
DEMO_BACKEND_APP_ROOT_USERNAME="root"
```

Please note, that whenever you start `demoBackend`, Zaruba will always run `demoDb` first.

To Access the book catalog API, you can open `http://localhost:3000/docs` from your browser.

![DemoBackend API](images/from-zero-to-cloud-backend-api.png)

To access the API, you need to authorize yourself by clicking `Authorize` button. Your username should be `root`, and your password should be `Alch3mist`. You can leave the other inputs blank.

Once authorized, you can start adding new books by expanding `POST /books/` panel, and clicking `Try it out` button.

Your request body should be in JSON format, like this one:

```json
{
  "title": "Doraemon",
  "author": "Fujiko F. Fujio",
  "synopsis": "Robot cat from 22 century, helping Nobita doing his job."
}
```

![DemoBackend Add Book](images/from-zero-to-cloud-backend-add-book.png)

Once you define the request body, you can click `Execute button`.

To see your book list, you can expand `GET /api/v1/books/` panel.

# Add Static Web Server

Now let's add 🐸 `Static Web Server` to your project.

```bash
zaruba please addNginx \
  appDirectory=demoFrontend \
  appPorts='["8080:80", "443"]
```

This command does several things at once:

* Create Nginx related resources under the `./demoFrontend` directory.
* Create scripts to manage `demoFrontend` under the `./zaruba-tasks/demoFrontend` directory.
  * Map port `80` of the `demoFrontend` container to port `8080` of the host.
  * Map port `443` of the `demoFrontend` container to port `443` of the host.
* Register the tasks into `./index.zaruba.yaml`.

You can run the `Static Web Server` by invoking:

```bash
zaruba please startDemoBackend
# or
# zaruba please startDemoBackendContainer
```

Or you can run the `Static Web Server` along with other components by invoking:

```bash
zaruba please start
# ctrl + c
# zaruba please stopContainers
```

To Access the book catalog API, you can open `http://localhost:8080` from your browser.

# Override Static Web Server's Init Script

Your 🐸 `Static Web Server` needs to let your users know where your 🐍 `Book Catalog API` is.

We can do this by:

* Add `API_HOST` environment variable to `demoFrontend` container.
* Overriding `demoFrontend` container's init script to create a JavaScript file.

## Add `API_HOST` Environment Variable

To add `API_HOST` environment variable, you have to create environment file under `./demoFrontend`:

```bash
# fileName: ./demoFrontend/template.env
API_HOST=http://localhost:3000
```

Once you did it, you should run `syncEnv` task:

```bash
zaruba please syncEnv
```

##  Overriding `demoFrontend` Init Script

Now, let's define your new init script under `./demoFronted`:

```bash
# fileName: ./demoFrontend/bootstrap.sh
echo "var apiHost=\"$API_HOST\";" > /opt/bitnami/nginx/html/apiHost.js && /opt/bitnami/scripts/nginx/run.sh
```

This script does three things:

* Read `$API_HOST` from environment variable, and put it into JavaScript.
* Creating a JavaScript file named `apiHost.js` under container's `/opt/bitnami/html`.
* Run original init script (`/opt/bitnami/scripts/nginx/run.sh`)

Next, let's add `bootstrap.sh` to your docker image by editing `./demoFrontend/Dockerfile`:

```
FROM docker.io/bitnami/nginx:1.21.6

# become root to install certbot
USER 0
RUN apt update && \
    apt install certbot -y && \
    apt-get autoremove -yqq --purge && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# stop become root
USER 1001

COPY html /opt/bitnami/nginx/html
COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf 
USER 0
COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
RUN chmod 755 /opt/bitnami/nginx/html
USER 1001
CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
```

If you have run `demoFrontend` before, you need to remove the container first:

```bash
zaruba please removeDemoFrontendContainer
```

Finally, you have to make your `./demoFrontend/html` accessible:

```bash
chmod 777 -R ./demoFrontend/html
```

Now, whenever you start `demoFrontend` container, `bootstrap.sh` will create `apiHost.js` based on `API_HOST` environment variable.

# Create User Interface Page

You already have `apiHost.js` in your 🐸 `Static Web Server`. Now, let's create a user interface on it.

To do this, you need to modify `./demoFrontend/html/index.html` as follow:

<details>
<summary>Show the script</summary>

```
<!DOCTYPE html>
<html>
<head>
<title>Book Catalog</title>
<style>
    html { color-scheme: light dark; }
    body {
        width: 35em; margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }

    .toggle-content {
        display: none;
    }

    .toggle-content.is-visible {
        display: block;
    }
</style>
</head>
<body>
<h1>Welcome</h1>
<p>A book catalog</p>

<div id="unauthenticated-container" class="toggle-content is-visible">
    <form>
        <label>User</label>
        <br />
        <input id="input-login-user" placeholder="user" />
        <br />

        <label>Password</label>
        <br />
        <input id="input-login-password" placeholder="password" type="password" />
        <br />

        <button id="btn-login">Login</button>
    </form>
</div>

<div id="authenticated-container" class="toggle-content">
    <a id="btn-logout" href="#">Logout</a>

    <ul id="book-list">
    </ul>

    <form>
        <label>Title</label>
        <br />
        <input id="input-new-book-title"/>
        <br />

        <label>Author</label>
        <br />
        <input id="input-new-book-author"/>
        <br />

        <label>Synopsis</label>
        <br />
        <textarea id="input-new-book-synopsis"></textarea>
        <br />

        <button id="btn-new-book">Add</button>
    </form>
</div>

<script src="/apiHost.js"></script>
<script>

    function showAuthenticatedContainer() {
        const authenticatedContainer = document.getElementById("authenticated-container");
        const unauthenticatedContainer = document.getElementById("unauthenticated-container");
        unauthenticatedContainer.classList.remove("is-visible");
        authenticatedContainer.classList.add("is-visible");
        populateBookList();
    }

    function showUnAuthenticatedContainer() {
        const authenticatedContainer = document.getElementById("authenticated-container");
        const unauthenticatedContainer = document.getElementById("unauthenticated-container");
        authenticatedContainer.classList.remove("is-visible");
        unauthenticatedContainer.classList.add("is-visible");
    }

    function populateBookList() {
        const xhttp = new XMLHttpRequest();
        xhttp.onload = function() {
            if (this.status != 200) {
                console.error(this.responseText);
                logout();
                return;
            }
            const bookList = JSON.parse(this.responseText);
            let html = "";
            for (let bookIndex = 0; bookIndex < bookList.length; bookIndex++) {
                const book = bookList[bookIndex];
                html += "<li>";
                html += "<p>" + book.title + "(" + book.author + ")" + "</p>";
                html += "<p>" + book.synopsis + "</p>";
                html += '<a class="btn-delete-book" href="#" onclick="deleteBook(\'' + book.id + '\');">Delete</a>';
                html += "</li>";
            }
            document.getElementById("book-list").innerHTML = html;
        }
        xhttp.open("GET", apiHost + "/api/v1/books/", true);
        xhttp.setRequestHeader("Content-type", "application/json")
        xhttp.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("accessToken"))
        xhttp.send();
    }

    function deleteBook(bookId) {
        const xhttp = new XMLHttpRequest();
        xhttp.onload = function() {
            if (this.status != 200) {
                console.error(this.responseText);
                return;
            }
            populateBookList();
        }
        xhttp.open("DELETE", apiHost + "/api/v1/books/" + bookId, true);
        xhttp.setRequestHeader("Content-type", "application/json")
        xhttp.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("accessToken"))
        xhttp.send();
    }

    function addNewBook() {
        const newBookTitle = document.getElementById("input-new-book-title").value;
        const newBookAuthor = document.getElementById("input-new-book-author").value;
        const newBookSynopsis = document.getElementById("input-new-book-synopsis").value;
        const xhttp = new XMLHttpRequest();
        xhttp.onload = function() {
            if (this.status != 200) {
                console.error(this.responseText);
                return;
            }
            populateBookList();
        }
        xhttp.open("POST", apiHost + "/api/v1/books/", true);
        xhttp.setRequestHeader("Content-type", "application/json")
        xhttp.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("accessToken"))
        xhttp.send(JSON.stringify({
            title: newBookTitle,
            author: newBookAuthor,
            synopsis: newBookSynopsis,
        }));
    }

    function login() {
        const user = document.getElementById("input-login-user").value;
        const password = document.getElementById("input-login-password").value;
        const xhttp = new XMLHttpRequest();
        xhttp.onload = function() {
            if (this.status != 200) {
                window.alert("invalid login");
                return;
            }
            const response = JSON.parse(this.responseText);
            const accessToken = response.access_token;
            localStorage.setItem("accessToken", accessToken);
            showAuthenticatedContainer();
        }
        xhttp.open("POST", apiHost + "/api/v1/create-oauth-access-token/", true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send("username=" + encodeURIComponent(user) + "&password=" + encodeURIComponent(password));
    }

    function logout() {
        localStorage.removeItem("accessToken");
        showUnAuthenticatedContainer();
    }

    function main() {

        document.getElementById("btn-login").addEventListener("click", function(e) {
            e.preventDefault();
            login();
        });

        document.getElementById("btn-logout").addEventListener("click", function(e) {
            e.preventDefault();
            logout();
        });

        document.getElementById("btn-new-book").addEventListener("click", function(e) {
            e.preventDefault();
            addNewBook();
        });

        if(localStorage.getItem("accessToken")) {
            showAuthenticatedContainer();
        }
    }

    main();

</script>
</body>
</html>
```
</details>

That's it. Now, let's rebuild `demoFrontend` image.

```
zaruba please removeDemoFrontendContainer
zaruba please buildDemoFrontendImage
```

# Run Project

Now, you can run your 🐍 `Book Catalog API`, 🐸 `Static Web Server`, and 🐬 `MySQL server` by invoking:

```bash
zaruba please start
# ctrl + c
# zaruba please stopContainers
```

Now, let's access your `Static Web Server` by opening `http://localhost:8080/`.

First of all, users need to log in. You can use `root`/`Alch3mist` as your credentials:

![Login page](images/from-zero-to-cloud-frontend-login.png)

Once users log in, they can start adding books:

![Catalog page](images/from-zero-to-cloud-frontend-add-book.png)

Please note that `Static Web Server` and `MySQL Server` are running as containers. Thus, to stop your applications, you need to press `Ctrl+C` and invoke `zaruba please stopContainers`.

# Run Project as Containers

To run your applications as containers, you can invoke:

```bash
zaruba please startContainers
# ctrl + c
# zaruba please stopContainers
```

While to stop your applications, you can press `Ctrl+C` and invoke `zaruba please stopContainers`.

# Build and Push Images

> __💡 HINT__ You can skip this step if your kubernetes-cluster runs on top of `docker-desktop`.

To deploy your applications in your kubernetes cluster, you need to build and push your images to image registry:

```bash
docker login -u <your-container-registry>
zaruba project setValue defaultImagePrefix <your-container-registry>/<your-dockerhub-user-name>
zaruba please pushImages
```

The commands do two things:

* Set your default image prefix to `<your-dockerhub-user-name>`
* Push images to `<your-container-registry>`.

If you are not sure, you can sign up to [hub.docker.com](https://hub.docker.com) and use `docker.io` as `<your-container-registry>`.

If you use `docker.io` as container registry, you probably have to push your images one by one:

```bash
zaruba please pushDemoDb
zaruba please pushDemoBackend
zaruba please pushDemoFrontend
```

![docker images](images/from-zero-to-cloud-docker-images.png)

# Add Kubernetes Deployments

To be able to deploy your applications, you need to create deployment tasks:

```bash
zaruba please addMysqlHelmDeployment deploymentDirectory=demoDbDeployment

zaruba please addAppHelmDeployment appDirectory=demoBackend
zaruba task setEnv deployDemoBackendDeployment APP_SQLALCHEMY_DATABASE_URL "mysql+pymysql://root:Alch3mist@demo-db/sample?charset=utf8mb4"
zaruba task setEnv deployDemoDbDeployment FULLNAME_OVERRIDE "demo-db"

zaruba please addAppHelmDeployment appDirectory=demoFrontend

zaruba please syncEnv
```

It is important to invoke `zaruba plese syncEnv`, since this command allows you to adjust environment definitions for your deployments.

Since your 🐍 `Book Catalog API` and 🐸 `Static Web Server` needs to be accessed from ouside the cluster, you need to set their service type into [LoadBalancer](https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/).

```bash
zaruba task setConfig prepareDemoBackendDeployment serviceType LoadBalancer
zaruba task setConfig prepareDemoFrontendDeployment serviceType LoadBalancer
zaruba task setConfig prepareDemoFrontendDeployment ports 80
zaruba project setValue defaultKubeContext <your-kube-context>
zaruba project setValue pulumiUseLocalBackend yes
```

# Deploy

Let's deploy your applications into Kubernetes

```bash
zaruba please deploy
```

![Kubernetes resources](images/from-zero-to-cloud-kubernetes-resources.png)

If you are using `docker-desktop`, you should be able to access:

* 🐍 `Book Catalog API` on port `3000`
* 🐸 `Static Web Server` on port `80`

![demoFrontend on Kubernetes](images/from-zero-to-cloud-kubernetes-demo-frontend.png)

[Now, you are prepared](https://www.youtube.com/watch?v=YRNBtaHPkZU).

# Wrap Up

Let's do everything at once.

```bash
# 💀 Make Project
mkdir -p examples/playground/use-cases/fromZeroToCloud
cd examples/playground/use-cases/fromZeroToCloud
zaruba please initProject

# 💀 Add DB
zaruba please addMysql \
  appDirectory=demoDb

# 💀 Add Backend
zaruba please addFastAppCrud \
  appDirectory=demoBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

zaruba task addDependencies runDemoBackend runDemoDb
zaruba task addDependencies runDemoBackendContainer runDemoDbContainer

# 💀 Add Frontend
zaruba please addNginx \
  appDirectory=demoFrontend \
  appPorts='["8080:80", "443"]'

chmod -R 777 demoFrontend/html

# 💀 Add .gitignore
echo '' >> demoFrontend/.gitignore
echo 'html/apiHost.js' >> demoFrontend/.gitignore

# 💀 Add environment and sync
echo "API_HOST=http://localhost:3000" > demoFrontend/template.env
zaruba please syncEnv

zaruba task setConfigs startDemoFrontendContainer localhost localhost

# 💀 Add bootstrap
echo 'echo "var apiHost=\"$API_HOST\";" > /opt/bitnami/nginx/html/apiHost.js && /opt/bitnami/scripts/nginx/run.sh' > demoFrontend/bootstrap.sh

# 💀 Overwrite index.html
cp ../../../use-cases/from-zero-to-cloud/index.html demoFrontend/html/index.html

# Modify Dockerfile
echo '' >> demoFrontend/Dockerfile
echo 'USER 0' >> demoFrontend/Dockerfile
echo 'COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh' >> demoFrontend/Dockerfile
echo 'RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh' >> demoFrontend/Dockerfile
echo 'RUN touch /opt/bitnami/nginx/html/apiHost.js' >> demoFrontend/Dockerfile
echo 'RUN chown -R 1001 /opt/bitnami/nginx/html/apiHost.js' >> demoFrontend/Dockerfile
echo 'USER 1001' >> demoFrontend/Dockerfile
echo 'CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]' >> demoFrontend/Dockerfile

zaruba please buildImages

zaruba please start
# <ctrl + c>

# zaruba please startContainers
zaruba please startContainers

zaruba please stopContainers
zaruba please removeContainers

# 💀 Create deployment
zaruba please addMysqlHelmDeployment deploymentDirectory=demoDbDeployment
zaruba please addAppHelmDeployment appDirectory=demoBackend
zaruba task setEnv deployDemoBackendDeployment APP_SQLALCHEMY_DATABASE_URL "mysql+pymysql://root:Alch3mist@demo-db/sample?charset=utf8mb4"
zaruba task setEnv deployDemoDbDeployment FULLNAME_OVERRIDE "demo-db"
zaruba please addAppHelmDeployment appDirectory=demoFrontend


# 💀 Synchronize environment
zaruba please syncEnv

# zaruba project setValue defaultImagePrefix gofrendi
# zaruba please pushImages

zaruba task setConfig prepareDemoBackendDeployment serviceType LoadBalancer
zaruba task setConfig prepareDemoFrontendDeployment serviceType LoadBalancer
zaruba task setConfig prepareDemoFrontendDeployment ports 80
zaruba project setValue defaultKubeContext docker-desktop
zaruba project setValue pulumiUseLocalBackend yes

zaruba please deploy
zaruba please destroy
```
