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
         Elapsed Time: 1.288µs
         Current Time: 08:14:54
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 08:14:54.687 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 08:14:54.691 🎉🎉🎉
💀    🚀 initProject          🚧 08:14:54.691 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 113.333935ms
         Current Time: 08:14:54
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 314.502955ms
         Current Time: 08:14:54
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.744µs
         Current Time: 08:14:55
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:14:55.136 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:14:55.139         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:14:55.139     
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:14:55.139   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:14:55.139   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:14:55.139   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:14:55.139 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 08:14:55.573 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 08:14:55.573 Preparing base variables
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Base variables prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing start command
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Start command prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing test command
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Test command prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing check command
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Check command prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.658 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 08:14:55.875 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 08:14:55.881 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 08:14:55.888 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 08:14:55.888 ✅ Validate
💀    🚀 makeMysqlApp         🐬 08:14:55.889 Validate app directory
💀    🚀 makeMysqlApp         🐬 08:14:55.889 Done validating app directory
💀    🚀 makeMysqlApp         🐬 08:14:55.889 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 08:14:55.893 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 08:14:55.893 Validate template locations
💀    🚀 makeMysqlApp         🐬 08:14:55.903 Done validating template locations
💀    🚀 makeMysqlApp         🐬 08:14:55.903 Validate app ports
💀    🚀 makeMysqlApp         🐬 08:14:55.907 Done validating app ports
💀    🚀 makeMysqlApp         🐬 08:14:55.907 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 08:14:55.91  Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 08:14:55.91  🚧 Generate
💀    🚀 makeMysqlApp         🐬 08:14:55.91  🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 08:14:55.91    "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 08:14:55.91  ]
💀    🚀 makeMysqlApp         🐬 08:14:55.91  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 08:14:55.924 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 08:14:55.924 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 08:14:55.924 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.301 🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.301 Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.447 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.448 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.448 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.448 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.448 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.703 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.711 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.718 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.718 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.718 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.718 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.718 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.722 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.722 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.74  Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.74  Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.743 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.743 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.746 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.746 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.746 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.746   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.747   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.747   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.747 ]
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.747 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.788 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.791 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.796 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:14:56.971 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.183 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.37  Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.561 Checking start
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.565 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.748 Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.917 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:57.92  Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.084 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.251 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.417 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.599 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.603 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.764 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.937 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:58.94  Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.1   Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.252 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.256 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.419 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.581 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.585 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.732 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.88  Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.884 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 08:14:59.884 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.853977565s
         Current Time: 08:14:59
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.964254139s
         Current Time: 08:15:00
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.11µs
         Current Time: 08:15:00
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:15:00.25  Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:15:00.252 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:15:00.252 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:15:00.252 
💀    🚀 zrbShowAdv           ☕ 08:15:00.252         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:15:00.252         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:15:00.252         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:15:00.252         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:15:00.252         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:15:00.253         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:15:00.253     
💀    🚀 zrbShowAdv           ☕ 08:15:00.253 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:15:00.253 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:15:00.253   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:15:00.253   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:15:00.253   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:15:00.253 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 08:15:00.702 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 08:15:00.702 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 08:15:00.856 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:00.857 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 08:15:01.094 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:15:01.102 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 08:15:01.108 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 08:15:01.108 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 08:15:01.108 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 08:15:01.108 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 08:15:01.108 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:15:01.111 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 08:15:01.111 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 08:15:01.122 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 08:15:01.122 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 08:15:01.126 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 08:15:01.126 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13  Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13  🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13  🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13    "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13  ]
💀    🚀 makeFastApiApp       ⚡ 08:15:01.13  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 08:15:01.663 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 08:15:01.663 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 08:15:01.664 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.135 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.135 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.934 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.935 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:02.935 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.163 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.17  Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.176 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.176 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.177 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.177 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.177 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.18  Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.18  Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.195 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.195 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.198 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.198 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 ]
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.201 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.246 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.25  Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.253 Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.403 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.406 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.555 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.706 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.71  Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 08:15:03.863 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.016 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.019 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.177 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.329 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.332 Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.483 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.637 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.64  Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.79  Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:04.94  Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.093 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.248 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.251 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.401 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.556 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.559 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.706 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.858 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:15:05.861 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.016 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.17  Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.173 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.325 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.478 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.635 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.788 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:06.943 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 08:15:07.091 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 08:15:07.244 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:07.398 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 08:15:07.553 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 08:15:07.553 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 08:15:08.044 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 08:15:08.044 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Preparing start command
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Start command prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Preparing test command
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Test command prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.122 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 08:15:09.123 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.123 Preparing check command
💀    🚀 addFastApiModule     ⚡ 08:15:09.123 Check command prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.123 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 08:15:09.4   Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 08:15:09.408 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 08:15:09.416 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 08:15:09.416 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 08:15:09.416 Validate app directory
💀    🚀 addFastApiModule     ⚡ 08:15:09.416 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 08:15:09.416 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 08:15:09.42  Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 08:15:09.42  Validate template locations
💀    🚀 addFastApiModule     ⚡ 08:15:09.435 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 08:15:09.435 Validate app ports
💀    🚀 addFastApiModule     ⚡ 08:15:09.441 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 08:15:09.441 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 08:15:09.446 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 08:15:09.446 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 08:15:09.446 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 08:15:09.446   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 08:15:09.446 ]
💀    🚀 addFastApiModule     ⚡ 08:15:09.446 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 08:15:09.469 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 08:15:09.469 Registering module
💀    🚀 addFastApiModule     ⚡ 08:15:09.503 Done registering module
💀    🚀 addFastApiModule     ⚡ 08:15:09.504 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 08:15:09.504 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 08:15:09.847 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 08:15:09.847 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:10.933 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:11.296 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:11.309 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:11.32  Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.32  Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:15:11.335 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 08:15:11.335 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:15:11.427 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 08:15:11.427 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:15:11.513 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 08:15:11.513 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 08:15:11.65  Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 08:15:11.651 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:15:11.736 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 08:15:11.736 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 08:15:11.736 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:11.737 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:11.99  Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:11.998 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 08:15:12.005 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 08:15:12.005 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 08:15:12.005 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 08:15:12.005 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 08:15:12.005 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:15:12.008 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 08:15:12.008 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 08:15:12.019 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 08:15:12.019 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 08:15:12.022 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 08:15:12.022 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 ]
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 
💀    🚀 addFastApiCrud       ⚡ 08:15:12.026 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 08:15:12.055 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 08:15:12.055 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:15:12.094 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 08:15:12.094 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:15:12.149 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 08:15:12.149 Registering repo
💀    🚀 addFastApiCrud       ⚡ 08:15:12.21  Done registering repo
💀    🚀 addFastApiCrud       ⚡ 08:15:12.211 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 08:15:12.211 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 12.067998915s
         Current Time: 08:15:12
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.269376435s
         Current Time: 08:15:12
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.423µs
         Current Time: 08:15:12
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:15:12.667 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:15:12.669         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:15:12.669     
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:15:12.669   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:15:12.669   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:15:12.669   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:15:12.669 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 08:15:13.113 🧰 Prepare
💀    🚀 makeNginxApp         📗 08:15:13.113 Preparing base variables
💀    🚀 makeNginxApp         📗 08:15:13.192 Base variables prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing start command
💀    🚀 makeNginxApp         📗 08:15:13.193 Start command prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing prepare command
💀    🚀 makeNginxApp         📗 08:15:13.193 Prepare command prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing test command
💀    🚀 makeNginxApp         📗 08:15:13.193 Test command prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing migrate command
💀    🚀 makeNginxApp         📗 08:15:13.193 Migrate command prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing check command
💀    🚀 makeNginxApp         📗 08:15:13.193 Check command prepared
💀    🚀 makeNginxApp         📗 08:15:13.193 Preparing replacement map
💀    🚀 makeNginxApp         📗 08:15:13.412 Add config to replacement map
💀    🚀 makeNginxApp         📗 08:15:13.419 Add env to replacement map
💀    🚀 makeNginxApp         📗 08:15:13.426 Replacement map prepared
💀    🚀 makeNginxApp         📗 08:15:13.426 ✅ Validate
💀    🚀 makeNginxApp         📗 08:15:13.426 Validate app directory
💀    🚀 makeNginxApp         📗 08:15:13.426 Done validating app directory
💀    🚀 makeNginxApp         📗 08:15:13.426 Validate app container volumes
💀    🚀 makeNginxApp         📗 08:15:13.429 Done validating app container volumes
💀    🚀 makeNginxApp         📗 08:15:13.429 Validate template locations
💀    🚀 makeNginxApp         📗 08:15:13.438 Done validating template locations
💀    🚀 makeNginxApp         📗 08:15:13.438 Validate app ports
💀    🚀 makeNginxApp         📗 08:15:13.441 Done validating app ports
💀    🚀 makeNginxApp         📗 08:15:13.441 Validate app crud fields
💀    🚀 makeNginxApp         📗 08:15:13.445 Done validating app crud fields
💀    🚀 makeNginxApp         📗 08:15:13.445 🚧 Generate
💀    🚀 makeNginxApp         📗 08:15:13.445 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 08:15:13.445   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 08:15:13.445 ]
💀    🚀 makeNginxApp         📗 08:15:13.445 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 08:15:13.465 🔩 Integrate
💀    🚀 makeNginxApp         📗 08:15:13.465 🎉🎉🎉
💀    🚀 makeNginxApp         📗 08:15:13.465 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 08:15:13.946 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 08:15:13.946 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 08:15:14.051 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing start command
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Start command prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing test command
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Test command prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing check command
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Check command prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.052 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 08:15:14.301 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 08:15:14.308 Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 08:15:14.315 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 08:15:14.315 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 08:15:14.315 Validate app directory
💀    🚀 makeNginxAppRunner   📗 08:15:14.315 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 08:15:14.315 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 08:15:14.319 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 08:15:14.319 Validate template locations
💀    🚀 makeNginxAppRunner   📗 08:15:14.338 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 08:15:14.338 Validate app ports
💀    🚀 makeNginxAppRunner   📗 08:15:14.342 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 08:15:14.342 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 08:15:14.347 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 08:15:14.347 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 08:15:14.347 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 08:15:14.347   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 08:15:14.347   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 08:15:14.347 ]
💀    🚀 makeNginxAppRunner   📗 08:15:14.347 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"8080:80\n443","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["8080:80","443"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"8080:80\", \"443\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 08:15:14.379 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 08:15:14.383 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 08:15:14.387 Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:15:14.566 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:15:14.734 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:15:14.897 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:15:15.102 Checking start
💀    🚀 makeNginxAppRunner   📗 08:15:15.106 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 08:15:15.28  Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:15:15.439 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 08:15:15.442 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 08:15:15.598 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 08:15:15.76  Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:15:15.92  Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:15:16.083 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 08:15:16.086 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 08:15:16.249 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 08:15:16.418 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 08:15:16.421 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 08:15:16.58  Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:15:16.742 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 08:15:16.746 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 08:15:16.907 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:15:17.068 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 08:15:17.073 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 08:15:17.232 Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 08:15:17.401 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 08:15:17.405 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 08:15:17.405 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.844775418s
         Current Time: 08:15:17
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.955509587s
         Current Time: 08:15:17
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["8080:80", "443"]' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.491µs
         Current Time: 08:15:17
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:15:17.773 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:15:17.879 Synchronize task environments
💀    🚀 syncEnv              🔄 08:15:18.052 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:15:18.217 🎉🎉🎉
💀    🚀 syncEnv              🔄 08:15:18.217 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 550.73624ms
         Current Time: 08:15:18
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 752.632304ms
         Current Time: 08:15:18
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.278µs
         Current Time: 08:15:18
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:15:18.674 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:15:18.674 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoFrontend... 🏭 08:15:18.935 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 08:15:18.935 Build image demo-db:latest
💀    🚀 buildDemoBackendI... 🏭 08:15:18.936 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 08:15:21.011 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:15:21.013 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoBackendI... 🏭 08:15:21.095 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoBackendI... 🏭 08:15:21.1   Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 08:15:21.101 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:15:21.102 Step 6/9 : COPY . .
💀    🚀 buildDemoDbImage     🏭 08:15:21.103 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:15:21.103  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:15:21.103 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:15:21.107 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:15:21.108 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109 Step 4/11 : USER 1001
💀    🚀 buildDemoDbImage     🏭 08:15:21.109 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:15:21.109 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:15:21.109 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:15:21.11   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.11   ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 08:15:21.11  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 08:15:21.111 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:15:21.112 Successfully built 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:15:21.115 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoBackendI... 🏭 08:15:21.116  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.116  ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 08:15:21.116 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:15:21.117  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.117  ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 08:15:21.117 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:15:21.117  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.117  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 08:15:21.118 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoFrontend... 🏭 08:15:21.118 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:15:21.118 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 08:15:21.123  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:15:21.123  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:15:21.125 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 08:15:21.128 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:15:21.131 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:15:21.131 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 🏭 'buildImages' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 buildImages          🏭 08:15:21.238 
💀 🎉 Successfully running 🏭 'buildImages' command
💀 🔎 Job Running...
         Elapsed Time: 2.670837832s
         Current Time: 08:15:21
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.781943976s
         Current Time: 08:15:21
zaruba please buildImages -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.713µs
         Current Time: 08:15:21
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:15:21.676 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:15:21.676 Links updated
💀    🚀 prepareDemoBackend   🔧 08:15:21.677 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 08:15:21.715 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 08:15:21.794 Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 08:15:21.945 Build image demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 08:15:23.934 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 08:15:23.935 Sending build context to Docker daemon  14.85kB
💀    🚀 buildDemoDbImage     🏭 08:15:24.012 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:15:24.012  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:15:24.012 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:15:24.016 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:15:24.018 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:15:24.018 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02  Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02   ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02  Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02   ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02  Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.02   ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021 Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021  ---> 1347440dac6a
💀    🚀 buildDemoFrontend... 🏭 08:15:24.021 Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> a31b560cf951
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022 Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> 2da3ba665444
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022 Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022  ---> 0084068762a4
💀    🚀 buildDemoFrontend... 🏭 08:15:24.022 Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> 20cca1eb6764
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> 7bb05f6d9d8b
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:15:24.023  ---> 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:15:24.024 Successfully built 40621c693b70
💀    🚀 buildDemoFrontend... 🏭 08:15:24.028 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:15:24.03  🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:15:24.03  Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:15:24.333 🔎 Waiting docker container 'demoFrontend' running status
💀 🔥 🔎 startDemoFrontend... 📗 08:15:24.376 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoFrontend... 📗 08:15:24.384 Error: No such container: demoFrontend
💀 🔥 🚀 startDemoFrontend... 📗 08:15:24.42  Error: No such container: demoFrontend
💀    🚀 startDemoFrontend... 📗 08:15:24.426 🐳 Creating and starting container 'demoFrontend'
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:15:24.468 🔎 Waiting docker container 'demoDb' running status
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:24.518 Error: No such container: demoDb
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:24.521 Error: No such container: demoDb
💀    🚀 startDemoFrontend... 📗 08:15:24.531 666e18f7201af30be741810c12764244293d01451f8b9f9c89dabc504ab89876
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:24.568 Error: No such container: demoDb
💀    🚀 startDemoDbContainer 🐬 08:15:24.574 🐳 Creating and starting container 'demoDb'
💀    🚀 startDemoDbContainer 🐬 08:15:24.653 a2945a04df5aaf67e5cff2919ffb94da2380b53028b9da7267481fea4e0569e1
💀    🚀 prepareDemoBackend   🔧 08:15:24.688 Activate venv
💀    🚀 prepareDemoBackend   🔧 08:15:24.689 Install dependencies
💀    🚀 prepareDemoBackend   🔧 08:15:25.224 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 08:15:25.541   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:25.56  Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 08:15:25.66    Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:25.671 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 08:15:25.798   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 startDemoFrontend... 📗 08:15:26.47  🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 08:15:26.472 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🚀 startDemoFrontend... 📗 08:15:26.499 
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.499 [38;5;6mnginx [38;5;5m01:15:26.48 
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.499 [38;5;6mnginx [38;5;5m01:15:26.48 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.499 [38;5;6mnginx [38;5;5m01:15:26.48 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.499 [38;5;6mnginx [38;5;5m01:15:26.48 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.499 [38;5;6mnginx [38;5;5m01:15:26.48 
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.505 [38;5;6mnginx [38;5;5m01:15:26.50 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🔎 startDemoFrontend... 📗 08:15:26.506 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 08:15:26.506 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 08:15:26.507 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 08:15:26.507 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 08:15:26.509 🔎 Host port '443' is ready
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.528 2022/05/11 01:15:26 [warn] 13#13: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀 🔥 🚀 startDemoFrontend... 📗 08:15:26.528 nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /opt/bitnami/nginx/conf/nginx.conf:2
💀    🚀 startDemoDbContainer 🐬 08:15:26.664 🐳 Logging 'demoDb'
💀    🔎 startDemoDbContainer 🐬 08:15:26.667 🔎 Waiting docker container 'demoDb' healthcheck
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.66 
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.67 Welcome to the Bitnami mysql container
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.67 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-mysql
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.67 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-mysql/issues
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.67 
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.67 [38;5;2mINFO  ==> ** Starting MySQL setup **
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.71 [38;5;2mINFO  ==> Validating settings in MYSQL_*/MARIADB_* env vars
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.724 [38;5;6mmysql [38;5;5m01:15:26.71 [38;5;2mINFO  ==> Initializing mysql database
💀    🔎 startDemoDbContainer 🐬 08:15:26.728 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 08:15:26.728 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 08:15:26.73  🔎 Host port '3306' is ready
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.733 [38;5;6mmysql [38;5;5m01:15:26.73 [38;5;2mINFO  ==> Updating 'my.cnf' with custom configuration
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.738 [38;5;6mmysql [38;5;5m01:15:26.73 [38;5;2mINFO  ==> Setting user option
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.749 [38;5;6mmysql [38;5;5m01:15:26.74 [38;5;2mINFO  ==> Setting slow_query_log option
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.755 [38;5;6mmysql [38;5;5m01:15:26.75 [38;5;2mINFO  ==> Setting long_query_time option
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:26.761 [38;5;6mmysql [38;5;5m01:15:26.76 [38;5;2mINFO  ==> Installing database
💀    🔎 startDemoFrontend... 📗 08:15:29.512 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoFrontend... 📗 08:15:29.701 check demoFrontend
💀    🔎 startDemoFrontend... 📗 08:15:29.711 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 08:15:29.733 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🚀 prepareDemoBackend   🔧 08:15:29.749 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:29.9   mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:29.9   ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 08:15:29.92    Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:29.945 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀    🚀 prepareDemoBackend   🔧 08:15:30.053   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:30.065 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀    🚀 prepareDemoBackend   🔧 08:15:30.154   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:30.187 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 08:15:30.297   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:30.309 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 08:15:30.506   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:30.62  Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 prepareDemoBackend   🔧 08:15:31.073   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:31.181 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 08:15:31.397   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:31.464 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀    🚀 prepareDemoBackend   🔧 08:15:32.005   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:15:32.355 Collecting greenlet==1.1.1 (from -r requirements.txt (line 13))
💀    🚀 prepareDemoBackend   🔧 08:15:32.631   Using cached https://files.pythonhosted.org/packages/32/7a/85cbb3374bef5cac1a2eebec9f6ff324a6758970c38a2825a7b89a2e9aec/greenlet-1.1.1-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:32.649 Collecting h11==0.12.0 (from -r requirements.txt (line 14))
💀    🚀 prepareDemoBackend   🔧 08:15:32.718   Using cached https://files.pythonhosted.org/packages/60/0f/7a0eeea938eaf61074f29fed9717f2010e8d0e0905d36b38d3275a1e4622/h11-0.12.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:32.728 Collecting idna==3.3 (from -r requirements.txt (line 15))
💀    🚀 prepareDemoBackend   🔧 08:15:32.882   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:32.891 Collecting jsons==1.5.1 (from -r requirements.txt (line 16))
💀    🚀 prepareDemoBackend   🔧 08:15:33.004   Using cached https://files.pythonhosted.org/packages/04/b9/7e174aeb2994076929ba28fb0a5191d4d6f21f02db9af00cd3a963e7f0a6/jsons-1.5.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.019 Collecting passlib==1.7.4 (from -r requirements.txt (line 17))
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:33.028 [38;5;6mmysql [38;5;5m01:15:33.02 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:33.048 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:33.05  ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 08:15:33.127   Using cached https://files.pythonhosted.org/packages/3b/a4/ab6b7589382ca3df236e03faa71deac88cae040af60c071a78d254a62172/passlib-1.7.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.166 Collecting pika==1.2.0 (from -r requirements.txt (line 18))
💀    🚀 prepareDemoBackend   🔧 08:15:33.271   Using cached https://files.pythonhosted.org/packages/f5/56/2590c41852df1212426bec3e5e312cba50170e12d083a0fb1e544a52d215/pika-1.2.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.292 Collecting pydantic==1.8.2 (from -r requirements.txt (line 19))
💀    🚀 prepareDemoBackend   🔧 08:15:33.553   Using cached https://files.pythonhosted.org/packages/ff/74/54e030641601112309f6d2af620774e9080f99c7a15742fc6a0b170c4076/pydantic-1.8.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.585 Collecting PyMySQL==1.0.2 (from -r requirements.txt (line 20))
💀    🚀 prepareDemoBackend   🔧 08:15:33.669   Using cached https://files.pythonhosted.org/packages/4f/52/a115fe175028b058df353c5a3d5290b71514a83f67078a6482cff24d6137/PyMySQL-1.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.68  Collecting pytest==6.2.5 (from -r requirements.txt (line 21))
💀    🔎 startDemoFrontend... 📗 08:15:33.713 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 08:15:33.713 📜 Task 'startDemoFrontendContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀    🚀 prepareDemoBackend   🔧 08:15:33.862   Using cached https://files.pythonhosted.org/packages/40/76/86f886e750b81a4357b6ed606b2bcf0ce6d6c27ad3c09ebf63ed674fc86e/pytest-6.2.5-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:33.896 Collecting pytest-cov==3.0.0 (from -r requirements.txt (line 22))
💀 🎉 Reach 📗 'startDemoFrontend' wrapper
💀    🚀 prepareDemoBackend   🔧 08:15:34.03    Using cached https://files.pythonhosted.org/packages/20/49/b3e0edec68d81846f519c602ac38af9db86e1e71275528b3e814ae236063/pytest_cov-3.0.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:34.046 Collecting python-jose==3.3.0 (from -r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:15:34.157   Using cached https://files.pythonhosted.org/packages/bd/2d/e94b2f7bab6773c70efc70a61d66e312e1febccd9e0db6b9e0adf58cbad1/python_jose-3.3.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:34.176 Collecting python-multipart==0.0.5 (from -r requirements.txt (line 24))
💀    🚀 prepareDemoBackend   🔧 08:15:34.285   Using cached https://files.pythonhosted.org/packages/46/40/a933ac570bf7aad12a298fc53458115cc74053474a72fbb8201d7dc06d3d/python-multipart-0.0.5.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:15:34.496 Collecting requests==2.27.1 (from -r requirements.txt (line 25))
💀    🚀 prepareDemoBackend   🔧 08:15:34.616   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:34.637 Collecting sqlalchemy==1.4.23 (from -r requirements.txt (line 26))
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:35.043 [38;5;6mmysql [38;5;5m01:15:35.04 [38;5;2mINFO  ==> Configuring authentication
💀    🚀 prepareDemoBackend   🔧 08:15:35.063   Using cached https://files.pythonhosted.org/packages/d0/6b/32b93b001ca2274ca0097a1f81be27fc69a805761049719531ac182427fe/SQLAlchemy-1.4.23-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:35.099 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:35.124 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:35.177 [38;5;6mmysql [38;5;5m01:15:35.17 [38;5;2mINFO  ==> Running mysql_upgrade
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:35.181 [38;5;6mmysql [38;5;5m01:15:35.18 [38;5;2mINFO  ==> Stopping mysql
💀    🚀 prepareDemoBackend   🔧 08:15:35.181 Collecting starlette==0.14.2 (from -r requirements.txt (line 27))
💀    🚀 prepareDemoBackend   🔧 08:15:35.297   Using cached https://files.pythonhosted.org/packages/15/34/db1890f442a1cd3a2c761f4109a0eb4e63503218d70a8c8e97faa09a5500/starlette-0.14.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:35.311 Collecting typing-extensions==3.10.0.2 (from -r requirements.txt (line 28))
💀    🚀 prepareDemoBackend   🔧 08:15:35.395   Using cached https://files.pythonhosted.org/packages/74/60/18783336cc7fcdd95dae91d73477830aa53f5d3181ae4fe20491d7fc3199/typing_extensions-3.10.0.2-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:35.402 Collecting typish==1.9.3 (from -r requirements.txt (line 29))
💀    🚀 prepareDemoBackend   🔧 08:15:35.481   Using cached https://files.pythonhosted.org/packages/9d/d6/3f56c9c0c12adf61dfcf4ed5c8ffd2c431db8dd85592067a57e8e1968565/typish-1.9.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:35.496 Collecting urllib3==1.26.8 (from -r requirements.txt (line 30))
💀    🚀 prepareDemoBackend   🔧 08:15:35.621   Using cached https://files.pythonhosted.org/packages/4e/b8/f5a25b22e803f0578e668daa33ba3701bb37858ec80e08a150bd7d2cf1b1/urllib3-1.26.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:35.643 Collecting uuid==1.30 (from -r requirements.txt (line 31))
💀    🚀 prepareDemoBackend   🔧 08:15:35.749   Using cached https://files.pythonhosted.org/packages/ce/63/f42f5aa951ebf2c8dac81f77a8edcc1c218640a2a35a03b9ff2d4aa64c3d/uuid-1.30.tar.gz
💀    🚀 prepareDemoBackend   🔧 08:15:35.909 Collecting uvicorn==0.15.0 (from -r requirements.txt (line 32))
💀    🚀 prepareDemoBackend   🔧 08:15:36.037   Using cached https://files.pythonhosted.org/packages/6f/d0/2c2f4e88d63a8f8891419ca02e029e3a7200ab8f64a3628517cf35ff0379/uvicorn-0.15.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.07  Collecting six>=1.4.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:15:36.148   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.155 Collecting cffi>=1.1 (from bcrypt==3.2.0->-r requirements.txt (line 4))
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:36.191 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:36.192 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 prepareDemoBackend   🔧 08:15:36.444   Using cached https://files.pythonhosted.org/packages/e5/fe/1dac7533ddb73767df8ba26183a9375dde2ee136aec7c92c9fb3038108e3/cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.46  Collecting packaging (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:36.55    Using cached https://files.pythonhosted.org/packages/05/8e/8de486cbd03baba4deef4142bd643a3e7bbe954a784dc1bb17142572d127/packaging-21.3-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.558 Collecting py>=1.8.2 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:36.642   Using cached https://files.pythonhosted.org/packages/f6/f0/10642828a8dfb741e5f3fbaac830550a518a775c7fff6f04a007259b0548/py-1.11.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.654 Collecting pluggy<2.0,>=0.12 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:36.752   Using cached https://files.pythonhosted.org/packages/9e/01/f38e2ff29715251cf25532b9082a1589ab7e4f571ced434f98d0139336dc/pluggy-1.0.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.764 Collecting attrs>=19.2.0 (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:36.9     Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:36.934 Collecting iniconfig (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:37.015   Using cached https://files.pythonhosted.org/packages/9b/dd/b3c12c6d707058fa947864b67f0c4e0c39ef8610988d7baea9578f3c48f3/iniconfig-1.1.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:37.021 Collecting toml (from pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:37.102   Using cached https://files.pythonhosted.org/packages/44/6f/7120676b6d73228c96e17f1f794d8ab046fc910d781c8d151120c3f1569e/toml-0.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:37.108 Collecting coverage[toml]>=5.2.1 (from pytest-cov==3.0.0->-r requirements.txt (line 22))
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:37.199 [38;5;6mmysql [38;5;5m01:15:37.19 [38;5;2mINFO  ==> Starting mysql in background
💀    🚀 prepareDemoBackend   🔧 08:15:37.579   Using cached https://files.pythonhosted.org/packages/0c/58/25b4d208e0f6f00e19440385f360dc9891f8fa5ab62c11da52eb226fd9cd/coverage-6.3.2-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_17_x86_64.manylinux2014_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 08:15:37.596 Collecting ecdsa!=0.15 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:15:37.71    Using cached https://files.pythonhosted.org/packages/4a/b6/b678b080967b2696e9a201c096dc076ad756fb35c87dca4e1d1a13496ff7/ecdsa-0.17.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:37.723 Collecting rsa (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:15:37.84    Using cached https://files.pythonhosted.org/packages/30/ab/8fd9e88e6fa5ec41afca995938bbefb72195278e0cfc5bd76a4f29b23fb2/rsa-4.8-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:37.9   Collecting pyasn1 (from python-jose==3.3.0->-r requirements.txt (line 23))
💀    🚀 prepareDemoBackend   🔧 08:15:38.018   Using cached https://files.pythonhosted.org/packages/62/1e/a94a8d635fa3ce4cfc7f506003548d0a2447ae76fd5ca53932970fe3053f/pyasn1-0.4.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:38.031 Collecting pycparser (from cffi>=1.1->bcrypt==3.2.0->-r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 08:15:38.12    Using cached https://files.pythonhosted.org/packages/62/d5/5f610ebe421e85889f2e55e33b7f9a6795bd982198517d912eb1c76e1a53/pycparser-2.21-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:38.131 Collecting pyparsing!=3.0.5,>=2.0.2 (from packaging->pytest==6.2.5->-r requirements.txt (line 21))
💀    🚀 prepareDemoBackend   🔧 08:15:38.294   Using cached https://files.pythonhosted.org/packages/6c/10/a7d0fa5baea8fe7b50f448ab742f26f52b80bfca85ac2be9d35cdd9a3246/pyparsing-3.0.9-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:38.307 Collecting tomli; extra == "toml" (from coverage[toml]>=5.2.1->pytest-cov==3.0.0->-r requirements.txt (line 22))
💀    🚀 prepareDemoBackend   🔧 08:15:38.397   Using cached https://files.pythonhosted.org/packages/97/75/10a9ebee3fd790d20926a90a2547f0bf78f371b2f13aa822c759680ca7b9/tomli-2.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 08:15:38.716 Installing collected packages: aiofiles, asgiref, avro-python3, six, pycparser, cffi, bcrypt, certifi, charset-normalizer, click, fastavro, urllib3, idna, requests, confluent-kafka, cryptography, typing-extensions, pydantic, starlette, fastapi, greenlet, h11, typish, jsons, passlib, pika, PyMySQL, pyparsing, packaging, py, pluggy, attrs, iniconfig, toml, pytest, tomli, coverage, pytest-cov, ecdsa, pyasn1, rsa, python-jose, python-multipart, sqlalchemy, uuid, uvicorn
💀    🚀 prepareDemoBackend   🔧 08:15:38.745   Running setup.py install for avro-python3: started
💀    🚀 prepareDemoBackend   🔧 08:15:39.012     Running setup.py install for avro-python3: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:15:39.379   Running setup.py install for fastavro: started
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:39.404 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:39.406 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:42.558 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:42.56  ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:45.213 [38;5;6mmysql [38;5;5m01:15:45.21 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:45.231 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:45.24  find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:45.245 [38;5;6mmysql [38;5;5m01:15:45.24 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:45.751 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:45.753 ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/opt/bitnami/mysql/tmp/mysql.sock' (2)
💀    🚀 startDemoDbContainer 🐬 08:15:47.26  
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:47.26  [38;5;6mmysql [38;5;5m01:15:47.25 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 08:15:47.282 [38;5;6mmysql [38;5;5m01:15:47.28 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 08:15:47.516 2022-05-11T01:15:47.509610Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 08:15:47.516 2022-05-11T01:15:47.511443Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 08:15:47.517 2022-05-11T01:15:47.511451Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 08:15:47.518 2022-05-11T01:15:47.517731Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 08:15:47.654 2022-05-11T01:15:47.653675Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 08:15:47.854 2022-05-11T01:15:47.853904Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 08:15:47.854 2022-05-11T01:15:47.853946Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 08:15:47.871 2022-05-11T01:15:47.870662Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 08:15:47.871 2022-05-11T01:15:47.870785Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀 🔥 🔎 startDemoDbContainer 🐬 08:15:48.895 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 Database
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 information_schema
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 mysql
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 performance_schema
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 sample
💀    🔎 startDemoDbContainer 🐬 08:15:48.903 sys
💀    🔎 startDemoDbContainer 🐬 08:15:48.908 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoDbContainer 🐬 08:15:52.911 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 08:15:52.911 📜 Task 'startDemoDbContainer' is ready
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀    🚀 prepareDemoBackend   🔧 08:16:07.688     Running setup.py install for fastavro: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:16:08.936   Running setup.py install for python-multipart: started
💀    🚀 prepareDemoBackend   🔧 08:16:09.1       Running setup.py install for python-multipart: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:16:09.678   Running setup.py install for uuid: started
💀    🚀 prepareDemoBackend   🔧 08:16:09.809     Running setup.py install for uuid: finished with status 'done'
💀    🚀 prepareDemoBackend   🔧 08:16:09.857 Successfully installed PyMySQL-1.0.2 aiofiles-0.7.0 asgiref-3.4.1 attrs-21.4.0 avro-python3-1.10.0 bcrypt-3.2.0 certifi-2021.10.8 cffi-1.15.0 charset-normalizer-2.0.12 click-8.0.1 confluent-kafka-1.8.2 coverage-6.3.2 cryptography-36.0.1 ecdsa-0.17.0 fastapi-0.68.1 fastavro-1.4.9 greenlet-1.1.1 h11-0.12.0 idna-3.3 iniconfig-1.1.1 jsons-1.5.1 packaging-21.3 passlib-1.7.4 pika-1.2.0 pluggy-1.0.0 py-1.11.0 pyasn1-0.4.8 pycparser-2.21 pydantic-1.8.2 pyparsing-3.0.9 pytest-6.2.5 pytest-cov-3.0.0 python-jose-3.3.0 python-multipart-0.0.5 requests-2.27.1 rsa-4.8 six-1.16.0 sqlalchemy-1.4.23 starlette-0.14.2 toml-0.10.2 tomli-2.0.1 typing-extensions-3.10.0.2 typish-1.9.3 urllib3-1.26.8 uuid-1.30 uvicorn-0.15.0
💀 🔥 🚀 prepareDemoBackend   🔧 08:16:09.914 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 08:16:09.914 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBackend   🔧 08:16:09.965 Prepare
💀    🚀 prepareDemoBackend   🔧 08:16:09.965 prepare command
💀    🚀 prepareDemoBackend   🔧 08:16:09.965 Preparation complete
💀 🎉 Successfully running 🔧 'prepareDemoBackend' command
💀 🏁 Run ⚡ 'startDemoBackend' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackend' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackend     ⚡ 08:16:10.176 🔎 Waiting for port '3000'
💀    🚀 startDemoBackend     ⚡ 08:16:10.176 Activate venv
💀    🚀 startDemoBackend     ⚡ 08:16:10.176 Start
💀    🚀 startDemoBackend     ⚡ 08:16:10.549 2022-05-11 08:16:10,549 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackend     ⚡ 08:16:10.549 2022-05-11 08:16:10,549 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.554 2022-05-11 08:16:10,553 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackend     ⚡ 08:16:10.554 2022-05-11 08:16:10,554 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.556 2022-05-11 08:16:10,556 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackend     ⚡ 08:16:10.556 2022-05-11 08:16:10,556 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.558 2022-05-11 08:16:10,558 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:10.559 2022-05-11 08:16:10,559 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:16:10.559 2022-05-11 08:16:10,559 INFO sqlalchemy.engine.Engine [generated in 0.00012s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 2022-05-11 08:16:10,561 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 CREATE TABLE books (
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	title VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	author VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	synopsis VARCHAR(255), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 )
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 
💀    🚀 startDemoBackend     ⚡ 08:16:10.562 2022-05-11 08:16:10,562 INFO sqlalchemy.engine.Engine [no key 0.00012s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.585 2022-05-11 08:16:10,585 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_synopsis ON books (synopsis)
💀    🚀 startDemoBackend     ⚡ 08:16:10.585 2022-05-11 08:16:10,585 INFO sqlalchemy.engine.Engine [no key 0.00016s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.604 2022-05-11 08:16:10,604 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_author ON books (author)
💀    🚀 startDemoBackend     ⚡ 08:16:10.604 2022-05-11 08:16:10,604 INFO sqlalchemy.engine.Engine [no key 0.00018s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.62  2022-05-11 08:16:10,620 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_title ON books (title)
💀    🚀 startDemoBackend     ⚡ 08:16:10.62  2022-05-11 08:16:10,620 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.638 2022-05-11 08:16:10,638 INFO sqlalchemy.engine.Engine CREATE INDEX ix_books_id ON books (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.639 2022-05-11 08:16:10,638 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.653 2022-05-11 08:16:10,653 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:16:10.654 2022-05-11 08:16:10,654 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:10.654 2022-05-11 08:16:10,654 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:16:10.654 2022-05-11 08:16:10,654 INFO sqlalchemy.engine.Engine [cached since 0.09587s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 2022-05-11 08:16:10,656 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 CREATE TABLE roles (
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	name VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	json_permissions VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 )
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 
💀    🚀 startDemoBackend     ⚡ 08:16:10.656 2022-05-11 08:16:10,656 INFO sqlalchemy.engine.Engine [no key 0.00008s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.678 2022-05-11 08:16:10,678 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_name ON roles (name)
💀    🚀 startDemoBackend     ⚡ 08:16:10.678 2022-05-11 08:16:10,678 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.696 2022-05-11 08:16:10,696 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_json_permissions ON roles (json_permissions)
💀    🚀 startDemoBackend     ⚡ 08:16:10.696 2022-05-11 08:16:10,696 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.712 2022-05-11 08:16:10,712 INFO sqlalchemy.engine.Engine CREATE INDEX ix_roles_id ON roles (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.712 2022-05-11 08:16:10,712 INFO sqlalchemy.engine.Engine [no key 0.00014s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.728 2022-05-11 08:16:10,728 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:16:10.729 2022-05-11 08:16:10,729 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:10.729 2022-05-11 08:16:10,729 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackend     ⚡ 08:16:10.729 2022-05-11 08:16:10,729 INFO sqlalchemy.engine.Engine [cached since 0.1707s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 2022-05-11 08:16:10,731 INFO sqlalchemy.engine.Engine 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 CREATE TABLE users (
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	id VARCHAR(36) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	username VARCHAR(50) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	email VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	phone_number VARCHAR(20), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	json_permissions TEXT NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	active BOOL NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	hashed_password VARCHAR(60) NOT NULL, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	full_name VARCHAR(50), 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	created_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	updated_at DATETIME, 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 	PRIMARY KEY (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 )
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 
💀    🚀 startDemoBackend     ⚡ 08:16:10.731 2022-05-11 08:16:10,731 INFO sqlalchemy.engine.Engine [no key 0.00010s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.752 2022-05-11 08:16:10,752 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_phone_number ON users (phone_number)
💀    🚀 startDemoBackend     ⚡ 08:16:10.752 2022-05-11 08:16:10,752 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.768 2022-05-11 08:16:10,768 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_active ON users (active)
💀    🚀 startDemoBackend     ⚡ 08:16:10.768 2022-05-11 08:16:10,768 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.784 2022-05-11 08:16:10,784 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_id ON users (id)
💀    🚀 startDemoBackend     ⚡ 08:16:10.784 2022-05-11 08:16:10,784 INFO sqlalchemy.engine.Engine [no key 0.00013s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.8   2022-05-11 08:16:10,800 INFO sqlalchemy.engine.Engine CREATE INDEX ix_users_full_name ON users (full_name)
💀    🚀 startDemoBackend     ⚡ 08:16:10.8   2022-05-11 08:16:10,800 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.819 2022-05-11 08:16:10,819 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_username ON users (username)
💀    🚀 startDemoBackend     ⚡ 08:16:10.819 2022-05-11 08:16:10,819 INFO sqlalchemy.engine.Engine [no key 0.00015s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.836 2022-05-11 08:16:10,836 INFO sqlalchemy.engine.Engine CREATE UNIQUE INDEX ix_users_email ON users (email)
💀    🚀 startDemoBackend     ⚡ 08:16:10.836 2022-05-11 08:16:10,836 INFO sqlalchemy.engine.Engine [no key 0.00017s] {}
💀    🚀 startDemoBackend     ⚡ 08:16:10.852 2022-05-11 08:16:10,852 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:16:10.854 2022-05-11 08:16:10,854 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:10.856 2022-05-11 08:16:10,856 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackend     ⚡ 08:16:10.856 FROM users 
💀    🚀 startDemoBackend     ⚡ 08:16:10.856 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackend     ⚡ 08:16:10.856  LIMIT %(param_1)s
💀    🚀 startDemoBackend     ⚡ 08:16:10.856 2022-05-11 08:16:10,856 INFO sqlalchemy.engine.Engine [generated in 0.00014s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackend     ⚡ 08:16:10.858 2022-05-11 08:16:10,858 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 08:16:11.057 2022-05-11 08:16:11,057 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:11.058 2022-05-11 08:16:11,058 INFO sqlalchemy.engine.Engine INSERT INTO users (id, username, email, phone_number, json_permissions, active, hashed_password, full_name, created_at, updated_at) VALUES (%(id)s, %(username)s, %(email)s, %(phone_number)s, %(json_permissions)s, %(active)s, %(hashed_password)s, %(full_name)s, %(created_at)s, %(updated_at)s)
💀    🚀 startDemoBackend     ⚡ 08:16:11.058 2022-05-11 08:16:11,058 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'id': 'e98ee9c8-5ed6-4628-ac2d-31f7b1e0c60e', 'username': 'root', 'email': 'root@innistrad.com', 'phone_number': '+621234567890', 'json_permissions': '["root"]', 'active': 1, 'hashed_password': '$2b$12$6IzsmeGbmmJisl9/BaeMl.I.9DasGZb9Eo9xc8ZeL7s4JMADdLhPq', 'full_name': 'root', 'created_at': datetime.datetime(2022, 5, 11, 8, 16, 11, 56672), 'updated_at': datetime.datetime(2022, 5, 11, 8, 16, 11, 58190)}
💀    🚀 startDemoBackend     ⚡ 08:16:11.06  2022-05-11 08:16:11,060 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackend     ⚡ 08:16:11.068 2022-05-11 08:16:11,068 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackend     ⚡ 08:16:11.069 2022-05-11 08:16:11,069 INFO sqlalchemy.engine.Engine SELECT users.id, users.username, users.email, users.phone_number, users.json_permissions, users.active, users.hashed_password, users.full_name, users.created_at, users.updated_at 
💀    🚀 startDemoBackend     ⚡ 08:16:11.069 FROM users 
💀    🚀 startDemoBackend     ⚡ 08:16:11.069 WHERE users.id = %(pk_1)s
💀    🚀 startDemoBackend     ⚡ 08:16:11.069 2022-05-11 08:16:11,069 INFO sqlalchemy.engine.Engine [generated in 0.00011s] {'pk_1': 'e98ee9c8-5ed6-4628-ac2d-31f7b1e0c60e'}
💀    🚀 startDemoBackend     ⚡ 08:16:11.07  2022-05-11 08:16:11,070 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackend     ⚡ 08:16:11.071 Register app shutdown handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.079 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Handle HTTP routes for auth.User
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Register auth route handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Register auth event handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Handle RPC for auth.Role
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Handle RPC for auth.User
💀    🚀 startDemoBackend     ⚡ 08:16:11.089 Register auth RPC handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.095 Handle HTTP routes for library.Book
💀    🚀 startDemoBackend     ⚡ 08:16:11.095 Register library route handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.095 Register library event handler
💀    🚀 startDemoBackend     ⚡ 08:16:11.095 Handle RPC for library.Book
💀    🚀 startDemoBackend     ⚡ 08:16:11.095 Register library RPC handler
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:11.095 INFO:     Started server process [1908]
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:11.095 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:11.096 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:11.096 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackend     ⚡ 08:16:11.179 🔎 Port '3000' is ready
💀    🔎 startDemoBackend     ⚡ 08:16:11.179 check demoBackend
💀    🔎 startDemoBackend     ⚡ 08:16:11.18  🎉🎉🎉
💀    🔎 startDemoBackend     ⚡ 08:16:11.18  📜 Task 'startDemoBackend' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackend' readiness check
💀 🏁 Run 🏁 'start' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 start                🏁 08:16:11.285 
💀 🎉 Successfully running 🏁 'start' command
💀 🔎 Job Running...
         Elapsed Time: 49.723541523s
         Current Time: 08:16:11
         Active Process:
           * (PID=1511) 📗 'startDemoFrontendContainer' service
           * (PID=1903) ⚡ 'startDemoBackend' service
           * (PID=1564) 🐬 'startDemoDbContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=1511)
💀 🔪 Kill ⚡ 'startDemoBackend' service (PID=1903)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=1564)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:12.904 INFO:     Shutting down
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:13.004 INFO:     Waiting for application shutdown.
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:13.004 INFO:     Application shutdown complete.
💀 🔥 🚀 startDemoBackend     ⚡ 08:16:13.004 INFO:     Finished server process [1908]
💀    🚀 startDemoBackend     ⚡ 08:16:13.076 🎉🎉🎉
💀    🚀 startDemoBackend     ⚡ 08:16:13.076 📜 Task 'startDemoBackend' is started
💀 🔎 Job Ended...
         Elapsed Time: 51.827593854s
         Current Time: 08:16:13
💀 🔥 ⚡ 'startDemoBackend' service exited without any error message
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.242µs
         Current Time: 08:16:13
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:16:13.641 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:16:13.641 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 08:16:13.664 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 08:16:13.752 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 08:16:13.9   Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 08:16:13.9   Build image demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 08:16:15.73  Sending build context to Docker daemon  16.38kB
💀    🚀 buildDemoDbImage     🏭 08:16:15.735 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoDbImage     🏭 08:16:15.834 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 08:16:15.834  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:16:15.835 Successfully built 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 08:16:15.84  Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 08:16:15.845 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 08:16:15.845 Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 08:16:15.847 Step 1/11 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 08:16:15.847  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 08:16:15.847 Step 2/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848 Step 3/11 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 08:16:15.848 Step 4/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85  Step 5/11 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 808ba8676c5f
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85  Step 6/11 : COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 0c9047d38d7d
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85  Step 7/11 : USER 0
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 99c8982165ff
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85  Step 8/11 : COPY bootstrap.sh /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 3bacbc306156
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85  Step 9/11 : RUN chmod 755 /opt/bitnami/scripts/nginx/bootstrap.sh
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.85   ---> 0e12772b83fe
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851 Step 10/11 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851  ---> 8072400998af
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851 Step 11/11 : CMD ["/opt/bitnami/scripts/nginx/bootstrap.sh"]
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851  ---> 00baf0e406aa
💀    🚀 buildDemoFrontend... 🏭 08:16:15.851 Successfully built 00baf0e406aa
💀    🚀 buildDemoBackendI... 🏭 08:16:15.857 Sending build context to Docker daemon  1.179MB
💀    🚀 buildDemoFrontend... 🏭 08:16:15.861 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 08:16:15.866 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 08:16:15.866 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868  ---> caf584a25606
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 08:16:15.868 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 08:16:15.869 Step 6/9 : COPY . .
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 buildDemoBackendI... 🏭 08:16:16.116  ---> 7e283971f4e8
💀    🚀 buildDemoBackendI... 🏭 08:16:16.116 Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 08:16:16.154  ---> Running in 010b70a01be4
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 08:16:16.2   🔎 Waiting docker container 'demoDb' running status
💀    🔎 startDemoDbContainer 🐬 08:16:16.261 🔎 Waiting docker container 'demoDb' healthcheck
💀    🚀 buildDemoBackendI... 🏭 08:16:16.263 Removing intermediate container 010b70a01be4
💀    🚀 buildDemoBackendI... 🏭 08:16:16.263  ---> d49560a56120
💀    🚀 buildDemoBackendI... 🏭 08:16:16.263 Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 startDemoDbContainer 🐬 08:16:16.264 🐳 Container 'demoDb' is already started
💀    🚀 startDemoDbContainer 🐬 08:16:16.264 🐳 Logging 'demoDb'
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 08:16:16.317 🔎 Waiting docker container 'demoFrontend' running status
💀    🔎 startDemoDbContainer 🐬 08:16:16.317 🔎 Docker container 'demoDb' is running
💀    🔎 startDemoDbContainer 🐬 08:16:16.317 🔎 Waiting for host port: '3306'
💀    🔎 startDemoDbContainer 🐬 08:16:16.319 🔎 Host port '3306' is ready
💀    🚀 buildDemoBackendI... 🏭 08:16:16.339  ---> Running in 1bac5224ed92
💀    🚀 startDemoFrontend... 📗 08:16:16.37  🐳 Container 'demoFrontend' is already started
💀    🚀 startDemoFrontend... 📗 08:16:16.37  🐳 Logging 'demoFrontend'
💀    🔎 startDemoFrontend... 📗 08:16:16.379 🔎 Waiting docker container 'demoFrontend' healthcheck
💀    🔎 startDemoFrontend... 📗 08:16:16.445 🔎 Docker container 'demoFrontend' is running
💀    🔎 startDemoFrontend... 📗 08:16:16.445 🔎 Waiting for host port: '8080'
💀    🔎 startDemoFrontend... 📗 08:16:16.446 🔎 Host port '8080' is ready
💀    🔎 startDemoFrontend... 📗 08:16:16.446 🔎 Waiting for host port: '443'
💀    🔎 startDemoFrontend... 📗 08:16:16.453 🔎 Host port '443' is ready
💀    🚀 buildDemoBackendI... 🏭 08:16:17.232 Removing intermediate container 1bac5224ed92
💀    🚀 buildDemoBackendI... 🏭 08:16:17.232  ---> c6fdf9e0e00c
💀    🚀 buildDemoBackendI... 🏭 08:16:17.232 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 08:16:17.259  ---> Running in da1ef5f2867e
💀    🚀 buildDemoBackendI... 🏭 08:16:17.328 Removing intermediate container da1ef5f2867e
💀    🚀 buildDemoBackendI... 🏭 08:16:17.328  ---> 441a4178e7b7
💀    🚀 buildDemoBackendI... 🏭 08:16:17.33  Successfully built 441a4178e7b7
💀    🚀 buildDemoBackendI... 🏭 08:16:17.337 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 08:16:17.341 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 08:16:17.341 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀    🔎 startDemoDbContainer 🐬 08:16:19.328 🔎 Run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 08:16:19.456 🔎 Run check in 'demoFrontend': 'echo check demoFrontend'
💀 🔥 🔎 startDemoDbContainer 🐬 08:16:20.371 mysql: [Warning] Using a password on the command line interface can be insecure.
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 Database
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 information_schema
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 mysql
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 performance_schema
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 sample
💀    🔎 startDemoDbContainer 🐬 08:16:20.376 sys
💀    🔎 startDemoFrontend... 📗 08:16:20.377 check demoFrontend
💀    🔎 startDemoDbContainer 🐬 08:16:20.383 🔎 Sucessfully run check in 'demoDb': 'mysql -u root -pAlch3mist -e SHOW SCHEMAS'
💀    🔎 startDemoFrontend... 📗 08:16:20.385 🔎 Sucessfully run check in 'demoFrontend': 'echo check demoFrontend'
💀    🔎 startDemoDbContainer 🐬 08:16:24.387 🎉🎉🎉
💀    🔎 startDemoDbContainer 🐬 08:16:24.387 📜 Task 'startDemoDbContainer' is ready
💀    🔎 startDemoFrontend... 📗 08:16:24.388 🎉🎉🎉
💀    🔎 startDemoFrontend... 📗 08:16:24.388 📜 Task 'startDemoFrontendContainer' is ready
💀 🎉 Successfully running 📗 'startDemoFrontendContainer' readiness check
💀 🎉 Successfully running 🐬 'startDemoDbContainer' readiness check
💀 🎉 Reach 🐬 'startDemoDb' wrapper
💀 🏁 Run ⚡ 'startDemoBackendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀 🏁 Check ⚡ 'startDemoBackendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🔎 startDemoBackendC... ⚡ 08:16:25.075 🔎 Waiting docker container 'demoBackend' running status
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:25.103 Error: No such container: demoBackend
💀 🔥 🔎 startDemoBackendC... ⚡ 08:16:25.105 Error: No such container: demoBackend
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:25.129 Error: No such container: demoBackend
💀    🚀 startDemoBackendC... ⚡ 08:16:25.131 🐳 Creating and starting container 'demoBackend'
💀    🚀 startDemoBackendC... ⚡ 08:16:25.184 dddddf1de0b42c297798334ee739c5297780b0027dce325886aa8f704fbc5bbc
💀    🚀 startDemoBackendC... ⚡ 08:16:26.586 🐳 Logging 'demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:16:26.598 🔎 Waiting docker container 'demoBackend' healthcheck
💀    🔎 startDemoBackendC... ⚡ 08:16:26.658 🔎 Docker container 'demoBackend' is running
💀    🔎 startDemoBackendC... ⚡ 08:16:26.658 🔎 Waiting for host port: '3000'
💀    🔎 startDemoBackendC... ⚡ 08:16:26.66  🔎 Host port '3000' is ready
💀    🚀 startDemoBackendC... ⚡ 08:16:27.671 2022-05-11 01:16:27,670 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'sql_mode'
💀    🚀 startDemoBackendC... ⚡ 08:16:27.671 2022-05-11 01:16:27,670 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.676 2022-05-11 01:16:27,675 INFO sqlalchemy.engine.Engine SHOW VARIABLES LIKE 'lower_case_table_names'
💀    🚀 startDemoBackendC... ⚡ 08:16:27.676 2022-05-11 01:16:27,676 INFO sqlalchemy.engine.Engine [generated in 0.00015s] {}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.682 2022-05-11 01:16:27,682 INFO sqlalchemy.engine.Engine SELECT DATABASE()
💀    🚀 startDemoBackendC... ⚡ 08:16:27.682 2022-05-11 01:16:27,682 INFO sqlalchemy.engine.Engine [raw sql] {}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.688 2022-05-11 01:16:27,687 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:16:27.688 2022-05-11 01:16:27,688 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:16:27.688 2022-05-11 01:16:27,688 INFO sqlalchemy.engine.Engine [generated in 0.00016s] {'table_schema': 'sample', 'table_name': 'books'}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.692 2022-05-11 01:16:27,691 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:16:27.696 2022-05-11 01:16:27,696 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:16:27.697 2022-05-11 01:16:27,696 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:16:27.697 2022-05-11 01:16:27,697 INFO sqlalchemy.engine.Engine [cached since 0.008874s ago] {'table_schema': 'sample', 'table_name': 'roles'}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.702 2022-05-11 01:16:27,701 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:16:27.706 2022-05-11 01:16:27,705 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:16:27.706 2022-05-11 01:16:27,705 INFO sqlalchemy.engine.Engine SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = %(table_schema)s AND table_name = %(table_name)s
💀    🚀 startDemoBackendC... ⚡ 08:16:27.706 2022-05-11 01:16:27,706 INFO sqlalchemy.engine.Engine [cached since 0.01797s ago] {'table_schema': 'sample', 'table_name': 'users'}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.712 2022-05-11 01:16:27,711 INFO sqlalchemy.engine.Engine COMMIT
💀    🚀 startDemoBackendC... ⚡ 08:16:27.719 2022-05-11 01:16:27,719 INFO sqlalchemy.engine.Engine BEGIN (implicit)
💀    🚀 startDemoBackendC... ⚡ 08:16:27.722 2022-05-11 01:16:27,721 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.username AS users_username, users.email AS users_email, users.phone_number AS users_phone_number, users.json_permissions AS users_json_permissions, users.active AS users_active, users.hashed_password AS users_hashed_password, users.full_name AS users_full_name, users.created_at AS users_created_at, users.updated_at AS users_updated_at 
💀    🚀 startDemoBackendC... ⚡ 08:16:27.722 FROM users 
💀    🚀 startDemoBackendC... ⚡ 08:16:27.722 WHERE users.username = %(username_1)s 
💀    🚀 startDemoBackendC... ⚡ 08:16:27.722  LIMIT %(param_1)s
💀    🚀 startDemoBackendC... ⚡ 08:16:27.722 2022-05-11 01:16:27,721 INFO sqlalchemy.engine.Engine [generated in 0.00018s] {'username_1': 'root', 'param_1': 1}
💀    🚀 startDemoBackendC... ⚡ 08:16:27.727 2022-05-11 01:16:27,726 INFO sqlalchemy.engine.Engine ROLLBACK
💀    🚀 startDemoBackendC... ⚡ 08:16:27.732 Register app shutdown handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.749 Handle HTTP routes for auth.Role
💀    🚀 startDemoBackendC... ⚡ 08:16:27.765 Handle HTTP routes for auth.User
💀    🚀 startDemoBackendC... ⚡ 08:16:27.765 Register auth route handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.766 Register auth event handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.766 Handle RPC for auth.Role
💀    🚀 startDemoBackendC... ⚡ 08:16:27.766 Handle RPC for auth.User
💀    🚀 startDemoBackendC... ⚡ 08:16:27.766 Register auth RPC handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.778 Handle HTTP routes for library.Book
💀    🚀 startDemoBackendC... ⚡ 08:16:27.778 Register library route handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.778 Register library event handler
💀    🚀 startDemoBackendC... ⚡ 08:16:27.778 Handle RPC for library.Book
💀    🚀 startDemoBackendC... ⚡ 08:16:27.778 Register library RPC handler
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:27.779 INFO:     Started server process [8]
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:27.779 INFO:     Waiting for application startup.
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:27.779 INFO:     Application startup complete.
💀 🔥 🚀 startDemoBackendC... ⚡ 08:16:27.779 INFO:     Uvicorn running on http://0.0.0.0:3000 (Press CTRL+C to quit)
💀    🔎 startDemoBackendC... ⚡ 08:16:29.668 🔎 Run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:16:29.761 check demoBackend
💀    🔎 startDemoBackendC... ⚡ 08:16:29.764 🔎 Sucessfully run check in 'demoBackend': 'echo check demoBackend'
💀    🔎 startDemoBackendC... ⚡ 08:16:30.765 🎉🎉🎉
💀    🔎 startDemoBackendC... ⚡ 08:16:30.765 📜 Task 'startDemoBackendContainer' is ready
💀 🎉 Successfully running ⚡ 'startDemoBackendContainer' readiness check
💀 🏁 Run 🐳 'startContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 startContainers      🐳 08:16:30.873 
💀 🎉 Successfully running 🐳 'startContainers' command
💀 🔎 Job Running...
         Elapsed Time: 17.338728566s
         Current Time: 08:16:30
         Active Process:
           * (PID=4482) 📗 'startDemoFrontendContainer' service
           * (PID=4435) 🐬 'startDemoDbContainer' service
           * (PID=4568) ⚡ 'startDemoBackendContainer' service
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=4435)
💀 🔪 Kill ⚡ 'startDemoBackendContainer' service (PID=4568)
💀 🔪 Kill 📗 'startDemoFrontendContainer' service (PID=4482)
💀 🔥 📗 'startDemoFrontendContainer' service exited: signal: interrupt
💀 🔥 🐬 'startDemoDbContainer' service exited: signal: interrupt
💀 🔥 ⚡ 'startDemoBackendContainer' service exited: signal: interrupt
💀 🔎 Job Ended...
         Elapsed Time: 19.440820997s
         Current Time: 08:16:33
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
💀 🔎 Job Starting...
         Elapsed Time: 1.574µs
         Current Time: 08:16:33
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 08:16:33.238 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:16:33.238 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoFrontendC... ✋ 08:16:33.636 Stop docker container demoFrontend
💀    🚀 stopDemoDbContainer  ✋ 08:16:33.636 Stop docker container demoDb
💀    🚀 stopDemoBackendCo... ✋ 08:16:33.683 Stop docker container demoBackend
💀    🚀 stopDemoDbContainer  ✋ 08:16:38.162 demoDb
💀    🚀 stopDemoDbContainer  ✋ 08:16:38.164 🎉🎉🎉
💀    🚀 stopDemoDbContainer  ✋ 08:16:38.164 Docker container demoDb stopped
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀    🚀 stopDemoBackendCo... ✋ 08:16:44.513 demoBackend
💀    🚀 stopDemoBackendCo... ✋ 08:16:44.515 🎉🎉🎉
💀    🚀 stopDemoBackendCo... ✋ 08:16:44.515 Docker container demoBackend stopped
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀    🚀 stopDemoFrontendC... ✋ 08:16:44.973 demoFrontend
💀    🚀 stopDemoFrontendC... ✋ 08:16:44.975 🎉🎉🎉
💀    🚀 stopDemoFrontendC... ✋ 08:16:44.975 Docker container demoFrontend stopped
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 08:16:45.083 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 11.950348244s
         Current Time: 08:16:45
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 12.061102538s
         Current Time: 08:16:45
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.324µs
         Current Time: 08:16:45
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:16:45.448 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:16:45.45          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:16:45.45      
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:16:45.45    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:16:45.45    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:16:45.45    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:16:45.45  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:45.889 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:45.89  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.049 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.049 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.05  Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.264 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.271 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.278 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.278 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.278 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.278 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.278 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.281 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.281 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.291 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.291 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.295 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.295 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.299 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.332 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.332 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.332 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.623 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.623 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.79  Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:46.791 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.047 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.053 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.06  Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.06  ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.06  Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.06  Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.06  Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.064 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.064 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.074 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.074 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.077 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.077 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08  ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.08  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.102 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.105 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.109 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.273 Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.432 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.436 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.607 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.773 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.776 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.932 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.932 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:47.932 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.590329976s
         Current Time: 08:16:48
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.701301785s
         Current Time: 08:16:48
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.542µs
         Current Time: 08:16:48
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbShowAdv           ☕ 08:16:48.316 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:16:48.317 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:16:48.317 
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:16:48.317         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:16:48.317     
💀    🚀 zrbShowAdv           ☕ 08:16:48.317 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:16:48.317 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:16:48.317   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:16:48.317   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:16:48.317   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:16:48.317 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:48.768 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:48.768 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.541 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.542 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.542 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.751 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.758 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.764 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.764 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.764 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.764 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.764 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.767 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.767 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.776 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.776 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.778 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.779 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.781 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.782 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.782 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.782   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.782 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.782 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.812 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.812 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:49.812 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:50.263 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:50.263 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.018 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.24  Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.246 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.253 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.253 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.253 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.253 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.253 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.257 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.257 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.267 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.267 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.27  Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.27  Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.273 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.273 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.273 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.273   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.273 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.274 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.312 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.319 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.329 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.611 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.828 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:51.831 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.062 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.247 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.252 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.482 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.482 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:52.482 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.277520893s
         Current Time: 08:16:52
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.388871025s
         Current Time: 08:16:52
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.33µs
         Current Time: 08:16:52
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:16:52.865 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 08:16:52.867         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 08:16:52.867     
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 08:16:52.867   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 08:16:52.867   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 08:16:52.867   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 08:16:52.867 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.308 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.308 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.426 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.657 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.665 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.672 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.672 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.672 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.672 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.672 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.675 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.675 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.685 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.685 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.688 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.688 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.691 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.724 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.724 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:53.724 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.109 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.109 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.218 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.218 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.218 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.218 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.218 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.219 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.535 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.543 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.552 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.552 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.553 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.553 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.553 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.556 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.556 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.568 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.568 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.572 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.572 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577 ]
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.577 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.599 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.605 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.609 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:54.823 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.031 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.034 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.234 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.426 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.43  Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.62  Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.62  🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 08:16:55.62  Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 2.862505234s
         Current Time: 08:16:55
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.973752726s
         Current Time: 08:16:55
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.362µs
         Current Time: 08:16:56
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:16:56.028 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 08:16:56.134 Synchronize task environments
💀    🚀 syncEnv              🔄 08:16:56.346 Synchronize project's environment files
💀    🚀 syncEnv              🔄 08:16:56.541 🎉🎉🎉
💀    🚀 syncEnv              🔄 08:16:56.542 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 621.647478ms
         Current Time: 08:16:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 823.116439ms
         Current Time: 08:16:56
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.59µs
         Current Time: 08:16:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:16:57.043 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 08:16:57.157 🎉🎉🎉
💀    🚀 setProjectValue      🔗 08:16:57.157 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 220.632357ms
         Current Time: 08:16:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 421.569008ms
         Current Time: 08:16:57
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.557µs
         Current Time: 08:16:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 08:16:57.661 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 08:16:57.774 🎉🎉🎉
💀    🚀 setProjectValue      🔗 08:16:57.774 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 220.723418ms
         Current Time: 08:16:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 422.213089ms
         Current Time: 08:16:58
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.694µs
         Current Time: 08:16:58
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 08:16:58.288 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 08:16:58.289 🚧 Create virtual environment.
💀    🚀 prepareDemoBacken... 🏁 08:16:58.292 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 08:17:00.883 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 08:17:00.95  🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:17:00.974 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:01.293 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:01.299 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:01.313 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.117   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.135 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:02.215   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:02.235 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:02.307   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:02.326 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.344   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoFronte... 🏁 08:17:02.459   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 08:17:02.586   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.652 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:02.776 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.777   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:02.784 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:02.884 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:02.962   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:02.98  Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:03.128   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:03.159   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:03.163 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:03.179 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:03.256   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:03.262 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:03.548   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:03.629 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:03.685   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:03.713 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:03.813   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:03.833 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 08:17:03.851   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:03.856 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:03.914   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:03.922 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:03.928   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.008 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.117   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.13  Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.257   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.281 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.297   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.328 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.403   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.414 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.436   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.443 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.48    Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.488 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.526   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.54  Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.567   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.585 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:04.624   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.669   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.692 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:04.713 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.717   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.735 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.769   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.785 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:04.837   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:04.838   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:04.853 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:04.858 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.876   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:04.911 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:04.974   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.002 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:05.028   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.066   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:05.072 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.093 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:05.118   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.141 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:05.17    Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:05.185 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.185   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.196 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:05.215   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.251 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.311   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.323 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:05.325   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:05.333 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 08:17:05.38    Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.389 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.428   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:05.455   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.473   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.483 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:05.489 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 08:17:05.515 Installing collected packages: six, protobuf, grpcio, dill, pyyaml, semver, pulumi, arpeggio, attrs, parver, urllib3, certifi, charset-normalizer, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 08:17:05.589   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 08:17:05.597   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.6   Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 08:17:05.682 Installing collected packages: six, grpcio, pyyaml, semver, protobuf, dill, pulumi, arpeggio, attrs, parver, idna, certifi, urllib3, charset-normalizer, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 08:17:05.737   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 08:17:05.857 Installing collected packages: dill, pyyaml, semver, protobuf, six, grpcio, pulumi, arpeggio, attrs, parver, certifi, idna, charset-normalizer, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoDbDepl... 🏁 08:17:06.239   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 08:17:06.522   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 08:17:06.81    Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.125     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.184 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:17:08.226 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:17:08.226 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.401     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 08:17:08.466 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoFronte... 🏁 08:17:08.504 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:17:08.504 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.547 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.547 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.628 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.629       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63    -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63        --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63  
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.63  Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:08.632 🚧 Preparation completed.
💀    🚀 prepareDemoBacken... 🏁 08:17:08.657     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 08:17:08.718 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 prepareDemoBacken... 🏁 08:17:08.767 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:17:08.767 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.783 🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 08:17:08.783 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 08:17:08.865       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Usage:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:17:08.866       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867 
💀    🚀 prepareDemoFronte... 🏁 08:17:08.867 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:17:08.868 🚧 Preparation completed.
💀 🔥 🚀 deployDemoDbDeplo... 🏁 08:17:08.926 error: no stack named 'dev' found
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 deployDemoDbDeplo... 🏁 08:17:09.081 Created stack 'dev'
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 08:17:09.284 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 08:17:09.454 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 08:17:10.246 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:17:10.323 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 08:17:10.323 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.379 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  
💀    🚀 prepareDemoBacken... 🏁 08:17:10.38  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 for this case.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 Usage:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 Flags:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:17:10.381       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382 
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 08:17:10.382 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 08:17:10.753 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 08:17:10.867 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 08:17:11.168 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 08:17:11.489 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:17:11.531 
💀    🚀 deployDemoFronten... 🏁 08:17:11.866 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:11.899  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:11.977  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.248  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.251  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 08:17:12.253  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:17:12.317  +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344  
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.344 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 08:17:12.537  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 08:17:12.538  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 08:17:12.617  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 08:17:12.617  
💀    🚀 deployDemoFronten... 🏁 08:17:12.617 Resources:
💀    🚀 deployDemoFronten... 🏁 08:17:12.617     + 4 to create
💀    🚀 deployDemoFronten... 🏁 08:17:12.617 
💀    🚀 deployDemoFronten... 🏁 08:17:12.617 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 08:17:12.679 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 08:17:12.729 
💀    🚀 deployDemoFronten... 🏁 08:17:13.067 
💀    🚀 deployDemoBackend... 🏁 08:17:13.124 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.163  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.235  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoFronten... 🏁 08:17:13.443  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.478  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.48   +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.493  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.497  +  kubernetes:core/v1:ServiceAccount default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.501  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.508  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoFronten... 🏁 08:17:13.53   +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:13.546  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:17:13.637  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.654  +  pulumi:pulumi:Stack demoDbDeployment-dev created 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.654  
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655 Outputs:
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655     app: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655         ready    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655             [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655             [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655         resources: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.655             apps/v1/Deployment:default/demo-db: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                 api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                 id         : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                 kind       : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                 metadata   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                             apiVersion: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                             kind      : "Deployment"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                             metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                             spec      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 replicas: 1
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                 selector: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                     matchLabels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                         app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.656                                         app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                 template: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                     metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                         labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                     spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                         containers        : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                             [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                 env            : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                         value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                 ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                 image          : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                                 name           : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                         ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                         serviceAccountName: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     generation        : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                         helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.657                     managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             api_version: "apps/v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     f:strategy               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         f:rollingUpdate: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     f:template               : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             f:labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         f:spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             f:containers                   : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                 k:{"name":"demo-db"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                     f:env                     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"MYSQL_USER"}              : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             time       : "2022-05-11T01:17:13Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     resource_version  : "15808"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     uid               : "bb1a2a5e-9561-48c9-b0cd-46fcde243187"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                 spec       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     progress_deadline_seconds: 600
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     replicas                 : 1
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     revision_history_limit   : 10
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     selector                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         match_labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     strategy                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         rolling_update: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             max_surge      : "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             max_unavailable: "25%"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         type          : "RollingUpdate"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                     template                 : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             labels: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 app.kubernetes.io/instance: "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 app.kubernetes.io/name    : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                         spec    : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                             containers                      : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                 [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                     env                       : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             name : "MYSQL_DATABASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             value: "sample"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         [1]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             name : "MYSQL_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         [2]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             name : "MYSQL_ROOT_PASSWORD"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             value: "Alch3mist"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         [3]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             name : "MYSQL_USER"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                             value: "mysql"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.658                                         [4]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                         [5]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                         [6]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                             value: "1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     image                     : "demo-db:latest"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     name                      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                     termination_message_policy: "File"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             restart_policy                  : "Always"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             service_account                 : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             service_account_name            : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             termination_grace_period_seconds: 30
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659             v1/ServiceAccount:default/demo-db : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 api_version                    : "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 id                             : "default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 kind                           : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                 metadata                       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                     annotations       : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             apiVersion: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.659                             kind      : "ServiceAccount"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                              metadata  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      pulumi.com/skipAwait: "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  name       : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  namespace  : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                              }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66  
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                      }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                      creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                      labels            : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          app.kubernetes.io/instance  : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          app.kubernetes.io/name      : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                      }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                      managed_fields    : [
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                          [0]: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                              api_version: "v1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                              fields_type: "FieldsV1"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                              fields_v1  : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  f:metadata: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      f:annotations: {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      f:labels     : {
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                      }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.66                                  }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                             operation  : "Update"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                             time       : "2022-05-11T01:17:13Z"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                     ]
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                     name              : "demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                     namespace         : "default"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                     resource_version  : "15807"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                     uid               : "3b316adf-d1ca-4640-ad71-688162e642f7"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                 }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661                 urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661             }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661         }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661         urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661     }
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661     + 4 created
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 Duration: 1s
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 
💀    🚀 deployDemoDbDeplo... 🏁 08:17:13.661 hello world
💀 🎉 Successfully running 🏁 'deployDemoDbDeployment' command
💀    🚀 deployDemoFronten... 🏁 08:17:13.79   +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:17:13.792  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:17:13.803  +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:17:13.803  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoFronten... 🏁 08:17:13.807  +  kubernetes:core/v1:ServiceAccount default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 08:17:13.809  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoBackend... 🏁 08:17:13.895  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:17:13.896  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 08:17:13.901  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoFronten... 🏁 08:17:13.921  +  pulumi:pulumi:Stack demoFrontendDeployment-dev created 
💀    🚀 deployDemoFronten... 🏁 08:17:13.921  
💀    🚀 deployDemoFronten... 🏁 08:17:13.922 Outputs:
💀    🚀 deployDemoFronten... 🏁 08:17:13.923     app: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923         ready    : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.923             [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923             [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923         ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.923         resources: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923             apps/v1/Deployment:default/demo-frontend: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                 api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                 id         : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                 kind       : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                 metadata   : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                             apiVersion: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                             kind      : "Deployment"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.923                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                             spec      : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 replicas: 1
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 selector: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     matchLabels: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 template: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     metadata: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         labels: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     spec    : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         containers        : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                             [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                 env            : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     [1]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     [2]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     [3]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                         value: "1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                 ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                 image          : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                                 name           : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                         serviceAccountName: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.924 
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 08:17:13.924                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     generation        : 1
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     labels            : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             api_version: "apps/v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                 f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     f:strategy               : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         f:rollingUpdate: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     f:template               : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                             f:labels: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         f:spec    : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                             f:containers                   : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                 k:{"name":"demo-frontend"}: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                     f:env                     : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         k:{"name":"API_HOST"}                : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             time       : "2022-05-11T01:17:13Z"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     resource_version  : "15825"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     uid               : "eb6e0a47-c017-48c0-afc7-95a46eb730ba"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                 spec       : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     progress_deadline_seconds: 600
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     replicas                 : 1
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     revision_history_limit   : 10
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     selector                 : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         match_labels: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     strategy                 : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         rolling_update: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             max_surge      : "25%"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                             max_unavailable: "25%"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                         type          : "RollingUpdate"
💀    🚀 deployDemoFronten... 🏁 08:17:13.925                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     template                 : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         metadata: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             labels: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 app.kubernetes.io/instance: "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 app.kubernetes.io/name    : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         spec    : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             containers                      : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     env                       : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             name : "API_HOST"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             value: "localhost:3000"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         [1]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         [2]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         [3]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                             value: "1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     image                     : "demo-frontend:latest"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     image_pull_policy         : "IfNotPresent"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     name                      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     termination_message_path  : "/dev/termination-log"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     termination_message_policy: "File"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             dns_policy                      : "ClusterFirst"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             restart_policy                  : "Always"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             scheduler_name                  : "default-scheduler"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             service_account                 : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             service_account_name            : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             termination_grace_period_seconds: 30
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926             v1/ServiceAccount:default/demo-frontend : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 api_version                    : "v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 id                             : "default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 kind                           : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                 metadata                       : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     annotations       : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             apiVersion: "v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             kind      : "ServiceAccount"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             metadata  : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 annotations: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 labels     : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                     helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 name       : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                                 namespace  : "default"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926 
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                     labels            : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoFronten... 🏁 08:17:13.926                         app.kubernetes.io/name      : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.927                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.927                         helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                     managed_fields    : [
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                         [0]: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                             api_version: "v1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                             fields_type: "FieldsV1"
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                             fields_v1  : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                 f:metadata: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                     f:annotations: {
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                     f:labels     : {
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.928                                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                             operation  : "Update"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                             time       : "2022-05-11T01:17:13Z"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                     ]
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                     name              : "demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                     namespace         : "default"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                     resource_version  : "15824"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                     uid               : "debf9039-62e4-4ecc-895d-327dc3d215ef"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                 }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929                 urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929             }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929         }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929         urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 deployDemoFronten... 🏁 08:17:13.929     }
💀    🚀 deployDemoFronten... 🏁 08:17:13.929 
💀    🚀 deployDemoFronten... 🏁 08:17:13.929 Resources:
💀    🚀 deployDemoFronten... 🏁 08:17:13.929     + 4 created
💀    🚀 deployDemoFronten... 🏁 08:17:13.929 
💀    🚀 deployDemoFronten... 🏁 08:17:13.929 Duration: 1s
💀    🚀 deployDemoFronten... 🏁 08:17:13.929 
💀    🚀 deployDemoBackend... 🏁 08:17:14.028  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 08:17:14.028  
💀 🎉 Successfully running 🏁 'deployDemoFrontendDeployment' command
💀    🚀 deployDemoBackend... 🏁 08:17:14.028 Resources:
💀    🚀 deployDemoBackend... 🏁 08:17:14.028     + 5 to create
💀    🚀 deployDemoBackend... 🏁 08:17:14.028 
💀    🚀 deployDemoBackend... 🏁 08:17:14.028 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 08:17:14.476 
💀    🚀 deployDemoBackend... 🏁 08:17:14.912  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 08:17:14.99   +  kubernetes:helm.sh/v3:Chart demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.298  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.299  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.309  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.317  +  kubernetes:core/v1:ServiceAccount default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.327  +  kubernetes:core/v1:ServiceAccount default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:17:15.375  +  kubernetes:core/v1:Service default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.376  +  kubernetes:apps/v1:Deployment default/demo-backend creating 
💀    🚀 deployDemoBackend... 🏁 08:17:15.385  +  kubernetes:core/v1:Service default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:17:15.397  +  kubernetes:apps/v1:Deployment default/demo-backend created 
💀    🚀 deployDemoBackend... 🏁 08:17:15.66   +  pulumi:pulumi:Stack demoBackendDeployment-dev created 
💀    🚀 deployDemoBackend... 🏁 08:17:15.66   
💀    🚀 deployDemoBackend... 🏁 08:17:15.662 Outputs:
💀    🚀 deployDemoBackend... 🏁 08:17:15.663     app: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663         ready    : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.663             [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663             [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663             [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663         ]
💀    🚀 deployDemoBackend... 🏁 08:17:15.663         resources: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663             apps/v1/Deployment:default/demo-backend: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                 api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                 id         : "default/demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                 kind       : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                 metadata   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                     annotations       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                         kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                             apiVersion: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                             kind      : "Deployment"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                             metadata  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 annotations: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     pulumi.com/skipAwait: "true"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 labels     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                     helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 name       : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 namespace  : "default"
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                             spec      : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.663                                 replicas: 1
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                 selector: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                     matchLabels: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                         app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                         app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                 template: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                     metadata: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                         labels: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                     spec    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                         containers        : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                             [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                 env            : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "HS256"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [1]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "30"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [2]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [3]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "/token/"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [4]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [5]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [6]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [7]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "10"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [8]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "guest"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [9]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "3000"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [10]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [11]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [12]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                     [13]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.664                                                         name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [14]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [15]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [16]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [17]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [18]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [19]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [20]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [21]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [22]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [23]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [24]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [25]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [26]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "local"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [27]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [28]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [29]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "/static"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [30]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                     [31]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.665                                                         value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [32]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [33]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [34]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [35]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [36]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [37]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [38]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [39]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [40]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [41]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [42]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "/"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [43]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         value: "sqlite:///test.db"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 ]
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 image          : "demo-backend:latest"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 imagePullPolicy: "IfNotPresent"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 name           : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 ports          : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         containerPort: 3000
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         name         : "port0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                         protocol     : "TCP"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                                 ]
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                         ]
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                         serviceAccountName: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666 
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         pulumi.com/skipAwait                            : "true"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     creation_timestamp: "2022-05-11T01:17:15Z"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     generation        : 1
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     labels            : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         app.kubernetes.io/instance  : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         app.kubernetes.io/managed-by: "pulumi"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         app.kubernetes.io/name      : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         app.kubernetes.io/version   : "1.16.0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                     managed_fields    : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                         [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                             api_version: "apps/v1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                             fields_type: "FieldsV1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                             fields_v1  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                 f:metadata: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     f:annotations: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     f:labels     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                 f:spec    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     f:strategy               : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                         f:rollingUpdate: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                     f:template               : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.666                                         f:metadata: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                             f:labels: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                         f:spec    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                             f:containers                   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                 k:{"name":"demo-backend"}: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                     f:env                     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                     f:ports                   : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                             manager    : "pulumi-resource-kubernetes"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                             operation  : "Update"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                             time       : "2022-05-11T01:17:15Z"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     ]
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     name              : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     namespace         : "default"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     resource_version  : "15852"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     uid               : "8770dc55-1468-4bd3-bda0-73bce86904f9"
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                 }
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                 spec       : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     progress_deadline_seconds: 600
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     replicas                 : 1
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     revision_history_limit   : 10
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                     selector                 : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                         match_labels: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.667                             app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                     strategy                 : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         rolling_update: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             max_surge      : "25%"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             max_unavailable: "25%"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         type          : "RollingUpdate"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                     }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                     template                 : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         metadata: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             labels: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                 app.kubernetes.io/instance: "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                 app.kubernetes.io/name    : "demo-backend"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                         spec    : {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                             containers                      : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                 [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                     env                       : [
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [0]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "HS256"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [1]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "30"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [2]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [3]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ACCESS_TOKEN_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "/token/"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [4]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [5]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [6]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [7]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_ERROR_THRESHOLD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "10"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [8]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_GUEST_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             value: "guest"
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                         [9]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.668                                             name : "APP_HTTP_PORT"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "3000"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [10]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [11]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [12]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [13]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [14]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "http://localhost:8081"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [15]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "local"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [16]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_RABBITMQ_HOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "localhost"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [17]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_RABBITMQ_PASS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [18]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_RABBITMQ_USER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [19]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_RABBITMQ_VHOST"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             value: "/"
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                         [20]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.669                                             name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "root@innistrad.com"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [21]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [22]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "Alch3mist"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [23]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "+621234567890"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [24]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_ROOT_PERMISSION"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [25]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_ROOT_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "root"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [26]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_RPC_TYPE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "local"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [27]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "sqlite:///database.db"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [28]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name: "APP_STATIC_DIRECTORY"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [29]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "APP_STATIC_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "/static"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [30]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name: "PULUMI_BACKEND_URL"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [31]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              value: "defaultLocalPulumiPassphrase"
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          }
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                          [32]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.67                                              name : "PYTHONUNBUFFERED"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             value: "1"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [33]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             name : "TEST_INTEGRATION"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             value: "0"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [34]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             value: "localhost:9092"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [35]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             value: "PLAIN"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [36]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [37]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                             name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         }
💀    🚀 deployDemoBackend... 🏁 08:17:15.671                                         [38]: {
💀    🚀 deployDemoBackend... 🏁 08:17:15.671    
💀 🎉 Successfully running 🏁 'deployDemoBackendDeployment' command
💀 🏁 Run 🏭 'deploy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🎉 Successfully running 🏭 'deploy' command
💀 🔎 Job Running...
         Elapsed Time: 17.608855595s
         Current Time: 08:17:15
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 17.720252907s
         Current Time: 08:17:15
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.892µs
         Current Time: 08:17:16
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoFronte... 🏁 08:17:16.35  🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.35  🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 08:17:16.354 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.845 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.851 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.851 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.86  Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.861 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.862 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.873 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.874 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.875 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.876 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.878 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.879 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.883 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.884 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.884 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.886 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.89  Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.892 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.892 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.894 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.9   Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.901 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.904 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.909 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.91  Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.915 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.921 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.924 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.926 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.932 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoFronte... 🏁 08:17:16.951 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:16.955 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.956 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 08:17:16.982 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 08:17:17.003 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 08:17:17.024 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.028 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.031 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 08:17:17.038 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 08:17:17.043 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoBacken... 🏁 08:17:17.059 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.072 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.077 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 08:17:17.084 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 08:17:17.088 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:17:17.105 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 08:17:17.105 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 08:17:17.11  Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 08:17:17.116 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀 🔥 🚀 prepareDemoFronte... 🏁 08:17:17.123 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 08:17:17.123 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoBacken... 🏁 08:17:17.125 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀 🔥 🚀 prepareDemoBacken... 🏁 08:17:17.155 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 08:17:17.155 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.58  🚧 Deployment config: {"env":[{"name":"API_HOST","value":"localhost:3000"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 08:17:17.58  🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.717 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.717 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.723 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.723 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 dependencies.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 08:17:17.724 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725     dependencies:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 for this case.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Usage:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Aliases:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Flags:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 
💀    🚀 prepareDemoFronte... 🏁 08:17:17.725 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 08:17:17.728 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.857     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.858   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.859 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 08:17:17.862 🚧 Preparation completed.
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 08:17:19.259 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 08:17:19.346 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 08:17:19.346 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.41  
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 dependencies.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 08:17:19.411 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412     dependencies:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 for this case.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Usage:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Aliases:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Flags:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 08:17:19.412 
💀    🚀 prepareDemoBacken... 🏁 08:17:19.413 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 08:17:19.413 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 08:17:19.821 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 08:17:19.917 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.918  -  kubernetes:core/v1:ServiceAccount default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.92   -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.923  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:19.924 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 08:17:19.925  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.927  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.927  
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929   - app: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929         ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929       - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.929               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                    - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                            - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                            - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                  }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                  }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                              }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                            - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                        - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                        - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                      }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                  }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                - template: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                        - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                            - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                            - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                          }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                      }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                    - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                        - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                        -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                      }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.93                                                        - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                                 ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                         ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.931                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.932                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                     ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - resource_version  : "15825"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - uid               : "eb6e0a47-c017-48c0-afc7-95a46eb730ba"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.933                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.934                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                             ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.935                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.936                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - api_version: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                     ]
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.937                   - resource_version  : "15824"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938                   - uid               : "debf9039-62e4-4ecc-895d-327dc3d215ef"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938               - urn                            : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938             }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938         }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938       - urn      : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart::demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938     }
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938 Resources:
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938     - 4 to delete
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938 
💀    🚀 destroyDemoFronte... 🏁 08:17:19.938 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 08:17:20.012 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.012 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.013  -  kubernetes:core/v1:ServiceAccount default/demo-db delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.014  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.014  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.015  -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.016  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021  
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021       - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021       -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021       -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021       - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021           - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021               - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021               - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021               - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021               - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                           - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.021                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                           - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                               - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                               - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                       - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                       - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.022                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.023                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.024                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.025                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - resource_version  : "15808"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - uid               : "bb1a2a5e-9561-48c9-b0cd-46fcde243187"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.026                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                             ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.027                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.029                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.029                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.029                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.029                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                                  }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                              }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                            - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                            - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                          }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                      ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                    - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                    - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                    - resource_version  : "15807"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                    - uid               : "3b316adf-d1ca-4640-ad71-688162e642f7"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                  }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03                - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03              }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03          }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03        - urn      : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart::demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03      }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03  
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03  Resources:
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03      - 4 to delete
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03  
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.03  Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.138 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.148  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.148  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.148  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.148  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.148  -  kubernetes:core/v1:ServiceAccount default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.148  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.149  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.151  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.151  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.151  
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153 Outputs:
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153   - app: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153       - ready    : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153       -     [0]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153       -     [1]: "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153         ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.153       - resources: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154           - apps/v1/Deployment:default/demo-frontend: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154               - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154               - id         : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154               - kind       : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154               - metadata   : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                           - kind      : "Deployment"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - name       : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - namespace  : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                           - spec      : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - replicas: 1
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                               - selector: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                   - matchLabels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.154                                       - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                       - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                               - template: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                   - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                       - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                   - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                       - containers        : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                       -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               - env            : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                       - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                                 ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               - image          : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                               - name           : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                         ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                       - serviceAccountName: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155 
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                   - generation        : 1
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                   - labels            : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.155                       - app.kubernetes.io/instance  : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - app.kubernetes.io/name      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - helm.sh/chart               : "demo-frontend-0.1.0"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - managed_fields    : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - api_version: "apps/v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - fields_v1  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                               - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                   - f:annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                   - f:labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                               - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                   - f:strategy               : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                       - f:rollingUpdate: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                   - f:template               : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                       - f:metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                           - f:labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                       - f:spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                           - f:containers                   : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                               - k:{"name":"demo-frontend"}: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                   - f:env                     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                       - k:{"name":"API_HOST"}                : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - operation  : "Update"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                     ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - name              : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - namespace         : "default"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - resource_version  : "15825"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - uid               : "eb6e0a47-c017-48c0-afc7-95a46eb730ba"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156               - spec       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - replicas                 : 1
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - revision_history_limit   : 10
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - selector                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - match_labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                   - strategy                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                       - rolling_update: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.156                           - max_surge      : "25%"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - max_unavailable: "25%"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                       - type          : "RollingUpdate"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                   - template                 : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                       - metadata: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - labels: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                               - app.kubernetes.io/instance: "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                               - app.kubernetes.io/name    : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                       - spec    : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - containers                      : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - env                       : [
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   -     [0]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - name : "API_HOST"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - value: "localhost:3000"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   -     [1]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   -     [2]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   -     [3]: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                           - value: "1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                     ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - image                     : "demo-frontend:latest"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - name                      : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                   - termination_message_policy: "File"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                             ]
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - restart_policy                  : "Always"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - service_account                 : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - service_account_name            : "demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                         }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                     }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157               - urn        : "urn:pulumi:dev::demoFrontendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157             }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157           - v1/ServiceAccount:default/demo-frontend : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157               - api_version                    : "v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.157               - id                             : "default/demo-frontend"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158               - metadata                       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                   - annotations       : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                           - apiVersion: "v1"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                           - metadata  : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                               - annotations: {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                                 }
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158                               - labels     : {
💀    🚀 destroyDemoFronte... 🏁 08:17:20.158         
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.262  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.262  -  kubernetes:core/v1:ServiceAccount default/demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.263  -  kubernetes:core/v1:ServiceAccount default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.264  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.27   -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.278  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.278  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.278  
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.279 Outputs:
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.279   - app: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28        - ready    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28        -     [0]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28        -     [1]: "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28          ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28        - resources: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28            - apps/v1/Deployment:default/demo-db: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                - id         : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                - kind       : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                - metadata   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                    - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                        - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                            - apiVersion: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                            - kind      : "Deployment"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                            - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                  }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                  }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                              }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                            - spec      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - replicas: 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                - selector: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                    - matchLabels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                        - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                        - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.28                                      }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                               - template: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                   - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                       - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                   - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                       - containers        : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                       -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               - env            : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                       - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                                 ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               - image          : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                               - name           : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                         ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                       - serviceAccountName: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                   - generation        : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                           - api_version: "apps/v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.281                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                               - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                   - f:strategy               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                       - f:rollingUpdate: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                   - f:template               : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                       - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                           - f:labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                       - f:spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                           - f:containers                   : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                               - k:{"name":"demo-db"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                   - f:env                     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"MYSQL_DATABASE"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"MYSQL_PASSWORD"}          : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"MYSQL_ROOT_PASSWORD"}     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"MYSQL_USER"}              : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"PULUMI_BACKEND_URL"}      : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                       - k:{"name":"PYTHONUNBUFFERED"}        : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - resource_version  : "15808"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - uid               : "bb1a2a5e-9561-48c9-b0cd-46fcde243187"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282               - spec       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - replicas                 : 1
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - revision_history_limit   : 10
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - selector                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                       - match_labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - strategy                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                       - rolling_update: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - max_surge      : "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - max_unavailable: "25%"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                       - type          : "RollingUpdate"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                   - template                 : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                       - metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - labels: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                               - app.kubernetes.io/instance: "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                               - app.kubernetes.io/name    : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                       - spec    : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.282                           - containers                      : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - env                       : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "MYSQL_DATABASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "sample"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [1]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "MYSQL_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [2]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "MYSQL_ROOT_PASSWORD"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "Alch3mist"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [3]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "MYSQL_USER"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "mysql"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [4]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [5]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   -     [6]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                           - value: "1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - image                     : "demo-db:latest"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - name                      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - termination_message_policy: "File"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                             ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - restart_policy                  : "Always"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - service_account                 : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - service_account_name            : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283               - urn        : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283           - v1/ServiceAccount:default/demo-db : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283               - api_version                    : "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283               - id                             : "default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283               - metadata                       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                   - annotations       : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - apiVersion: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                           - metadata  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                               - annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                               - labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.283                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                                   - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                                   - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                               - name       : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                               - namespace  : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284 
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                   - creation_timestamp: "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.284                   - labels            : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                       - app.kubernetes.io/instance  : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                       - app.kubernetes.io/name      : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                       - helm.sh/chart               : "demo-db-0.1.0"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                   - managed_fields    : [
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                   -     [0]: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                           - api_version: "v1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                           - fields_v1  : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.286                               - f:metadata: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                                   - f:annotations: {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                                   - f:labels     : {
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                                     }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                           - operation  : "Update"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                           - time       : "2022-05-11T01:17:13Z"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                     ]
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                   - name              : "demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                   - namespace         : "default"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                   - resource_version  : "15807"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                   - uid               : "3b316adf-d1ca-4640-ad71-688162e642f7"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287                 }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287               - urn                            : "urn:pulumi:dev::demoDbDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-db"
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287             }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287         }
💀    🚀 destroyDemoDbDepl... 🏁 08:17:20.287   
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 08:17:21.103 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 08:17:21.211 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.214  -  kubernetes:core/v1:Service default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.217  -  kubernetes:core/v1:ServiceAccount default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.222  -  kubernetes:apps/v1:Deployment default/demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.226  -  kubernetes:helm.sh/v3:Chart demo-backend delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.229  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.232  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.232  
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236         ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.236                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.237                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.238                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.239                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                        - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.24                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                               -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.241                                                       - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                               -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.242                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                               -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                                       - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                                       - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                               - image          : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                               - imagePullPolicy: "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.243                                               - name           : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                               - ports          : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                                       - containerPort: 3000
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                                       - name         : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                                       - protocol     : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                                 ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                         ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                       - serviceAccountName: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                   - creation_timestamp: "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                   - generation        : 1
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                           - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.246                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                               - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                   - f:strategy               : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                       - f:rollingUpdate: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                   - f:template               : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                       - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                           - f:labels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                       - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                           - f:containers                   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                               - k:{"name":"demo-backend"}: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                   - f:env                     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ACCESS_TOKEN_ALGORITHM"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES"}: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ACCESS_TOKEN_SECRET_KEY"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ACCESS_TOKEN_URL"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ENABLE_EVENT_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ENABLE_ROUTE_HANDLER"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ENABLE_RPC_HANDLER"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_ERROR_THRESHOLD"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_GUEST_USERNAME"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_HTTP_PORT"}                  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                       - k:{"name":"APP_KAFKA_BOOTSTRAP_SERVERS"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.247                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_KAFKA_SASL_MECHANISM"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_KAFKA_SASL_PLAIN_USERNAME"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_KAFKA_SCHEMA_REGISTRY"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_MESSAGE_BUS_TYPE"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_RABBITMQ_HOST"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_RABBITMQ_PASS"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_RABBITMQ_USER"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_RABBITMQ_VHOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_INITIAL_EMAIL"}         : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_INITIAL_FULL_NAME"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_INITIAL_PASSWORD"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_INITIAL_PHONE_NUMBER"}  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_PERMISSION"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                       - k:{"name":"APP_ROOT_USERNAME"}              : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.248                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"APP_RPC_TYPE"}                   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"APP_SQLALCHEMY_DATABASE_URL"}    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"APP_STATIC_DIRECTORY"}           : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"APP_STATIC_URL"}                 : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"PULUMI_BACKEND_URL"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"PULUMI_CONFIG_PASSPHRASE"}       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"PYTHONUNBUFFERED"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"TEST_INTEGRATION"}               : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                       - k:{"name":"TEST_KAFKA_SASL_MECHANISM"}      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.249                                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD"} : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME"} : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_KAFKA_SCHEMA_REGISTRY"}     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_RABBITMQ_HOST"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_RABBITMQ_PASS"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_RABBITMQ_USER"}             : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_RABBITMQ_VHOST"}            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"name":"TEST_SQLALCHEMY_DATABASE_URL"}   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                    - f:ports                   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                        - k:{"containerPort":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                                  }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                              }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.25                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - time       : "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - resource_version  : "15852"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - uid               : "8770dc55-1468-4bd3-bda0-73bce86904f9"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - progress_deadline_seconds: 600
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - replicas                 : 1
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - revision_history_limit   : 10
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - selector                 : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                       - match_labels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - strategy                 : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                       - rolling_update: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - max_surge      : "25%"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - max_unavailable: "25%"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                       - type          : "RollingUpdate"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                   - template                 : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                       - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                               - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                               - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                       - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           - containers                      : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                           -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                   - env                       : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                   -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                   -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                   -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                           - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.251                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name: "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name: "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                   -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.252                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.253                                   -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                   -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                   -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                   -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                   -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.254                                           - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                   -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                   -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - name: "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                   -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                   -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - name: "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                   -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.255                                           - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [33]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - name : "TEST_INTEGRATION"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - value: "0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [34]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - name : "TEST_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [35]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - name : "TEST_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [36]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                           - name: "TEST_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.256                                   -     [37]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name: "TEST_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [38]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name : "TEST_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [39]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name : "TEST_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [40]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name : "TEST_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [41]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name : "TEST_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [42]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - name : "TEST_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                           - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.257                                   -     [43]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                           - name : "TEST_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                           - value: "sqlite:///test.db"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - image                     : "demo-backend:latest"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - image_pull_policy         : "IfNotPresent"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - name                      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - ports                     : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                           - container_port: 3000
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                           - name          : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                           - protocol      : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - termination_message_path  : "/dev/termination-log"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                   - termination_message_policy: "File"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                             ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - dns_policy                      : "ClusterFirst"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - restart_policy                  : "Always"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - scheduler_name                  : "default-scheduler"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - service_account                 : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - service_account_name            : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                           - termination_grace_period_seconds: 30
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.258             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259           - v1/Service:default/demo-backend        : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259               - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259               - kind       : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                           - kind      : "Service"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               - ports   : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.259                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                        - name      : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                        - port      : 3000
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                        - protocol  : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                        - targetPort: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                  ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                    - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                    - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                  }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                - type    : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                              }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26  
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                    - creation_timestamp: "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                    - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                        - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                    - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                    -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                    - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                    - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                  }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                - f:spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                    - f:ports                : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                        - k:{"port":3000,"protocol":"TCP"}: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                                  }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                              }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                            - time       : "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                          }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.26                      ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - resource_version  : "15853"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - uid               : "0f17918d-5aee-4b83-a30c-44ff9f1843c9"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261               - spec       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - cluster_ip             : "10.102.83.46"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - cluster_ips            : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   -     [0]: "10.102.83.46"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - internal_traffic_policy: "Cluster"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - ip_families            : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   -     [0]: "IPv4"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - ip_family_policy       : "SingleStack"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - ports                  : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                           - name       : "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                           - port       : 3000
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                           - protocol   : "TCP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                           - target_port: "port0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                   - selector               : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.261                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                   - session_affinity       : "None"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                   - type                   : "ClusterIP"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - status     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - urn        : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262           - v1/ServiceAccount:default/demo-backend : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - api_version                    : "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - id                             : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - kind                           : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262               - metadata                       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                           - apiVersion: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                           - kind      : "ServiceAccount"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.262                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                       - pulumi.com/skipAwait                            : "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                   - creation_timestamp: "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                   - labels            : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                       - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                       - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                       - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.263                       - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                       - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                   - managed_fields    : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                   -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - api_version: "v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - fields_type: "FieldsV1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - fields_v1  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                               - f:metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                                   - f:annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                                   - f:labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - manager    : "pulumi-resource-kubernetes"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - operation  : "Update"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                           - time       : "2022-05-11T01:17:15Z"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.264                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                     ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                   - name              : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                   - namespace         : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                   - resource_version  : "15848"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                   - uid               : "14163926-5937-4af7-850c-c391319c8d89"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265               - urn                            : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265       - urn      : "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart::demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265 Resources:
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265     - 5 to delete
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.265 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 08:17:21.364 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.377  -  kubernetes:apps/v1:Deployment default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.377  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.377  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.518  -  kubernetes:core/v1:ServiceAccount default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.52   -  kubernetes:core/v1:ServiceAccount default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.536  -  kubernetes:apps/v1:Deployment default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.545  -  kubernetes:core/v1:Service default/demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.546  -  kubernetes:core/v1:Service default/demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.548  -  kubernetes:helm.sh/v3:Chart demo-backend deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.569  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.575  -  kubernetes:helm.sh/v3:Chart demo-backend deleted 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.575  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 08:17:21.575  
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578 Outputs:
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578   - app: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578       - ready    : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578       -     [0]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:ServiceAccount::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578       -     [1]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:core/v1:Service::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578       -     [2]: "urn:pulumi:dev::demoBackendDeployment::kubernetes:helm.sh/v3:Chart$kubernetes:apps/v1:Deployment::default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578         ]
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578       - resources: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578           - apps/v1/Deployment:default/demo-backend: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578               - api_version: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578               - id         : "default/demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578               - kind       : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578               - metadata   : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578                   - annotations       : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578                       - kubectl.kubernetes.io/last-applied-configuration: (json) {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578                           - apiVersion: "apps/v1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.578                           - kind      : "Deployment"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                           - metadata  : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - annotations: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - pulumi.com/skipAwait: "true"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - labels     : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - app.kubernetes.io/instance  : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - app.kubernetes.io/managed-by: "pulumi"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - app.kubernetes.io/name      : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - app.kubernetes.io/version   : "1.16.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - helm.sh/chart               : "demo-backend-0.1.0"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - name       : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - namespace  : "default"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                             }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                           - spec      : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - replicas: 1
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - selector: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - matchLabels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                       - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                       - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                 }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                               - template: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - metadata: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                       - labels: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                           - app.kubernetes.io/instance: "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                           - app.kubernetes.io/name    : "demo-backend"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                         }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                   - spec    : {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                       - containers        : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                       -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                               - env            : [
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                               -     [0]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - name : "APP_ACCESS_TOKEN_ALGORITHM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - value: "HS256"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                               -     [1]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - name : "APP_ACCESS_TOKEN_EXPIRE_MINUTES"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - value: "30"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                               -     [2]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - name : "APP_ACCESS_TOKEN_SECRET_KEY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - value: "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                               -     [3]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.579                                                       - name : "APP_ACCESS_TOKEN_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "/token/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [4]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_ENABLE_EVENT_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [5]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_ENABLE_ROUTE_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [6]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_ENABLE_RPC_HANDLER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "1"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [7]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_ERROR_THRESHOLD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "10"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [8]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_GUEST_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "guest"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [9]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_HTTP_PORT"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "3000"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [10]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_KAFKA_BOOTSTRAP_SERVERS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "localhost:9092"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [11]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_KAFKA_SASL_MECHANISM"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "PLAIN"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [12]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_KAFKA_SASL_PLAIN_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [13]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_KAFKA_SASL_PLAIN_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [14]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_KAFKA_SCHEMA_REGISTRY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "http://localhost:8081"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                -     [15]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - name : "APP_MESSAGE_BUS_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                        - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.58                                                      }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [16]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_RABBITMQ_HOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "localhost"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [17]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_RABBITMQ_PASS"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [18]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_RABBITMQ_USER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [19]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_RABBITMQ_VHOST"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "/"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [20]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_INITIAL_EMAIL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "root@innistrad.com"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [21]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_INITIAL_FULL_NAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [22]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_INITIAL_PASSWORD"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "Alch3mist"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [23]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_INITIAL_PHONE_NUMBER"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "+621234567890"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [24]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_PERMISSION"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [25]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_ROOT_USERNAME"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "root"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [26]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - name : "APP_RPC_TYPE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                       - value: "local"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.581                                               -     [27]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "APP_SQLALCHEMY_DATABASE_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - value: "sqlite:///database.db"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                               -     [28]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "APP_STATIC_DIRECTORY"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                               -     [29]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "APP_STATIC_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - value: "/static"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                               -     [30]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "PULUMI_BACKEND_URL"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                               -     [31]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "PULUMI_CONFIG_PASSPHRASE"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - value: "defaultLocalPulumiPassphrase"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                     }
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                               -     [32]: {
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                                                       - name : "PYTHONUNBUFFERED"
💀    🚀 destroyDemoBacken... 🏁 08:17:21.582                         
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 08:17:21.695 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 5.473127455s
         Current Time: 08:17:21
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.583873639s
         Current Time: 08:17:21
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

