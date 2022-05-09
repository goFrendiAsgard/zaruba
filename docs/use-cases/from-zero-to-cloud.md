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

zaruba please addMysql appDirectory=demoDb

zaruba please addFastApiCrud \
  appDirectory=demoBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["demoDb"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'

zaruba please addNginx \
  appDirectory=demoFrontend \
  appPorts='["80:80"]' \
  appEnvs='{"API_HOST":"localhost:3000"}'

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
         Elapsed Time: 1.153µs
         Current Time: 06:53:10
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 06:53:10.719 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 06:53:10.725 🎉🎉🎉
💀    🚀 initProject          🚧 06:53:10.725 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 118.145439ms
         Current Time: 06:53:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 319.691578ms
         Current Time: 06:53:11
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 2.096µs
         Current Time: 06:53:11
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:53:11.244 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:53:11.249         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:53:11.249     
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:53:11.249   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:53:11.249   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:53:11.249   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:53:11.249 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 06:53:11.72  🧰 Prepare
💀    🚀 makeMysqlApp         🐬 06:53:11.72  Preparing base variables
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Base variables prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Preparing start command
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Start command prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Preparing test command
💀    🚀 makeMysqlApp         🐬 06:53:11.841 Test command prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.842 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 06:53:11.842 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.842 Preparing check command
💀    🚀 makeMysqlApp         🐬 06:53:11.842 Check command prepared
💀    🚀 makeMysqlApp         🐬 06:53:11.842 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 06:53:12.149 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 06:53:12.158 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 06:53:12.168 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 06:53:12.168 ✅ Validate
💀    🚀 makeMysqlApp         🐬 06:53:12.168 Validate app directory
💀    🚀 makeMysqlApp         🐬 06:53:12.168 Done validating app directory
💀    🚀 makeMysqlApp         🐬 06:53:12.168 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 06:53:12.172 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 06:53:12.172 Validate template locations
💀    🚀 makeMysqlApp         🐬 06:53:12.184 Done validating template locations
💀    🚀 makeMysqlApp         🐬 06:53:12.185 Validate app ports
💀    🚀 makeMysqlApp         🐬 06:53:12.188 Done validating app ports
💀    🚀 makeMysqlApp         🐬 06:53:12.189 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 06:53:12.193 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 06:53:12.193 🚧 Generate
💀    🚀 makeMysqlApp         🐬 06:53:12.193 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 06:53:12.193   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 06:53:12.193 ]
💀    🚀 makeMysqlApp         🐬 06:53:12.193 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 06:53:12.211 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 06:53:12.211 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 06:53:12.211 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.654 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.654 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.883 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.883 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.883 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.883 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.883 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:12.884 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.19  Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.2   Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.208 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.208 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.209 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.209 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.209 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.213 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.213 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.236 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.236 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.242 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.242 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.247 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248 ]
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.248 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.297 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.302 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.307 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.485 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.673 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:53:13.855 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.047 Checking start
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.051 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.24  Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.436 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.441 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.656 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:53:14.849 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.036 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.223 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.227 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.423 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.619 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.624 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.811 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.994 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 06:53:15.999 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.2   Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.387 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.392 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.587 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.774 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.779 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 06:53:16.779 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.644058188s
         Current Time: 06:53:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.755718292s
         Current Time: 06:53:16
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.316µs
         Current Time: 06:53:17
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:53:17.176 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:53:17.181         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:53:17.181     
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:53:17.181   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:53:17.181   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:53:17.181   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:53:17.181 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 06:53:17.659 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 06:53:17.66  Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 06:53:17.867 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.868 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 06:53:17.868 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.868 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 06:53:17.868 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:17.868 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 06:53:18.18  Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 06:53:18.189 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 06:53:18.198 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 06:53:18.198 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 06:53:18.199 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 06:53:18.199 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 06:53:18.199 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 06:53:18.204 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 06:53:18.204 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 06:53:18.216 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 06:53:18.216 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 06:53:18.219 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 06:53:18.219 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224 ]
💀    🚀 makeFastApiApp       ⚡ 06:53:18.224 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 06:53:18.83  🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 06:53:18.831 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 06:53:18.831 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 06:53:19.318 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:53:19.318 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.401 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.401 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.402 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.772 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.782 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.791 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.791 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.791 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.791 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.791 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.795 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.795 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.82  Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.82  Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.825 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.825 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83    "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  ]
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.83  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.892 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.897 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 06:53:20.902 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.091 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.095 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.278 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.573 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.576 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.756 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.909 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 06:53:21.913 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.061 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.206 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.209 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.357 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.506 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.509 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.66  Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.808 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:22.957 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.107 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.11  Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.253 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.398 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.401 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.552 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.699 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.703 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.846 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.992 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 06:53:23.995 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.148 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.294 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.452 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.601 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.748 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:53:24.904 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 06:53:25.067 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:25.234 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:53:25.385 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 06:53:25.385 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 06:53:25.875 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 06:53:25.875 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 06:53:26.616 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing start command
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Start command prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing test command
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Test command prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing check command
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Check command prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.617 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 06:53:26.839 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 06:53:26.847 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 06:53:26.855 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 06:53:26.855 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 06:53:26.855 Validate app directory
💀    🚀 addFastApiModule     ⚡ 06:53:26.855 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 06:53:26.855 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 06:53:26.858 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 06:53:26.858 Validate template locations
💀    🚀 addFastApiModule     ⚡ 06:53:26.867 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 06:53:26.867 Validate app ports
💀    🚀 addFastApiModule     ⚡ 06:53:26.87  Done validating app ports
💀    🚀 addFastApiModule     ⚡ 06:53:26.87  Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 06:53:26.873 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 06:53:26.873 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 06:53:26.873 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 06:53:26.873   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 06:53:26.873 ]
💀    🚀 addFastApiModule     ⚡ 06:53:26.874 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 06:53:26.892 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 06:53:26.892 Registering module
💀    🚀 addFastApiModule     ⚡ 06:53:26.917 Done registering module
💀    🚀 addFastApiModule     ⚡ 06:53:26.917 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 06:53:26.918 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 06:53:27.227 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 06:53:27.227 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Preparing start command
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Start command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Preparing test command
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Test command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Preparing check command
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Check command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:27.98  Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.194 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.2   Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.207 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.207 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 06:53:28.217 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 06:53:28.217 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 06:53:28.277 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 06:53:28.277 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 06:53:28.336 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 06:53:28.336 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 06:53:28.431 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 06:53:28.431 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 06:53:28.49  Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.491 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.724 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.73  Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:53:28.737 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 06:53:28.737 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 06:53:28.737 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 06:53:28.737 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 06:53:28.737 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 06:53:28.742 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 06:53:28.742 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 06:53:28.751 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 06:53:28.751 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 06:53:28.754 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 06:53:28.754 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 ]
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 
💀    🚀 addFastApiCrud       ⚡ 06:53:28.757 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 06:53:28.785 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 06:53:28.785 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 06:53:28.82  Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 06:53:28.82  Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 06:53:28.867 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 06:53:28.867 Registering repo
💀    🚀 addFastApiCrud       ⚡ 06:53:28.921 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 06:53:28.922 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 06:53:28.922 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 11.854035346s
         Current Time: 06:53:29
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.05621029s
         Current Time: 06:53:29
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.297µs
         Current Time: 06:53:29
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:53:29.374 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:53:29.376         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:53:29.376     
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:53:29.376   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:53:29.376   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:53:29.376   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:53:29.376 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 06:53:29.802 🧰 Prepare
💀    🚀 makeNginxApp         📗 06:53:29.802 Preparing base variables
💀    🚀 makeNginxApp         📗 06:53:29.892 Base variables prepared
💀    🚀 makeNginxApp         📗 06:53:29.892 Preparing start command
💀    🚀 makeNginxApp         📗 06:53:29.893 Start command prepared
💀    🚀 makeNginxApp         📗 06:53:29.893 Preparing prepare command
💀    🚀 makeNginxApp         📗 06:53:29.893 Prepare command prepared
💀    🚀 makeNginxApp         📗 06:53:29.893 Preparing test command
💀    🚀 makeNginxApp         📗 06:53:29.893 Test command prepared
💀    🚀 makeNginxApp         📗 06:53:29.893 Preparing migrate command
💀    🚀 makeNginxApp         📗 06:53:29.893 Migrate command prepared
💀    🚀 makeNginxApp         📗 06:53:29.893 Preparing check command
💀    🚀 makeNginxApp         📗 06:53:29.893 Check command prepared
💀    🚀 makeNginxApp         📗 06:53:29.893 Preparing replacement map
💀    🚀 makeNginxApp         📗 06:53:30.102 Add config to replacement map
💀    🚀 makeNginxApp         📗 06:53:30.109 Add env to replacement map
💀    🚀 makeNginxApp         📗 06:53:30.115 Replacement map prepared
💀    🚀 makeNginxApp         📗 06:53:30.115 ✅ Validate
💀    🚀 makeNginxApp         📗 06:53:30.115 Validate app directory
💀    🚀 makeNginxApp         📗 06:53:30.115 Done validating app directory
💀    🚀 makeNginxApp         📗 06:53:30.115 Validate app container volumes
💀    🚀 makeNginxApp         📗 06:53:30.118 Done validating app container volumes
💀    🚀 makeNginxApp         📗 06:53:30.118 Validate template locations
💀    🚀 makeNginxApp         📗 06:53:30.127 Done validating template locations
💀    🚀 makeNginxApp         📗 06:53:30.127 Validate app ports
💀    🚀 makeNginxApp         📗 06:53:30.13  Done validating app ports
💀    🚀 makeNginxApp         📗 06:53:30.13  Validate app crud fields
💀    🚀 makeNginxApp         📗 06:53:30.132 Done validating app crud fields
💀    🚀 makeNginxApp         📗 06:53:30.132 🚧 Generate
💀    🚀 makeNginxApp         📗 06:53:30.132 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 06:53:30.132   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 06:53:30.133 ]
💀    🚀 makeNginxApp         📗 06:53:30.133 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 06:53:30.152 🔩 Integrate
💀    🚀 makeNginxApp         📗 06:53:30.152 🎉🎉🎉
💀    🚀 makeNginxApp         📗 06:53:30.152 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 06:53:30.541 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 06:53:30.541 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Preparing start command
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Start command prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Preparing test command
💀    🚀 makeNginxAppRunner   📗 06:53:30.644 Test command prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.645 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 06:53:30.645 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.645 Preparing check command
💀    🚀 makeNginxAppRunner   📗 06:53:30.645 Check command prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.645 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 06:53:30.845 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 06:53:30.851 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 06:53:30.856 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 06:53:30.856 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 06:53:30.856 Validate app directory
💀    🚀 makeNginxAppRunner   📗 06:53:30.856 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 06:53:30.857 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 06:53:30.859 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 06:53:30.859 Validate template locations
💀    🚀 makeNginxAppRunner   📗 06:53:30.872 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 06:53:30.872 Validate app ports
💀    🚀 makeNginxAppRunner   📗 06:53:30.875 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 06:53:30.875 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 06:53:30.878 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 06:53:30.878 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 06:53:30.878 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 06:53:30.878   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 06:53:30.878   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 06:53:30.878 ]
💀    🚀 makeNginxAppRunner   📗 06:53:30.878 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 06:53:30.901 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 06:53:30.905 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 06:53:30.908 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:53:31.056 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:53:31.205 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:53:31.36  Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:53:31.512 Checking start
💀    🚀 makeNginxAppRunner   📗 06:53:31.516 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 06:53:31.669 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:53:31.827 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 06:53:31.831 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 06:53:31.992 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:53:32.181 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:53:32.395 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:53:32.608 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 06:53:32.612 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 06:53:32.795 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:53:32.949 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 06:53:32.953 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 06:53:33.121 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:53:33.391 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 06:53:33.399 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 06:53:33.624 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:53:33.863 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 06:53:33.869 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 06:53:34.039 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:53:34.195 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 06:53:34.199 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 06:53:34.199 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.931523146s
         Current Time: 06:53:34
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.042242251s
         Current Time: 06:53:34
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["80:80"]' -v 'appEnvs={"API_HOST":"localhost:3000"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.353µs
         Current Time: 06:53:34
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:53:34.572 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:53:34.572 Links updated
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 prepareDemoBackend   🔧 06:53:34.576 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 06:53:34.604 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 06:53:34.687 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 06:53:34.831 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 06:53:35.81  Sending build context to Docker daemon  13.31kB
💀    🚀 buildDemoDbImage     🏭 06:53:35.811 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoDbImage     🏭 06:53:35.868 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868 Step 2/6 : USER 0
💀    🚀 buildDemoDbImage     🏭 06:53:35.868  ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> Using cache
💀    🚀 buildDemoDbImage     🏭 06:53:35.868 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:53:35.868  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 06:53:35.869 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 06:53:35.869  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:53:35.869  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 06:53:35.87  Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 06:53:35.87   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:53:35.87   ---> 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 06:53:35.871 Successfully built 736550e2d78d
💀    🚀 buildDemoDbImage     🏭 06:53:35.876 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 06:53:35.879 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 06:53:35.879 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 06:53:35.88  Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 06:53:35.883 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 06:53:35.883 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 06:53:36.272 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 06:53:36.31  🔎 Waiting docker container 'demoDb' running status
💀 🔥 🚀 startDemoFrontend... 📗 06:53:36.318 Error: No such container: demoFrontend
💀 🔥 🔎 startDemoFrontend... 📗 06:53:36.318 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:36.35  Error: No such container: demoDb
💀 🔥 🚀 startDemoFrontend... 📗 06:53:36.358 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 06:53:36.361 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:36.371 Error: No such container: demoDb
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:36.393 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 06:53:36.398 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 06:53:36.446 488e8948cde2b56ff1ccf509216c94056d357112c3990420707a944f7a4fb0dc
💀    🚀 startDemoDbContainer 🐬 06:53:36.478 fa283d0cd3b2659a0ab4111e01951bc1d8ea0773bb68195b81cc582b3c830c2c
💀    🚀 prepareDemoBackend   🔧 06:53:36.819 Activate venv
💀    🚀 prepareDemoBackend   🔧 06:53:36.82  Install dependencies
💀    🚀 prepareDemoBackend   🔧 06:53:37.171 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 06:53:37.406   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:37.416 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 06:53:37.5     Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:37.511 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 06:53:37.572   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoFrontend... 📗 06:53:38.107 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 06:53:38.113 🔎 Waiting docker container 'demoFrontend' healthcheck
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.12 
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.12 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.13 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.13 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.13 
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.153 [38;5;6mnginx [38;5;5m23:53:38.13 [38;5;2mINFO  ==> ** Starting NGINX setup **
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.154 [38;5;6mnginx [38;5;5m23:53:38.15 [38;5;2mINFO  ==> Validating settings in NGINX_* env vars
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.165 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> No custom scripts in /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.168 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> Initializing NGINX
💀    🔎 startDemoFrontend... 📗 06:53:38.17  🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 06:53:38.17  🔎 Waiting for host port: '80'
💀    🔎 startDemoFrontend... 📗 06:53:38.172 🔎 Host port '80' is ready
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.173 realpath: /bitnami/nginx/conf/vhosts: No such file or directory
💀    🚀 startDemoFrontend... 📗 06:53:38.189 
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.189 [38;5;6mnginx [38;5;5m23:53:38.18 [38;5;2mINFO  ==> ** NGINX setup finished! **
💀 🔥 🚀 startDemoFrontend... 📗 06:53:38.2   [38;5;6mnginx [38;5;5m23:53:38.19 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🚀 startDemoDbContainer 🐬 06:53:38.44  🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 06:53:38.441 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.45 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.483 [38;5;6mmysql [38;5;5m23:53:38.47 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.485 [38;5;6mmysql [38;5;5m23:53:38.48 [38;5;2mINFO  ==> Initializing mysql database
💀    🔎 startDemoDbContainer 🐬 06:53:38.486 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 06:53:38.486 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 06:53:38.487 🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.495 [38;5;6mmysql [38;5;5m23:53:38.49 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.5   [38;5;6mmysql [38;5;5m23:53:38.50 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.51  [38;5;6mmysql [38;5;5m23:53:38.50 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.515 [38;5;6mmysql [38;5;5m23:53:38.51 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:38.522 [38;5;6mmysql [38;5;5m23:53:38.52 [38;5;2mINFO  ==> Installing database
💀    🚀 prepareDemoBackend   🔧 06:53:40.333 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 06:53:40.44    Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:40.452 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 06:53:40.535   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:40.547 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 06:53:40.627   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:40.653 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 06:53:40.735   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:40.746 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 06:53:40.877   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:40.983 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🔎 startDemoFrontend... 📗 06:53:41.175 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoFrontend... 📗 06:53:41.334 check demoFrontend
💀    🔎 startDemoFrontend... 📗 06:53:41.34  🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 06:53:41.427   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🔎 startDemoDbContainer 🐬 06:53:41.491 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🚀 prepareDemoBackend   🔧 06:53:41.528 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 06:53:41.67    Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:41.682 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:41.683 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 06:53:41.733 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 06:53:42.07    Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:53:42.335 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 06:53:42.567   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:42.583 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧 06:53:42.653   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:42.664 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 06:53:42.729   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:42.739 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 06:53:42.821   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:42.836 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 06:53:42.903   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:42.935 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 06:53:43.009   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.022 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 06:53:43.263   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.282 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 06:53:43.377   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.39  Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:43.522   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.549 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:43.604 [38;5;6mmysql [38;5;5m23:53:43.60 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 06:53:43.615   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.628 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:53:43.693   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:43.704 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 06:53:43.762   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:53:44.019 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 06:53:44.168   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:44.191 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀    🚀 prepareDemoBackend   🔧 06:53:44.576   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:44.673 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 06:53:44.778   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:44.792 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:44.815 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:44.818 ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: YES)
💀    🚀 prepareDemoBackend   🔧 06:53:44.862   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:44.869 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧 06:53:44.928   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:44.942 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 06:53:45.048   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.074 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 06:53:45.124   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:53:45.273 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🔎 startDemoFrontend... 📗 06:53:45.343 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 06:53:45.343 📜 Task 'startDemoFrontendContainer' is ready
💀    🚀 prepareDemoBackend   🔧 06:53:45.354   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.376 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:45.619 [38;5;6mmysql [38;5;5m23:53:45.61 [38;5;2mINFO  ==> Configuring authentication
💀    🚀 prepareDemoBackend   🔧 06:53:45.648   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.665 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:45.668 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:45.692 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 prepareDemoBackend   🔧 06:53:45.731   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.741 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:45.742 [38;5;6mmysql [38;5;5m23:53:45.74 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:45.746 [38;5;6mmysql [38;5;5m23:53:45.74 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 06:53:45.828   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.838 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:45.909   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:45.942 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:45.998   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.009 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:46.086   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.092 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:46.195   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.204 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:46.29    Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.303 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 06:53:46.747   Using cached https://files.pythonhosted.org/packages/0c/58/25b4d208e0f6f00e19440385f360dc9891f8fa5ab62c11da52eb226fd9cd/coverage-6.3.2-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.763 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:53:46.885   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:46.934 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:53:47.007   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:47.016 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:53:47.103   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:47.115 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 06:53:47.177   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:47.189 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:53:47.307   Using cached https://files.pythonhosted.org/packages/d9/41/d9cfb4410589805cd787f8a82cddd13142d9bf7449d12adf2d05a4a7d633/pyparsing-3.0.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:47.322 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 06:53:47.392   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:53:47.735 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, fastavro, urllib3, idna, requests, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, pluggy, attrs, toml, iniconfig, pyparsing, packaging, py, pytest, tomli, coverage, pytest-cov, pyasn1, rsa, ecdsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:47.769 [38;5;6mmysql [38;5;5m23:53:47.76 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 06:53:47.775   Running setup.py install for avro-python3: started
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:47.999 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:48.001 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 06:53:48.137     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:53:48.531   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:51.161 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:51.164 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:53.779 [38;5;6mmysql [38;5;5m23:53:53.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:53.792 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:53.799 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:53.803 [38;5;6mmysql [38;5;5m23:53:53.80 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:54.283 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:54.284 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:56.819 [38;5;6mmysql [38;5;5m23:53:56.81 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 06:53:56.819 
💀 🔥 🚀 startDemoDbContainer 🐬 06:53:56.839 [38;5;6mmysql [38;5;5m23:53:56.83 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 06:53:57.045 2022-05-08T23:53:57.039634Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 06:53:57.045 2022-05-08T23:53:57.041494Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 06:53:57.045 2022-05-08T23:53:57.041504Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 06:53:57.046 2022-05-08T23:53:57.045770Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 06:53:57.172 2022-05-08T23:53:57.171704Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 06:53:57.351 2022-05-08T23:53:57.350325Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 06:53:57.351 2022-05-08T23:53:57.350378Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 06:53:57.417 2022-05-08T23:53:57.416412Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 06:53:57.417 2022-05-08T23:53:57.416514Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 06:53:57.433 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 Database
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 information_schema
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 mysql
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 performance_schema
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 sample
💀    🔎 startDemoDbContainer 🐬 06:53:57.439 sys
💀    🔎 startDemoDbContainer 🐬 06:53:57.443 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 06:54:01.445 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 06:54:01.445 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 06:54:12.721     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:54:13.978   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 06:54:14.152     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:54:14.681   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 06:54:14.833     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:54:14.888 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.2 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.8 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 06:54:14.948 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 06:54:14.948 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 06:54:14.998 Prepare
💀    🚀 prepareDemoBackend   🔧 06:54:14.998 prepare command
💀    🚀 prepareDemoBackend   🔧 06:54:14.998 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 startDemoBackend     ⚡ 06:54:15.186 Activate venv
💀    🔎 startDemoBackend     ⚡ 06:54:15.186 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 06:54:15.186 Start
💀    🚀 startDemoBackend     ⚡ 06:54:15.586 2022-05-09 06:54:15,586 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 06:54:15.586 2022-05-09 06:54:15,586 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.59  2022-05-09 06:54:15,590 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 06:54:15.59  2022-05-09 06:54:15,590 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.593 2022-05-09 06:54:15,592 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 06:54:15.593 2022-05-09 06:54:15,593 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.594 2022-05-09 06:54:15,594 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:15.595 2022-05-09 06:54:15,594 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:54:15.595 2022-05-09 06:54:15,595 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 2022-05-09 06:54:15,598 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 )
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 
💀    🚀 startDemoBackend     ⚡ 06:54:15.598 2022-05-09 06:54:15,598 INFO sqlalchemy.engine.Engine [no key 0.00012s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.628 2022-05-09 06:54:15,628 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.628 2022-05-09 06:54:15,628 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.648 2022-05-09 06:54:15,647 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 06:54:15.648 2022-05-09 06:54:15,648 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.665 2022-05-09 06:54:15,665 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 06:54:15.665 2022-05-09 06:54:15,665 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.68  2022-05-09 06:54:15,680 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 06:54:15.68  2022-05-09 06:54:15,680 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.696 2022-05-09 06:54:15,696 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:54:15.697 2022-05-09 06:54:15,697 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:15.697 2022-05-09 06:54:15,697 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:54:15.697 2022-05-09 06:54:15,697 INFO sqlalchemy.engine.Engine [cached since 0.1029s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 2022-05-09 06:54:15,699 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 )
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 
💀    🚀 startDemoBackend     ⚡ 06:54:15.699 2022-05-09 06:54:15,699 INFO sqlalchemy.engine.Engine [no key 0.00010s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.719 2022-05-09 06:54:15,719 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 06:54:15.72  2022-05-09 06:54:15,719 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.735 2022-05-09 06:54:15,735 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.735 2022-05-09 06:54:15,735 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.751 2022-05-09 06:54:15,751 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 06:54:15.751 2022-05-09 06:54:15,751 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.77  2022-05-09 06:54:15,769 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:54:15.771 2022-05-09 06:54:15,771 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:15.771 2022-05-09 06:54:15,771 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:54:15.771 2022-05-09 06:54:15,771 INFO sqlalchemy.engine.Engine [cached since 0.1765s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 2022-05-09 06:54:15,773 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 )
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 
💀    🚀 startDemoBackend     ⚡ 06:54:15.773 2022-05-09 06:54:15,773 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.797 2022-05-09 06:54:15,797 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 06:54:15.798 2022-05-09 06:54:15,797 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.814 2022-05-09 06:54:15,814 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 06:54:15.814 2022-05-09 06:54:15,814 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.831 2022-05-09 06:54:15,831 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 06:54:15.831 2022-05-09 06:54:15,831 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.85  2022-05-09 06:54:15,850 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 06:54:15.85  2022-05-09 06:54:15,850 INFO sqlalchemy.engine.Engine [no key 0.00012s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.868 2022-05-09 06:54:15,868 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 06:54:15.868 2022-05-09 06:54:15,868 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.889 2022-05-09 06:54:15,889 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 06:54:15.889 2022-05-09 06:54:15,889 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 06:54:15.91  2022-05-09 06:54:15,910 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:54:15.912 2022-05-09 06:54:15,912 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:15.914 2022-05-09 06:54:15,914 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 06:54:15.914 FROM users 
💀    🚀 startDemoBackend     ⚡ 06:54:15.914 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 06:54:15.914  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 06:54:15.914 2022-05-09 06:54:15,914 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 06:54:15.915 2022-05-09 06:54:15,915 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 06:54:16.11  2022-05-09 06:54:16,109 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:16.111 2022-05-09 06:54:16,111 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 06:54:16.111 2022-05-09 06:54:16,111 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'id': 'd2892409-1ad8-4134-b919-e47025f84746', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$Ol67Uh8gDa.6WnjOlWgtuunfQlePALEWVpv7LGgZI8QWEFIbgD58S', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 9, 6, 54, 16, 109428), 'updated_at': datetime.datetime(2022, 5, 9, 6, 54, 16, 111078)}
💀    🚀 startDemoBackend     ⚡ 06:54:16.112 2022-05-09 06:54:16,112 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:54:16.119 2022-05-09 06:54:16,119 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:54:16.12  2022-05-09 06:54:16,120 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 06:54:16.12  FROM users 
💀    🚀 startDemoBackend     ⚡ 06:54:16.12  WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 06:54:16.12  2022-05-09 06:54:16,120 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {'pk_1': 'd2892409-1ad8-4134-b919-e47025f84746'}
💀    🚀 startDemoBackend     ⚡ 06:54:16.122 2022-05-09 06:54:16,121 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 06:54:16.123 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.131 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 06:54:16.141 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.148 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 06:54:16.148 Register library route handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.148 Register library event handler
💀    🚀 startDemoBackend     ⚡ 06:54:16.148 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 06:54:16.148 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:16.149 INFO:     Started server process [24370]
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:16.149 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:16.149 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:16.149 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 06:54:16.189 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 06:54:16.189 check demoBackend
💀    🔎 startDemoBackend     ⚡ 06:54:16.189 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 06:54:16.189 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 06:54:16.296 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 41.830866847s
         Current Time: 06:54:16
         Active Process:
           * (PID=23989) 📗 'startDemoFrontendContainer' service
           * (PID=24366) ⚡ 'startDemoBackend' service
           * (PID=24018) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=23989)
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=24366)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=24018)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:17.956 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:18.057 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:18.057 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 06:54:18.057 INFO:     Finished server process [24370]
💀    🚀 startDemoBackend     ⚡ 06:54:18.133 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 06:54:18.133 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 43.934285774s
         Current Time: 06:54:18
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.184µs
         Current Time: 06:54:18
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:54:18.663 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:54:18.663 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 06:54:18.692 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 06:54:18.78  Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoFrontend... 🏭 06:54:18.923 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 06:54:18.923 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 06:54:19.725 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 06:54:19.725 Sending build context to Docker daemon  13.82kB
💀    🚀 buildDemoDbImage     🏭 06:54:19.773 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 06:54:19.773 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoBackendI... 🏭 06:54:19.773 Sending build context to Docker daemon  1.179MB
💀    🚀 buildDemoDbImage     🏭 06:54:19.773  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 06:54:19.773 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 06:54:19.773  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 06:54:19.773 Step 2/6 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777  ---> 562078b73ebf
💀    🚀 buildDemoDbImage     🏭 06:54:19.777 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 06:54:19.777 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778 Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:54:19.778  ---> 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 06:54:19.779 Successfully built 736550e2d78d
💀    🚀 buildDemoBackendI... 🏭 06:54:19.779 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoDbImage     🏭 06:54:19.78  🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 06:54:19.78  Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 06:54:19.781  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 06:54:19.781 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoFrontend... 🏭 06:54:19.781 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 06:54:19.781  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:54:19.781  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 06:54:19.781 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 06:54:19.782 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 06:54:19.783 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 06:54:19.783 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 06:54:19.907  ---> 05c43ef8b25a
💀    🚀 buildDemoBackendI... 🏭 06:54:19.907 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 06:54:19.927  ---> Running in f711f857eb81
💀    🚀 buildDemoBackendI... 🏭 06:54:19.984 Removing intermediate container f711f857eb81
💀    🚀 buildDemoBackendI... 🏭 06:54:19.985  ---> 02a88491abae
💀    🚀 buildDemoBackendI... 🏭 06:54:19.985 Step 8/9 : RUN chmod 755 ./start.sh
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 06:54:20.004 🔎 Waiting docker container 'demoDb' running status
💀    🚀 buildDemoBackendI... 🏭 06:54:20.009  ---> Running in 743b958b2854
💀    🚀 startDemoDbContainer 🐬 06:54:20.041 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 06:54:20.041 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 06:54:20.041 🔎 Waiting docker container 'demoDb' healthcheck
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 06:54:20.127 🔎 Waiting docker container 'demoFrontend' running status
💀    🚀 startDemoFrontend... 📗 06:54:20.411 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 06:54:20.411 🐳 Logging 'demoFrontend'
💀    🔎 startDemoDbContainer 🐬 06:54:20.412 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 06:54:20.412 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 06:54:20.414 🔎 Host port '3306' is ready
💀    🔎 startDemoFrontend... 📗 06:54:20.416 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🔎 startDemoFrontend... 📗 06:54:20.449 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 06:54:20.449 🔎 Waiting for host port: '80'
💀    🔎 startDemoFrontend... 📗 06:54:20.45  🔎 Host port '80' is ready
💀    🚀 buildDemoBackendI... 🏭 06:54:20.768 Removing intermediate container 743b958b2854
💀    🚀 buildDemoBackendI... 🏭 06:54:20.768  ---> cb63e33c39a2
💀    🚀 buildDemoBackendI... 🏭 06:54:20.768 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 06:54:20.792  ---> Running in 075b90d23658
💀    🚀 buildDemoBackendI... 🏭 06:54:20.846 Removing intermediate container 075b90d23658
💀    🚀 buildDemoBackendI... 🏭 06:54:20.846  ---> c1bbbe186033
💀    🚀 buildDemoBackendI... 🏭 06:54:20.848 Successfully built c1bbbe186033
💀    🚀 buildDemoBackendI... 🏭 06:54:20.854 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 06:54:20.857 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 06:54:20.857 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoDbContainer 🐬 06:54:23.417 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 06:54:23.453 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 06:54:23.545 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 Database
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 information_schema
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 mysql
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 performance_schema
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 sample
💀    🔎 startDemoDbContainer 🐬 06:54:23.549 sys
💀    🔎 startDemoDbContainer 🐬 06:54:23.554 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 06:54:23.583 check demoFrontend
💀    🔎 startDemoFrontend... 📗 06:54:23.587 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 06:54:27.556 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 06:54:27.556 📜 Task 'startDemoDbContainer' is ready
💀    🔎 startDemoFrontend... 📗 06:54:27.589 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 06:54:27.589 📜 Task 'startDemoFrontendContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 06:54:28.19  🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🔎 startDemoBackendC... ⚡ 06:54:28.218 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:28.22  Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:28.246 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 06:54:28.249 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 06:54:28.293 cac8078a7f494e0adc2e2eb474cbcb4d9f53612857bb5825c556d9ae859b1d2e
💀    🚀 startDemoBackendC... ⚡ 06:54:29.346 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:54:29.349 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 06:54:29.391 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 06:54:29.391 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 06:54:29.392 🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 06:54:29.94  2022-05-08 23:54:29,939 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 06:54:29.94  2022-05-08 23:54:29,939 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.943 2022-05-08 23:54:29,942 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 06:54:29.943 2022-05-08 23:54:29,943 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.949 2022-05-08 23:54:29,948 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 06:54:29.949 2022-05-08 23:54:29,949 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.953 2022-05-08 23:54:29,952 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:54:29.954 2022-05-08 23:54:29,953 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:54:29.954 2022-05-08 23:54:29,953 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.956 2022-05-08 23:54:29,956 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:54:29.959 2022-05-08 23:54:29,958 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:54:29.959 2022-05-08 23:54:29,959 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:54:29.959 2022-05-08 23:54:29,959 INFO sqlalchemy.engine.Engine [cached since 0.005577s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.961 2022-05-08 23:54:29,961 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:54:29.963 2022-05-08 23:54:29,963 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:54:29.964 2022-05-08 23:54:29,963 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:54:29.964 2022-05-08 23:54:29,963 INFO sqlalchemy.engine.Engine [cached since 0.01033s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.966 2022-05-08 23:54:29,965 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:54:29.971 2022-05-08 23:54:29,970 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:54:29.973 2022-05-08 23:54:29,973 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 06:54:29.973 FROM users 
💀    🚀 startDemoBackendC... ⚡ 06:54:29.973 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 06:54:29.973  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 06:54:29.973 2022-05-08 23:54:29,973 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 06:54:29.976 2022-05-08 23:54:29,976 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 06:54:29.979 Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 06:54:29.988 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 06:54:30.004 Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 06:54:30.011 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 06:54:30.011 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 06:54:30.011 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 06:54:30.011 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 06:54:30.011 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:30.011 INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:30.011 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:30.011 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 06:54:30.012 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 06:54:32.395 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:54:32.521 check demoBackend
💀    🔎 startDemoBackendC... ⚡ 06:54:32.525 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:54:33.526 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 06:54:33.526 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 06:54:33.633 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 15.07661422s
         Current Time: 06:54:33
         Active Process:
           * (PID=26007) ⚡ 'startDemoBackendContainer' service
           * (PID=25843) 🐬 'startDemoDbContainer' service
           * (PID=25905) 📗 'startDemoFrontendContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=26007)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=25843)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=25905)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 17.179950135s
         Current Time: 06:54:35
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.934µs
         Current Time: 06:54:35
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:54:35.995 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:54:35.995 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoDbContainer  ✋ 06:54:36.356 Stop docker container demoDb
💀    🚀 stopDemoBackendCo... ✋ 06:54:36.358 Stop docker container demoBackend
💀    🚀 stopDemoFrontendC... ✋ 06:54:36.36  Stop docker container demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 06:54:37.075 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 06:54:37.077 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 06:54:37.077 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀    🚀 stopDemoDbContainer  ✋ 06:54:40.573 demoDb
💀    🚀 stopDemoDbContainer  ✋ 06:54:40.575 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 06:54:40.575 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ 06:54:46.965 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 06:54:46.967 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 06:54:46.967 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 06:54:47.074 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.184830321s
         Current Time: 06:54:47
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.295730897s
         Current Time: 06:54:47
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 3.132µs
         Current Time: 06:54:47
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:47.449 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:54:47.452         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:54:47.452     
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:54:47.452   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:54:47.452   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:54:47.452   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:54:47.452 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:47.892 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:47.892 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.044 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.264 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.271 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.277 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.277 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.277 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.278 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.278 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.28  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.281 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.29  Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.29  Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.293 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.293 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.296 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.327 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.327 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.327 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.619 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.619 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.776 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.991 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:48.999 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.006 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.006 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.006 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.006 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.006 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.009 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.009 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.019 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.019 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.022 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.022 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.025 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.045 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.049 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.052 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.207 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.364 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.368 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.526 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.684 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.687 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.844 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.844 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:49.844 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.501408343s
         Current Time: 06:54:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.612203334s
         Current Time: 06:54:50
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.417µs
         Current Time: 06:54:50
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:50.217 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:54:50.22          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:54:50.22      
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:54:50.22    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:54:50.22    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:54:50.22    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:54:50.22  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:50.655 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:50.655 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.384 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.384 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.385 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.609 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.616 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.622 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.622 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.622 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.622 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.622 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.626 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.626 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.635 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.635 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.638 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.638 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.641 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.641 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.641 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.641   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.642 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.642 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.674 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.675 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:51.675 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.059 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.06  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.794 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.795 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.795 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:52.795 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.024 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.032 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.039 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.039 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.039 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.039 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.039 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.042 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.042 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.051 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.051 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.054 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.054 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.057 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.076 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.081 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.085 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.241 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.404 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.407 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.573 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.734 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.738 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.905 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.905 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:53.905 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.794894558s
         Current Time: 06:54:54
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.906210905s
         Current Time: 06:54:54
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.23µs
         Current Time: 06:54:54
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:54.281 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:54:54.284         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:54:54.284     
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:54:54.284   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:54:54.284   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:54:54.284   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:54:54.284 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.718 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.718 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.81  Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:54.811 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.028 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.036 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.042 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.042 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.042 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.042 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.042 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.045 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.045 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.054 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.054 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.057 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.057 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06  ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.06  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.092 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.092 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.092 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.522 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.522 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.617 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.831 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.838 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.845 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.845 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.845 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.845 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.845 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.848 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.848 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.858 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.858 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.861 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.861 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.865 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.884 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.887 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:55.89  Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.054 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.216 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.22  Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.385 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.552 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.556 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.726 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.726 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:54:56.726 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.5501392s
         Current Time: 06:54:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.660913026s
         Current Time: 06:54:56
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.003µs
         Current Time: 06:54:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:57.102 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 06:54:57.208 Synchronize task environments
💀    🚀 syncEnv              🔄 06:54:57.386 Synchronize project's environment files
💀    🚀 syncEnv              🔄 06:54:57.553 🎉🎉🎉
💀    🚀 syncEnv              🔄 06:54:57.553 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 556.781648ms
         Current Time: 06:54:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 757.994936ms
         Current Time: 06:54:57
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.471µs
         Current Time: 06:54:58
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:58.022 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 06:54:58.134 🎉🎉🎉
💀    🚀 setProjectValue      🔗 06:54:58.134 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 217.075016ms
         Current Time: 06:54:58
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 417.433775ms
         Current Time: 06:54:58
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.325µs
         Current Time: 06:54:58
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:54:58.613 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 06:54:58.724 🎉🎉🎉
💀    🚀 setProjectValue      🔗 06:54:58.724 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.452599ms
         Current Time: 06:54:58
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.302602ms
         Current Time: 06:54:59
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.333µs
         Current Time: 06:54:59
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 06:54:59.205 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 06:54:59.205 🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 06:54:59.207 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 06:55:01.053 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:01.094 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 06:55:01.128 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 06:55:01.348 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:01.368 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:01.393 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:02.014   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:02.033 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.041   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.064 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:02.089   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:02.105 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:02.215   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.238   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 06:55:02.313   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 06:55:02.503 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.521 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:02.57  Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.697   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.717 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.779   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.783 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.85    Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:02.858 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:02.872   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:02.881   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:02.896 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:02.904 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:02.984   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:02.988   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:02.992 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:02.993 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:03.055   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:03.06  Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:03.105   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:03.122 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:03.194   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:03.204 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:03.283   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:03.29  Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:03.542   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:03.615 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:03.76    Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:03.835 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:03.934   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:03.945   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:03.957 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:55:03.958 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:55:03.998   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.024   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.038 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.057   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.072 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.075 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.1     Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.111 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.127   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.139 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.14    Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.153 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.213   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.229 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.24    Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.256 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.266   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.281 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.286   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.316 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.317   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.346 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.363   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.374   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.387 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.395 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.403   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.417 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.444   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.448   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.455 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.458 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.48    Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.487 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.53    Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.531   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.539 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.541 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.556   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.563 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:55:04.61    Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.617 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.618   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.644 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.666   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.688 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.699   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.706   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:55:04.725 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:55:04.754   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:55:04.757 Installing collected packages: pyyaml, semver, six, grpcio, protobuf, dill, pulumi, attrs, arpeggio, parver, certifi, charset-normalizer, urllib3, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 06:55:04.789   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:55:04.822 Installing collected packages: protobuf, semver, pyyaml, dill, six, grpcio, pulumi, attrs, arpeggio, parver, certifi, charset-normalizer, urllib3, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 06:55:04.871 Installing collected packages: protobuf, semver, six, grpcio, dill, pyyaml, pulumi, attrs, arpeggio, parver, idna, certifi, urllib3, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 06:55:05.37    Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 06:55:05.44    Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 06:55:05.443   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 06:55:06.684     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 06:55:06.731 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoBacken... 🏁 06:55:06.739     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:55:06.759 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:55:06.759 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 06:55:06.786 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 06:55:06.797     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀 🔥 🚀 prepareDemoBacken... 🏁 06:55:06.812 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 06:55:06.812 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 06:55:06.847 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 06:55:06.88  WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 06:55:06.88  You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:06.976 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 06:55:06.976 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.027 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.028 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:07.029 🚧 Preparation completed.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.03  🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 06:55:07.03  🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.077     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 for this case.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 Usage:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078 Aliases:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.078   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 Flags:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 06:55:07.079       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 06:55:07.08    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 06:55:07.08        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 06:55:07.08        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 06:55:07.081       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 06:55:07.081 
💀    🚀 prepareDemoFronte... 🏁 06:55:07.081 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:07.293 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:07.293    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:07.293 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:07.293 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:07.307 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:07.307    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:07.307 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:07.308 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 06:55:07.395 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 06:55:07.415 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 06:55:07.759 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 06:55:07.814 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 06:55:07.814 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 for this case.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Usage:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Aliases:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859 Flags:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.859   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86  
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86  Global Flags:
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86  
💀    🚀 prepareDemoBacken... 🏁 06:55:07.86  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 06:55:07.861 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:08.214 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:08.214    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:08.214 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:08.214 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 06:55:08.303 Created stack 'dev'
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:08.356 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:08.356    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:08.357 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:08.383 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:08.383    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:08.383 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 deployDemoFronten... 🏁 06:55:09.15  Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 06:55:09.153 Previewing update (dev):
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:09.52  warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:09.52     $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:09.52  or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 deployDemoDbDeplo... 🏁 06:55:09.609 
💀    🚀 deployDemoFronten... 🏁 06:55:09.609 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:09.924  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 06:55:09.931  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:09.989  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 06:55:09.99   +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoBackend... 🏁 06:55:10.098 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.206  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.209  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 06:55:10.225  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 06:55:10.227  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319  
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.319 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 06:55:10.345  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 06:55:10.345  
💀    🚀 deployDemoFronten... 🏁 06:55:10.345 Resources:
💀    🚀 deployDemoFronten... 🏁 06:55:10.345     + 4 to create
💀    🚀 deployDemoFronten... 🏁 06:55:10.345 
💀    🚀 deployDemoFronten... 🏁 06:55:10.345 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 06:55:10.48  
💀    🚀 deployDemoDbDeplo... 🏁 06:55:10.895 
💀    🚀 deployDemoFronten... 🏁 06:55:10.932 
💀    🚀 deployDemoBackend... 🏁 06:55:11.082  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 06:55:11.145  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.315  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 06:55:11.32   +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.391  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:55:11.394  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 06:55:11.458  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:55:11.459  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:55:11.464  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:55:11.595  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 06:55:11.595  
💀    🚀 deployDemoBackend... 🏁 06:55:11.595 Resources:
💀    🚀 deployDemoBackend... 🏁 06:55:11.595     + 5 to create
💀    🚀 deployDemoBackend... 🏁 06:55:11.595 
💀    🚀 deployDemoBackend... 🏁 06:55:11.595 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.654  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.657  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:55:11.67   +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.672  +  kubernetes:core/v1:ServiceAccount default/demo-db creating Retry #0; creation failed: serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.672  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.673  +  kubernetes:core/v1:ServiceAccount default/demo-db creating error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.673  +  kubernetes:core/v1:ServiceAccount default/demo-db **creating failed** error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoFronten... 🏁 06:55:11.677  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.685  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.685  +  pulumi:pulumi:Stack demoDbDeployment-dev creating error: update failed
💀    🚀 deployDemoFronten... 🏁 06:55:11.689  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.689  +  pulumi:pulumi:Stack demoDbDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.689  +  kubernetes:helm.sh/v3:Chart demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.689  
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.689 Diagnostics:
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69    pulumi:pulumi:Stack (demoDbDeployment-dev):
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69      error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69   
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69    kubernetes:core/v1:ServiceAccount (default/demo-db):
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69      error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69   
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69  Resources:
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69      + 3 created
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69  
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69  Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 06:55:11.69  
💀    🚀 deployDemoFronten... 🏁 06:55:11.692  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating Retry #0; creation failed: serviceaccounts "demo-frontend" already exists
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:11.692 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:11.692    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:55:11.692 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 deployDemoFronten... 🏁 06:55:11.693  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 06:55:11.693  +  kubernetes:core/v1:ServiceAccount default/demo-frontend **creating failed** error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀 🔥 Error running 🏁 'deployDemoDbDeployment' command:
        * bash
        * -c
        *    1 | set -e
             2 | . /home/gofrendi/zaruba/zaruba-tasks/_base/run/bash/shellUtil.sh
             3 | _NORMAL='';_BOLD='';_FAINT='';_ITALIC='';_UNDERLINE='';_BLINK_SLOW='';_BLINK_RAPID='';_INVERSE='';_CONCEAL='';_CROSSED_OUT='';_BLACK='';_RED='';_GREEN='';_YELLOW='';_BLUE='';_MAGENTA='';_CYAN='';_WHITE='';_BG_BLACK='';_BG_RED='';_BG_GREEN='';_BG_YELLOW='';_BG_BLUE='';_BG_MAGENTA='';_BG_CYAN='';_BG_WHITE='';_NO_UNDERLINE='';_NO_INVERSE='';_NO_COLOR='';_SKULL='💀';_SUCCESS='🎉';_ERROR='🔥';_START='🏁';_KILL='🔪';_INSPECT='🔎';_RUN='🚀';_EMPTY='  '
             4 | 
             5 | 
             6 | mkdir -p ./pulumiLock
             7 | PULUMI_BACKEND_URL="file://./pulumiLock"
             8 | pulumi stack select "dev" || pulumi stack init "dev" 
             9 | pulumi up -y
            10 | echo hello world
            11 | 
            12 | 
            13 | 
            14 | 
exit status 255
💀    🚀 deployDemoFronten... 🏁 06:55:11.702  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 06:55:11.702  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating error: update failed
💀    🚀 deployDemoFronten... 🏁 06:55:11.709  +  pulumi:pulumi:Stack demoFrontendDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoFronten... 🏁 06:55:11.709  +  kubernetes:helm.sh/v3:Chart demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 06:55:11.709  
💀    🚀 deployDemoFronten... 🏁 06:55:11.709 Diagnostics:
💀    🚀 deployDemoFronten... 🏁 06:55:11.709   pulumi:pulumi:Stack (demoFrontendDeployment-dev):
💀    🚀 deployDemoFronten... 🏁 06:55:11.709     error: update failed
💀    🚀 deployDemoFronten... 🏁 06:55:11.71   
💀    🚀 deployDemoFronten... 🏁 06:55:11.71    kubernetes:core/v1:ServiceAccount (default/demo-frontend):
💀    🚀 deployDemoFronten... 🏁 06:55:11.71      error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 06:55:11.71   
💀    🚀 deployDemoFronten... 🏁 06:55:11.71  Resources:
💀    🚀 deployDemoFronten... 🏁 06:55:11.71      + 3 created
💀    🚀 deployDemoFronten... 🏁 06:55:11.71  
💀    🚀 deployDemoFronten... 🏁 06:55:11.71  Duration: 1s
💀    🚀 deployDemoFronten... 🏁 06:55:11.71  
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:11.711 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:11.711    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoFronten... 🏁 06:55:11.711 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 Error running 🏁 'deployDemoFrontendDeployment' command:
        * bash
        * -c
        *    1 | set -e
             2 | . /home/gofrendi/zaruba/zaruba-tasks/_base/run/bash/shellUtil.sh
             3 | _NORMAL='';_BOLD='';_FAINT='';_ITALIC='';_UNDERLINE='';_BLINK_SLOW='';_BLINK_RAPID='';_INVERSE='';_CONCEAL='';_CROSSED_OUT='';_BLACK='';_RED='';_GREEN='';_YELLOW='';_BLUE='';_MAGENTA='';_CYAN='';_WHITE='';_BG_BLACK='';_BG_RED='';_BG_GREEN='';_BG_YELLOW='';_BG_BLUE='';_BG_MAGENTA='';_BG_CYAN='';_BG_WHITE='';_NO_UNDERLINE='';_NO_INVERSE='';_NO_COLOR='';_SKULL='💀';_SUCCESS='🎉';_ERROR='🔥';_START='🏁';_KILL='🔪';_INSPECT='🔎';_RUN='🚀';_EMPTY='  '
             4 | 
             5 | 
             6 | mkdir -p ./pulumiLock
             7 | PULUMI_BACKEND_URL="file://./pulumiLock"
             8 | pulumi stack select "dev" || pulumi stack init "dev" 
             9 | pulumi up -y
            10 | echo hello world
            11 | 
            12 | 
            13 | 
            14 | 
exit status 255
💀 🔥 Terminating
💀 🔪 Kill 🏁 'deployDemoBackendDeployment' command (PID=3396)
💀    🚀 deployDemoBackend... 🏁 06:55:11.982 
💀    🚀 deployDemoBackend... 🏁 06:55:12.097     pulumi:pulumi:Stack demoBackendDeployment-dev  error: update canceled
💀    🚀 deployDemoBackend... 🏁 06:55:12.099     pulumi:pulumi:Stack demoBackendDeployment-dev **failed** 1 error
💀    🚀 deployDemoBackend... 🏁 06:55:12.099  
💀    🚀 deployDemoBackend... 🏁 06:55:12.099 Diagnostics:
💀    🚀 deployDemoBackend... 🏁 06:55:12.099   pulumi:pulumi:Stack (demoBackendDeployment-dev):
💀    🚀 deployDemoBackend... 🏁 06:55:12.099     error: update canceled
💀    🚀 deployDemoBackend... 🏁 06:55:12.099  
💀    🚀 deployDemoBackend... 🏁 06:55:12.099 Resources:
💀    🚀 deployDemoBackend... 🏁 06:55:12.099 
💀    🚀 deployDemoBackend... 🏁 06:55:12.099 Duration: 1s
💀    🚀 deployDemoBackend... 🏁 06:55:12.099 
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:12.1   warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:12.1      $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 deployDemoBackend... 🏁 06:55:12.1   or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀 🔥 Error running 🏁 'deployDemoBackendDeployment' command: exit status 255
💀 🔎 Job Ended...
         Elapsed Time: 13.507551996s
         Current Time: 06:55:12
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["deploy"]
🔥 Stderr    : exit status 255
💀 🔎 Job Starting...
         Elapsed Time: 1.785µs
         Current Time: 06:55:12
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 06:55:12.891 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:12.893 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 06:55:12.893 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 06:55:13.226 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.231 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.237 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.238 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.238 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.239 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.243 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.244 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.246 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.247 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.248 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.248 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.249 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.249 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.25  Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.25  Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.253 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.254 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.255 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.256 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.259 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.26  Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.261 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.261 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.263 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.264 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.266 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.269 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.274 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.278 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.278 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.288 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.291 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.305 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.315 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.315 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.32  Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.322 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.33  Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.335 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.338 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.346 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.347 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.347 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 06:55:13.35  Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.351 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 06:55:13.356 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoBacken... 🏁 06:55:13.382 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 06:55:13.382 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.383 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoFronte... 🏁 06:55:13.395 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 06:55:13.395 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:55:13.41  WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:55:13.41  You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.757 🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 06:55:13.757 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.862     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 for this case.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 Usage:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 Aliases:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 Flags:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 06:55:13.863       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 06:55:13.864       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 06:55:13.864 
💀    🚀 prepareDemoFronte... 🏁 06:55:13.864 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 06:55:13.866 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.905 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.905 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.954 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.955 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 06:55:13.956 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:14.727 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:14.728    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:14.728 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.736 PARTS: ["3000"]
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:14.809 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:14.809    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:14.809 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.828 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 06:55:14.828 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 06:55:14.884 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 for this case.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 Usage:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 Aliases:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.885 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 Flags:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 06:55:14.886 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.208 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.208    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.208 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 destroyDemoFronte... 🏁 06:55:15.498 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 06:55:15.598 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.601  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603  
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603 Resources:
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603     - 3 to delete
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.603 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.612 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 06:55:15.716 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.719  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.72  
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.728  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.729  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.729  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.73   -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.73   
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.73  Resources:
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.731     - 3 to delete
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.731 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.731 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 06:55:15.819  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.82   -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.82   -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821  
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821 Resources:
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821     - 3 deleted
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 06:55:15.821 
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:15.822 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:15.822    $ curl -sSL https://get.pulumi.com | sh
💀    🚀 destroyDemoFronte... 🏁 06:55:15.822 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 06:55:15.822 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀 🔥 🚀 destroyDemoFronte... 🏁 06:55:15.822 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 destroyDemoFronte... 🏁 06:55:15.824 hello world
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.835 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.838  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.924  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.924  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.925  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927  
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927     - 3 deleted
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927 Duration: 1s
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.927 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.928 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.928 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:15.928 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:15.928    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 destroyDemoDbDepl... 🏁 06:55:15.928 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 destroyDemoDbDepl... 🏁 06:55:15.929 hello world
💀    🚀 destroyDemoBacken... 🏁 06:55:15.947 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 06:55:15.948 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.948  
💀    🚀 destroyDemoBacken... 🏁 06:55:15.948 Resources:
💀    🚀 destroyDemoBacken... 🏁 06:55:15.948 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.948 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949  
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 Resources:
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 Duration: 1s
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoBacken... 🏁 06:55:15.949 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.949 warning: A new version of Pulumi is available. To upgrade from version '3.26.1' to '3.32.1', run 
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.949    $ curl -sSL https://get.pulumi.com | sh
💀 🔥 🚀 destroyDemoBacken... 🏁 06:55:15.949 or visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes.
💀    🚀 destroyDemoBacken... 🏁 06:55:15.95  hello world
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 06:55:16.061 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 3.284322044s
         Current Time: 06:55:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.395687074s
         Current Time: 06:55:16
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

