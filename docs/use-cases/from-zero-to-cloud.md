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

zaruba task setConfig index.zaruba.yaml startDemoFrontendContainer '{"localhost": "localhost"}'

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
         Elapsed Time: 1.22µs
         Current Time: 06:00:26
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 06:00:26.915 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 06:00:26.919 🎉🎉🎉
💀    🚀 initProject          🚧 06:00:26.919 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 112.922185ms
         Current Time: 06:00:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 313.928727ms
         Current Time: 06:00:27
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.178µs
         Current Time: 06:00:27
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:00:27.398 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:00:27.401 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:00:27.401 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:00:27.401 
💀    🚀 zrbShowAdv           ☕ 06:00:27.401         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:00:27.402         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:00:27.402     
💀    🚀 zrbShowAdv           ☕ 06:00:27.402 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:00:27.402 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:00:27.402   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:00:27.402   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:00:27.402   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:00:27.402 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 06:00:27.865 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 06:00:27.865 Preparing base variables
💀    🚀 makeMysqlApp         🐬 06:00:27.938 Base variables prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.938 Preparing start command
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Start command prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Preparing test command
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Test command prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Preparing check command
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Check command prepared
💀    🚀 makeMysqlApp         🐬 06:00:27.939 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 06:00:28.134 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 06:00:28.14  Add env to replacement map
💀    🚀 makeMysqlApp         🐬 06:00:28.146 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 06:00:28.146 ✅ Validate
💀    🚀 makeMysqlApp         🐬 06:00:28.146 Validate app directory
💀    🚀 makeMysqlApp         🐬 06:00:28.146 Done validating app directory
💀    🚀 makeMysqlApp         🐬 06:00:28.146 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 06:00:28.149 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 06:00:28.149 Validate template locations
💀    🚀 makeMysqlApp         🐬 06:00:28.157 Done validating template locations
💀    🚀 makeMysqlApp         🐬 06:00:28.157 Validate app ports
💀    🚀 makeMysqlApp         🐬 06:00:28.16  Done validating app ports
💀    🚀 makeMysqlApp         🐬 06:00:28.16  Validate app crud fields
💀    🚀 makeMysqlApp         🐬 06:00:28.162 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 06:00:28.162 🚧 Generate
💀    🚀 makeMysqlApp         🐬 06:00:28.162 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 06:00:28.162   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 06:00:28.162 ]
💀    🚀 makeMysqlApp         🐬 06:00:28.162 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 06:00:28.176 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 06:00:28.176 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 06:00:28.176 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.642 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.642 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.84  Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:28.841 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.122 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.128 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.134 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.134 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.135 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.135 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.135 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.138 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.138 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.155 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.155 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.159 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.159 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162 ]
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.162 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.204 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.207 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.21  Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.362 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.534 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.688 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.844 Checking start
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.849 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 06:00:29.997 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.16  Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.164 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.328 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.488 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.656 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.815 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:30.819 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.014 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.203 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.206 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.389 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.582 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.585 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.77  Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.957 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 06:00:31.96  Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 06:00:32.157 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 06:00:32.353 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 06:00:32.357 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 06:00:32.357 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.065085951s
         Current Time: 06:00:32
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.176860809s
         Current Time: 06:00:32
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.285µs
         Current Time: 06:00:32
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:00:32.766 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:00:32.77          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:00:32.77      
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:00:32.77    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:00:32.77    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:00:32.77    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:00:32.77  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 06:00:33.25  🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 06:00:33.25  Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.464 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 06:00:33.764 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 06:00:33.774 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 06:00:33.786 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 06:00:33.786 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 06:00:33.786 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 06:00:33.786 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 06:00:33.786 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 06:00:33.791 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 06:00:33.791 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 06:00:33.806 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 06:00:33.806 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 06:00:33.811 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 06:00:33.812 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 06:00:33.815 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 06:00:33.816 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 06:00:33.816 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 06:00:33.816   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 06:00:33.816 ]
💀    🚀 makeFastApiApp       ⚡ 06:00:33.816 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 06:00:34.554 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 06:00:34.555 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 06:00:34.555 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 06:00:35.079 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:00:35.079 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.404 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.405 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.643 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.65  Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.657 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.657 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.657 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.657 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.657 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.661 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.661 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.676 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.676 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.679 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.679 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 ]
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.682 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.725 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.729 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.732 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.889 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:00:36.893 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.066 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.265 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.269 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.462 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.693 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.697 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 06:00:37.884 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.082 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.086 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.283 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.469 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.473 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.665 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:38.857 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.053 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.269 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.274 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.467 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.648 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.652 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 06:00:39.86  Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.059 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.063 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.261 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.457 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.461 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.655 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 06:00:40.873 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 06:00:41.076 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 06:00:41.276 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:41.481 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 06:00:41.67  Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 06:00:41.868 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:42.067 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 06:00:42.266 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 06:00:42.266 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 06:00:42.832 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 06:00:42.832 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing start command
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Start command prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing test command
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Test command prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing check command
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Check command prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.114 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 06:00:44.468 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 06:00:44.476 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 06:00:44.487 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 06:00:44.487 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 06:00:44.487 Validate app directory
💀    🚀 addFastApiModule     ⚡ 06:00:44.487 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 06:00:44.487 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 06:00:44.49  Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 06:00:44.49  Validate template locations
💀    🚀 addFastApiModule     ⚡ 06:00:44.503 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 06:00:44.503 Validate app ports
💀    🚀 addFastApiModule     ⚡ 06:00:44.506 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 06:00:44.506 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 06:00:44.51  Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 06:00:44.51  🚧 Generate
💀    🚀 addFastApiModule     ⚡ 06:00:44.51  🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 06:00:44.51    "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 06:00:44.51  ]
💀    🚀 addFastApiModule     ⚡ 06:00:44.51  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 06:00:44.528 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 06:00:44.528 Registering module
💀    🚀 addFastApiModule     ⚡ 06:00:44.56  Done registering module
💀    🚀 addFastApiModule     ⚡ 06:00:44.561 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 06:00:44.561 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 06:00:44.91  🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 06:00:44.91  Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:45.818 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.105 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.112 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.12  Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.12  Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 06:00:46.133 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 06:00:46.133 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 06:00:46.212 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 06:00:46.212 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 06:00:46.293 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 06:00:46.293 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 06:00:46.406 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 06:00:46.406 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 06:00:46.468 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 06:00:46.468 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 06:00:46.468 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.468 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 06:00:46.468 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.469 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.703 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.709 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 06:00:46.714 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 06:00:46.715 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 06:00:46.715 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 06:00:46.715 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 06:00:46.715 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 06:00:46.717 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 06:00:46.717 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 06:00:46.726 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 06:00:46.726 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 06:00:46.728 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 06:00:46.728 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 ]
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 
💀    🚀 addFastApiCrud       ⚡ 06:00:46.731 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 06:00:46.758 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 06:00:46.758 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 06:00:46.791 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 06:00:46.791 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 06:00:46.836 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 06:00:46.836 Registering repo
💀    🚀 addFastApiCrud       ⚡ 06:00:46.887 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 06:00:46.888 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 06:00:46.888 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 14.230576617s
         Current Time: 06:00:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.431022177s
         Current Time: 06:00:47
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.195µs
         Current Time: 06:00:47
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:00:47.376 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:00:47.378         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:00:47.378     
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:00:47.378   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:00:47.378   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:00:47.378   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:00:47.378 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 06:00:47.828 🧰 Prepare
💀    🚀 makeNginxApp         📗 06:00:47.828 Preparing base variables
💀    🚀 makeNginxApp         📗 06:00:47.899 Base variables prepared
💀    🚀 makeNginxApp         📗 06:00:47.899 Preparing start command
💀    🚀 makeNginxApp         📗 06:00:47.899 Start command prepared
💀    🚀 makeNginxApp         📗 06:00:47.899 Preparing prepare command
💀    🚀 makeNginxApp         📗 06:00:47.899 Prepare command prepared
💀    🚀 makeNginxApp         📗 06:00:47.899 Preparing test command
💀    🚀 makeNginxApp         📗 06:00:47.899 Test command prepared
💀    🚀 makeNginxApp         📗 06:00:47.899 Preparing migrate command
💀    🚀 makeNginxApp         📗 06:00:47.899 Migrate command prepared
💀    🚀 makeNginxApp         📗 06:00:47.899 Preparing check command
💀    🚀 makeNginxApp         📗 06:00:47.9   Check command prepared
💀    🚀 makeNginxApp         📗 06:00:47.9   Preparing replacement map
💀    🚀 makeNginxApp         📗 06:00:48.093 Add config to replacement map
💀    🚀 makeNginxApp         📗 06:00:48.099 Add env to replacement map
💀    🚀 makeNginxApp         📗 06:00:48.106 Replacement map prepared
💀    🚀 makeNginxApp         📗 06:00:48.106 ✅ Validate
💀    🚀 makeNginxApp         📗 06:00:48.106 Validate app directory
💀    🚀 makeNginxApp         📗 06:00:48.106 Done validating app directory
💀    🚀 makeNginxApp         📗 06:00:48.106 Validate app container volumes
💀    🚀 makeNginxApp         📗 06:00:48.109 Done validating app container volumes
💀    🚀 makeNginxApp         📗 06:00:48.109 Validate template locations
💀    🚀 makeNginxApp         📗 06:00:48.117 Done validating template locations
💀    🚀 makeNginxApp         📗 06:00:48.117 Validate app ports
💀    🚀 makeNginxApp         📗 06:00:48.119 Done validating app ports
💀    🚀 makeNginxApp         📗 06:00:48.119 Validate app crud fields
💀    🚀 makeNginxApp         📗 06:00:48.122 Done validating app crud fields
💀    🚀 makeNginxApp         📗 06:00:48.122 🚧 Generate
💀    🚀 makeNginxApp         📗 06:00:48.122 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 06:00:48.122   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 06:00:48.122 ]
💀    🚀 makeNginxApp         📗 06:00:48.122 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 06:00:48.142 🔩 Integrate
💀    🚀 makeNginxApp         📗 06:00:48.142 🎉🎉🎉
💀    🚀 makeNginxApp         📗 06:00:48.142 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 06:00:48.541 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 06:00:48.541 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Preparing start command
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Start command prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Preparing test command
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Test command prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Preparing check command
💀    🚀 makeNginxAppRunner   📗 06:00:48.627 Check command prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.628 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 06:00:48.813 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 06:00:48.818 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 06:00:48.824 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 06:00:48.824 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 06:00:48.824 Validate app directory
💀    🚀 makeNginxAppRunner   📗 06:00:48.824 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 06:00:48.824 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 06:00:48.827 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 06:00:48.827 Validate template locations
💀    🚀 makeNginxAppRunner   📗 06:00:48.837 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 06:00:48.837 Validate app ports
💀    🚀 makeNginxAppRunner   📗 06:00:48.84  Done validating app ports
💀    🚀 makeNginxAppRunner   📗 06:00:48.84  Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 06:00:48.842 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 06:00:48.842 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 06:00:48.842 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 06:00:48.843   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 06:00:48.843   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 06:00:48.843 ]
💀    🚀 makeNginxAppRunner   📗 06:00:48.843 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 06:00:48.864 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 06:00:48.867 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 06:00:48.87  Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:00:49.015 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:00:49.201 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:00:49.402 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:00:49.6   Checking start
💀    🚀 makeNginxAppRunner   📗 06:00:49.604 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 06:00:49.821 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:00:50.012 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 06:00:50.017 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 06:00:50.234 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 06:00:50.429 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:00:50.627 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:00:50.811 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 06:00:50.814 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 06:00:51.001 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 06:00:51.193 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 06:00:51.196 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 06:00:51.379 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:00:51.559 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 06:00:51.563 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 06:00:51.734 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:00:51.916 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 06:00:51.921 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 06:00:52.116 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 06:00:52.264 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 06:00:52.267 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 06:00:52.267 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.997963274s
         Current Time: 06:00:52
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.109592113s
         Current Time: 06:00:52
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.383µs
         Current Time: 06:00:52
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:00:52.7   Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 06:00:52.808 Synchronize task environments
💀    🚀 syncEnv              🔄 06:00:53.032 Synchronize project's environment files
💀    🚀 syncEnv              🔄 06:00:53.24  🎉🎉🎉
💀    🚀 syncEnv              🔄 06:00:53.24  Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 647.132047ms
         Current Time: 06:00:53
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 848.113579ms
         Current Time: 06:00:53
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.556µs
         Current Time: 06:00:53
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:00:53.973 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:00:53.973 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoFrontend... 🏭 06:00:54.236 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 06:00:54.236 Build image demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 06:00:54.237 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 06:00:57.362 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 06:00:57.362 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 06:00:57.418 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 06:00:57.418  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 06:00:57.418 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 06:00:57.421 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 06:00:57.423 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 06:00:57.423 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 06:00:57.425 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 06:00:57.426 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoBackendI... 🏭 06:00:57.427 Sending build context to Docker daemon   1.03MB
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 06:00:57.432 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 06:00:57.433 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> 97fdfef7cb48
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> bf9c545afbe0
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439  ---> a62a483a9091
💀    🚀 buildDemoBackendI... 🏭 06:00:57.439 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 06:00:57.44   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:00:57.44   ---> db465fe79375
💀    🚀 buildDemoBackendI... 🏭 06:00:57.441 Successfully built db465fe79375
💀    🚀 buildDemoBackendI... 🏭 06:00:57.447 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 06:00:57.448 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 06:00:57.448 Docker image demo-backend built
💀    🚀 buildDemoFrontend... 🏭 06:00:57.51   ---> bf21cc72e86b
💀    🚀 buildDemoFrontend... 🏭 06:00:57.51  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🚀 buildDemoFrontend... 🏭 06:00:57.583  ---> 4ac66cb44f82
💀    🚀 buildDemoFrontend... 🏭 06:00:57.583 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:00:57.604  ---> Running in f0e2268df083
💀    🚀 buildDemoFrontend... 🏭 06:00:57.696 Removing intermediate container f0e2268df083
💀    🚀 buildDemoFrontend... 🏭 06:00:57.696  ---> 8314179b044e
💀    🚀 buildDemoFrontend... 🏭 06:00:57.696 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 06:00:57.791  ---> 822e80eea4b0
💀    🚀 buildDemoFrontend... 🏭 06:00:57.791 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 06:00:57.811  ---> Running in 821b5a142915
💀    🚀 buildDemoFrontend... 🏭 06:00:58.611 Removing intermediate container 821b5a142915
💀    🚀 buildDemoFrontend... 🏭 06:00:58.611  ---> ca885dd12698
💀    🚀 buildDemoFrontend... 🏭 06:00:58.611 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:00:58.634  ---> Running in 2762365670ac
💀    🚀 buildDemoFrontend... 🏭 06:00:58.699 Removing intermediate container 2762365670ac
💀    🚀 buildDemoFrontend... 🏭 06:00:58.699  ---> 1739ea4e7fbe
💀    🚀 buildDemoFrontend... 🏭 06:00:58.7   Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 06:00:58.722  ---> Running in 0645996351b3
💀    🚀 buildDemoFrontend... 🏭 06:00:58.79  Removing intermediate container 0645996351b3
💀    🚀 buildDemoFrontend... 🏭 06:00:58.79   ---> 0130b4a39d50
💀    🚀 buildDemoFrontend... 🏭 06:00:58.791 Successfully built 0130b4a39d50
💀    🚀 buildDemoFrontend... 🏭 06:00:58.797 Successfully tagged demo-frontend:latest
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 06:00:58.905 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 5.041036372s
         Current Time: 06:00:59
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.152575843s
         Current Time: 06:00:59
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.074µs
         Current Time: 06:00:59
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:00:59.276 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:00:59.276 Links updated
💀    🚀 prepareDemoBackend   🔧 06:00:59.277 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 06:00:59.301 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 06:00:59.391 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 06:00:59.534 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 06:01:00.042 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 06:01:00.043 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 06:01:00.087 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> bf21cc72e86b
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> 4ac66cb44f82
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089  ---> 8314179b044e
💀    🚀 buildDemoFrontend... 🏭 06:01:00.089 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoDbImage     🏭 06:01:00.089 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 06:01:00.09   ---> 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 06:01:00.09   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.09   ---> 822e80eea4b0
💀    🚀 buildDemoFrontend... 🏭 06:01:00.09  Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoDbImage     🏭 06:01:00.09  Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 06:01:00.092 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092  ---> ca885dd12698
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092  ---> 1739ea4e7fbe
💀    🚀 buildDemoFrontend... 🏭 06:01:00.092 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 06:01:00.093  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:00.093  ---> 0130b4a39d50
💀    🚀 buildDemoFrontend... 🏭 06:01:00.093 Successfully built 0130b4a39d50
💀    🚀 buildDemoDbImage     🏭 06:01:00.094 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 06:01:00.094 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 06:01:00.095 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 06:01:00.097 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 06:01:00.097 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 06:01:00.418 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🔥 🚀 startDemoFrontend... 📗 06:01:00.447 Error: No such container: demoFrontend
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🔥 🔎 startDemoFrontend... 📗 06:01:00.448 Error: No such container: demoFrontend
💀    🔎 startDemoDbContainer 🐬 06:01:00.449 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🚀 startDemoFrontend... 📗 06:01:00.473 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 06:01:00.476 🐳 Creating and starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:00.478 Error: No such container: demoDb
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:00.491 Error: No such container: demoDb
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:00.506 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 06:01:00.507 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoFrontend... 📗 06:01:00.533 6fcd27d21801f30dfc06f61805c79bc06ff815e030275488788f36d01a33a61f
💀    🚀 startDemoDbContainer 🐬 06:01:00.555 7e8a3b699ee0ed79d91080a4fe1db6d369346c56cff66b1254eacea9a161493e
💀    🚀 prepareDemoBackend   🔧 06:01:00.946 Activate venv
💀    🚀 prepareDemoBackend   🔧 06:01:00.946 Install dependencies
💀    🚀 prepareDemoBackend   🔧 06:01:01.209 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 06:01:01.426   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:01.433 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 06:01:01.555   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:01.562 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 06:01:01.635   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoFrontend... 📗 06:01:02.533 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 06:01:02.545 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoDbContainer 🐬 06:01:02.554 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 06:01:02.554 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.53 
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.53 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.53 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀    🚀 startDemoFrontend... 📗 06:01:02.568 
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.53 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.54 
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 [38;5;6mnginx [38;5;5m23:01:02.55 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 2022/05/11 23:01:02 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 06:01:02.568 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🔎 startDemoFrontend... 📗 06:01:02.579 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 06:01:02.579 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 06:01:02.58  🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 06:01:02.58  🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 06:01:02.582 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.53 
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.53 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.54 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.54 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.54 
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.54 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.56 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.56 [38;5;2mINFO  ==> Initializing mysql database
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.58 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.588 [38;5;6mmysql [38;5;5m23:01:02.58 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.594 [38;5;6mmysql [38;5;5m23:01:02.59 [38;5;2mINFO  ==> Setting slow_query_log option
💀    🔎 startDemoDbContainer 🐬 06:01:02.595 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 06:01:02.595 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 06:01:02.596 🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.599 [38;5;6mmysql [38;5;5m23:01:02.59 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:02.604 [38;5;6mmysql [38;5;5m23:01:02.60 [38;5;2mINFO  ==> Installing database
💀    🚀 prepareDemoBackend   🔧 06:01:03.54  Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 06:01:03.648   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:03.656 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 06:01:03.719   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:03.728 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 06:01:03.827   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:03.845 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 06:01:03.956   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:03.966 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 06:01:04.124   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:04.201 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 06:01:04.54    Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:04.614 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 06:01:04.77    Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:04.857 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 06:01:05.229   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:01:05.416 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 06:01:05.576   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:05.586 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🔎 startDemoFrontend... 📗 06:01:05.586 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 06:01:05.599 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🚀 prepareDemoBackend   🔧 06:01:05.681   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:05.688 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🔎 startDemoFrontend... 📗 06:01:05.709 check demoFrontend
💀    🔎 startDemoFrontend... 📗 06:01:05.711 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:05.723 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:05.724 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 06:01:05.745   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:05.75  Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 06:01:05.835   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:05.845 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀    🚀 prepareDemoBackend   🔧 06:01:05.922   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:05.946 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 06:01:06.017   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.027 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 06:01:06.225   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.247 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 06:01:06.387   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.397 Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:06.555   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.589 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 06:01:06.67    Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.684 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:01:06.754   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:06.767 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 06:01:06.853   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:01:07.073 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 06:01:07.213   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:07.234 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:07.481 [38;5;6mmysql [38;5;5m23:01:07.48 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 06:01:07.723   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:07.844 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 06:01:07.95    Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:07.965 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 06:01:08.077   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.082 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧 06:01:08.165   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.177 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 06:01:08.284   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.298 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 06:01:08.347   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 06:01:08.45  Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧 06:01:08.551   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.566 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 06:01:08.77    Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.783 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:08.852 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:08.854 ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: YES)
💀    🚀 prepareDemoBackend   🔧 06:01:08.903   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.907 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:08.976   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:08.983 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:09.068   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.078 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:09.16    Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.188 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:09.263   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.268 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:09.332   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.337 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:09.433   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.441 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:09.494 [38;5;6mmysql [38;5;5m23:01:09.49 [38;5;2mINFO  ==> Configuring authentication
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:09.536 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:09.555 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:09.592 [38;5;6mmysql [38;5;5m23:01:09.59 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:09.595 [38;5;6mmysql [38;5;5m23:01:09.59 [38;5;2mINFO  ==> Stopping mysql
💀    🔎 startDemoFrontend... 📗 06:01:09.713 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 06:01:09.713 📜 Task 'startDemoFrontendContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 06:01:09.818   Using cached https://files.pythonhosted.org/packages/0c/58/25b4d208e0f6f00e19440385f360dc9891f8fa5ab62c11da52eb226fd9cd/coverage-6.3.2-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.829 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:01:09.94    Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:09.972 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:01:10.063   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:10.069 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 06:01:10.152   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:10.16  Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 06:01:10.236   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:10.243 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 06:01:10.363   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:10.373 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 06:01:10.515   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 06:01:10.754 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, fastavro, idna, urllib3, requests, confluent-kafka, cryptography, starlette, typing-extensions, pydantic, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, toml, pluggy, attrs, pyparsing, packaging, iniconfig, py, pytest, tomli, coverage, pytest-cov, pyasn1, rsa, ecdsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 06:01:10.772   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 06:01:10.944     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:01:11.152   Running setup.py install for fastavro: started
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:11.61  [38;5;6mmysql [38;5;5m23:01:11.60 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:12.032 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:12.033 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:15.142 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:15.143 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:17.62  [38;5;6mmysql [38;5;5m23:01:17.61 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:17.632 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:17.636 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:17.64  [38;5;6mmysql [38;5;5m23:01:17.63 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:18.291 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:18.293 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 startDemoDbContainer 🐬 06:01:19.654 
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:19.654 [38;5;6mmysql [38;5;5m23:01:19.65 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 06:01:19.67  [38;5;6mmysql [38;5;5m23:01:19.67 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 06:01:19.855 2022-05-11T23:01:19.850755Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 06:01:19.855 2022-05-11T23:01:19.851921Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 06:01:19.855 2022-05-11T23:01:19.851927Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 06:01:19.856 2022-05-11T23:01:19.855916Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 06:01:19.96  2022-05-11T23:01:19.960299Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 06:01:20.063 2022-05-11T23:01:20.062587Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 06:01:20.063 2022-05-11T23:01:20.062632Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 06:01:20.076 2022-05-11T23:01:20.075328Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 06:01:20.076 2022-05-11T23:01:20.075406Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:21.414 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 Database
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 information_schema
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 mysql
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 performance_schema
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 sample
💀    🔎 startDemoDbContainer 🐬 06:01:21.421 sys
💀    🔎 startDemoDbContainer 🐬 06:01:21.425 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 06:01:25.427 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 06:01:25.428 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 06:01:32.285     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:01:33.46    Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 06:01:33.649     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:01:34.168   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 06:01:34.322     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 06:01:34.382 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.2 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 06:01:34.445 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 06:01:34.445 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 06:01:34.492 Prepare
💀    🚀 prepareDemoBackend   🔧 06:01:34.492 prepare command
💀    🚀 prepareDemoBackend   🔧 06:01:34.492 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 startDemoBackend     ⚡ 06:01:34.751 Activate venv
💀    🔎 startDemoBackend     ⚡ 06:01:34.752 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 06:01:34.752 Start
💀    🚀 startDemoBackend     ⚡ 06:01:35.103 2022-05-12 06:01:35,103 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 06:01:35.104 2022-05-12 06:01:35,103 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.108 2022-05-12 06:01:35,108 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 06:01:35.108 2022-05-12 06:01:35,108 INFO sqlalchemy.engine.Engine [generated in 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.11  2022-05-12 06:01:35,110 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 06:01:35.11  2022-05-12 06:01:35,110 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.111 2022-05-12 06:01:35,111 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.112 2022-05-12 06:01:35,112 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:01:35.112 2022-05-12 06:01:35,112 INFO sqlalchemy.engine.Engine [generated in 0.00011s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 2022-05-12 06:01:35,114 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 )
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 
💀    🚀 startDemoBackend     ⚡ 06:01:35.115 2022-05-12 06:01:35,115 INFO sqlalchemy.engine.Engine [no key 0.00008s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.134 2022-05-12 06:01:35,134 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.134 2022-05-12 06:01:35,134 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.151 2022-05-12 06:01:35,151 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 06:01:35.151 2022-05-12 06:01:35,151 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.167 2022-05-12 06:01:35,167 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 06:01:35.167 2022-05-12 06:01:35,167 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.183 2022-05-12 06:01:35,182 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 06:01:35.183 2022-05-12 06:01:35,183 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.198 2022-05-12 06:01:35,198 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:01:35.199 2022-05-12 06:01:35,199 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.2   2022-05-12 06:01:35,199 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:01:35.2   2022-05-12 06:01:35,200 INFO sqlalchemy.engine.Engine [cached since 0.08804s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 2022-05-12 06:01:35,201 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 )
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 
💀    🚀 startDemoBackend     ⚡ 06:01:35.201 2022-05-12 06:01:35,201 INFO sqlalchemy.engine.Engine [no key 0.00010s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.221 2022-05-12 06:01:35,221 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.221 2022-05-12 06:01:35,221 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.238 2022-05-12 06:01:35,238 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 06:01:35.238 2022-05-12 06:01:35,238 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.257 2022-05-12 06:01:35,257 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 06:01:35.257 2022-05-12 06:01:35,257 INFO sqlalchemy.engine.Engine [no key 0.00020s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.274 2022-05-12 06:01:35,274 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:01:35.275 2022-05-12 06:01:35,275 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.275 2022-05-12 06:01:35,275 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 06:01:35.275 2022-05-12 06:01:35,275 INFO sqlalchemy.engine.Engine [cached since 0.1638s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 2022-05-12 06:01:35,277 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 06:01:35.277 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 )
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 
💀    🚀 startDemoBackend     ⚡ 06:01:35.278 2022-05-12 06:01:35,277 INFO sqlalchemy.engine.Engine [no key 0.00009s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.298 2022-05-12 06:01:35,298 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 06:01:35.298 2022-05-12 06:01:35,298 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.312 2022-05-12 06:01:35,312 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 06:01:35.312 2022-05-12 06:01:35,312 INFO sqlalchemy.engine.Engine [no key 0.00012s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.333 2022-05-12 06:01:35,333 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 06:01:35.333 2022-05-12 06:01:35,333 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.348 2022-05-12 06:01:35,348 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 06:01:35.348 2022-05-12 06:01:35,348 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.364 2022-05-12 06:01:35,364 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 06:01:35.364 2022-05-12 06:01:35,364 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.381 2022-05-12 06:01:35,381 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 06:01:35.381 2022-05-12 06:01:35,381 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 06:01:35.399 2022-05-12 06:01:35,399 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:01:35.401 2022-05-12 06:01:35,401 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.402 2022-05-12 06:01:35,402 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 06:01:35.402 FROM users 
💀    🚀 startDemoBackend     ⚡ 06:01:35.403 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 06:01:35.403  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 06:01:35.403 2022-05-12 06:01:35,402 INFO sqlalchemy.engine.Engine [generated in 0.00012s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 06:01:35.404 2022-05-12 06:01:35,404 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 06:01:35.587 2022-05-12 06:01:35,587 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.588 2022-05-12 06:01:35,588 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 06:01:35.588 2022-05-12 06:01:35,588 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {'id': 'd7a48640-fc5e-4c65-a2a6-9233427d376b', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$H0VSZ/tVEKA7UytyYxWHxuHbHrDvwZI8azgElvockwJvGWmEl4yKG', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 12, 6, 1, 35, 586884), 'updated_at': datetime.datetime(2022, 5, 12, 6, 1, 35, 588393)}
💀    🚀 startDemoBackend     ⚡ 06:01:35.59  2022-05-12 06:01:35,589 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 06:01:35.598 2022-05-12 06:01:35,597 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 06:01:35.598 2022-05-12 06:01:35,598 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 06:01:35.599 FROM users 
💀    🚀 startDemoBackend     ⚡ 06:01:35.599 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 06:01:35.599 2022-05-12 06:01:35,598 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {'pk_1': 'd7a48640-fc5e-4c65-a2a6-9233427d376b'}
💀    🚀 startDemoBackend     ⚡ 06:01:35.6   2022-05-12 06:01:35,600 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 06:01:35.602 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.61  Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 06:01:35.618 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.623 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 06:01:35.623 Register library route handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.623 Register library event handler
💀    🚀 startDemoBackend     ⚡ 06:01:35.623 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 06:01:35.623 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:35.624 INFO:     Started server process [26365]
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:35.624 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:35.624 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:35.624 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 06:01:35.755 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 06:01:35.755 check demoBackend
💀    🔎 startDemoBackend     ⚡ 06:01:35.755 🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 06:01:35.755 📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 06:01:35.863 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 36.695477077s
         Current Time: 06:01:35
         Active Process:
           * (PID=25565) 🐬 'startDemoDbContainer' service
           * (PID=25534) 📗 'startDemoFrontendContainer' service
           * (PID=26361) ⚡ 'startDemoBackend' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=26361)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=25565)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=25534)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:37.533 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:37.634 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:37.634 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 06:01:37.634 INFO:     Finished server process [26365]
💀    🚀 startDemoBackend     ⚡ 06:01:37.704 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 06:01:37.704 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 38.800737527s
         Current Time: 06:01:38
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.358µs
         Current Time: 06:01:38
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:01:38.306 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:01:38.306 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 06:01:38.35  🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 06:01:38.421 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 06:01:38.574 Build image demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 06:01:38.574 Build image demo-backend:latest
💀    🚀 buildDemoFrontend... 🏭 06:01:39.227 Sending build context to Docker daemon  16.38kB
💀    🚀 buildDemoDbImage     🏭 06:01:39.227 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoDbImage     🏭 06:01:39.282 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 06:01:39.282 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 06:01:39.282  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 06:01:39.282 Step 2/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 06:01:39.282  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 06:01:39.282 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 06:01:39.286 Sending build context to Docker daemon   1.18MB
💀    🚀 buildDemoDbImage     🏭 06:01:39.287 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 06:01:39.287  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.287  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 06:01:39.287 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289  ---> Using cache
💀    🚀 buildDemoDbImage     🏭 06:01:39.289 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289 Step 4/11 : USER 1001
💀    🚀 buildDemoDbImage     🏭 06:01:39.289 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 06:01:39.289 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> e64417d5ed16
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> bd888e5b5a62
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29  Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> d3adc6868f86
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29  Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29   ---> 9a5ca54187f1
💀    🚀 buildDemoFrontend... 🏭 06:01:39.29  Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> fd772be255c8
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> 863cd50f5bf7
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 06:01:39.291  ---> b1cc733d54af
💀    🚀 buildDemoFrontend... 🏭 06:01:39.292 Successfully built b1cc733d54af
💀    🚀 buildDemoBackendI... 🏭 06:01:39.294 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoFrontend... 🏭 06:01:39.297 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 06:01:39.297  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 06:01:39.297 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 06:01:39.298 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 06:01:39.299 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 06:01:39.3   🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 06:01:39.3   Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 06:01:39.432  ---> 0fa7d4234ea2
💀    🚀 buildDemoBackendI... 🏭 06:01:39.432 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 06:01:39.454  ---> Running in e0c861059561
💀    🚀 buildDemoBackendI... 🏭 06:01:39.519 Removing intermediate container e0c861059561
💀    🚀 buildDemoBackendI... 🏭 06:01:39.519  ---> b61c6e58d9e3
💀    🚀 buildDemoBackendI... 🏭 06:01:39.522 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 06:01:39.55   ---> Running in 6170074a9b74
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 06:01:39.576 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 06:01:39.605 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 06:01:39.611 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 06:01:39.611 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 06:01:39.617 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🔎 startDemoDbContainer 🐬 06:01:39.638 🔎 Waiting docker container 'demoDb' healthcheck
💀    🚀 startDemoDbContainer 🐬 06:01:39.64  🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 06:01:39.64  🐳 Logging 'demoDb'
💀    🔎 startDemoFrontend... 📗 06:01:39.651 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 06:01:39.651 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 06:01:39.653 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 06:01:39.653 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 06:01:39.655 🔎 Host port '443' is ready
💀    🔎 startDemoDbContainer 🐬 06:01:39.672 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 06:01:39.672 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 06:01:39.673 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 06:01:40.257 Removing intermediate container 6170074a9b74
💀    🚀 buildDemoBackendI... 🏭 06:01:40.257  ---> 395f4a5d3be8
💀    🚀 buildDemoBackendI... 🏭 06:01:40.257 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 06:01:40.279  ---> Running in f686c0aeac4a
💀    🚀 buildDemoBackendI... 🏭 06:01:40.33  Removing intermediate container f686c0aeac4a
💀    🚀 buildDemoBackendI... 🏭 06:01:40.331  ---> 643386eb8a8e
💀    🚀 buildDemoBackendI... 🏭 06:01:40.333 Successfully built 643386eb8a8e
💀    🚀 buildDemoBackendI... 🏭 06:01:40.346 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 06:01:40.348 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 06:01:40.348 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 06:01:42.66  🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 06:01:42.679 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 06:01:42.811 check demoFrontend
💀    🔎 startDemoFrontend... 📗 06:01:42.815 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 06:01:42.83  mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 Database
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 information_schema
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 mysql
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 performance_schema
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 sample
💀    🔎 startDemoDbContainer 🐬 06:01:42.832 sys
💀    🔎 startDemoDbContainer 🐬 06:01:42.835 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 06:01:46.818 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 06:01:46.818 📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 06:01:46.839 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 06:01:46.839 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 06:01:47.348 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:47.373 Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ 06:01:47.376 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:47.395 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 06:01:47.396 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 06:01:47.44  f4451fbf64bcd23feb002127f176b911105c066b74a50bdd1fded08c0a2fd09a
💀    🚀 startDemoBackendC... ⚡ 06:01:48.468 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:01:48.47  🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 06:01:48.498 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 06:01:48.498 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 06:01:48.5   🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 06:01:48.88  2022-05-11 23:01:48,879 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 06:01:48.88  2022-05-11 23:01:48,879 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.884 2022-05-11 23:01:48,883 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 06:01:48.884 2022-05-11 23:01:48,883 INFO sqlalchemy.engine.Engine [generated in 0.00028s] {}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.889 2022-05-11 23:01:48,888 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 06:01:48.889 2022-05-11 23:01:48,888 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.893 2022-05-11 23:01:48,892 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:01:48.893 2022-05-11 23:01:48,893 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:01:48.893 2022-05-11 23:01:48,893 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.896 2022-05-11 23:01:48,895 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:01:48.899 2022-05-11 23:01:48,899 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:01:48.9   2022-05-11 23:01:48,899 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:01:48.9   2022-05-11 23:01:48,899 INFO sqlalchemy.engine.Engine [cached since 0.006386s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.902 2022-05-11 23:01:48,901 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:01:48.904 2022-05-11 23:01:48,904 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:01:48.905 2022-05-11 23:01:48,904 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 06:01:48.905 2022-05-11 23:01:48,904 INFO sqlalchemy.engine.Engine [cached since 0.01168s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.907 2022-05-11 23:01:48,907 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 06:01:48.912 2022-05-11 23:01:48,912 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 06:01:48.914 2022-05-11 23:01:48,914 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 06:01:48.914 FROM users 
💀    🚀 startDemoBackendC... ⚡ 06:01:48.914 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 06:01:48.915  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 06:01:48.915 2022-05-11 23:01:48,914 INFO sqlalchemy.engine.Engine [generated in 0.00017s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 06:01:48.917 2022-05-11 23:01:48,917 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 06:01:48.92  Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.93  Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 06:01:48.94  Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.947 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 06:01:48.947 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.947 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 06:01:48.947 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 06:01:48.947 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:48.947 INFO:     Started server process [9]
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:48.948 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:48.948 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 06:01:48.948 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 06:01:51.503 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:01:51.633 check demoBackend
💀    🔎 startDemoBackendC... ⚡ 06:01:51.636 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 06:01:52.637 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 06:01:52.637 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 06:01:52.745 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 14.547784636s
         Current Time: 06:01:52
         Active Process:
           * (PID=27398) 📗 'startDemoFrontendContainer' service
           * (PID=27543) ⚡ 'startDemoBackendContainer' service
           * (PID=27430) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=27398)
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=27543)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=27430)
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 16.652045339s
         Current Time: 06:01:54
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.191µs
         Current Time: 06:01:55
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:01:55.113 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:01:55.113 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoDbContainer  ✋ 06:01:55.463 Stop docker container demoDb
💀    🚀 stopDemoBackendCo... ✋ 06:01:55.463 Stop docker container demoBackend
💀    🚀 stopDemoFrontendC... ✋ 06:01:55.465 Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ 06:01:59.184 demoDb
💀    🚀 stopDemoDbContainer  ✋ 06:01:59.185 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 06:01:59.185 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ 06:02:06.044 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 06:02:06.046 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 06:02:06.046 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀    🚀 stopDemoFrontendC... ✋ 06:02:06.542 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 06:02:06.544 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 06:02:06.544 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 06:02:06.653 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.646318446s
         Current Time: 06:02:06
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.758389803s
         Current Time: 06:02:06
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.15µs
         Current Time: 06:02:07
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 06:02:07.035 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 06:02:07.035 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🔥 🚀 removeDemoDbConta... ❌ 06:02:07.347 Error: No such container: 
💀 🔥 🚀 removeDemoFronten... ❌ 06:02:07.349 Error: No such container: 
💀    🚀 removeDemoDbConta... ❌ 06:02:07.349 Stop docker container demoDb
💀    🚀 removeDemoFronten... ❌ 06:02:07.351 Stop docker container demoFrontend
💀 🔥 🚀 removeDemoBackend... ❌ 06:02:07.352 Error: No such container: 
💀    🚀 removeDemoBackend... ❌ 06:02:07.354 Stop docker container demoBackend
💀    🚀 removeDemoDbConta... ❌ 06:02:07.447 Docker container demoDb stopped
💀    🚀 removeDemoDbConta... ❌ 06:02:07.447 Remove docker container demoDb
💀    🚀 removeDemoFronten... ❌ 06:02:07.452 Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ 06:02:07.452 Remove docker container demoFrontend
💀    🚀 removeDemoBackend... ❌ 06:02:07.454 Docker container demoBackend stopped
💀    🚀 removeDemoBackend... ❌ 06:02:07.454 Remove docker container demoBackend
💀    🚀 removeDemoFronten... ❌ 06:02:07.519 demoFrontend
💀    🚀 removeDemoBackend... ❌ 06:02:07.519 demoBackend
💀    🚀 removeDemoFronten... ❌ 06:02:07.521 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ 06:02:07.521 Docker container demoFrontend removed
💀    🚀 removeDemoBackend... ❌ 06:02:07.521 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ 06:02:07.521 Docker container demoBackend removed
💀    🚀 removeDemoDbConta... ❌ 06:02:07.529 demoDb
💀    🚀 removeDemoDbConta... ❌ 06:02:07.531 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ 06:02:07.531 Docker container demoDb removed
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 removeContainers     ❌ 06:02:07.638 
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 709.873476ms
         Current Time: 06:02:07
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 821.047304ms
         Current Time: 06:02:07
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 4.961µs
         Current Time: 06:02:08
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:08.049 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:02:08.052         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:02:08.052     
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:02:08.052   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:02:08.052   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:02:08.052   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:02:08.052 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.481 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.481 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.635 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.635 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.635 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.635 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.635 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.636 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.861 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.867 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.872 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.872 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.872 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.873 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.873 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.875 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.875 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.884 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.884 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.888 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.888 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.891 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.927 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.927 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:08.927 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.355 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.355 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.634 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.635 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.97  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.984 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.994 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.994 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.994 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.994 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.994 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.997 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:09.997 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.006 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.006 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.011 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.011 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.014 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.015 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.015 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.015   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.015 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.015 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.035 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.038 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.042 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.195 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.359 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.362 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.525 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.674 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.677 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.824 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.824 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:10.824 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.884418332s
         Current Time: 06:02:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.995685947s
         Current Time: 06:02:11
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.051µs
         Current Time: 06:02:11
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:11.187 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:02:11.191         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:02:11.191     
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:02:11.191   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:02:11.191   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:02:11.191   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:02:11.191 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:11.609 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:11.609 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.345 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.55  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.556 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.562 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.562 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.562 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.562 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.562 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.565 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.565 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.574 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.574 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.576 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.576 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.579 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.61  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.611 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:12.611 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.033 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.034 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.815 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:13.816 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.024 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.03  Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.036 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.036 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.036 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.036 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.036 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.038 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.038 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.046 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.046 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.049 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.049 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.052 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.07  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.073 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.076 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.223 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.369 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.372 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.522 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.669 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.672 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.82  Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.82  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:14.82  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.739395485s
         Current Time: 06:02:14
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.850338435s
         Current Time: 06:02:15
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.149µs
         Current Time: 06:02:15
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:15.193 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 06:02:15.197         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 06:02:15.197     
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 06:02:15.197   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 06:02:15.197   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 06:02:15.197   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 06:02:15.197 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.615 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.615 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.715 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.715 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.715 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.715 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.716 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.917 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.923 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.928 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.928 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.928 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.928 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.928 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.931 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.931 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.939 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.939 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.942 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.942 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.945 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.975 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.975 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:15.975 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.328 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.328 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.468 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.468 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.469 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.74  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.748 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.757 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.757 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.757 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.757 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.757 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.76  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.76  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.772 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.772 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.777 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.777 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782 ]
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.782 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.808 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.813 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.817 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:16.985 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.138 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.141 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.291 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.44  Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.442 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.592 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.592 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 06:02:17.592 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.508007425s
         Current Time: 06:02:17
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.619516846s
         Current Time: 06:02:17
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.119µs
         Current Time: 06:02:17
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:17.962 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 06:02:18.067 Synchronize task environments
💀    🚀 syncEnv              🔄 06:02:18.237 Synchronize project's environment files
💀    🚀 syncEnv              🔄 06:02:18.404 🎉🎉🎉
💀    🚀 syncEnv              🔄 06:02:18.404 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 550.572992ms
         Current Time: 06:02:18
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 751.743509ms
         Current Time: 06:02:18
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.541µs
         Current Time: 06:02:18
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:18.913 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 06:02:19.025 🎉🎉🎉
💀    🚀 setProjectValue      🔗 06:02:19.025 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 219.398776ms
         Current Time: 06:02:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 421.48525ms
         Current Time: 06:02:19
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.072µs
         Current Time: 06:02:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 06:02:19.491 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 06:02:19.602 🎉🎉🎉
💀    🚀 setProjectValue      🔗 06:02:19.602 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.895419ms
         Current Time: 06:02:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.152038ms
         Current Time: 06:02:19
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.219µs
         Current Time: 06:02:20
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 06:02:20.079 🚧 Create virtual environment.
💀    🚀 prepareDemoFronte... 🏁 06:02:20.079 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 06:02:20.082 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 06:02:21.75  🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:21.97  🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 06:02:22.037 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:22.042 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:22.211 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:22.249 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:22.761   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:22.775 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:22.95    Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:22.972 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:22.989   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:22.99    Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 06:02:23.008 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.164   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 06:02:23.178   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 06:02:23.234 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.378 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:23.395 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:23.484   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:23.488 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:23.489   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:23.497 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.568   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.585 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:23.591   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:23.596 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.716   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.72  Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:23.743   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:23.753   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:23.762 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:23.772 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.809   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.813 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.888   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:23.896 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:23.902   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:23.909 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:24.02    Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:24.037 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:24.124   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:24.131 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:24.154   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:24.182 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:24.219   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:24.238 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 06:02:24.256   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:24.261 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 06:02:24.934   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.015 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.094   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.107 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.123   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.144   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.204 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.209   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.219 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.224 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.285   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.296 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.316   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.327 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.385   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.4   Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.407   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.425   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.43  Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.438 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.495   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.507 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.511   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.522 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.528   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.541 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.602   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.608   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.62  Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.66  Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.66    Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.721 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.733   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.734   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 06:02:25.747 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:25.749 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.81    Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.823 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.837   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.849   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:25.858 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.901   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:25.909 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 06:02:25.937 Installing collected packages: six, protobuf, semver, pyyaml, dill, grpcio, pulumi, arpeggio, attrs, parver, charset-normalizer, certifi, idna, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 06:02:25.98    Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:26.01  Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.034   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.042 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 06:02:26.093   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 06:02:26.196 Installing collected packages: dill, six, pyyaml, protobuf, semver, grpcio, pulumi, arpeggio, attrs, parver, idna, charset-normalizer, urllib3, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.198   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.222 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.375   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 06:02:26.508 Installing collected packages: pyyaml, six, semver, dill, protobuf, grpcio, pulumi, attrs, arpeggio, parver, idna, certifi, urllib3, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 06:02:26.7     Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 06:02:26.921   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 06:02:27.286   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 06:02:28.273     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 06:02:28.325 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoBacken... 🏁 06:02:28.355 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 06:02:28.355 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.536     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 06:02:28.584 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 06:02:28.621 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 06:02:28.621 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.787 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 06:02:28.788 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:28.793     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 06:02:28.835 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 06:02:28.837 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.837 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 06:02:28.838 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 for this case.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 Usage:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 Aliases:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839 Flags:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.839   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84  
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84  
💀    🚀 prepareDemoFronte... 🏁 06:02:28.84  Use "helm dependency [command] --help" for more information about a command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:02:28.864 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:02:28.865 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 06:02:29.16  error: no stack named 'dev' found
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.176 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.177 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.261       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.262 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:29.263 🚧 Preparation completed.
💀    🚀 deployDemoFronten... 🏁 06:02:29.335 Created stack 'dev'
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoDbDeplo... 🏁 06:02:29.668 error: no stack named 'dev' found
💀    🚀 prepareDemoBacken... 🏁 06:02:29.782 PARTS: ["3000"]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:29.826 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 06:02:29.851 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 06:02:29.851 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.901 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.901 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.901 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.901 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 for this case.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 Usage:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 Aliases:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.902   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 Flags:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 
💀    🚀 prepareDemoBacken... 🏁 06:02:29.903 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 06:02:29.904 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 06:02:30.154 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 06:02:30.257 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 06:02:31.355 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 06:02:31.367 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 06:02:31.731 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:31.731 
💀    🚀 deployDemoBackend... 🏁 06:02:32.075 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 06:02:32.129  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.135  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 06:02:32.197  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.2    +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.411  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.414  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 06:02:32.427  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 06:02:32.433  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoBackend... 🏁 06:02:32.541 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662  
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:32.662 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 06:02:32.692  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 06:02:32.692  
💀    🚀 deployDemoFronten... 🏁 06:02:32.692 Resources:
💀    🚀 deployDemoFronten... 🏁 06:02:32.692     + 4 to create
💀    🚀 deployDemoFronten... 🏁 06:02:32.692 
💀    🚀 deployDemoFronten... 🏁 06:02:32.692 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 06:02:33.057  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 06:02:33.135  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.168 
💀    🚀 deployDemoFronten... 🏁 06:02:33.181 
💀    🚀 deployDemoBackend... 🏁 06:02:33.372  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:02:33.373  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:02:33.38   +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 06:02:33.512  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 06:02:33.512  
💀    🚀 deployDemoBackend... 🏁 06:02:33.512 Resources:
💀    🚀 deployDemoBackend... 🏁 06:02:33.512     + 5 to create
💀    🚀 deployDemoBackend... 🏁 06:02:33.512 
💀    🚀 deployDemoBackend... 🏁 06:02:33.512 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.525  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.539  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.611  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.612  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.868  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.869  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.87   +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.871  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.885  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.887  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.887  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.888  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 06:02:33.891  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.897  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:33.9    +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoFronten... 🏁 06:02:33.906  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoBackend... 🏁 06:02:34.007 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.128  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.128  
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.131 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132     app: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132         ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.132                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.133                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             time       : "2022-05-11T23:02:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     resource_version  : "72922"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     uid               : "49df096b-814d-41cd-89c0-36e2b7ac154e"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.134                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                             time       : "2022-05-11T23:02:33Z"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     ]
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     resource_version  : "72919"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                     uid               : "0a8d8a53-6730-462d-b5d1-a8821b928c68"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135             }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135         }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135     }
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.135 
💀    🚀 deployDemoDbDeplo... 🏁 06:02:34.136 hello world
💀    🚀 deployDemoFronten... 🏁 06:02:34.147  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 06:02:34.147  
💀    🚀 deployDemoFronten... 🏁 06:02:34.148 Outputs:
💀    🚀 deployDemoFronten... 🏁 06:02:34.148     app: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148         ready    : [
💀    🚀 deployDemoFronten... 🏁 06:02:34.148             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.148             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.148         ]
💀    🚀 deployDemoFronten... 🏁 06:02:34.148         resources: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 06:02:34.148                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                             spec      : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                 selector: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.149                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                          app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                          app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                      }
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                  }
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                  template: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                      metadata: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                          labels: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.15                                              app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                 ]
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.151                                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                         ]
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152 
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     labels            : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                         [0]: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.152                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                         }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                     }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                                 }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                             }
💀    🚀 deployDemoFronten... 🏁 06:02:34.153                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 06:02:34.153       
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoBackend... 🏁 06:02:34.405  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.472  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.753  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.755  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.766  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.778  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.778  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.783  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 06:02:34.791  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 06:02:34.815  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 06:02:34.816  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 06:02:34.957  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 06:02:34.958  
💀    🚀 deployDemoBackend... 🏁 06:02:34.96  Outputs:
💀    🚀 deployDemoBackend... 🏁 06:02:34.961     app: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961         ready    : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.961             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961         ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.961         resources: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.961                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                             spec      : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 selector: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                 template: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                         labels: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "false"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.962 
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.962 
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.962 
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: (json) []
💀    🚀 deployDemoBackend... 🏁 06:02:34.962 
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "600"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.962                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.963                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.964                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [44]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [45]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [46]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [47]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [48]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [49]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [50]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                                 ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                         ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965 
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                     creation_timestamp: "2022-05-11T23:02:34Z"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                     generation        : 1
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                     labels            : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.965                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                         [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                     f:template               : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                         f:metadata: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                             f:labels: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                         f:spec    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                             f:containers                   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                 k:{"name":"demo-backend"}: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                     f:env                     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.966                                                         k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.967                                                         k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.968                                                         k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                     f:ports                   : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             operation  : "Update"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             time       : "2022-05-11T23:02:34Z"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     name              : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     namespace         : "default"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     resource_version  : "72953"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     uid               : "76ff6947-7d30-4da2-bb19-c21c8cdf77ee"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                 }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                 spec       : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     progress_deadline_seconds: 600
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     replicas                 : 1
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     revision_history_limit   : 10
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     selector                 : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         match_labels: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     strategy                 : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         rolling_update: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             max_surge      : "25%"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             max_unavailable: "25%"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         type          : "RollingUpdate"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                     template                 : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         metadata: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             labels: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                 app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                 app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                         spec    : {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                             containers                      : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                 [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                     env                       : [
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [0]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: "HS256"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [1]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: "30"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [2]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [3]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: "/token/"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [4]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: "false"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [5]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                 [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.969 
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [6]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                 [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.969 
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                         [7]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             value: (json) [
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                                 [0]: "*"
💀    🚀 deployDemoBackend... 🏁 06:02:34.969                                             ]
💀    🚀 deployDemoBackend... 🏁 06:02:34.97  
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [8]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [9]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: (json) []
💀    🚀 deployDemoBackend... 🏁 06:02:34.97  
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [10]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "600"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [11]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [12]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [13]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "1"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [14]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "10"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [15]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                              value: "guest"
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          }
💀    🚀 deployDemoBackend... 🏁 06:02:34.97                                          [16]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.971                                             name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 06:02:34.971                                             value: "3000"
💀    🚀 deployDemoBackend... 🏁 06:02:34.971                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.971                                         [17]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         [18]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         [19]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         [20]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         }
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                         [21]: {
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 06:02:34.972                                             value: "htt
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 06:02:35.077 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 15.115820779s
         Current Time: 06:02:35
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 15.226396033s
         Current Time: 06:02:35
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.305µs
         Current Time: 06:02:35
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.565 🚧 Install pip packages.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 06:02:35.568 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 06:02:35.578 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.878 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.882 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.887 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.888 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.889 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.893 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.894 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.895 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.896 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.898 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.899 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.903 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.905 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.905 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.908 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.909 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.909 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.91  Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.914 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.915 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.917 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.917 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.919 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.921 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.922 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.922 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.923 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.925 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.928 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.93  Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.931 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.938 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.941 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.957 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.963 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.965 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:35.968 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.968 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.977 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.981 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 06:02:35.984 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 06:02:35.991 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 06:02:36.001 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.003 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 06:02:36.005 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 06:02:36.02  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:02:36.025 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 06:02:36.025 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 06:02:36.039 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 06:02:36.039 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 06:02:36.039 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 06:02:36.04  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoBacken... 🏁 06:02:36.057 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 06:02:36.057 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.362 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 06:02:36.362 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.445 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.446 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 dependencies.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.453 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     dependencies:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 for this case.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 Usage:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 06:02:36.454 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 Aliases:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 Flags:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 06:02:36.455 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.509 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.509 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.509 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.509 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51      - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  for this case.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Usage:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51    -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51  Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.51        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511 
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.511 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 06:02:36.513 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 06:02:37.577 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 06:02:37.658 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 06:02:37.658 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  dependencies.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.71  
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     dependencies:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 for this case.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 Usage:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711 Aliases:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.711   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 Flags:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 06:02:37.712       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713 
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 06:02:37.713 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 06:02:38.061 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 06:02:38.168 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.171  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.172  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.173  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.174  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.175  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.175  
💀    🚀 destroyDemoFronte... 🏁 06:02:38.177 Outputs:
💀    🚀 destroyDemoFronte... 🏁 06:02:38.177   - app: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.177       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178         ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178       - resources: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.178               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                               - template: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                                 ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                         ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.179                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    - labels            : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                        - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                        - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                        - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                    - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                    - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                  }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                    - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                        - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                    - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                        - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                            - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                              }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                        - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                            - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                    - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                        - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                        - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                        - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                        - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                                  }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                              }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                      }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                                  }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                              }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                            - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                          }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                      ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.18                    - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - resource_version  : "72921"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - uid               : "1a58d9e4-f345-4c40-84e6-74e968d3be3f"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                     ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                             ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.181                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                     ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.182                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183                   - resource_version  : "72920"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183                   - uid               : "3b09e960-127c-47cc-805f-4271f0c4f440"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183 Resources:
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.183 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.224 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 06:02:38.302 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.306  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.306  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.357 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.357  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.359  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.362  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.364  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.365  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.365  
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.366 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.367   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.367       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.367       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.367       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368         ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.368                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.369                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                      }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.37                                                        - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.371                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                     ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - resource_version  : "72922"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - uid               : "49df096b-814d-41cd-89c0-36e2b7ac154e"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.372                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.373                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                             ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.374                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.375                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                     ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.376                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377                   - resource_version  : "72919"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377                   - uid               : "0a8d8a53-6730-462d-b5d1-a8821b928c68"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.377 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 06:02:38.426  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.426  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.428  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.429  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.438  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.444  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.444  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.445  
💀    🚀 destroyDemoFronte... 🏁 06:02:38.445 Outputs:
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446   - app: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446         ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446       - resources: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                               - template: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.446                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                 ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                         ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                     ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - resource_version  : "72921"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - uid               : "1a58d9e4-f345-4c40-84e6-74e968d3be3f"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.447                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                     ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                             ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.448                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                                     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                     ]
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   - resource_version  : "72920"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                   - uid               : "3b09e960-127c-47cc-805f-4271f0c4f440"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449                 }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449             }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449         }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449     }
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 Resources:
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449     - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 06:02:38.449 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.499 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.503  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.504  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.609  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.613  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.613  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.614  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.616  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.622  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.622  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.623  
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624         ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.624                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.625                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626 
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                   - creation_timestamp: "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.626                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.627                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - time       : "2022-05-11T23:02:33Z"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                     ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - resource_version  : "72922"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - uid               : "49df096b-814d-41cd-89c0-36e2b7ac154e"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                 }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                     }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                             }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.628                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                         }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.629                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                    -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                            - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                    -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                            - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                            - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                    -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                            - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                            - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                          }
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                      ]
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                    - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 06:02:38.63                                    - image_pull_policy         : "IfNotPres
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 06:02:39.298 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 06:02:39.366 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.367  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.367  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.371  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.375  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.377  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.378  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.379  
💀    🚀 destroyDemoBacken... 🏁 06:02:39.38  Outputs:
💀    🚀 destroyDemoBacken... 🏁 06:02:39.38    - app: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.38        - ready    : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381       - resources: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.381                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                               - template: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.382                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.383                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.384                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.385                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.386                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.387                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.388                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                 ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                                 ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.389                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                    - creation_timestamp: "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                    - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                    - labels            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                        - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                        - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                        - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                      }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                    - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                            - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                            - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                    - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                      }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                    - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                      }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                  }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                    - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                        - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                          }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                      }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                    - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.39                                        - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.391                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.392                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.393                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.394                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - time       : "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - resource_version  : "72953"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - uid               : "76ff6947-7d30-4da2-bb19-c21c8cdf77ee"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.395                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.396                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - value: "false"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                             ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                             ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                             ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                           - name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.397                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - value: "600"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.398                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.399                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "local"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                           }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                     -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.4                                             - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.401                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.402                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                             ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.403                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                 ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.404                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                   - creation_timestamp: "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.405                           - time       : "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - resource_version  : "72955"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - uid               : "ce8b863b-610c-43c1-aa17-31d431c50b08"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - cluster_ip             : "10.107.52.169"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   -     [0]: "10.107.52.169"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.406                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - status     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.407 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - creation_timestamp: "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                           - time       : "2022-05-11T23:02:34Z"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                     ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - resource_version  : "72954"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.408                   - uid               : "022532a4-080a-4e32-8d76-35702f9de206"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409 Resources:
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.409 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 06:02:39.45  
💀    🚀 destroyDemoBacken... 🏁 06:02:39.451  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.451  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.457  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.523  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.529  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.53   -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.53   -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.531  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.536  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.543  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.549  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.549  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.549  
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554 Outputs:
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554   - app: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554       - resources: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.554                             }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                 }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                               - template: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                         }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.555                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                         ]
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556 
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.556                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.557                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                     }
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 06:02:39.558              
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 06:02:39.663 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 4.215041528s
         Current Time: 06:02:39
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.325666685s
         Current Time: 06:02:39
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

