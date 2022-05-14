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
echo "API_HOST=http://localhost:3000" > demoFrontend/template.env
zaruba please syncEnv

zaruba task setConfigs startDemoFrontendContainer localhost localhost

# Add bootstrap
echo 'echo "var apiHost=\"$API_HOST\";" > /opt/bitnami/nginx/html/apiHost.js && /opt/bitnami/scripts/nginx/run.sh' > demoFrontend/bootstrap.sh

# Overwrite index.html
cp ../../use-cases/from-zero-to-cloud/index.html demoFrontend/html/index.html

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
         Elapsed Time: 1.647µs
         Current Time: 07:50:46
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 07:50:46.124 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 07:50:46.129 🎉🎉🎉
💀    🚀 initProject          🚧 07:50:46.129 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 117.055328ms
         Current Time: 07:50:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 318.384137ms
         Current Time: 07:50:46
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 981ns
         Current Time: 07:50:46
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:50:46.586 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:50:46.589         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:50:46.589     
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:50:46.589   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:50:46.589   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:50:46.589   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:50:46.589 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 07:50:46.906 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 07:50:46.906 Preparing base variables
💀    🚀 makeMysqlApp         🐬 07:50:46.993 Base variables prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing start command
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Start command prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing test command
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Test command prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing check command
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Check command prepared
💀    🚀 makeMysqlApp         🐬 07:50:46.994 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 07:50:47.23  Add config to replacement map
💀    🚀 makeMysqlApp         🐬 07:50:47.238 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 07:50:47.246 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 07:50:47.246 ✅ Validate
💀    🚀 makeMysqlApp         🐬 07:50:47.246 Validate app directory
💀    🚀 makeMysqlApp         🐬 07:50:47.246 Done validating app directory
💀    🚀 makeMysqlApp         🐬 07:50:47.246 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 07:50:47.249 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 07:50:47.249 Validate template locations
💀    🚀 makeMysqlApp         🐬 07:50:47.259 Done validating template locations
💀    🚀 makeMysqlApp         🐬 07:50:47.259 Validate app ports
💀    🚀 makeMysqlApp         🐬 07:50:47.264 Done validating app ports
💀    🚀 makeMysqlApp         🐬 07:50:47.264 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 07:50:47.268 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 07:50:47.268 🚧 Generate
💀    🚀 makeMysqlApp         🐬 07:50:47.268 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 07:50:47.268   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 07:50:47.268 ]
💀    🚀 makeMysqlApp         🐬 07:50:47.268 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 07:50:47.287 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 07:50:47.287 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 07:50:47.288 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.776 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.776 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.949 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.949 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:47.95  Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.196 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.204 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.21  Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.21  ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.21  Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.21  Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.21  Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.214 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.214 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.23  Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.23  Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.234 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.234 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.24  Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.24  🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.24  🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.24    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.24    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.241   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.241 ]
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.241 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.279 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.284 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.288 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.456 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.626 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.796 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.971 Checking start
💀    🚀 makeMysqlAppRunner   🐬 07:50:48.974 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.138 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.305 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.309 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.474 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.629 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.796 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.956 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:49.96  Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.123 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.289 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.292 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.458 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.629 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.633 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.801 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.971 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:50:50.974 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:50:51.144 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:50:51.319 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:50:51.322 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 07:50:51.322 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.842172572s
         Current Time: 07:50:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.953554433s
         Current Time: 07:50:51
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 967ns
         Current Time: 07:50:51
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:50:51.702 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:50:51.705         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:50:51.705     
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:50:51.705   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:50:51.705   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:50:51.705   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:50:51.705 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 07:50:52.148 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 07:50:52.149 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 07:50:52.291 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.291 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 07:50:52.291 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.291 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 07:50:52.291 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.292 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 07:50:52.519 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:50:52.526 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:50:52.534 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 07:50:52.534 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 07:50:52.534 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 07:50:52.534 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 07:50:52.534 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:50:52.538 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:50:52.538 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 07:50:52.548 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 07:50:52.548 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 07:50:52.551 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 07:50:52.551 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554 ]
💀    🚀 makeFastApiApp       ⚡ 07:50:52.554 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 07:50:53.224 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 07:50:53.225 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 07:50:53.225 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 07:50:53.705 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:50:53.705 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.737 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.737 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.737 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:54.738 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.011 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.021 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.03  Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.03  ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.03  Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.03  Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.03  Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.034 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.034 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.052 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.052 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.055 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.055 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.058 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.058 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.058 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059 ]
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059 
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.059 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.108 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.111 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.115 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.284 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.287 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.456 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.623 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.626 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.799 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.977 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:50:55.979 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.151 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.322 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.325 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.502 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.68  Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.684 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:56.905 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.107 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.313 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.503 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.507 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.701 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.907 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:57.912 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.108 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.297 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.301 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.495 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.682 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.685 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:50:58.852 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.042 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.231 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.402 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.598 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.793 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:50:59.987 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:51:00.162 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:51:00.331 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 07:51:00.331 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 07:51:00.835 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 07:51:00.835 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 07:51:01.884 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.884 Preparing start command
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Start command prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Preparing test command
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Test command prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Preparing check command
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Check command prepared
💀    🚀 addFastApiModule     ⚡ 07:51:01.885 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 07:51:02.161 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 07:51:02.169 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 07:51:02.178 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 07:51:02.178 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 07:51:02.178 Validate app directory
💀    🚀 addFastApiModule     ⚡ 07:51:02.178 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 07:51:02.178 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 07:51:02.181 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 07:51:02.182 Validate template locations
💀    🚀 addFastApiModule     ⚡ 07:51:02.192 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 07:51:02.192 Validate app ports
💀    🚀 addFastApiModule     ⚡ 07:51:02.196 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 07:51:02.196 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 07:51:02.2   Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 07:51:02.2   🚧 Generate
💀    🚀 addFastApiModule     ⚡ 07:51:02.2   🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 07:51:02.2     "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 07:51:02.2   ]
💀    🚀 addFastApiModule     ⚡ 07:51:02.2   🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 07:51:02.217 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 07:51:02.217 Registering module
💀    🚀 addFastApiModule     ⚡ 07:51:02.247 Done registering module
💀    🚀 addFastApiModule     ⚡ 07:51:02.248 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 07:51:02.248 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 07:51:02.613 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 07:51:02.613 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 07:51:03.744 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:03.744 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:51:03.744 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:03.744 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:51:03.744 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:03.745 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:51:03.745 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:03.745 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:51:03.745 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:03.745 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.016 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.024 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.031 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.031 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:51:04.043 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:51:04.043 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:51:04.117 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:51:04.117 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:51:04.186 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:51:04.186 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 07:51:04.307 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 07:51:04.307 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.378 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.662 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.669 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:51:04.678 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:51:04.678 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 07:51:04.678 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 07:51:04.678 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 07:51:04.678 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:51:04.682 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:51:04.682 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 07:51:04.692 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 07:51:04.692 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 07:51:04.696 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 07:51:04.696 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7     "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   ]
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   
💀    🚀 addFastApiCrud       ⚡ 07:51:04.7   🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 07:51:04.735 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 07:51:04.735 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:51:04.78  Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:51:04.78  Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:51:04.842 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:51:04.842 Registering repo
💀    🚀 addFastApiCrud       ⚡ 07:51:04.914 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 07:51:04.915 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 07:51:04.915 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 13.31891986s
         Current Time: 07:51:05
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 13.520212444s
         Current Time: 07:51:05
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.306µs
         Current Time: 07:51:05
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:51:05.406 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:51:05.409         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:51:05.409     
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:51:05.409   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:51:05.409   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:51:05.409   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:51:05.409 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 07:51:05.856 🧰 Prepare
💀    🚀 makeNginxApp         📗 07:51:05.856 Preparing base variables
💀    🚀 makeNginxApp         📗 07:51:05.937 Base variables prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing start command
💀    🚀 makeNginxApp         📗 07:51:05.937 Start command prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing prepare command
💀    🚀 makeNginxApp         📗 07:51:05.937 Prepare command prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing test command
💀    🚀 makeNginxApp         📗 07:51:05.937 Test command prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing migrate command
💀    🚀 makeNginxApp         📗 07:51:05.937 Migrate command prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing check command
💀    🚀 makeNginxApp         📗 07:51:05.937 Check command prepared
💀    🚀 makeNginxApp         📗 07:51:05.937 Preparing replacement map
💀    🚀 makeNginxApp         📗 07:51:06.168 Add config to replacement map
💀    🚀 makeNginxApp         📗 07:51:06.176 Add env to replacement map
💀    🚀 makeNginxApp         📗 07:51:06.183 Replacement map prepared
💀    🚀 makeNginxApp         📗 07:51:06.183 ✅ Validate
💀    🚀 makeNginxApp         📗 07:51:06.183 Validate app directory
💀    🚀 makeNginxApp         📗 07:51:06.183 Done validating app directory
💀    🚀 makeNginxApp         📗 07:51:06.183 Validate app container volumes
💀    🚀 makeNginxApp         📗 07:51:06.187 Done validating app container volumes
💀    🚀 makeNginxApp         📗 07:51:06.187 Validate template locations
💀    🚀 makeNginxApp         📗 07:51:06.197 Done validating template locations
💀    🚀 makeNginxApp         📗 07:51:06.197 Validate app ports
💀    🚀 makeNginxApp         📗 07:51:06.201 Done validating app ports
💀    🚀 makeNginxApp         📗 07:51:06.201 Validate app crud fields
💀    🚀 makeNginxApp         📗 07:51:06.204 Done validating app crud fields
💀    🚀 makeNginxApp         📗 07:51:06.204 🚧 Generate
💀    🚀 makeNginxApp         📗 07:51:06.204 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 07:51:06.204   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 07:51:06.204 ]
💀    🚀 makeNginxApp         📗 07:51:06.204 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 07:51:06.232 🔩 Integrate
💀    🚀 makeNginxApp         📗 07:51:06.232 🎉🎉🎉
💀    🚀 makeNginxApp         📗 07:51:06.232 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 07:51:06.73  🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 07:51:06.73  Preparing base variables
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing start command
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Start command prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing test command
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Test command prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing check command
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Check command prepared
💀    🚀 makeNginxAppRunner   📗 07:51:06.844 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 07:51:07.096 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 07:51:07.105 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 07:51:07.113 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 07:51:07.113 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 07:51:07.113 Validate app directory
💀    🚀 makeNginxAppRunner   📗 07:51:07.113 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 07:51:07.113 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 07:51:07.116 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 07:51:07.116 Validate template locations
💀    🚀 makeNginxAppRunner   📗 07:51:07.131 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 07:51:07.131 Validate app ports
💀    🚀 makeNginxAppRunner   📗 07:51:07.136 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 07:51:07.136 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 07:51:07.139 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 07:51:07.139 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 07:51:07.139 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 07:51:07.139   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 07:51:07.139   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 07:51:07.139 ]
💀    🚀 makeNginxAppRunner   📗 07:51:07.139 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 07:51:07.167 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 07:51:07.171 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:51:07.175 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:51:07.356 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:51:07.536 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:51:07.708 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:51:07.887 Checking start
💀    🚀 makeNginxAppRunner   📗 07:51:07.891 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 07:51:08.081 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:51:08.277 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 07:51:08.28  Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 07:51:08.483 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:51:08.678 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:51:08.865 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:51:09.054 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 07:51:09.057 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 07:51:09.241 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:51:09.429 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 07:51:09.433 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 07:51:09.609 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:51:09.787 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 07:51:09.79  Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 07:51:09.966 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:51:10.135 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 07:51:10.138 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 07:51:10.309 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:51:10.485 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:51:10.489 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 07:51:10.489 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.189109435s
         Current Time: 07:51:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.299713641s
         Current Time: 07:51:10
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.037µs
         Current Time: 07:51:10
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:51:10.872 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 07:51:10.978 Synchronize task environments
💀    🚀 syncEnv              🔄 07:51:11.142 Synchronize project's environment files
💀    🚀 syncEnv              🔄 07:51:11.313 🎉🎉🎉
💀    🚀 syncEnv              🔄 07:51:11.313 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 547.311486ms
         Current Time: 07:51:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 747.952636ms
         Current Time: 07:51:11
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.064µs
         Current Time: 07:51:11
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:51:11.986 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:51:11.986 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoBackendI... 🏭 07:51:12.249 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 07:51:12.249 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:12.25  Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:51:18.129 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:51:18.129 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 07:51:18.188 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:51:18.189  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:51:18.189 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 07:51:18.19  Sending build context to Docker daemon   1.03MB
💀    🚀 buildDemoFrontend... 🏭 07:51:18.193 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoDbImage     🏭 07:51:18.194 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:51:18.194 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:51:18.195  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.195  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:51:18.195 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoDbImage     🏭 07:51:18.196 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:51:18.196 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 07:51:18.196 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 07:51:18.196  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 07:51:18.196 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 07:51:18.197 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 07:51:18.197 Step 7/11 : USER 0
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198  ---> 16e3e46a7774
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.198 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198  ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 07:51:18.198 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 07:51:18.199 Successfully built 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 07:51:18.206 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:18.208 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:51:18.208 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> 97fdfef7cb48
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> bf9c545afbe0
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> a62a483a9091
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:51:18.209  ---> db465fe79375
💀    🚀 buildDemoBackendI... 🏭 07:51:18.211 Successfully built db465fe79375
💀    🚀 buildDemoBackendI... 🏭 07:51:18.216 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 07:51:18.218 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 07:51:18.218 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 07:51:18.323 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 6.444274138s
         Current Time: 07:51:18
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.555166453s
         Current Time: 07:51:18
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.345µs
         Current Time: 07:51:18
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:51:18.726 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:51:18.726 Links updated
💀    🚀 prepareDemoBackend   🔧 07:51:18.727 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 07:51:18.756 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:51:18.842 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 07:51:18.988 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:19.584 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 07:51:19.585 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:51:19.631 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:51:19.632 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoDbImage     🏭 07:51:19.633 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:51:19.633  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:51:19.633 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:51:19.637 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639 Step 7/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 07:51:19.639 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> Using cache
💀    🚀 buildDemoDbImage     🏭 07:51:19.639 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 07:51:19.639 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64   ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64  Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64   ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 07:51:19.64  Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:51:19.641  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:51:19.641  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 07:51:19.641 Successfully built 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 07:51:19.645 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:51:19.646 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:51:19.647 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:51:19.998 🔎 Waiting docker container 'demoFrontend' running status
💀 🔥 🔎 startDemoFrontend... 📗 07:51:20.025 Error: No such container: demoFrontend
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 07:51:20.038 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🚀 startDemoFrontend... 📗 07:51:20.042 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:20.069 Error: No such container: demoDb
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:20.086 Error: No such container: demoDb
💀 🔥 🚀 startDemoFrontend... 📗 07:51:20.087 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 07:51:20.095 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:20.11  Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 07:51:20.112 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoDbContainer 🐬 07:51:20.186 f9130c31edce744dad3462fc66a8e10221598e9aa96d1634705c8e0a3fd128db
💀    🚀 startDemoFrontend... 📗 07:51:20.187 b616fa0fd9c82be197a00413377fc54483a0307304a58b704f16c2b81eca59d2
💀    🚀 prepareDemoBackend   🔧 07:51:20.871 Activate venv
💀    🚀 prepareDemoBackend   🔧 07:51:20.871 Install dependencies
💀    🚀 prepareDemoBackend   🔧 07:51:21.224 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 07:51:21.5     Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:21.507 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 07:51:21.613   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:21.623 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 07:51:21.759   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoDbContainer 🐬 07:51:23.58  🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 07:51:23.585 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.628 [38;5;6mmysql [38;5;5m00:51:23.62 
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.628 [38;5;6mmysql [38;5;5m00:51:23.62 Welcome to the Bitnami mysql container
💀    🔎 startDemoDbContainer 🐬 07:51:23.629 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 07:51:23.629 🔎 Waiting for host port: '3306'
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.63  [38;5;6mmysql [38;5;5m00:51:23.62 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀    🔎 startDemoDbContainer 🐬 07:51:23.631 🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.631 [38;5;6mmysql [38;5;5m00:51:23.63 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.633 [38;5;6mmysql [38;5;5m00:51:23.63 
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.634 [38;5;6mmysql [38;5;5m00:51:23.63 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.654 [38;5;6mmysql [38;5;5m00:51:23.65 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.663 [38;5;6mmysql [38;5;5m00:51:23.66 [38;5;2mINFO  ==> Initializing mysql database
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.675 [38;5;6mmysql [38;5;5m00:51:23.67 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.682 [38;5;6mmysql [38;5;5m00:51:23.68 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.698 [38;5;6mmysql [38;5;5m00:51:23.69 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.704 [38;5;6mmysql [38;5;5m00:51:23.70 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:23.71  [38;5;6mmysql [38;5;5m00:51:23.70 [38;5;2mINFO  ==> Installing database
💀    🚀 startDemoFrontend... 📗 07:51:24.122 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 07:51:24.128 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoFrontend... 📗 07:51:24.168 
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.168 [38;5;6mnginx [38;5;5m00:51:24.14 
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.168 [38;5;6mnginx [38;5;5m00:51:24.14 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.168 [38;5;6mnginx [38;5;5m00:51:24.15 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.168 [38;5;6mnginx [38;5;5m00:51:24.15 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.168 [38;5;6mnginx [38;5;5m00:51:24.15 
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.176 [38;5;6mnginx [38;5;5m00:51:24.17 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🔎 startDemoFrontend... 📗 07:51:24.185 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 07:51:24.185 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 07:51:24.186 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 07:51:24.186 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 07:51:24.188 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.212 2022/05/14 00:51:24 [warn] 12#12: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 07:51:24.212 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 prepareDemoBackend   🔧 07:51:24.53  Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 07:51:24.756   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:24.767 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 07:51:24.886   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:24.893 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 07:51:24.974   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:24.993 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 07:51:25.102   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:25.113 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 07:51:25.273   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:25.374 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 07:51:25.831   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:25.943 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 07:51:26.13    Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:26.202 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 07:51:26.606   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🔎 startDemoDbContainer 🐬 07:51:26.635 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:26.804 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:26.806 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 07:51:26.88  Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 07:51:27.128   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.142 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🔎 startDemoFrontend... 📗 07:51:27.191 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 07:51:27.278   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.289 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🔎 startDemoFrontend... 📗 07:51:27.317 check demoFrontend
💀    🔎 startDemoFrontend... 📗 07:51:27.322 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 07:51:27.368   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.376 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 07:51:27.465   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.476 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 07:51:27.557   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.586 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 07:51:27.694   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.708 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 07:51:27.904   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:27.92  Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 07:51:28.014   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:28.024 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 07:51:28.18    Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:28.22  Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 07:51:28.368   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:28.379 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 07:51:28.454   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:28.468 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 07:51:28.561   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 07:51:28.754 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:28.821 [38;5;6mmysql [38;5;5m00:51:28.82 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 07:51:28.939   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:28.964 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀    🚀 prepareDemoBackend   🔧 07:51:29.55    Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:29.649 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 07:51:29.746   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:29.762 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 07:51:29.853   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:29.861 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:29.934 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:29.936 ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: YES)
💀    🚀 prepareDemoBackend   🔧 07:51:29.973   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:29.987 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 07:51:30.121   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:30.139 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 07:51:30.203   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 07:51:30.366 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧 07:51:30.649   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:30.67  Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:30.837 [38;5;6mmysql [38;5;5m00:51:30.83 [38;5;2mINFO  ==> Configuring authentication
💀    🚀 prepareDemoBackend   🔧 07:51:30.84    Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:30.846 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:30.876 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:30.897 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:30.94  [38;5;6mmysql [38;5;5m00:51:30.93 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:30.943 [38;5;6mmysql [38;5;5m00:51:30.94 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 07:51:31.184   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.2   Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 07:51:31.285   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.291 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🔎 startDemoFrontend... 📗 07:51:31.324 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 07:51:31.324 📜 Task 'startDemoFrontendContainer' is ready
💀    🚀 prepareDemoBackend   🔧 07:51:31.395   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.403 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀    🚀 prepareDemoBackend   🔧 07:51:31.521   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.533 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 07:51:31.638   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.65  Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 07:51:31.788   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.832 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 07:51:31.959   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:31.967 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 07:51:32.662   Using cached https://files.pythonhosted.org/packages/c1/38/a9fd8c7bb151325d8b3d9108ce791348c84171b5d9f346b0bf0639de603f/coverage-6.3.3-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:51:32.676 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 07:51:32.75    Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:32.758 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 07:51:32.893   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:32.902 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:32.962 [38;5;6mmysql [38;5;5m00:51:32.96 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:33.042 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:33.044 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 07:51:33.068   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:33.083 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 07:51:33.181   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:33.194 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 07:51:33.321   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:33.331 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 07:51:33.457   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:51:33.764 Installing collected packages: aiofiles, asgiref, avro-python3, six, pycparser, cffi, bcrypt, certifi, charset-normalizer, click, urllib3, idna, requests, fastavro, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, iniconfig, pyparsing, packaging, py, toml, attrs, pluggy, pytest, tomli, coverage, pytest-cov, pyasn1, rsa, ecdsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 07:51:33.793   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 07:51:34.054     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 07:51:34.551   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:36.196 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:36.198 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:38.973 [38;5;6mmysql [38;5;5m00:51:38.97 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:38.991 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:39     find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:39.005 [38;5;6mmysql [38;5;5m00:51:39.00 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:39.336 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:39.338 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 startDemoDbContainer 🐬 07:51:42.022 
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:42.022 [38;5;6mmysql [38;5;5m00:51:42.02 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 07:51:42.076 [38;5;6mmysql [38;5;5m00:51:42.07 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 07:51:42.314 2022-05-14T00:51:42.310272Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 07:51:42.315 2022-05-14T00:51:42.311905Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 07:51:42.315 2022-05-14T00:51:42.311912Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 07:51:42.316 2022-05-14T00:51:42.315764Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:42.483 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:42.484 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 startDemoDbContainer 🐬 07:51:42.515 2022-05-14T00:51:42.515271Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 07:51:42.722 2022-05-14T00:51:42.721793Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 07:51:42.722 2022-05-14T00:51:42.721854Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 07:51:42.741 2022-05-14T00:51:42.741126Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 07:51:42.741 2022-05-14T00:51:42.741216Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 07:51:45.626 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 07:51:45.633 Database
💀    🔎 startDemoDbContainer 🐬 07:51:45.634 information_schema
💀    🔎 startDemoDbContainer 🐬 07:51:45.634 mysql
💀    🔎 startDemoDbContainer 🐬 07:51:45.634 performance_schema
💀    🔎 startDemoDbContainer 🐬 07:51:45.634 sample
💀    🔎 startDemoDbContainer 🐬 07:51:45.634 sys
💀    🔎 startDemoDbContainer 🐬 07:51:45.637 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 07:51:49.64  🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 07:51:49.64  📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 07:52:04.509     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 07:52:05.875   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 07:52:06.133     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 07:52:06.837   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 07:52:07.024     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 07:52:07.09  Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.3 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 07:52:07.163 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 07:52:07.163 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 07:52:07.225 Prepare
💀    🚀 prepareDemoBackend   🔧 07:52:07.225 prepare command
💀    🚀 prepareDemoBackend   🔧 07:52:07.225 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackend     ⚡ 07:52:07.445 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 07:52:07.446 Activate venv
💀    🚀 startDemoBackend     ⚡ 07:52:07.446 Start
💀    🚀 startDemoBackend     ⚡ 07:52:07.922 2022-05-14 07:52:07,922 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 07:52:07.922 2022-05-14 07:52:07,922 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.927 2022-05-14 07:52:07,927 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 07:52:07.927 2022-05-14 07:52:07,927 INFO sqlalchemy.engine.Engine [generated in 0.00020s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.93  2022-05-14 07:52:07,930 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 07:52:07.93  2022-05-14 07:52:07,930 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.932 2022-05-14 07:52:07,932 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:07.933 2022-05-14 07:52:07,933 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 07:52:07.933 2022-05-14 07:52:07,933 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 07:52:07.936 2022-05-14 07:52:07,936 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 07:52:07.936 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 07:52:07.936 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:07.936 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 )
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 
💀    🚀 startDemoBackend     ⚡ 07:52:07.937 2022-05-14 07:52:07,936 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.96  2022-05-14 07:52:07,960 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 07:52:07.96  2022-05-14 07:52:07,960 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.979 2022-05-14 07:52:07,979 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 07:52:07.979 2022-05-14 07:52:07,979 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:07.996 2022-05-14 07:52:07,996 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 07:52:07.996 2022-05-14 07:52:07,996 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.014 2022-05-14 07:52:08,014 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 07:52:08.014 2022-05-14 07:52:08,014 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.033 2022-05-14 07:52:08,033 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 07:52:08.034 2022-05-14 07:52:08,034 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:08.034 2022-05-14 07:52:08,034 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 07:52:08.035 2022-05-14 07:52:08,034 INFO sqlalchemy.engine.Engine [cached since 0.102s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 07:52:08.036 2022-05-14 07:52:08,036 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 07:52:08.036 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 07:52:08.036 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.036 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 07:52:08.036 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 )
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 
💀    🚀 startDemoBackend     ⚡ 07:52:08.037 2022-05-14 07:52:08,036 INFO sqlalchemy.engine.Engine [no key 0.00011s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.061 2022-05-14 07:52:08,061 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 07:52:08.061 2022-05-14 07:52:08,061 INFO sqlalchemy.engine.Engine [no key 0.00031s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.08  2022-05-14 07:52:08,080 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 07:52:08.08  2022-05-14 07:52:08,080 INFO sqlalchemy.engine.Engine [no key 0.00020s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.097 2022-05-14 07:52:08,097 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 07:52:08.097 2022-05-14 07:52:08,097 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.118 2022-05-14 07:52:08,118 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 07:52:08.119 2022-05-14 07:52:08,119 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:08.12  2022-05-14 07:52:08,120 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 07:52:08.12  2022-05-14 07:52:08,120 INFO sqlalchemy.engine.Engine [cached since 0.1876s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 2022-05-14 07:52:08,123 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 )
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 
💀    🚀 startDemoBackend     ⚡ 07:52:08.123 2022-05-14 07:52:08,123 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.153 2022-05-14 07:52:08,153 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 07:52:08.153 2022-05-14 07:52:08,153 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.171 2022-05-14 07:52:08,171 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 07:52:08.171 2022-05-14 07:52:08,171 INFO sqlalchemy.engine.Engine [no key 0.00024s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.192 2022-05-14 07:52:08,191 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 07:52:08.192 2022-05-14 07:52:08,191 INFO sqlalchemy.engine.Engine [no key 0.00025s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.217 2022-05-14 07:52:08,216 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 07:52:08.217 2022-05-14 07:52:08,217 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.234 2022-05-14 07:52:08,234 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 07:52:08.234 2022-05-14 07:52:08,234 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.257 2022-05-14 07:52:08,257 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 07:52:08.257 2022-05-14 07:52:08,257 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 07:52:08.277 2022-05-14 07:52:08,277 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 07:52:08.28  2022-05-14 07:52:08,280 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:08.282 2022-05-14 07:52:08,282 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 07:52:08.282 FROM users 
💀    🚀 startDemoBackend     ⚡ 07:52:08.282 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 07:52:08.282  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 07:52:08.282 2022-05-14 07:52:08,282 INFO sqlalchemy.engine.Engine [generated in 0.00025s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 07:52:08.284 2022-05-14 07:52:08,284 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 07:52:08.535 2022-05-14 07:52:08,535 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:08.537 2022-05-14 07:52:08,537 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 07:52:08.537 2022-05-14 07:52:08,537 INFO sqlalchemy.engine.Engine [generated in 0.00020s] {'id': '928864cf-5416-4df0-a1ea-93c81f371a78', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$MnY.971H/BhOIqjy6NeS9uwUlDfFPcROno5TaAmizGLk2osn6tpg2', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 14, 7, 52, 8, 535235), 'updated_at': datetime.datetime(2022, 5, 14, 7, 52, 8, 537037)}
💀    🚀 startDemoBackend     ⚡ 07:52:08.538 2022-05-14 07:52:08,538 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 07:52:08.546 2022-05-14 07:52:08,546 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 07:52:08.547 2022-05-14 07:52:08,547 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 07:52:08.547 FROM users 
💀    🚀 startDemoBackend     ⚡ 07:52:08.547 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 07:52:08.547 2022-05-14 07:52:08,547 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {'pk_1': '928864cf-5416-4df0-a1ea-93c81f371a78'}
💀    🚀 startDemoBackend     ⚡ 07:52:08.548 2022-05-14 07:52:08,548 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 07:52:08.55  Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.56  Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 07:52:08.574 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.584 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 07:52:08.584 Register library route handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.585 Register library event handler
💀    🚀 startDemoBackend     ⚡ 07:52:08.585 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 07:52:08.585 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:08.586 INFO:     Started server process [7662]
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:08.586 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:08.586 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:08.586 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 07:52:09.451 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 07:52:09.451 check demoBackend
💀    🔎 startDemoBackend     ⚡ 07:52:09.451 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 07:52:09.451 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 07:52:09.558 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 50.940227474s
         Current Time: 07:52:09
         Active Process:
           * (PID=28332) 🐬 'startDemoDbContainer' service
           * (PID=28301) 📗 'startDemoFrontendContainer' service
           * (PID=7657) ⚡ 'startDemoBackend' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=7657)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=28332)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=28301)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:11.196 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:11.296 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:11.296 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 07:52:11.296 INFO:     Finished server process [7662]
💀    🚀 startDemoBackend     ⚡ 07:52:11.412 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 07:52:11.412 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 53.043782995s
         Current Time: 07:52:11
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.233µs
         Current Time: 07:52:11
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:52:11.937 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:52:11.937 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 07:52:11.969 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:52:12.049 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoBackendI... 🏭 07:52:12.197 Build image demo-backend:latest
💀    🚀 buildDemoFrontend... 🏭 07:52:12.197 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:52:12.779 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:52:12.78  Sending build context to Docker daemon  22.02kB
💀    🚀 buildDemoDbImage     🏭 07:52:12.842 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:52:12.842  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:52:12.843 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 07:52:12.843 Sending build context to Docker daemon   1.18MB
💀    🚀 buildDemoDbImage     🏭 07:52:12.852 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.852  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 11c677f847bc
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 776095918b33
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 48dc42a93a8a
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 0beee76410dd
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 68555ae22bc5
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 992fa94aa2f2
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853  ---> 02304e445f6f
💀    🚀 buildDemoFrontend... 🏭 07:52:12.853 Successfully built 02304e445f6f
💀    🚀 buildDemoDbImage     🏭 07:52:12.854 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:52:12.854 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 07:52:12.855 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 07:52:12.857  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 07:52:12.857 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoFrontend... 🏭 07:52:12.857 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 07:52:12.861 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 07:52:12.862  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:52:12.862  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 07:52:12.862 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 07:52:12.863 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:52:12.863 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 07:52:13.008  ---> 821efbdd6c49
💀    🚀 buildDemoBackendI... 🏭 07:52:13.008 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 07:52:13.037  ---> Running in 51b34a04fca3
💀    🚀 buildDemoBackendI... 🏭 07:52:13.091 Removing intermediate container 51b34a04fca3
💀    🚀 buildDemoBackendI... 🏭 07:52:13.091  ---> d874165bc0b8
💀    🚀 buildDemoBackendI... 🏭 07:52:13.091 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:52:13.112  ---> Running in 127b0d15ce27
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:52:13.194 🔎 Waiting docker container 'demoFrontend' running status
💀    🔎 startDemoFrontend... 📗 07:52:13.223 🔎 Waiting docker container 'demoFrontend' healthcheck
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 startDemoFrontend... 📗 07:52:13.226 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 07:52:13.226 🐳 Logging 'demoFrontend'
💀    🔎 startDemoDbContainer 🐬 07:52:13.226 🔎 Waiting docker container 'demoDb' running status
💀    🔎 startDemoFrontend... 📗 07:52:13.264 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 07:52:13.264 🔎 Waiting for host port: '8080'
💀    🔎 startDemoDbContainer 🐬 07:52:13.265 🔎 Waiting docker container 'demoDb' healthcheck
💀    🔎 startDemoFrontend... 📗 07:52:13.266 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 07:52:13.266 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 07:52:13.267 🔎 Host port '443' is ready
💀    🚀 startDemoDbContainer 🐬 07:52:13.284 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 07:52:13.284 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 07:52:13.311 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 07:52:13.311 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 07:52:13.312 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 07:52:13.95  Removing intermediate container 127b0d15ce27
💀    🚀 buildDemoBackendI... 🏭 07:52:13.95   ---> 6408d7923bc6
💀    🚀 buildDemoBackendI... 🏭 07:52:13.95  Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:52:13.981  ---> Running in 0a6797c25320
💀    🚀 buildDemoBackendI... 🏭 07:52:14.06  Removing intermediate container 0a6797c25320
💀    🚀 buildDemoBackendI... 🏭 07:52:14.06   ---> 711ddb58dfa1
💀    🚀 buildDemoBackendI... 🏭 07:52:14.063 Successfully built 711ddb58dfa1
💀    🚀 buildDemoBackendI... 🏭 07:52:14.07  Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 07:52:14.071 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 07:52:14.071 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 07:52:16.272 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 07:52:16.317 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 07:52:16.412 check demoFrontend
💀    🔎 startDemoFrontend... 📗 07:52:16.418 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 07:52:16.458 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 Database
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 information_schema
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 mysql
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 performance_schema
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 sample
💀    🔎 startDemoDbContainer 🐬 07:52:16.461 sys
💀    🔎 startDemoDbContainer 🐬 07:52:16.465 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 07:52:20.42  🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 07:52:20.42  📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 07:52:20.467 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 07:52:20.467 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 07:52:20.969 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:20.996 Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ 07:52:21.002 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:21.022 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 07:52:21.025 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 07:52:21.076 a92833ecac8960032c11703e81353b7fe797f7196b0148f5bdce9806fa4650c3
💀    🚀 startDemoBackendC... ⚡ 07:52:22.332 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 07:52:22.335 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 07:52:22.374 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 07:52:22.374 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 07:52:22.375 🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 07:52:23.232 2022-05-14 00:52:23,231 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 07:52:23.232 2022-05-14 00:52:23,231 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.238 2022-05-14 00:52:23,238 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 07:52:23.238 2022-05-14 00:52:23,238 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.243 2022-05-14 00:52:23,242 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 07:52:23.243 2022-05-14 00:52:23,242 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.247 2022-05-14 00:52:23,246 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 07:52:23.248 2022-05-14 00:52:23,247 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 07:52:23.248 2022-05-14 00:52:23,247 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.251 2022-05-14 00:52:23,250 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 07:52:23.255 2022-05-14 00:52:23,254 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 07:52:23.255 2022-05-14 00:52:23,255 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 07:52:23.255 2022-05-14 00:52:23,255 INFO sqlalchemy.engine.Engine [cached since 0.007713s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.258 2022-05-14 00:52:23,258 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 07:52:23.261 2022-05-14 00:52:23,260 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 07:52:23.261 2022-05-14 00:52:23,261 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 07:52:23.261 2022-05-14 00:52:23,261 INFO sqlalchemy.engine.Engine [cached since 0.01362s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.264 2022-05-14 00:52:23,263 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 07:52:23.269 2022-05-14 00:52:23,268 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 07:52:23.271 2022-05-14 00:52:23,270 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 07:52:23.271 FROM users 
💀    🚀 startDemoBackendC... ⚡ 07:52:23.271 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 07:52:23.271  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 07:52:23.271 2022-05-14 00:52:23,270 INFO sqlalchemy.engine.Engine [generated in 0.00020s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 07:52:23.274 2022-05-14 00:52:23,273 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 07:52:23.277 Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.288 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 07:52:23.299 Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.307 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 07:52:23.307 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.307 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 07:52:23.307 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 07:52:23.307 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:23.308 INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:23.308 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:23.308 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 07:52:23.308 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 07:52:25.378 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 07:52:25.545 check demoBackend
💀    🔎 startDemoBackendC... ⚡ 07:52:25.549 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 07:52:26.55  🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 07:52:26.55  📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 07:52:26.657 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 14.827934406s
         Current Time: 07:52:26
         Active Process:
           * (PID=10180) 📗 'startDemoFrontendContainer' service
           * (PID=12731) ⚡ 'startDemoBackendContainer' service
           * (PID=10216) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=10180)
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=12731)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=10216)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 16.931079765s
         Current Time: 07:52:28
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.775µs
         Current Time: 07:52:29
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:52:29.052 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:52:29.052 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoDbContainer  ✋ 07:52:29.425 Stop docker container demoDb
💀    🚀 stopDemoBackendCo... ✋ 07:52:29.427 Stop docker container demoBackend
💀    🚀 stopDemoFrontendC... ✋ 07:52:29.427 Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ 07:52:32.904 demoDb
💀    🚀 stopDemoDbContainer  ✋ 07:52:32.905 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 07:52:32.905 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoFrontendC... ✋ 07:52:40.242 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 07:52:40.244 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 07:52:40.244 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀    🚀 stopDemoBackendCo... ✋ 07:52:40.434 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 07:52:40.435 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 07:52:40.435 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 07:52:40.542 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.597288408s
         Current Time: 07:52:40
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.707613226s
         Current Time: 07:52:40
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.247µs
         Current Time: 07:52:40
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:52:40.953 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:52:40.953 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🔥 🚀 removeDemoFronten... ❌ 07:52:41.282 Error: No such container: 
💀    🚀 removeDemoFronten... ❌ 07:52:41.285 Stop docker container demoFrontend
💀 🔥 🚀 removeDemoDbConta... ❌ 07:52:41.287 Error: No such container: 
💀 🔥 🚀 removeDemoBackend... ❌ 07:52:41.287 Error: No such container: 
💀    🚀 removeDemoDbConta... ❌ 07:52:41.29  Stop docker container demoDb
💀    🚀 removeDemoBackend... ❌ 07:52:41.293 Stop docker container demoBackend
💀    🚀 removeDemoDbConta... ❌ 07:52:41.388 Docker container demoDb stopped
💀    🚀 removeDemoDbConta... ❌ 07:52:41.388 Remove docker container demoDb
💀    🚀 removeDemoFronten... ❌ 07:52:41.392 Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ 07:52:41.392 Remove docker container demoFrontend
💀    🚀 removeDemoBackend... ❌ 07:52:41.394 Docker container demoBackend stopped
💀    🚀 removeDemoBackend... ❌ 07:52:41.394 Remove docker container demoBackend
💀    🚀 removeDemoFronten... ❌ 07:52:41.457 demoFrontend
💀    🚀 removeDemoBackend... ❌ 07:52:41.458 demoBackend
💀    🚀 removeDemoBackend... ❌ 07:52:41.465 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ 07:52:41.465 Docker container demoBackend removed
💀    🚀 removeDemoFronten... ❌ 07:52:41.466 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ 07:52:41.466 Docker container demoFrontend removed
💀    🚀 removeDemoDbConta... ❌ 07:52:41.471 demoDb
💀    🚀 removeDemoDbConta... ❌ 07:52:41.473 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ 07:52:41.473 Docker container demoDb removed
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 removeContainers     ❌ 07:52:41.58  
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 733.242446ms
         Current Time: 07:52:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 845.129186ms
         Current Time: 07:52:41
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.463µs
         Current Time: 07:52:41
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:52:41.989 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:52:41.992         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:52:41.992     
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:52:41.992   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:52:41.992   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:52:41.992   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:52:41.992 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.446 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.446 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.626 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:42.627 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.024 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.051 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.065 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.065 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.065 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.065 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.066 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.071 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.071 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.09  Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.09  Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.118 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.118 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.126 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.127 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.127 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.127   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.127 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.127 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.198 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.198 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.198 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.658 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.658 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.861 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.861 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.861 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:43.862 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.127 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.135 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.141 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.141 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.142 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.142 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.142 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.146 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.146 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.155 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.155 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.159 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.159 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.163 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.183 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.188 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.192 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.414 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.588 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.592 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.773 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.954 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:44.959 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:45.135 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:45.135 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:45.135 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.253942081s
         Current Time: 07:52:45
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.365785373s
         Current Time: 07:52:45
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.031µs
         Current Time: 07:52:45
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:52:45.627 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:52:45.629         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:52:45.629     
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:52:45.629   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:52:45.629   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:52:45.629   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:52:45.629 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:46.106 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:46.106 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.248 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.249 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.544 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.553 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.56  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.561 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.561 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.561 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.561 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.565 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.565 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.577 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.577 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.581 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.581 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.585 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.624 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.625 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:47.625 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:48.015 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:48.015 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.102 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.102 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.102 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.102 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.102 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.103 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.357 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.365 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.373 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.373 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.373 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.374 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.374 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.378 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.378 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.387 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.387 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.391 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.391 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.394 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.421 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.428 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.433 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.649 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.844 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:49.849 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.043 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.237 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.24  Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.432 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.432 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:50.432 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.912294652s
         Current Time: 07:52:50
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.023321629s
         Current Time: 07:52:50
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.472µs
         Current Time: 07:52:50
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:52:50.839 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:52:50.845         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:52:50.845     
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:52:50.845   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:52:50.845   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:52:50.845   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:52:50.845 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.28  🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.28  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.408 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.764 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.771 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.78  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.78  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.78  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.78  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.78  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.785 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.785 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.799 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.799 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.803 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.804 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.808 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.863 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.863 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:51.863 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:52:53.639 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:52:53.639 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:54.571 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.485 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.498 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.511 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.511 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.511 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.511 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.511 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.524 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.524 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.546 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.546 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.551 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.551 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.556 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.596 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.603 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.608 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:55.917 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:56.301 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:56.336 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:56.918 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:52:57.362 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:57.372 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:52:57.786 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:52:57.786 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:52:57.786 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 7.053976014s
         Current Time: 07:52:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 7.164725411s
         Current Time: 07:52:57
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.43µs
         Current Time: 07:52:58
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:52:58.351 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 07:52:58.461 Synchronize task environments
💀    🚀 syncEnv              🔄 07:52:58.778 Synchronize project's environment files
💀    🚀 syncEnv              🔄 07:52:59.027 🎉🎉🎉
💀    🚀 syncEnv              🔄 07:52:59.027 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 787.15434ms
         Current Time: 07:52:59
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 987.367237ms
         Current Time: 07:52:59
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.031µs
         Current Time: 07:52:59
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:52:59.518 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:52:59.632 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:52:59.632 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 221.35016ms
         Current Time: 07:52:59
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 421.839336ms
         Current Time: 07:52:59
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.664µs
         Current Time: 07:53:00
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:53:00.141 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:53:00.257 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:53:00.257 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 222.387002ms
         Current Time: 07:53:00
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 429.475112ms
         Current Time: 07:53:00
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 3.03µs
         Current Time: 07:53:00
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 07:53:00.995 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:00.995 🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 07:53:00.997 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 07:53:03.496 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:03.506 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 07:53:03.566 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:53:03.857 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:03.858 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:03.901 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:04.753   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:04.767   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:04.775 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:04.788 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:04.933   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:04.951 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:04.993   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 07:53:05.08    Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 07:53:05.225   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 07:53:05.288 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:05.367   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:05.371 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:05.372 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:05.511 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:05.611   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:05.626 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:05.722   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:05.75  Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:05.754   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:05.773 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:05.857   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:05.862 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.22    Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.318 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:06.574   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.599   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.605 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:06.65  Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:06.701   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.765   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:06.771 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.775 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:06.835   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:06.841 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.949   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:06.959   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:06.969 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:06.98  Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:07.114   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:07.118 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:53:07.147   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:07.171 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:53:07.186   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:07.194 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:07.261   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:07.266 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.316   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.338 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:07.525   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:07.537 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:07.624   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:07.636 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:07.642   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:07.656 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.724   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.736 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:07.737   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:07.754 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.833   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.85  Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.937   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:07.97  Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:07.994   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:08.002   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.007 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.017 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.07    Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.079   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.085 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.096   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.112 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.125 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.158   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.165 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:08.211   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:08.212   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.218 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.218 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.244   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.25  Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:08.341   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:08.344   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.35  Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.354 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:53:08.425   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.43    Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:08.432 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:53:08.437 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.44    Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.446 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.545   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:53:08.545   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:53:08.555   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:53:08.634 Installing collected packages: six, grpcio, semver, dill, pyyaml, protobuf, pulumi, attrs, arpeggio, parver, charset-normalizer, idna, certifi, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 07:53:08.651 Installing collected packages: protobuf, six, grpcio, semver, pyyaml, dill, pulumi, arpeggio, attrs, parver, idna, certifi, charset-normalizer, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 07:53:08.659 Installing collected packages: dill, pyyaml, six, grpcio, protobuf, semver, pulumi, arpeggio, attrs, parver, certifi, charset-normalizer, idna, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 07:53:09.326   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 07:53:09.344   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 07:53:09.348   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 07:53:10.969     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 07:53:10.971     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 07:53:10.978     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 07:53:11.015 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 07:53:11.016 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.02  Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 07:53:11.051 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:53:11.051 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:53:11.053 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:53:11.053 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:53:11.055 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:53:11.055 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.247 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoFronte... 🏁 07:53:11.248 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.316 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.316 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.415 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.415 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.415 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.416       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.416 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:53:11.417 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:11.418 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.418 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 Flags:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 
💀    🚀 prepareDemoFronte... 🏁 07:53:11.419 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:53:11.42  🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 07:53:11.705 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 07:53:11.722 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 07:53:11.837 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 07:53:11.847 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 07:53:12.182 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:53:12.274 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"ClusterIP"}
💀    🚀 prepareDemoBacken... 🏁 07:53:12.274 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:53:12.332       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 
💀    🚀 prepareDemoBacken... 🏁 07:53:12.333 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 07:53:12.334 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 07:53:12.576 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 07:53:12.701 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 07:53:13.697 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:53:13.745 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:53:14.156 
💀    🚀 deployDemoFronten... 🏁 07:53:14.199 
💀    🚀 deployDemoBackend... 🏁 07:53:14.529 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:53:14.592  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 07:53:14.733  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:14.733  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:14.737  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoBackend... 🏁 07:53:14.982 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.057  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.082  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 07:53:15.106  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 07:53:15.113  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227  
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.227 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 07:53:15.269  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 07:53:15.269  
💀    🚀 deployDemoFronten... 🏁 07:53:15.27  Resources:
💀    🚀 deployDemoFronten... 🏁 07:53:15.27      + 4 to create
💀    🚀 deployDemoFronten... 🏁 07:53:15.27  
💀    🚀 deployDemoFronten... 🏁 07:53:15.27  Updating (dev):
💀    🚀 deployDemoBackend... 🏁 07:53:15.47   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:53:15.538  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:15.668 
💀    🚀 deployDemoFronten... 🏁 07:53:15.707 
💀    🚀 deployDemoBackend... 🏁 07:53:15.839  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:53:15.841  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:53:15.845  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:53:15.99   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:53:15.99   
💀    🚀 deployDemoBackend... 🏁 07:53:15.99  Resources:
💀    🚀 deployDemoBackend... 🏁 07:53:15.99      + 5 to create
💀    🚀 deployDemoBackend... 🏁 07:53:15.99  
💀    🚀 deployDemoBackend... 🏁 07:53:15.99  Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.081  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.133  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.153  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.221  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.433  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.435  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.457  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.46   +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.463  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.471  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoBackend... 🏁 07:53:16.472 
💀    🚀 deployDemoFronten... 🏁 07:53:16.531  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.534  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.545  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.549  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 07:53:16.553  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 07:53:16.559  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.644  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.644  
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645     app: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645         ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.645                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.646                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             time       : "2022-05-14T00:53:16Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     resource_version  : "185526"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     uid               : "3bc70e0f-6b7e-4418-84b3-d1307e0da783"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.647                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.648                                     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                             time       : "2022-05-14T00:53:16Z"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                     ]
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                     resource_version  : "185525"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                     uid               : "9fc8572b-64c6-4e72-aca0-b16299f8d071"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                 }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649             }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649         }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649     }
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649 Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.649 
💀    🚀 deployDemoDbDeplo... 🏁 07:53:16.65  hello world
💀    🚀 deployDemoFronten... 🏁 07:53:16.702  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 07:53:16.702  
💀    🚀 deployDemoFronten... 🏁 07:53:16.703 Outputs:
💀    🚀 deployDemoFronten... 🏁 07:53:16.703     app: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.703         ready    : [
💀    🚀 deployDemoFronten... 🏁 07:53:16.703             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704         ]
💀    🚀 deployDemoFronten... 🏁 07:53:16.704         resources: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.704                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                             spec      : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                 selector: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                 template: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                         labels: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.705                                                 ]
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                         ]
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706 
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     labels            : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.706                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                         [0]: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.707                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                         }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                     }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                                 }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                             }
💀    🚀 deployDemoFronten... 🏁 07:53:16.708                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 07:53:16.708       
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoBackend... 🏁 07:53:16.916  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 07:53:16.985  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.255  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.256  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.262  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.268  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.27   +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.273  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 07:53:17.281  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 07:53:17.285  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 07:53:17.293  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 07:53:17.461  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 07:53:17.461  
💀    🚀 deployDemoBackend... 🏁 07:53:17.463 Outputs:
💀    🚀 deployDemoBackend... 🏁 07:53:17.464     app: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464         ready    : [
💀    🚀 deployDemoBackend... 🏁 07:53:17.464             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464         ]
💀    🚀 deployDemoBackend... 🏁 07:53:17.464         resources: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                 }
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.464                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 }
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                             }
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                             spec      : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 selector: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 }
💀    🚀 deployDemoBackend... 🏁 07:53:17.465                                 template: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                         labels: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                         }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.466                                                         value: "false"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         ]
💀    🚀 deployDemoBackend... 🏁 07:53:17.467 
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         ]
💀    🚀 deployDemoBackend... 🏁 07:53:17.467 
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         ]
💀    🚀 deployDemoBackend... 🏁 07:53:17.467 
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         value: (json) []
💀    🚀 deployDemoBackend... 🏁 07:53:17.467 
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                         value: "600"
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.467                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 07:53:17.468                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     }
💀    🚀 deployDemoBackend... 🏁 07:53:17.469                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 07:53:17.469   
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 07:53:17.574 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 16.700226401s
         Current Time: 07:53:17
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 16.810807494s
         Current Time: 07:53:17
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.156µs
         Current Time: 07:53:17
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.013 🚧 Install pip packages.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 07:53:18.013 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:53:18.016 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.381 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.394 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.395 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.399 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.403 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.404 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.407 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.411 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.412 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.414 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.414 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.416 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.417 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.417 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.419 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.419 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.42  Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.424 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.424 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.426 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.426 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.426 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.43  Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.435 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.436 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.439 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.441 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.443 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.449 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.457 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.463 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.467 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.472 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.475 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.476 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.484 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.514 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.514 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.519 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.523 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.531 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.558 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.562 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.565 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 07:53:18.569 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.573 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.577 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 07:53:18.58  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:53:18.587 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:53:18.587 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:53:18.594 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:53:18.594 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:53:18.61  WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:53:18.61  You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.864 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoFronte... 🏁 07:53:18.864 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.941 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 07:53:18.944 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 07:53:18.949 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95      dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  for this case.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Usage:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Flags:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  
💀    🚀 prepareDemoFronte... 🏁 07:53:18.95  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:53:18.952 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.008 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.009       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01  
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.01  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 07:53:19.011 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 07:53:20.092 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:53:20.183 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"ClusterIP"}
💀    🚀 prepareDemoBacken... 🏁 07:53:20.183 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.255 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 
💀    🚀 prepareDemoBacken... 🏁 07:53:20.256 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 07:53:20.258 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 07:53:20.712 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 07:53:20.815 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.818  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.818  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.82   -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.821  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.822  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.822  
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824 Outputs:
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824   - app: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824         ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824       - resources: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.824                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                               - template: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.825                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                                 ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                         ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.826                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.827                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                           - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                     ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - resource_version  : "185542"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - uid               : "f05b03df-38db-4569-883e-7f27a495179d"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.828                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.829                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                            - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                            - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                          }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                      ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                    - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                                  }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                              ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                            - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                          }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                      }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                  }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83              }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83            - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.83                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                           - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                     ]
💀    🚀 destroyDemoFronte... 🏁 07:53:20.831                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832                   - resource_version  : "185541"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832                   - uid               : "6524bd41-2058-4cef-8659-8505eb6dbc11"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832             }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832         }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832     }
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832 Resources:
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.832 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.865 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 07:53:20.939 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.941  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:53:20.941  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.976 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.979  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.98   -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.982  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.983  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.985  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.985  
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.987                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.988                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.989                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                        - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                        - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                        - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                    - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                        - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                                  }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                  }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                      ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - resource_version  : "185526"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - uid               : "3bc70e0f-6b7e-4418-84b3-d1307e0da783"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                  }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                        - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                        - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                        - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                    - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                        - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                        - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                            -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                    -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                            - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.99                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                             ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.991                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                           - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                   - resource_version  : "185525"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                   - uid               : "9fc8572b-64c6-4e72-aca0-b16299f8d071"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:20.992 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 07:53:21.068  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.07   -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.071  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.075  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.081  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.084  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.084  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.085  
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086 Outputs:
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086   - app: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086         ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086       - resources: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 07:53:21.086                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                               - template: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                                 ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                         ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087 
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.087                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                           - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                     ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                   - resource_version  : "185542"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                   - uid               : "f05b03df-38db-4569-883e-7f27a495179d"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.088               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                             }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                     ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                             ]
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                         }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                     }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.089                 }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09              }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09            - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                            - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                            - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                                  }
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 07:53:21.09   
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.124 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.126  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.126  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.229  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.233  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.236  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.238  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.242  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247  
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.247                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248 
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                   - creation_timestamp: "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.248                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.249                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                    - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                        - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                                  }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                      }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                                  }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                              }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                            - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                            - time       : "2022-05-14T00:53:16Z"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                          }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                      ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                    - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                    - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.25                    - resource_version  : "185526"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - uid               : "3bc70e0f-6b7e-4418-84b3-d1307e0da783"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                 }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                     }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                             }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.251                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                         }
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 07:53:21.252                                   - image_pull_policy         : "IfNotPre
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 07:53:21.867 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 07:53:21.943 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.944  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.946  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.952  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.955  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.957  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.959  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.959  
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961 Outputs:
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961   - app: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961       - resources: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.961               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.962                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                               - template: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.963                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.964 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.965                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.966                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.967                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.968                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.969                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                                      }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                                  ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                              }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                          ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                        - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                      }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                                  }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                              }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97  
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                      }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                    - creation_timestamp: "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                    - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                    - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                        - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                      }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                    - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.97                            - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.971                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.972                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.973                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.974                                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - time       : "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - resource_version  : "185558"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - uid               : "f4ceee80-3aa4-4e82-870d-ba08e5d91319"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.975                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: "false"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                             ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                             ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.976                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                             ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "600"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.977                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.978                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                   -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.979                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "root"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "/"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                      ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                            - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                      ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                    - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                                  }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                              ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                            - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                          }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                      }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                  }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98                - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98              }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.98            - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                 ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                   - creation_timestamp: "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.981                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - time       : "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - resource_version  : "185561"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - uid               : "919d2062-dbbf-4194-ab2e-939d8f826d86"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - cluster_ip             : "10.96.119.53"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   -     [0]: "10.96.119.53"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.982                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - status     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.983                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                   - creation_timestamp: "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.984                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.985                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                           - time       : "2022-05-14T00:53:17Z"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                     ]
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                   - resource_version  : "185557"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                   - uid               : "94c4e3b6-529b-44d0-8eca-d59b0c9f1ab9"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986             }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986         }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986     }
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986 Resources:
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986 
💀    🚀 destroyDemoBacken... 🏁 07:53:21.986 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 07:53:22.047 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.048  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.048  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.055  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.138  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.14   -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.145  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.148  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.149  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.152  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.16   -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.167  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.167  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.169  
💀    🚀 destroyDemoBacken... 🏁 07:53:22.169 Outputs:
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17    - app: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17        - ready    : [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17        -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17        -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17        -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17          ]
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17        - resources: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17            - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                    - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                            - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                            - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - annotations: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                  }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - labels     : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                  }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                              }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                            - spec      : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                - selector: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                    - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.17                                        - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                 }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                               - template: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                         }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.171                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                         ]
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.172                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173 
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.173                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.174                                                     }
💀    🚀 destroyDemoBacken... 🏁 07:53:22.174                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 07:53:22.174                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 07:53:22.174                                         
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 07:53:22.278 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 4.390992741s
         Current Time: 07:53:22
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.501724831s
         Current Time: 07:53:22
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

