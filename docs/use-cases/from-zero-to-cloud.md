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
         Elapsed Time: 1.16µs
         Current Time: 08:03:13
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 08:03:13.026 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 08:03:13.03  🎉🎉🎉
💀    🚀 initProject          🚧 08:03:13.031 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 116.120276ms
         Current Time: 08:03:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 317.952147ms
         Current Time: 08:03:13
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 2.869µs
         Current Time: 08:03:13
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:03:13.525         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:03:13.525     
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:03:13.525   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:03:13.525   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:03:13.525   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:03:13.525 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 08:03:14.009 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 08:03:14.009 Preparing base variables
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Base variables prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing start command
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Start command prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing test command
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Test command prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing check command
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Check command prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.125 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 08:03:14.429 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 08:03:14.439 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 08:03:14.447 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 08:03:14.447 ✅ Validate
💀    🚀 makeMysqlApp         🐬 08:03:14.447 Validate app directory
💀    🚀 makeMysqlApp         🐬 08:03:14.447 Done validating app directory
💀    🚀 makeMysqlApp         🐬 08:03:14.447 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 08:03:14.452 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 08:03:14.452 Validate template locations
💀    🚀 makeMysqlApp         🐬 08:03:14.464 Done validating template locations
💀    🚀 makeMysqlApp         🐬 08:03:14.464 Validate app ports
💀    🚀 makeMysqlApp         🐬 08:03:14.468 Done validating app ports
💀    🚀 makeMysqlApp         🐬 08:03:14.468 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 08:03:14.473 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 08:03:14.473 🚧 Generate
💀    🚀 makeMysqlApp         🐬 08:03:14.473 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 08:03:14.473   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 08:03:14.473 ]
💀    🚀 makeMysqlApp         🐬 08:03:14.473 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 08:03:14.491 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 08:03:14.491 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 08:03:14.491 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 08:03:14.943 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 08:03:14.943 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.156 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.157 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.46  Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.469 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.478 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.478 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.478 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.478 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.478 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.482 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.482 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.504 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.504 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.508 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.508 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.512 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.512 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.512 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.512   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.513   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.513   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.513 ]
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.513 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.561 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.567 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.572 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.76  Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:03:15.953 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.141 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.323 Checking start
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.328 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.512 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.695 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.699 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:16.893 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.075 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.256 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.446 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.451 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.631 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.812 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:17.815 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.004 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.188 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.192 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.379 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.566 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.571 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.749 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.941 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.946 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 08:03:18.946 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.531862349s
         Current Time: 08:03:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.643683641s
         Current Time: 08:03:19
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 5.014µs
         Current Time: 08:03:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:03:19.36  Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:03:19.363         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:03:19.363     
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:03:19.363   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:03:19.363   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:03:19.363   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:03:19.363 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 08:03:19.839 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 08:03:19.839 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.044 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 08:03:20.352 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:03:20.363 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:03:20.374 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 08:03:20.375 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 08:03:20.375 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 08:03:20.375 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 08:03:20.375 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:03:20.38  Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:03:20.38  Validate template locations
💀    🚀 makeFastApiApp       ⚡ 08:03:20.394 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 08:03:20.394 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 08:03:20.399 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 08:03:20.399 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403 ]
💀    🚀 makeFastApiApp       ⚡ 08:03:20.403 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 08:03:21.042 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 08:03:21.044 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 08:03:21.044 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 08:03:21.501 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:03:21.501 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.626 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.968 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.979 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.988 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.988 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.988 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.988 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.988 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.992 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:03:22.992 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.014 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.014 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.02  Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.02  Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 ]
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.024 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.085 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.089 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.093 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.291 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.294 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.486 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.675 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.68  Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 08:03:23.872 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.053 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.057 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.251 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.445 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.449 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.63  Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.819 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:24.824 Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.018 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.21  Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.408 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.59  Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.594 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.777 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.965 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:25.968 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.159 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.347 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.352 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.545 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.728 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.732 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:03:26.926 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:03:27.112 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:03:27.31  Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:03:27.503 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:27.685 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:03:27.873 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:03:28.062 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:28.25  Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:03:28.451 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 08:03:28.451 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 08:03:28.946 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 08:03:28.946 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing start command
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Start command prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing test command
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Test command prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing check command
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Check command prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.408 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 08:03:30.732 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 08:03:30.741 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 08:03:30.75  Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 08:03:30.75  ✅ Validate
💀    🚀 addFastApiModule     ⚡ 08:03:30.751 Validate app directory
💀    🚀 addFastApiModule     ⚡ 08:03:30.751 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 08:03:30.751 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 08:03:30.754 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 08:03:30.754 Validate template locations
💀    🚀 addFastApiModule     ⚡ 08:03:30.769 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 08:03:30.769 Validate app ports
💀    🚀 addFastApiModule     ⚡ 08:03:30.773 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 08:03:30.773 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 08:03:30.777 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 08:03:30.778 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 08:03:30.778 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 08:03:30.778   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 08:03:30.778 ]
💀    🚀 addFastApiModule     ⚡ 08:03:30.778 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 08:03:30.798 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 08:03:30.798 Registering module
💀    🚀 addFastApiModule     ⚡ 08:03:30.839 Done registering module
💀    🚀 addFastApiModule     ⚡ 08:03:30.841 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 08:03:30.841 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 08:03:31.389 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 08:03:31.389 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:32.85  Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:33.199 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:33.209 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:33.22  Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.22  Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:03:33.24  Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:03:33.24  Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:03:33.385 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:03:33.385 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:03:33.526 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:03:33.527 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 08:03:33.678 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 08:03:33.678 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:33.782 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:34.233 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:34.244 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:03:34.254 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:03:34.254 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 08:03:34.254 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 08:03:34.254 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 08:03:34.254 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:03:34.26  Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:03:34.26  Validate template locations
💀    🚀 addFastApiCrud       ⚡ 08:03:34.277 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 08:03:34.277 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 08:03:34.281 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 08:03:34.281 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:03:34.287 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288 ]
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288 
💀    🚀 addFastApiCrud       ⚡ 08:03:34.288 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 08:03:34.329 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 08:03:34.329 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:03:34.381 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:03:34.381 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:03:34.455 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:03:34.455 Registering repo
💀    🚀 addFastApiCrud       ⚡ 08:03:34.539 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 08:03:34.539 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 08:03:34.539 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 15.297481077s
         Current Time: 08:03:34
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 15.498178975s
         Current Time: 08:03:34
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.197µs
         Current Time: 08:03:35
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:03:35.044 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:03:35.048 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:03:35.048 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:03:35.048 
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:03:35.048         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:03:35.049         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:03:35.049         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:03:35.049     
💀    🚀 zrbShowAdv           ☕ 08:03:35.049 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:03:35.049 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:03:35.049   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:03:35.049   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:03:35.049   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:03:35.049 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 08:03:35.633 🧰 Prepare
💀    🚀 makeNginxApp         📗 08:03:35.634 Preparing base variables
💀    🚀 makeNginxApp         📗 08:03:35.801 Base variables prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing start command
💀    🚀 makeNginxApp         📗 08:03:35.802 Start command prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing prepare command
💀    🚀 makeNginxApp         📗 08:03:35.802 Prepare command prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing test command
💀    🚀 makeNginxApp         📗 08:03:35.802 Test command prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing migrate command
💀    🚀 makeNginxApp         📗 08:03:35.802 Migrate command prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing check command
💀    🚀 makeNginxApp         📗 08:03:35.802 Check command prepared
💀    🚀 makeNginxApp         📗 08:03:35.802 Preparing replacement map
💀    🚀 makeNginxApp         📗 08:03:36.267 Add config to replacement map
💀    🚀 makeNginxApp         📗 08:03:36.277 Add env to replacement map
💀    🚀 makeNginxApp         📗 08:03:36.285 Replacement map prepared
💀    🚀 makeNginxApp         📗 08:03:36.285 ✅ Validate
💀    🚀 makeNginxApp         📗 08:03:36.285 Validate app directory
💀    🚀 makeNginxApp         📗 08:03:36.285 Done validating app directory
💀    🚀 makeNginxApp         📗 08:03:36.285 Validate app container volumes
💀    🚀 makeNginxApp         📗 08:03:36.29  Done validating app container volumes
💀    🚀 makeNginxApp         📗 08:03:36.29  Validate template locations
💀    🚀 makeNginxApp         📗 08:03:36.303 Done validating template locations
💀    🚀 makeNginxApp         📗 08:03:36.303 Validate app ports
💀    🚀 makeNginxApp         📗 08:03:36.308 Done validating app ports
💀    🚀 makeNginxApp         📗 08:03:36.308 Validate app crud fields
💀    🚀 makeNginxApp         📗 08:03:36.314 Done validating app crud fields
💀    🚀 makeNginxApp         📗 08:03:36.314 🚧 Generate
💀    🚀 makeNginxApp         📗 08:03:36.314 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 08:03:36.315   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 08:03:36.315 ]
💀    🚀 makeNginxApp         📗 08:03:36.315 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 08:03:36.343 🔩 Integrate
💀    🚀 makeNginxApp         📗 08:03:36.343 🎉🎉🎉
💀    🚀 makeNginxApp         📗 08:03:36.343 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 08:03:36.871 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 08:03:36.871 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 08:03:37.019 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing start command
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Start command prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing test command
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Test command prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing check command
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Check command prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.02  Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 08:03:37.727 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 08:03:37.746 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 08:03:37.762 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 08:03:37.762 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 08:03:37.763 Validate app directory
💀    🚀 makeNginxAppRunner   📗 08:03:37.763 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 08:03:37.763 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 08:03:37.775 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 08:03:37.775 Validate template locations
💀    🚀 makeNginxAppRunner   📗 08:03:37.811 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 08:03:37.811 Validate app ports
💀    🚀 makeNginxAppRunner   📗 08:03:37.819 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 08:03:37.819 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 08:03:37.829 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 08:03:37.829 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 08:03:37.83  🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 08:03:37.83    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 08:03:37.83    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 08:03:37.83  ]
💀    🚀 makeNginxAppRunner   📗 08:03:37.83  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 08:03:37.904 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 08:03:37.912 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 08:03:37.918 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:03:38.151 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:03:38.382 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:03:38.59  Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:03:38.809 Checking start
💀    🚀 makeNginxAppRunner   📗 08:03:38.812 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 08:03:39.052 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:03:39.343 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 08:03:39.346 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 08:03:39.549 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:03:39.76  Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:03:39.974 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:03:40.186 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 08:03:40.19  Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 08:03:40.399 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:03:40.627 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 08:03:40.632 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 08:03:40.896 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:03:41.094 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 08:03:41.097 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 08:03:41.296 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:03:41.509 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 08:03:41.514 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 08:03:41.751 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:03:41.981 Done registering app runner tasks
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 7.04968756s
         Current Time: 08:03:42
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 7.161055692s
         Current Time: 08:03:42
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.638µs
         Current Time: 08:03:42
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:03:42.421 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:03:42.53  Synchronize task environments
💀    🚀 syncEnv              🔄 08:03:42.751 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:03:42.95  🎉🎉🎉
💀    🚀 syncEnv              🔄 08:03:42.95  Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 636.935281ms
         Current Time: 08:03:43
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 838.632332ms
         Current Time: 08:03:43
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 8.63µs
         Current Time: 08:03:43
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:03:43.762 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:03:43.763 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoBackendI... 🏭 08:03:44.04  Build image demo-backend:latest
💀    🚀 buildDemoFrontend... 🏭 08:03:44.04  Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 08:03:44.042 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:03:52.144 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:03:52.146 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 08:03:52.259 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:03:52.259  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:03:52.26  Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 08:03:52.263 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:03:52.266  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:03:52.266 Step 2/11 : USER 0
💀    🚀 buildDemoDbImage     🏭 08:03:52.271 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:03:52.271  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.271  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.272  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 08:03:52.273 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoBackendI... 🏭 08:03:52.275 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 08:03:52.275 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:03:52.276 Successfully built 40621c693b70
💀    🚀 buildDemoDbImage     🏭 08:03:52.278 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:03:52.278 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:03:52.283 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:03:52.289 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:03:52.289 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoFrontend... 🏭 08:03:52.289 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 08:03:52.289  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:03:52.289 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29   ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29  Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29   ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 08:03:52.29  Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 08:03:52.291  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.291  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:03:52.291 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 08:03:52.295  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.295  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:03:52.295 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 08:03:52.303  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.303  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 08:03:52.303 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:03:52.306  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.306  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:03:52.307 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:03:52.315 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:03:52.321 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:03:52.321 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 08:03:52.434 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 8.781458151s
         Current Time: 08:03:52
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 8.892623978s
         Current Time: 08:03:52
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.754µs
         Current Time: 08:03:53
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 updateProjectLinks   🔗 08:03:53.049 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:03:53.049 Links updated
💀    🚀 prepareDemoBackend   🔧 08:03:53.051 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 08:03:53.108 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 08:03:53.176 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 08:03:53.322 Build image demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:03:57.616 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:03:57.618 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 08:03:57.696 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:03:57.696  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:03:57.696 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:03:57.701 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:03:57.706 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:03:57.706 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:03:57.749 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75   ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75  Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75   ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 08:03:57.75  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 08:03:57.751 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:03:57.755  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.755  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 08:03:57.755 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:03:57.756 Successfully built 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:03:57.76  Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:03:57.767 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:03:57.767 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 prepareDemoBackend   🔧 08:03:58.037 Activate venv
💀    🚀 prepareDemoBackend   🔧 08:03:58.038 Install dependencies
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:03:58.296 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:03:58.314 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 08:03:58.489 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 08:03:58.532 🐳 Retrieve previous log of 'demoDb'
💀    🚀 prepareDemoBackend   🔧 08:03:58.729 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 08:03:59.055   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:03:59.063 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 08:03:59.382   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:03:59.394 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.72 
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.72 Welcome to the Bitnami nginx container
💀    🚀 startDemoFrontend... 📗 08:03:59.513 
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.72 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.73 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.73 
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 [38;5;6mnginx [38;5;5m00:23:16.75 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 2022/05/11 00:23:16 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 08:03:59.513 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoFrontend... 📗 08:03:59.514 🐳 Starting container 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.017311Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.561 [38;5;6mmysql [38;5;5m00:23:23.69 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 [38;5;6mmysql [38;5;5m00:23:29.70 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.020297Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.020305Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 [38;5;6mmysql [38;5;5m00:23:29.72 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 [38;5;6mmysql [38;5;5m00:23:31.74 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.026741Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.139970Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.324017Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀 🔥 🚀 startDemoDbContainer 🐬 08:03:59.562 [38;5;6mmysql [38;5;5m00:23:31.80 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.324064Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.345040Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:23:32.345499Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:24:08.135554Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:24:10.136993Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 08:03:59.562 2022-05-11T00:24:10.862380Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 08:03:59.563 🐳 Starting container 'demoDb'
💀    🚀 prepareDemoBackend   🔧 08:03:59.606   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀 🔥 🚀 startDemoFrontend... 📗 08:04:00.644 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 08:04:00.644 Error: failed to start containers: demoFrontend
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
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=29775)
💀 🔪 Kill 🔧 'prepareDemoBackend' command (PID=27157)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=29786)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=29787)
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:00.795 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:00.795 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 🚀 prepareDemoBackend   🔧 08:04:01.05  ERROR: Operation cancelled by user
💀 🔥 🚀 prepareDemoBackend   🔧 08:04:01.055 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 08:04:01.055 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 Error running 🔧 'prepareDemoBackend' command: exit status 1
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 8.613220053s
         Current Time: 08:04:01
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["start"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.577µs
         Current Time: 08:04:01
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:04:01.798 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:04:01.798 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 08:04:01.819 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 08:04:01.912 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoFrontend... 🏭 08:04:02.06  Build image demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 08:04:02.06  Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 08:04:03.37  Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:04:03.371 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 08:04:03.415 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:04:03.415  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:04:03.415 Successfully built 188ba73f5790
💀    🚀 buildDemoBackendI... 🏭 08:04:03.417 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoDbImage     🏭 08:04:03.419 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421  ---> 562078b73ebf
💀    🚀 buildDemoDbImage     🏭 08:04:03.421 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:04:03.421 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:04:03.421 Step 4/11 : USER 1001
💀    🚀 buildDemoBackendI... 🏭 08:04:03.422 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 08:04:03.422  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:04:03.422 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 08:04:03.423 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:04:03.424 Step 6/9 : COPY . .
💀    🚀 buildDemoFrontend... 🏭 08:04:03.425  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.425  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:04:03.425 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:04:03.426  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.426  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 08:04:03.426 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:04:03.427  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.427  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 08:04:03.427 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.428  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 08:04:03.429 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43   ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43  Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43   ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:04:03.43  Successfully built 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:04:03.434 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:04:03.436 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:04:03.436 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.437  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 08:04:03.438 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:04:03.438  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:04:03.438  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:04:03.439 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:04:03.442 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:04:03.444 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:04:03.444 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:04:03.777 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:04:03.789 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 08:04:03.847 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 08:04:03.868 🐳 Retrieve previous log of 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.875 [38;5;6mnginx [38;5;5m00:23:16.72 
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.875 [38;5;6mnginx [38;5;5m00:23:16.72 Welcome to the Bitnami nginx container
💀    🚀 startDemoFrontend... 📗 08:04:04.875 
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 [38;5;6mnginx [38;5;5m00:23:16.72 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 [38;5;6mnginx [38;5;5m00:23:16.73 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 [38;5;6mnginx [38;5;5m00:23:16.73 
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 [38;5;6mnginx [38;5;5m00:23:16.75 [38;5;2mINFO  ==> ** Starting NGINX **
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 2022/05/11 00:23:16 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 08:04:04.876 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoFrontend... 📗 08:04:04.879 🐳 Starting container 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.017311Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 [38;5;6mmysql [38;5;5m00:23:23.69 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 [38;5;6mmysql [38;5;5m00:23:29.70 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.020297Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.020305Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.026741Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.139970Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 [38;5;6mmysql [38;5;5m00:23:29.72 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.324017Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.324064Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 [38;5;6mmysql [38;5;5m00:23:31.74 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:04.903 [38;5;6mmysql [38;5;5m00:23:31.80 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 08:04:04.903 2022-05-11T00:23:32.345040Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 08:04:04.904 2022-05-11T00:23:32.345499Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 08:04:04.904 2022-05-11T00:24:08.135554Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 08:04:04.904 2022-05-11T00:24:10.136993Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 08:04:04.904 2022-05-11T00:24:10.862380Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 08:04:04.905 🐳 Starting container 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 08:04:05.992 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 08:04:05.992 Error: failed to start containers: demoFrontend
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
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=32510)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=32526)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=32527)
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:06.152 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 08:04:06.152 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 5.198560937s
         Current Time: 08:04:06
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["startContainers"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.153µs
         Current Time: 08:04:07
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:04:07.155 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:04:07.155 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoDbContainer  ✋ 08:04:07.499 Docker container demoDb is not running
💀    🚀 stopDemoFrontendC... ✋ 08:04:07.501 Docker container demoFrontend is not running
💀    🚀 stopDemoBackendCo... ✋ 08:04:07.503 Docker container demoBackend is not running
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 08:04:07.61  
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 560.990206ms
         Current Time: 08:04:07
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 672.047503ms
         Current Time: 08:04:07
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.099µs
         Current Time: 08:04:07
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:07.99  Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:04:07.993         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:04:07.993     
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:04:07.993   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:04:07.993   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:04:07.993   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:04:07.993 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.286 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.286 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.444 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.444 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.445 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.673 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.679 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.686 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.686 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.686 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.686 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.686 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.689 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.689 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.699 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.699 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.702 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.702 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.705 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.738 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.738 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:08.738 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.141 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.141 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.289 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.289 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.29  Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.518 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.525 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.533 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.533 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.533 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.533 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.533 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.536 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.536 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.547 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.548 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.551 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.551 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.556 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.586 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.59  Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.595 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.764 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.925 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:09.929 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.086 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.251 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.255 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.417 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.417 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:10.417 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.534797324s
         Current Time: 08:04:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.645338092s
         Current Time: 08:04:10
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.244µs
         Current Time: 08:04:10
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:10.793 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:04:10.795         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:04:10.795     
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:04:10.795   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:04:10.795   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:04:10.795   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:04:10.795 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.224 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.225 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:11.987 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.215 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.221 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.228 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.228 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.228 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.228 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.228 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.231 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.231 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.241 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.241 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.245 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.245 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.249 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.282 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.283 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.283 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.648 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:12.648 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.706 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.707 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.974 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.981 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.988 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.988 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.988 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.988 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.988 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.993 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:13.993 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.003 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.003 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.007 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.007 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01  ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.01  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.031 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.035 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.041 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.215 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.375 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.378 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.54  Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.689 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.692 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.841 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.841 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:14.841 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.15389878s
         Current Time: 08:04:14
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.265315075s
         Current Time: 08:04:15
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.017µs
         Current Time: 08:04:15
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:15.204 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:04:15.206         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:04:15.206     
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:04:15.206   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:04:15.206   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:04:15.206   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:04:15.206 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.63  🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.63  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.733 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.734 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.734 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.734 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.942 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.948 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.957 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.957 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.957 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.957 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.957 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.96  Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.96  Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.969 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.969 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.972 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.972 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:15.975 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.004 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.004 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.004 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.443 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.443 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.549 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.756 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.763 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.769 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.769 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.769 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.769 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.769 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.772 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.772 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.781 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.781 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.784 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.784 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.787 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.806 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.809 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.812 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:16.971 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.128 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.132 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.292 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.451 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.454 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.61  Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.61  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:04:17.61  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.513036331s
         Current Time: 08:04:17
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.623985393s
         Current Time: 08:04:17
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.15µs
         Current Time: 08:04:17
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:17.989 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:04:18.094 Synchronize task environments
💀    🚀 syncEnv              🔄 08:04:18.264 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:04:18.424 🎉🎉🎉
💀    🚀 syncEnv              🔄 08:04:18.424 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 542.161659ms
         Current Time: 08:04:18
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 742.813927ms
         Current Time: 08:04:18
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.162µs
         Current Time: 08:04:18
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:18.881 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 08:04:18.992 🎉🎉🎉
💀    🚀 setProjectValue      🔗 08:04:18.992 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 216.785279ms
         Current Time: 08:04:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.866577ms
         Current Time: 08:04:19
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.061µs
         Current Time: 08:04:19
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:04:19.455 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 08:04:19.566 🎉🎉🎉
💀    🚀 setProjectValue      🔗 08:04:19.566 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 217.565711ms
         Current Time: 08:04:19
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 418.693991ms
         Current Time: 08:04:19
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.026µs
         Current Time: 08:04:20
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 08:04:20.046 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:20.046 🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 08:04:20.047 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 08:04:21.867 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:21.882 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:04:21.903 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:22.139 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:22.161 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:22.174 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:23.373   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:23.393 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:23.591   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:23.605 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:23.958   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 08:04:24.071   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 08:04:24.091   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:24.109 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:24.192 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:24.284 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:24.404   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 08:04:24.557   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:24.562 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:24.643 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:24.909   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:24.933 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.055   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:25.065   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:25.085 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.115 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:25.128   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:25.137 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:25.331   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:25.351 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.416   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.421 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:25.526   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:25.531 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.609   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.614 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.855   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:25.866 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:26.063   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:26.153 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:26.256   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:26.285 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:04:26.435   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:26.454 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:04:26.456   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:26.519 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:26.791   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:26.806   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:26.808 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:26.811 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:26.844   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:26.852 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:27.016   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:27.028 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:27.07    Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:27.079 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:27.141   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:27.146 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:27.348   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:27.36  Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:27.428   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:27.443 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.29    Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.3   Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.305   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.311   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:28.323 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.344 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.405   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.409   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.413   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.427 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.433 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.435 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.517   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.544 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.557   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:28.564 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.649   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:28.652   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.655   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.661 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.662 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.667 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.737   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.744 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.757   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.761   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:04:28.774 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:04:28.783 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.911   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.916   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:28.917 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.917   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:28.924 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:04:28.977 Installing collected packages: six, protobuf, dill, grpcio, pyyaml, semver, pulumi, arpeggio, attrs, parver, charset-normalizer, idna, urllib3, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 08:04:29.008   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:29.018 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:04:29.034   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:04:29.114   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:04:29.121 Installing collected packages: six, grpcio, semver, dill, protobuf, pyyaml, pulumi, attrs, arpeggio, parver, charset-normalizer, certifi, idna, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 08:04:29.202 Installing collected packages: protobuf, pyyaml, semver, six, grpcio, dill, pulumi, attrs, arpeggio, parver, urllib3, charset-normalizer, idna, certifi, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 08:04:29.544   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 08:04:29.71    Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 08:04:29.759   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 08:04:30.771     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 08:04:30.815 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 08:04:30.854 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:04:30.854 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:30.935     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 08:04:30.976 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 08:04:31.011 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 08:04:31.011 🚧 Prepare chart dependencies.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:04:31.012 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:04:31.012 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.025     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06  
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06      dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.06        repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.061 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Usage:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:04:31.062 🚧 Preparation completed.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.065 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoBacken... 🏁 08:04:31.092 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:04:31.092 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.196 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.196 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.238 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.239   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24    list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24    -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  
💀    🚀 prepareDemoDbDepl... 🏁 08:04:31.24  Use "helm dependency [command] --help" for more information about a command.
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🔥 🚀 deployDemoFronten... 🏁 08:04:31.366 error: no stack named 'dev' found
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 deployDemoFronten... 🏁 08:04:31.461 Created stack 'dev'
💀 🔥 🚀 deployDemoDbDeplo... 🏁 08:04:31.565 error: no stack named 'dev' found
💀    🚀 deployDemoDbDeplo... 🏁 08:04:31.659 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 08:04:31.878 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:04:31.929 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 08:04:31.93  🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:04:31.979       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98      dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  for this case.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  Usage:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98  Flags:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.98    -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981 
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 08:04:31.981 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 08:04:32.19  error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 08:04:32.292 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 08:04:33.112 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:04:33.246 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 08:04:33.452 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:33.618 
💀    🚀 deployDemoFronten... 🏁 08:04:33.79   +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:04:33.86  Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 08:04:33.861  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:33.969  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.029  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 08:04:34.1    +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 08:04:34.103  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 08:04:34.213  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:04:34.213  
💀    🚀 deployDemoFronten... 🏁 08:04:34.213 Resources:
💀    🚀 deployDemoFronten... 🏁 08:04:34.213     + 4 to create
💀    🚀 deployDemoFronten... 🏁 08:04:34.213 
💀    🚀 deployDemoFronten... 🏁 08:04:34.214 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 08:04:34.257 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.288  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.29   +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38   +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38   
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38  Resources:
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38      + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38  
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.38  Updating (dev):
💀    🚀 deployDemoFronten... 🏁 08:04:34.571 
💀    🚀 deployDemoBackend... 🏁 08:04:34.594  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:04:34.654  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:34.737 
💀    🚀 deployDemoBackend... 🏁 08:04:34.856  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:04:34.857  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:04:34.861  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoFronten... 🏁 08:04:34.893  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 08:04:34.959  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:34.985  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:04:34.985  
💀    🚀 deployDemoBackend... 🏁 08:04:34.985 Resources:
💀    🚀 deployDemoBackend... 🏁 08:04:34.985     + 5 to create
💀    🚀 deployDemoBackend... 🏁 08:04:34.985 
💀    🚀 deployDemoBackend... 🏁 08:04:34.985 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.066  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.125  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 08:04:35.19   +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:04:35.191  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:04:35.208  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:04:35.209  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:04:35.214  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 08:04:35.214  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 08:04:35.358  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 08:04:35.358  
💀    🚀 deployDemoFronten... 🏁 08:04:35.359 Outputs:
💀    🚀 deployDemoFronten... 🏁 08:04:35.359     app: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.359         ready    : [
💀    🚀 deployDemoFronten... 🏁 08:04:35.359             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.359             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36          ]
💀    🚀 deployDemoFronten... 🏁 08:04:35.36          resources: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36              apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                  api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                  id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                  kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                  metadata   : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                      annotations       : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                          kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                              apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                              kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                              metadata  : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                  annotations: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                  }
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                  labels     : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                      helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                  }
💀    🚀 deployDemoFronten... 🏁 08:04:35.36                                  name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                             spec      : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                 selector: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                 }
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                 template: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                         labels: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 08:04:35.361                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                 ]
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                         ]
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                                 }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362 
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     labels            : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.362                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                         [0]: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                 }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.363                                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                                 }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                         }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                     }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                                 }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                             }
💀    🚀 deployDemoFronten... 🏁 08:04:35.364                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 08:04:35.364       
💀    🚀 deployDemoBackend... 🏁 08:04:35.376 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.381  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.383  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.393  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.396  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.4    +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.401  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.512  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.512  
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513     app: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.513                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.514                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515 
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.515                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                             time       : "2022-05-11T01:04:35Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.516                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     resource_version  : "14718"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     uid               : "2659e91b-4e54-4e05-a4d4-d3cc2860a247"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.517                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.518                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:04:35.519                                             name :
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoBackend... 🏁 08:04:35.734  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 08:04:35.799  +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.078  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.078  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.083  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.092  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.092  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.095  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:04:36.1    +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:04:36.1    +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:04:36.112  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:04:36.275  +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 08:04:36.275  
💀    🚀 deployDemoBackend... 🏁 08:04:36.277 Outputs:
💀    🚀 deployDemoBackend... 🏁 08:04:36.278     app: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278         ready    : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.278             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278         ]
💀    🚀 deployDemoBackend... 🏁 08:04:36.278         resources: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 08:04:36.278                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                             spec      : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 selector: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                 template: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                         labels: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.279                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "/token/"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [4]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [5]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [6]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [7]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "10"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [8]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "guest"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [9]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "3000"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [10]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [11]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [12]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [13]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [14]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [15]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "local"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [16]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [17]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [18]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [19]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "/"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [20]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [21]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [22]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [23]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      [24]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                          value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.28                                                      }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                 ]
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.281                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                 ]
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         ]
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282 
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     creation_timestamp: "2022-05-11T01:04:36Z"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     generation        : 1
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     labels            : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                         [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                     f:template               : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         f:metadata: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                             f:labels: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                         f:spec    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                             f:containers                   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                 k:{"name":"demo-backend"}: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                     f:env                     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.282                                                         k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.283                                                         k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                     f:ports                   : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             operation  : "Update"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             time       : "2022-05-11T01:04:36Z"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     ]
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     name              : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     namespace         : "default"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     resource_version  : "14734"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     uid               : "18e3c890-12b6-4137-b66d-23e15dac152d"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                 }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                 spec       : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     progress_deadline_seconds: 600
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     replicas                 : 1
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     revision_history_limit   : 10
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     selector                 : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         match_labels: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     strategy                 : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         rolling_update: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             max_surge      : "25%"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             max_unavailable: "25%"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         type          : "RollingUpdate"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                     template                 : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         metadata: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             labels: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                 app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                                 app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                             }
💀    🚀 deployDemoBackend... 🏁 08:04:36.284                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                         spec    : {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                             containers                      : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                 [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                     env                       : [
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [0]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "HS256"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [1]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "30"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [2]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [3]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "/token/"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [4]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [5]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [6]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [7]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "10"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                         [8]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.285                                             value: "guest"
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                         [9]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                             name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                             value: "3000"
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.286                                         [10]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [11]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [12]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [13]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [14]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [15]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "local"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [16]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [17]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [18]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                             value: "root"
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         }
💀    🚀 deployDemoBackend... 🏁 08:04:36.287                                         [19]: {
💀    🚀 deployDemoBackend... 🏁 08:04:36.287     
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 deploy               🏭 08:04:36.394 
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 16.463745317s
         Current Time: 08:04:36
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 16.57425463s
         Current Time: 08:04:36
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.175µs
         Current Time: 08:04:36
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 08:04:36.83  🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 08:04:36.831 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:04:36.832 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 08:04:37.138 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.138 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.139 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.145 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.145 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.147 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.152 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.153 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.154 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.156 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.157 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.158 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.158 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.163 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.165 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.165 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.166 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.166 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.167 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.168 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.169 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.169 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.17  Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.171 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.171 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.172 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.175 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.183 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.184 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.187 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.202 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.204 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.205 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.213 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.242 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.243 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.249 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.252 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.253 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.253 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.254 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.256 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.256 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.258 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.26  Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoBacken... 🏁 08:04:37.262 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.29  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 08:04:37.292 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:04:37.302 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:04:37.302 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:04:37.311 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:04:37.311 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:04:37.316 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:04:37.316 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.567 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 08:04:37.567 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.633 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.634 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      - name: memcached
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  for this case.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Usage:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Flags:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64    -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  
💀    🚀 prepareDemoFronte... 🏁 08:04:37.64  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641 
💀    🚀 prepareDemoFronte... 🏁 08:04:37.641 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:04:37.642 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.708 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.709       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71  
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.71  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 08:04:37.713 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoBacken... 🏁 08:04:38.581 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:04:38.659 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 08:04:38.659 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.734 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.734 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.734 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 for this case.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.735 Usage:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 Flags:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 08:04:38.736 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 08:04:39.504 Previewing destroy (dev):
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.513 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 08:04:39.62  
💀    🚀 destroyDemoFronte... 🏁 08:04:39.621  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.623  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.626  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.628  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.628  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.628  
💀    🚀 destroyDemoFronte... 🏁 08:04:39.629 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63    - app: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63        - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63        -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63        -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63          ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63        - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63            - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                            - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                  }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                    - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                  }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                              }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                            - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.63                                - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                               - template: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                                 ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                         ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.631                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                     ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - resource_version  : "14701"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - uid               : "2e6a7669-a110-4e49-936f-01d263fb0371"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.632                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.633 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                             ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.633               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.634  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.634                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                     ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                   - resource_version  : "14702"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                   - uid               : "84c7e26f-1db0-4250-926b-e1757a7668e2"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.636     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.637 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.637 Resources:
💀    🚀 destroyDemoFronte... 🏁 08:04:39.637     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 08:04:39.637 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.637 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.637  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.637  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.639  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.64   -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.64   
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.641 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.642                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.643                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - resource_version  : "14718"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - uid               : "2659e91b-4e54-4e05-a4d4-d3cc2860a247"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.644                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                             ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.645 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - resource_version  : "14717"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                   - uid               : "b52237ec-530d-4e68-9aba-174088a77f1d"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646       - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.646 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.647     - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.647 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.647 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.739 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.74  
💀    🚀 destroyDemoFronte... 🏁 08:04:39.744  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.746  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.752  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.754  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.881  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.881  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.889  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.889  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.889  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.89   -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.89   -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894  
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894   - app: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.894         ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895       - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                               - template: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                                 ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                         ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.895                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                     ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - resource_version  : "14701"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - uid               : "2e6a7669-a110-4e49-936f-01d263fb0371"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.896                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.897                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                             ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.898                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                             }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                         }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                     }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.899                                 }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                               }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                             - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                             - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                             - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                           }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                       ]
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                     - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                     - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                     - resource_version  : "14702"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                     - uid               : "84c7e26f-1db0-4250-926b-e1757a7668e2"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                   }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9                 - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9               }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9           }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9         - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9       }
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   Resources:
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9       - 4 deleted
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 08:04:39.9   If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.903  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.903  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.91   -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.911  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.911  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.911  
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.912               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.913                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.914                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.915                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                   - resource_version  : "14718"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                   - uid               : "2659e91b-4e54-4e05-a4d4-d3cc2860a247"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.916               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.917                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                             ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918 
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                   - creation_timestamp: "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.918                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                           - time       : "2022-05-11T01:04:35Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                   - resource_version  : "14717"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                   - uid               : "b52237ec-530d-4e68-9aba-174088a77f1d"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919             }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919         }
💀    🚀 destroyDemoDbDepl... 🏁 08:04:39.919   
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 08:04:40.564 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 08:04:40.642 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.645  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.645  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.651  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.655  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.657  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.659  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.659  
💀    🚀 destroyDemoBacken... 🏁 08:04:40.661 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662         ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.662                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.663                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.664                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.665                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.666                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.667                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                   - creation_timestamp: "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.668                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.669                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                        - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.67                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                   - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                       - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.671                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                           - time       : "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - resource_version  : "14734"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - uid               : "18e3c890-12b6-4137-b66d-23e15dac152d"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.672                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.673                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.674                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.675                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.676                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.677                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.678                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.679                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                                    - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                                  }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                              ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                          }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                      }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                  }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68              }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68            - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                    - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                            - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                                - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.68                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                       - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                       - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                       - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                       - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                   - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                               - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.681                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                   - creation_timestamp: "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.682                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                   - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                       - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.683                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - time       : "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - resource_version  : "14736"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - uid               : "add505c5-3067-4ea7-bb4d-cce21ee4610b"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - cluster_ip             : "10.104.21.206"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   -     [0]: "10.104.21.206"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - status     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.684                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                   - creation_timestamp: "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.685                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                           - time       : "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                     ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                   - resource_version  : "14733"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                   - uid               : "ecce5f54-a59d-4334-a41f-a14302fe8ded"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686 Resources:
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.686 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 08:04:40.748 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.749  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.749  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.756  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.843  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.844  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.849  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.855  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.856  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.858  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.869  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.875  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.875  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.876  
💀    🚀 destroyDemoBacken... 🏁 08:04:40.882 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:04:40.882   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.882       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.882       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883         ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.883                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.884                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.885                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.886                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.887                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888 
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                   - creation_timestamp: "2022-05-11T01:04:36Z"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                 }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                     }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                             }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.888                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:04:40.889                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERN
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 08:04:40.997 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 4.286657192s
         Current Time: 08:04:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.397713817s
         Current Time: 08:04:41
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

