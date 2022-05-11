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
mkdir -p examples/playground/myEndToEndDemo
cd examples/playground/myEndToEndDemo
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
echo "API_HOST=localhost:3000" > demoFrontend/template.env
zaruba please syncEnv

# Add bootstrap
echo 'echo "var apiHost=\"$API_HOST\";" > /opt/bitnami/nginx/html/apiHost.js && /opt/bitnami/scripts/nginx/run.sh' > demoFrontend/bootstrap.sh

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

zaruba please addAppHelmDeployment appDirectory=demoDb
zaruba please addAppHelmDeployment appDirectory=demoBackend
zaruba please addAppHelmDeployment appDirectory=demoFrontend
zaruba please syncEnv

# zaruba please setProjectValue \
#    variableName=defaultImagePrefix \
#    variableValue=gofrendi
#
# zaruba please pushImages

zaruba please setProjectValue \
   variableName=defaultKubeContext \
   variableValue=docker-desktop

zaruba please setProjectValue \
    variableName=pulumiUseLocalBackend \
    variableValue=true

zaruba please deploy
zaruba please destroy
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.847µs
         Current Time: 07:56:46
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 07:56:46.078 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 07:56:46.085 🎉🎉🎉
💀    🚀 initProject          🚧 07:56:46.085 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 149.487857ms
         Current Time: 07:56:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 351.242793ms
         Current Time: 07:56:46
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.363µs
         Current Time: 07:56:46
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:56:46.555 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:56:46.558 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:56:46.558 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:56:46.558 
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:56:46.558         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:56:46.559     
💀    🚀 zrbShowAdv           ☕ 07:56:46.559 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:56:46.559 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:56:46.559   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:56:46.559   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:56:46.559   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:56:46.559 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 07:56:47.027 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 07:56:47.027 Preparing base variables
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Base variables prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing start command
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Start command prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing test command
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Test command prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing check command
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Check command prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.111 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 07:56:47.323 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 07:56:47.329 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 07:56:47.337 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 07:56:47.337 ✅ Validate
💀    🚀 makeMysqlApp         🐬 07:56:47.338 Validate app directory
💀    🚀 makeMysqlApp         🐬 07:56:47.338 Done validating app directory
💀    🚀 makeMysqlApp         🐬 07:56:47.338 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 07:56:47.341 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 07:56:47.341 Validate template locations
💀    🚀 makeMysqlApp         🐬 07:56:47.351 Done validating template locations
💀    🚀 makeMysqlApp         🐬 07:56:47.351 Validate app ports
💀    🚀 makeMysqlApp         🐬 07:56:47.354 Done validating app ports
💀    🚀 makeMysqlApp         🐬 07:56:47.354 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 07:56:47.357 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 07:56:47.357 🚧 Generate
💀    🚀 makeMysqlApp         🐬 07:56:47.357 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 07:56:47.357   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 07:56:47.357 ]
💀    🚀 makeMysqlApp         🐬 07:56:47.357 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 07:56:47.373 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 07:56:47.373 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 07:56:47.373 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 07:56:47.849 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 07:56:47.849 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.008 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.224 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.232 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.238 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.238 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.238 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.238 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.238 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.241 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.241 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.257 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.257 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.26  Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.26  Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.263 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264 ]
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.264 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.302 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.305 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.308 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.462 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.605 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.747 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.947 Checking start
💀    🚀 makeMysqlAppRunner   🐬 07:56:48.952 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.097 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.228 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.231 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.371 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.529 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.74  Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.932 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:49.935 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.108 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.295 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.299 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.484 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.628 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.632 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.776 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.92  Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:56:50.923 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:56:51.083 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:56:51.227 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:56:51.23  🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 07:56:51.23  Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.783880157s
         Current Time: 07:56:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.894965861s
         Current Time: 07:56:51
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.686µs
         Current Time: 07:56:51
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:56:51.586 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:56:51.588         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:56:51.588     
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:56:51.588   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:56:51.588   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:56:51.588   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:56:51.588 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 07:56:52.032 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 07:56:52.032 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 07:56:52.186 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.187 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 07:56:52.187 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.187 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 07:56:52.517 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:56:52.524 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:56:52.534 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 07:56:52.534 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 07:56:52.534 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 07:56:52.534 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 07:56:52.534 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:56:52.538 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:56:52.538 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 07:56:52.549 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 07:56:52.549 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 07:56:52.553 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 07:56:52.553 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558 ]
💀    🚀 makeFastApiApp       ⚡ 07:56:52.558 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 07:56:53.148 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 07:56:53.149 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 07:56:53.149 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 07:56:53.624 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:56:53.624 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.505 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.79  Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.798 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.807 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.808 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.808 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.808 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.808 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.811 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.811 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.831 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.831 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.834 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.834 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.838 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.838 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.838 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.838   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.838   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.839   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.839 ]
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.839 
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.839 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.892 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.896 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:56:54.9   Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.057 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.06  Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.2   Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.35  Checking test
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.353 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.57  Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.811 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:56:55.814 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.032 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.25  Checking start
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.253 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.397 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.543 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.546 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.689 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.831 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:56.972 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.122 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.126 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.338 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.518 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.521 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.7   Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.877 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:56:57.88  Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.072 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.261 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.266 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.456 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.633 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:56:58.829 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:56:59.04  Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:59.244 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:56:59.423 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:56:59.675 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:56:59.844 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:57:00.013 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 07:57:00.013 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 07:57:00.365 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 07:57:00.365 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 07:57:01.195 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.195 Preparing start command
💀    🚀 addFastApiModule     ⚡ 07:57:01.195 Start command prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.195 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 07:57:01.195 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Preparing test command
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Test command prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Preparing check command
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Check command prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.196 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 07:57:01.396 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 07:57:01.402 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 07:57:01.408 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 07:57:01.408 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 07:57:01.408 Validate app directory
💀    🚀 addFastApiModule     ⚡ 07:57:01.409 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 07:57:01.409 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 07:57:01.411 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 07:57:01.411 Validate template locations
💀    🚀 addFastApiModule     ⚡ 07:57:01.42  Done validating template locations
💀    🚀 addFastApiModule     ⚡ 07:57:01.42  Validate app ports
💀    🚀 addFastApiModule     ⚡ 07:57:01.423 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 07:57:01.423 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 07:57:01.426 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 07:57:01.426 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 07:57:01.426 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 07:57:01.426   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 07:57:01.426 ]
💀    🚀 addFastApiModule     ⚡ 07:57:01.426 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 07:57:01.44  🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 07:57:01.44  Registering module
💀    🚀 addFastApiModule     ⚡ 07:57:01.463 Done registering module
💀    🚀 addFastApiModule     ⚡ 07:57:01.464 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 07:57:01.464 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 07:57:01.786 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 07:57:01.786 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 07:57:02.797 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:02.797 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:02.798 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.091 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.098 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.104 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.104 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:57:03.113 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:57:03.113 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:57:03.168 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:57:03.168 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:57:03.234 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:57:03.234 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 07:57:03.365 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 07:57:03.365 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.436 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.742 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.751 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:57:03.759 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:57:03.759 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 07:57:03.759 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 07:57:03.759 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 07:57:03.759 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:57:03.764 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:57:03.764 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 07:57:03.776 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 07:57:03.776 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 07:57:03.78  Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 07:57:03.78  Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784 ]
💀    🚀 addFastApiCrud       ⚡ 07:57:03.784 
💀    🚀 addFastApiCrud       ⚡ 07:57:03.785 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 07:57:03.82  🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 07:57:03.82  Registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:57:03.868 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:57:03.868 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:57:03.951 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:57:03.951 Registering repo
💀    🚀 addFastApiCrud       ⚡ 07:57:04.037 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 07:57:04.038 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 07:57:04.038 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 12.559220257s
         Current Time: 07:57:04
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.760794865s
         Current Time: 07:57:04
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.243µs
         Current Time: 07:57:04
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:04.587 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:57:04.593         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:57:04.593     
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:57:04.593   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:57:04.593   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:57:04.593   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:57:04.593 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 07:57:05.067 🧰 Prepare
💀    🚀 makeNginxApp         📗 07:57:05.067 Preparing base variables
💀    🚀 makeNginxApp         📗 07:57:05.164 Base variables prepared
💀    🚀 makeNginxApp         📗 07:57:05.164 Preparing start command
💀    🚀 makeNginxApp         📗 07:57:05.164 Start command prepared
💀    🚀 makeNginxApp         📗 07:57:05.164 Preparing prepare command
💀    🚀 makeNginxApp         📗 07:57:05.164 Prepare command prepared
💀    🚀 makeNginxApp         📗 07:57:05.164 Preparing test command
💀    🚀 makeNginxApp         📗 07:57:05.164 Test command prepared
💀    🚀 makeNginxApp         📗 07:57:05.164 Preparing migrate command
💀    🚀 makeNginxApp         📗 07:57:05.165 Migrate command prepared
💀    🚀 makeNginxApp         📗 07:57:05.165 Preparing check command
💀    🚀 makeNginxApp         📗 07:57:05.165 Check command prepared
💀    🚀 makeNginxApp         📗 07:57:05.165 Preparing replacement map
💀    🚀 makeNginxApp         📗 07:57:05.369 Add config to replacement map
💀    🚀 makeNginxApp         📗 07:57:05.375 Add env to replacement map
💀    🚀 makeNginxApp         📗 07:57:05.381 Replacement map prepared
💀    🚀 makeNginxApp         📗 07:57:05.381 ✅ Validate
💀    🚀 makeNginxApp         📗 07:57:05.381 Validate app directory
💀    🚀 makeNginxApp         📗 07:57:05.381 Done validating app directory
💀    🚀 makeNginxApp         📗 07:57:05.381 Validate app container volumes
💀    🚀 makeNginxApp         📗 07:57:05.384 Done validating app container volumes
💀    🚀 makeNginxApp         📗 07:57:05.384 Validate template locations
💀    🚀 makeNginxApp         📗 07:57:05.392 Done validating template locations
💀    🚀 makeNginxApp         📗 07:57:05.392 Validate app ports
💀    🚀 makeNginxApp         📗 07:57:05.396 Done validating app ports
💀    🚀 makeNginxApp         📗 07:57:05.396 Validate app crud fields
💀    🚀 makeNginxApp         📗 07:57:05.399 Done validating app crud fields
💀    🚀 makeNginxApp         📗 07:57:05.399 🚧 Generate
💀    🚀 makeNginxApp         📗 07:57:05.399 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 07:57:05.399   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 07:57:05.399 ]
💀    🚀 makeNginxApp         📗 07:57:05.399 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 07:57:05.42  🔩 Integrate
💀    🚀 makeNginxApp         📗 07:57:05.42  🎉🎉🎉
💀    🚀 makeNginxApp         📗 07:57:05.42  Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 07:57:05.841 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 07:57:05.841 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Preparing start command
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Start command prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.927 Preparing test command
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Test command prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Preparing check command
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Check command prepared
💀    🚀 makeNginxAppRunner   📗 07:57:05.928 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 07:57:06.132 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 07:57:06.138 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 07:57:06.144 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 07:57:06.144 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 07:57:06.144 Validate app directory
💀    🚀 makeNginxAppRunner   📗 07:57:06.144 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 07:57:06.144 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 07:57:06.147 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 07:57:06.147 Validate template locations
💀    🚀 makeNginxAppRunner   📗 07:57:06.157 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 07:57:06.157 Validate app ports
💀    🚀 makeNginxAppRunner   📗 07:57:06.16  Done validating app ports
💀    🚀 makeNginxAppRunner   📗 07:57:06.16  Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 07:57:06.163 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 07:57:06.163 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 07:57:06.163 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 07:57:06.163   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 07:57:06.163   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 07:57:06.163 ]
💀    🚀 makeNginxAppRunner   📗 07:57:06.164 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 07:57:06.188 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 07:57:06.191 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:57:06.194 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:57:06.341 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:57:06.491 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:57:06.638 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:57:06.781 Checking start
💀    🚀 makeNginxAppRunner   📗 07:57:06.784 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 07:57:06.931 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:57:07.077 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 07:57:07.08  Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 07:57:07.224 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:57:07.372 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:57:07.52  Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:57:07.662 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 07:57:07.665 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 07:57:07.811 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:57:07.953 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 07:57:07.956 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 07:57:08.127 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:57:08.275 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 07:57:08.279 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 07:57:08.429 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:57:08.581 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 07:57:08.584 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 07:57:08.733 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:57:08.883 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:57:08.887 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 07:57:08.887 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.409279757s
         Current Time: 07:57:08
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.519872118s
         Current Time: 07:57:09
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.88µs
         Current Time: 07:57:09
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:09.247 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 07:57:09.352 Synchronize task environments
💀    🚀 syncEnv              🔄 07:57:09.506 Synchronize project's environment files
💀    🚀 syncEnv              🔄 07:57:09.647 🎉🎉🎉
💀    🚀 syncEnv              🔄 07:57:09.647 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 505.664001ms
         Current Time: 07:57:09
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 707.655592ms
         Current Time: 07:57:09
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.948µs
         Current Time: 07:57:10
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:57:10.098 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:57:10.098 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:57:10.363 Build image demo-db:latest
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoBackendI... 🏭 07:57:10.365 Build image demo-backend:latest
💀    🚀 buildDemoFrontend... 🏭 07:57:10.366 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:57:16.933 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:57:16.934 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 07:57:16.979 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:57:16.979  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:57:16.979 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 07:57:16.979 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoBackendI... 🏭 07:57:16.981 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoFrontend... 🏭 07:57:16.983  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:57:16.983 Step 2/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 07:57:16.983 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:57:16.984 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoDbImage     🏭 07:57:16.985 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:57:16.985 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 07:57:16.985 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 07:57:16.986 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 07:57:16.987 Successfully built 40621c693b70
💀    🚀 buildDemoBackendI... 🏭 07:57:16.988 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 07:57:16.993  ---> caf584a25606
💀    🚀 buildDemoFrontend... 🏭 07:57:16.993 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 07:57:16.993 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 07:57:16.994 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 07:57:16.995  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:16.995  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 07:57:16.995 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 07:57:16.996  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:16.996  ---> 16e3e46a7774
💀    🚀 buildDemoFrontend... 🏭 07:57:16.996 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:57:16.996 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 07:57:16.996 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:17.003  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:57:17.004 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:57:17.01  Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 07:57:17.012 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 07:57:17.012 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 07:57:17.12  
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 7.126855411s
         Current Time: 07:57:17
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 7.237392549s
         Current Time: 07:57:17
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.44µs
         Current Time: 07:57:17
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:57:17.506 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:57:17.506 Links updated
💀    🚀 prepareDemoBackend   🔧 07:57:17.506 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 07:57:17.537 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:57:17.623 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 07:57:17.763 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:57:19.148 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:57:19.149 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 07:57:19.195 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:57:19.195  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:57:19.195 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:57:19.2   Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 07:57:19.202 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:57:19.202 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:57:19.202  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 07:57:19.203 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:19.204  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 07:57:19.205 Successfully built 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 07:57:19.208 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:57:19.209 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:57:19.209 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 prepareDemoBackend   🔧 07:57:19.317 Activate venv
💀    🚀 prepareDemoBackend   🔧 07:57:19.317 Install dependencies
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 07:57:19.445 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoDbContainer 🐬 07:57:19.539 🐳 Retrieve previous log of 'demoDb'
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:57:19.584 🔎 Waiting docker container 'demoFrontend' running status
💀    🚀 startDemoFrontend... 📗 07:57:19.643 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 prepareDemoBackend   🔧 07:57:19.667 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 07:57:19.979   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:57:19.984 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 07:57:20.137   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:57:20.145 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 07:57:20.285   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.017311Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.020297Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.562 [38;5;6mmysql [38;5;5m00:23:23.69 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 [38;5;6mmysql [38;5;5m00:23:29.70 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.020305Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.026741Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.139970Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 [38;5;6mmysql [38;5;5m00:23:29.72 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 [38;5;6mmysql [38;5;5m00:23:31.74 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.324017Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.324064Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.345040Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:20.563 [38;5;6mmysql [38;5;5m00:23:31.80 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:23:32.345499Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:24:08.135554Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:24:10.136993Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 07:57:20.563 2022-05-11T00:24:10.862380Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:57:20.565 🐳 Starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 07:57:20.667 
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.72 
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.72 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.72 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.73 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.73 
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 [38;5;6mnginx [38;5;5m00:23:16.75 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 2022/05/11 00:23:16 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 07:57:20.667 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoFrontend... 📗 07:57:20.668 🐳 Starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:21.535 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:21.535 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited:
        * bash
        * -c
        *    1 | set -e
             2 | . /home/gofrendi/zaruba/zaruba-tasks/_base/run/bash/shellUtil.sh
             3 | _NORMAL='';_BOLD='';_FAINT='';_ITALIC='';_UNDERLINE='';_BLINK_SLOW='';_BLINK_RAPID='';_INVERSE='';_CONCEAL='';_CROSSED_OUT='';_BLACK='';_RED='';_GREEN='';_YELLOW='';_BLUE='';_MAGENTA='';_CYAN='';_WHITE='';_BG_BLACK='';_BG_RED='';_BG_GREEN='';_BG_YELLOW='';_BG_BLUE='';_BG_MAGENTA='';_BG_CYAN='';_BG_WHITE='';_NO_UNDERLINE='';_NO_INVERSE='';_NO_COLOR='';_SKULL='💀';_SUCCESS='🎉';_ERROR='🔥';_START='🏁';_KILL='🔪';_INSPECT='🔎';_RUN='🚀';_EMPTY='  ' 
             4 | CONTAINER_NAME="demoDb"
             5 | if [ -z "${CONTAINER_NAME}" ]
             6 | then
             7 |   echo "${_BOLD}${_RED}containerName is not provided${_NORMAL}"
             8 |   exit 1
             9 | fi 
            10 | DOCKER_IMAGE_NAME="demo-db"
            11 | if [ -z "${DOCKER_IMAGE_NAME}" ]
            12 | then
            13 |   echo "${_BOLD}${_RED}imageName is not provided${_NORMAL}"
            14 |   exit 1
            15 | fi
            16 | 
            17 | 
            18 | if [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
            19 | then
            20 |   echo "🐳 ${_BOLD}${_YELLOW}Container '${CONTAINER_NAME}' is already started${_NORMAL}"
            21 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            22 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            23 | 
            24 | elif [ ! -z $(inspectDocker "container" ".Name" "${CONTAINER_NAME}") ]
            25 | then
            26 |   echo "🐳 ${_BOLD}${_YELLOW}Retrieve previous log of '${CONTAINER_NAME}'${_NORMAL}"
            27 |   sleep 1
            28 |   docker logs --tail 20 "${CONTAINER_NAME}"
            29 |   echo "🐳 ${_BOLD}${_YELLOW}Starting container '${CONTAINER_NAME}'${_NORMAL}"
            30 |   docker start "${CONTAINER_NAME}"
            31 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            32 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            33 | 
            34 | else
            35 |   echo "🐳 ${_BOLD}${_YELLOW}Creating and starting container '${CONTAINER_NAME}'${_NORMAL}"
            36 |   docker run --name "${CONTAINER_NAME}" --hostname "${CONTAINER_NAME}" --network "zaruba"  --shm-size "100m" -e 'MYSQL_DATABASE=sample' -e 'MYSQL_PASSWORD=mysql' -e 'MYSQL_ROOT_PASSWORD=Alch3mist' -e 'MYSQL_USER=mysql' -e 'PYTHONUNBUFFERED=1' -p 3306:3306 -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb/initdb.d:/docker-entrypoint-initdb.d"  --restart no -d "${DOCKER_IMAGE_NAME}" 
            37 | 
            38 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            39 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            40 | 
            41 | fi
            42 | 
            43 | 
            44 | 
            45 | 
            46 | echo 🎉🎉🎉
            47 | echo "📜 ${_BOLD}${_YELLOW}Task 'startDemoDbContainer' is started${_NORMAL}"
            48 | 

exit status 1
💀 🔥 Terminating
💀 🔪 Kill 🔧 'prepareDemoBackend' command (PID=21902)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=24421)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=24464)
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=24465)
💀 🔥 🚀 startDemoFrontend... 📗 07:57:21.882 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 07:57:21.882 Error: failed to start containers: demoFrontend
💀 🔥 📗 'startDemoFrontendContainer' service exited: exit status 1
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 🚀 prepareDemoBackend   🔧 07:57:21.943 ERROR: Operation cancelled by user
💀 🔥 🚀 prepareDemoBackend   🔧 07:57:21.948 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 07:57:21.948 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 Error running 🔧 'prepareDemoBackend' command: exit status 1
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 5.041720263s
         Current Time: 07:57:22
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["start"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.412µs
         Current Time: 07:57:22
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:57:22.71  🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:57:22.71  Links updated
💀    🚀 zrbCreateDockerNe... 🐳 07:57:22.734 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 07:57:22.823 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoFrontend... 🏭 07:57:22.969 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:57:22.969 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 07:57:24.321 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:57:24.323 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 07:57:24.36  Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:57:24.36   ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:57:24.36  Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 07:57:24.365 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 07:57:24.366  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:57:24.366 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:57:24.366  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.366  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:57:24.366 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoDbImage     🏭 07:57:24.367 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:57:24.367 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 07:57:24.368 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369 Step 10/11 : USER 1001
💀    🚀 buildDemoDbImage     🏭 07:57:24.369 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:57:24.369 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:57:24.369  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 07:57:24.37  Successfully built 40621c693b70
💀    🚀 buildDemoBackendI... 🏭 07:57:24.37  Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoFrontend... 🏭 07:57:24.373 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:57:24.375 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:57:24.375 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 07:57:24.375 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 07:57:24.376 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 07:57:24.386 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:57:24.387  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:57:24.387  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:57:24.388 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:57:24.392 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 07:57:24.394 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 07:57:24.394 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:57:24.668 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 07:57:24.687 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 07:57:24.734 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 07:57:24.753 🐳 Retrieve previous log of 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.72 
💀    🚀 startDemoFrontend... 📗 07:57:25.763 
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.72 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.72 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.73 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.73 
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 [38;5;6mnginx [38;5;5m00:23:16.75 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 2022/05/11 00:23:16 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 07:57:25.763 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoFrontend... 📗 07:57:25.764 🐳 Starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  [38;5;6mmysql [38;5;5m00:23:23.69 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  [38;5;6mmysql [38;5;5m00:23:29.70 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.017311Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.020297Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.020305Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.026741Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.139970Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  [38;5;6mmysql [38;5;5m00:23:29.72 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  [38;5;6mmysql [38;5;5m00:23:31.74 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:25.78  [38;5;6mmysql [38;5;5m00:23:31.80 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 07:57:25.78  2022-05-11T00:23:32.324017Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:23:32.324064Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:23:32.345040Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:23:32.345499Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:24:08.135554Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:24:10.136993Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 07:57:25.781 2022-05-11T00:24:10.862380Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:57:25.782 🐳 Starting container 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 07:57:26.724 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 07:57:26.724 Error: failed to start containers: demoFrontend
💀 🔥 📗 'startDemoFrontendContainer' service exited:
        * bash
        * -c
        *    1 | set -e
             2 | . /home/gofrendi/zaruba/zaruba-tasks/_base/run/bash/shellUtil.sh
             3 | _NORMAL='';_BOLD='';_FAINT='';_ITALIC='';_UNDERLINE='';_BLINK_SLOW='';_BLINK_RAPID='';_INVERSE='';_CONCEAL='';_CROSSED_OUT='';_BLACK='';_RED='';_GREEN='';_YELLOW='';_BLUE='';_MAGENTA='';_CYAN='';_WHITE='';_BG_BLACK='';_BG_RED='';_BG_GREEN='';_BG_YELLOW='';_BG_BLUE='';_BG_MAGENTA='';_BG_CYAN='';_BG_WHITE='';_NO_UNDERLINE='';_NO_INVERSE='';_NO_COLOR='';_SKULL='💀';_SUCCESS='🎉';_ERROR='🔥';_START='🏁';_KILL='🔪';_INSPECT='🔎';_RUN='🚀';_EMPTY='  ' 
             4 | CONTAINER_NAME="demoFrontend"
             5 | if [ -z "${CONTAINER_NAME}" ]
             6 | then
             7 |   echo "${_BOLD}${_RED}containerName is not provided${_NORMAL}"
             8 |   exit 1
             9 | fi 
            10 | DOCKER_IMAGE_NAME="demo-frontend"
            11 | if [ -z "${DOCKER_IMAGE_NAME}" ]
            12 | then
            13 |   echo "${_BOLD}${_RED}imageName is not provided${_NORMAL}"
            14 |   exit 1
            15 | fi
            16 | 
            17 | 
            18 | if [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
            19 | then
            20 |   echo "🐳 ${_BOLD}${_YELLOW}Container '${CONTAINER_NAME}' is already started${_NORMAL}"
            21 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            22 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            23 | 
            24 | elif [ ! -z $(inspectDocker "container" ".Name" "${CONTAINER_NAME}") ]
            25 | then
            26 |   echo "🐳 ${_BOLD}${_YELLOW}Retrieve previous log of '${CONTAINER_NAME}'${_NORMAL}"
            27 |   sleep 1
            28 |   docker logs --tail 20 "${CONTAINER_NAME}"
            29 |   echo "🐳 ${_BOLD}${_YELLOW}Starting container '${CONTAINER_NAME}'${_NORMAL}"
            30 |   docker start "${CONTAINER_NAME}"
            31 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            32 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            33 | 
            34 | else
            35 |   echo "🐳 ${_BOLD}${_YELLOW}Creating and starting container '${CONTAINER_NAME}'${_NORMAL}"
            36 |   docker run --name "${CONTAINER_NAME}" --hostname "${CONTAINER_NAME}" --network "zaruba"  --shm-size "100m" -e 'API_HOST=host.docker.internal:3000' -e 'PYTHONUNBUFFERED=1' -p 8080:80 -p 443:443 -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/letsencrypt:/etc/letsencrypt" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/html:/opt/bitnami/nginx/html" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/server_blocks:/opt/bitnami/nginx/conf/server_blocks"  --restart no -d "${DOCKER_IMAGE_NAME}" 
            37 | 
            38 |   echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
            39 | docker logs --since 0m --follow "${CONTAINER_NAME}"
            40 | 
            41 | fi
            42 | 
            43 | 
            44 | 
            45 | 
            46 | echo 🎉🎉🎉
            47 | echo "📜 ${_BOLD}${_YELLOW}Task 'startDemoFrontendContainer' is started${_NORMAL}"
            48 | 

exit status 1
💀 🔥 Terminating
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=27222)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=27223)
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=27197)
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:26.902 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 07:57:26.902 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 5.02695511s
         Current Time: 07:57:27
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["startContainers"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.698µs
         Current Time: 07:57:27
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:57:27.898 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:57:27.898 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoFrontendC... ✋ 07:57:28.244 Docker container demoFrontend is not running
💀    🚀 stopDemoBackendCo... ✋ 07:57:28.245 Docker container demoBackend is not running
💀    🚀 stopDemoDbContainer  ✋ 07:57:28.245 Docker container demoDb is not running
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 07:57:28.351 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 558.632835ms
         Current Time: 07:57:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 670.406333ms
         Current Time: 07:57:28
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.485µs
         Current Time: 07:57:28
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:28.735 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:57:28.737 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:57:28.737 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:57:28.737 
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:57:28.737         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:57:28.738         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:57:28.738     
💀    🚀 zrbShowAdv           ☕ 07:57:28.738 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:57:28.738 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:57:28.738   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:57:28.738   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:57:28.738   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:57:28.738 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.205 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.205 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.439 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.439 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.439 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.439 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.439 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.44  Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.721 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.728 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.734 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.734 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.735 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.735 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.735 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.738 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.738 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.747 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.747 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.751 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.751 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.754 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.792 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.792 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:29.792 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.088 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.088 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.245 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.245 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.246 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.461 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.466 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.472 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.472 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.472 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.472 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.472 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.475 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.475 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.484 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.484 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.486 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.487 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49  ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.49  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.509 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.514 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.517 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.669 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.816 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.819 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:30.965 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:31.114 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:31.116 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:31.264 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:31.265 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:31.265 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.637072234s
         Current Time: 07:57:31
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.748821215s
         Current Time: 07:57:31
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.443µs
         Current Time: 07:57:31
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:31.624 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:57:31.626         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:57:31.626     
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:57:31.626   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:57:31.626   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:57:31.626   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:57:31.626 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.049 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.049 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.715 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.92  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.926 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.932 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.932 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.932 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.932 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.932 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.935 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.935 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.943 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.943 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.947 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.947 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95  ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.95  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.98  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.98  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:32.98  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:33.422 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:33.422 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.311 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.311 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.312 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.595 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.601 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.607 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.607 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.607 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.607 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.607 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.61  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.61  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.618 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.618 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.621 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.621 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.624 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.645 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.649 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.653 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.812 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.97  Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:34.973 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.125 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.278 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.282 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.443 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.443 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:35.443 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.925999117s
         Current Time: 07:57:35
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.03638072s
         Current Time: 07:57:35
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.381µs
         Current Time: 07:57:35
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:35.804 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:57:35.807         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:57:35.807     
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:57:35.807   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:57:35.807   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:57:35.807   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:57:35.807 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.226 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.226 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.326 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.327 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.529 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.536 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.543 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.543 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.543 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.544 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.544 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.547 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.547 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.556 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.556 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.558 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.558 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.561 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.561 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.561 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.561   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.562 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.562 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.594 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.595 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.595 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.932 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:57:36.932 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.034 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.233 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.239 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.246 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.246 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.246 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.246 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.246 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.248 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.248 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.257 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.257 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.261 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.261 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.264 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.265 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.265 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.265   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.265 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.265 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.282 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.286 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.289 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.442 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.598 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.601 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.749 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.9   Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:37.903 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:57:38.089 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:57:38.089 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:57:38.089 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.391416369s
         Current Time: 07:57:38
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.50282166s
         Current Time: 07:57:38
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.278µs
         Current Time: 07:57:38
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:38.456 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 07:57:38.562 Synchronize task environments
💀    🚀 syncEnv              🔄 07:57:38.728 Synchronize project's environment files
💀    🚀 syncEnv              🔄 07:57:38.882 🎉🎉🎉
💀    🚀 syncEnv              🔄 07:57:38.882 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 530.779704ms
         Current Time: 07:57:38
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 731.663473ms
         Current Time: 07:57:39
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.497µs
         Current Time: 07:57:39
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:39.344 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:57:39.455 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:57:39.455 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 217.32825ms
         Current Time: 07:57:39
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.947499ms
         Current Time: 07:57:39
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.35µs
         Current Time: 07:57:39
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:57:39.915 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:57:40.028 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:57:40.028 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 219.364087ms
         Current Time: 07:57:40
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 420.625072ms
         Current Time: 07:57:40
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.23µs
         Current Time: 07:57:40
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoDbDepl... 🏁 07:57:40.499 🚧 Create virtual environment.
💀    🚀 prepareDemoFronte... 🏁 07:57:40.499 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 07:57:40.501 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 07:57:42.292 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:57:42.306 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 07:57:42.368 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:42.578 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:42.582 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:42.625 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:43.422   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:43.436 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:43.544   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:43.558 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:43.662   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 07:57:43.711   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:43.724 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:43.831   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 07:57:43.884 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:43.985   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.031   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.039 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:44.056 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:44.191 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.208   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.213 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:44.269   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:44.279 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:44.382   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:44.398 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.67    Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.692 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:44.723   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:44.746 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.861   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:44.879 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:44.935   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:44.952 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:45.029   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:45.082   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:45.089 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:45.091 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:45.204   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:45.214 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:45.636   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:45.665 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:57:45.751   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:45.806   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:45.811 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:45.812 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:45.815   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:45.876 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:45.964   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:45.969 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:57:46.084   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:46.089 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.147   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.158 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:46.229   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:46.235 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:46.269   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:46.278 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.344   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.358 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:46.373   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:46.382 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:46.458   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:46.471 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.501   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.526 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:46.566   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:46.58  Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:46.702   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:46.728 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.769   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:46.782 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:46.809   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:46.82  Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:46.949   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:46.958 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.039   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.045 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:47.084   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:47.111 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:47.145   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:47.151 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.225   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.23  Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:47.358   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:47.365 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:47.43    Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:47.447 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.505   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:47.52    Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:47.529 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.531 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:47.595   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:57:47.602 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:57:47.718   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.722   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:47.736 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:57:47.763   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:57:47.793 Installing collected packages: six, semver, protobuf, pyyaml, dill, grpcio, pulumi, attrs, arpeggio, parver, idna, certifi, urllib3, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 07:57:47.833 Installing collected packages: pyyaml, six, grpcio, dill, protobuf, semver, pulumi, attrs, arpeggio, parver, certifi, urllib3, charset-normalizer, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 07:57:47.978   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:57:48.061 Installing collected packages: dill, protobuf, pyyaml, six, grpcio, semver, pulumi, arpeggio, attrs, parver, certifi, idna, urllib3, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 07:57:48.328   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 07:57:48.411   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 07:57:48.629   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 07:57:51.456     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 07:57:51.486 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoBacken... 🏁 07:57:51.508 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:57:51.508 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.639     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.653     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 07:57:51.68  Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.695 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 07:57:51.705 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:57:51.705 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:57:51.723 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:57:51.723 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.861 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 07:57:51.861 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.928 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.928 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98      dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98      # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98      dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98        version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98      - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98        repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98      - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98        repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98      - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98        version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98        version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.98  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 07:57:51.98  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 for this case.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Usage:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Flags:
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.981       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:57:51.981       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 07:57:51.982 🚧 Preparation completed.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:57:51.982       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:57:51.983 
💀    🚀 prepareDemoFronte... 🏁 07:57:51.983 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:57:51.983 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 07:57:52.204 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 07:57:52.207 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 07:57:52.306 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 07:57:52.306 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 07:57:52.333 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:57:52.388 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 07:57:52.388 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.427     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:57:52.428 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 
💀    🚀 prepareDemoBacken... 🏁 07:57:52.429 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 07:57:52.653 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 07:57:52.763 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 07:57:54.024 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:57:54.04  Previewing update (dev):
💀    🚀 deployDemoBackend... 🏁 07:57:54.245 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:57:54.393 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:54.393 
💀    🚀 deployDemoBackend... 🏁 07:57:54.606 
💀    🚀 deployDemoFronten... 🏁 07:57:54.751  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:54.765  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 07:57:54.886  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:54.886  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoBackend... 🏁 07:57:54.979  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:57:55.039  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoFronten... 🏁 07:57:55.154  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 07:57:55.156  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.159  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.161  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252  
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.252 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 07:57:55.275  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 07:57:55.275  
💀    🚀 deployDemoFronten... 🏁 07:57:55.275 Resources:
💀    🚀 deployDemoFronten... 🏁 07:57:55.275     + 4 to create
💀    🚀 deployDemoFronten... 🏁 07:57:55.275 
💀    🚀 deployDemoFronten... 🏁 07:57:55.275 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 07:57:55.311  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:57:55.312  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:57:55.319  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:57:55.413  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:57:55.413  
💀    🚀 deployDemoBackend... 🏁 07:57:55.413 Resources:
💀    🚀 deployDemoBackend... 🏁 07:57:55.413     + 5 to create
💀    🚀 deployDemoBackend... 🏁 07:57:55.413 
💀    🚀 deployDemoBackend... 🏁 07:57:55.414 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:57:55.689 
💀    🚀 deployDemoFronten... 🏁 07:57:55.695 
💀    🚀 deployDemoBackend... 🏁 07:57:55.751 
💀    🚀 deployDemoFronten... 🏁 07:57:56.028  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.038  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.097  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.097  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.111  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.192  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.318  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.32   +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.323  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.333  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.338  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.34   +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.349  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.351  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:57:56.353  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 07:57:56.354  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.364  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.364  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoBackend... 🏁 07:57:56.455  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.457  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.463  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.474  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.474  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.483  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:57:56.492  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 07:57:56.495  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 07:57:56.501  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoFronten... 🏁 07:57:56.522  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 07:57:56.522  
💀    🚀 deployDemoFronten... 🏁 07:57:56.524 Outputs:
💀    🚀 deployDemoFronten... 🏁 07:57:56.524     app: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524         ready    : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.524             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524         ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.524         resources: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:57:56.524                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                             spec      : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 selector: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                 template: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                         labels: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                         value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 07:57:56.525                                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                 ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                         ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527 
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                     creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                     labels            : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:57:56.527                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             time       : "2022-05-11T00:57:56Z"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     resource_version  : "14103"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     uid               : "e2ad759f-8e0f-432f-bda9-ef3b2cd1ce9a"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                 spec       : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     replicas                 : 1
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     selector                 : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         match_labels: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                     template                 : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.528                         metadata: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             labels: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                 app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                 app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                         spec    : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             containers                      : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                 [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     env                       : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         [1]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         [2]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         [3]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                             value: "1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.529             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             apiVersion: "v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             kind      : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 07:57:56.529                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                  annotations: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                  }
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                  labels     : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                      helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                  }
💀    🚀 deployDemoFronten... 🏁 07:57:56.53                                  name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     labels            : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         [0]: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             api_version: "v1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                             time       : "2022-05-11T00:57:56Z"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     ]
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     resource_version  : "14102"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                     uid               : "0cfee1e9-aa8c-48ad-b542-1fc22dd10666"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                 }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531                 urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531             }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531         }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531         urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:57:56.531     }
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 Resources:
💀    🚀 deployDemoFronten... 🏁 07:57:56.531     + 4 created
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 Duration: 1s
💀    🚀 deployDemoFronten... 🏁 07:57:56.531 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.539  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.539  
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.54  Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541     app: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541         ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                     creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.541                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             time       : "2022-05-11T00:57:56Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     resource_version  : "14105"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     uid               : "0285043b-87b2-4ffd-90bc-f280d1fd6664"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.542                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.543                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                             time       : "2022-05-11T00:57:56Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                     resource_version  : "14104"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                     uid               : "8d5390be-4194-4552-a142-294f32ca71a1"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544             }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544         }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544     }
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 
💀    🚀 deployDemoDbDeplo... 🏁 07:57:56.544 hello world
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoBackend... 🏁 07:57:56.641  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 07:57:56.641  
💀    🚀 deployDemoBackend... 🏁 07:57:56.643 Outputs:
💀    🚀 deployDemoBackend... 🏁 07:57:56.644     app: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.644         ready    : [
💀    🚀 deployDemoBackend... 🏁 07:57:56.644             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644         ]
💀    🚀 deployDemoBackend... 🏁 07:57:56.644         resources: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.644             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.644                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 annotations: {
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                             }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                             spec      : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 selector: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                 template: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                         labels: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 07:57:56.645                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.646                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.647                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 ]
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                                 ]
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                             }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                         ]
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                 }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                             }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648 
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     generation        : 1
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     labels            : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                         [0]: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                 }
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.648                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                     }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                     f:template               : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                         f:metadata: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                             f:labels: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                             }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                         f:spec    : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                             f:containers                   : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                 k:{"name":"demo-backend"}: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                     f:env                     : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         }
💀    🚀 deployDemoBackend... 🏁 07:57:56.649                                                         k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 deployDemoBackend... 🏁 07:57:56.649          
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 07:57:56.754 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 16.371233371s
         Current Time: 07:57:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 16.49054879s
         Current Time: 07:57:56
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.386µs
         Current Time: 07:57:57
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 07:57:57.186 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.187 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:57:57.189 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 07:57:57.495 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.499 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.502 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.502 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.503 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.508 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.508 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.509 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.509 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.51  Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.513 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.513 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.514 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.515 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.516 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.516 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.517 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.518 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.519 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.519 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.52  Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.52  Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.521 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.522 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.524 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.526 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.527 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.527 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.536 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.538 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.541 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.548 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.552 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.555 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.573 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.583 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.586 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.588 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.589 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.589 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.59  Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.599 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.608 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.61  Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 07:57:57.624 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.624 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 07:57:57.627 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoFronte... 🏁 07:57:57.639 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:57:57.639 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.639 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoBacken... 🏁 07:57:57.643 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:57:57.643 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:57:57.656 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:57:57.656 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.792 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 07:57:57.792 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.848 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 
💀    🚀 prepareDemoFronte... 🏁 07:57:57.849 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  for this case.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Usage:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Flags:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  
💀    🚀 prepareDemoFronte... 🏁 07:57:57.85  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:57:57.851 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.879 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.879 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.923 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.923 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.923 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.924 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 07:57:57.925 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 07:57:58.785 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:57:58.864 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 07:57:58.865 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.923 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 
💀    🚀 prepareDemoBacken... 🏁 07:57:58.924 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 07:57:58.926 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 07:57:59.511 Previewing destroy (dev):
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.613 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 07:57:59.651 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.654  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.654  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.658  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.66   -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.661  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.661  
💀    🚀 destroyDemoFronte... 🏁 07:57:59.662 Outputs:
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663   - app: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663         ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663       - resources: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                               - template: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.663                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                                 ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                         ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.664                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - resource_version  : "14103"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - uid               : "e2ad759f-8e0f-432f-bda9-ef3b2cd1ce9a"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.665                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                             ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.666                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.667                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                   - resource_version  : "14102"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                   - uid               : "0cfee1e9-aa8c-48ad-b542-1fc22dd10666"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668 Resources:
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.668 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.74  
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.742  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.745  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.746  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.748  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.749  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.749  
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.751                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.752                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.753                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - resource_version  : "14105"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - uid               : "0285043b-87b2-4ffd-90bc-f280d1fd6664"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.754                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                             ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.755                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - resource_version  : "14104"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                   - uid               : "8d5390be-4194-4552-a142-294f32ca71a1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.756 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 07:57:59.789 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.789  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.789  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.853 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.857  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.857  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.888  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.888  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.889  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.89   -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.903  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.908  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.908  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909  
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909 Outputs:
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909   - app: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909         ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909       - resources: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.909               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                            - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                  }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                  }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                              }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                            - spec      : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - selector: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                  }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                - template: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                            - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                            - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                          }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                    - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        - containers        : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                - env            : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                        - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                  ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                                - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                              }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                          ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                        - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                      }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                                  }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                              }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                          }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91  
💀    🚀 destroyDemoFronte... 🏁 07:57:59.91                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - resource_version  : "14103"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - uid               : "e2ad759f-8e0f-432f-bda9-ef3b2cd1ce9a"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.911                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                             ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.912                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.913                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                                     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                     ]
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - resource_version  : "14102"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                   - uid               : "0cfee1e9-aa8c-48ad-b542-1fc22dd10666"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914                 }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914             }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914         }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914     }
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 Resources:
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914     - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 07:57:59.914 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoFronte... 🏁 07:57:59.922 hello world
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.957  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.958  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.96   -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.962  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.965  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.972  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.972  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974  
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.974                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.975                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - resource_version  : "14105"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - uid               : "0285043b-87b2-4ffd-90bc-f280d1fd6664"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.976                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                             ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977 
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.977                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   - resource_version  : "14104"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                   - uid               : "8d5390be-4194-4552-a142-294f32ca71a1"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978             }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978         }
💀    🚀 destroyDemoDbDepl... 🏁 07:57:59.978   
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 07:58:00.572 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 07:58:00.657 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.658  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.658  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.662  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.665  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.669  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.671  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.671  
💀    🚀 destroyDemoBacken... 🏁 07:58:00.678 Outputs:
💀    🚀 destroyDemoBacken... 🏁 07:58:00.678   - app: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.678       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679         ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679       - resources: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                               - template: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.679                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                        - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.68                                                -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.681                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.682                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.683                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.684                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.685                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - resource_version  : "14135"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - uid               : "e815dd0a-c9d0-423d-a5a8-3ed7182d6641"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.686                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.687                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                             ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.688                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - resource_version  : "14138"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - uid               : "3f714daa-b867-404f-8c37-c46d045e8280"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - cluster_ip             : "10.101.175.71"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   -     [0]: "10.101.175.71"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.689                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                          }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                      ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                    - selector               : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                        - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                        - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                      }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                    - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                    - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                  }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - status     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                  }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69              }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69            - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                    - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                            - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                  }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                    - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                  }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                                - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.69                              }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - creation_timestamp: "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                           - time       : "2022-05-11T00:57:56Z"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                     ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.691                   - resource_version  : "14134"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692                   - uid               : "5bf01092-a9ef-4fdb-bd8b-7b035fe5dd7c"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692 Resources:
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.692 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 07:58:00.763 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.764  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.764  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.768  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.851  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.852  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.863  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.873  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.874  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.876  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.884  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.891  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.891  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.892  
💀    🚀 destroyDemoBacken... 🏁 07:58:00.895 Outputs:
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896   - app: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896         ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896       - resources: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                               - template: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.896                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.897                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.898                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.899                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                       }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                 -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.9                                                         - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                     }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                                 }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                             }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901                         }
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901 
💀    🚀 destroyDemoBacken... 🏁 07:58:00.901              
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 07:58:01.006 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 3.939919308s
         Current Time: 07:58:01
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.051217524s
         Current Time: 07:58:01
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

