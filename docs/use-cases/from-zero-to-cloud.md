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
         Elapsed Time: 1.088µs
         Current Time: 09:31:36
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 09:31:36.762 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 09:31:36.77  🎉🎉🎉
💀    🚀 initProject          🚧 09:31:36.77  Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 117.749312ms
         Current Time: 09:31:36
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 318.482314ms
         Current Time: 09:31:37
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.982µs
         Current Time: 09:31:37
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:31:37.249 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:31:37.256         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:31:37.256     
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:31:37.256   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:31:37.256   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:31:37.256   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:31:37.256 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 09:31:37.72  🧰 Prepare
💀    🚀 makeMysqlApp         🐬 09:31:37.721 Preparing base variables
💀    🚀 makeMysqlApp         🐬 09:31:37.828 Base variables prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Preparing start command
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Start command prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Preparing test command
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Test command prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 09:31:37.829 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.83  Preparing check command
💀    🚀 makeMysqlApp         🐬 09:31:37.83  Check command prepared
💀    🚀 makeMysqlApp         🐬 09:31:37.83  Preparing replacement map
💀    🚀 makeMysqlApp         🐬 09:31:38.1   Add config to replacement map
💀    🚀 makeMysqlApp         🐬 09:31:38.106 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 09:31:38.114 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 09:31:38.114 ✅ Validate
💀    🚀 makeMysqlApp         🐬 09:31:38.114 Validate app directory
💀    🚀 makeMysqlApp         🐬 09:31:38.115 Done validating app directory
💀    🚀 makeMysqlApp         🐬 09:31:38.115 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 09:31:38.118 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 09:31:38.118 Validate template locations
💀    🚀 makeMysqlApp         🐬 09:31:38.127 Done validating template locations
💀    🚀 makeMysqlApp         🐬 09:31:38.127 Validate app ports
💀    🚀 makeMysqlApp         🐬 09:31:38.13  Done validating app ports
💀    🚀 makeMysqlApp         🐬 09:31:38.13  Validate app crud fields
💀    🚀 makeMysqlApp         🐬 09:31:38.133 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 09:31:38.133 🚧 Generate
💀    🚀 makeMysqlApp         🐬 09:31:38.133 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 09:31:38.133   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 09:31:38.133 ]
💀    🚀 makeMysqlApp         🐬 09:31:38.133 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 09:31:38.152 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 09:31:38.152 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 09:31:38.152 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.637 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.637 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.904 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.904 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:38.905 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.259 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.268 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.275 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.275 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.275 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.275 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.275 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.279 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.279 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.309 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.309 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.313 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.313 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318 ]
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.318 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.381 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.385 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.388 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.604 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.768 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 09:31:39.954 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.105 Checking start
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.108 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.289 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.473 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.477 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.63  Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.808 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 09:31:40.986 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.144 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.147 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.319 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.475 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.479 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.657 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.839 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 09:31:41.842 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.013 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.212 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.215 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.401 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.592 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.597 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 09:31:42.597 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.456714634s
         Current Time: 09:31:42
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.567831746s
         Current Time: 09:31:42
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.224µs
         Current Time: 09:31:42
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:31:42.973 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:31:42.982         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:31:42.982     
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:31:42.982   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:31:42.982   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:31:42.982   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:31:42.982 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 09:31:43.407 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 09:31:43.407 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 09:31:43.588 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.588 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 09:31:43.588 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.589 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 09:31:43.8   Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 09:31:43.807 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 09:31:43.813 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 09:31:43.813 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 09:31:43.813 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 09:31:43.813 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 09:31:43.813 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 09:31:43.816 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 09:31:43.816 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 09:31:43.827 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 09:31:43.827 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 09:31:43.83  Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 09:31:43.83  Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833 ]
💀    🚀 makeFastApiApp       ⚡ 09:31:43.833 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 09:31:44.454 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 09:31:44.456 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 09:31:44.456 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 09:31:44.901 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 09:31:44.901 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.756 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.756 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.757 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.989 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 09:31:45.996 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.003 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.003 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.004 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.004 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.004 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.007 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.007 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.022 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.022 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.027 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.027 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 ]
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.032 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.087 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.089 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.092 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.256 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.259 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.419 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.578 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.581 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.744 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.898 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 09:31:46.901 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.064 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.261 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.264 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.426 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.597 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.601 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.764 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:47.928 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.084 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.246 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.249 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.405 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.567 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.571 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.733 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.923 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 09:31:48.928 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.089 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.248 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.251 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.407 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.571 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.748 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 09:31:49.922 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.091 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.257 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.417 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.586 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.749 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 09:31:50.749 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 09:31:51.276 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 09:31:51.276 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing start command
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Start command prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing test command
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Test command prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing check command
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Check command prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.088 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 09:31:52.358 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 09:31:52.367 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 09:31:52.375 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 09:31:52.375 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 09:31:52.375 Validate app directory
💀    🚀 addFastApiModule     ⚡ 09:31:52.375 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 09:31:52.375 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 09:31:52.379 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 09:31:52.38  Validate template locations
💀    🚀 addFastApiModule     ⚡ 09:31:52.391 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 09:31:52.391 Validate app ports
💀    🚀 addFastApiModule     ⚡ 09:31:52.398 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 09:31:52.398 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 09:31:52.403 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 09:31:52.403 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 09:31:52.403 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 09:31:52.403   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 09:31:52.404 ]
💀    🚀 addFastApiModule     ⚡ 09:31:52.404 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 09:31:52.446 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 09:31:52.447 Registering module
💀    🚀 addFastApiModule     ⚡ 09:31:52.48  Done registering module
💀    🚀 addFastApiModule     ⚡ 09:31:52.481 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 09:31:52.481 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 09:31:52.844 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 09:31:52.844 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.751 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:53.974 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:53.981 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:53.987 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:53.987 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 09:31:53.996 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 09:31:53.997 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 09:31:54.056 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 09:31:54.056 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 09:31:54.151 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 09:31:54.153 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 09:31:54.258 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 09:31:54.258 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.322 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:54.59  Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:54.597 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 09:31:54.603 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 09:31:54.603 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 09:31:54.603 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 09:31:54.603 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 09:31:54.603 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 09:31:54.606 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 09:31:54.606 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 09:31:54.616 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 09:31:54.616 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 09:31:54.619 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 09:31:54.619 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 09:31:54.622 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623 ]
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623 
💀    🚀 addFastApiCrud       ⚡ 09:31:54.623 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 09:31:54.672 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 09:31:54.673 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 09:31:54.736 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 09:31:54.737 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 09:31:54.802 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 09:31:54.802 Registering repo
💀    🚀 addFastApiCrud       ⚡ 09:31:54.859 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 09:31:54.859 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 09:31:54.859 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 11.993224085s
         Current Time: 09:31:54
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.193923007s
         Current Time: 09:31:55
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.776µs
         Current Time: 09:31:55
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:31:55.347 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:31:55.348 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:31:55.348 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:31:55.349 
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:31:55.349         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:31:55.349     
💀    🚀 zrbShowAdv           ☕ 09:31:55.349 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:31:55.349 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:31:55.349   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:31:55.349   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:31:55.349   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:31:55.349 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 09:31:55.81  🧰 Prepare
💀    🚀 makeNginxApp         📗 09:31:55.81  Preparing base variables
💀    🚀 makeNginxApp         📗 09:31:55.904 Base variables prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing start command
💀    🚀 makeNginxApp         📗 09:31:55.904 Start command prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing prepare command
💀    🚀 makeNginxApp         📗 09:31:55.904 Prepare command prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing test command
💀    🚀 makeNginxApp         📗 09:31:55.904 Test command prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing migrate command
💀    🚀 makeNginxApp         📗 09:31:55.904 Migrate command prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing check command
💀    🚀 makeNginxApp         📗 09:31:55.904 Check command prepared
💀    🚀 makeNginxApp         📗 09:31:55.904 Preparing replacement map
💀    🚀 makeNginxApp         📗 09:31:56.137 Add config to replacement map
💀    🚀 makeNginxApp         📗 09:31:56.149 Add env to replacement map
💀    🚀 makeNginxApp         📗 09:31:56.16  Replacement map prepared
💀    🚀 makeNginxApp         📗 09:31:56.16  ✅ Validate
💀    🚀 makeNginxApp         📗 09:31:56.16  Validate app directory
💀    🚀 makeNginxApp         📗 09:31:56.16  Done validating app directory
💀    🚀 makeNginxApp         📗 09:31:56.16  Validate app container volumes
💀    🚀 makeNginxApp         📗 09:31:56.164 Done validating app container volumes
💀    🚀 makeNginxApp         📗 09:31:56.164 Validate template locations
💀    🚀 makeNginxApp         📗 09:31:56.177 Done validating template locations
💀    🚀 makeNginxApp         📗 09:31:56.177 Validate app ports
💀    🚀 makeNginxApp         📗 09:31:56.182 Done validating app ports
💀    🚀 makeNginxApp         📗 09:31:56.182 Validate app crud fields
💀    🚀 makeNginxApp         📗 09:31:56.187 Done validating app crud fields
💀    🚀 makeNginxApp         📗 09:31:56.187 🚧 Generate
💀    🚀 makeNginxApp         📗 09:31:56.187 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 09:31:56.187   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 09:31:56.187 ]
💀    🚀 makeNginxApp         📗 09:31:56.187 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 09:31:56.209 🔩 Integrate
💀    🚀 makeNginxApp         📗 09:31:56.21  🎉🎉🎉
💀    🚀 makeNginxApp         📗 09:31:56.21  Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 09:31:56.636 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 09:31:56.636 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 09:31:56.738 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.738 Preparing start command
💀    🚀 makeNginxAppRunner   📗 09:31:56.738 Start command prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.738 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 09:31:56.738 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Preparing test command
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Test command prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Preparing check command
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Check command prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.739 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 09:31:56.965 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 09:31:56.972 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 09:31:56.977 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 09:31:56.978 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 09:31:56.978 Validate app directory
💀    🚀 makeNginxAppRunner   📗 09:31:56.978 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 09:31:56.978 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 09:31:56.981 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 09:31:56.981 Validate template locations
💀    🚀 makeNginxAppRunner   📗 09:31:56.993 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 09:31:56.993 Validate app ports
💀    🚀 makeNginxAppRunner   📗 09:31:56.996 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 09:31:56.996 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 09:31:56.999 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 09:31:56.999 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 09:31:56.999 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 09:31:56.999   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 09:31:56.999   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 09:31:56.999 ]
💀    🚀 makeNginxAppRunner   📗 09:31:56.999 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 09:31:57.025 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 09:31:57.028 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 09:31:57.031 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 09:31:57.214 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 09:31:57.375 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 09:31:57.572 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 09:31:57.761 Checking start
💀    🚀 makeNginxAppRunner   📗 09:31:57.764 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 09:31:57.962 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 09:31:58.144 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 09:31:58.148 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 09:31:58.378 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 09:31:58.605 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 09:31:58.802 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 09:31:58.973 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 09:31:58.977 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 09:31:59.153 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 09:31:59.341 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 09:31:59.344 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 09:31:59.538 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 09:31:59.727 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 09:31:59.73  Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 09:31:59.931 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 09:32:00.157 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 09:32:00.163 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 09:32:00.342 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 09:32:00.517 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 09:32:00.521 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 09:32:00.521 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.282488057s
         Current Time: 09:32:00
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.392783013s
         Current Time: 09:32:00
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["80:80"]' -v 'appEnvs={"API_HOST":"localhost:3000"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.666µs
         Current Time: 09:32:00
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 09:32:00.918 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 09:32:00.918 Links updated
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 prepareDemoBackend   🔧 09:32:00.922 Create venv
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 09:32:01.04  Build image demo-db:latest
💀    🚀 zrbCreateDockerNe... 🐳 09:32:01.16  🐳 Network 'zaruba' is already exist
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 09:32:01.183 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀    🚀 prepareDemoBackend   🔧 09:32:03.179 Activate venv
💀    🚀 prepareDemoBackend   🔧 09:32:03.18  Install dependencies
💀    🚀 prepareDemoBackend   🔧 09:32:03.447 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 09:32:03.579   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:03.585 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 09:32:03.643   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:03.651 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 09:32:03.687   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 prepareDemoBackend   🔧 09:32:06.111 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 09:32:06.373   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 09:32:06.385 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 09:32:06.458   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:06.467 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 09:32:06.674   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:06.69  Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 09:32:06.865   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:06.879 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 09:32:07.201   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 09:32:07.369 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 buildDemoFrontend... 🏭 09:32:07.613 Sending build context to Docker daemon  13.31kB
💀    🚀 buildDemoDbImage     🏭 09:32:07.613 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoDbImage     🏭 09:32:07.829 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 09:32:07.829 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoDbImage     🏭 09:32:07.829  ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 09:32:07.829  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 09:32:07.829 Step 2/6 : USER 0
💀    🚀 buildDemoDbImage     🏭 09:32:07.83  Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoDbImage     🏭 09:32:07.833 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 09:32:07.833  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834 Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:07.834  ---> 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 09:32:07.835 Successfully built 736550e2d78d
💀    🚀 buildDemoDbImage     🏭 09:32:07.836 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 09:32:07.836 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 09:32:07.838 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 09:32:07.84  🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 09:32:07.84  Docker image demo-frontend built
💀    🚀 prepareDemoBackend   🔧 09:32:07.891   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 prepareDemoBackend   🔧 09:32:07.974 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 09:32:08.141   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 09:32:08.153 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 09:32:08.206 🔎 Waiting docker container 'demoDb' running status
💀    🚀 prepareDemoBackend   🔧 09:32:08.239 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 startDemoFrontend... 📗 09:32:08.283 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 09:32:08.305 🐳 Retrieve previous log of 'demoDb'
💀    🚀 prepareDemoBackend   🔧 09:32:08.507   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 09:32:08.702 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 09:32:08.845   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 09:32:08.856 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧 09:32:08.89    Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:08.896 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 09:32:08.942   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:08.947 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 09:32:08.99    Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:08.999 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 09:32:09.041   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.066 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 09:32:09.107   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.12  Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 09:32:09.247   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.262 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 09:32:09.307   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.316 [38;5;6mnginx [38;5;5m23:53:38.12 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.12 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.13 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀    🚀 startDemoFrontend... 📗 09:32:09.317 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.13 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.13 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.13 [38;5;2mINFO  ==> ** Starting NGINX setup **
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.15 [38;5;2mINFO  ==> Validating settings in NGINX_* env vars
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> No custom scripts in /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> Initializing NGINX
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 realpath: /bitnami/nginx/conf/vhosts: No such file or directory
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.18 [38;5;2mINFO  ==> ** NGINX setup finished! **
💀 🔥 🚀 startDemoFrontend... 📗 09:32:09.317 [38;5;6mnginx [38;5;5m23:53:38.19 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🚀 prepareDemoBackend   🔧 09:32:09.318 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 startDemoFrontend... 📗 09:32:09.319 🐳 Starting container 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.039634Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 [38;5;6mmysql [38;5;5m23:53:47.76 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 [38;5;6mmysql [38;5;5m23:53:53.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.041494Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.041504Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.045770Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 [38;5;6mmysql [38;5;5m23:53:53.80 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.171704Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.350325Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 [38;5;6mmysql [38;5;5m23:53:56.81 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:09.337 [38;5;6mmysql [38;5;5m23:53:56.83 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.350378Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.416412Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:53:57.416514Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:54:36.519804Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:54:38.521030Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 09:32:09.337 2022-05-08T23:54:39.918045Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 09:32:09.339 🐳 Starting container 'demoDb'
💀    🚀 prepareDemoBackend   🔧 09:32:09.426   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.445 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 09:32:09.492   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.501 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 09:32:09.541   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.552 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 09:32:09.604   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 09:32:09.824 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 09:32:09.891   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:09.907 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🔥 🚀 startDemoFrontend... 📗 09:32:10.084 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 09:32:10.084 Error: failed to start containers: demoFrontend
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
            36 |   docker run --name "${CONTAINER_NAME}" --hostname "${CONTAINER_NAME}" --network "zaruba"  --shm-size "100m" -e 'API_HOST=host.docker.internal:3000' -e 'PYTHONUNBUFFERED=1' -p 80:80 -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/letsencrypt:/etc/letsencrypt" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/html:/opt/bitnami/nginx/html" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/server_blocks:/opt/bitnami/nginx/conf/server_blocks"  --restart no -d "${DOCKER_IMAGE_NAME}" 
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
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=16241)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=16242)
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=16211)
💀 🔪 Kill 🔧 'prepareDemoBackend' command (PID=14049)
💀    🚀 prepareDemoBackend   🔧 09:32:10.255   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 09:32:10.354 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:10.371 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:10.371 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀    🚀 prepareDemoBackend   🔧 09:32:10.413   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:10.424 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 09:32:10.478   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 09:32:10.483 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 🚀 prepareDemoBackend   🔧 09:32:10.51  ERROR: Operation cancelled by user
💀 🔥 🚀 prepareDemoBackend   🔧 09:32:10.516 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 09:32:10.516 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 Error running 🔧 'prepareDemoBackend' command: exit status 1
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 10.180784241s
         Current Time: 09:32:11
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["start"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.127µs
         Current Time: 09:32:11
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 09:32:11.251 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 09:32:11.251 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 09:32:11.281 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 09:32:11.369 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 09:32:11.515 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 09:32:11.515 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 09:32:12.425 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 09:32:12.425 Sending build context to Docker daemon  13.31kB
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476 Step 2/6 : USER 0
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 09:32:12.476 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477 Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477  ---> 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 09:32:12.477 Successfully built 736550e2d78d
💀    🚀 buildDemoBackendI... 🏭 09:32:12.478 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoDbImage     🏭 09:32:12.479 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 09:32:12.48  Successfully tagged demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 09:32:12.48   ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 09:32:12.483 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 09:32:12.483 Docker image demo-frontend built
💀    🚀 buildDemoDbImage     🏭 09:32:12.483 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 09:32:12.485 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoDbImage     🏭 09:32:12.485 Successfully tagged demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 09:32:12.485  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 09:32:12.485 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 09:32:12.486 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 09:32:12.487  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.487  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 09:32:12.487 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 09:32:12.488  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.488  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 09:32:12.488 Step 6/9 : COPY . .
💀    🚀 buildDemoDbImage     🏭 09:32:12.488 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 09:32:12.488 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 09:32:12.491 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 09:32:12.492  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.492  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 09:32:12.492 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 09:32:12.492  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 09:32:12.492  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 09:32:12.493 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 09:32:12.5   Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 09:32:12.502 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 09:32:12.502 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 09:32:12.828 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 09:32:12.869 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 09:32:12.918 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 09:32:12.945 🐳 Retrieve previous log of 'demoDb'
💀    🚀 startDemoFrontend... 📗 09:32:13.949 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.949 [38;5;6mnginx [38;5;5m23:53:38.12 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.949 [38;5;6mnginx [38;5;5m23:53:38.12 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.949 [38;5;6mnginx [38;5;5m23:53:38.13 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.949 [38;5;6mnginx [38;5;5m23:53:38.13 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.949 [38;5;6mnginx [38;5;5m23:53:38.13 
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.13 [38;5;2mINFO  ==> ** Starting NGINX setup **
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.15 [38;5;2mINFO  ==> Validating settings in NGINX_* env vars
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> No custom scripts in /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> Initializing NGINX
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  realpath: /bitnami/nginx/conf/vhosts: No such file or directory
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.18 [38;5;2mINFO  ==> ** NGINX setup finished! **
💀 🔥 🚀 startDemoFrontend... 📗 09:32:13.95  [38;5;6mnginx [38;5;5m23:53:38.19 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🚀 startDemoFrontend... 📗 09:32:13.953 🐳 Starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 [38;5;6mmysql [38;5;5m23:53:47.76 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 [38;5;6mmysql [38;5;5m23:53:53.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 [38;5;6mmysql [38;5;5m23:53:53.80 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 [38;5;6mmysql [38;5;5m23:53:56.81 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:13.983 [38;5;6mmysql [38;5;5m23:53:56.83 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.039634Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.041494Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.041504Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.045770Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.171704Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.350325Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.350378Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.416412Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:53:57.416514Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:54:36.519804Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:54:38.521030Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 09:32:13.983 2022-05-08T23:54:39.918045Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 09:32:13.986 🐳 Starting container 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 09:32:14.753 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 09:32:14.753 Error: failed to start containers: demoFrontend
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
            36 |   docker run --name "${CONTAINER_NAME}" --hostname "${CONTAINER_NAME}" --network "zaruba"  --shm-size "100m" -e 'API_HOST=host.docker.internal:3000' -e 'PYTHONUNBUFFERED=1' -p 80:80 -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/letsencrypt:/etc/letsencrypt" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/html:/opt/bitnami/nginx/html" -v "/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend/server_blocks:/opt/bitnami/nginx/conf/server_blocks"  --restart no -d "${DOCKER_IMAGE_NAME}" 
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
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=17949)
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=17919)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=17948)
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:14.964 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 09:32:14.964 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 4.514648441s
         Current Time: 09:32:15
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["startContainers"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.157µs
         Current Time: 09:32:15
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 09:32:15.917 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 09:32:15.917 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoBackendCo... ✋ 09:32:16.395 Docker container demoBackend is not running
💀    🚀 stopDemoFrontendC... ✋ 09:32:16.398 Docker container demoFrontend is not running
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 09:32:16.505 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 695.210265ms
         Current Time: 09:32:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 805.85999ms
         Current Time: 09:32:16
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.991µs
         Current Time: 09:32:16
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:16.959 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:32:16.961         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:32:16.961     
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:32:16.961   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:32:16.961   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:32:16.961   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:32:16.961 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.394 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.395 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.564 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.78  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.788 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.794 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.794 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.794 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.794 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.794 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.798 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.798 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.807 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.808 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.81  Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.81  Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813 ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.813 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.851 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.851 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:17.851 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.187 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.187 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.331 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.538 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.545 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.55  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.55  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.55  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.551 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.551 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.554 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.554 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.562 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.562 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.565 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.565 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568 ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.568 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.592 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.596 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.599 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:18.805 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.083 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.086 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.252 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.439 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.442 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.603 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.603 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:19.603 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.757752848s
         Current Time: 09:32:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.869258157s
         Current Time: 09:32:19
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.103µs
         Current Time: 09:32:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:19.997 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:32:19.999         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:32:19.999     
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:32:19.999   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:32:19.999   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:32:19.999   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:32:19.999 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:20.441 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:20.441 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.396 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.396 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.396 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.399 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.921 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.931 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.939 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.939 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.94  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.94  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.94  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.944 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.944 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.958 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.958 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.962 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.962 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966 ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:21.966 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:22.003 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:22.003 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:22.003 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:22.46  🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:22.46  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.243 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.478 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.485 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.491 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.491 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.492 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.492 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.492 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.495 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.495 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.504 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.504 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.507 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.507 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51  ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.51  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.531 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.534 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.537 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.715 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.887 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:23.89  Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.083 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.257 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.261 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.426 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.426 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:24.426 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.538416929s
         Current Time: 09:32:24
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.649064291s
         Current Time: 09:32:24
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.103µs
         Current Time: 09:32:24
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:24.809 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 09:32:24.812         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 09:32:24.812     
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 09:32:24.812   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 09:32:24.812   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 09:32:24.812   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 09:32:24.812 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.255 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.255 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.352 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.353 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.353 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.576 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.582 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.588 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.588 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.588 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.588 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.588 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.591 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.591 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.599 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.599 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.602 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.602 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.605 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.606 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.606 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.606   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.606 ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.606 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.636 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.636 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:25.636 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.053 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.053 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.163 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.164 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.379 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.386 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.392 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.392 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.392 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.393 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.393 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.395 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.395 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.404 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.404 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.407 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.407 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409 ]
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.409 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.428 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.431 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.434 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.638 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.8   Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.803 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:26.993 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 09:32:27.155 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:27.158 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 09:32:27.354 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 09:32:27.354 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 09:32:27.354 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.653009047s
         Current Time: 09:32:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.764257589s
         Current Time: 09:32:27
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.219µs
         Current Time: 09:32:27
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:27.759 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 09:32:27.868 Synchronize task environments
💀    🚀 syncEnv              🔄 09:32:28.071 Synchronize project's environment files
💀    🚀 syncEnv              🔄 09:32:28.25  🎉🎉🎉
💀    🚀 syncEnv              🔄 09:32:28.25  Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 596.558691ms
         Current Time: 09:32:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 798.49287ms
         Current Time: 09:32:28
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.138µs
         Current Time: 09:32:28
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:28.73  Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 09:32:28.846 🎉🎉🎉
💀    🚀 setProjectValue      🔗 09:32:28.846 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 223.199958ms
         Current Time: 09:32:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 424.527043ms
         Current Time: 09:32:29
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.828µs
         Current Time: 09:32:29
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 09:32:29.354 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 09:32:29.465 🎉🎉🎉
💀    🚀 setProjectValue      🔗 09:32:29.465 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 217.266393ms
         Current Time: 09:32:29
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.945498ms
         Current Time: 09:32:29
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.07µs
         Current Time: 09:32:29
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 09:32:29.959 🚧 Create virtual environment.
💀    🚀 prepareDemoFronte... 🏁 09:32:29.959 🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 09:32:29.96  🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 09:32:32.112 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 09:32:32.22  🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 09:32:32.227 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:32.417 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:32.493 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:32.51  Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:33.101   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:33.118 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:33.304   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:33.325 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:33.525   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 09:32:33.535   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:33.552 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:33.563   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 09:32:33.822 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:33.873 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:33.927   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 09:32:34.057   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:34.062 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:34.112   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:34.121 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:34.197 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:34.239   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:34.248 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:34.299   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:34.322 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:34.872   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:34.898 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:34.995   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:35.075   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:35.083 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:35.094 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:35.23    Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:35.308 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:35.501   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:35.538 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:35.556   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:35.578 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:35.801   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:35.806 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:35.823   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:35.838 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 09:32:35.869   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.002   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.007 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.027 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 09:32:36.188   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.21  Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.237   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.243   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.252 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.252 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:36.331   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.335   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.342 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:36.343 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:36.375   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.392 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:36.425   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.454 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:36.578   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.6   Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:36.635   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.668 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.688   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.7   Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:36.788   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.791   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:36.81  Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.822   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:36.835 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:36.846 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:37.02    Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.024   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.035 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.057 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:37.144   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:37.153 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:37.202   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.203   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.213 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.226 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:37.3     Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:37.309 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:37.328   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.335 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.425   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.432 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:37.508   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 09:32:37.516 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 09:32:37.565   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.589 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.627   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.637   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.639 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 09:32:37.658   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 09:32:37.736 Installing collected packages: six, grpcio, pyyaml, dill, protobuf, semver, pulumi, attrs, arpeggio, parver, idna, charset-normalizer, urllib3, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 09:32:37.772 Installing collected packages: dill, pyyaml, protobuf, six, grpcio, semver, pulumi, attrs, arpeggio, parver, charset-normalizer, idna, certifi, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.818   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:37.829 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 09:32:38.083   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 09:32:38.214 Installing collected packages: six, grpcio, protobuf, semver, pyyaml, dill, pulumi, attrs, arpeggio, parver, certifi, idna, charset-normalizer, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 09:32:38.616   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 09:32:38.633   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 09:32:39.206   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 09:32:40.689     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 09:32:40.689     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 09:32:40.75  Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoBacken... 🏁 09:32:40.755 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 09:32:40.79  WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 09:32:40.79  You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 09:32:40.803 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 09:32:40.803 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 09:32:40.967 🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 09:32:40.967 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.148     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 dependencies.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 09:32:41.188 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189 
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.189       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  for this case.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Usage:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Aliases:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Available Commands:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Flags:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  
💀    🚀 prepareDemoFronte... 🏁 09:32:41.19  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 09:32:41.191 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.229 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoDbDepl... 🏁 09:32:41.275 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 09:32:41.275 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 09:32:41.469 error: no stack named 'dev' found
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.584 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.584 🚧 Prepare chart dependencies.
💀    🚀 deployDemoFronten... 🏁 09:32:41.611 Created stack 'dev'
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.643     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.644 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:41.646 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoDbDeplo... 🏁 09:32:41.882 error: no stack named 'dev' found
💀    🚀 prepareDemoBacken... 🏁 09:32:41.993 PARTS: ["3000"]
💀    🚀 deployDemoDbDeplo... 🏁 09:32:41.996 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 09:32:42.053 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 09:32:42.053 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.103 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.103 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.103 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 dependencies.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104     dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 09:32:42.104 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105     dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 for this case.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 Usage:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 Aliases:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 Flags:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 09:32:42.105   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106 
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 09:32:42.106 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 09:32:42.352 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 09:32:42.467 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 09:32:43.444 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 09:32:43.923 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 09:32:43.989 
💀    🚀 deployDemoFronten... 🏁 09:32:44.399  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:44.41  
💀    🚀 deployDemoBackend... 🏁 09:32:44.414 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 09:32:44.599  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:44.79   +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 09:32:44.85  
💀    🚀 deployDemoDbDeplo... 🏁 09:32:44.882  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 09:32:44.914  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 09:32:44.916  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 09:32:45.026  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 09:32:45.026  
💀    🚀 deployDemoFronten... 🏁 09:32:45.026 Resources:
💀    🚀 deployDemoFronten... 🏁 09:32:45.026     + 4 to create
💀    🚀 deployDemoFronten... 🏁 09:32:45.026 
💀    🚀 deployDemoFronten... 🏁 09:32:45.026 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.176  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.179  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoBackend... 🏁 09:32:45.279  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.29   +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.29   
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.291 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.291     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.291 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.291 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 09:32:45.364  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoFronten... 🏁 09:32:45.479 
💀    🚀 deployDemoBackend... 🏁 09:32:45.676  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 09:32:45.677  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 09:32:45.683  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:45.741 
💀    🚀 deployDemoBackend... 🏁 09:32:45.83   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 09:32:45.83   
💀    🚀 deployDemoBackend... 🏁 09:32:45.83  Resources:
💀    🚀 deployDemoBackend... 🏁 09:32:45.83      + 5 to create
💀    🚀 deployDemoBackend... 🏁 09:32:45.83  
💀    🚀 deployDemoBackend... 🏁 09:32:45.83  Updating (dev):
💀    🚀 deployDemoFronten... 🏁 09:32:45.907  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 09:32:45.985  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.213  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 09:32:46.275  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 09:32:46.279  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 09:32:46.288 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.315  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 09:32:46.324  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating Retry #0; creation failed: serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 09:32:46.326  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 09:32:46.326  +  kubernetes:core/v1:ServiceAccount default/demo-frontend **creating failed** error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 09:32:46.329  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 09:32:46.335  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 09:32:46.335  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating error: update failed
💀    🚀 deployDemoFronten... 🏁 09:32:46.343  +  pulumi:pulumi:Stack demoFrontendDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoFronten... 🏁 09:32:46.343  +  kubernetes:helm.sh/v3:Chart demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 09:32:46.343  
💀    🚀 deployDemoFronten... 🏁 09:32:46.343 Diagnostics:
💀    🚀 deployDemoFronten... 🏁 09:32:46.343   kubernetes:core/v1:ServiceAccount (default/demo-frontend):
💀    🚀 deployDemoFronten... 🏁 09:32:46.343     error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 09:32:46.343  
💀    🚀 deployDemoFronten... 🏁 09:32:46.343   pulumi:pulumi:Stack (demoFrontendDeployment-dev):
💀    🚀 deployDemoFronten... 🏁 09:32:46.343     error: update failed
💀    🚀 deployDemoFronten... 🏁 09:32:46.343  
💀    🚀 deployDemoFronten... 🏁 09:32:46.343 Resources:
💀    🚀 deployDemoFronten... 🏁 09:32:46.343     + 3 created
💀    🚀 deployDemoFronten... 🏁 09:32:46.343 
💀    🚀 deployDemoFronten... 🏁 09:32:46.343 Duration: 1s
💀    🚀 deployDemoFronten... 🏁 09:32:46.343 
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
💀 🔪 Kill 🏁 'deployDemoDbDeployment' command (PID=27053)
💀 🔪 Kill 🏁 'deployDemoBackendDeployment' command (PID=27546)
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.596  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.6    +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.612  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.612  +  kubernetes:core/v1:ServiceAccount default/demo-db creating Retry #0; creation failed: serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.613  +  kubernetes:core/v1:ServiceAccount default/demo-db creating error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.613  +  kubernetes:core/v1:ServiceAccount default/demo-db **creating failed** error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.618  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.618  +  pulumi:pulumi:Stack demoDbDeployment-dev creating error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.626  +  pulumi:pulumi:Stack demoDbDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.626  +  kubernetes:helm.sh/v3:Chart demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627  
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627 Diagnostics:
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627   kubernetes:core/v1:ServiceAccount (default/demo-db):
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627     error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627  
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627   pulumi:pulumi:Stack (demoDbDeployment-dev):
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627     error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627  
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627     + 3 created
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627 
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627 Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 09:32:46.627 
💀 🔥 Error running 🏁 'deployDemoDbDeployment' command: exit status 255
💀    🚀 deployDemoBackend... 🏁 09:32:46.726  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 09:32:46.792  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating error: update canceled
💀    🚀 deployDemoBackend... 🏁 09:32:46.796  +  pulumi:pulumi:Stack demoBackendDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoBackend... 🏁 09:32:46.796  
💀    🚀 deployDemoBackend... 🏁 09:32:46.796 Diagnostics:
💀    🚀 deployDemoBackend... 🏁 09:32:46.796   pulumi:pulumi:Stack (demoBackendDeployment-dev):
💀    🚀 deployDemoBackend... 🏁 09:32:46.796     error: update canceled
💀    🚀 deployDemoBackend... 🏁 09:32:46.796  
💀    🚀 deployDemoBackend... 🏁 09:32:46.796 Resources:
💀    🚀 deployDemoBackend... 🏁 09:32:46.796     + 1 created
💀    🚀 deployDemoBackend... 🏁 09:32:46.796 
💀    🚀 deployDemoBackend... 🏁 09:32:46.796 Duration: 1s
💀    🚀 deployDemoBackend... 🏁 09:32:46.796 
💀 🔥 Error running 🏁 'deployDemoBackendDeployment' command: exit status 255
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 17.404877183s
         Current Time: 09:32:47
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["deploy"]
🔥 Stderr    : exit status 255
💀 🔎 Job Starting...
         Elapsed Time: 1.275µs
         Current Time: 09:32:47
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 09:32:47.574 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.574 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 09:32:47.575 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 09:32:47.907 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.911 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.915 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.917 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.92  Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.922 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.923 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.923 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.923 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.924 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.927 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.928 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.929 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.929 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.93  Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.931 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.932 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.933 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.934 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.934 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.936 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.938 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.94  Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.942 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.946 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.947 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.949 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.951 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.953 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.96  Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.962 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.963 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:47.966 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.969 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.98  Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 09:32:47.987 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 09:32:47.998 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 09:32:48     Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.006 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.01  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.012 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 09:32:48.027 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 09:32:48.038 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 09:32:48.042 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.048 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 09:32:48.065 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoFronte... 🏁 09:32:48.068 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 09:32:48.068 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 09:32:48.068 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 09:32:48.073 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 09:32:48.076 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 09:32:48.077 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 09:32:48.092 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 09:32:48.092 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.247 🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 09:32:48.247 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.332 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.332 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.332 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 dependencies.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333     dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.333 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334     dependencies:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 for this case.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Usage:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Aliases:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Flags:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 09:32:48.334       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335 
💀    🚀 prepareDemoFronte... 🏁 09:32:48.335 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 09:32:48.336 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.38  🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.38  🚧 Prepare chart dependencies.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.445 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.446 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 09:32:48.447 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 09:32:49.581 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 09:32:49.671 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 09:32:49.671 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 dependencies.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732     dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.732 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733     dependencies:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 for this case.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Usage:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Aliases:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Flags:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 09:32:49.733       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734 
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 09:32:49.734 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 09:32:50.12  Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 09:32:50.243 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.247  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.249  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.249  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.251  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.251  
💀    🚀 destroyDemoFronte... 🏁 09:32:50.251 Resources:
💀    🚀 destroyDemoFronte... 🏁 09:32:50.251     - 3 to delete
💀    🚀 destroyDemoFronte... 🏁 09:32:50.252 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.252 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.315 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 09:32:50.401 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.404  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.471 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.475  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.475  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.475  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.478  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.478  
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.478 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.479     - 3 to delete
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.479 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.479 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 09:32:50.55   -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.55   -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.552  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.558  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.558  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559  
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 Resources:
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559     - 3 deleted
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 09:32:50.559 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoFronte... 🏁 09:32:50.562 hello world
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.612 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.617  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.737  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.737  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.737  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742  
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742     - 3 deleted
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742 Duration: 1s
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.742 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.743 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.743 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 09:32:50.746 hello world
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 09:32:51.39  Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 09:32:51.39  
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391  
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391 Resources:
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391     - 1 to delete
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.391  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392  
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392 Resources:
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392     - 1 deleted
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392 Duration: 1s
💀    🚀 destroyDemoBacken... 🏁 09:32:51.392 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.393 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoBacken... 🏁 09:32:51.393 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoBacken... 🏁 09:32:51.395 hello world
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 09:32:51.501 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 4.045219175s
         Current Time: 09:32:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.155560787s
         Current Time: 09:32:51
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

