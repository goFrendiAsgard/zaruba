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

# zaruba project setValue defaultImagePrefix gofrendi
# zaruba please pushImages

zaruba project setValue defaultKubeContext docker-desktop
zaruba project setValue pulumiUseLocalBackend true

zaruba please deploy
zaruba please destroy
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.98µs
         Current Time: 08:45:57
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 08:45:57.021 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 08:45:57.025 🎉🎉🎉
💀    🚀 initProject          🚧 08:45:57.025 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 114.111057ms
         Current Time: 08:45:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 315.602577ms
         Current Time: 08:45:57
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.203µs
         Current Time: 08:45:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:45:57.482 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:45:57.484 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:45:57.484 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:45:57.484 
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:45:57.484         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:45:57.485     
💀    🚀 zrbShowAdv           ☕ 08:45:57.485 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:45:57.485 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:45:57.485   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:45:57.485   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:45:57.485   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:45:57.485 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 08:45:57.926 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 08:45:57.926 Preparing base variables
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Base variables prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing start command
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Start command prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing test command
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Test command prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing check command
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Check command prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.008 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 08:45:58.237 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 08:45:58.243 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 08:45:58.249 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 08:45:58.249 ✅ Validate
💀    🚀 makeMysqlApp         🐬 08:45:58.249 Validate app directory
💀    🚀 makeMysqlApp         🐬 08:45:58.249 Done validating app directory
💀    🚀 makeMysqlApp         🐬 08:45:58.249 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 08:45:58.252 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 08:45:58.252 Validate template locations
💀    🚀 makeMysqlApp         🐬 08:45:58.261 Done validating template locations
💀    🚀 makeMysqlApp         🐬 08:45:58.261 Validate app ports
💀    🚀 makeMysqlApp         🐬 08:45:58.264 Done validating app ports
💀    🚀 makeMysqlApp         🐬 08:45:58.264 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 08:45:58.267 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 08:45:58.267 🚧 Generate
💀    🚀 makeMysqlApp         🐬 08:45:58.267 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 08:45:58.268   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 08:45:58.268 ]
💀    🚀 makeMysqlApp         🐬 08:45:58.268 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 08:45:58.281 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 08:45:58.281 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 08:45:58.281 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.755 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.755 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:58.927 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.171 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.178 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.186 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.186 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.186 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.186 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.186 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.19  Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.19  Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.208 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.208 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.211 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.212 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215 ]
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.215 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.255 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.258 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.262 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.41  Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.559 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.721 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.905 Checking start
💀    🚀 makeMysqlAppRunner   🐬 08:45:59.91  Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.086 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.246 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.249 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.392 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.541 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.685 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.832 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.835 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:00.998 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.155 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.158 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.305 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.444 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.447 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.586 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.728 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.731 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:46:01.883 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:46:02.028 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:46:02.031 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 08:46:02.031 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.657257598s
         Current Time: 08:46:02
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.768515176s
         Current Time: 08:46:02
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.104µs
         Current Time: 08:46:02
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:46:02.387 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:46:02.389         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:46:02.389     
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:46:02.389   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:46:02.389   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:46:02.389   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:46:02.389 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 08:46:02.819 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 08:46:02.819 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:02.966 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 08:46:03.178 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:46:03.185 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:46:03.191 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 08:46:03.191 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 08:46:03.192 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 08:46:03.192 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 08:46:03.192 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:46:03.195 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:46:03.195 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 08:46:03.203 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 08:46:03.203 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 08:46:03.206 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 08:46:03.206 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209 ]
💀    🚀 makeFastApiApp       ⚡ 08:46:03.209 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 08:46:03.719 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 08:46:03.72  🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 08:46:03.72  Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 08:46:04.366 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:46:04.366 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.692 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.693 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.693 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.693 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.693 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.914 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.922 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.929 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.929 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.929 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.929 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.929 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.932 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.932 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.947 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.947 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.951 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.951 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.954 ]
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.955 
💀    🚀 makeFastApiAppRunner ⚡ 08:46:05.955 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.005 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.009 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.014 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.186 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.189 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.341 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.49  Checking test
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.494 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.653 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.803 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.806 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:46:06.957 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.123 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.127 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.321 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.509 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.513 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.698 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:07.875 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.07  Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.216 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.219 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.375 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.527 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.53  Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.679 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.827 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.83  Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:46:08.995 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.146 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.149 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.304 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.462 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.637 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.789 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:09.942 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:46:10.093 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:46:10.241 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:10.398 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:46:10.567 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 08:46:10.567 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 08:46:11.118 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 08:46:11.118 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 08:46:11.955 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.955 Preparing start command
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Start command prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Preparing test command
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Test command prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Preparing check command
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Check command prepared
💀    🚀 addFastApiModule     ⚡ 08:46:11.956 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 08:46:12.172 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 08:46:12.179 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 08:46:12.185 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 08:46:12.185 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 08:46:12.185 Validate app directory
💀    🚀 addFastApiModule     ⚡ 08:46:12.185 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 08:46:12.185 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 08:46:12.188 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 08:46:12.188 Validate template locations
💀    🚀 addFastApiModule     ⚡ 08:46:12.196 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 08:46:12.196 Validate app ports
💀    🚀 addFastApiModule     ⚡ 08:46:12.199 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 08:46:12.199 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 08:46:12.202 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 08:46:12.202 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 08:46:12.202 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 08:46:12.202   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 08:46:12.202 ]
💀    🚀 addFastApiModule     ⚡ 08:46:12.202 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 08:46:12.216 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 08:46:12.216 Registering module
💀    🚀 addFastApiModule     ⚡ 08:46:12.239 Done registering module
💀    🚀 addFastApiModule     ⚡ 08:46:12.24  🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 08:46:12.24  Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 08:46:12.551 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 08:46:12.551 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.429 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:13.656 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:13.663 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:13.669 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.669 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:46:13.679 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:46:13.679 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:46:13.742 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:46:13.742 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:46:13.81  Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:46:13.81  Set repo field update
💀    🚀 addFastApiCrud       ⚡ 08:46:13.906 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 08:46:13.906 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:13.967 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:14.232 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:14.238 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:46:14.244 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:46:14.244 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 08:46:14.245 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 08:46:14.245 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 08:46:14.245 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:46:14.247 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:46:14.247 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 08:46:14.257 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 08:46:14.257 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 08:46:14.26  Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 08:46:14.26  Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 ]
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 
💀    🚀 addFastApiCrud       ⚡ 08:46:14.263 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 08:46:14.292 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 08:46:14.292 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:46:14.331 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:46:14.331 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:46:14.381 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:46:14.381 Registering repo
💀    🚀 addFastApiCrud       ⚡ 08:46:14.438 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 08:46:14.439 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 08:46:14.439 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 12.1583609s
         Current Time: 08:46:14
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.359746742s
         Current Time: 08:46:14
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.409µs
         Current Time: 08:46:14
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:46:14.901 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:46:14.903         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:46:14.903     
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:46:14.903   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:46:14.903   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:46:14.903   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:46:14.903 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 08:46:15.351 🧰 Prepare
💀    🚀 makeNginxApp         📗 08:46:15.351 Preparing base variables
💀    🚀 makeNginxApp         📗 08:46:15.431 Base variables prepared
💀    🚀 makeNginxApp         📗 08:46:15.431 Preparing start command
💀    🚀 makeNginxApp         📗 08:46:15.431 Start command prepared
💀    🚀 makeNginxApp         📗 08:46:15.431 Preparing prepare command
💀    🚀 makeNginxApp         📗 08:46:15.431 Prepare command prepared
💀    🚀 makeNginxApp         📗 08:46:15.431 Preparing test command
💀    🚀 makeNginxApp         📗 08:46:15.432 Test command prepared
💀    🚀 makeNginxApp         📗 08:46:15.432 Preparing migrate command
💀    🚀 makeNginxApp         📗 08:46:15.432 Migrate command prepared
💀    🚀 makeNginxApp         📗 08:46:15.432 Preparing check command
💀    🚀 makeNginxApp         📗 08:46:15.432 Check command prepared
💀    🚀 makeNginxApp         📗 08:46:15.432 Preparing replacement map
💀    🚀 makeNginxApp         📗 08:46:15.662 Add config to replacement map
💀    🚀 makeNginxApp         📗 08:46:15.669 Add env to replacement map
💀    🚀 makeNginxApp         📗 08:46:15.676 Replacement map prepared
💀    🚀 makeNginxApp         📗 08:46:15.676 ✅ Validate
💀    🚀 makeNginxApp         📗 08:46:15.676 Validate app directory
💀    🚀 makeNginxApp         📗 08:46:15.676 Done validating app directory
💀    🚀 makeNginxApp         📗 08:46:15.676 Validate app container volumes
💀    🚀 makeNginxApp         📗 08:46:15.679 Done validating app container volumes
💀    🚀 makeNginxApp         📗 08:46:15.679 Validate template locations
💀    🚀 makeNginxApp         📗 08:46:15.688 Done validating template locations
💀    🚀 makeNginxApp         📗 08:46:15.688 Validate app ports
💀    🚀 makeNginxApp         📗 08:46:15.692 Done validating app ports
💀    🚀 makeNginxApp         📗 08:46:15.692 Validate app crud fields
💀    🚀 makeNginxApp         📗 08:46:15.695 Done validating app crud fields
💀    🚀 makeNginxApp         📗 08:46:15.695 🚧 Generate
💀    🚀 makeNginxApp         📗 08:46:15.695 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 08:46:15.695   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 08:46:15.695 ]
💀    🚀 makeNginxApp         📗 08:46:15.695 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 08:46:15.718 🔩 Integrate
💀    🚀 makeNginxApp         📗 08:46:15.718 🎉🎉🎉
💀    🚀 makeNginxApp         📗 08:46:15.718 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 08:46:16.18  🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 08:46:16.18  Preparing base variables
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing start command
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Start command prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing test command
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Test command prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing check command
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Check command prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.267 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 08:46:16.474 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 08:46:16.48  Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 08:46:16.486 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 08:46:16.487 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 08:46:16.487 Validate app directory
💀    🚀 makeNginxAppRunner   📗 08:46:16.487 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 08:46:16.487 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 08:46:16.489 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 08:46:16.489 Validate template locations
💀    🚀 makeNginxAppRunner   📗 08:46:16.501 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 08:46:16.501 Validate app ports
💀    🚀 makeNginxAppRunner   📗 08:46:16.503 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 08:46:16.503 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 08:46:16.506 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 08:46:16.506 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 08:46:16.506 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 08:46:16.506   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 08:46:16.506   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 08:46:16.506 ]
💀    🚀 makeNginxAppRunner   📗 08:46:16.506 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 08:46:16.531 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 08:46:16.536 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 08:46:16.54  Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:46:16.694 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:46:16.845 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:46:16.997 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:46:17.152 Checking start
💀    🚀 makeNginxAppRunner   📗 08:46:17.155 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 08:46:17.317 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:46:17.468 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 08:46:17.471 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 08:46:17.624 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:46:17.78  Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:46:17.933 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:46:18.085 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 08:46:18.089 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 08:46:18.238 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:46:18.384 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 08:46:18.387 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 08:46:18.541 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:46:18.69  Checking buildImages
💀    🚀 makeNginxAppRunner   📗 08:46:18.693 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 08:46:18.845 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:46:18.996 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 08:46:18.999 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 08:46:19.153 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:46:19.302 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 08:46:19.306 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 08:46:19.306 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.513501386s
         Current Time: 08:46:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.625531413s
         Current Time: 08:46:19
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.075µs
         Current Time: 08:46:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:46:19.683 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:46:19.788 Synchronize task environments
💀    🚀 syncEnv              🔄 08:46:20.019 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:46:20.302 🎉🎉🎉
💀    🚀 syncEnv              🔄 08:46:20.302 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 726.137791ms
         Current Time: 08:46:20
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 927.62219ms
         Current Time: 08:46:20
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.531µs
         Current Time: 08:46:20
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:46:20.91  🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:46:20.91  Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoBackendI... 🏭 08:46:21.172 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 08:46:21.173 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:46:21.173 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:46:21.841 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 08:46:21.841 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:46:21.891 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:46:21.891  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:46:21.891 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892  ---> 162e06eadcfd
💀    🚀 buildDemoBackendI... 🏭 08:46:21.892 Sending build context to Docker daemon   1.03MB
💀    🚀 buildDemoFrontend... 🏭 08:46:21.892 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoDbImage     🏭 08:46:21.895 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:46:21.895  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:46:21.896 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 08:46:21.896 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9    ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9   Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoDbImage     🏭 08:46:21.9   Successfully tagged demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9    ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9    ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9   Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9    ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9    ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 08:46:21.9   Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:46:21.901 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 08:46:21.901  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.901  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 08:46:21.901 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902 Step 7/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 08:46:21.902 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:46:21.902 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902  ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 08:46:21.902 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 08:46:21.903 Successfully built 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 08:46:21.909 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:46:21.911 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:46:21.911 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> 97fdfef7cb48
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> bf9c545afbe0
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> a62a483a9091
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:46:21.912  ---> db465fe79375
💀    🚀 buildDemoBackendI... 🏭 08:46:21.913 Successfully built db465fe79375
💀    🚀 buildDemoBackendI... 🏭 08:46:21.919 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:46:21.921 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:46:21.921 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 08:46:22.028 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 1.2237247s
         Current Time: 08:46:22
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.334897621s
         Current Time: 08:46:22
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 5.107µs
         Current Time: 08:46:22
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:46:22.394 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:46:22.394 Links updated
💀    🚀 prepareDemoBackend   🔧 08:46:22.395 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 08:46:22.422 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 08:46:22.506 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 08:46:22.652 Build image demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:46:23.187 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 08:46:23.188 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:46:23.226 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:46:23.226  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:46:23.226 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:46:23.226  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.226  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:46:23.227 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 08:46:23.229 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23  Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23  Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23  Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23  Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23   ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 08:46:23.23  Successfully built 6ea76668c578
💀    🚀 buildDemoDbImage     🏭 08:46:23.233 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:46:23.234  ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 08:46:23.235 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 08:46:23.235 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 08:46:23.236 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:46:23.236 Docker image demo-frontend built
💀    🚀 buildDemoDbImage     🏭 08:46:23.238 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:46:23.239 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:46:23.239 Docker image demo-db built
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:46:23.535 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:46:23.557 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🚀 startDemoFrontend... 📗 08:46:23.56  Error: No such container: demoFrontend
💀 🔥 🔎 startDemoFrontend... 📗 08:46:23.563 Error: No such container: demoFrontend
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:23.585 Error: No such container: demoDb
💀 🔥 🚀 startDemoFrontend... 📗 08:46:23.591 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 08:46:23.593 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:23.603 Error: No such container: demoDb
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:23.63  Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 08:46:23.632 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 08:46:23.65  809210155ac7150e174bd50f7ee683093cec68b20acc2e19e21032ef124daf26
💀    🚀 startDemoDbContainer 🐬 08:46:23.691 a409c5c3fe5fab0f885a25cc926f563d4eb4fdc30a4afcfd17dd40d93d926097
💀    🚀 prepareDemoBackend   🔧 08:46:24.138 Activate venv
💀    🚀 prepareDemoBackend   🔧 08:46:24.139 Install dependencies
💀    🚀 prepareDemoBackend   🔧 08:46:24.429 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 08:46:24.622   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:24.628 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 08:46:24.702   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:24.71  Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 08:46:24.78    Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoFrontend... 📗 08:46:25.941 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 08:46:25.944 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoDbContainer 🐬 08:46:25.958 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 08:46:25.968 🔎 Waiting docker container 'demoDb' healthcheck
💀    🔎 startDemoFrontend... 📗 08:46:26.03  🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 08:46:26.03  🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 08:46:26.035 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 08:46:26.035 🔎 Waiting for host port: '443'
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.035 [38;5;6mmysql [38;5;5m01:46:26.00 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.035 [38;5;6mmysql [38;5;5m01:46:26.01 [38;5;2mINFO  ==> Initializing mysql database
💀    🔎 startDemoFrontend... 📗 08:46:26.037 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.042 [38;5;6mmysql [38;5;5m01:46:26.03 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.047 [38;5;6mmysql [38;5;5m01:46:26.04 [38;5;2mINFO  ==> Setting user option
💀    🔎 startDemoDbContainer 🐬 08:46:26.053 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 08:46:26.053 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 08:46:26.054 🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.059 [38;5;6mmysql [38;5;5m01:46:26.05 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.067 [38;5;6mmysql [38;5;5m01:46:26.06 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:26.076 [38;5;6mmysql [38;5;5m01:46:26.07 [38;5;2mINFO  ==> Installing database
💀    🚀 prepareDemoBackend   🔧 08:46:27.468 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:46:27.606   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:27.622 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 08:46:27.735   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:27.746 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 08:46:27.838   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:27.866 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 08:46:27.94    Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:27.95  Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 08:46:28.138   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:28.22  Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 08:46:28.591   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:28.671 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 08:46:28.832   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:28.875 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🔎 startDemoFrontend... 📗 08:46:29.043 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 08:46:29.058 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 08:46:29.186 check demoFrontend
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:29.191 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoFrontend... 📗 08:46:29.192 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:29.193 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 08:46:29.251   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:46:29.448 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 08:46:29.591   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:29.601 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧 08:46:29.683   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:29.69  Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 08:46:29.744   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:29.751 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 08:46:29.827   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:29.837 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 08:46:29.917   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:29.939 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 08:46:30.008   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.019 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 08:46:30.197   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.214 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 08:46:30.327   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.336 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:46:30.454   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.475 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 08:46:30.555   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.564 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:46:30.643   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:30.652 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 08:46:30.715   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:46:30.86  Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 08:46:30.989   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:31.009 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:31.134 [38;5;6mmysql [38;5;5m01:46:31.13 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 08:46:31.465   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:31.56  Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 08:46:31.676   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:31.694 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 08:46:31.784   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:31.793 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧 08:46:31.91    Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:31.924 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 08:46:32.031   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:32.049 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 08:46:32.136   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:46:32.351 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:32.365 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:32.368 ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: YES)
💀    🚀 prepareDemoBackend   🔧 08:46:32.475   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:32.492 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:46:32.797   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:32.812 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:46:32.898   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:32.905 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:46:32.995   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.006 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:46:33.094   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.107 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:33.148 [38;5;6mmysql [38;5;5m01:46:33.14 [38;5;2mINFO  ==> Configuring authentication
💀    🚀 prepareDemoBackend   🔧 08:46:33.179   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.188 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🔎 startDemoFrontend... 📗 08:46:33.194 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 08:46:33.194 📜 Task 'startDemoFrontendContainer' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:33.197 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:33.221 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:33.271 [38;5;6mmysql [38;5;5m01:46:33.27 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:33.276 [38;5;6mmysql [38;5;5m01:46:33.27 [38;5;2mINFO  ==> Stopping mysql
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀    🚀 prepareDemoBackend   🔧 08:46:33.295   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.309 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 08:46:33.443   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.479 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:46:33.619   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:33.629 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 08:46:34.055   Using cached https://files.pythonhosted.org/packages/c1/38/a9fd8c7bb151325d8b3d9108ce791348c84171b5d9f346b0bf0639de603f/coverage-6.3.3-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.073 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:46:34.169   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.176 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:34.292 [38;5;6mmysql [38;5;5m01:46:34.29 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 08:46:34.369   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.382 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:46:34.491   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.507 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:46:34.573   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.583 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:46:34.769   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:34.784 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 08:46:34.882   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:46:35.139 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, urllib3, idna, requests, fastavro, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, pluggy, py, iniconfig, pyparsing, packaging, attrs, toml, pytest, tomli, coverage, pytest-cov, pyasn1, rsa, ecdsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 08:46:35.162   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 08:46:35.348     Running setup.py install for avro-python3: finished with status 'done'
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:35.515 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:35.516 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 08:46:35.728   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:38.663 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:38.664 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:40.302 [38;5;6mmysql [38;5;5m01:46:40.30 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:40.316 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:40.32  find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:40.323 [38;5;6mmysql [38;5;5m01:46:40.32 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:41.77  mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:41.771 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:42.336 [38;5;6mmysql [38;5;5m01:46:42.33 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 08:46:42.336 
💀 🔥 🚀 startDemoDbContainer 🐬 08:46:42.359 [38;5;6mmysql [38;5;5m01:46:42.35 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 08:46:42.558 2022-05-14T01:46:42.553677Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 08:46:42.558 2022-05-14T01:46:42.554608Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 08:46:42.558 2022-05-14T01:46:42.554615Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 08:46:42.559 2022-05-14T01:46:42.559424Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 08:46:42.668 2022-05-14T01:46:42.668207Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 08:46:42.82  2022-05-14T01:46:42.819529Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 08:46:42.82  2022-05-14T01:46:42.819574Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 08:46:42.832 2022-05-14T01:46:42.831675Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 08:46:42.832 2022-05-14T01:46:42.831701Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀 🔥 🔎 startDemoDbContainer 🐬 08:46:44.902 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 Database
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 information_schema
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 mysql
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 performance_schema
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 sample
💀    🔎 startDemoDbContainer 🐬 08:46:44.906 sys
💀    🔎 startDemoDbContainer 🐬 08:46:44.91  🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 08:46:48.912 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 08:46:48.912 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 08:46:57.68      Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:46:58.763   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 08:46:58.923     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:46:59.442   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 08:46:59.579     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:46:59.63  Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.3 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 08:46:59.687 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 08:46:59.688 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 08:46:59.738 Prepare
💀    🚀 prepareDemoBackend   🔧 08:46:59.738 prepare command
💀    🚀 prepareDemoBackend   🔧 08:46:59.738 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 startDemoBackend     ⚡ 08:46:59.973 Activate venv
💀    🔎 startDemoBackend     ⚡ 08:46:59.973 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 08:46:59.974 Start
💀    🚀 startDemoBackend     ⚡ 08:47:00.362 2022-05-14 08:47:00,362 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 08:47:00.362 2022-05-14 08:47:00,362 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.366 2022-05-14 08:47:00,366 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 08:47:00.366 2022-05-14 08:47:00,366 INFO sqlalchemy.engine.Engine [generated in 0.00022s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.369 2022-05-14 08:47:00,368 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 08:47:00.369 2022-05-14 08:47:00,369 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.37  2022-05-14 08:47:00,370 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.371 2022-05-14 08:47:00,370 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:47:00.371 2022-05-14 08:47:00,370 INFO sqlalchemy.engine.Engine [generated in 0.00012s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 08:47:00.373 2022-05-14 08:47:00,373 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:47:00.373 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 08:47:00.373 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 )
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 
💀    🚀 startDemoBackend     ⚡ 08:47:00.374 2022-05-14 08:47:00,373 INFO sqlalchemy.engine.Engine [no key 0.00011s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.393 2022-05-14 08:47:00,393 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 08:47:00.393 2022-05-14 08:47:00,393 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.417 2022-05-14 08:47:00,416 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 08:47:00.417 2022-05-14 08:47:00,417 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.435 2022-05-14 08:47:00,435 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 08:47:00.435 2022-05-14 08:47:00,435 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.453 2022-05-14 08:47:00,453 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.453 2022-05-14 08:47:00,453 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.476 2022-05-14 08:47:00,476 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:47:00.478 2022-05-14 08:47:00,478 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.478 2022-05-14 08:47:00,478 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:47:00.478 2022-05-14 08:47:00,478 INFO sqlalchemy.engine.Engine [cached since 0.1075s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  2022-05-14 08:47:00,479 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  )
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  
💀    🚀 startDemoBackend     ⚡ 08:47:00.48  2022-05-14 08:47:00,480 INFO sqlalchemy.engine.Engine [no key 0.00009s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.501 2022-05-14 08:47:00,501 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 08:47:00.501 2022-05-14 08:47:00,501 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.519 2022-05-14 08:47:00,519 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 08:47:00.519 2022-05-14 08:47:00,519 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.537 2022-05-14 08:47:00,537 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.537 2022-05-14 08:47:00,537 INFO sqlalchemy.engine.Engine [no key 0.00023s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.554 2022-05-14 08:47:00,554 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:47:00.555 2022-05-14 08:47:00,555 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.555 2022-05-14 08:47:00,555 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:47:00.556 2022-05-14 08:47:00,555 INFO sqlalchemy.engine.Engine [cached since 0.1851s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 2022-05-14 08:47:00,557 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.557 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:47:00.558 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.558 )
💀    🚀 startDemoBackend     ⚡ 08:47:00.558 
💀    🚀 startDemoBackend     ⚡ 08:47:00.558 
💀    🚀 startDemoBackend     ⚡ 08:47:00.558 2022-05-14 08:47:00,557 INFO sqlalchemy.engine.Engine [no key 0.00010s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.598 2022-05-14 08:47:00,598 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 08:47:00.598 2022-05-14 08:47:00,598 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.625 2022-05-14 08:47:00,625 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 08:47:00.625 2022-05-14 08:47:00,625 INFO sqlalchemy.engine.Engine [no key 0.00029s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.658 2022-05-14 08:47:00,658 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 08:47:00.658 2022-05-14 08:47:00,658 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.695 2022-05-14 08:47:00,694 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 08:47:00.695 2022-05-14 08:47:00,695 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.711 2022-05-14 08:47:00,711 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 08:47:00.711 2022-05-14 08:47:00,711 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.729 2022-05-14 08:47:00,729 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 08:47:00.729 2022-05-14 08:47:00,729 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:47:00.749 2022-05-14 08:47:00,749 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:47:00.751 2022-05-14 08:47:00,751 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.753 2022-05-14 08:47:00,753 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 08:47:00.753 FROM users 
💀    🚀 startDemoBackend     ⚡ 08:47:00.754 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 08:47:00.754  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 08:47:00.754 2022-05-14 08:47:00,753 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 08:47:00.755 2022-05-14 08:47:00,755 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 08:47:00.94  2022-05-14 08:47:00,940 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.941 2022-05-14 08:47:00,941 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 08:47:00.941 2022-05-14 08:47:00,941 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {'id': '5d164c8c-3b56-4883-a646-98ea48138100', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$L7PUjMsOc6pafymw.s/WSewwuhFRBenthMVaU.aGougjBso0ZovUG', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 14, 8, 47, 0, 940210), 'updated_at': datetime.datetime(2022, 5, 14, 8, 47, 0, 941688)}
💀    🚀 startDemoBackend     ⚡ 08:47:00.943 2022-05-14 08:47:00,943 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:47:00.95  2022-05-14 08:47:00,950 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:47:00.951 2022-05-14 08:47:00,951 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 08:47:00.951 FROM users 
💀    🚀 startDemoBackend     ⚡ 08:47:00.951 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 08:47:00.951 2022-05-14 08:47:00,951 INFO sqlalchemy.engine.Engine [generated in 0.00012s] {'pk_1': '5d164c8c-3b56-4883-a646-98ea48138100'}
💀    🚀 startDemoBackend     ⚡ 08:47:00.952 2022-05-14 08:47:00,952 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 08:47:00.954 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.962 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 08:47:00.971 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.977 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 08:47:00.977 Register library route handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.977 Register library event handler
💀    🚀 startDemoBackend     ⚡ 08:47:00.977 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 08:47:00.977 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:00.978 INFO:     Started server process [5805]
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:00.978 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:00.978 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:00.979 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 08:47:01.979 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 08:47:01.979 check demoBackend
💀    🔎 startDemoBackend     ⚡ 08:47:01.979 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 08:47:01.979 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 08:47:02.085 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 39.798997481s
         Current Time: 08:47:02
         Active Process:
           * (PID=5422) 📗 'startDemoFrontendContainer' service
           * (PID=5801) ⚡ 'startDemoBackend' service
           * (PID=5450) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=5422)
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=5801)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=5450)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:03.69  INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:03.791 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:03.791 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 08:47:03.791 INFO:     Finished server process [5805]
💀    🚀 startDemoBackend     ⚡ 08:47:03.865 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 08:47:03.865 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 41.901734096s
         Current Time: 08:47:04
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.116µs
         Current Time: 08:47:04
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:47:04.441 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:47:04.441 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 08:47:04.468 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 08:47:04.553 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoDbImage     🏭 08:47:04.7   Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:47:04.7   Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 08:47:05.177 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:47:05.178 Sending build context to Docker daemon  22.02kB
💀    🚀 buildDemoDbImage     🏭 08:47:05.233 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:47:05.233  ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 08:47:05.233 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:47:05.233  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:47:05.233 Step 2/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 08:47:05.234 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 08:47:05.235 Sending build context to Docker daemon   1.18MB
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239 Step 4/11 : USER 1001
💀    🚀 buildDemoDbImage     🏭 08:47:05.239 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:47:05.239 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoDbImage     🏭 08:47:05.242 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:47:05.242 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoFrontend... 🏭 08:47:05.243  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.243  ---> 11c677f847bc
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243  ---> c9a3cbe90f60
💀    🚀 buildDemoFrontend... 🏭 08:47:05.243 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoBackendI... 🏭 08:47:05.243 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> 776095918b33
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> 48dc42a93a8a
💀    🚀 buildDemoBackendI... 🏭 08:47:05.244  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoBackendI... 🏭 08:47:05.244  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:47:05.244 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:47:05.244  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.244  ---> 0beee76410dd
💀    🚀 buildDemoBackendI... 🏭 08:47:05.245  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:47:05.245 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245  ---> 68555ae22bc5
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245  ---> 992fa94aa2f2
💀    🚀 buildDemoFrontend... 🏭 08:47:05.245 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:47:05.246  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:47:05.246  ---> 02304e445f6f
💀    🚀 buildDemoFrontend... 🏭 08:47:05.246 Successfully built 02304e445f6f
💀    🚀 buildDemoFrontend... 🏭 08:47:05.254 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:47:05.256 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:47:05.256 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 08:47:05.422  ---> a963024b3190
💀    🚀 buildDemoBackendI... 🏭 08:47:05.422 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:47:05.446  ---> Running in fc2ed96ca8fc
💀    🚀 buildDemoBackendI... 🏭 08:47:05.501 Removing intermediate container fc2ed96ca8fc
💀    🚀 buildDemoBackendI... 🏭 08:47:05.501  ---> 0b22791006bf
💀    🚀 buildDemoBackendI... 🏭 08:47:05.501 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:47:05.531  ---> Running in f5c69c1e3a8e
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:47:05.603 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:47:05.623 🔎 Waiting docker container 'demoDb' running status
💀    🔎 startDemoFrontend... 📗 08:47:05.639 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoFrontend... 📗 08:47:05.642 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 08:47:05.642 🐳 Logging 'demoFrontend'
💀    🔎 startDemoDbContainer 🐬 08:47:05.677 🔎 Waiting docker container 'demoDb' healthcheck
💀    🚀 startDemoDbContainer 🐬 08:47:05.685 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 08:47:05.685 🐳 Logging 'demoDb'
💀    🔎 startDemoFrontend... 📗 08:47:05.688 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 08:47:05.688 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 08:47:05.689 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 08:47:05.689 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 08:47:05.691 🔎 Host port '443' is ready
💀    🔎 startDemoDbContainer 🐬 08:47:05.72  🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 08:47:05.72  🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 08:47:05.721 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 08:47:06.222 Removing intermediate container f5c69c1e3a8e
💀    🚀 buildDemoBackendI... 🏭 08:47:06.222  ---> 3e7f8e114cfb
💀    🚀 buildDemoBackendI... 🏭 08:47:06.222 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:47:06.242  ---> Running in 0837ba35fd47
💀    🚀 buildDemoBackendI... 🏭 08:47:06.298 Removing intermediate container 0837ba35fd47
💀    🚀 buildDemoBackendI... 🏭 08:47:06.298  ---> 218bfcbb31ee
💀    🚀 buildDemoBackendI... 🏭 08:47:06.3   Successfully built 218bfcbb31ee
💀    🚀 buildDemoBackendI... 🏭 08:47:06.306 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:47:06.309 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:47:06.309 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 08:47:08.694 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 08:47:08.724 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 08:47:08.821 check demoFrontend
💀    🔎 startDemoFrontend... 📗 08:47:08.826 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 08:47:08.839 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 Database
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 information_schema
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 mysql
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 performance_schema
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 sample
💀    🔎 startDemoDbContainer 🐬 08:47:08.842 sys
💀    🔎 startDemoDbContainer 🐬 08:47:08.846 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 08:47:12.828 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 08:47:12.828 📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 08:47:12.848 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 08:47:12.848 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 08:47:13.466 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🔎 startDemoBackendC... ⚡ 08:47:13.492 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:13.492 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:13.515 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 08:47:13.517 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 08:47:13.562 e854c32b31a59fbbbf78330a3fceed9585215cbca86608e99ff5b0cc813cea7a
💀    🚀 startDemoBackendC... ⚡ 08:47:14.859 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:47:14.861 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 08:47:14.891 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 08:47:14.891 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 08:47:14.893 🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 08:47:15.262 2022-05-14 01:47:15,261 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 08:47:15.262 2022-05-14 01:47:15,262 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.266 2022-05-14 01:47:15,265 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 08:47:15.266 2022-05-14 01:47:15,265 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.269 2022-05-14 01:47:15,269 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 08:47:15.269 2022-05-14 01:47:15,269 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.273 2022-05-14 01:47:15,272 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:47:15.273 2022-05-14 01:47:15,273 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:47:15.273 2022-05-14 01:47:15,273 INFO sqlalchemy.engine.Engine [generated in 0.00013s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.276 2022-05-14 01:47:15,275 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:47:15.278 2022-05-14 01:47:15,278 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:47:15.278 2022-05-14 01:47:15,278 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:47:15.278 2022-05-14 01:47:15,278 INFO sqlalchemy.engine.Engine [cached since 0.005334s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.28  2022-05-14 01:47:15,280 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:47:15.282 2022-05-14 01:47:15,282 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:47:15.283 2022-05-14 01:47:15,282 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:47:15.283 2022-05-14 01:47:15,282 INFO sqlalchemy.engine.Engine [cached since 0.009695s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.286 2022-05-14 01:47:15,285 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:47:15.291 2022-05-14 01:47:15,290 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:47:15.293 2022-05-14 01:47:15,293 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 08:47:15.293 FROM users 
💀    🚀 startDemoBackendC... ⚡ 08:47:15.293 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 08:47:15.294  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 08:47:15.294 2022-05-14 01:47:15,293 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 08:47:15.296 2022-05-14 01:47:15,296 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 08:47:15.3   Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.309 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 08:47:15.32  Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.327 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 08:47:15.327 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.327 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 08:47:15.327 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 08:47:15.327 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:15.328 INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:15.328 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:15.328 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 08:47:15.328 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 08:47:17.896 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:47:18.02  check demoBackend
💀    🔎 startDemoBackendC... ⚡ 08:47:18.023 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:47:19.025 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 08:47:19.025 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 08:47:19.132 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 14.79840567s
         Current Time: 08:47:19
         Active Process:
           * (PID=7003) ⚡ 'startDemoBackendContainer' service
           * (PID=6846) 📗 'startDemoFrontendContainer' service
           * (PID=6869) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=6869)
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=7003)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=6846)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 16.901914226s
         Current Time: 08:47:21
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.084µs
         Current Time: 08:47:21
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:47:21.491 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:47:21.491 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoFrontendC... ✋ 08:47:21.864 Stop docker container demoFrontend
💀    🚀 stopDemoBackendCo... ✋ 08:47:21.866 Stop docker container demoBackend
💀    🚀 stopDemoDbContainer  ✋ 08:47:21.867 Stop docker container demoDb
💀    🚀 stopDemoDbContainer  ✋ 08:47:25.914 demoDb
💀    🚀 stopDemoDbContainer  ✋ 08:47:25.916 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 08:47:25.916 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ 08:47:32.461 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 08:47:32.463 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 08:47:32.463 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀    🚀 stopDemoFrontendC... ✋ 08:47:32.9   demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 08:47:32.901 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 08:47:32.901 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 08:47:33.008 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.622756518s
         Current Time: 08:47:33
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.73344322s
         Current Time: 08:47:33
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.163µs
         Current Time: 08:47:33
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:47:33.372 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:47:33.372 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🔥 🚀 removeDemoDbConta... ❌ 08:47:33.686 Error: No such container: 
💀 🔥 🚀 removeDemoBackend... ❌ 08:47:33.687 Error: No such container: 
💀    🚀 removeDemoDbConta... ❌ 08:47:33.689 Stop docker container demoDb
💀    🚀 removeDemoBackend... ❌ 08:47:33.689 Stop docker container demoBackend
💀 🔥 🚀 removeDemoFronten... ❌ 08:47:33.7   Error: No such container: 
💀    🚀 removeDemoFronten... ❌ 08:47:33.702 Stop docker container demoFrontend
💀    🚀 removeDemoDbConta... ❌ 08:47:33.791 Docker container demoDb stopped
💀    🚀 removeDemoBackend... ❌ 08:47:33.791 Docker container demoBackend stopped
💀    🚀 removeDemoDbConta... ❌ 08:47:33.791 Remove docker container demoDb
💀    🚀 removeDemoBackend... ❌ 08:47:33.791 Remove docker container demoBackend
💀    🚀 removeDemoFronten... ❌ 08:47:33.815 Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ 08:47:33.815 Remove docker container demoFrontend
💀    🚀 removeDemoBackend... ❌ 08:47:33.868 demoBackend
💀    🚀 removeDemoBackend... ❌ 08:47:33.876 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ 08:47:33.876 Docker container demoBackend removed
💀    🚀 removeDemoDbConta... ❌ 08:47:33.89  demoDb
💀    🚀 removeDemoDbConta... ❌ 08:47:33.892 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ 08:47:33.892 Docker container demoDb removed
💀    🚀 removeDemoFronten... ❌ 08:47:33.906 demoFrontend
💀    🚀 removeDemoFronten... ❌ 08:47:33.909 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ 08:47:33.909 Docker container demoFrontend removed
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 removeContainers     ❌ 08:47:34.015 
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 749.52011ms
         Current Time: 08:47:34
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 860.333027ms
         Current Time: 08:47:34
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.302µs
         Current Time: 08:47:34
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:47:34.405 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:47:34.41          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:47:34.41      
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:47:34.41    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:47:34.41    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:47:34.41    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:47:34.41  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.832 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.832 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.994 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.995 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.995 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.995 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:34.995 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.214 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.22  Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.227 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.227 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.227 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.227 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.227 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.23  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.23  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.239 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.239 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.242 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.242 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.245 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.276 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.276 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.276 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.559 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.559 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:35.744 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.043 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.049 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.056 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.056 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.056 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.056 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.056 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.06  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.06  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.071 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.071 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.075 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.075 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.078 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.096 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.1   Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.103 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.259 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.412 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.415 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.57  Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.729 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.732 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.883 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.883 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:36.883 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.586776187s
         Current Time: 08:47:36
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.697413567s
         Current Time: 08:47:37
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.796µs
         Current Time: 08:47:37
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:47:37.253 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:47:37.256         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:47:37.256     
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:47:37.256   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:47:37.256   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:47:37.256   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:47:37.256 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:37.682 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:37.683 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.541 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.541 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.542 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.771 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.778 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.784 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.785 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.785 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.785 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.785 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.788 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.788 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.798 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.798 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.801 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.801 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.805 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.839 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.84  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:38.84  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:39.193 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:39.193 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.04  Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.267 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.273 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.28  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.28  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.28  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.28  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.28  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.283 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.283 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.294 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.294 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.297 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.297 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3   Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3   🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3   🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3     "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3   ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.3   🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.318 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.322 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.325 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.481 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.64  Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.644 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.802 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.978 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:40.981 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:41.177 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:41.177 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:41.177 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.031160888s
         Current Time: 08:47:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.142429838s
         Current Time: 08:47:41
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.293µs
         Current Time: 08:47:41
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:47:41.603 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:47:41.606         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:47:41.606     
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:47:41.606   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:47:41.606   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:47:41.606   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:47:41.606 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.062 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.063 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.201 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.202 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.428 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.437 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.445 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.445 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.445 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.445 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.445 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.449 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.449 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.46  Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.46  Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.463 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.463 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.466 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.497 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.497 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.497 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.839 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.839 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.954 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:42.955 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.177 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.184 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.19  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.19  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.19  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.19  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.19  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.193 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.193 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.202 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.202 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.205 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.205 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.209 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.21  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.21  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.21    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.21  ]
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.21  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.23  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.234 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.237 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.392 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.546 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.549 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.707 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.86  Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:43.863 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:47:44.02  Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:47:44.02  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:47:44.021 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.525710407s
         Current Time: 08:47:44
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.637090141s
         Current Time: 08:47:44
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.1µs
         Current Time: 08:47:44
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:47:44.398 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:47:44.503 Synchronize task environments
💀    🚀 syncEnv              🔄 08:47:44.683 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:47:44.842 🎉🎉🎉
💀    🚀 syncEnv              🔄 08:47:44.842 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 549.745706ms
         Current Time: 08:47:44
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 750.626289ms
         Current Time: 08:47:45
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.053µs
         Current Time: 08:47:45
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoFronte... 🏁 08:47:45.332 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:45.332 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 08:47:45.337 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 08:47:47.259 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 08:47:47.264 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:47:47.309 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:47.538 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:47.541 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:47.592 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.303   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.322 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:48.364   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:48.369   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:48.385 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:48.389 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.485   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 08:47:48.599   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 08:47:48.605   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.748 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:48.861 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:48.867 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.878   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:48.883 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:49.006   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:49.024 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:49.116   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:49.121 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:49.553   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:49.613   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:49.625 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:49.681 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:49.727   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:49.746 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:49.807   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:49.813   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:49.826 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:49.886 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:47:49.991   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.006 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.056   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.083 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.128   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.15  Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.166   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.175 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.221   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.228 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.26    Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.265 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.301   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.31    Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.311 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.331 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.348   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.353 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.389   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.4   Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.403   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.407 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.431   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.44  Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.486   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.506 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.527   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.547 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.559   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.58  Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.613   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.632   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.633 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.67  Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.672   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.687 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.712   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.743 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.76    Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.785 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.794   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.808 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.817   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.83  Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:50.887   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.892   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:50.896 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.903 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:50.946   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:50.952 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.984   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:50.994 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:51.014   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:51.038 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:51.045   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:51.068 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:51.086   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:51.107 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:47:51.122   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:51.128 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:47:51.163   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:51.17  Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:47:51.187   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:47:51.266   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:47:51.269   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:47:51.288 Installing collected packages: six, grpcio, pyyaml, protobuf, semver, dill, pulumi, attrs, arpeggio, parver, charset-normalizer, idna, urllib3, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 08:47:51.374 Installing collected packages: pyyaml, six, grpcio, dill, protobuf, semver, pulumi, attrs, arpeggio, parver, charset-normalizer, urllib3, certifi, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 08:47:51.38  Installing collected packages: six, grpcio, pyyaml, protobuf, dill, semver, pulumi, arpeggio, attrs, parver, idna, urllib3, certifi, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 08:47:51.927   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 08:47:51.965   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 08:47:52.035   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.32      Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.368 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 08:47:53.381     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:47:53.398 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:47:53.398 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.429 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 08:47:53.463 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:47:53.463 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 08:47:53.531     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 08:47:53.583 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoBacken... 🏁 08:47:53.615 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:47:53.615 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.678 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.678 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoFronte... 🏁 08:47:53.678 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.678 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.722 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.723 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.723       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 Usage:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.724 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:47:53.725 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 
💀    🚀 prepareDemoFronte... 🏁 08:47:53.725 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 08:47:53.945 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 08:47:54.042 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 08:47:54.053 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 08:47:54.143 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 08:47:54.722 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:47:54.811 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"ClusterIP"}
💀    🚀 prepareDemoBacken... 🏁 08:47:54.811 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.866 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 for this case.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Usage:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Flags:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 
💀    🚀 prepareDemoBacken... 🏁 08:47:54.867 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 08:47:54.868 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 08:47:55.105 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 08:47:55.202 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 08:47:55.765 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:47:55.83  Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 08:47:56.136 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.189 
💀    🚀 deployDemoFronten... 🏁 08:47:56.469  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.532  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:47:56.539  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.607  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 08:47:56.771  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 08:47:56.774  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.852  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.856  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoBackend... 🏁 08:47:56.892 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 08:47:56.905  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:47:56.905  
💀    🚀 deployDemoFronten... 🏁 08:47:56.905 Resources:
💀    🚀 deployDemoFronten... 🏁 08:47:56.905     + 4 to create
💀    🚀 deployDemoFronten... 🏁 08:47:56.905 
💀    🚀 deployDemoFronten... 🏁 08:47:56.905 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988  
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:56.988 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 08:47:57.326 
💀    🚀 deployDemoFronten... 🏁 08:47:57.329 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:57.403 
💀    🚀 deployDemoFronten... 🏁 08:47:57.69   +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 08:47:57.711  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:47:57.758  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:57.768  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 08:47:57.786  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:57.867  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 08:47:58.051  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:47:58.054  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:47:58.078  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:47:58.084  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:47:58.09   +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 08:47:58.102  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoBackend... 🏁 08:47:58.15   +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:47:58.151  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:47:58.162  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.21   +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.215  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.228  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.232  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.254  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.254  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoFronten... 🏁 08:47:58.324  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 08:47:58.324  
💀    🚀 deployDemoFronten... 🏁 08:47:58.325 Outputs:
💀    🚀 deployDemoFronten... 🏁 08:47:58.325     app: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326         ready    : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.326             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326         ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.326         resources: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.326                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                             }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                             spec      : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                 selector: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                 template: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         labels: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                 ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                             }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                             }
💀    🚀 deployDemoFronten... 🏁 08:47:58.327                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.328 
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     labels            : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                         [0]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.328                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.329                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                              }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                          f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                              f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                  k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                      f:env                     : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                      }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                                  }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                              }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                      }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                  }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              time       : "2022-05-14T01:47:58Z"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      resource_version  : "190261"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      uid               : "6dad8f2b-760f-4dfa-992a-032c9469b7aa"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                  }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                  spec       : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      replicas                 : 1
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      selector                 : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          match_labels: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      strategy                 : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          rolling_update: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                      template                 : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          metadata: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              labels: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                  app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                  app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                          spec    : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                              containers                      : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                  [0]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                      env                       : [
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                          [0]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                              name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                              value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 08:47:58.33                                          }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         [1]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         [2]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         [3]: {
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                             value: "1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             ]
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                         }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                     }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                 }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331             }
💀    🚀 deployDemoFronten... 🏁 08:47:58.331             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁 08:47:58.331                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:47:58.332                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 08:47:58.332                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.332                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 08:47:58.332        
💀    🚀 deployDemoBackend... 🏁 08:47:58.35   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:47:58.35   
💀    🚀 deployDemoBackend... 🏁 08:47:58.35  Resources:
💀    🚀 deployDemoBackend... 🏁 08:47:58.35      + 5 to create
💀    🚀 deployDemoBackend... 🏁 08:47:58.35  
💀    🚀 deployDemoBackend... 🏁 08:47:58.35  Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.422  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.422  
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423     app: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.423                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.424                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             time       : "2022-05-14T01:47:58Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     resource_version  : "190277"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     uid               : "9b69bd60-2b49-4dc6-bb07-fabfb225fcea"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.425                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                             time       : "2022-05-14T01:47:58Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     resource_version  : "190278"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                     uid               : "8ac06424-9384-48f6-b060-b6797074b3dc"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.426             }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427         }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427     }
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427 Duration: 2s
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.427 
💀    🚀 deployDemoDbDeplo... 🏁 08:47:58.428 hello world
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoBackend... 🏁 08:47:58.834 
💀    🚀 deployDemoBackend... 🏁 08:47:59.205  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.268  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.523  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.526  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.53   +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.542  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.545  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.554  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:47:59.555  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:47:59.577  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:47:59.588  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:47:59.796  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 08:47:59.796  
💀    🚀 deployDemoBackend... 🏁 08:47:59.798 Outputs:
💀    🚀 deployDemoBackend... 🏁 08:47:59.799     app: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799         ready    : [
💀    🚀 deployDemoBackend... 🏁 08:47:59.799             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799         ]
💀    🚀 deployDemoBackend... 🏁 08:47:59.799         resources: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 }
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 }
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                             }
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                             spec      : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 08:47:59.799                                 selector: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                       matchLabels: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                           app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                           app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                       }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                   }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                   template: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                       metadata: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                           labels: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                               app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                               app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                           }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                       }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                       spec    : {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                           containers        : [
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                               [0]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                   env            : [
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       [0]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           value: "HS256"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       [1]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           value: "30"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       [2]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       }
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                       [3]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.8                                                           name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                         name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                         value: "false"
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.801                                                         name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         ]
💀    🚀 deployDemoBackend... 🏁 08:47:59.802 
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         ]
💀    🚀 deployDemoBackend... 🏁 08:47:59.802 
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         ]
💀    🚀 deployDemoBackend... 🏁 08:47:59.802 
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.802                                                         name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: (json) []
💀    🚀 deployDemoBackend... 🏁 08:47:59.803 
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "600"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.803                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:47:59.804                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     }
💀    🚀 deployDemoBackend... 🏁 08:47:59.805                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 08:47:59.805   
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 08:47:59.912 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 14.704796038s
         Current Time: 08:48:00
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.816486906s
         Current Time: 08:48:00
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.261µs
         Current Time: 08:48:00
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.446 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 08:48:00.446 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:48:00.448 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 08:48:00.832 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.832 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.836 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.838 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.839 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.842 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.845 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.849 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.85  Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.852 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.852 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.852 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.854 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.854 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.857 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.858 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.859 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.859 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.861 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.863 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.864 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.866 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.866 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.868 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.868 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.87  Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.872 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.876 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.88  Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.885 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.893 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.898 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.905 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.906 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.92  Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.942 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.946 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.949 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 08:48:00.952 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:00.954 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.959 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.963 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.973 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 08:48:00.977 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.003 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 08:48:01.003 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 08:48:01.007 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 08:48:01.013 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:48:01.043 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:48:01.043 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:48:01.053 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:48:01.053 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:48:01.078 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:48:01.078 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.494 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoFronte... 🏁 08:48:01.494 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.573 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 Usage:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:48:01.574 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 
💀    🚀 prepareDemoFronte... 🏁 08:48:01.575 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:48:01.576 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.591 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.enabled":true,"service.ports":[],"service.type":"ClusterIP"}
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.591 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.667 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.668 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 08:48:01.669 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 08:48:02.749 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:48:02.849 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.enabled":true,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}],"service.type":"ClusterIP"}
💀    🚀 prepareDemoBacken... 🏁 08:48:02.849 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.909 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91      dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  for this case.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  Usage:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91    helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91    dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91  Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.91    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 Flags:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 
💀    🚀 prepareDemoBacken... 🏁 08:48:02.911 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 08:48:03.283 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 08:48:03.366 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.368  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.368  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.37   -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.373  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.375  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.375  
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376   - app: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376         ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376       - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.376               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                               - template: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.377                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.378                                                 ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                         ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                   - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.379                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                              }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                        - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                            - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                    - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                        - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                        - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                        - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                        - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                                  }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                              }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                                  }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                              }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                            - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                            - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                      ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                    - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                    - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                    - resource_version  : "190261"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                    - uid               : "6dad8f2b-760f-4dfa-992a-032c9469b7aa"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.38                  }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.381                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.382                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                             ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                   - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.383                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                     ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                   - resource_version  : "190262"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                   - uid               : "be52a82d-ad68-4c89-b2aa-f9f77ac29258"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384 Resources:
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.384 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.393 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 08:48:03.462 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.462  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.463  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.48  
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.481  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.483  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.485  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.486  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.487  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.487  
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.488 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.488   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.488       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.488       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.489                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                    - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                        - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                        -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                        - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                  ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                                - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                              }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                          ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                        - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                                  }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                              }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                          }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49  
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                    - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                    - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                    - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                        - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                    - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                    -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                            - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.49                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - resource_version  : "190277"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - uid               : "9b69bd60-2b49-4dc6-bb07-fabfb225fcea"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.491                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                             ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.492                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                   - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.493                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.494                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.494                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                   - resource_version  : "190278"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                   - uid               : "8ac06424-9384-48f6-b060-b6797074b3dc"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.495 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 08:48:03.566  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.571  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.575  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.576  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.581  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.588  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.588  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.588  
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589   - app: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589         ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589       - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.589                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                    - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                  }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                - template: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                    - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                            - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                            - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                    - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                        - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                  ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                                - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                              }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                          ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                        - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                                  }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                              }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                          }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59  
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                      }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                    - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                    - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.59                    - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                     ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - resource_version  : "190261"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - uid               : "6dad8f2b-760f-4dfa-992a-032c9469b7aa"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.591                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                             ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                     ]
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - resource_version  : "190262"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                   - uid               : "be52a82d-ad68-4c89-b2aa-f9f77ac29258"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592                 }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.592               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593             }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593         }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593     }
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 Resources:
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593     - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 08:48:03.593 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoFronte... 🏁 08:48:03.594 hello world
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.599 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.601  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.607  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.702  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.703  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.703  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.705  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.709  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715  
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.715                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.716                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.717                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718 
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                   - creation_timestamp: "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.718                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.719                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - time       : "2022-05-14T01:47:58Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - resource_version  : "190277"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - uid               : "9b69bd60-2b49-4dc6-bb07-fabfb225fcea"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.722                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.723                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.724                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:48:03.725                                   - image_pull_policy         : "IfNotPre
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 08:48:04.42  Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 08:48:04.508 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.508  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.508  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.509  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.512  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.515  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.518  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.518  
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.521       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.522                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.523                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.524                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.525                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                   - creation_timestamp: "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.526                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.527                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - time       : "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - resource_version  : "190306"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - uid               : "6c1da0a9-6c7b-4c81-aeb4-37701456f09a"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.528                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: "false"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                             ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                             ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                             ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                           - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.529                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                    -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - value: "600"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                          }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                    -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                          }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                    -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                          }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                    -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.53                                            - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.531                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.532                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.533                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                   -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.534                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                             ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.535               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.536                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   - creation_timestamp: "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                           - time       : "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.537                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - resource_version  : "190308"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - uid               : "935991b4-52b8-4e8b-976f-426a8521fdeb"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - cluster_ip             : "10.106.122.19"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   -     [0]: "10.106.122.19"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                     ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538               - status     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.538               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                   - creation_timestamp: "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.539                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                        - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                      }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                    - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                      }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                    - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                      }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                                  }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                              }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                            - time       : "2022-05-14T01:47:59Z"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                          }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                      ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    - resource_version  : "190305"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                    - uid               : "778ecfc7-6182-4f28-b88b-9954eeb1ab0c"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                  }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54                - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54              }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54          }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54        - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54      }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54  
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54  Resources:
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54      - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54  
💀    🚀 destroyDemoBacken... 🏁 08:48:04.54  Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 08:48:04.621 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.622  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.623  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.628  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.728  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.737  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.737  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.741  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.743  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.75   -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.76   -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.768  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.768  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.77   
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.771               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                             }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                 }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                         }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.772                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 08:48:04.773 
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.774                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.775                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.776                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:48:04.777              
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 08:48:04.881 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 4.55510389s
         Current Time: 08:48:04
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.666890719s
         Current Time: 08:48:05
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

