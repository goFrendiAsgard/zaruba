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
zaruba please removeContainers

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
         Elapsed Time: 1.333µs
         Current Time: 20:55:28
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 20:55:28.071 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 20:55:28.075 🎉🎉🎉
💀    🚀 initProject          🚧 20:55:28.075 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 113.06514ms
         Current Time: 20:55:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 314.429461ms
         Current Time: 20:55:28
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 2.463µs
         Current Time: 20:55:28
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:55:28.521 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:55:28.526         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:55:28.526     
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:55:28.526   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:55:28.526   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:55:28.526   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:55:28.526 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 20:55:28.964 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 20:55:28.964 Preparing base variables
💀    🚀 makeMysqlApp         🐬 20:55:29.046 Base variables prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing start command
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Start command prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing test command
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Test command prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing check command
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Check command prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.047 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 20:55:29.262 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 20:55:29.269 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 20:55:29.275 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 20:55:29.275 ✅ Validate
💀    🚀 makeMysqlApp         🐬 20:55:29.275 Validate app directory
💀    🚀 makeMysqlApp         🐬 20:55:29.275 Done validating app directory
💀    🚀 makeMysqlApp         🐬 20:55:29.275 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 20:55:29.278 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 20:55:29.278 Validate template locations
💀    🚀 makeMysqlApp         🐬 20:55:29.287 Done validating template locations
💀    🚀 makeMysqlApp         🐬 20:55:29.287 Validate app ports
💀    🚀 makeMysqlApp         🐬 20:55:29.291 Done validating app ports
💀    🚀 makeMysqlApp         🐬 20:55:29.291 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 20:55:29.294 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 20:55:29.294 🚧 Generate
💀    🚀 makeMysqlApp         🐬 20:55:29.294 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 20:55:29.294   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 20:55:29.294 ]
💀    🚀 makeMysqlApp         🐬 20:55:29.294 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 20:55:29.309 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 20:55:29.309 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 20:55:29.309 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.686 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.686 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.839 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.839 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.839 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.839 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.839 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:29.84  Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.052 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.059 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.065 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.065 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.065 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.065 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.065 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.069 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.069 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.085 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.085 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.088 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.088 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092 ]
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.092 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.129 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.133 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.135 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.288 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.449 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.603 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.754 Checking start
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.757 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 20:55:30.918 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.066 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.069 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.216 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.361 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.508 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.653 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.656 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.802 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.95  Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:31.953 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.1   Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.243 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.246 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.389 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.535 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.538 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.683 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.832 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.835 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 20:55:32.835 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.419970333s
         Current Time: 20:55:32
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.531804043s
         Current Time: 20:55:33
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.333µs
         Current Time: 20:55:33
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:55:33.197 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:55:33.199 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:55:33.199 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:55:33.2   
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:55:33.2           '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:55:33.2       
💀    🚀 zrbShowAdv           ☕ 20:55:33.2   Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:55:33.2   You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:55:33.2     * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:55:33.2     * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:55:33.2     * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:55:33.2   
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 20:55:33.632 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 20:55:33.632 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing start command
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Start command prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing test command
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Test command prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing check command
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Check command prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.77  Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 20:55:33.984 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 20:55:33.991 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 20:55:33.997 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 20:55:33.997 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 20:55:33.998 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 20:55:33.998 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 20:55:33.998 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 20:55:34     Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 20:55:34     Validate template locations
💀    🚀 makeFastApiApp       ⚡ 20:55:34.01  Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 20:55:34.01  Validate app ports
💀    🚀 makeFastApiApp       ⚡ 20:55:34.013 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 20:55:34.013 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 20:55:34.015 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 20:55:34.016 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 20:55:34.016 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 20:55:34.016   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 20:55:34.016 ]
💀    🚀 makeFastApiApp       ⚡ 20:55:34.016 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 20:55:34.524 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 20:55:34.525 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 20:55:34.525 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 20:55:34.982 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 20:55:34.982 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.739 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.739 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.739 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.74  Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.975 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.982 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.988 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.988 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.988 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.988 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.988 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.991 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 20:55:35.991 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.007 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.007 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.011 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.011 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.013 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014 ]
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014 
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.014 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.057 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.06  Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.063 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.223 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.225 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.377 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.53  Checking test
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.533 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.683 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.839 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.842 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 20:55:36.991 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.141 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.144 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.298 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.475 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.478 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.624 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.768 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:37.942 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.088 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.092 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.24  Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.383 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.386 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.538 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.684 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.687 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.834 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.985 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 20:55:38.988 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.138 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.287 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.444 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.593 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.738 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 20:55:39.887 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 20:55:40.036 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:40.187 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 20:55:40.339 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 20:55:40.34  Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 20:55:40.801 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 20:55:40.801 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing start command
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Start command prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing test command
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Test command prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing check command
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Check command prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.521 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 20:55:41.732 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 20:55:41.74  Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 20:55:41.746 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 20:55:41.746 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 20:55:41.746 Validate app directory
💀    🚀 addFastApiModule     ⚡ 20:55:41.746 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 20:55:41.746 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 20:55:41.749 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 20:55:41.749 Validate template locations
💀    🚀 addFastApiModule     ⚡ 20:55:41.757 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 20:55:41.757 Validate app ports
💀    🚀 addFastApiModule     ⚡ 20:55:41.76  Done validating app ports
💀    🚀 addFastApiModule     ⚡ 20:55:41.76  Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 20:55:41.763 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 20:55:41.763 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 20:55:41.763 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 20:55:41.763   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 20:55:41.763 ]
💀    🚀 addFastApiModule     ⚡ 20:55:41.763 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 20:55:41.777 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 20:55:41.777 Registering module
💀    🚀 addFastApiModule     ⚡ 20:55:41.799 Done registering module
💀    🚀 addFastApiModule     ⚡ 20:55:41.8   🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 20:55:41.8   Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 20:55:42.109 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 20:55:42.109 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 20:55:42.822 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:42.823 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.028 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.035 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.041 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.041 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 20:55:43.05  Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 20:55:43.05  Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 20:55:43.107 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 20:55:43.107 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 20:55:43.167 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 20:55:43.167 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 20:55:43.262 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 20:55:43.262 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 20:55:43.32  Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 20:55:43.32  Preparing start command
💀    🚀 addFastApiCrud       ⚡ 20:55:43.32  Start command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.32  Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.321 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.548 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.555 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 20:55:43.561 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 20:55:43.561 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 20:55:43.561 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 20:55:43.561 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 20:55:43.561 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 20:55:43.564 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 20:55:43.564 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 20:55:43.573 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 20:55:43.573 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 20:55:43.576 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 20:55:43.576 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 ]
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 
💀    🚀 addFastApiCrud       ⚡ 20:55:43.579 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 20:55:43.603 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 20:55:43.603 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 20:55:43.637 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 20:55:43.637 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 20:55:43.683 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 20:55:43.683 Registering repo
💀    🚀 addFastApiCrud       ⚡ 20:55:43.739 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 20:55:43.74  🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 20:55:43.74  Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 10.649441827s
         Current Time: 20:55:43
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 10.851310525s
         Current Time: 20:55:44
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.481µs
         Current Time: 20:55:44
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:55:44.195 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:55:44.197 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:55:44.197 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:55:44.197 
💀    🚀 zrbShowAdv           ☕ 20:55:44.197         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:55:44.197         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:55:44.197         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:55:44.197         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:55:44.197         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:55:44.198         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:55:44.198     
💀    🚀 zrbShowAdv           ☕ 20:55:44.198 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:55:44.198 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:55:44.198   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:55:44.198   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:55:44.198   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:55:44.198 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 20:55:44.627 🧰 Prepare
💀    🚀 makeNginxApp         📗 20:55:44.627 Preparing base variables
💀    🚀 makeNginxApp         📗 20:55:44.702 Base variables prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing start command
💀    🚀 makeNginxApp         📗 20:55:44.702 Start command prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing prepare command
💀    🚀 makeNginxApp         📗 20:55:44.702 Prepare command prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing test command
💀    🚀 makeNginxApp         📗 20:55:44.702 Test command prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing migrate command
💀    🚀 makeNginxApp         📗 20:55:44.702 Migrate command prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing check command
💀    🚀 makeNginxApp         📗 20:55:44.702 Check command prepared
💀    🚀 makeNginxApp         📗 20:55:44.702 Preparing replacement map
💀    🚀 makeNginxApp         📗 20:55:44.912 Add config to replacement map
💀    🚀 makeNginxApp         📗 20:55:44.919 Add env to replacement map
💀    🚀 makeNginxApp         📗 20:55:44.925 Replacement map prepared
💀    🚀 makeNginxApp         📗 20:55:44.925 ✅ Validate
💀    🚀 makeNginxApp         📗 20:55:44.925 Validate app directory
💀    🚀 makeNginxApp         📗 20:55:44.926 Done validating app directory
💀    🚀 makeNginxApp         📗 20:55:44.926 Validate app container volumes
💀    🚀 makeNginxApp         📗 20:55:44.928 Done validating app container volumes
💀    🚀 makeNginxApp         📗 20:55:44.929 Validate template locations
💀    🚀 makeNginxApp         📗 20:55:44.937 Done validating template locations
💀    🚀 makeNginxApp         📗 20:55:44.937 Validate app ports
💀    🚀 makeNginxApp         📗 20:55:44.94  Done validating app ports
💀    🚀 makeNginxApp         📗 20:55:44.94  Validate app crud fields
💀    🚀 makeNginxApp         📗 20:55:44.943 Done validating app crud fields
💀    🚀 makeNginxApp         📗 20:55:44.943 🚧 Generate
💀    🚀 makeNginxApp         📗 20:55:44.943 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 20:55:44.943   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 20:55:44.943 ]
💀    🚀 makeNginxApp         📗 20:55:44.944 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 20:55:44.967 🔩 Integrate
💀    🚀 makeNginxApp         📗 20:55:44.967 🎉🎉🎉
💀    🚀 makeNginxApp         📗 20:55:44.967 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 20:55:45.359 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 20:55:45.359 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 20:55:45.449 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.449 Preparing start command
💀    🚀 makeNginxAppRunner   📗 20:55:45.449 Start command prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.449 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 20:55:45.449 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Preparing test command
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Test command prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Preparing check command
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Check command prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.45  Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 20:55:45.649 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 20:55:45.655 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 20:55:45.661 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 20:55:45.661 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 20:55:45.661 Validate app directory
💀    🚀 makeNginxAppRunner   📗 20:55:45.661 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 20:55:45.661 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 20:55:45.664 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 20:55:45.664 Validate template locations
💀    🚀 makeNginxAppRunner   📗 20:55:45.677 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 20:55:45.677 Validate app ports
💀    🚀 makeNginxAppRunner   📗 20:55:45.68  Done validating app ports
💀    🚀 makeNginxAppRunner   📗 20:55:45.68  Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 20:55:45.683 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 20:55:45.683 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 20:55:45.683 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 20:55:45.684   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 20:55:45.684   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 20:55:45.684 ]
💀    🚀 makeNginxAppRunner   📗 20:55:45.684 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 20:55:45.707 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 20:55:45.711 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 20:55:45.714 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 20:55:45.869 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 20:55:46.016 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 20:55:46.2   Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 20:55:46.35  Checking start
💀    🚀 makeNginxAppRunner   📗 20:55:46.354 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 20:55:46.517 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 20:55:46.664 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 20:55:46.667 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 20:55:46.84  Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 20:55:46.994 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 20:55:47.144 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 20:55:47.297 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 20:55:47.3   Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 20:55:47.45  Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 20:55:47.601 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 20:55:47.605 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 20:55:47.758 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 20:55:47.915 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 20:55:47.919 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 20:55:48.069 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 20:55:48.218 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 20:55:48.221 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 20:55:48.368 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 20:55:48.518 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 20:55:48.521 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 20:55:48.521 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.431686856s
         Current Time: 20:55:48
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.543676501s
         Current Time: 20:55:48
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.567µs
         Current Time: 20:55:48
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:55:48.885 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 20:55:48.991 Synchronize task environments
💀    🚀 syncEnv              🔄 20:55:49.149 Synchronize project's environment files
💀    🚀 syncEnv              🔄 20:55:49.303 🎉🎉🎉
💀    🚀 syncEnv              🔄 20:55:49.303 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 524.225095ms
         Current Time: 20:55:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 726.21524ms
         Current Time: 20:55:49
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.771µs
         Current Time: 20:55:49
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 20:55:49.758 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 20:55:49.758 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 20:55:50.018 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 20:55:50.018 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 20:55:50.019 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 20:55:55.677 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 20:55:55.678 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoFrontend... 🏭 20:55:55.732 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 20:55:55.732  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 20:55:55.732 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 20:55:55.733 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 20:55:55.734 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoBackendI... 🏭 20:55:55.735 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:55.735  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 20:55:55.736 Successfully built 40621c693b70
💀    🚀 buildDemoDbImage     🏭 20:55:55.737 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoBackendI... 🏭 20:55:55.739 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoDbImage     🏭 20:55:55.741  ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 20:55:55.741 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 20:55:55.741 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 20:55:55.741  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 20:55:55.741 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoFrontend... 🏭 20:55:55.742 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 20:55:55.742 Docker image demo-frontend built
💀    🚀 buildDemoDbImage     🏭 20:55:55.744 Successfully tagged demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 20:55:55.744 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745  ---> 16e3e46a7774
💀    🚀 buildDemoDbImage     🏭 20:55:55.745 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 20:55:55.745 Step 6/9 : COPY . .
💀    🚀 buildDemoDbImage     🏭 20:55:55.745 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75  Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75  Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75   ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 20:55:55.75  Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 20:55:55.751  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:55:55.751  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 20:55:55.752 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 20:55:55.756 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 20:55:55.757 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 20:55:55.757 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 20:55:55.863 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 6.210771505s
         Current Time: 20:55:55
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.322028093s
         Current Time: 20:55:56
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.805µs
         Current Time: 20:55:56
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 updateProjectLinks   🔗 20:55:56.235 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 20:55:56.235 Links updated
💀    🚀 prepareDemoBackend   🔧 20:55:56.236 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 20:55:56.257 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 20:55:56.349 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 20:55:56.496 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 20:55:57.213 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 20:55:57.213 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 20:55:57.255 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 20:55:57.255  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 20:55:57.255 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 20:55:57.256 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 20:55:57.256  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 20:55:57.256 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:55:57.259  ---> Using cache
💀    🚀 buildDemoDbImage     🏭 20:55:57.259 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 20:55:57.259  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 20:55:57.259 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 20:55:57.259  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26   ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26  Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26   ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26  Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26   ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 20:55:57.26  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 20:55:57.261 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 20:55:57.262 Successfully built 40621c693b70
💀    🚀 buildDemoDbImage     🏭 20:55:57.264 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 20:55:57.264 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 20:55:57.265 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 20:55:57.267 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 20:55:57.267 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 20:55:57.614 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 20:55:57.631 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🔎 startDemoFrontend... 📗 20:55:57.644 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoFrontend... 📗 20:55:57.645 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoDbContainer 🐬 20:55:57.659 Error: No such container: demoDb
💀 🔥 🔎 startDemoDbContainer 🐬 20:55:57.666 Error: No such container: demoDb
💀 🔥 🚀 startDemoFrontend... 📗 20:55:57.671 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 20:55:57.672 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 20:55:57.689 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 20:55:57.69  🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 20:55:57.731 2b4abcdc6a5b7ce7c1314f4362849d8c65c984abe9090bada9fd0a1d13667799
💀    🚀 startDemoDbContainer 🐬 20:55:57.749 923d2f700cadf2e7c8deefd29a5cc5fd0decfd17122e22e9be25d5d6097ea21d
💀    🚀 prepareDemoBackend   🔧 20:55:58.105 Activate venv
💀    🚀 prepareDemoBackend   🔧 20:55:58.105 Install dependencies
💀    🚀 prepareDemoBackend   🔧 20:55:58.337 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 20:55:58.667   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:55:58.674 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 20:55:58.762   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:55:58.769 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 20:55:58.853   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoDbContainer 🐬 20:56:00.679 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 20:56:00.681 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.68 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.706 [38;5;6mmysql [38;5;5m13:56:00.70 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.709 [38;5;6mmysql [38;5;5m13:56:00.70 [38;5;2mINFO  ==> Initializing mysql database
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.719 [38;5;6mmysql [38;5;5m13:56:00.71 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀    🔎 startDemoDbContainer 🐬 20:56:00.719 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 20:56:00.719 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 20:56:00.72  🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.722 [38;5;6mmysql [38;5;5m13:56:00.72 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.73  [38;5;6mmysql [38;5;5m13:56:00.72 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.735 [38;5;6mmysql [38;5;5m13:56:00.73 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:00.74  [38;5;6mmysql [38;5;5m13:56:00.74 [38;5;2mINFO  ==> Installing database
💀    🚀 startDemoFrontend... 📗 20:56:01.094 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 20:56:01.096 🔎 Waiting docker container 'demoFrontend' healthcheck
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.126 [38;5;6mnginx [38;5;5m13:56:01.09 
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.126 [38;5;6mnginx [38;5;5m13:56:01.09 Welcome to the Bitnami nginx container
💀    🚀 startDemoFrontend... 📗 20:56:01.126 
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.126 [38;5;6mnginx [38;5;5m13:56:01.09 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.127 [38;5;6mnginx [38;5;5m13:56:01.10 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.127 [38;5;6mnginx [38;5;5m13:56:01.10 
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.127 [38;5;6mnginx [38;5;5m13:56:01.11 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🔎 startDemoFrontend... 📗 20:56:01.128 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 20:56:01.128 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 20:56:01.129 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 20:56:01.129 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 20:56:01.131 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.148 2022/05/11 13:56:01 [warn] 12#12: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 20:56:01.148 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 prepareDemoBackend   🔧 20:56:01.909 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 20:56:02.013   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:02.021 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 20:56:02.145   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:02.151 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 20:56:02.239   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:02.256 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 20:56:02.37    Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:02.38  Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 20:56:02.574   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:02.67  Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 20:56:03.056   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:03.141 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 20:56:03.28    Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:03.335 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 20:56:03.714   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🔎 startDemoDbContainer 🐬 20:56:03.723 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:03.844 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:03.846 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 20:56:03.977 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🔎 startDemoFrontend... 📗 20:56:04.133 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 20:56:04.185   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:04.204 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🔎 startDemoFrontend... 📗 20:56:04.245 check demoFrontend
💀    🔎 startDemoFrontend... 📗 20:56:04.249 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 20:56:04.286   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:04.294 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 20:56:04.386   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:04.392 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:04.635 [38;5;6mmysql [38;5;5m13:56:04.63 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 20:56:04.793   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:04.807 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 20:56:04.946   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:04.978 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 20:56:05.093   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.11  Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 20:56:05.296   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.313 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 20:56:05.418   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.428 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:05.625   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.652 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 20:56:05.775   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.804 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 20:56:05.908   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:05.92  Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 20:56:06.023   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 20:56:06.227 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 20:56:06.362   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:06.375 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:06.65  [38;5;6mmysql [38;5;5m13:56:06.64 [38;5;2mINFO  ==> Configuring authentication
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:06.691 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:06.711 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:06.749 [38;5;6mmysql [38;5;5m13:56:06.74 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:06.752 [38;5;6mmysql [38;5;5m13:56:06.75 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 20:56:06.814   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:06.915 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:06.985 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:06.986 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 20:56:07.025   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:07.038 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 20:56:07.149   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:07.154 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧 20:56:07.469   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:07.481 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 20:56:07.621   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:07.636 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 20:56:07.714   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 20:56:07.831 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧 20:56:07.981   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.003 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🔎 startDemoFrontend... 📗 20:56:08.251 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 20:56:08.251 📜 Task 'startDemoFrontendContainer' is ready
💀    🚀 prepareDemoBackend   🔧 20:56:08.289   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.301 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 20:56:08.377   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.38  Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:08.455   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.46  Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:08.596   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.605 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:08.704   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.729 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:08.766 [38;5;6mmysql [38;5;5m13:56:08.76 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 20:56:08.824   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.829 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:08.927   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:08.935 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:09.015   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.023 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 20:56:09.472   Using cached https://files.pythonhosted.org/packages/0c/58/25b4d208e0f6f00e19440385f360dc9891f8fa5ab62c11da52eb226fd9cd/coverage-6.3.2-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.485 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 20:56:09.585   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.591 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 20:56:09.753   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.761 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 20:56:09.85    Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.862 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 20:56:09.956   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:09.965 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 20:56:10.093   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:10.102 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:10.114 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:10.115 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 20:56:10.174   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 20:56:10.402 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, idna, urllib3, requests, fastavro, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, iniconfig, py, attrs, toml, pyparsing, packaging, pluggy, pytest, tomli, coverage, pytest-cov, pyasn1, rsa, ecdsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 20:56:10.425   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 20:56:10.597     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 20:56:10.932   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:13.224 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:13.225 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:14.775 [38;5;6mmysql [38;5;5m13:56:14.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:14.788 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:14.795 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:14.798 [38;5;6mmysql [38;5;5m13:56:14.79 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:16.332 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:16.334 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:16.81  [38;5;6mmysql [38;5;5m13:56:16.80 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 20:56:16.81  
💀 🔥 🚀 startDemoDbContainer 🐬 20:56:16.827 [38;5;6mmysql [38;5;5m13:56:16.82 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 20:56:17.011 2022-05-11T13:56:17.007216Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 20:56:17.011 2022-05-11T13:56:17.008428Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 20:56:17.011 2022-05-11T13:56:17.008435Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 20:56:17.013 2022-05-11T13:56:17.012601Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 20:56:17.143 2022-05-11T13:56:17.142947Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 20:56:17.299 2022-05-11T13:56:17.299279Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 20:56:17.299 2022-05-11T13:56:17.299326Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 20:56:17.311 2022-05-11T13:56:17.311219Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 20:56:17.311 2022-05-11T13:56:17.311257Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:19.453 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 Database
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 information_schema
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 mysql
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 performance_schema
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 sample
💀    🔎 startDemoDbContainer 🐬 20:56:19.458 sys
💀    🔎 startDemoDbContainer 🐬 20:56:19.461 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 20:56:23.463 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 20:56:23.463 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 20:56:32.346     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 20:56:33.419   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 20:56:33.589     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 20:56:34.079   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 20:56:34.217     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 20:56:34.264 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.2 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 20:56:34.316 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 20:56:34.316 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 20:56:34.36  Prepare
💀    🚀 prepareDemoBackend   🔧 20:56:34.36  prepare command
💀    🚀 prepareDemoBackend   🔧 20:56:34.36  Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 startDemoBackend     ⚡ 20:56:34.618 Activate venv
💀    🔎 startDemoBackend     ⚡ 20:56:34.618 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 20:56:34.618 Start
💀    🚀 startDemoBackend     ⚡ 20:56:34.978 2022-05-11 20:56:34,978 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 20:56:34.978 2022-05-11 20:56:34,978 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 20:56:34.984 2022-05-11 20:56:34,984 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 20:56:34.984 2022-05-11 20:56:34,984 INFO sqlalchemy.engine.Engine [generated in 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:34.986 2022-05-11 20:56:34,986 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 20:56:34.986 2022-05-11 20:56:34,986 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 20:56:34.988 2022-05-11 20:56:34,988 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:34.988 2022-05-11 20:56:34,988 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 20:56:34.988 2022-05-11 20:56:34,988 INFO sqlalchemy.engine.Engine [generated in 0.00011s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 2022-05-11 20:56:34,991 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 )
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 
💀    🚀 startDemoBackend     ⚡ 20:56:34.991 2022-05-11 20:56:34,991 INFO sqlalchemy.engine.Engine [no key 0.00009s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.011 2022-05-11 20:56:35,011 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 20:56:35.011 2022-05-11 20:56:35,011 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.032 2022-05-11 20:56:35,032 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 20:56:35.032 2022-05-11 20:56:35,032 INFO sqlalchemy.engine.Engine [no key 0.00051s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.07  2022-05-11 20:56:35,070 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 20:56:35.07  2022-05-11 20:56:35,070 INFO sqlalchemy.engine.Engine [no key 0.00022s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.093 2022-05-11 20:56:35,093 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 20:56:35.093 2022-05-11 20:56:35,093 INFO sqlalchemy.engine.Engine [no key 0.00024s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.112 2022-05-11 20:56:35,112 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 20:56:35.113 2022-05-11 20:56:35,113 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:35.114 2022-05-11 20:56:35,114 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 20:56:35.114 2022-05-11 20:56:35,114 INFO sqlalchemy.engine.Engine [cached since 0.1257s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 2022-05-11 20:56:35,115 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 )
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 
💀    🚀 startDemoBackend     ⚡ 20:56:35.115 2022-05-11 20:56:35,115 INFO sqlalchemy.engine.Engine [no key 0.00009s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.135 2022-05-11 20:56:35,135 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 20:56:35.135 2022-05-11 20:56:35,135 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.151 2022-05-11 20:56:35,151 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 20:56:35.151 2022-05-11 20:56:35,151 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.167 2022-05-11 20:56:35,166 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 20:56:35.167 2022-05-11 20:56:35,167 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.183 2022-05-11 20:56:35,183 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 20:56:35.184 2022-05-11 20:56:35,184 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:35.185 2022-05-11 20:56:35,185 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 20:56:35.185 2022-05-11 20:56:35,185 INFO sqlalchemy.engine.Engine [cached since 0.1968s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 2022-05-11 20:56:35,186 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.186 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 )
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 
💀    🚀 startDemoBackend     ⚡ 20:56:35.187 2022-05-11 20:56:35,186 INFO sqlalchemy.engine.Engine [no key 0.00010s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.208 2022-05-11 20:56:35,208 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 20:56:35.208 2022-05-11 20:56:35,208 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.222 2022-05-11 20:56:35,222 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 20:56:35.222 2022-05-11 20:56:35,222 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.239 2022-05-11 20:56:35,239 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 20:56:35.239 2022-05-11 20:56:35,239 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.255 2022-05-11 20:56:35,255 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 20:56:35.255 2022-05-11 20:56:35,255 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.271 2022-05-11 20:56:35,271 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 20:56:35.271 2022-05-11 20:56:35,271 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.289 2022-05-11 20:56:35,289 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 20:56:35.289 2022-05-11 20:56:35,289 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 20:56:35.305 2022-05-11 20:56:35,305 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 20:56:35.307 2022-05-11 20:56:35,307 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:35.309 2022-05-11 20:56:35,309 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 20:56:35.309 FROM users 
💀    🚀 startDemoBackend     ⚡ 20:56:35.309 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 20:56:35.309  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 20:56:35.309 2022-05-11 20:56:35,309 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 20:56:35.31  2022-05-11 20:56:35,310 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 20:56:35.496 2022-05-11 20:56:35,496 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:35.498 2022-05-11 20:56:35,497 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 20:56:35.498 2022-05-11 20:56:35,498 INFO sqlalchemy.engine.Engine [generated in 0.00020s] {'id': 'd9687cf1-035d-4db0-9838-396d6f7dcedc', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$Q7qJjCoFymnCvwItZti9iOKUe4K4CzCYSA0geUum0/i61hZQZUTfu', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 11, 20, 56, 35, 496218), 'updated_at': datetime.datetime(2022, 5, 11, 20, 56, 35, 497873)}
💀    🚀 startDemoBackend     ⚡ 20:56:35.499 2022-05-11 20:56:35,499 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 20:56:35.507 2022-05-11 20:56:35,507 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 20:56:35.508 2022-05-11 20:56:35,508 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 20:56:35.508 FROM users 
💀    🚀 startDemoBackend     ⚡ 20:56:35.508 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 20:56:35.508 2022-05-11 20:56:35,508 INFO sqlalchemy.engine.Engine [generated in 0.00011s] {'pk_1': 'd9687cf1-035d-4db0-9838-396d6f7dcedc'}
💀    🚀 startDemoBackend     ⚡ 20:56:35.509 2022-05-11 20:56:35,509 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 20:56:35.511 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.519 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 20:56:35.527 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 20:56:35.527 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.528 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.528 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 20:56:35.528 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 20:56:35.528 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.535 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 20:56:35.535 Register library route handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.535 Register library event handler
💀    🚀 startDemoBackend     ⚡ 20:56:35.535 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 20:56:35.535 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:35.536 INFO:     Started server process [22847]
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:35.536 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:35.536 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:35.536 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 20:56:35.621 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 20:56:35.621 check demoBackend
💀    🔎 startDemoBackend     ⚡ 20:56:35.621 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 20:56:35.621 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 20:56:35.728 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 39.600436435s
         Current Time: 20:56:35
         Active Process:
           * (PID=22425) 📗 'startDemoFrontendContainer' service
           * (PID=22842) ⚡ 'startDemoBackend' service
           * (PID=22447) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=22425)
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=22842)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=22447)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:37.346 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:37.446 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:37.447 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 20:56:37.447 INFO:     Finished server process [22847]
💀    🚀 startDemoBackend     ⚡ 20:56:37.569 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 20:56:37.569 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 41.703954504s
         Current Time: 20:56:37
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 2.426µs
         Current Time: 20:56:38
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 20:56:38.132 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 20:56:38.132 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 20:56:38.161 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 20:56:38.246 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoDbImage     🏭 20:56:38.391 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 20:56:38.392 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 20:56:38.99  Sending build context to Docker daemon  16.38kB
💀    🚀 buildDemoDbImage     🏭 20:56:38.991 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoDbImage     🏭 20:56:39.058 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 20:56:39.058  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 20:56:39.058 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 20:56:39.059 Sending build context to Docker daemon  1.179MB
💀    🚀 buildDemoDbImage     🏭 20:56:39.062 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 20:56:39.066 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 20:56:39.066 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07  Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07   ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07  Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07   ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07  Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07   ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 20:56:39.07  Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 20:56:39.071 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 20:56:39.078 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 20:56:39.078  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 20:56:39.078 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:56:39.078  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.078  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 20:56:39.079 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 20:56:39.081 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 20:56:39.082  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.082  ---> 808ba8676c5f
💀    🚀 buildDemoFrontend... 🏭 20:56:39.082 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083  ---> 0c9047d38d7d
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083  ---> 99c8982165ff
💀    🚀 buildDemoFrontend... 🏭 20:56:39.083 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> 3bacbc306156
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> 0e12772b83fe
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084  ---> 8072400998af
💀    🚀 buildDemoFrontend... 🏭 20:56:39.084 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 20:56:39.085  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 20:56:39.085  ---> 00baf0e406aa
💀    🚀 buildDemoFrontend... 🏭 20:56:39.086 Successfully built 00baf0e406aa
💀    🚀 buildDemoFrontend... 🏭 20:56:39.092 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 20:56:39.094 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 20:56:39.094 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 20:56:39.232  ---> fa1fa6639e90
💀    🚀 buildDemoBackendI... 🏭 20:56:39.232 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 20:56:39.253  ---> Running in a423dc7ac509
💀    🚀 buildDemoBackendI... 🏭 20:56:39.318 Removing intermediate container a423dc7ac509
💀    🚀 buildDemoBackendI... 🏭 20:56:39.318  ---> 7e70d2f60475
💀    🚀 buildDemoBackendI... 🏭 20:56:39.318 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 20:56:39.37   ---> Running in c894c239e25a
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 20:56:39.432 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 20:56:39.463 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 20:56:39.476 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 20:56:39.476 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 20:56:39.477 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoDbContainer 🐬 20:56:39.507 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 20:56:39.508 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 20:56:39.519 🔎 Waiting docker container 'demoDb' healthcheck
💀    🔎 startDemoFrontend... 📗 20:56:39.537 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 20:56:39.537 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 20:56:39.538 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 20:56:39.538 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 20:56:39.54  🔎 Host port '443' is ready
💀    🔎 startDemoDbContainer 🐬 20:56:39.56  🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 20:56:39.56  🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 20:56:39.561 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 20:56:40.211 Removing intermediate container c894c239e25a
💀    🚀 buildDemoBackendI... 🏭 20:56:40.211  ---> 6a02e101d011
💀    🚀 buildDemoBackendI... 🏭 20:56:40.211 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 20:56:40.248  ---> Running in 46883e20ca8c
💀    🚀 buildDemoBackendI... 🏭 20:56:40.375 Removing intermediate container 46883e20ca8c
💀    🚀 buildDemoBackendI... 🏭 20:56:40.375  ---> 984da7f983bf
💀    🚀 buildDemoBackendI... 🏭 20:56:40.378 Successfully built 984da7f983bf
💀    🚀 buildDemoBackendI... 🏭 20:56:40.384 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 20:56:40.385 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 20:56:40.385 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 20:56:42.544 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 20:56:42.564 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 20:56:42.634 check demoFrontend
💀    🔎 startDemoFrontend... 📗 20:56:42.637 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 20:56:42.648 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 20:56:42.65  Database
💀    🔎 startDemoDbContainer 🐬 20:56:42.65  information_schema
💀    🔎 startDemoDbContainer 🐬 20:56:42.651 mysql
💀    🔎 startDemoDbContainer 🐬 20:56:42.651 performance_schema
💀    🔎 startDemoDbContainer 🐬 20:56:42.651 sample
💀    🔎 startDemoDbContainer 🐬 20:56:42.651 sys
💀    🔎 startDemoDbContainer 🐬 20:56:42.654 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 20:56:46.639 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 20:56:46.639 📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 20:56:46.657 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 20:56:46.657 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 20:56:47.196 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:47.218 Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ 20:56:47.219 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:47.24  Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 20:56:47.241 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 20:56:47.281 7f2426f86d5ebc19a86000ef732da934cdff7f26215684675cd1aee8f85e5c2b
💀    🚀 startDemoBackendC... ⚡ 20:56:48.371 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 20:56:48.373 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 20:56:48.398 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 20:56:48.398 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 20:56:48.4   🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 20:56:48.796 2022-05-11 13:56:48,795 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 20:56:48.796 2022-05-11 13:56:48,796 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.8   2022-05-11 13:56:48,800 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 20:56:48.8   2022-05-11 13:56:48,800 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.805 2022-05-11 13:56:48,805 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 20:56:48.805 2022-05-11 13:56:48,805 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.809 2022-05-11 13:56:48,809 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 20:56:48.81  2022-05-11 13:56:48,810 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 20:56:48.81  2022-05-11 13:56:48,810 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.813 2022-05-11 13:56:48,812 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 20:56:48.815 2022-05-11 13:56:48,815 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 20:56:48.816 2022-05-11 13:56:48,816 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 20:56:48.816 2022-05-11 13:56:48,816 INFO sqlalchemy.engine.Engine [cached since 0.006268s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.819 2022-05-11 13:56:48,818 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 20:56:48.822 2022-05-11 13:56:48,821 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 20:56:48.823 2022-05-11 13:56:48,822 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 20:56:48.823 2022-05-11 13:56:48,823 INFO sqlalchemy.engine.Engine [cached since 0.01321s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.826 2022-05-11 13:56:48,826 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 20:56:48.831 2022-05-11 13:56:48,831 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 20:56:48.834 2022-05-11 13:56:48,833 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 20:56:48.834 FROM users 
💀    🚀 startDemoBackendC... ⚡ 20:56:48.834 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 20:56:48.834  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 20:56:48.834 2022-05-11 13:56:48,833 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 20:56:48.837 2022-05-11 13:56:48,836 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 20:56:48.84  Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.85  Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 20:56:48.86  Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.867 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 20:56:48.867 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.867 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 20:56:48.867 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 20:56:48.867 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:48.867 INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:48.867 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:48.868 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 20:56:48.868 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 20:56:51.402 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 20:56:51.507 check demoBackend
💀    🔎 startDemoBackendC... ⚡ 20:56:51.513 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 20:56:52.514 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 20:56:52.515 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 20:56:52.622 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 14.599485945s
         Current Time: 20:56:52
         Active Process:
           * (PID=23940) 🐬 'startDemoDbContainer' service
           * (PID=24073) ⚡ 'startDemoBackendContainer' service
           * (PID=23911) 📗 'startDemoFrontendContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=24073)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=23911)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=23940)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 16.703389208s
         Current Time: 20:56:54
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 2.05µs
         Current Time: 20:56:55
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 20:56:55.027 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 20:56:55.027 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoDbContainer  ✋ 20:56:55.376 Stop docker container demoDb
💀    🚀 stopDemoBackendCo... ✋ 20:56:55.381 Stop docker container demoBackend
💀    🚀 stopDemoFrontendC... ✋ 20:56:55.387 Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ 20:56:59.368 demoDb
💀    🚀 stopDemoDbContainer  ✋ 20:56:59.369 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 20:56:59.369 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoFrontendC... ✋ 20:57:06.395 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 20:57:06.397 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 20:57:06.397 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀    🚀 stopDemoBackendCo... ✋ 20:57:06.526 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 20:57:06.528 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 20:57:06.528 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 20:57:06.633 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.712079319s
         Current Time: 20:57:06
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.823515813s
         Current Time: 20:57:06
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.328µs
         Current Time: 20:57:07
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 20:57:07.015 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 20:57:07.015 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🔥 🚀 removeDemoFronten... ❌ 20:57:07.326 Error: No such container: 
💀 🔥 🚀 removeDemoBackend... ❌ 20:57:07.329 Error: No such container: 
💀    🚀 removeDemoBackend... ❌ 20:57:07.331 Stop docker container demoBackend
💀    🚀 removeDemoFronten... ❌ 20:57:07.332 Stop docker container demoFrontend
💀 🔥 🚀 removeDemoDbConta... ❌ 20:57:07.346 Error: No such container: 
💀    🚀 removeDemoDbConta... ❌ 20:57:07.346 Stop docker container demoDb
💀    🚀 removeDemoBackend... ❌ 20:57:07.433 Docker container demoBackend stopped
💀    🚀 removeDemoBackend... ❌ 20:57:07.433 Remove docker container demoBackend
💀    🚀 removeDemoFronten... ❌ 20:57:07.438 Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ 20:57:07.438 Remove docker container demoFrontend
💀    🚀 removeDemoDbConta... ❌ 20:57:07.445 Docker container demoDb stopped
💀    🚀 removeDemoDbConta... ❌ 20:57:07.445 Remove docker container demoDb
💀    🚀 removeDemoBackend... ❌ 20:57:07.501 demoBackend
💀    🚀 removeDemoBackend... ❌ 20:57:07.503 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ 20:57:07.503 Docker container demoBackend removed
💀    🚀 removeDemoFronten... ❌ 20:57:07.504 demoFrontend
💀    🚀 removeDemoFronten... ❌ 20:57:07.505 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ 20:57:07.505 Docker container demoFrontend removed
💀    🚀 removeDemoDbConta... ❌ 20:57:07.524 demoDb
💀    🚀 removeDemoDbConta... ❌ 20:57:07.526 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ 20:57:07.526 Docker container demoDb removed
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 removeContainers     ❌ 20:57:07.634 
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 724.889176ms
         Current Time: 20:57:07
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 836.35812ms
         Current Time: 20:57:07
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.884µs
         Current Time: 20:57:08
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:08.187 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:57:08.192         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:57:08.192     
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:57:08.192   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:57:08.192   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:57:08.192   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:57:08.192 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:08.764 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:08.796 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.092 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.092 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.093 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.094 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.095 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.095 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.095 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.413 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.434 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.441 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.441 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.441 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.441 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.441 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.445 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.445 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.456 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.456 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.459 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.46  Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.463 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.495 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.495 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:09.496 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10     🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10     Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.262 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.521 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.528 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.534 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.534 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.534 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.534 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.534 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.537 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.537 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.546 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.546 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.549 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.549 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.552 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.571 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.574 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.577 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.739 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.894 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:10.898 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.079 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.246 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.25  Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.416 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.416 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:11.416 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.339239943s
         Current Time: 20:57:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.45100908s
         Current Time: 20:57:11
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.003µs
         Current Time: 20:57:11
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:11.787 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:57:11.791         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:57:11.791     
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:57:11.791   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:57:11.791   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:57:11.791   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:57:11.791 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.216 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.216 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:12.906 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.114 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.121 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.127 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.127 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.127 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.127 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.127 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.13  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.13  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.139 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.139 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.142 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.142 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.145 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.175 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.176 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.176 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.528 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:13.528 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.221 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.221 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.222 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.434 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.44  Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.446 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.446 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.446 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.446 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.446 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.449 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.449 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.458 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.458 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.461 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.461 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.464 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.482 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.486 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.489 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.647 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.798 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.801 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:14.954 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:15.109 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:15.112 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:15.272 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:15.272 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:15.272 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.590762903s
         Current Time: 20:57:15
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.702730437s
         Current Time: 20:57:15
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.31µs
         Current Time: 20:57:15
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:15.648 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 20:57:15.65          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 20:57:15.65      
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 20:57:15.65    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 20:57:15.65    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 20:57:15.65    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 20:57:15.65  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.081 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.081 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.199 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.199 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.2   Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.443 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.45  Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.457 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.457 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.457 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.457 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.457 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.46  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.461 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.472 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.472 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.476 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.476 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.48  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.48  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.48  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.481   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.481 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.481 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.517 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.517 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.517 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.804 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.804 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.921 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:16.922 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.155 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.162 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.17  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.17  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.17  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.17  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.17  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.173 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.173 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.182 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.182 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.185 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.186 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189 ]
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.189 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.208 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.212 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.216 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.397 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.575 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.58  Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.755 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.921 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:17.924 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 20:57:18.1   Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 20:57:18.1   🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 20:57:18.1   Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.558755354s
         Current Time: 20:57:18
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.669564842s
         Current Time: 20:57:18
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.339µs
         Current Time: 20:57:18
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:18.486 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 20:57:18.591 Synchronize task environments
💀    🚀 syncEnv              🔄 20:57:18.767 Synchronize project's environment files
💀    🚀 syncEnv              🔄 20:57:18.924 🎉🎉🎉
💀    🚀 syncEnv              🔄 20:57:18.924 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 544.430592ms
         Current Time: 20:57:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 745.196987ms
         Current Time: 20:57:19
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.51µs
         Current Time: 20:57:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:19.392 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 20:57:19.503 🎉🎉🎉
💀    🚀 setProjectValue      🔗 20:57:19.503 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.769259ms
         Current Time: 20:57:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.438322ms
         Current Time: 20:57:19
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.261µs
         Current Time: 20:57:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 20:57:19.975 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 20:57:20.085 🎉🎉🎉
💀    🚀 setProjectValue      🔗 20:57:20.085 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.97237ms
         Current Time: 20:57:20
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 417.448652ms
         Current Time: 20:57:20
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.844µs
         Current Time: 20:57:20
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoDbDepl... 🏁 20:57:20.57  🚧 Create virtual environment.
💀    🚀 prepareDemoFronte... 🏁 20:57:20.57  🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 20:57:20.57  🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 20:57:22.467 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 20:57:22.499 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 20:57:22.547 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:22.739 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:22.755 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:22.811 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:23.519   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:23.523   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:23.534 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:23.539 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:23.636   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:23.653 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:23.739   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 20:57:23.742   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 20:57:23.912   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 20:57:23.995 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:24.009 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.065   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.071 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:24.087   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:24.095 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:24.21  Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:24.244   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:24.257 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:24.345   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:24.365 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.4     Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.423 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:24.546   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.557   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:24.569 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.576 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:24.636   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:24.658 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:24.699   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.705   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:24.708 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.713 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:24.741   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:24.751 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.79    Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:24.801 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:24.838   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:24.843 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 20:57:25.005   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:25.01  Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:25.379   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.451 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 20:57:25.575   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.599 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.6     Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.679 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:25.683   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.695 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.768   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.788 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:25.802   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.808   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.824 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:25.866 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:25.912   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.92    Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:25.925 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:25.933 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:25.946   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:25.956 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:26.004   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:26.032 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.043   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.055 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:26.071   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.085 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.139   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:26.155   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.171 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:26.178 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:26.218   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.229 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.266   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.284 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:26.314   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:26.317   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:26.324 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:26.338 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.406   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.414 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:26.425   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.43    Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 20:57:26.434 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 20:57:26.44  Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.505   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.511 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:26.526   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.543   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.549 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 20:57:26.589 Installing collected packages: six, dill, protobuf, semver, grpcio, pyyaml, pulumi, arpeggio, attrs, parver, urllib3, charset-normalizer, certifi, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.633   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.682   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.692 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 20:57:26.708 Installing collected packages: six, protobuf, pyyaml, semver, dill, grpcio, pulumi, arpeggio, attrs, parver, urllib3, charset-normalizer, idna, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 20:57:26.797   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 20:57:26.919 Installing collected packages: pyyaml, protobuf, dill, six, semver, grpcio, pulumi, arpeggio, attrs, parver, charset-normalizer, idna, certifi, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 20:57:27.214   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 20:57:27.363   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 20:57:27.673   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 20:57:28.491     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 20:57:28.529 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 20:57:28.554 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 20:57:28.554 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.662     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 20:57:28.693 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 20:57:28.693 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.699 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoDbDepl... 🏁 20:57:28.729 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 20:57:28.729 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 dependencies.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751     dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 20:57:28.751       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752     dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 for this case.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 Usage:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 Aliases:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 20:57:28.752 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 Flags:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 
💀    🚀 prepareDemoFronte... 🏁 20:57:28.753 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 prepareDemoBacken... 🏁 20:57:28.905     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.93  🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.93  🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 20:57:28.947 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.977       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.978   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:28.979 Use "helm dependency [command] --help" for more information about a command.
💀 🔥 🚀 prepareDemoBacken... 🏁 20:57:28.98  WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 20:57:28.98  You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 deployDemoFronten... 🏁 20:57:29.07  error: no stack named 'dev' found
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 deployDemoFronten... 🏁 20:57:29.175 Created stack 'dev'
💀 🔥 🚀 deployDemoDbDeplo... 🏁 20:57:29.282 error: no stack named 'dev' found
💀    🚀 deployDemoDbDeplo... 🏁 20:57:29.374 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 20:57:29.81  PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 20:57:29.889 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 20:57:29.889 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 dependencies.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 for this case.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Usage:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Aliases:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Flags:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 
💀    🚀 prepareDemoBacken... 🏁 20:57:29.955 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 20:57:29.957 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 20:57:30.296 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 20:57:30.408 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 20:57:30.799 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 20:57:31.019 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 20:57:31.188 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:31.392 
💀    🚀 deployDemoFronten... 🏁 20:57:31.532  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 20:57:31.593  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:31.754  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:31.838  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 20:57:31.851  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 20:57:31.855  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 20:57:31.97   +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 20:57:31.97   
💀    🚀 deployDemoFronten... 🏁 20:57:31.97  Resources:
💀    🚀 deployDemoFronten... 🏁 20:57:31.97      + 4 to create
💀    🚀 deployDemoFronten... 🏁 20:57:31.97  
💀    🚀 deployDemoFronten... 🏁 20:57:31.97  Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.103  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.104  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoBackend... 🏁 20:57:32.127 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257  
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.257 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 20:57:32.461 
💀    🚀 deployDemoBackend... 🏁 20:57:32.607 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:32.746 
💀    🚀 deployDemoFronten... 🏁 20:57:33.003  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.091  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:33.187  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 20:57:33.267  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.316  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.394  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.415  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.418  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.436  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.439  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 20:57:33.442  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 20:57:33.449  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 20:57:33.599  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 20:57:33.599  
💀    🚀 deployDemoFronten... 🏁 20:57:33.6   Outputs:
💀    🚀 deployDemoFronten... 🏁 20:57:33.6       app: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6           ready    : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.6               [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6               [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6           ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.6           resources: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6               apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                   api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                   id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                   kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                   metadata   : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                       annotations       : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                           kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                               apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                               kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                               metadata  : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   annotations: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   }
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   labels     : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   }
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                               }
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                               spec      : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   replicas: 1
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                   selector: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                       matchLabels: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                           app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.6                                           app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                 template: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                         labels: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                 ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                         ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601 
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                     creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                     labels            : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.601                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                         [0]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.602                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             time       : "2022-05-11T13:57:33Z"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     resource_version  : "60385"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     uid               : "820b45eb-f57f-4be8-8ca5-91e75d9b7f0c"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                 spec       : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     replicas                 : 1
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     selector                 : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         match_labels: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                     template                 : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         metadata: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             labels: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                 app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                 app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                         spec    : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                             containers                      : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                 [0]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                     env                       : [
💀    🚀 deployDemoFronten... 🏁 20:57:33.603                                         [0]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         [1]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         [2]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         [3]: {
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                             value: "1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             ]
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                         }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                     }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                 }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604             }
💀    🚀 deployDemoFronten... 🏁 20:57:33.604             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 20:57:33.604                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 20:57:33.605                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.605                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 20:57:33.605                
💀    🚀 deployDemoBackend... 🏁 20:57:33.609  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 20:57:33.61   +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 20:57:33.613  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.692  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.695  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.707  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.709  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.716  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.721  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoBackend... 🏁 20:57:33.79   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 20:57:33.79   
💀    🚀 deployDemoBackend... 🏁 20:57:33.79  Resources:
💀    🚀 deployDemoBackend... 🏁 20:57:33.79      + 5 to create
💀    🚀 deployDemoBackend... 🏁 20:57:33.79  
💀    🚀 deployDemoBackend... 🏁 20:57:33.791 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.959  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.959  
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961     app: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961         ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.961                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             time       : "2022-05-11T13:57:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     resource_version  : "60399"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     uid               : "6a593876-bf7a-4925-b33c-701083bb6672"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.962                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.963                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.964 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                     creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                             api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.965                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                                     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                             time       : "2022-05-11T13:57:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                     ]
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                     resource_version  : "60400"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                     uid               : "0da82e0f-01c4-4c67-9dae-d1ff684b7558"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                 }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966             }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966         }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966     }
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966 Duration: 2s
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.966 
💀    🚀 deployDemoDbDeplo... 🏁 20:57:33.967 hello world
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoBackend... 🏁 20:57:34.326 
💀    🚀 deployDemoBackend... 🏁 20:57:34.795  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 20:57:34.902  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.135  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.137  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.141  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.147  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.149  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.153  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 20:57:35.156  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 20:57:35.166  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 20:57:35.166  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 20:57:35.447  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 20:57:35.447  
💀    🚀 deployDemoBackend... 🏁 20:57:35.449 Outputs:
💀    🚀 deployDemoBackend... 🏁 20:57:35.449     app: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.449         ready    : [
💀    🚀 deployDemoBackend... 🏁 20:57:35.449             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449         ]
💀    🚀 deployDemoBackend... 🏁 20:57:35.449         resources: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.449             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.449                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 20:57:35.449                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                      annotations       : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                          kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                              apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                              kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                              metadata  : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  annotations: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  labels     : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                              }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                              spec      : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  replicas: 1
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  selector: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      matchLabels: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                          app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                          app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                  template: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      metadata: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                          labels: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                              app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                              app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                          }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                      spec    : {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                          containers        : [
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                              [0]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                  env            : [
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [0]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "HS256"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [1]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "30"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [2]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [3]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "/token/"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [4]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [5]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          value: "1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      }
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                      [6]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.45                                                          name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.451                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.452                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 20:57:35.453                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 20:57:35.454                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     }
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                 ]
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 20:57:35.455                                                         contain
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 20:57:35.561 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 15.107040319s
         Current Time: 20:57:35
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 15.219216407s
         Current Time: 20:57:35
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.424µs
         Current Time: 20:57:35
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 20:57:35.999 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:35.999 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 20:57:36     🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 20:57:36.374 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.384 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.384 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.389 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.39  Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.391 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.391 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.393 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.396 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.397 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.398 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.399 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.4   Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.401 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.402 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.403 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.404 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.404 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.407 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.408 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.41  Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.411 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.413 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.418 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.418 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.421 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.422 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.424 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.433 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.436 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.438 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.446 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.452 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.452 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.488 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.491 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.492 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.493 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.496 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.5   Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.503 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.504 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.507 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.507 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.51  Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 20:57:36.514 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 20:57:36.532 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.543 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoBacken... 🏁 20:57:36.554 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 20:57:36.554 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 20:57:36.561 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 20:57:36.561 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 20:57:36.563 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 20:57:36.563 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.796 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 20:57:36.797 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 dependencies.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887     dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 20:57:36.887 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888     dependencies:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 for this case.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Usage:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Aliases:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Flags:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.888       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889 
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 20:57:36.889 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.906 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.906 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.956 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.956 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.956 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.957 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.958 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959 
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 20:57:36.959 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 20:57:37.753 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 20:57:37.821 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 20:57:37.822 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 dependencies.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872     dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.872 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873     dependencies:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 for this case.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Usage:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Aliases:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Flags:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.873       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874 
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 20:57:37.874 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 20:57:38.499 Previewing destroy (dev):
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.522 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 20:57:38.583 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.584  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.585  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.587  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.588  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.589  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.589  
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59  Outputs:
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59    - app: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59        - ready    : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59        -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59        -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59          ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59        - resources: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59            - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                            - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                  }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                  }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                              }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                            - spec      : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                - selector: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                    - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                        - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                        - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.59                                      }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                               - template: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                                 ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                         ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.591                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                     ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   - resource_version  : "60385"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.592                   - uid               : "820b45eb-f57f-4be8-8ca5-91e75d9b7f0c"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.593                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                     ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.594                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                             ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.595                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                     ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   - resource_version  : "60384"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                   - uid               : "117f8846-2648-4333-9a98-489ce83fc923"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596 Resources:
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.596 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.615 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.616  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.617  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.618  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.619  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.621  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.621  
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623         ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.623                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.624                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.625                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.626                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.627                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                     ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.628                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - resource_version  : "60399"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - uid               : "6a593876-bf7a-4925-b33c-701083bb6672"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.629                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                            - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                      ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                    - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                                  }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                              ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.63                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.631                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                     ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                   - resource_version  : "60400"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                   - uid               : "0da82e0f-01c4-4c67-9dae-d1ff684b7558"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.632 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 20:57:38.67  
💀    🚀 destroyDemoFronte... 🏁 20:57:38.671  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.672  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.703 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.709  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.709  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.767  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.771  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.771  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.771  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.774  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.777  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.777  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.778  
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779 Outputs:
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779   - app: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779         ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779       - resources: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                               - template: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.779                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                      }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                        - value: "1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.78                                                      }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                                 ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                         ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781 
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.781                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                     ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - resource_version  : "60385"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - uid               : "820b45eb-f57f-4be8-8ca5-91e75d9b7f0c"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.782                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                     ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                             ]
💀    🚀 destroyDemoFronte... 🏁 20:57:38.783                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                         }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                     }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784             }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                                 }
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 20:57:38.784         
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.814  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.821  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.821  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.824  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.829  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.834  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.834  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.834  
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835         ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.835                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.836                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837 
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                   - creation_timestamp: "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.837                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - time       : "2022-05-11T13:57:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                     ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - resource_version  : "60399"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - uid               : "6a593876-bf7a-4925-b33c-701083bb6672"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                 }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.838                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                     }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                             }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                         }
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.839                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 20:57:38.84                                    - image_pull_policy         : "IfNotPres
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 20:57:39.454 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 20:57:39.529 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.531  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.531  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.535  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.537  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.539  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.541  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.541  
💀    🚀 destroyDemoBacken... 🏁 20:57:39.543 Outputs:
💀    🚀 destroyDemoBacken... 🏁 20:57:39.543   - app: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.543       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.543       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544         ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544       - resources: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.544                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                               - template: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.545                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.546                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.547                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.548                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.549                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.55                                                        - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                 ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                                 ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                         ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.551                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                   - creation_timestamp: "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.552                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.553                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.554                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                           - time       : "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - resource_version  : "60424"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - uid               : "01c2e109-71ea-430e-8bb5-72374085cc4e"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 20:57:39.555                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.556                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.557                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.558                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.559                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                          }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                    -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.56                                            - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.561                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                             ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.562                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.563                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                 ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                   - creation_timestamp: "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.564                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.565                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                           - time       : "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - resource_version  : "60426"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - uid               : "75c414b0-fc5f-4842-b390-d67a81d9aff9"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - cluster_ip             : "10.106.154.158"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   -     [0]: "10.106.154.158"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 20:57:39.566                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - status     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.567                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                   - creation_timestamp: "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.568                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                           - time       : "2022-05-11T13:57:35Z"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                     ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                   - resource_version  : "60423"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                   - uid               : "0429445c-fd65-4f12-814e-c8a65ac9623f"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569 Resources:
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.569 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 20:57:39.628 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.629  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.629  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.636  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.721  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.723  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.723  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.724  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.724  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.729  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.736  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.746  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.746  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 20:57:39.746  
💀    🚀 destroyDemoBacken... 🏁 20:57:39.746 Outputs:
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747   - app: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747         ]
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747       - resources: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                             }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                 }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                               - template: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.747                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                         }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.748                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                     }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.749                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "root"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "local"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                      }
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 20:57:39.75                          
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 20:57:39.855 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 3.972891626s
         Current Time: 20:57:39
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.084145593s
         Current Time: 20:57:40
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

