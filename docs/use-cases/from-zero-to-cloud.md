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

zaruba task setConfig startDemoFrontendContainer '{"localhost": "localhost"}'

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
         Elapsed Time: 1.043µs
         Current Time: 23:54:49
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 23:54:49.786 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 23:54:49.79  🎉🎉🎉
💀    🚀 initProject          🚧 23:54:49.79  Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 115.177529ms
         Current Time: 23:54:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 316.345974ms
         Current Time: 23:54:50
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.128µs
         Current Time: 23:54:50
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:54:50.25  Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:54:50.253         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:54:50.253     
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:54:50.253   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:54:50.253   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:54:50.253   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:54:50.253 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 23:54:50.721 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 23:54:50.722 Preparing base variables
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Base variables prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Preparing start command
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Start command prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.826 Preparing test command
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Test command prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Preparing check command
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Check command prepared
💀    🚀 makeMysqlApp         🐬 23:54:50.827 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 23:54:51.095 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 23:54:51.104 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 23:54:51.111 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 23:54:51.112 ✅ Validate
💀    🚀 makeMysqlApp         🐬 23:54:51.112 Validate app directory
💀    🚀 makeMysqlApp         🐬 23:54:51.112 Done validating app directory
💀    🚀 makeMysqlApp         🐬 23:54:51.112 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 23:54:51.116 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 23:54:51.116 Validate template locations
💀    🚀 makeMysqlApp         🐬 23:54:51.128 Done validating template locations
💀    🚀 makeMysqlApp         🐬 23:54:51.128 Validate app ports
💀    🚀 makeMysqlApp         🐬 23:54:51.134 Done validating app ports
💀    🚀 makeMysqlApp         🐬 23:54:51.135 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 23:54:51.139 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 23:54:51.139 🚧 Generate
💀    🚀 makeMysqlApp         🐬 23:54:51.139 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 23:54:51.139   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 23:54:51.139 ]
💀    🚀 makeMysqlApp         🐬 23:54:51.139 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 23:54:51.159 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 23:54:51.159 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 23:54:51.159 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.637 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.637 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.798 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.799 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:51.799 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.047 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.054 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.061 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.061 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.061 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.061 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.061 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.065 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.065 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.081 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.081 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.084 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.084 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087 ]
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.087 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.127 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.13  Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.133 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.295 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.454 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.617 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.782 Checking start
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.786 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 23:54:52.947 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.115 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.118 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.281 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.444 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.607 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.773 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.777 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:53.941 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.095 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.098 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.267 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.429 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.432 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.582 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.752 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.756 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 23:54:54.93  Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 23:54:55.103 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 23:54:55.106 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 23:54:55.106 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.962834418s
         Current Time: 23:54:55
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.07386825s
         Current Time: 23:54:55
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.149µs
         Current Time: 23:54:55
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:54:55.487 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:54:55.49          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:54:55.49      
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:54:55.49    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:54:55.49    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:54:55.49    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:54:55.49  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 23:54:55.956 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 23:54:55.956 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.112 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 23:54:56.354 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 23:54:56.36  Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 23:54:56.367 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 23:54:56.367 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 23:54:56.367 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 23:54:56.367 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 23:54:56.367 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 23:54:56.37  Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 23:54:56.37  Validate template locations
💀    🚀 makeFastApiApp       ⚡ 23:54:56.379 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 23:54:56.38  Validate app ports
💀    🚀 makeFastApiApp       ⚡ 23:54:56.383 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 23:54:56.383 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386 ]
💀    🚀 makeFastApiApp       ⚡ 23:54:56.386 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 23:54:56.956 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 23:54:56.957 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 23:54:56.957 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 23:54:57.394 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 23:54:57.394 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.342 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.343 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.578 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.586 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.594 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.594 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.594 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.594 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.594 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.597 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.597 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.615 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.615 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.618 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.618 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.621 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.621 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.621 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.621   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.622   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.622   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.622 ]
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.622 
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.622 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.666 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.671 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.674 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.845 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 23:54:58.848 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.011 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.191 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.195 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.367 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.536 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.539 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.712 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.884 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 23:54:59.888 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.07  Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.251 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.255 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.423 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.591 Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.771 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.946 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:00.95  Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.143 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.324 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.328 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.522 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.703 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.706 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 23:55:01.883 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.068 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.072 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.261 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.437 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.621 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 23:55:02.823 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.013 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.198 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.385 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.58  Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.763 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 23:55:03.763 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 23:55:04.305 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 23:55:04.305 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Preparing start command
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Start command prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Preparing test command
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Test command prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.355 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 23:55:05.356 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.356 Preparing check command
💀    🚀 addFastApiModule     ⚡ 23:55:05.356 Check command prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.356 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 23:55:05.629 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 23:55:05.639 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 23:55:05.647 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 23:55:05.647 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 23:55:05.647 Validate app directory
💀    🚀 addFastApiModule     ⚡ 23:55:05.647 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 23:55:05.647 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 23:55:05.651 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 23:55:05.651 Validate template locations
💀    🚀 addFastApiModule     ⚡ 23:55:05.662 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 23:55:05.662 Validate app ports
💀    🚀 addFastApiModule     ⚡ 23:55:05.666 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 23:55:05.666 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 23:55:05.669 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 23:55:05.67  🚧 Generate
💀    🚀 addFastApiModule     ⚡ 23:55:05.67  🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 23:55:05.67    "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 23:55:05.67  ]
💀    🚀 addFastApiModule     ⚡ 23:55:05.67  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 23:55:05.688 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 23:55:05.688 Registering module
💀    🚀 addFastApiModule     ⚡ 23:55:05.729 Done registering module
💀    🚀 addFastApiModule     ⚡ 23:55:05.73  🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 23:55:05.73  Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 23:55:06.098 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 23:55:06.098 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.196 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:07.482 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:07.49  Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:07.497 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.498 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 23:55:07.51  Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 23:55:07.51  Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 23:55:07.584 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 23:55:07.584 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 23:55:07.649 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 23:55:07.649 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 23:55:07.768 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 23:55:07.768 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:07.846 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:08.131 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:08.139 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 23:55:08.146 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 23:55:08.146 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 23:55:08.146 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 23:55:08.146 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 23:55:08.146 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 23:55:08.151 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 23:55:08.151 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 23:55:08.162 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 23:55:08.162 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 23:55:08.166 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 23:55:08.166 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 ]
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 
💀    🚀 addFastApiCrud       ⚡ 23:55:08.169 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 23:55:08.199 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 23:55:08.199 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 23:55:08.239 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 23:55:08.239 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 23:55:08.289 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 23:55:08.289 Registering repo
💀    🚀 addFastApiCrud       ⚡ 23:55:08.348 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 23:55:08.348 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 23:55:08.348 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 12.971179411s
         Current Time: 23:55:08
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 13.172019333s
         Current Time: 23:55:08
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.452µs
         Current Time: 23:55:08
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:55:08.828 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:55:08.831         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:55:08.831     
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:55:08.831   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:55:08.831   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:55:08.831   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:55:08.831 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 23:55:09.267 🧰 Prepare
💀    🚀 makeNginxApp         📗 23:55:09.267 Preparing base variables
💀    🚀 makeNginxApp         📗 23:55:09.345 Base variables prepared
💀    🚀 makeNginxApp         📗 23:55:09.345 Preparing start command
💀    🚀 makeNginxApp         📗 23:55:09.345 Start command prepared
💀    🚀 makeNginxApp         📗 23:55:09.345 Preparing prepare command
💀    🚀 makeNginxApp         📗 23:55:09.345 Prepare command prepared
💀    🚀 makeNginxApp         📗 23:55:09.345 Preparing test command
💀    🚀 makeNginxApp         📗 23:55:09.346 Test command prepared
💀    🚀 makeNginxApp         📗 23:55:09.346 Preparing migrate command
💀    🚀 makeNginxApp         📗 23:55:09.346 Migrate command prepared
💀    🚀 makeNginxApp         📗 23:55:09.346 Preparing check command
💀    🚀 makeNginxApp         📗 23:55:09.346 Check command prepared
💀    🚀 makeNginxApp         📗 23:55:09.346 Preparing replacement map
💀    🚀 makeNginxApp         📗 23:55:09.559 Add config to replacement map
💀    🚀 makeNginxApp         📗 23:55:09.565 Add env to replacement map
💀    🚀 makeNginxApp         📗 23:55:09.572 Replacement map prepared
💀    🚀 makeNginxApp         📗 23:55:09.572 ✅ Validate
💀    🚀 makeNginxApp         📗 23:55:09.572 Validate app directory
💀    🚀 makeNginxApp         📗 23:55:09.572 Done validating app directory
💀    🚀 makeNginxApp         📗 23:55:09.572 Validate app container volumes
💀    🚀 makeNginxApp         📗 23:55:09.575 Done validating app container volumes
💀    🚀 makeNginxApp         📗 23:55:09.575 Validate template locations
💀    🚀 makeNginxApp         📗 23:55:09.585 Done validating template locations
💀    🚀 makeNginxApp         📗 23:55:09.585 Validate app ports
💀    🚀 makeNginxApp         📗 23:55:09.589 Done validating app ports
💀    🚀 makeNginxApp         📗 23:55:09.589 Validate app crud fields
💀    🚀 makeNginxApp         📗 23:55:09.593 Done validating app crud fields
💀    🚀 makeNginxApp         📗 23:55:09.593 🚧 Generate
💀    🚀 makeNginxApp         📗 23:55:09.593 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 23:55:09.593   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 23:55:09.593 ]
💀    🚀 makeNginxApp         📗 23:55:09.593 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 23:55:09.616 🔩 Integrate
💀    🚀 makeNginxApp         📗 23:55:09.616 🎉🎉🎉
💀    🚀 makeNginxApp         📗 23:55:09.617 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 23:55:10.05  🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 23:55:10.05  Preparing base variables
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing start command
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Start command prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing test command
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Test command prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing check command
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Check command prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.168 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 23:55:10.4   Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 23:55:10.407 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 23:55:10.413 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 23:55:10.413 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 23:55:10.414 Validate app directory
💀    🚀 makeNginxAppRunner   📗 23:55:10.414 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 23:55:10.414 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 23:55:10.417 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 23:55:10.417 Validate template locations
💀    🚀 makeNginxAppRunner   📗 23:55:10.428 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 23:55:10.428 Validate app ports
💀    🚀 makeNginxAppRunner   📗 23:55:10.431 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 23:55:10.431 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 23:55:10.434 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 23:55:10.434 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 23:55:10.434 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 23:55:10.434   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 23:55:10.434   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 23:55:10.434 ]
💀    🚀 makeNginxAppRunner   📗 23:55:10.434 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 23:55:10.46  🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 23:55:10.463 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 23:55:10.467 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 23:55:10.636 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 23:55:10.826 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 23:55:11.02  Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 23:55:11.204 Checking start
💀    🚀 makeNginxAppRunner   📗 23:55:11.207 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 23:55:11.37  Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 23:55:11.535 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 23:55:11.538 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 23:55:11.7   Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 23:55:11.864 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 23:55:12.023 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 23:55:12.193 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 23:55:12.196 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 23:55:12.367 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 23:55:12.557 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 23:55:12.56  Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 23:55:12.731 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 23:55:12.903 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 23:55:12.906 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 23:55:13.093 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 23:55:13.267 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 23:55:13.27  Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 23:55:13.451 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 23:55:13.653 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 23:55:13.657 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 23:55:13.657 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.934997723s
         Current Time: 23:55:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.045700051s
         Current Time: 23:55:13
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.733µs
         Current Time: 23:55:14
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:55:14.065 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 23:55:14.171 Synchronize task environments
💀    🚀 syncEnv              🔄 23:55:14.37  Synchronize project's environment files
💀    🚀 syncEnv              🔄 23:55:14.559 🎉🎉🎉
💀    🚀 syncEnv              🔄 23:55:14.559 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 600.360644ms
         Current Time: 23:55:14
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 801.718779ms
         Current Time: 23:55:14
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 5.237µs
         Current Time: 23:55:15
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 23:55:15.228 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 23:55:15.228 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoDbImage     🏭 23:55:15.496 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 23:55:15.496 Build image demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 23:55:15.496 Build image demo-backend:latest
💀    🚀 buildDemoFrontend... 🏭 23:55:16.486 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 23:55:16.488 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoDbImage     🏭 23:55:16.555 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 23:55:16.555  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 23:55:16.555 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 23:55:16.563 Successfully tagged demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 23:55:16.566 Sending build context to Docker daemon   1.03MB
💀    🚀 buildDemoFrontend... 🏭 23:55:16.571 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 23:55:16.571  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 23:55:16.572 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 23:55:16.573  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.573  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 23:55:16.573 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 23:55:16.577  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.577  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 23:55:16.578 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:55:16.579  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.579  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 23:55:16.579 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:55:16.58   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.58   ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 23:55:16.581 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 23:55:16.583  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.583  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 23:55:16.584 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 23:55:16.584  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:16.584  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 23:55:16.585 Successfully built 6ea76668c578
💀    🚀 buildDemoBackendI... 🏭 23:55:16.591 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 23:55:16.596  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 23:55:16.597 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 23:55:16.597  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.597  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 23:55:16.597 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoFrontend... 🏭 23:55:16.597 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 23:55:16.6   🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 23:55:16.6   Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 23:55:16.6    ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.6    ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 23:55:16.6   Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 23:55:16.601  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.601  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 23:55:16.602 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 23:55:16.602  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.602  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 23:55:16.602 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616  ---> 97fdfef7cb48
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616  ---> bf9c545afbe0
💀    🚀 buildDemoBackendI... 🏭 23:55:16.616 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 23:55:16.617  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.617  ---> a62a483a9091
💀    🚀 buildDemoBackendI... 🏭 23:55:16.617 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 23:55:16.617  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:55:16.617  ---> db465fe79375
💀    🚀 buildDemoBackendI... 🏭 23:55:16.619 Successfully built db465fe79375
💀    🚀 buildDemoBackendI... 🏭 23:55:16.626 Successfully tagged demo-backend:latest
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 23:55:16.742 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 1.620801806s
         Current Time: 23:55:16
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.731299843s
         Current Time: 23:55:16
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.279µs
         Current Time: 23:55:17
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 updateProjectLinks   🔗 23:55:17.392 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 23:55:17.392 Links updated
💀    🚀 prepareDemoBackend   🔧 23:55:17.395 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 23:55:17.44  🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 23:55:17.524 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 23:55:17.647 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 23:55:18.668 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 23:55:18.673 Sending build context to Docker daemon  20.48kB
💀    🚀 buildDemoDbImage     🏭 23:55:18.744 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 23:55:18.744  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 23:55:18.744 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 23:55:18.745 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoDbImage     🏭 23:55:18.747 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 23:55:18.747  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 23:55:18.747 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 23:55:18.747  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.747  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 23:55:18.747 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 23:55:18.748  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.748  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 23:55:18.748 Step 4/11 : USER 1001
💀    🚀 buildDemoDbImage     🏭 23:55:18.749 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 23:55:18.749 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 23:55:18.749  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.75   ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 23:55:18.75  Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 23:55:18.754  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.754  ---> 2eae8b6cd23a
💀    🚀 buildDemoFrontend... 🏭 23:55:18.754 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> f7779f873da5
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> 695d610f8d47
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> 55dee5d4680a
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> 78f649e6f9d4
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> 01b89502a453
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764  ---> 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Successfully built 6ea76668c578
💀    🚀 buildDemoFrontend... 🏭 23:55:18.764 Successfully tagged demo-frontend:latest
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoDbContainer 🐬 23:55:19.212 🔎 Waiting docker container 'demoDb' running status
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 23:55:19.22  🔎 Waiting docker container 'demoFrontend' running status
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:19.27  Error: No such container: demoDb
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:19.271 Error: No such container: demoDb
💀 🔥 🔎 startDemoFrontend... 📗 23:55:19.3   Error: No such container: demoFrontend
💀 🔥 🚀 startDemoFrontend... 📗 23:55:19.301 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:19.318 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 23:55:19.322 🐳 Creating and starting container 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 23:55:19.348 Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 23:55:19.35  🐳 Creating and starting container 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 23:55:19.759 0741a5ddbb687f48ce390320e37636c1e375a09dfe206338b04a563f68c9ffd2
💀    🚀 startDemoFrontend... 📗 23:55:19.835 9cccb44c2f40eb5bdd85835d89785ec8ba3196873b9ea97cc659ef9ed95cbed8
💀    🚀 prepareDemoBackend   🔧 23:55:21.386 Activate venv
💀    🚀 prepareDemoBackend   🔧 23:55:21.386 Install dependencies
💀    🚀 startDemoDbContainer 🐬 23:55:21.69  🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 23:55:21.696 🔎 Waiting docker container 'demoDb' healthcheck
💀    🔎 startDemoDbContainer 🐬 23:55:21.752 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 23:55:21.752 🔎 Waiting for host port: '3306'
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.757 [38;5;6mmysql [38;5;5m16:55:21.75 
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.758 [38;5;6mmysql [38;5;5m16:55:21.75 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.76  [38;5;6mmysql [38;5;5m16:55:21.75 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.761 [38;5;6mmysql [38;5;5m16:55:21.76 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.763 [38;5;6mmysql [38;5;5m16:55:21.76 
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.764 [38;5;6mmysql [38;5;5m16:55:21.76 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀    🔎 startDemoDbContainer 🐬 23:55:21.77  🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.8   [38;5;6mmysql [38;5;5m16:55:21.79 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.814 [38;5;6mmysql [38;5;5m16:55:21.81 [38;5;2mINFO  ==> Initializing mysql database
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.838 [38;5;6mmysql [38;5;5m16:55:21.83 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.846 [38;5;6mmysql [38;5;5m16:55:21.84 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.864 [38;5;6mmysql [38;5;5m16:55:21.86 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.872 [38;5;6mmysql [38;5;5m16:55:21.87 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:21.88  [38;5;6mmysql [38;5;5m16:55:21.87 [38;5;2mINFO  ==> Installing database
💀    🚀 prepareDemoBackend   🔧 23:55:21.881 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 startDemoFrontend... 📗 23:55:22.136 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 23:55:22.147 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 prepareDemoBackend   🔧 23:55:22.223   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 startDemoFrontend... 📗 23:55:22.226 
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.226 [38;5;6mnginx [38;5;5m16:55:22.17 
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.226 [38;5;6mnginx [38;5;5m16:55:22.18 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.226 [38;5;6mnginx [38;5;5m16:55:22.21 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.226 [38;5;6mnginx [38;5;5m16:55:22.22 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.226 [38;5;6mnginx [38;5;5m16:55:22.22 
💀    🚀 prepareDemoBackend   🔧 23:55:22.238 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🔎 startDemoFrontend... 📗 23:55:22.243 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 23:55:22.243 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 23:55:22.249 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 23:55:22.25  🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 23:55:22.253 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.255 [38;5;6mnginx [38;5;5m16:55:22.25 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.28  2022/05/13 16:55:22 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 23:55:22.28  nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 prepareDemoBackend   🔧 23:55:22.357   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:22.371 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 23:55:22.453   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🔎 startDemoDbContainer 🐬 23:55:24.775 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:24.991 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:25.009 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🔎 startDemoFrontend... 📗 23:55:25.257 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 23:55:25.488 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🔎 startDemoFrontend... 📗 23:55:25.491 check demoFrontend
💀    🔎 startDemoFrontend... 📗 23:55:25.501 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🚀 prepareDemoBackend   🔧 23:55:25.684   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:25.699 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 23:55:25.807   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:25.819 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 23:55:25.955   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:25.994 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 23:55:26.097   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:26.111 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 23:55:26.271   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:26.384 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 23:55:26.757   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:26.843 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 23:55:27.014   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:27.066 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 23:55:27.427   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 23:55:27.688 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 23:55:27.884   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:27.902 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧 23:55:27.981   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:27.993 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 23:55:28.091   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.099 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:28.148 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:28.149 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 23:55:28.18    Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.192 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:28.246 [38;5;6mmysql [38;5;5m16:55:28.24 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 23:55:28.281   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.31  Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 23:55:28.404   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.419 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 23:55:28.649   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.665 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 23:55:28.739   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.75  Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:28.889   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:28.911 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 23:55:28.999   Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:29.009 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 23:55:29.07    Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:29.082 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 23:55:29.173   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 23:55:29.357 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 23:55:29.487   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🔎 startDemoFrontend... 📗 23:55:29.504 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 23:55:29.504 📜 Task 'startDemoFrontendContainer' is ready
💀    🚀 prepareDemoBackend   🔧 23:55:29.505 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 23:55:29.964   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:30.065 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 23:55:30.171   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:30.184 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 23:55:30.259   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:30.263 [38;5;6mmysql [38;5;5m16:55:30.26 [38;5;2mINFO  ==> Configuring authentication
💀    🚀 prepareDemoBackend   🔧 23:55:30.266 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:30.305 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:30.326 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 prepareDemoBackend   🔧 23:55:30.346   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:30.363 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:30.371 [38;5;6mmysql [38;5;5m16:55:30.37 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:30.374 [38;5;6mmysql [38;5;5m16:55:30.37 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 23:55:30.459   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:30.477 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 23:55:30.545   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 23:55:30.689 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧 23:55:30.826   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:30.844 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 23:55:31.136   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.152 Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 23:55:31.241   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.247 Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:31.301 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:31.302 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 23:55:31.336   Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.343 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:31.423   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.428 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:31.494   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.499 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:31.585   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.594 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:31.668   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.678 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:31.745   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:31.775 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 23:55:32.249   Using cached https://files.pythonhosted.org/packages/c1/38/a9fd8c7bb151325d8b3d9108ce791348c84171b5d9f346b0bf0639de603f/coverage-6.3.3-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.264 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 23:55:32.346   Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.355 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:32.389 [38;5;6mmysql [38;5;5m16:55:32.38 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 23:55:32.425   Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.431 Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 23:55:32.59    Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.604 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 23:55:32.717   Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.727 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 23:55:32.864   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:32.881 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 23:55:32.952   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 23:55:33.288 Installing collected packages: aiofiles, asgiref, avro-python3, pycparser, cffi, six, bcrypt, certifi, charset-normalizer, click, idna, urllib3, requests, fastavro, confluent-kafka, cryptography, starlette, typing-extensions, pydantic, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, pyparsing, packaging, toml, iniconfig, py, pluggy, attrs, pytest, tomli, coverage, pytest-cov, ecdsa, pyasn1, rsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 23:55:33.321   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 23:55:33.568     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 23:55:34.018   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:34.465 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:34.467 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:37.624 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:37.626 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:38.452 [38;5;6mmysql [38;5;5m16:55:38.45 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:38.467 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:38.475 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:38.478 [38;5;6mmysql [38;5;5m16:55:38.47 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:40.741 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:40.743 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 startDemoDbContainer 🐬 23:55:41.494 
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:41.494 [38;5;6mmysql [38;5;5m16:55:41.49 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 23:55:41.552 [38;5;6mmysql [38;5;5m16:55:41.55 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 23:55:41.755 2022-05-13T16:55:41.751272Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 23:55:41.756 2022-05-13T16:55:41.752814Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 23:55:41.756 2022-05-13T16:55:41.752822Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 23:55:41.757 2022-05-13T16:55:41.756866Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 23:55:41.882 2022-05-13T16:55:41.882350Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 23:55:42.066 2022-05-13T16:55:42.065911Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 23:55:42.066 2022-05-13T16:55:42.065954Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 23:55:42.08  2022-05-13T16:55:42.079577Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 23:55:42.08  2022-05-13T16:55:42.079787Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 23:55:43.906 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 Database
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 information_schema
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 mysql
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 performance_schema
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 sample
💀    🔎 startDemoDbContainer 🐬 23:55:43.912 sys
💀    🔎 startDemoDbContainer 🐬 23:55:43.917 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 23:55:47.919 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 23:55:47.919 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 23:56:02.439     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 23:56:03.902   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 23:56:04.127     Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 23:56:04.778   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 23:56:04.941     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 23:56:05.026 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.3 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 23:56:05.105 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 23:56:05.105 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 23:56:05.139 Prepare
💀    🚀 prepareDemoBackend   🔧 23:56:05.139 prepare command
💀    🚀 prepareDemoBackend   🔧 23:56:05.139 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 startDemoBackend     ⚡ 23:56:05.365 Activate venv
💀    🚀 startDemoBackend     ⚡ 23:56:05.365 Start
💀    🔎 startDemoBackend     ⚡ 23:56:05.365 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 23:56:05.895 2022-05-13 23:56:05,895 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 23:56:05.895 2022-05-13 23:56:05,895 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.9   2022-05-13 23:56:05,900 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 23:56:05.901 2022-05-13 23:56:05,900 INFO sqlalchemy.engine.Engine [generated in 0.00024s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.904 2022-05-13 23:56:05,904 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 23:56:05.904 2022-05-13 23:56:05,904 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.907 2022-05-13 23:56:05,906 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:05.907 2022-05-13 23:56:05,907 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 23:56:05.907 2022-05-13 23:56:05,907 INFO sqlalchemy.engine.Engine [generated in 0.00013s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  2022-05-13 23:56:05,910 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  )
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  
💀    🚀 startDemoBackend     ⚡ 23:56:05.91  2022-05-13 23:56:05,910 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.933 2022-05-13 23:56:05,933 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 23:56:05.933 2022-05-13 23:56:05,933 INFO sqlalchemy.engine.Engine [no key 0.00024s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.956 2022-05-13 23:56:05,956 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 23:56:05.956 2022-05-13 23:56:05,956 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.972 2022-05-13 23:56:05,972 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 23:56:05.973 2022-05-13 23:56:05,972 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:05.99  2022-05-13 23:56:05,990 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 23:56:05.99  2022-05-13 23:56:05,990 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.008 2022-05-13 23:56:06,008 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 23:56:06.009 2022-05-13 23:56:06,009 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:06.01  2022-05-13 23:56:06,010 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 23:56:06.01  2022-05-13 23:56:06,010 INFO sqlalchemy.engine.Engine [cached since 0.1032s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 23:56:06.012 2022-05-13 23:56:06,012 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 23:56:06.012 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 )
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 
💀    🚀 startDemoBackend     ⚡ 23:56:06.013 2022-05-13 23:56:06,012 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.031 2022-05-13 23:56:06,031 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 23:56:06.031 2022-05-13 23:56:06,031 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.053 2022-05-13 23:56:06,053 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 23:56:06.053 2022-05-13 23:56:06,053 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.072 2022-05-13 23:56:06,072 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 23:56:06.072 2022-05-13 23:56:06,072 INFO sqlalchemy.engine.Engine [no key 0.00021s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.09  2022-05-13 23:56:06,090 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 23:56:06.091 2022-05-13 23:56:06,091 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:06.091 2022-05-13 23:56:06,091 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 23:56:06.091 2022-05-13 23:56:06,091 INFO sqlalchemy.engine.Engine [cached since 0.1848s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 2022-05-13 23:56:06,094 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.094 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 )
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 
💀    🚀 startDemoBackend     ⚡ 23:56:06.095 2022-05-13 23:56:06,094 INFO sqlalchemy.engine.Engine [no key 0.00024s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.121 2022-05-13 23:56:06,121 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 23:56:06.121 2022-05-13 23:56:06,121 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.139 2022-05-13 23:56:06,139 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 23:56:06.139 2022-05-13 23:56:06,139 INFO sqlalchemy.engine.Engine [no key 0.00019s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.156 2022-05-13 23:56:06,156 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 23:56:06.157 2022-05-13 23:56:06,156 INFO sqlalchemy.engine.Engine [no key 0.00022s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.179 2022-05-13 23:56:06,179 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 23:56:06.179 2022-05-13 23:56:06,179 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.2   2022-05-13 23:56:06,200 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 23:56:06.2   2022-05-13 23:56:06,200 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.22  2022-05-13 23:56:06,220 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 23:56:06.22  2022-05-13 23:56:06,220 INFO sqlalchemy.engine.Engine [no key 0.00041s] {}
💀    🚀 startDemoBackend     ⚡ 23:56:06.24  2022-05-13 23:56:06,240 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 23:56:06.243 2022-05-13 23:56:06,243 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:06.246 2022-05-13 23:56:06,246 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 23:56:06.246 FROM users 
💀    🚀 startDemoBackend     ⚡ 23:56:06.246 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 23:56:06.246  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 23:56:06.246 2022-05-13 23:56:06,246 INFO sqlalchemy.engine.Engine [generated in 0.00021s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 23:56:06.247 2022-05-13 23:56:06,247 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 23:56:06.469 2022-05-13 23:56:06,469 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:06.471 2022-05-13 23:56:06,471 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 23:56:06.471 2022-05-13 23:56:06,471 INFO sqlalchemy.engine.Engine [generated in 0.00024s] {'id': '2c5c8827-caac-4479-a60f-f158d096f2f9', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$LZe8IIQwSOpwefW.MbW6V.qflirq0CyoxZRYSBQ7kpN5DtdSOT2TW', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 13, 23, 56, 6, 469361), 'updated_at': datetime.datetime(2022, 5, 13, 23, 56, 6, 470967)}
💀    🚀 startDemoBackend     ⚡ 23:56:06.472 2022-05-13 23:56:06,472 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 23:56:06.48  2022-05-13 23:56:06,480 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 23:56:06.481 2022-05-13 23:56:06,481 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 23:56:06.481 FROM users 
💀    🚀 startDemoBackend     ⚡ 23:56:06.481 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 23:56:06.481 2022-05-13 23:56:06,481 INFO sqlalchemy.engine.Engine [generated in 0.00014s] {'pk_1': '2c5c8827-caac-4479-a60f-f158d096f2f9'}
💀    🚀 startDemoBackend     ⚡ 23:56:06.482 2022-05-13 23:56:06,482 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 23:56:06.484 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.493 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 23:56:06.506 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.514 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 23:56:06.514 Register library route handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.515 Register library event handler
💀    🚀 startDemoBackend     ⚡ 23:56:06.515 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 23:56:06.515 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:06.515 INFO:     Started server process [31343]
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:06.515 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:06.516 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:06.516 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 23:56:07.37  🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 23:56:07.37  check demoBackend
💀    🔎 startDemoBackend     ⚡ 23:56:07.37  🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 23:56:07.37  📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 23:56:07.478 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 50.199467822s
         Current Time: 23:56:07
         Active Process:
           * (PID=19971) 📗 'startDemoFrontendContainer' service
           * (PID=19964) 🐬 'startDemoDbContainer' service
           * (PID=31339) ⚡ 'startDemoBackend' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=19964)
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=31339)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=19971)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:09.126 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:09.226 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:09.227 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 23:56:09.227 INFO:     Finished server process [31343]
💀    🚀 startDemoBackend     ⚡ 23:56:09.324 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 23:56:09.324 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 52.303060826s
         Current Time: 23:56:09
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 866ns
         Current Time: 23:56:09
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 23:56:09.853 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 23:56:09.853 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 23:56:09.881 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 23:56:09.966 Build image demo-backend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoDbImage     🏭 23:56:10.118 Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 23:56:10.118 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 23:56:12.878 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 23:56:12.878 Sending build context to Docker daemon  22.02kB
💀    🚀 buildDemoDbImage     🏭 23:56:12.928 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 23:56:12.929  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 23:56:12.929 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 23:56:12.932 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoBackendI... 🏭 23:56:12.934 Sending build context to Docker daemon   1.18MB
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935 Step 2/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 23:56:12.935 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.935  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 23:56:12.936 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 23:56:12.936  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.936  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 23:56:12.936 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoDbImage     🏭 23:56:12.937 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 23:56:12.937 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 23:56:12.939  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.939  ---> 11c677f847bc
💀    🚀 buildDemoFrontend... 🏭 23:56:12.939 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 23:56:12.939  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.939  ---> 776095918b33
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94  Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> 48dc42a93a8a
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94  Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> 0beee76410dd
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94  Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94   ---> 68555ae22bc5
💀    🚀 buildDemoFrontend... 🏭 23:56:12.94  Step 10/11 : USER 1001
💀    🚀 buildDemoBackendI... 🏭 23:56:12.941 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 23:56:12.941  ---> caf584a25606
💀    🚀 buildDemoFrontend... 🏭 23:56:12.941  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.941  ---> 992fa94aa2f2
💀    🚀 buildDemoBackendI... 🏭 23:56:12.941 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 23:56:12.941  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 23:56:12.941 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 23:56:12.941  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoFrontend... 🏭 23:56:12.942  ---> 02304e445f6f
💀    🚀 buildDemoFrontend... 🏭 23:56:12.942 Successfully built 02304e445f6f
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 23:56:12.942  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:56:12.943  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 23:56:12.943 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 23:56:12.943  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 23:56:12.943  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 23:56:12.943 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 23:56:12.95  Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 23:56:12.952 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 23:56:12.952 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 23:56:13.185  ---> eef18c6041ce
💀    🚀 buildDemoBackendI... 🏭 23:56:13.185 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 23:56:13.298  ---> Running in ab1c03c14964
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 23:56:13.374 🔎 Waiting docker container 'demoFrontend' running status
💀    🚀 buildDemoBackendI... 🏭 23:56:13.4   Removing intermediate container ab1c03c14964
💀    🚀 buildDemoBackendI... 🏭 23:56:13.4    ---> 508eb15b31a0
💀    🚀 buildDemoBackendI... 🏭 23:56:13.4   Step 8/9 : RUN chmod 755 ./start.sh
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 23:56:13.436 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 23:56:13.441 🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 23:56:13.443 🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 23:56:13.446 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 buildDemoBackendI... 🏭 23:56:13.453  ---> Running in 2000132dbfce
💀    🔎 startDemoDbContainer 🐬 23:56:13.49  🔎 Waiting docker container 'demoDb' healthcheck
💀    🚀 startDemoDbContainer 🐬 23:56:13.504 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 23:56:13.504 🐳 Logging 'demoDb'
💀    🔎 startDemoFrontend... 📗 23:56:13.505 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 23:56:13.505 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 23:56:13.513 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 23:56:13.513 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 23:56:13.515 🔎 Host port '443' is ready
💀    🔎 startDemoDbContainer 🐬 23:56:13.551 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 23:56:13.551 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 23:56:13.552 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 23:56:14.33  Removing intermediate container 2000132dbfce
💀    🚀 buildDemoBackendI... 🏭 23:56:14.33   ---> 49cab60d1873
💀    🚀 buildDemoBackendI... 🏭 23:56:14.33  Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 23:56:14.356  ---> Running in 034f01455346
💀    🚀 buildDemoBackendI... 🏭 23:56:14.418 Removing intermediate container 034f01455346
💀    🚀 buildDemoBackendI... 🏭 23:56:14.418  ---> 510db6bad25e
💀    🚀 buildDemoBackendI... 🏭 23:56:14.421 Successfully built 510db6bad25e
💀    🚀 buildDemoBackendI... 🏭 23:56:14.428 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 23:56:14.43  🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 23:56:14.43  Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoFrontend... 📗 23:56:16.52  🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 23:56:16.556 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 23:56:16.64  check demoFrontend
💀    🔎 startDemoFrontend... 📗 23:56:16.643 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 23:56:16.656 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 Database
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 information_schema
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 mysql
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 performance_schema
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 sample
💀    🔎 startDemoDbContainer 🐬 23:56:16.659 sys
💀    🔎 startDemoDbContainer 🐬 23:56:16.666 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 23:56:20.647 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 23:56:20.647 📜 Task 'startDemoFrontendContainer' is ready
💀    🔎 startDemoDbContainer 🐬 23:56:20.668 🎉🎉🎉
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 23:56:21.408 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:21.439 Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ 23:56:21.44  Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:21.465 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 23:56:21.467 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 23:56:21.522 a8dd52b6d7928ca233566780912196ef8e0a642202cb6c2f8807cf378c561602
💀    🚀 startDemoBackendC... ⚡ 23:56:22.753 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 23:56:22.774 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 23:56:22.821 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 23:56:22.821 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 23:56:22.823 🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 23:56:23.82  2022-05-13 16:56:23,819 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 23:56:23.82  2022-05-13 16:56:23,820 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.824 2022-05-13 16:56:23,823 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 23:56:23.824 2022-05-13 16:56:23,823 INFO sqlalchemy.engine.Engine [generated in 0.00023s] {}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.829 2022-05-13 16:56:23,829 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 23:56:23.829 2022-05-13 16:56:23,829 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.835 2022-05-13 16:56:23,834 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 23:56:23.836 2022-05-13 16:56:23,835 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 23:56:23.836 2022-05-13 16:56:23,835 INFO sqlalchemy.engine.Engine [generated in 0.00019s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.838 2022-05-13 16:56:23,838 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 23:56:23.842 2022-05-13 16:56:23,841 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 23:56:23.842 2022-05-13 16:56:23,842 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 23:56:23.842 2022-05-13 16:56:23,842 INFO sqlalchemy.engine.Engine [cached since 0.007011s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.845 2022-05-13 16:56:23,845 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 23:56:23.849 2022-05-13 16:56:23,848 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 23:56:23.849 2022-05-13 16:56:23,849 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 23:56:23.849 2022-05-13 16:56:23,849 INFO sqlalchemy.engine.Engine [cached since 0.01397s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.853 2022-05-13 16:56:23,852 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 23:56:23.859 2022-05-13 16:56:23,859 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 23:56:23.863 2022-05-13 16:56:23,862 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 23:56:23.863 FROM users 
💀    🚀 startDemoBackendC... ⚡ 23:56:23.863 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 23:56:23.863  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 23:56:23.863 2022-05-13 16:56:23,862 INFO sqlalchemy.engine.Engine [generated in 0.00022s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 23:56:23.866 2022-05-13 16:56:23,865 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 23:56:23.869 Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.881 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 23:56:23.894 Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.903 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 23:56:23.903 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.903 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 23:56:23.903 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 23:56:23.903 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:23.903 INFO:     Started server process [8]
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:23.903 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:23.904 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 23:56:23.904 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 23:56:25.827 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 23:56:25.93  check demoBackend
💀    🔎 startDemoBackendC... ⚡ 23:56:25.936 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 23:56:26.936 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 23:56:26.937 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 23:56:27.045 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 17.299098655s
         Current Time: 23:56:27
         Active Process:
           * (PID=4658) ⚡ 'startDemoBackendContainer' service
           * (PID=2824) 📗 'startDemoFrontendContainer' service
           * (PID=2853) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=4658)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=2824)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=2853)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 19.402530307s
         Current Time: 23:56:29
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 2.1µs
         Current Time: 23:56:29
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 23:56:29.429 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 23:56:29.429 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoBackendCo... ✋ 23:56:29.816 Stop docker container demoBackend
💀    🚀 stopDemoDbContainer  ✋ 23:56:29.827 Stop docker container demoDb
💀    🚀 stopDemoFrontendC... ✋ 23:56:29.829 Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ 23:56:34.291 demoDb
💀    🚀 stopDemoDbContainer  ✋ 23:56:34.293 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 23:56:34.293 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ 23:56:40.489 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 23:56:40.491 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 23:56:40.491 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀    🚀 stopDemoFrontendC... ✋ 23:56:40.992 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 23:56:40.994 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 23:56:40.994 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 23:56:41.103 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.78148049s
         Current Time: 23:56:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 11.892621848s
         Current Time: 23:56:41
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.219µs
         Current Time: 23:56:41
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 23:56:41.498 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 23:56:41.498 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ❌ 'removeDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run ❌ 'removeDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run ❌ 'removeDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🔥 🚀 removeDemoDbConta... ❌ 23:56:41.816 Error: No such container: 
💀 🔥 🚀 removeDemoBackend... ❌ 23:56:41.819 Error: No such container: 
💀    🚀 removeDemoBackend... ❌ 23:56:41.821 Stop docker container demoBackend
💀    🚀 removeDemoDbConta... ❌ 23:56:41.823 Stop docker container demoDb
💀 🔥 🚀 removeDemoFronten... ❌ 23:56:41.823 Error: No such container: 
💀    🚀 removeDemoFronten... ❌ 23:56:41.826 Stop docker container demoFrontend
💀    🚀 removeDemoFronten... ❌ 23:56:41.93  Docker container demoFrontend stopped
💀    🚀 removeDemoFronten... ❌ 23:56:41.93  Remove docker container demoFrontend
💀    🚀 removeDemoBackend... ❌ 23:56:41.931 Docker container demoBackend stopped
💀    🚀 removeDemoBackend... ❌ 23:56:41.931 Remove docker container demoBackend
💀    🚀 removeDemoDbConta... ❌ 23:56:41.941 Docker container demoDb stopped
💀    🚀 removeDemoDbConta... ❌ 23:56:41.941 Remove docker container demoDb
💀    🚀 removeDemoFronten... ❌ 23:56:42.002 demoFrontend
💀    🚀 removeDemoBackend... ❌ 23:56:42.002 demoBackend
💀    🚀 removeDemoBackend... ❌ 23:56:42.005 🎉🎉🎉
💀    🚀 removeDemoBackend... ❌ 23:56:42.005 Docker container demoBackend removed
💀    🚀 removeDemoFronten... ❌ 23:56:42.012 🎉🎉🎉
💀    🚀 removeDemoFronten... ❌ 23:56:42.012 Docker container demoFrontend removed
💀    🚀 removeDemoDbConta... ❌ 23:56:42.039 demoDb
💀    🚀 removeDemoDbConta... ❌ 23:56:42.042 🎉🎉🎉
💀    🚀 removeDemoDbConta... ❌ 23:56:42.042 Docker container demoDb removed
💀 🎉 Successfully running ❌ 'removeDemoBackendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoFrontendContainer' command
💀 🎉 Successfully running ❌ 'removeDemoDbContainer' command
💀 🏁 Run ❌ 'removeContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 removeContainers     ❌ 23:56:42.152 
💀 🎉 Successfully running ❌ 'removeContainers' command
💀 🔎 Job Running...
         Elapsed Time: 760.998468ms
         Current Time: 23:56:42
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 871.352594ms
         Current Time: 23:56:42
zaruba please removeContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.029µs
         Current Time: 23:56:42
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:42.589 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:56:42.591         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:56:42.591     
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:56:42.591   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:56:42.591   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:56:42.591   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:56:42.591 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.083 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.083 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.317 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.647 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.658 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.668 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.668 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.668 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.668 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.668 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.673 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.673 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.687 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.687 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.693 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.693 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699 ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.699 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.765 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.765 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:43.765 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.363 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.363 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.598 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.921 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.932 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.941 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.941 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.942 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.942 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.942 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.947 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.947 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.961 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.961 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.966 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.966 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97  ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:44.97  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.001 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.008 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.012 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.244 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.449 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.454 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:45.679 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:46.001 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:46.005 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:46.242 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:46.243 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:46.243 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.760813421s
         Current Time: 23:56:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.871966267s
         Current Time: 23:56:46
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.329µs
         Current Time: 23:56:46
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:46.651 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:56:46.653         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:56:46.653     
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:56:46.653   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:56:46.653   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:56:46.653   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:56:46.653 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:47.116 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:47.116 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.115 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.388 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.395 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.406 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.406 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.406 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.406 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.406 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.409 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.409 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.42  Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.421 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.424 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.424 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428 ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.428 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.465 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.466 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.466 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.926 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:48.926 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:49.959 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.228 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.236 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.244 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.244 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.244 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.244 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.244 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.248 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.248 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.262 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.262 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.267 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.267 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.271 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.272 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.272 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.272   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.272 ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.272 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_CORS_ALLOW_CREDENTIALS:\n  default: \"false\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_CREDENTIALS\nAPP_CORS_ALLOW_HEADERS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_HEADERS\nAPP_CORS_ALLOW_METHODS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_METHODS\nAPP_CORS_ALLOW_ORIGIN_REGEX:\n  default: \"\"\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGIN_REGEX\nAPP_CORS_ALLOW_ORIGINS:\n  default: '[\"*\"]'\n  from: DEMO_BACKEND_APP_CORS_ALLOW_ORIGINS\nAPP_CORS_EXPOSE_HEADERS:\n  default: '[]'\n  from: DEMO_BACKEND_APP_CORS_EXPOSE_HEADERS\nAPP_CORS_MAX_AGE:\n  default: \"600\"\n  from: DEMO_BACKEND_APP_CORS_MAX_AGE\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_CORS_ALLOW_CREDENTIALS":"false","APP_CORS_ALLOW_HEADERS":"[\"*\"]","APP_CORS_ALLOW_METHODS":"[\"*\"]","APP_CORS_ALLOW_ORIGINS":"[\"*\"]","APP_CORS_ALLOW_ORIGIN_REGEX":"","APP_CORS_EXPOSE_HEADERS":"[]","APP_CORS_MAX_AGE":"600","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.302 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.306 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.31  Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.504 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.7   Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.704 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:50.934 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:51.13  Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:51.134 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:51.32  Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:51.321 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:51.321 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.777233137s
         Current Time: 23:56:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.888530798s
         Current Time: 23:56:51
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.083µs
         Current Time: 23:56:51
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:51.735 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 23:56:51.739         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 23:56:51.739     
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 23:56:51.739   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 23:56:51.739   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 23:56:51.739   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 23:56:51.739 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.208 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.208 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.332 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.333 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.597 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.605 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.612 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.612 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.612 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.612 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.612 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.616 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.616 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.626 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.626 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.63  Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.63  Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.633 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.633 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.633 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.634   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.634 ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.634 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.67  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.67  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:52.67  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.106 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.106 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.233 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.475 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.482 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.49  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.49  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.49  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.49  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.49  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.494 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.494 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.504 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.504 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.507 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.507 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511 ]
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.511 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: http://localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"http://localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.533 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.537 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.54  Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.734 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.939 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:53.942 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.161 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.368 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.372 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.577 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.577 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 23:56:54.577 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.949044167s
         Current Time: 23:56:54
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.060216822s
         Current Time: 23:56:54
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.249µs
         Current Time: 23:56:54
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:54.985 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 23:56:55.09  Synchronize task environments
💀    🚀 syncEnv              🔄 23:56:55.302 Synchronize project's environment files
💀    🚀 syncEnv              🔄 23:56:55.516 🎉🎉🎉
💀    🚀 syncEnv              🔄 23:56:55.516 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 639.330783ms
         Current Time: 23:56:55
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 840.228571ms
         Current Time: 23:56:55
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.22µs
         Current Time: 23:56:56
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:56.016 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 23:56:56.127 🎉🎉🎉
💀    🚀 setProjectValue      🔗 23:56:56.127 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 217.910186ms
         Current Time: 23:56:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 419.029338ms
         Current Time: 23:56:56
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.094µs
         Current Time: 23:56:56
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 23:56:56.607 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 23:56:56.717 🎉🎉🎉
💀    🚀 setProjectValue      🔗 23:56:56.717 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.603606ms
         Current Time: 23:56:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 417.811734ms
         Current Time: 23:56:57
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 7.605µs
         Current Time: 23:56:57
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoFronte... 🏁 23:56:57.216 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 23:56:57.217 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 23:56:57.22  🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 23:56:59.805 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 23:56:59.808 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 23:56:59.843 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 23:57:00.9   Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:01.107 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:02.539 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:03.89    Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:04     Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:04.174   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:04.284 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:04.46    Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 23:57:04.572   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:04.633 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:04.733   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 23:57:05.011   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 23:57:05.136 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:05.549   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:05.591 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:05.598 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:05.757   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:05.822   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:05.822 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:05.852 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:05.94    Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:05.948   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:05.965 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:06     Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:06.084 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:06.745   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:06.836 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:06.938   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:06.994 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:07.062   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:07.123 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:07.132   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:07.148 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:08.323   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:08.512 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:08.621   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:08.679   Using cached https://files.pythonhosted.org/packages/88/e4/dd895e84b3baaa8826963ad1e8e9a8c83c8c435b602a8c47bca33d5972d6/grpcio-1.46.1-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:08.75    Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:08.776 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:08.776 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:08.809 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:08.926   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:08.927   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:08.94  Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:08.942 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.088   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.117 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 23:57:09.15    Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:09.192 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.209   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.231 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.246   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:09.285 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.34    Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.363 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:09.382   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.444   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:09.452 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.459   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.479 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.506 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:09.548   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:09.598 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.635   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.661   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:09.717   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.734 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.74  Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:09.804 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:09.837   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.852   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:09.887 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:09.937 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:09.979   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:10.046 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.053   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:10.056   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:10.076 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.081 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:10.141   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:10.182 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.217   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:10.233   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.236 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:10.298 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:10.348   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.351   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:10.366 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.379 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 23:57:10.427   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:10.453 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 23:57:10.5     Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.554   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 23:57:10.562   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.641 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.728   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 23:57:10.753 Installing collected packages: pyyaml, semver, dill, protobuf, six, grpcio, pulumi, arpeggio, attrs, parver, urllib3, idna, certifi, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 23:57:10.802 Installing collected packages: dill, six, protobuf, grpcio, pyyaml, semver, pulumi, attrs, arpeggio, parver, charset-normalizer, urllib3, certifi, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 23:57:10.915 Installing collected packages: protobuf, six, grpcio, pyyaml, dill, semver, pulumi, attrs, arpeggio, parver, certifi, charset-normalizer, urllib3, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 23:57:11.941   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 23:57:11.989   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 23:57:12.13    Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 23:57:14.31      Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 23:57:14.314     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 23:57:14.38  Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 23:57:14.383 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoDbDepl... 🏁 23:57:14.394     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀 🔥 🚀 prepareDemoFronte... 🏁 23:57:14.423 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 23:57:14.423 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 23:57:14.426 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 23:57:14.426 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:14.459 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.1 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoDbDepl... 🏁 23:57:14.513 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 23:57:14.513 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 23:57:14.743 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 23:57:14.743 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:14.905 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 23:57:14.905 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 dependencies.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 23:57:15.094 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 for this case.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Usage:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Aliases:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Flags:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 23:57:15.095       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096 
💀    🚀 prepareDemoFronte... 🏁 23:57:15.096 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 23:57:15.097 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.098 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.099 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:15.102 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 23:57:15.399 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 23:57:15.479 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 23:57:15.555 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 23:57:15.628 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 23:57:16.296 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 23:57:16.398 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 23:57:16.398 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  dependencies.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46  
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46      dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46        repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.46      - name: memcached
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461     dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 23:57:16.461       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 for this case.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 Usage:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 Aliases:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 Flags:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 23:57:16.462       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463 
💀    🚀 prepareDemoBacken... 🏁 23:57:16.463 Use "helm dependency [command] --help" for more information about a command.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 23:57:16.745 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 23:57:16.906 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 23:57:17.99  Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 23:57:18.075 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 23:57:18.489 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:18.521 
💀    🚀 deployDemoFronten... 🏁 23:57:19.016  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:19.086  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 23:57:19.182 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 23:57:19.454  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 23:57:19.454  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 23:57:19.802  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:19.804  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 23:57:19.809  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:19.81   +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoBackend... 🏁 23:57:19.837 
💀    🚀 deployDemoFronten... 🏁 23:57:19.947  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 23:57:19.947  
💀    🚀 deployDemoFronten... 🏁 23:57:19.947 Resources:
💀    🚀 deployDemoFronten... 🏁 23:57:19.947     + 4 to create
💀    🚀 deployDemoFronten... 🏁 23:57:19.947 
💀    🚀 deployDemoFronten... 🏁 23:57:19.947 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001  
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.001 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 23:57:20.38   +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 23:57:20.468  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoFronten... 🏁 23:57:20.487 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:20.578 
💀    🚀 deployDemoBackend... 🏁 23:57:20.827  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 23:57:20.829  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 23:57:20.833  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 23:57:20.983  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 23:57:20.983  
💀    🚀 deployDemoBackend... 🏁 23:57:20.983 Resources:
💀    🚀 deployDemoBackend... 🏁 23:57:20.983     + 5 to create
💀    🚀 deployDemoBackend... 🏁 23:57:20.983 
💀    🚀 deployDemoBackend... 🏁 23:57:20.983 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 23:57:20.989  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.053  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.09   +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.158  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.463  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.468  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.49   +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.493  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 23:57:21.508  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 23:57:21.508  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.553  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.556  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.571  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.574  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.577  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.583  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoBackend... 🏁 23:57:21.628 
💀    🚀 deployDemoFronten... 🏁 23:57:21.697  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 23:57:21.697  
💀    🚀 deployDemoFronten... 🏁 23:57:21.699 Outputs:
💀    🚀 deployDemoFronten... 🏁 23:57:21.699     app: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699         ready    : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.699             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699         ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.699         resources: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.699                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                               }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                               spec      : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   replicas: 1
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   selector: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       matchLabels: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                           app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                           app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                   template: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       metadata: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                           labels: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                               app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                               app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                           }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                       spec    : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                           containers        : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                               [0]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                   env            : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                       [0]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                           name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                           value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                       }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                       [1]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                           name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                       }
💀    🚀 deployDemoFronten... 🏁 23:57:21.7                                                       [2]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                 ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                         ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701 
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                     creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                     labels            : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.701                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                         [0]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.702                                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             time       : "2022-05-13T16:57:21Z"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     resource_version  : "178353"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     uid               : "c42b7682-bb0a-4d4a-982d-999e90d38804"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                 spec       : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     replicas                 : 1
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     selector                 : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         match_labels: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                     template                 : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                         metadata: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             labels: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                                 app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                                 app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.703                             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                         spec    : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                             containers                      : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                 [0]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     env                       : [
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         [0]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             value: "http://localhost:3000"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         [1]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         [2]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         [3]: {
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                             value: "1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁 23:57:21.704                                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             ]
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                         }
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                     }
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 }
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705             }
💀    🚀 deployDemoFronten... 🏁 23:57:21.705             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.705                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 23:57:21.705        
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.753  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.753  
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754     app: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754         ]
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.754                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                             }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.755                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                             }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.756                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                                 }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                             }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757 
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                     creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.757                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                 }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                             }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.758                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 23:57:21.759         
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoBackend... 🏁 23:57:22.076  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.171  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.517  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.518  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.53   +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.539  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.543  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.546  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 23:57:22.548  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 23:57:22.566  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 23:57:22.574  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 23:57:22.794  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 23:57:22.794  
💀    🚀 deployDemoBackend... 🏁 23:57:22.796 Outputs:
💀    🚀 deployDemoBackend... 🏁 23:57:22.797     app: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797         ready    : [
💀    🚀 deployDemoBackend... 🏁 23:57:22.797             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797         ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.797         resources: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                 }
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.797                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                             }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                             spec      : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 selector: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                 template: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                         labels: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                         }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.798                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "false"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.799 
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_ALLOW_METHODS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.799 
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: (json) [
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                             [0]: "*"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.799 
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: (json) []
💀    🚀 deployDemoBackend... 🏁 23:57:22.799 
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_CORS_MAX_AGE"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "600"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.799                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "guest"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [16]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "3000"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [17]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [18]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [19]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [20]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [21]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [22]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "local"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [23]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "localhost"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [24]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [25]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "root"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [26]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "/"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [27]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [28]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "root"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [29]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [30]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [31]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "root"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [32]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "root"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [33]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "local"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [34]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [35]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [36]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                           value: "/static"
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       }
💀    🚀 deployDemoBackend... 🏁 23:57:22.8                                                       [37]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 23:57:22.801                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [44]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [45]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [46]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [47]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [48]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     [49]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 23:57:22.802                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                     [50]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                                 ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                             }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                         ]
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                                 }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                             }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803 
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     creation_timestamp: "2022-05-13T16:57:22Z"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     generation        : 1
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     labels            : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                         [0]: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.803                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                 }
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                         }
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     }
💀    🚀 deployDemoBackend... 🏁 23:57:22.804                                     f:t
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 23:57:22.913 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 25.811528753s
         Current Time: 23:57:23
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 25.923615934s
         Current Time: 23:57:23
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.607µs
         Current Time: 23:57:23
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.472 🚧 Install pip packages.
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoBacken... 🏁 23:57:23.476 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 23:57:23.477 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 23:57:23.902 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.903 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.91  Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.911 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.911 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.917 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.918 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.918 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.92  Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.92  Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.926 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.927 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.927 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.929 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.929 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.93  Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.932 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.933 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.934 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.935 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.941 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.941 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.943 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.943 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.943 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.945 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.95  Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.951 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.953 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.969 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:23.972 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.977 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 23:57:23.99  Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 23:57:23.996 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 23:57:24.012 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.029 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.044 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.049 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 23:57:24.053 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.058 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.058 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 23:57:24.058 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoFronte... 🏁 23:57:24.064 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 23:57:24.084 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 23:57:24.107 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 23:57:24.125 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 23:57:24.125 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 23:57:24.131 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 23:57:24.131 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 23:57:24.132 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 23:57:24.139 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 23:57:24.142 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoBacken... 🏁 23:57:24.171 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 23:57:24.171 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.463 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"http://localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 23:57:24.463 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.547 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.547 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.547 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.547 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 dependencies.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     dependencies:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 for this case.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 Usage:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 Aliases:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.548 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 Flags:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 
💀    🚀 prepareDemoFronte... 🏁 23:57:24.549 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 23:57:24.551 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.599 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.599 🚧 Prepare chart dependencies.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.689 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69      - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  for this case.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  Usage:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69    helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69  Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.69    dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.691       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692 
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 23:57:24.692 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 23:57:26.147 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 23:57:26.245 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_CORS_ALLOW_CREDENTIALS","value":"false"},{"name":"APP_CORS_ALLOW_HEADERS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_METHODS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGINS","value":"[\"*\"]"},{"name":"APP_CORS_ALLOW_ORIGIN_REGEX","value":""},{"name":"APP_CORS_EXPOSE_HEADERS","value":"[]"},{"name":"APP_CORS_MAX_AGE","value":"600"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 23:57:26.245 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 dependencies.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324     dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.324       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325     dependencies:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 for this case.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Usage:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Aliases:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Flags:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 
💀    🚀 prepareDemoBacken... 🏁 23:57:26.325 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 23:57:26.327 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 23:57:26.64  Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 23:57:26.77  
💀    🚀 destroyDemoFronte... 🏁 23:57:26.77   -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.772  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.774  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.775  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.776  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.776  
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777 Outputs:
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777   - app: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777         ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777       - resources: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.777                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                               - template: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                                 ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.778                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                         ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                     ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - resource_version  : "178353"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - uid               : "c42b7682-bb0a-4d4a-982d-999e90d38804"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.779                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                            - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                          }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                            - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                            - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                          }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                            - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                            - value: "1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                          }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                      ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                  }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                              ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                          }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                      }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                  }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78              }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78            - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                  }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                  }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                              }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                          }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78  
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                      }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                    - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                    - labels            : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                        - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                      }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                    - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                    -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                      }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                    - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                      }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                                  }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                              }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.78                            - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                     ]
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                   - resource_version  : "178352"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                   - uid               : "0f0d5214-f455-4515-bc1e-4aa47499c78c"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781             }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781         }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781     }
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781 Resources:
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.781 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.861 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 23:57:26.912 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.916  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 23:57:26.922  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.986 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.988  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.99   -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.993  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.995  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.996  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.996  
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.997 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998         ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.998                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:26.999                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                   - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                   - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                             ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                           - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27     
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                       - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                       - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                       - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                           - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                       - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                               - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                               - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                   - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                       - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                       - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                   - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                       - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                           - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                       - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                           - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                               - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                           - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                               - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                   - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                       - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                           - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27                                                           - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.001                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                     ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - resource_version  : "178369"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - uid               : "f199521b-a11b-4270-ba70-bd7d3219cd9f"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.002                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                             ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.003               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                   - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.004                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                     ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                   - resource_version  : "178368"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                   - uid               : "a6cb195b-8ec5-44e1-9730-b132d862cb3f"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.005 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 23:57:27.047  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.048  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.049  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.053  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.066  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.066  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.066  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.07   
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071 Outputs:
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071   - app: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071         ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071       - resources: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.071                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - template: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                                 ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                         ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                   - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.072                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                     ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - resource_version  : "178353"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - uid               : "c42b7682-bb0a-4d4a-982d-999e90d38804"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.073                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - value: "http://localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                     ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                             ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.074                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                   - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.075                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                                     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                     ]
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   - resource_version  : "178352"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                   - uid               : "0f0d5214-f455-4515-bc1e-4aa47499c78c"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076                 }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076             }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076         }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076     }
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076 Resources:
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076     - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.076 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 23:57:27.077 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.077 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 23:57:27.077 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.125 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.125  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.125  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.227  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.227  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.228  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.231  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.233  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.24   -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.24   -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.24   
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.242 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.242   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.242       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.242       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243         ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.243                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.244                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.245                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246 
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                   - creation_timestamp: "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.246                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - time       : "2022-05-13T16:57:21Z"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                     ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - resource_version  : "178369"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - uid               : "f199521b-a11b-4270-ba70-bd7d3219cd9f"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                 }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                     }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                             }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.247                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                         }
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 23:57:27.248                                   - image_pull_policy         : "IfNotPre
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 23:57:28.118 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 23:57:28.252 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.253  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.258  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.258  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.262  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.265  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.267  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.267  
💀    🚀 destroyDemoBacken... 🏁 23:57:28.271 Outputs:
💀    🚀 destroyDemoBacken... 🏁 23:57:28.271   - app: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.271       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.271       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272       - resources: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.272                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                               - template: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.273 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.274                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.275                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.276                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                 ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                                 ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.277                                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                   - creation_timestamp: "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.278                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_ALLOW_CREDENTIALS"}     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_ALLOW_HEADERS"}         : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_ALLOW_METHODS"}         : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_ALLOW_ORIGINS"}         : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_ALLOW_ORIGIN_REGEX"}    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_EXPOSE_HEADERS"}        : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_CORS_MAX_AGE"}               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.279                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                        - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.28                                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.281                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.282                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - time       : "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - resource_version  : "178385"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - uid               : "1e37dc8d-4924-4172-9df7-953c2b5ca8ed"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.283                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "false"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                             ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                             ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                             ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name: "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.284                                           - value: "600"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.285                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.286                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.287                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                   -     [44]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.288                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [45]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [46]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [47]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [48]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [49]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   -     [50]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.289                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                            - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                            - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                            - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                      ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                    - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                    - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                  }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                              ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                  }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29              }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29            - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                    - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                            - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                - annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.29                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                 ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.291                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                   - creation_timestamp: "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.292                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - time       : "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - resource_version  : "178387"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - uid               : "6fd7738f-f969-4706-8654-74c90a2968f6"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - cluster_ip             : "10.98.203.53"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   -     [0]: "10.98.203.53"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293               - status     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.293           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                   - creation_timestamp: "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.294                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                           - time       : "2022-05-13T16:57:22Z"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                     ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                   - resource_version  : "178384"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                   - uid               : "9cbb2af6-dc9a-4683-97b1-8c8bae82eda9"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295         }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295 Resources:
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.295 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 23:57:28.372 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.374  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.374  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.379  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.481  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.493  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.494  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.499  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.503  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.508  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.515  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.521  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.521  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.521  
💀    🚀 destroyDemoBacken... 🏁 23:57:28.525 Outputs:
💀    🚀 destroyDemoBacken... 🏁 23:57:28.526   - app: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.526       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528       - resources: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.528                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                             }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                 }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                               - template: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.529                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                          }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                    - spec    : {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                        - containers        : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                        -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                - env            : [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - value: "30"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                        - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                      }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.53                                                -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_ALLOW_CREDENTIALS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: "false"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_ALLOW_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_ALLOW_METHODS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_ALLOW_ORIGINS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: (json) [
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       -     [0]: "*"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                         ]
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_ALLOW_ORIGIN_REGEX"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_EXPOSE_HEADERS"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: (json) []
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531 
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_CORS_MAX_AGE"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: "600"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.531                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                     }
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 23:57:28.532                                         
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 23:57:28.639 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 5.290563444s
         Current Time: 23:57:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.40159629s
         Current Time: 23:57:28
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

