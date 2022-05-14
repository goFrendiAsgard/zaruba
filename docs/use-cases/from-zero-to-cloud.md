<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

# A Use Case

Suppose you want to build a simple book catalog system.

You want to deploy your book catalog as a web application in your first iteration. But in the future, you also want to build a mobile app version of your web.

Furthermore, you also want to show some relevant information on your website. For example, you want to show your company profile, office location, etc.

Thus, you decide to split up your system into three components:

* 🐍 `Book Catalog API`
* 🐸 `Static web server`
* 🐬 `MySQL server`.

![Application components](images/from-zero-to-cloud-architecture.png)

# Discover Dependencies

Your 🐸 `Static web server` might not only serve book catalog. It also shows your company profile and other information. Thus, you want your 🐸 `Static web server` to be independent of other components.

But, your 🐍 `Book Catalog API` is unusable once the 🐬 `MySQL server` is down. In this case, you can say that your `Book Catalog API` __depends on__ `MySQL Server`.

![Component dependencies](images/from-zero-to-cloud-dependencies.png)

# Create a Project

# Add MySQL

# Add Book Catalog API

# Add Static Web Server

# Create Front Page

# Run Project

# Run Project as Containers

# Build and Push Images

# Add Kubernetes Deployments

# Deploy

# Wrap Up

Let's do everything at once.

> __💡 NOTE:__ You can remove `-t` and `-w` parameters

<!--startCode-->
```bash
mkdir -p examples/playground/use-cases/fromZeroToCloud
cd examples/playground/use-cases/fromZeroToCloud
zaruba please initProject

# Add DB
zaruba please addMysql \
  appDirectory=demoDb

# Add Backend
zaruba please addFastApiCrud \
  appDirectory=demoBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["demoDb"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

# Add Frontend
zaruba please addNginx \
  appDirectory=demoFrontend \
  appPorts='["8080:80", "443"]'

chmod -R 777 demoFrontend/html

# Add .gitignore
echo '' >> demoFrontend/.gitignore
echo 'html/apiHost.js' >> demoFrontend/.gitignore

# Add environment and sync
echo "API_HOST=http://localhost:3000" > demoFrontend/template.env
zaruba please syncEnv

zaruba task setConfigs startDemoFrontendContainer localhost localhost

# Add bootstrap
echo 'echo "var apiHost=\"$API_HOST\";" > /opt/bitnami/nginx/html/apiHost.js && /opt/bitnami/scripts/nginx/run.sh' > demoFrontend/bootstrap.sh

# Overwrite index.html
cp ../../../use-cases/from-zero-to-cloud/index.html demoFrontend/html/index.html

# Modify Dockerfile
echo '' >> demoFrontend/Dockerfile
echo 'USER 0' >> demoFrontend/Dockerfile
echo 'COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh' >> demoFrontend/Dockerfile
echo 'RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh' >> demoFrontend/Dockerfile
echo 'USER 1001' >> demoFrontend/Dockerfile
echo 'CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]' >> demoFrontend/Dockerfile

zaruba please buildImages

# zaruba please start
# <ctrl + c>
zaruba please start -t -w 1s

# zaruba please startContainers
zaruba please startContainers -t -w 1s

zaruba please stopContainers
zaruba please removeContainers

zaruba please addAppHelmDeployment appDirectory=demoDb
zaruba please addAppHelmDeployment appDirectory=demoBackend
zaruba please addAppHelmDeployment appDirectory=demoFrontend
zaruba please syncEnv

# zaruba project setValue defaultImagePrefix gofrendi
# zaruba please pushImages

zaruba task setConfig prepareDemoBackendDeployment serviceType LoadBalancer
zaruba task setConfig prepareDemoFrontendDeployment serviceType LoadBalancer
zaruba project setValue defaultKubeContext docker-desktop
zaruba project setValue pulumiUseLocalBackend true

zaruba please deploy
zaruba please destroy
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.228µs
         Current Time: 15:02:15
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 initProject          🚧 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.git/
💀    🚀 initProject          🚧 🎉🎉🎉
💀    🚀 initProject          🚧 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 113.757561ms
         Current Time: 15:02:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 314.022609ms
         Current Time: 15:02:16
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.041µs
         Current Time: 15:02:16
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeMysqlApp         🐬 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 Preparing base variables
💀    🚀 makeMysqlApp         🐬 Base variables prepared
💀    🚀 makeMysqlApp         🐬 Preparing start command
💀    🚀 makeMysqlApp         🐬 Start command prepared
💀    🚀 makeMysqlApp         🐬 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 Preparing test command
💀    🚀 makeMysqlApp         🐬 Test command prepared
💀    🚀 makeMysqlApp         🐬 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 Preparing check command
💀    🚀 makeMysqlApp         🐬 Check command prepared
💀    🚀 makeMysqlApp         🐬 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 ✅ Validate
💀    🚀 makeMysqlApp         🐬 Validate app directory
💀    🚀 makeMysqlApp         🐬 Done validating app directory
💀    🚀 makeMysqlApp         🐬 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 Validate template locations
💀    🚀 makeMysqlApp         🐬 Done validating template locations
💀    🚀 makeMysqlApp         🐬 Validate app ports
💀    🚀 makeMysqlApp         🐬 Done validating app ports
💀    🚀 makeMysqlApp         🐬 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 🚧 Generate
💀    🚀 makeMysqlApp         🐬 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 ]
💀    🚀 makeMysqlApp         🐬 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"fromZeroToCloudMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeMysqlAppRunner   🐬 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 ]
💀    🚀 makeMysqlAppRunner   🐬 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"fromZeroToCloudMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 Checking start
💀    🚀 makeMysqlAppRunner   🐬 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.003773893s
         Current Time: 15:02:21
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.114209458s
         Current Time: 15:02:21
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.13µs
         Current Time: 15:02:21
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeFastApiApp       ⚡ 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ Preparing base variables
💀    🚀 makeFastApiApp       ⚡ Base variables prepared
💀    🚀 makeFastApiApp       ⚡ Preparing start command
💀    🚀 makeFastApiApp       ⚡ Start command prepared
💀    🚀 makeFastApiApp       ⚡ Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ Preparing test command
💀    🚀 makeFastApiApp       ⚡ Test command prepared
💀    🚀 makeFastApiApp       ⚡ Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ Preparing check command
💀    🚀 makeFastApiApp       ⚡ Check command prepared
💀    🚀 makeFastApiApp       ⚡ Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ ✅ Validate
💀    🚀 makeFastApiApp       ⚡ Validate app directory
💀    🚀 makeFastApiApp       ⚡ Done validating app directory
💀    🚀 makeFastApiApp       ⚡ Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ Validate template locations
💀    🚀 makeFastApiApp       ⚡ Done validating template locations
💀    🚀 makeFastApiApp       ⚡ Validate app ports
💀    🚀 makeFastApiApp       ⚡ Done validating app ports
💀    🚀 makeFastApiApp       ⚡ Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ ]
💀    🚀 makeFastApiApp       ⚡ 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"fromZeroToCloudFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeFastApiAppRunner ⚡ 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ ]
💀    🚀 makeFastApiAppRunner ⚡ 
💀    🚀 makeFastApiAppRunner ⚡ 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Checking test
💀    🚀 makeFastApiAppRunner ⚡ Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Checking start
💀    🚀 makeFastApiAppRunner ⚡ Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 addFastApiModule     ⚡ 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ Preparing base variables
💀    🚀 addFastApiModule     ⚡ Base variables prepared
💀    🚀 addFastApiModule     ⚡ Preparing start command
💀    🚀 addFastApiModule     ⚡ Start command prepared
💀    🚀 addFastApiModule     ⚡ Preparing prepare command
💀    🚀 addFastApiModule     ⚡ Prepare command prepared
💀    🚀 addFastApiModule     ⚡ Preparing test command
💀    🚀 addFastApiModule     ⚡ Test command prepared
💀    🚀 addFastApiModule     ⚡ Preparing migrate command
💀    🚀 addFastApiModule     ⚡ Migrate command prepared
💀    🚀 addFastApiModule     ⚡ Preparing check command
💀    🚀 addFastApiModule     ⚡ Check command prepared
💀    🚀 addFastApiModule     ⚡ Preparing replacement map
💀    🚀 addFastApiModule     ⚡ Add config to replacement map
💀    🚀 addFastApiModule     ⚡ Add env to replacement map
💀    🚀 addFastApiModule     ⚡ Replacement map prepared
💀    🚀 addFastApiModule     ⚡ ✅ Validate
💀    🚀 addFastApiModule     ⚡ Validate app directory
💀    🚀 addFastApiModule     ⚡ Done validating app directory
💀    🚀 addFastApiModule     ⚡ Validate app container volumes
💀    🚀 addFastApiModule     ⚡ Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ Validate template locations
💀    🚀 addFastApiModule     ⚡ Done validating template locations
💀    🚀 addFastApiModule     ⚡ Validate app ports
💀    🚀 addFastApiModule     ⚡ Done validating app ports
💀    🚀 addFastApiModule     ⚡ Validate app crud fields
💀    🚀 addFastApiModule     ⚡ Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ ]
💀    🚀 addFastApiModule     ⚡ 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"fromZeroToCloudFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ Registering module
💀    🚀 addFastApiModule     ⚡ Done registering module
💀    🚀 addFastApiModule     ⚡ 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 addFastApiCrud       ⚡ 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ Preparing base variables
💀    🚀 addFastApiCrud       ⚡ Base variables prepared
💀    🚀 addFastApiCrud       ⚡ Preparing start command
💀    🚀 addFastApiCrud       ⚡ Start command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing test command
💀    🚀 addFastApiCrud       ⚡ Test command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing check command
💀    🚀 addFastApiCrud       ⚡ Check command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ Set repo field insert
💀    🚀 addFastApiCrud       ⚡ Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ Set repo field update
💀    🚀 addFastApiCrud       ⚡ Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ Preparing start command
💀    🚀 addFastApiCrud       ⚡ Start command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing test command
💀    🚀 addFastApiCrud       ⚡ Test command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing check command
💀    🚀 addFastApiCrud       ⚡ Check command prepared
💀    🚀 addFastApiCrud       ⚡ Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ ✅ Validate
💀    🚀 addFastApiCrud       ⚡ Validate app directory
💀    🚀 addFastApiCrud       ⚡ Done validating app directory
💀    🚀 addFastApiCrud       ⚡ Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ Validate template locations
💀    🚀 addFastApiCrud       ⚡ Done validating template locations
💀    🚀 addFastApiCrud       ⚡ Validate app ports
💀    🚀 addFastApiCrud       ⚡ Done validating app ports
💀    🚀 addFastApiCrud       ⚡ Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ ]
💀    🚀 addFastApiCrud       ⚡ 
💀    🚀 addFastApiCrud       ⚡ 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"fromZeroToCloudFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ Registering route handler
💀    🚀 addFastApiCrud       ⚡ Done registering route handler
💀    🚀 addFastApiCrud       ⚡ Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ Registering repo
💀    🚀 addFastApiCrud       ⚡ Done registering repo
💀    🚀 addFastApiCrud       ⚡ 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 13.165708334s
         Current Time: 15:02:34
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 13.367576462s
         Current Time: 15:02:35
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.056µs
         Current Time: 15:02:35
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeNginxApp         📗 🧰 Prepare
💀    🚀 makeNginxApp         📗 Preparing base variables
💀    🚀 makeNginxApp         📗 Base variables prepared
💀    🚀 makeNginxApp         📗 Preparing start command
💀    🚀 makeNginxApp         📗 Start command prepared
💀    🚀 makeNginxApp         📗 Preparing prepare command
💀    🚀 makeNginxApp         📗 Prepare command prepared
💀    🚀 makeNginxApp         📗 Preparing test command
💀    🚀 makeNginxApp         📗 Test command prepared
💀    🚀 makeNginxApp         📗 Preparing migrate command
💀    🚀 makeNginxApp         📗 Migrate command prepared
💀    🚀 makeNginxApp         📗 Preparing check command
💀    🚀 makeNginxApp         📗 Check command prepared
💀    🚀 makeNginxApp         📗 Preparing replacement map
💀    🚀 makeNginxApp         📗 Add config to replacement map
💀    🚀 makeNginxApp         📗 Add env to replacement map
💀    🚀 makeNginxApp         📗 Replacement map prepared
💀    🚀 makeNginxApp         📗 ✅ Validate
💀    🚀 makeNginxApp         📗 Validate app directory
💀    🚀 makeNginxApp         📗 Done validating app directory
💀    🚀 makeNginxApp         📗 Validate app container volumes
💀    🚀 makeNginxApp         📗 Done validating app container volumes
💀    🚀 makeNginxApp         📗 Validate template locations
💀    🚀 makeNginxApp         📗 Done validating template locations
💀    🚀 makeNginxApp         📗 Validate app ports
💀    🚀 makeNginxApp         📗 Done validating app ports
💀    🚀 makeNginxApp         📗 Validate app crud fields
💀    🚀 makeNginxApp         📗 Done validating app crud fields
💀    🚀 makeNginxApp         📗 🚧 Generate
💀    🚀 makeNginxApp         📗 🚧 Template Location: [
💀    🚀 makeNginxApp         📗   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 ]
💀    🚀 makeNginxApp         📗 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"fromZeroToCloudNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 🔩 Integrate
💀    🚀 makeNginxApp         📗 🎉🎉🎉
💀    🚀 makeNginxApp         📗 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeNginxAppRunner   📗 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 Preparing start command
💀    🚀 makeNginxAppRunner   📗 Start command prepared
💀    🚀 makeNginxAppRunner   📗 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 Preparing test command
💀    🚀 makeNginxAppRunner   📗 Test command prepared
💀    🚀 makeNginxAppRunner   📗 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 Preparing check command
💀    🚀 makeNginxAppRunner   📗 Check command prepared
💀    🚀 makeNginxAppRunner   📗 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 Validate app directory
💀    🚀 makeNginxAppRunner   📗 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 Validate template locations
💀    🚀 makeNginxAppRunner   📗 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 Validate app ports
💀    🚀 makeNginxAppRunner   📗 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 ]
💀    🚀 makeNginxAppRunner   📗 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"fromZeroToCloudNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"fromZeroToCloudNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 Checking start
💀    🚀 makeNginxAppRunner   📗 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.133389931s
         Current Time: 15:02:40
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.243866316s
         Current Time: 15:02:40
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.342µs
         Current Time: 15:02:40
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 syncEnv              🔄 Synchronize task environments
💀    🚀 syncEnv              🔄 Synchronize project's environment files
💀    🚀 syncEnv              🔄 🎉🎉🎉
💀    🚀 syncEnv              🔄 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 532.201114ms
         Current Time: 15:02:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 733.783972ms
         Current Time: 15:02:41
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.52µs
         Current Time: 15:02:41
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🚀 buildDemoBackendI... 🏭 Build image demo-backend:latest
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀    🚀 buildDemoDbImage     🏭 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoBackendI... 🏭 Sending build context to Docker daemon   1.03MB
💀    🚀 buildDemoFrontend... 🏭 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoBackendI... 🏭 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> f7779f873da5
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 Step 7/11 : USER 0
💀    🚀 buildDemoBackendI... 🏭  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoBackendI... 🏭  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 Successfully built 6ea76668c578
💀    🚀 buildDemoDbImage     🏭 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 Docker image demo-frontend built
💀    🚀 buildDemoDbImage     🏭 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> 97fdfef7cb48
💀    🚀 buildDemoBackendI... 🏭 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> bf9c545afbe0
💀    🚀 buildDemoBackendI... 🏭 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> a62a483a9091
💀    🚀 buildDemoBackendI... 🏭 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> db465fe79375
💀    🚀 buildDemoBackendI... 🏭 Successfully built db465fe79375
💀    🚀 buildDemoBackendI... 🏭 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 buildImages          🏭 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 1.423897921s
         Current Time: 15:02:43
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.53437361s
         Current Time: 15:02:43
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.02µs
         Current Time: 15:02:43
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 prepareDemoBackend   🔧 Create venv
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🚀 buildDemoDbImage     🏭 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 Successfully built 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀    🔎 startDemoFrontend... 📗 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🚀 startDemoFrontend... 📗 Error: No such container: demoFrontend
💀 🔥 🔎 startDemoFrontend... 📗 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoDbContainer 🐬 Error: No such container: demoDb
💀 🔥 🚀 startDemoFrontend... 📗 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 Error: No such container: demoDb
💀 🔥 🚀 startDemoDbContainer 🐬 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 fb43e94a4c9a777d5426b465bfd102f7b08b5e81fcafe0c5d20a050869eaffd2
💀    🚀 startDemoDbContainer 🐬 4d5929f4948a5e1ae38204ba539d08a2c26fd0008bd1d6d23f06380b8ad7a315
💀    🚀 prepareDemoBackend   🔧 Activate venv
💀    🚀 prepareDemoBackend   🔧 Install dependencies
💀    🚀 prepareDemoBackend   🔧 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoFrontend... 📗 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 🔎 Waiting docker container 'demoFrontend' healthcheck
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.74 
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.74 Welcome to the Bitnami nginx container
💀    🔎 startDemoFrontend... 📗 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.75 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.75 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀    🚀 startDemoFrontend... 📗 
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.76 
💀 🔥 🚀 startDemoFrontend... 📗 [38;5;6mnginx [38;5;5m08:02:47.77 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 2022/05/14 08:02:47 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoDbContainer 🐬 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.84 
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.85 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.87 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.88 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.89 
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.89 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀    🔎 startDemoDbContainer 🐬 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.94 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.95 [38;5;2mINFO  ==> Initializing mysql database
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.96 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.97 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:47.99 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:48.00 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:48.00 [38;5;2mINFO  ==> Installing database
💀    🔎 startDemoFrontend... 📗 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 check demoFrontend
💀    🔎 startDemoFrontend... 📗 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:53.70 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🔎 startDemoFrontend... 📗 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 📜 Task 'startDemoFrontendContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:55.74 [38;5;2mINFO  ==> Configuring authentication
💀 🔥 🚀 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:55.88 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:55.88 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:02:57.90 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/c1/38/a9fd8c7bb151325d8b3d9108ce791348c84171b5d9f346b0bf0639de603f/coverage-6.3.3-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, idna, urllib3, requests, fastavro, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, py, attrs, toml, iniconfig, pluggy, pyparsing, packaging, pytest, tomli, coverage, pytest-cov, pyasn1, ecdsa, rsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:03:05.92 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:03:05.97 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:03:08.01 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 
💀 🔥 🚀 startDemoDbContainer 🐬 [38;5;6mmysql [38;5;5m08:03:08.07 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.342876Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.344757Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.344765Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.349684Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.546664Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.793746Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.793819Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.819905Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 2022-05-14T08:03:08.820023Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 Database
💀    🔎 startDemoDbContainer 🐬 information_schema
💀    🔎 startDemoDbContainer 🐬 mysql
💀    🔎 startDemoDbContainer 🐬 performance_schema
💀    🔎 startDemoDbContainer 🐬 sample
💀    🔎 startDemoDbContainer 🐬 sys
💀    🔎 startDemoDbContainer 🐬 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.3 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 Prepare
💀    🚀 prepareDemoBackend   🔧 prepare command
💀    🚀 prepareDemoBackend   🔧 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀    🔎 startDemoBackend     ⚡ 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ Activate venv
💀    🚀 startDemoBackend     ⚡ Start
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,207 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,207 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,215 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,215 INFO sqlalchemy.engine.Engine [generated in 0.00021s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,218 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,218 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,220 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,221 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,221 INFO sqlalchemy.engine.Engine [generated in 0.00013s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,226 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ )
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,226 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,253 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,253 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,278 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,278 INFO sqlalchemy.engine.Engine [no key 0.00020s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,298 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,298 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,315 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,316 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,336 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,338 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,338 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,338 INFO sqlalchemy.engine.Engine [cached since 0.1176s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,340 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ )
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,340 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,364 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,365 INFO sqlalchemy.engine.Engine [no key 0.00020s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,398 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,398 INFO sqlalchemy.engine.Engine [no key 0.00032s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,428 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,428 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,464 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,465 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,466 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,466 INFO sqlalchemy.engine.Engine [cached since 0.2453s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,468 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ )
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,468 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,494 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,494 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,516 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,516 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,535 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,535 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,553 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,553 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,575 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,576 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,605 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,605 INFO sqlalchemy.engine.Engine [no key 0.00025s] {}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,625 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,628 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,630 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ FROM users 
💀    🚀 startDemoBackend     ⚡ WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,630 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,632 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,846 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,847 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,848 INFO sqlalchemy.engine.Engine [generated in 0.00025s] {'id': '88717311-b308-41e1-9961-4346108e0330', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$HPpZC1nCoTWkkAU8F/dU/eEjYUDXGiauL/CPcAGmU9W89gxW5391G', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 14, 15, 3, 37, 846214), 'updated_at': datetime.datetime(2022, 5, 14, 15, 3, 37, 847879)}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,849 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,857 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,858 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ FROM users 
💀    🚀 startDemoBackend     ⚡ WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,858 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'pk_1': '88717311-b308-41e1-9961-4346108e0330'}
💀    🚀 startDemoBackend     ⚡ 2022-05-14 15:03:37,860 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ Register auth route handler
💀    🚀 startDemoBackend     ⚡ Register auth event handler
💀    🚀 startDemoBackend     ⚡ Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ Register library route handler
💀    🚀 startDemoBackend     ⚡ Register library event handler
💀    🚀 startDemoBackend     ⚡ Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Started server process [15299]
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ check demoBackend
💀    🔎 startDemoBackend     ⚡ 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 start                🏁 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 55.263564286s
         Current Time: 15:03:38
         Active Process:
           * (PID=15290) ⚡ 'startDemoBackend' service
           * (PID=3519) 📗 'startDemoFrontendContainer' service
           * (PID=3561) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=15290)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=3519)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=3561)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ INFO:     Finished server process [15299]
💀    🚀 startDemoBackend     ⚡ 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 57.366733309s
         Current Time: 15:03:40
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.037µs
         Current Time: 15:03:41
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀    🚀 buildDemoBackendI... 🏭 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🚀 buildDemoFrontend... 🏭 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 Sending build context to Docker daemon  22.02kB
💀    🚀 buildDemoDbImage     🏭 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 562078b73ebf
💀    🚀 buildDemoDbImage     🏭 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoBackendI... 🏭 Sending build context to Docker daemon   1.18MB
💀    🚀 buildDemoDbImage     🏭 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 11c677f847bc
💀    🚀 buildDemoFrontend... 🏭 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 776095918b33
💀    🚀 buildDemoFrontend... 🏭 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 48dc42a93a8a
💀    🚀 buildDemoFrontend... 🏭 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 0beee76410dd
💀    🚀 buildDemoFrontend... 🏭 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 68555ae22bc5
💀    🚀 buildDemoFrontend... 🏭 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 992fa94aa2f2
💀    🚀 buildDemoFrontend... 🏭 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭  ---> 02304e445f6f
💀    🚀 buildDemoFrontend... 🏭 Successfully built 02304e445f6f
💀    🚀 buildDemoBackendI... 🏭 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭  ---> b0401cd9a34b
💀    🚀 buildDemoBackendI... 🏭 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭  ---> Running in 219e928b2c47
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀    🔎 startDemoFrontend... 📗 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🚀 buildDemoBackendI... 🏭 Removing intermediate container 219e928b2c47
💀    🚀 buildDemoBackendI... 🏭  ---> ad0a6bf78252
💀    🚀 buildDemoBackendI... 🏭 Step 8/9 : RUN chmod 755 ./start.sh
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀    🔎 startDemoFrontend... 📗 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 🐳 Logging 'demoFrontend'
💀    🚀 buildDemoBackendI... 🏭  ---> Running in af755cbeb87b
💀    🚀 startDemoDbContainer 🐬 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting docker container 'demoDb' healthcheck
💀    🔎 startDemoFrontend... 📗 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 🔎 Host port '443' is ready
💀    🔎 startDemoDbContainer 🐬 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 Removing intermediate container af755cbeb87b
💀    🚀 buildDemoBackendI... 🏭  ---> b3f08a826115
💀    🚀 buildDemoBackendI... 🏭 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭  ---> Running in ebeda935b579
💀    🚀 buildDemoBackendI... 🏭 Removing intermediate container ebeda935b579
💀    🚀 buildDemoBackendI... 🏭  ---> f4946be8866a
💀    🚀 buildDemoBackendI... 🏭 Successfully built f4946be8866a
💀    🚀 buildDemoBackendI... 🏭 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀 🔥 🔎 startDemoDbContainer 🐬 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoFrontend... 📗 check demoFrontend
💀    🔎 startDemoDbContainer 🐬 Database
💀    🔎 startDemoDbContainer 🐬 information_schema
💀    🔎 startDemoDbContainer 🐬 mysql
💀    🔎 startDemoDbContainer 🐬 performance_schema
💀    🔎 startDemoDbContainer 🐬 sample
💀    🔎 startDemoDbContainer 🐬 sys
💀    🔎 startDemoFrontend... 📗 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀    🔎 startDemoBackendC... ⚡ 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 5e5297a41afb4d1399e777734e78a87fefa2cf214364232a6075cf72408beef6
💀    🚀 startDemoBackendC... ⚡ 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,950 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,950 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,954 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,954 INFO sqlalchemy.engine.Engine [generated in 0.00027s] {}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,961 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,961 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,966 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,967 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,967 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,971 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,974 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,975 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,975 INFO sqlalchemy.engine.Engine [cached since 0.008079s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,978 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,980 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,981 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,981 INFO sqlalchemy.engine.Engine [cached since 0.01377s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,983 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,992 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,995 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ FROM users 
💀    🚀 startDemoBackendC... ⚡ WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,995 INFO sqlalchemy.engine.Engine [generated in 0.00025s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 2022-05-14 08:03:53,998 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ Register auth route handler
💀    🚀 startDemoBackendC... ⚡ Register auth event handler
💀    🚀 startDemoBackendC... ⚡ Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ Register library route handler
💀    🚀 startDemoBackendC... ⚡ Register library event handler
💀    🚀 startDemoBackendC... ⚡ Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ check demoBackend
💀    🔎 startDemoBackendC... ⚡ 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 startContainers      🐳 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 16.165931123s
         Current Time: 15:03:57
         Active Process:
           * (PID=19851) ⚡ 'startDemoBackendContainer' service
           * (PID=17988) 📗 'startDemoFrontendContainer' service
           * (PID=18017) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=19851)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=17988)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=18017)
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 18.270079989s
         Current Time: 15:03:59
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 2.444µs
         Current Time: 15:03:59
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 stopDemoBackendCo... ✋ Stop docker container demoBackend
💀    🚀 stopDemoFrontendC... ✋ Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ Stop docker container demoDb
💀    🚀 stopDemoDbContainer  ✋ demoDb
💀    🚀 stopDemoDbContainer  ✋ 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ demoBackend
💀    🚀 stopDemoBackendCo... ✋ 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀    🚀 stopDemoFrontendC... ✋ demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 stopContainers       ✋ 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.857894605s
         Current Time: 15:04:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.969294957s
         Current Time: 15:04:11
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.075µs
         Current Time: 15:04:11
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontend
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDb
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackend
💀 🔥 🚀 removeDemoDbConta... ❌ Error: No such container: 
💀 🔥 🚀 removeDemoBackend... ❌ Error: No such container: 
💀 🔥 🚀 removeDemoFronten... ❌ Error: No such container: 
💀    🚀 removeDemoDbConta... ❌ Stop docker container demoDb
💀    🚀 removeDemoBackend... ❌ Stop docker container demoBackend
💀    🚀 removeDemoFronten... ❌ Stop docker container demoFrontend
💀    🚀 removeDemoDbConta... ❌ Docker container demoDb stopped
💀    🚀 removeDemoDbConta... ❌ Remove docker container demoDb
💀    🚀 removeDemoFronten... ❌ Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ Remove docker container demoFrontend
💀    🚀 removeDemoBackend... ❌ Docker container demoBackend stopped
💀    🚀 removeDemoBackend... ❌ Remove docker container demoBackend
💀    🚀 removeDemoBackend... ❌ demoBackend
💀    🚀 removeDemoFronten... ❌ demoFrontend
💀    🚀 removeDemoFronten... ❌ 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ Docker container demoFrontend removed
💀    🚀 removeDemoBackend... ❌ 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ Docker container demoBackend removed
💀    🚀 removeDemoDbConta... ❌ demoDb
💀    🚀 removeDemoDbConta... ❌ 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ Docker container demoDb removed
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 removeContainers     ❌ 
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 1.206791396s
         Current Time: 15:04:12
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.317709764s
         Current Time: 15:04:13
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.372µs
         Current Time: 15:04:13
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.564643678s
         Current Time: 15:04:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.675862961s
         Current Time: 15:04:16
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.179µs
         Current Time: 15:04:17
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.980426931s
         Current Time: 15:04:22
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.091244376s
         Current Time: 15:04:22
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.267µs
         Current Time: 15:04:22
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ Hello Human, 
💀    🚀 zrbShowAdv           ☕ Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 
💀    🚀 zrbShowAdv           ☕         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕     
💀    🚀 zrbShowAdv           ☕ Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 makeAppHelmDeploy... 🚢 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 ]
💀    🚀 makeAppHelmDeploy... 🚢 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.133428351s
         Current Time: 15:04:25
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.243777363s
         Current Time: 15:04:25
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.513µs
         Current Time: 15:04:25
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 syncEnv              🔄 Synchronize task environments
💀    🚀 syncEnv              🔄 Synchronize project's environment files
💀    🚀 syncEnv              🔄 🎉🎉🎉
💀    🚀 syncEnv              🔄 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 630.791293ms
         Current Time: 15:04:26
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 832.124583ms
         Current Time: 15:04:26
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.852µs
         Current Time: 15:04:27
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Create virtual environment.
💀    🚀 prepareDemoFronte... 🏁 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 Installing collected packages: six, grpcio, pyyaml, protobuf, semver, dill, pulumi, arpeggio, attrs, parver, urllib3, certifi, charset-normalizer, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 Installing collected packages: protobuf, six, grpcio, semver, dill, pyyaml, pulumi, arpeggio, attrs, parver, idna, charset-normalizer, certifi, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 Installing collected packages: dill, six, grpcio, protobuf, semver, pyyaml, pulumi, arpeggio, attrs, parver, certifi, charset-normalizer, idna, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoBacken... 🏁 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoDbDepl... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"LoadBalancer"}
💀    🚀 prepareDemoFronte... 🏁 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Usage:
💀    🚀 prepareDemoDbDepl... 🏁   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Flags:
💀    🚀 prepareDemoDbDepl... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Chart prepared.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 dependencies.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁     dependencies:
💀    🚀 prepareDemoFronte... 🏁     - name: nginx
💀    🚀 prepareDemoFronte... 🏁       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁     - name: memcached
💀    🚀 prepareDemoFronte... 🏁       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁     dependencies:
💀    🚀 prepareDemoFronte... 🏁     - name: nginx
💀    🚀 prepareDemoFronte... 🏁       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 for this case.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Usage:
💀    🚀 prepareDemoFronte... 🏁   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Aliases:
💀    🚀 prepareDemoFronte... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Available Commands:
💀    🚀 prepareDemoFronte... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Flags:
💀    🚀 prepareDemoFronte... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Global Flags:
💀    🚀 prepareDemoFronte... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 🚧 Chart prepared.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontendDeployment
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDbDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"LoadBalancer"}
💀    🚀 prepareDemoBacken... 🏁 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 dependencies.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁     dependencies:
💀    🚀 prepareDemoBacken... 🏁     - name: nginx
💀    🚀 prepareDemoBacken... 🏁       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁     - name: memcached
💀    🚀 prepareDemoBacken... 🏁       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁     dependencies:
💀    🚀 prepareDemoBacken... 🏁     - name: nginx
💀    🚀 prepareDemoBacken... 🏁       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 for this case.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Usage:
💀    🚀 prepareDemoBacken... 🏁   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Aliases:
💀    🚀 prepareDemoBacken... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Available Commands:
💀    🚀 prepareDemoBacken... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Flags:
💀    🚀 prepareDemoBacken... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Global Flags:
💀    🚀 prepareDemoBacken... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 🚧 Chart prepared.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 
💀    🚀 deployDemoDbDeplo... 🏁 
💀    🚀 deployDemoFronten... 🏁  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoBackend... 🏁 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁  
💀    🚀 deployDemoFronten... 🏁 Resources:
💀    🚀 deployDemoFronten... 🏁     + 4 to create
💀    🚀 deployDemoFronten... 🏁 
💀    🚀 deployDemoFronten... 🏁 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁  
💀    🚀 deployDemoDbDeplo... 🏁 Resources:
💀    🚀 deployDemoDbDeplo... 🏁     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 
💀    🚀 deployDemoDbDeplo... 🏁 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoFronten... 🏁 
💀    🚀 deployDemoDbDeplo... 🏁 
💀    🚀 deployDemoBackend... 🏁  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoFronten... 🏁  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoBackend... 🏁  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁  
💀    🚀 deployDemoBackend... 🏁 Resources:
💀    🚀 deployDemoBackend... 🏁     + 5 to create
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁 Updating (dev):
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoFronten... 🏁  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁  
💀    🚀 deployDemoDbDeplo... 🏁 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁     app: {
💀    🚀 deployDemoDbDeplo... 🏁         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁         ]
💀    🚀 deployDemoDbDeplo... 🏁         resources: {
💀    🚀 deployDemoDbDeplo... 🏁             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁  
💀    🚀 deployDemoDbDeplo... 🏁                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁                             }
💀    🚀 deployDemoDbDeplo... 🏁                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoFronten... 🏁 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoFronten... 🏁     app: {
💀    🚀 deployDemoFronten... 🏁         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁                                                 image          : "demo-db:latest"
💀    🚀 deployDemoFronten... 🏁             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁         ]
💀    🚀 deployDemoFronten... 🏁         resources: {
💀    🚀 deployDemoFronten... 🏁             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoDbDeplo... 🏁                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁                                                 name           : "demo-db"
💀    🚀 deployDemoFronten... 🏁                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁                 id         : "default/demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                                             }
💀    🚀 deployDemoDbDeplo... 🏁                                         ]
💀    🚀 deployDemoFronten... 🏁                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                     annotations       : {
💀    🚀 deployDemoFronten... 🏁                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                             }
💀    🚀 deployDemoFronten... 🏁                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁                         }
💀    🚀 deployDemoDbDeplo... 🏁 
💀    🚀 deployDemoDbDeplo... 🏁                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁                                 annotations: {
💀    🚀 deployDemoFronten... 🏁                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁                     }
💀    🚀 deployDemoDbDeplo... 🏁                     creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoFronten... 🏁                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoFronten... 🏁                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁                     }
💀    🚀 deployDemoFronten... 🏁                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁                                 name       : "demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁                             }
💀    🚀 deployDemoFronten... 🏁                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁                                 selector: {
💀    🚀 deployDemoFronten... 🏁                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoFronten... 🏁                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                                     metadata: {
💀    🚀 deployDemoFronten... 🏁                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁                                             }
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁                                     spec    : {
💀    🚀 deployDemoFronten... 🏁                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoFronten... 🏁                                             [0]: {
💀    🚀 deployDemoFronten... 🏁                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoFronten... 🏁                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁                                                         name : "API_HOST"
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoFronten... 🏁                                                         value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoFronten... 🏁                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoFronten... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁                                                     }
💀    🚀 deployDemoFronten... 🏁                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoFronten... 🏁                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁                                                         }
💀    🚀 deployDemoFronten... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                             }
💀    🚀 deployDemoFronten... 🏁                                                 ]
💀    🚀 deployDemoFronten... 🏁                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                             }
💀    🚀 deployDemoDbDeplo... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                 }
💀    🚀 deployDemoFronten... 🏁                                         ]
💀    🚀 deployDemoFronten... 🏁                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                             }
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoFronten... 🏁                             }
💀    🚀 deployDemoDbDeplo... 🏁                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoFronten... 🏁 
💀    🚀 deployDemoFronten... 🏁                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁                     }
💀    🚀 deployDemoDbDeplo... 🏁                             time       : "2022-05-14T08:04:50Z"
💀    🚀 deployDemoDbDeplo... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                     creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 deployDemoFronten... 🏁                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁                     ]
💀    🚀 deployDemoDbDeplo... 🏁                     name              : "demo-db"
💀    🚀 deployDemoFronten... 🏁                     labels            : {
💀    🚀 deployDemoFronten... 🏁                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁                     resource_version  : "1940"
💀    🚀 deployDemoDbDeplo... 🏁                     uid               : "ce198cfe-0c62-461d-ad69-f825b0ede8e8"
💀    🚀 deployDemoFronten... 🏁                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁                 }
💀    🚀 deployDemoFronten... 🏁                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁                     }
💀    🚀 deployDemoFronten... 🏁                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁                     selector                 : {
💀    🚀 deployDemoFronten... 🏁                         [0]: {
💀    🚀 deployDemoFronten... 🏁                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoFronten... 🏁                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                     }
💀    🚀 deployDemoDbDeplo... 🏁                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁                         }
💀    🚀 deployDemoDbDeplo... 🏁                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoFronten... 🏁                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁                     }
💀    🚀 deployDemoDbDeplo... 🏁                     template                 : {
💀    🚀 deployDemoFronten... 🏁                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁                             labels: {
💀    🚀 deployDemoFronten... 🏁                                             }
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁                             }
💀    🚀 deployDemoFronten... 🏁                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoDbDeplo... 🏁                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁                             containers                      : [
💀    🚀 deployDemoFronten... 🏁                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoDbDeplo... 🏁                                 [0]: {
💀    🚀 deployDemoFronten... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁                                         [0]: {
💀    🚀 deployDemoFronten... 🏁                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁                                             value: "sample"
💀    🚀 deployDemoFronten... 🏁                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                         [1]: {
💀    🚀 deployDemoFronten... 🏁                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁                                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                                     }
💀    🚀 deployDemoFronten... 🏁                                                 }
💀    🚀 deployDemoFronten... 🏁                                             }
💀    🚀 deployDemoDbDeplo... 🏁                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                     }
💀    🚀 deployDemoDbDeplo... 🏁                                             value: "Alch3mist"
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                             }
💀    🚀 deployDemoDbDeplo... 🏁                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁                                             name : "MYSQL_USER"
💀    🚀 deployDemoFronten... 🏁                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                             time       : "2022-05-14T08:04:50Z"
💀    🚀 deployDemoDbDeplo... 🏁                                         [4]: {
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoDbDeplo... 🏁                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁                     ]
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                     name              : "demo-frontend"
💀    🚀 deployDemoDbDeplo... 🏁                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁                     resource_version  : "1938"
💀    🚀 deployDemoDbDeplo... 🏁                                         }
💀    🚀 deployDemoDbDeplo... 🏁                                         [6]: {
💀    🚀 deployDemoFronten... 🏁                     uid               : "b316a0ca-e88e-466d-8d10-caad4dd38532"
💀    🚀 deployDemoFronten... 🏁                 }
💀    🚀 deployDemoDbDeplo... 🏁                                             name : 
💀    🚀 deployDemoFronten... 🏁                 spec       : {
💀    🚀 deployDemoFronten... 🏁                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁                     replicas                 : 1
💀    🚀 deployDemoFronten... 🏁                     revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁                     selector                 : {
💀    🚀 deployDemoFronten... 🏁                         match_labels: {
💀    🚀 deployDemoFronten... 🏁                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                     }
💀    🚀 deployDemoFronten... 🏁                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁                             max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁                     }
💀    🚀 deployDemoFronten... 🏁                     template                 : {
💀    🚀 deployDemoFronten... 🏁                         metadata: {
💀    🚀 deployDemoFronten... 🏁                             labels: {
💀    🚀 deployDemoFronten... 🏁                                 app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                 app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                             }
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                         spec    : {
💀    🚀 deployDemoFronten... 🏁                             containers                      : [
💀    🚀 deployDemoFronten... 🏁                                 [0]: {
💀    🚀 deployDemoFronten... 🏁                                     env                       : [
💀    🚀 deployDemoFronten... 🏁                                         [0]: {
💀    🚀 deployDemoFronten... 🏁                                             name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁                                             value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                         [1]: {
💀    🚀 deployDemoFronten... 🏁                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                         [2]: {
💀    🚀 deployDemoFronten... 🏁                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                         [3]: {
💀    🚀 deployDemoFronten... 🏁                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁                                             value: "1"
💀    🚀 deployDemoFronten... 🏁                                         }
💀    🚀 deployDemoFronten... 🏁                                     ]
💀    🚀 deployDemoFronten... 🏁                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁                                 }
💀    🚀 deployDemoFronten... 🏁                             ]
💀    🚀 deployDemoFronten... 🏁                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁                         }
💀    🚀 deployDemoFronten... 🏁                     }
💀    🚀 deployDemoFronten... 🏁                 }
💀    🚀 deployDemoFronten... 🏁                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁             }
💀    🚀 deployDemoFronten... 🏁             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁                     annotations       : {
💀    🚀 deployDemoFronten... 🏁          
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁  
💀    🚀 deployDemoBackend... 🏁 Outputs:
💀    🚀 deployDemoBackend... 🏁     app: {
💀    🚀 deployDemoBackend... 🏁         ready    : [
💀    🚀 deployDemoBackend... 🏁             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁         ]
💀    🚀 deployDemoBackend... 🏁         resources: {
💀    🚀 deployDemoBackend... 🏁             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁                 metadata   : {
💀    🚀 deployDemoBackend... 🏁                     annotations       : {
💀    🚀 deployDemoBackend... 🏁                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁                             metadata  : {
💀    🚀 deployDemoBackend... 🏁                                 annotations: {
💀    🚀 deployDemoBackend... 🏁                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁                                 }
💀    🚀 deployDemoBackend... 🏁                                 labels     : {
💀    🚀 deployDemoBackend... 🏁                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁                                 }
💀    🚀 deployDemoBackend... 🏁                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁                             }
💀    🚀 deployDemoBackend... 🏁                             spec      : {
💀    🚀 deployDemoBackend... 🏁                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁                                 selector: {
💀    🚀 deployDemoBackend... 🏁                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                 }
💀    🚀 deployDemoBackend... 🏁                                 template: {
💀    🚀 deployDemoBackend... 🏁                                     metadata: {
💀    🚀 deployDemoBackend... 🏁                                         labels: {
💀    🚀 deployDemoBackend... 🏁                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                         }
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                     spec    : {
💀    🚀 deployDemoBackend... 🏁                                         containers        : [
💀    🚀 deployDemoBackend... 🏁                                             [0]: {
💀    🚀 deployDemoBackend... 🏁                                                 env            : [
💀    🚀 deployDemoBackend... 🏁                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁                                                         value: "false"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁                                                         ]
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁                                                         ]
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁                                                         ]
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁                                                         value: (json) []
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁                                                         value: "600"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "621234567890"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [44]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [45]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [46]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [47]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [48]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [49]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                     [50]: {
💀    🚀 deployDemoBackend... 🏁                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                 ]
💀    🚀 deployDemoBackend... 🏁                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁                                                     }
💀    🚀 deployDemoBackend... 🏁                                                 ]
💀    🚀 deployDemoBackend... 🏁                                             }
💀    🚀 deployDemoBackend... 🏁                                         ]
💀    🚀 deployDemoBackend... 🏁                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                 }
💀    🚀 deployDemoBackend... 🏁                             }
💀    🚀 deployDemoBackend... 🏁                         }
💀    🚀 deployDemoBackend... 🏁 
💀    🚀 deployDemoBackend... 🏁                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁                     }
💀    🚀 deployDemoBackend... 🏁                     creation_timestamp: "2022-05-14T08:04:51Z"
💀    🚀 deployDemoBackend... 🏁                     generation        : 1
💀    🚀 deployDemoBackend... 🏁                     labels            : {
💀    🚀 deployDemoBackend... 🏁                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁                     }
💀    🚀 deployDemoBackend... 🏁                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁                         [0]: {
💀    🚀 deployDemoBackend... 🏁                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                 }
💀    🚀 deployDemoBackend... 🏁                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁                                         }
💀    🚀 deployDemoBackend... 🏁                                     }
💀    🚀 deployDemoBackend... 🏁                                     f:t
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 deploy               🏭 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 24.994017802s
         Current Time: 15:04:52
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 25.105034195s
         Current Time: 15:04:52
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.548µs
         Current Time: 15:04:52
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Install pip packages.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoFronte... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 WARNING: You are using pip version 19.2.3, however version 22.1 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"LoadBalancer"}
💀    🚀 prepareDemoFronte... 🏁 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 dependencies.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁     dependencies:
💀    🚀 prepareDemoFronte... 🏁     - name: nginx
💀    🚀 prepareDemoFronte... 🏁       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁     - name: memcached
💀    🚀 prepareDemoFronte... 🏁       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁     dependencies:
💀    🚀 prepareDemoFronte... 🏁     - name: nginx
💀    🚀 prepareDemoFronte... 🏁       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 for this case.
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Usage:
💀    🚀 prepareDemoFronte... 🏁   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Aliases:
💀    🚀 prepareDemoFronte... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Available Commands:
💀    🚀 prepareDemoFronte... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Flags:
💀    🚀 prepareDemoFronte... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Global Flags:
💀    🚀 prepareDemoFronte... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 
💀    🚀 prepareDemoFronte... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 🚧 Chart prepared.
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Prepare chart dependencies.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Usage:
💀    🚀 prepareDemoDbDepl... 🏁   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Flags:
💀    🚀 prepareDemoDbDepl... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 
💀    🚀 prepareDemoDbDepl... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 🚧 Chart prepared.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoFrontendDeployment
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"LoadBalancer"}
💀    🚀 prepareDemoBacken... 🏁 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 dependencies.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁     dependencies:
💀    🚀 prepareDemoBacken... 🏁     - name: nginx
💀    🚀 prepareDemoBacken... 🏁       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁     - name: memcached
💀    🚀 prepareDemoBacken... 🏁       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁     dependencies:
💀    🚀 prepareDemoBacken... 🏁     - name: nginx
💀    🚀 prepareDemoBacken... 🏁       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 for this case.
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Usage:
💀    🚀 prepareDemoBacken... 🏁   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Aliases:
💀    🚀 prepareDemoBacken... 🏁   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Available Commands:
💀    🚀 prepareDemoBacken... 🏁   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Flags:
💀    🚀 prepareDemoBacken... 🏁   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Global Flags:
💀    🚀 prepareDemoBacken... 🏁       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 
💀    🚀 prepareDemoBacken... 🏁 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 🚧 Chart prepared.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀    🚀 destroyDemoFronte... 🏁 Previewing destroy (dev):
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁  
💀    🚀 destroyDemoFronte... 🏁 Outputs:
💀    🚀 destroyDemoFronte... 🏁   - app: {
💀    🚀 destroyDemoFronte... 🏁       - ready    : [
💀    🚀 destroyDemoFronte... 🏁       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁         ]
💀    🚀 destroyDemoFronte... 🏁       - resources: {
💀    🚀 destroyDemoFronte... 🏁           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁                               - selector: {
💀    🚀 destroyDemoFronte... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - template: {
💀    🚀 destroyDemoFronte... 🏁                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                                 ]
💀    🚀 destroyDemoFronte... 🏁                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         ]
💀    🚀 destroyDemoFronte... 🏁                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                                 }
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     ]
💀    🚀 destroyDemoFronte... 🏁                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁                   - resource_version  : "1938"
💀    🚀 destroyDemoFronte... 🏁                   - uid               : "b316a0ca-e88e-466d-8d10-caad4dd38532"
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - spec       : {
💀    🚀 destroyDemoFronte... 🏁                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁                           - labels: {
💀    🚀 destroyDemoFronte... 🏁                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     ]
💀    🚀 destroyDemoFronte... 🏁                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             ]
💀    🚀 destroyDemoFronte... 🏁                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁             }
💀    🚀 destroyDemoFronte... 🏁           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     ]
💀    🚀 destroyDemoFronte... 🏁                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁                   - resource_version  : "1937"
💀    🚀 destroyDemoFronte... 🏁                   - uid               : "3941da09-4853-4a17-adc6-202537adb65a"
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁             }
💀    🚀 destroyDemoFronte... 🏁         }
💀    🚀 destroyDemoFronte... 🏁       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁     }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁 Resources:
💀    🚀 destroyDemoFronte... 🏁     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁  
💀    🚀 destroyDemoDbDepl... 🏁 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁   - app: {
💀    🚀 destroyDemoDbDepl... 🏁       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁         ]
💀    🚀 destroyDemoDbDepl... 🏁       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         ]
💀    🚀 destroyDemoDbDepl... 🏁                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                                 }
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     ]
💀    🚀 destroyDemoDbDepl... 🏁                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁                   - resource_version  : "1940"
💀    🚀 destroyDemoDbDepl... 🏁                   - uid               : "ce198cfe-0c62-461d-ad69-f825b0ede8e8"
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     ]
💀    🚀 destroyDemoDbDepl... 🏁                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁                             ]
💀    🚀 destroyDemoDbDepl... 🏁                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁             }
💀    🚀 destroyDemoDbDepl... 🏁           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     ]
💀    🚀 destroyDemoDbDepl... 🏁                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁                   - resource_version  : "1939"
💀    🚀 destroyDemoDbDepl... 🏁                   - uid               : "71393373-5313-4eee-a96e-0130662704a1"
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁             }
💀    🚀 destroyDemoDbDepl... 🏁         }
💀    🚀 destroyDemoDbDepl... 🏁       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁     }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁 Resources:
💀    🚀 destroyDemoDbDepl... 🏁     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁  
💀    🚀 destroyDemoFronte... 🏁 Outputs:
💀    🚀 destroyDemoFronte... 🏁   - app: {
💀    🚀 destroyDemoFronte... 🏁       - ready    : [
💀    🚀 destroyDemoFronte... 🏁       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁         ]
💀    🚀 destroyDemoFronte... 🏁       - resources: {
💀    🚀 destroyDemoFronte... 🏁           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁                               - selector: {
💀    🚀 destroyDemoFronte... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - template: {
💀    🚀 destroyDemoFronte... 🏁                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                                 ]
💀    🚀 destroyDemoFronte... 🏁                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         ]
💀    🚀 destroyDemoFronte... 🏁                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁                                                         }
💀    🚀 destroyDemoFronte... 🏁                                                     }
💀    🚀 destroyDemoFronte... 🏁                                                 }
💀    🚀 destroyDemoFronte... 🏁                                             }
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     ]
💀    🚀 destroyDemoFronte... 🏁                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁                   - resource_version  : "1938"
💀    🚀 destroyDemoFronte... 🏁                   - uid               : "b316a0ca-e88e-466d-8d10-caad4dd38532"
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - spec       : {
💀    🚀 destroyDemoFronte... 🏁                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁                           - labels: {
💀    🚀 destroyDemoFronte... 🏁                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁                                         }
💀    🚀 destroyDemoFronte... 🏁                                     ]
💀    🚀 destroyDemoFronte... 🏁                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             ]
💀    🚀 destroyDemoFronte... 🏁                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁             }
💀    🚀 destroyDemoFronte... 🏁           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁                     }
💀    🚀 destroyDemoFronte... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁                                     }
💀    🚀 destroyDemoFronte... 🏁                                 }
💀    🚀 destroyDemoFronte... 🏁                             }
💀    🚀 destroyDemoFronte... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoFronte... 🏁                         }
💀    🚀 destroyDemoFronte... 🏁                     ]
💀    🚀 destroyDemoFronte... 🏁                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁                   - resource_version  : "1937"
💀    🚀 destroyDemoFronte... 🏁                   - uid               : "3941da09-4853-4a17-adc6-202537adb65a"
💀    🚀 destroyDemoFronte... 🏁                 }
💀    🚀 destroyDemoFronte... 🏁               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁             }
💀    🚀 destroyDemoFronte... 🏁         }
💀    🚀 destroyDemoFronte... 🏁       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁     }
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁 Resources:
💀    🚀 destroyDemoFronte... 🏁     - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 
💀    🚀 destroyDemoFronte... 🏁 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoFronte... 🏁 hello world
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁  
💀    🚀 destroyDemoDbDepl... 🏁 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁   - app: {
💀    🚀 destroyDemoDbDepl... 🏁       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁         ]
💀    🚀 destroyDemoDbDepl... 🏁       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         ]
💀    🚀 destroyDemoDbDepl... 🏁                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁                                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                                 }
💀    🚀 destroyDemoDbDepl... 🏁                                             }
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     ]
💀    🚀 destroyDemoDbDepl... 🏁                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁                   - resource_version  : "1940"
💀    🚀 destroyDemoDbDepl... 🏁                   - uid               : "ce198cfe-0c62-461d-ad69-f825b0ede8e8"
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁                                         }
💀    🚀 destroyDemoDbDepl... 🏁                                     ]
💀    🚀 destroyDemoDbDepl... 🏁                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             ]
💀    🚀 destroyDemoDbDepl... 🏁                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁             }
💀    🚀 destroyDemoDbDepl... 🏁           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - creation_timestamp: "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁                     }
💀    🚀 destroyDemoDbDepl... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁                                     }
💀    🚀 destroyDemoDbDepl... 🏁                                 }
💀    🚀 destroyDemoDbDepl... 🏁                             }
💀    🚀 destroyDemoDbDepl... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁                           - time       : "2022-05-14T08:04:50Z"
💀    🚀 destroyDemoDbDepl... 🏁                         }
💀    🚀 destroyDemoDbDepl... 🏁                     ]
💀    🚀 destroyDemoDbDepl... 🏁                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁                   - resource_version  : "1939"
💀    🚀 destroyDemoDbDepl... 🏁                   - uid               : "71393373-5313-4eee-a96e-0130662704a1"
💀    🚀 destroyDemoDbDepl... 🏁                 }
💀    🚀 destroyDemoDbDepl... 🏁               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁             }
💀    🚀 destroyDemoDbDepl... 🏁         }
💀    🚀 destroyDemoDbDepl... 🏁       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁     }
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁 Resources:
💀    🚀 destroyDemoDbDepl... 🏁     - 4 deleted
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁 Duration: 1s
💀    🚀 destroyDemoDbDepl... 🏁 
💀    🚀 destroyDemoDbDepl... 🏁 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoDbDepl... 🏁 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 hello world
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁  
💀    🚀 destroyDemoBacken... 🏁 Outputs:
💀    🚀 destroyDemoBacken... 🏁   - app: {
💀    🚀 destroyDemoBacken... 🏁       - ready    : [
💀    🚀 destroyDemoBacken... 🏁       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁         ]
💀    🚀 destroyDemoBacken... 🏁       - resources: {
💀    🚀 destroyDemoBacken... 🏁           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁                               - selector: {
💀    🚀 destroyDemoBacken... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - template: {
💀    🚀 destroyDemoBacken... 🏁                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "621234567890"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                 ]
💀    🚀 destroyDemoBacken... 🏁                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                 ]
💀    🚀 destroyDemoBacken... 🏁                                             }
💀    🚀 destroyDemoBacken... 🏁                                         ]
💀    🚀 destroyDemoBacken... 🏁                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - creation_timestamp: "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁                                             }
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁                                                         }
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                 }
💀    🚀 destroyDemoBacken... 🏁                                             }
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁                           - time       : "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁                   - resource_version  : "1969"
💀    🚀 destroyDemoBacken... 🏁                   - uid               : "c40e21de-81f2-4e75-b94d-27ef11d5e53a"
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - spec       : {
💀    🚀 destroyDemoBacken... 🏁                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁                           - labels: {
💀    🚀 destroyDemoBacken... 🏁                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "false"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                             ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                             ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                             ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "600"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "621234567890"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [44]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [45]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [46]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [47]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [48]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [49]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                   -     [50]: {
💀    🚀 destroyDemoBacken... 🏁                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     ]
💀    🚀 destroyDemoBacken... 🏁                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     ]
💀    🚀 destroyDemoBacken... 🏁                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             ]
💀    🚀 destroyDemoBacken... 🏁                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁             }
💀    🚀 destroyDemoBacken... 🏁           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 ]
💀    🚀 destroyDemoBacken... 🏁                               - selector: {
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - type    : "LoadBalancer"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - creation_timestamp: "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁                                   - f:ports                        : {
💀    🚀 destroyDemoBacken... 🏁                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁                           - time       : "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁                   - resource_version  : "1973"
💀    🚀 destroyDemoBacken... 🏁                   - uid               : "d6cb1a50-051b-4b13-9458-7bc090f126e4"
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - spec       : {
💀    🚀 destroyDemoBacken... 🏁                   - allocate_load_balancer_node_ports: true
💀    🚀 destroyDemoBacken... 🏁                   - cluster_ip                       : "10.104.71.130"
💀    🚀 destroyDemoBacken... 🏁                   - cluster_ips                      : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: "10.104.71.130"
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - external_traffic_policy          : "Cluster"
💀    🚀 destroyDemoBacken... 🏁                   - internal_traffic_policy          : "Cluster"
💀    🚀 destroyDemoBacken... 🏁                   - ip_families                      : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - ip_family_policy                 : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁                   - ports                            : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁                           - node_port  : 31530
💀    🚀 destroyDemoBacken... 🏁                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - selector                         : {
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - session_affinity                 : "None"
💀    🚀 destroyDemoBacken... 🏁                   - type                             : "LoadBalancer"
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - status     : {
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁             }
💀    🚀 destroyDemoBacken... 🏁           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - creation_timestamp: "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁                           - time       : "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁                     ]
💀    🚀 destroyDemoBacken... 🏁                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁                   - resource_version  : "1968"
💀    🚀 destroyDemoBacken... 🏁                   - uid               : "50be05fe-6c69-4da3-a120-1d9674ced5d5"
💀    🚀 destroyDemoBacken... 🏁                 }
💀    🚀 destroyDemoBacken... 🏁               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁             }
💀    🚀 destroyDemoBacken... 🏁         }
💀    🚀 destroyDemoBacken... 🏁       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁     }
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁 Resources:
💀    🚀 destroyDemoBacken... 🏁     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁  
💀    🚀 destroyDemoBacken... 🏁 Outputs:
💀    🚀 destroyDemoBacken... 🏁   - app: {
💀    🚀 destroyDemoBacken... 🏁       - ready    : [
💀    🚀 destroyDemoBacken... 🏁       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁         ]
💀    🚀 destroyDemoBacken... 🏁       - resources: {
💀    🚀 destroyDemoBacken... 🏁           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁                               - selector: {
💀    🚀 destroyDemoBacken... 🏁                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - template: {
💀    🚀 destroyDemoBacken... 🏁                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁                                                         ]
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "621234567890"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                 ]
💀    🚀 destroyDemoBacken... 🏁                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁                                                     }
💀    🚀 destroyDemoBacken... 🏁                                                 ]
💀    🚀 destroyDemoBacken... 🏁                                             }
💀    🚀 destroyDemoBacken... 🏁                                         ]
💀    🚀 destroyDemoBacken... 🏁                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                             }
💀    🚀 destroyDemoBacken... 🏁                         }
💀    🚀 destroyDemoBacken... 🏁 
💀    🚀 destroyDemoBacken... 🏁                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - creation_timestamp: "2022-05-14T08:04:51Z"
💀    🚀 destroyDemoBacken... 🏁                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁                     }
💀    🚀 destroyDemoBacken... 🏁                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                 }
💀    🚀 destroyDemoBacken... 🏁                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                     }
💀    🚀 destroyDemoBacken... 🏁                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁                                             }
💀    🚀 destroyDemoBacken... 🏁                                         }
💀    🚀 destroyDemoBacken... 🏁                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁    
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud
💀    🚀 destroy              ❌ 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 5.411504192s
         Current Time: 15:04:58
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.522318133s
         Current Time: 15:04:58
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/fromZeroToCloud/default.values.yaml'
```````
</details>
<!--endCode-->

