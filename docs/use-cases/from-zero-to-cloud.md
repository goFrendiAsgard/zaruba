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
         Elapsed Time: 1.866µs
         Current Time: 07:54:23
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 07:54:23.12  Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 07:54:23.128 🎉🎉🎉
💀    🚀 initProject          🚧 07:54:23.128 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 128.658725ms
         Current Time: 07:54:23
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 329.903128ms
         Current Time: 07:54:23
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 2.431µs
         Current Time: 07:54:23
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:54:23.666 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:54:23.672         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:54:23.672     
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:54:23.672   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:54:23.672   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:54:23.672   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:54:23.672 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 07:54:24.207 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 07:54:24.207 Preparing base variables
💀    🚀 makeMysqlApp         🐬 07:54:24.335 Base variables prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing start command
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Start command prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing test command
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Test command prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing check command
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Check command prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.336 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 07:54:24.671 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 07:54:24.682 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 07:54:24.692 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 07:54:24.692 ✅ Validate
💀    🚀 makeMysqlApp         🐬 07:54:24.692 Validate app directory
💀    🚀 makeMysqlApp         🐬 07:54:24.692 Done validating app directory
💀    🚀 makeMysqlApp         🐬 07:54:24.692 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 07:54:24.697 Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 07:54:24.697 Validate template locations
💀    🚀 makeMysqlApp         🐬 07:54:24.712 Done validating template locations
💀    🚀 makeMysqlApp         🐬 07:54:24.712 Validate app ports
💀    🚀 makeMysqlApp         🐬 07:54:24.716 Done validating app ports
💀    🚀 makeMysqlApp         🐬 07:54:24.716 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 07:54:24.721 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 07:54:24.721 🚧 Generate
💀    🚀 makeMysqlApp         🐬 07:54:24.721 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 07:54:24.721   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 07:54:24.721 ]
💀    🚀 makeMysqlApp         🐬 07:54:24.721 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlApp         🐬 07:54:24.745 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 07:54:24.745 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 07:54:24.746 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.26  🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.26  Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.509 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.51  Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.51  Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.51  Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.51  Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.51  Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.846 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.861 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.872 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.872 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.872 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.872 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.872 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.877 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.877 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.904 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.904 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.908 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.908 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914 🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914 🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914 ]
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.914 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.972 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.977 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:54:25.986 Checking prepareDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:54:26.228 Checking testDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:54:26.458 Checking migrateDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:54:26.665 Checking startDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:54:26.895 Checking start
💀    🚀 makeMysqlAppRunner   🐬 07:54:26.901 Adding startDemoDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 07:54:27.15  Checking startDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:54:27.359 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:27.364 Adding startDemoDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:27.593 Checking runDemoDb
💀    🚀 makeMysqlAppRunner   🐬 07:54:27.855 Checking runDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.107 Checking stopDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.369 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.373 Adding stopDemoDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.589 Checking removeDemoDbContainer
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.831 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:28.837 Adding removeDemoDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.106 Checking buildDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.357 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.362 Adding buildDemoDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.606 Checking pushDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.866 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:54:29.872 Adding pushDemoDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 07:54:30.118 Checking pullDemoDbImage
💀    🚀 makeMysqlAppRunner   🐬 07:54:30.357 Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 07:54:30.362 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 07:54:30.362 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 6.804260915s
         Current Time: 07:54:30
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.91454429s
         Current Time: 07:54:30
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 3.477µs
         Current Time: 07:54:30
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:54:30.811 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:54:30.815         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:54:30.815     
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:54:30.815   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:54:30.815   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:54:30.815   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:54:30.815 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 07:54:31.361 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 07:54:31.361 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 07:54:31.594 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.595 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 07:54:31.595 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.595 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 07:54:31.595 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.595 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 07:54:31.963 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:54:31.971 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 07:54:31.98  Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 07:54:31.98  ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 07:54:31.98  Validate app directory
💀    🚀 makeFastApiApp       ⚡ 07:54:31.98  Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 07:54:31.98  Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:54:31.985 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 07:54:31.985 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 07:54:32.002 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 07:54:32.002 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 07:54:32.007 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 07:54:32.007 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013 ]
💀    🚀 makeFastApiApp       ⚡ 07:54:32.013 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiApp       ⚡ 07:54:32.823 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 07:54:32.824 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 07:54:32.824 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 07:54:33.374 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:54:33.374 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.374 Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.661 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.67  Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.68  Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.68  ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.68  Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.68  Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.68  Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.684 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.685 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.704 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.704 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.708 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.708 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 ]
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.713 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.772 🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.776 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.78  Checking prepareDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:34.998 Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.001 Adding prepareDemoBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.212 Checking testDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.387 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.391 Adding testDemoBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.594 Checking migrateDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.794 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.797 Adding migrateDemoBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 07:54:35.974 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.176 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.18  Adding startDemoBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.403 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.596 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.6   Adding startDemoBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.786 Checking runDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:36.98  Checking runDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.186 Checking stopDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.4   Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.404 Adding stopDemoBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.608 Checking removeDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.828 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:37.832 Adding removeDemoBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.037 Checking buildDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.241 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.245 Adding buildDemoBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.447 Checking pushDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.643 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.646 Adding pushDemoBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 07:54:38.845 Checking pullDemoBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 07:54:39.054 Done registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 07:54:39.257 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:54:39.463 Checking startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:39.682 Adding startDemoDb as dependency of startDemoBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 07:54:39.879 Checking startDemoDb
💀    🚀 makeFastApiAppRunner ⚡ 07:54:40.082 Checking startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:40.293 Adding startDemoDb as dependency of startDemoBackend
💀    🚀 makeFastApiAppRunner ⚡ 07:54:40.503 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 07:54:40.503 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 07:54:40.981 🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 07:54:40.981 Preparing base variables
💀    🚀 addFastApiModule     ⚡ 07:54:42.316 Base variables prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.316 Preparing start command
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Start command prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Preparing test command
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Test command prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Preparing check command
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Check command prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.317 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 07:54:42.687 Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 07:54:42.698 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 07:54:42.708 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 07:54:42.708 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 07:54:42.708 Validate app directory
💀    🚀 addFastApiModule     ⚡ 07:54:42.708 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 07:54:42.708 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 07:54:42.712 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 07:54:42.712 Validate template locations
💀    🚀 addFastApiModule     ⚡ 07:54:42.727 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 07:54:42.727 Validate app ports
💀    🚀 addFastApiModule     ⚡ 07:54:42.732 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 07:54:42.732 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 07:54:42.736 Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 07:54:42.736 🚧 Generate
💀    🚀 addFastApiModule     ⚡ 07:54:42.736 🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 07:54:42.736   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 07:54:42.736 ]
💀    🚀 addFastApiModule     ⚡ 07:54:42.737 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiModule     ⚡ 07:54:42.763 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 07:54:42.763 Registering module
💀    🚀 addFastApiModule     ⚡ 07:54:42.802 Done registering module
💀    🚀 addFastApiModule     ⚡ 07:54:42.803 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 07:54:42.803 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 07:54:43.27  🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 07:54:43.27  Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 07:54:44.581 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.581 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.582 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:44.953 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:44.964 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:44.979 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:44.979 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:54:45     Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 07:54:45     Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:54:45.106 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 07:54:45.107 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:54:45.205 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 07:54:45.205 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 07:54:45.362 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 07:54:45.362 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:54:45.461 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 07:54:45.461 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.462 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:45.955 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:45.966 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 07:54:45.978 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 07:54:45.978 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 07:54:45.978 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 07:54:45.978 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 07:54:45.978 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:54:45.983 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 07:54:45.983 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 07:54:46.001 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 07:54:46.001 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 07:54:46.009 Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 07:54:46.009 Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:54:46.014 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 07:54:46.014 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 07:54:46.014 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 07:54:46.015   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 07:54:46.015 ]
💀    🚀 addFastApiCrud       ⚡ 07:54:46.015 
💀    🚀 addFastApiCrud       ⚡ 07:54:46.015 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"library","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["demoDb"],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"demoDb\"]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 addFastApiCrud       ⚡ 07:54:46.06  🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 07:54:46.061 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:54:46.121 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 07:54:46.121 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:54:46.202 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 07:54:46.203 Registering repo
💀    🚀 addFastApiCrud       ⚡ 07:54:46.303 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 07:54:46.304 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 07:54:46.304 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 15.603520033s
         Current Time: 07:54:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 15.804024864s
         Current Time: 07:54:46
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["demoDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.447µs
         Current Time: 07:54:46
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:54:46.85  Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:54:46.854         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:54:46.854     
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:54:46.854   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:54:46.854   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:54:46.854   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:54:46.854 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 07:54:47.393 🧰 Prepare
💀    🚀 makeNginxApp         📗 07:54:47.394 Preparing base variables
💀    🚀 makeNginxApp         📗 07:54:47.548 Base variables prepared
💀    🚀 makeNginxApp         📗 07:54:47.548 Preparing start command
💀    🚀 makeNginxApp         📗 07:54:47.549 Start command prepared
💀    🚀 makeNginxApp         📗 07:54:47.549 Preparing prepare command
💀    🚀 makeNginxApp         📗 07:54:47.549 Prepare command prepared
💀    🚀 makeNginxApp         📗 07:54:47.549 Preparing test command
💀    🚀 makeNginxApp         📗 07:54:47.549 Test command prepared
💀    🚀 makeNginxApp         📗 07:54:47.549 Preparing migrate command
💀    🚀 makeNginxApp         📗 07:54:47.549 Migrate command prepared
💀    🚀 makeNginxApp         📗 07:54:47.549 Preparing check command
💀    🚀 makeNginxApp         📗 07:54:47.549 Check command prepared
💀    🚀 makeNginxApp         📗 07:54:47.549 Preparing replacement map
💀    🚀 makeNginxApp         📗 07:54:47.897 Add config to replacement map
💀    🚀 makeNginxApp         📗 07:54:47.906 Add env to replacement map
💀    🚀 makeNginxApp         📗 07:54:47.916 Replacement map prepared
💀    🚀 makeNginxApp         📗 07:54:47.916 ✅ Validate
💀    🚀 makeNginxApp         📗 07:54:47.916 Validate app directory
💀    🚀 makeNginxApp         📗 07:54:47.916 Done validating app directory
💀    🚀 makeNginxApp         📗 07:54:47.916 Validate app container volumes
💀    🚀 makeNginxApp         📗 07:54:47.921 Done validating app container volumes
💀    🚀 makeNginxApp         📗 07:54:47.921 Validate template locations
💀    🚀 makeNginxApp         📗 07:54:47.936 Done validating template locations
💀    🚀 makeNginxApp         📗 07:54:47.936 Validate app ports
💀    🚀 makeNginxApp         📗 07:54:47.941 Done validating app ports
💀    🚀 makeNginxApp         📗 07:54:47.941 Validate app crud fields
💀    🚀 makeNginxApp         📗 07:54:47.946 Done validating app crud fields
💀    🚀 makeNginxApp         📗 07:54:47.946 🚧 Generate
💀    🚀 makeNginxApp         📗 07:54:47.946 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 07:54:47.946   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 07:54:47.946 ]
💀    🚀 makeNginxApp         📗 07:54:47.946 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxApp         📗 07:54:47.982 🔩 Integrate
💀    🚀 makeNginxApp         📗 07:54:47.982 🎉🎉🎉
💀    🚀 makeNginxApp         📗 07:54:47.982 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 07:54:48.558 🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 07:54:48.558 Preparing base variables
💀    🚀 makeNginxAppRunner   📗 07:54:48.733 Base variables prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.733 Preparing start command
💀    🚀 makeNginxAppRunner   📗 07:54:48.733 Start command prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Preparing test command
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Test command prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Preparing check command
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Check command prepared
💀    🚀 makeNginxAppRunner   📗 07:54:48.734 Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 07:54:49.098 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 07:54:49.11  Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 07:54:49.121 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 07:54:49.121 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 07:54:49.121 Validate app directory
💀    🚀 makeNginxAppRunner   📗 07:54:49.121 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 07:54:49.121 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 07:54:49.127 Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 07:54:49.127 Validate template locations
💀    🚀 makeNginxAppRunner   📗 07:54:49.146 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 07:54:49.146 Validate app ports
💀    🚀 makeNginxAppRunner   📗 07:54:49.151 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 07:54:49.151 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 07:54:49.156 Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 07:54:49.156 🚧 Generate
💀    🚀 makeNginxAppRunner   📗 07:54:49.156 🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 07:54:49.156   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 07:54:49.156   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 07:54:49.156 ]
💀    🚀 makeNginxAppRunner   📗 07:54:49.157 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: DEMO_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeNginxAppRunner   📗 07:54:49.199 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 07:54:49.205 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:54:49.21  Checking prepareDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:54:49.472 Checking testDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:54:49.732 Checking migrateDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:54:49.982 Checking startDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:54:50.272 Checking start
💀    🚀 makeNginxAppRunner   📗 07:54:50.277 Adding startDemoFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 07:54:50.545 Checking startDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:54:50.8   Checking startContainers
💀    🚀 makeNginxAppRunner   📗 07:54:50.805 Adding startDemoFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 07:54:51.072 Checking runDemoFrontend
💀    🚀 makeNginxAppRunner   📗 07:54:51.342 Checking runDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:54:51.584 Checking stopDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:54:51.843 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 07:54:51.848 Adding stopDemoFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 07:54:52.124 Checking removeDemoFrontendContainer
💀    🚀 makeNginxAppRunner   📗 07:54:52.392 Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 07:54:52.396 Adding removeDemoFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 07:54:52.644 Checking buildDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:54:52.873 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 07:54:52.879 Adding buildDemoFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 07:54:53.137 Checking pushDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:54:53.396 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 07:54:53.401 Adding pushDemoFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 07:54:53.62  Checking pullDemoFrontendImage
💀    🚀 makeNginxAppRunner   📗 07:54:53.871 Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 07:54:53.877 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 07:54:53.877 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 7.137773627s
         Current Time: 07:54:53
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 7.249650736s
         Current Time: 07:54:54
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v 'appPorts=["80:80"]' -v 'appEnvs={"API_HOST":"localhost:3000"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.97µs
         Current Time: 07:54:54
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:54:54.383 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:54:54.383 Links updated
💀 🏁 Run 🔧 'prepareDemoBackend' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 prepareDemoBackend   🔧 07:54:54.387 Create venv
💀    🚀 zrbCreateDockerNe... 🐳 07:54:54.436 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 07:54:54.514 Build image demo-frontend:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:54:54.656 Build image demo-db:latest
💀    🚀 prepareDemoBackend   🔧 07:54:57.544 Activate venv
💀    🚀 prepareDemoBackend   🔧 07:54:57.544 Install dependencies
💀    🚀 prepareDemoBackend   🔧 07:54:57.982 Collecting aiofiles==0.7.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBackend   🔧 07:54:58.337   Using cached https://files.pythonhosted.org/packages/e7/61/007ac6f27fe1c2dc44d3a62f429a8440de1601428b4d0291eae1a3494d1f/aiofiles-0.7.0-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:54:58.349 Collecting asgiref==3.4.1 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBackend   🔧 07:54:58.463   Using cached https://files.pythonhosted.org/packages/fe/66/577f32b54c50dcd8dec38447258e82ed327ecb86820d67ae7b3dea784f13/asgiref-3.4.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:54:58.477 Collecting avro-python3==1.10.0 (from -r requirements.txt (line 3))
💀    🚀 prepareDemoBackend   🔧 07:54:58.556   Using cached https://files.pythonhosted.org/packages/b2/5a/819537be46d65a01f8b8c6046ed05603fb9ef88c663b8cca840263788d58/avro-python3-1.10.0.tar.gz
💀    🚀 buildDemoDbImage     🏭 07:55:01.576 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:55:01.577 Sending build context to Docker daemon  13.31kB
💀    🚀 buildDemoDbImage     🏭 07:55:01.659 Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoDbImage     🏭 07:55:01.66   ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:55:01.66  Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 07:55:01.661 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoDbImage     🏭 07:55:01.665 Successfully tagged demo-db:latest
💀    🚀 buildDemoFrontend... 🏭 07:55:01.665  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:55:01.665 Step 2/6 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:55:01.666  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:01.666  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:55:01.666 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:55:01.666  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:01.667  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:55:01.667 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:55:01.667  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:01.667  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:55:01.667 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoFrontend... 🏭 07:55:01.668  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:01.668  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 07:55:01.668 Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:55:01.668  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:01.668  ---> 736550e2d78d
💀    🚀 buildDemoDbImage     🏭 07:55:01.67  🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:55:01.67  Docker image demo-db built
💀    🚀 buildDemoFrontend... 🏭 07:55:01.671 Successfully built 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 07:55:01.674 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoFrontend... 🏭 07:55:01.679 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:55:01.679 Docker image demo-frontend built
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀    🚀 prepareDemoBackend   🔧 07:55:01.826 Collecting bcrypt==3.2.0 (from -r requirements.txt (line 4))
💀    🚀 prepareDemoBackend   🔧 07:55:02.021   Using cached https://files.pythonhosted.org/packages/26/70/6d218afbe4c73538053c1016dd631e8f25fffc10cd01f5c272d7acf3c03d/bcrypt-3.2.0-cp36-abi3-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:55:02.043 Collecting certifi==2021.10.8 (from -r requirements.txt (line 5))
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:55:02.111 🔎 Waiting docker container 'demoFrontend' running status
💀    🚀 prepareDemoBackend   🔧 07:55:02.177   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:55:02.188 Collecting charset-normalizer==2.0.12 (from -r requirements.txt (line 6))
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 07:55:02.204 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 07:55:02.277 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 prepareDemoBackend   🔧 07:55:02.319   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 startDemoDbContainer 🐬 07:55:02.347 🐳 Retrieve previous log of 'demoDb'
💀    🚀 prepareDemoBackend   🔧 07:55:02.361 Collecting click==8.0.1 (from -r requirements.txt (line 7))
💀    🚀 prepareDemoBackend   🔧 07:55:02.473   Using cached https://files.pythonhosted.org/packages/76/0a/b6c5f311e32aeb3b406e03c079ade51e905ea630fc19d1262a46249c1c86/click-8.0.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:55:02.489 Collecting confluent-kafka[avro]==1.8.2 (from -r requirements.txt (line 8))
💀    🚀 prepareDemoBackend   🔧 07:55:02.676   Using cached https://files.pythonhosted.org/packages/da/9b/f09a614e6b6b5e892c7aa50240ffe4e132664abb7f7b9fcdc89a4dddf35b/confluent_kafka-1.8.2-cp38-cp38-manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:55:02.835 Collecting cryptography==36.0.1 (from -r requirements.txt (line 10))
💀    🚀 startDemoFrontend... 📗 07:55:03.34  
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.34  [38;5;6mnginx [38;5;5m23:53:38.12 
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.34  [38;5;6mnginx [38;5;5m23:53:38.12 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.34  [38;5;6mnginx [38;5;5m23:53:38.13 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.34  [38;5;6mnginx [38;5;5m23:53:38.13 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.34  [38;5;6mnginx [38;5;5m23:53:38.13 
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.13 [38;5;2mINFO  ==> ** Starting NGINX setup **
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.15 [38;5;2mINFO  ==> Validating settings in NGINX_* env vars
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> No custom scripts in /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> Initializing NGINX
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 realpath: /bitnami/nginx/conf/vhosts: No such file or directory
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.18 [38;5;2mINFO  ==> ** NGINX setup finished! **
💀 🔥 🚀 startDemoFrontend... 📗 07:55:03.341 [38;5;6mnginx [38;5;5m23:53:38.19 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🚀 startDemoFrontend... 📗 07:55:03.346 🐳 Starting container 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 07:55:03.434 
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.039634Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.041494Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.041504Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.045770Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.171704Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.350325Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.350378Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.416412Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:53:57.416514Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:54:36.519804Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:54:38.521030Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 07:55:03.435 2022-05-08T23:54:39.918045Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 [38;5;6mmysql [38;5;5m23:53:47.76 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 [38;5;6mmysql [38;5;5m23:53:53.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 [38;5;6mmysql [38;5;5m23:53:53.80 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 [38;5;6mmysql [38;5;5m23:53:56.81 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:03.435 [38;5;6mmysql [38;5;5m23:53:56.83 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 07:55:03.439 🐳 Starting container 'demoDb'
💀    🚀 prepareDemoBackend   🔧 07:55:03.498   Using cached https://files.pythonhosted.org/packages/0c/12/a55cf1ed39c2fa9a22448b82c984152fdeb7b30a66e3544eee3bd52b08fc/cryptography-36.0.1-cp36-abi3-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBackend   🔧 07:55:03.656 Collecting fastapi==0.68.1 (from -r requirements.txt (line 11))
💀    🚀 prepareDemoBackend   🔧 07:55:03.894   Using cached https://files.pythonhosted.org/packages/df/44/ee1976b03404318590bbe4b0ef27007ce2c42b15757aa0c72bc99a4ebae7/fastapi-0.68.1-py3-none-any.whl
💀    🚀 prepareDemoBackend   🔧 07:55:04.001 Collecting fastavro==1.4.9 (from -r requirements.txt (line 12))
💀 🔥 🚀 startDemoFrontend... 📗 07:55:04.209 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 07:55:04.209 Error: failed to start containers: demoFrontend
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
💀 🔪 Kill 🔧 'prepareDemoBackend' command (PID=11560)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=14826)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=14827)
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=14790)
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:04.347 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:04.347 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀    🚀 prepareDemoBackend   🔧 07:55:04.558   Using cached https://files.pythonhosted.org/packages/9f/d4/0a04211257324a27ef39e0309989f10d05227be63b601c7789a156b23623/fastavro-1.4.9.tar.gz
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 🚀 prepareDemoBackend   🔧 07:55:04.63  ERROR: Operation cancelled by user
💀 🔥 🚀 prepareDemoBackend   🔧 07:55:04.641 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBackend   🔧 07:55:04.641 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 Error running 🔧 'prepareDemoBackend' command: exit status 1
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 10.851456149s
         Current Time: 07:55:05
zaruba please start -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["start"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.657µs
         Current Time: 07:55:05
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🐳 'zrbCreateDockerNetwork' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:55:05.464 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:55:05.464 Links updated
💀    🚀 zrbCreateDockerNe... 🐳 07:55:05.497 🐳 Network 'zaruba' is already exist
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 🏭 'buildDemoDbImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🚀 buildDemoDbImage     🏭 07:55:05.58  Build image demo-db:latest
💀 🎉 Successfully running 🐳 'zrbCreateDockerNetwork' command
💀 🏁 Run 🏭 'buildDemoFrontendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🚀 buildDemoFrontend... 🏭 07:55:05.737 Build image demo-frontend:latest
💀 🏁 Run 🏭 'buildDemoBackendImage' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackend
💀    🚀 buildDemoBackendI... 🏭 07:55:05.741 Build image demo-backend:latest
💀    🚀 buildDemoDbImage     🏭 07:55:07.039 Sending build context to Docker daemon  3.072kB
💀    🚀 buildDemoFrontend... 🏭 07:55:07.041 Sending build context to Docker daemon  13.31kB
💀    🚀 buildDemoFrontend... 🏭 07:55:07.134 Step 1/6 : FROM docker.io/bitnami/nginx:1.21.6
💀    🚀 buildDemoFrontend... 🏭 07:55:07.134  ---> 0b9593fe1d77
💀    🚀 buildDemoFrontend... 🏭 07:55:07.135 Step 2/6 : USER 0
💀    🚀 buildDemoFrontend... 🏭 07:55:07.135  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:07.135  ---> 562078b73ebf
💀    🚀 buildDemoFrontend... 🏭 07:55:07.135 Step 3/6 : RUN apt update &&     apt install certbot -y &&     apt-get autoremove -yqq --purge &&     apt-get clean &&     rm -rf /var/lib/apt/lists/*
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136  ---> c0b95731b707
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136 Step 4/6 : USER 1001
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136  ---> 162e06eadcfd
💀    🚀 buildDemoFrontend... 🏭 07:55:07.136 Step 5/6 : COPY html /opt/bitnami/nginx/html
💀    🚀 buildDemoBackendI... 🏭 07:55:07.137 Sending build context to Docker daemon  1.029MB
💀    🚀 buildDemoFrontend... 🏭 07:55:07.137  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:07.137  ---> 0b7a8e3dd34d
💀    🚀 buildDemoFrontend... 🏭 07:55:07.137 Step 6/6 : COPY /server_blocks/my_server_block.cnf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf
💀    🚀 buildDemoFrontend... 🏭 07:55:07.137  ---> Using cache
💀    🚀 buildDemoFrontend... 🏭 07:55:07.137  ---> 736550e2d78d
💀    🚀 buildDemoFrontend... 🏭 07:55:07.138 Successfully built 736550e2d78d
💀    🚀 buildDemoDbImage     🏭 07:55:07.14  Step 1/1 : FROM docker.io/bitnami/mysql:8.0.29
💀    🚀 buildDemoFrontend... 🏭 07:55:07.142 Successfully tagged demo-frontend:latest
💀    🚀 buildDemoDbImage     🏭 07:55:07.143  ---> 188ba73f5790
💀    🚀 buildDemoDbImage     🏭 07:55:07.144 Successfully built 188ba73f5790
💀    🚀 buildDemoFrontend... 🏭 07:55:07.146 🎉🎉🎉
💀    🚀 buildDemoFrontend... 🏭 07:55:07.146 Docker image demo-frontend built
💀    🚀 buildDemoBackendI... 🏭 07:55:07.148 Step 1/9 : FROM python:3.8-slim
💀    🚀 buildDemoBackendI... 🏭 07:55:07.148  ---> caf584a25606
💀    🚀 buildDemoDbImage     🏭 07:55:07.153 Successfully tagged demo-db:latest
💀    🚀 buildDemoDbImage     🏭 07:55:07.153 🎉🎉🎉
💀    🚀 buildDemoDbImage     🏭 07:55:07.153 Docker image demo-db built
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153 Step 2/9 : ENV PYTHONUNBUFFERED 1
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> 7296d7455c56
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153 Step 3/9 : WORKDIR /app
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> c9a3cbe90f60
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153 Step 4/9 : COPY requirements.txt ./
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153  ---> 90b390a57c9c
💀    🚀 buildDemoBackendI... 🏭 07:55:07.153 Step 5/9 : RUN pip install -r requirements.txt
💀    🚀 buildDemoBackendI... 🏭 07:55:07.154  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.154  ---> 16e3e46a7774
💀    🚀 buildDemoBackendI... 🏭 07:55:07.154 Step 6/9 : COPY . .
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17   ---> 8eab2e0c1eec
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17  Step 7/9 : EXPOSE 3000
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17   ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17   ---> 3bdcbd278244
💀    🚀 buildDemoBackendI... 🏭 07:55:07.17  Step 8/9 : RUN chmod 755 ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:55:07.171  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.171  ---> 0109fee8acf7
💀    🚀 buildDemoBackendI... 🏭 07:55:07.171 Step 9/9 : CMD ./start.sh
💀    🚀 buildDemoBackendI... 🏭 07:55:07.171  ---> Using cache
💀    🚀 buildDemoBackendI... 🏭 07:55:07.171  ---> 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:55:07.174 Successfully built 8ce3e60f57bf
💀    🚀 buildDemoBackendI... 🏭 07:55:07.177 Successfully tagged demo-backend:latest
💀    🚀 buildDemoBackendI... 🏭 07:55:07.181 🎉🎉🎉
💀    🚀 buildDemoBackendI... 🏭 07:55:07.181 Docker image demo-backend built
💀 🎉 Successfully running 🏭 'buildDemoFrontendImage' command
💀 🎉 Successfully running 🏭 'buildDemoDbImage' command
💀 🎉 Successfully running 🏭 'buildDemoBackendImage' command
💀 🏁 Run 📗 'startDemoFrontendContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀 🏁 Check 📗 'startDemoFrontendContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontend
💀    🔎 startDemoFrontend... 📗 07:55:07.525 🔎 Waiting docker container 'demoFrontend' running status
💀 🏁 Run 🐬 'startDemoDbContainer' service on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀 🏁 Check 🐬 'startDemoDbContainer' readiness on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDb
💀    🔎 startDemoDbContainer 🐬 07:55:07.581 🔎 Waiting docker container 'demoDb' running status
💀    🚀 startDemoFrontend... 📗 07:55:07.673 🐳 Retrieve previous log of 'demoFrontend'
💀    🚀 startDemoDbContainer 🐬 07:55:07.709 🐳 Retrieve previous log of 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.12 
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.12 Welcome to the Bitnami nginx container
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.13 Subscribe to project updates by watching https://github.com/bitnami/bitnami-docker-nginx
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.13 Submit issues and feature requests at https://github.com/bitnami/bitnami-docker-nginx/issues
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.13 
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.13 [38;5;2mINFO  ==> ** Starting NGINX setup **
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.15 [38;5;2mINFO  ==> Validating settings in NGINX_* env vars
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> No custom scripts in /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.16 [38;5;2mINFO  ==> Initializing NGINX
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 realpath: /bitnami/nginx/conf/vhosts: No such file or directory
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.18 [38;5;2mINFO  ==> ** NGINX setup finished! **
💀 🔥 🚀 startDemoFrontend... 📗 07:55:08.714 [38;5;6mnginx [38;5;5m23:53:38.19 [38;5;2mINFO  ==> ** Starting NGINX **
💀    🚀 startDemoFrontend... 📗 07:55:08.714 
💀    🚀 startDemoFrontend... 📗 07:55:08.716 🐳 Starting container 'demoFrontend'
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 [38;5;6mmysql [38;5;5m23:53:47.76 [38;5;2mINFO  ==> Starting mysql in background
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 [38;5;6mmysql [38;5;5m23:53:53.77 [38;5;2mINFO  ==> Loading user's custom files from /docker-entrypoint-initdb.d
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 mysql: [Warning] Using a password on the command line interface can be insecure.
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 find: '/docker-entrypoint-startdb.d/': No such file or directory
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 [38;5;6mmysql [38;5;5m23:53:53.80 [38;5;2mINFO  ==> Stopping mysql
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 [38;5;6mmysql [38;5;5m23:53:56.81 [38;5;2mINFO  ==> ** MySQL setup finished! **
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:08.758 [38;5;6mmysql [38;5;5m23:53:56.83 [38;5;2mINFO  ==> ** Starting MySQL **
💀    🚀 startDemoDbContainer 🐬 07:55:08.758 
💀    🚀 startDemoDbContainer 🐬 07:55:08.758 2022-05-08T23:53:57.039634Z 0 [System] [MY-010116] [Server] /opt/bitnami/mysql/bin/mysqld (mysqld 8.0.29) starting as process 1
💀    🚀 startDemoDbContainer 🐬 07:55:08.758 2022-05-08T23:53:57.041494Z 0 [Warning] [MY-013242] [Server] --character-set-server: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.041504Z 0 [Warning] [MY-013244] [Server] --collation-server: 'utf8_general_ci' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.045770Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.171704Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.350325Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.350378Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.416412Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /tmp/mysqlx.sock
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:53:57.416514Z 0 [System] [MY-010931] [Server] /opt/bitnami/mysql/bin/mysqld: ready for connections. Version: '8.0.29'  socket: '/opt/bitnami/mysql/tmp/mysql.sock'  port: 3306  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:54:36.519804Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.29).
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:54:38.521030Z 0 [Warning] [MY-010909] [Server] /opt/bitnami/mysql/bin/mysqld: Forcing close of thread 12  user: 'root'.
💀    🚀 startDemoDbContainer 🐬 07:55:08.759 2022-05-08T23:54:39.918045Z 0 [System] [MY-010910] [Server] /opt/bitnami/mysql/bin/mysqld: Shutdown complete (mysqld 8.0.29)  Source distribution.
💀    🚀 startDemoDbContainer 🐬 07:55:08.762 🐳 Starting container 'demoDb'
💀 🔥 🚀 startDemoFrontend... 📗 07:55:09.565 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/113d97acbd7f34c2b509379488ec44e364666e0b41ab486771cae22b3d6d1a01" to rootfs at "/etc/letsencrypt" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoFrontend... 📗 07:55:09.565 Error: failed to start containers: demoFrontend
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
💀 🔪 Kill 📗 'startDemoFrontendContainer' readiness check (PID=17275)
💀 🔪 Kill 🐬 'startDemoDbContainer' service (PID=17304)
💀 🔪 Kill 🐬 'startDemoDbContainer' readiness check (PID=17306)
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:09.764 Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: process_linux.go:545: container init caused: rootfs_linux.go:75: mounting "/run/desktop/mnt/host/wsl/docker-desktop-bind-mounts/Ubuntu/d3cfdac06a492498f225b55a282f20c76d878ba74f0cd0a3093feb76d542d506" to rootfs at "/docker-entrypoint-initdb.d" caused: mount through procfd: no such file or directory: unknown
💀 🔥 🚀 startDemoDbContainer 🐬 07:55:09.764 Error: failed to start containers: demoDb
💀 🔥 🐬 'startDemoDbContainer' service exited: exit status 1
💀 🔥 Error running 📗 'startDemoFrontendContainer' readiness check: signal: interrupt
💀 🔥 Error running 🐬 'startDemoDbContainer' readiness check: signal: interrupt
      no such process
💀 🔎 Job Ended...
         Elapsed Time: 5.115801089s
         Current Time: 07:55:10
zaruba please startContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml' -t -w 1s
🔥 Command   : zaruba please
🔥 Arguments : ["startContainers"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.861µs
         Current Time: 07:55:10
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 updateProjectLinks   🔗 07:55:10.859 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:55:10.859 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run ✋ 'stopDemoFrontendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoDbContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ✋ 'stopDemoBackendContainer' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopDemoBackendCo... ✋ 07:55:11.359 Docker container demoBackend is not running
💀    🚀 stopDemoFrontendC... ✋ 07:55:11.369 Docker container demoFrontend is not running
💀    🚀 stopDemoDbContainer  ✋ 07:55:11.371 Docker container demoDb is not running
💀 🎉 Successfully running ✋ 'stopDemoBackendContainer' command
💀 🎉 Successfully running ✋ 'stopDemoFrontendContainer' command
💀 🎉 Successfully running ✋ 'stopDemoDbContainer' command
💀 🏁 Run ✋ 'stopContainers' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 stopContainers       ✋ 07:55:11.478 
💀 🎉 Successfully running ✋ 'stopContainers' command
💀 🔎 Job Running...
         Elapsed Time: 727.468211ms
         Current Time: 07:55:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 837.792018ms
         Current Time: 07:55:11
zaruba please stopContainers -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.27µs
         Current Time: 07:55:11
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:11.937 Current directory is a valid zaruba project
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbShowAdv           ☕ 07:55:11.944 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:55:11.944 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:55:11.944 
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:55:11.944         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:55:11.944     
💀    🚀 zrbShowAdv           ☕ 07:55:11.944 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:55:11.944 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:55:11.944   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:55:11.944   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:55:11.944   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:55:11.945 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.511 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.511 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.777 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.777 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.777 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.777 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.777 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:12.778 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.115 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.126 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.136 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.137 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.137 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.137 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.137 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.141 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.141 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.156 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.156 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.161 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.161 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.164 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.165 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.165 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.165   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.165 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.165 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.224 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.224 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.224 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.659 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.659 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.91  Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:13.911 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.265 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.276 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.289 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.289 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.289 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.289 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.289 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.293 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.294 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.312 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.312 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.317 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.318 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.322 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoDbDeployment","ZtplTaskName":"DemoDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: DEMO_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: DEMO_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: DEMO_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: DEMO_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-db","ztpl-app-event-name":"","ztpl-app-image-name":"demo-db","ztpl-app-module-name":"","ztpl-app-name":"demo-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-db-deployment","ztpl-task-name":"demo-db","ztplAppContainerName":"demoDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoDbDeployment","ztplDeploymentName":"demoDbDeployment","ztplDeploymentTaskLocation":"../../demoDbDeployment","ztplTaskName":"demoDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_db_deployment","ztpl_task_name":"demo_db"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.354 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.358 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.364 Checking prepareDemoDbDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.61  Checking deployDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.861 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:14.867 Adding deployDemoDbDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.118 Checking destroyDemoDbDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.386 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.391 Adding destroyDemoDbDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.658 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.658 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:55:15.658 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 3.831388705s
         Current Time: 07:55:15
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.941844293s
         Current Time: 07:55:15
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.24µs
         Current Time: 07:55:16
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:16.124 Current directory is a valid zaruba project
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbShowAdv           ☕ 07:55:16.13  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:55:16.13  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:55:16.13  
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:55:16.13          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:55:16.13      
💀    🚀 zrbShowAdv           ☕ 07:55:16.13  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:55:16.13  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:55:16.13    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:55:16.13    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:55:16.13    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:55:16.131 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:16.66  🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:16.66  Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.823 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.823 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:17.824 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.222 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.235 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.249 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.249 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.249 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.249 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.249 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.257 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.258 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.277 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.278 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.284 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.284 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.291 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.291 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.291 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.291   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.291 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.292 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.361 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.362 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.362 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.905 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:18.906 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.195 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.196 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.196 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.196 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.196 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.196 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.563 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.573 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.582 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.582 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.582 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.582 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.582 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.587 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.587 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.602 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.602 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.605 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.605 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61  Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61  🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61  🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61    "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61  ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.61  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoBackend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoBackendDeployment","ZtplTaskName":"DemoBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: DEMO_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: DEMO_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: DEMO_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: DEMO_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: DEMO_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: DEMO_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: DEMO_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: DEMO_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///database.db\n  from: DEMO_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: DEMO_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: DEMO_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: DEMO_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: DEMO_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: DEMO_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: DEMO_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: DEMO_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: DEMO_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: DEMO_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: DEMO_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: DEMO_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-backend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-backend","ztpl-app-module-name":"","ztpl-app-name":"demo-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-backend-deployment","ztpl-task-name":"demo-backend","ztplAppContainerName":"demoBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"sqlite:///database.db","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoBackendDeployment","ztplDeploymentName":"demoBackendDeployment","ztplDeploymentTaskLocation":"../../demoBackendDeployment","ztplTaskName":"demoBackend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_backend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_backend_deployment","ztpl_task_name":"demo_backend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.64  🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.644 Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.649 Checking prepareDemoBackendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:20.915 Checking deployDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:21.212 Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:21.229 Adding deployDemoBackendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:21.855 Checking destroyDemoBackendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:22.404 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:22.415 Adding destroyDemoBackendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:22.915 Done registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:55:22.915 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:55:22.915 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 6.900578681s
         Current Time: 07:55:23
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 7.012446358s
         Current Time: 07:55:23
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoBackend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 2.945µs
         Current Time: 07:55:23
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:23.777 Current directory is a valid zaruba project
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 07:55:23.794         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 07:55:23.794     
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 07:55:23.794   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 07:55:23.794   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 07:55:23.794   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 07:55:23.794 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🚢 'makeAppHelmDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.535 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.535 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.785 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:24.786 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.351 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.378 Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.394 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.394 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.394 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.394 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.394 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.403 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.404 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.426 Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.426 Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.433 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.433 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.441 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.527 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.527 🎉🎉🎉
💀    🚀 makeAppHelmDeploy... 🚢 07:55:25.527 Done
💀 🎉 Successfully running 🚢 'makeAppHelmDeployment' command
💀 🏁 Run 🚢 'makeAppHelmDeploymentTask' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.145 🧰 Prepare
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.145 Preparing base variables
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.411 Base variables prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.411 Preparing start command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.411 Start command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Preparing prepare command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Prepare command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Preparing test command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Test command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Preparing migrate command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Migrate command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Preparing check command
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Check command prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.412 Preparing replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:26.992 Add config to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.01  Add env to replacement map
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.028 Replacement map prepared
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.028 ✅ Validate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.028 Validate app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.028 Done validating app directory
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.028 Validate app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.035 Done validating app container volumes
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.035 Validate template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.06  Done validating template locations
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.06  Validate app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.073 Done validating app ports
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.073 Validate app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082 Done validating app crud fields
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082 🚧 Generate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082 🚧 Template Location: [
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082   "/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate"
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082 ]
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.082 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"DEMO_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"DemoFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"DemoFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"DemoFrontendDeployment","ZtplTaskName":"DemoFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check demoFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate demoFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare demoFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start demoFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test demoFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"demo-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"demo-frontend","ztpl-app-module-name":"","ztpl-app-name":"demo-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"demo-frontend-deployment","ztpl-task-name":"demo-frontend","ztplAppContainerName":"demoFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"demoFrontend","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"demoFrontend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../demoFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"demoFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDeploymentDirectory":"demoFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appHelmDeployment/deploymentTaskTemplate\"\n]","ztplDeploymentDirectory":"demoFrontendDeployment","ztplDeploymentName":"demoFrontendDeployment","ztplDeploymentTaskLocation":"../../demoFrontendDeployment","ztplTaskName":"demoFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"demo_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"demo_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"demo_frontend_deployment","ztpl_task_name":"demo_frontend"}
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.142 🔩 Integrate
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.15  Registering deployment tasks
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.157 Checking prepareDemoFrontendDeploymentDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:27.573 Checking deployDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:28.12  Checking deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:28.131 Adding deployDemoFrontendDeployment as dependency of deploy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:28.564 Checking destroyDemoFrontendDeployment
💀    🚀 makeAppHelmDeploy... 🚢 07:55:28.998 Checking destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:29.005 Adding destroyDemoFrontendDeployment as dependency of destroy
💀    🚀 makeAppHelmDeploy... 🚢 07:55:29.389 Done registering deployment tasks
💀 🎉 Successfully running 🚢 'makeAppHelmDeploymentTask' command
💀 🎉 Reach 🚢 'addAppHelmDeployment' wrapper
💀 🔎 Job Running...
         Elapsed Time: 5.725765918s
         Current Time: 07:55:29
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 5.836888847s
         Current Time: 07:55:29
zaruba please addAppHelmDeployment -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=demoFrontend' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.387µs
         Current Time: 07:55:29
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:29.974 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔄 'syncEnv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 syncEnv              🔄 07:55:30.084 Synchronize task environments
💀    🚀 syncEnv              🔄 07:55:30.55  Synchronize project's environment files
💀    🚀 syncEnv              🔄 07:55:30.944 🎉🎉🎉
💀    🚀 syncEnv              🔄 07:55:30.944 Environment synchronized
💀 🎉 Successfully running 🔄 'syncEnv' command
💀 🔎 Job Running...
         Elapsed Time: 1.08285423s
         Current Time: 07:55:31
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.284245722s
         Current Time: 07:55:31
zaruba please syncEnv -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.629µs
         Current Time: 07:55:31
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:31.639 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:55:31.769 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:55:31.769 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 241.815262ms
         Current Time: 07:55:31
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 443.927407ms
         Current Time: 07:55:32
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=defaultKubeContext' -v 'variableValue=docker-desktop' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.935µs
         Current Time: 07:55:32
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 07:55:32.554 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🔗 'setProjectValue' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 setProjectValue      🔗 07:55:32.672 🎉🎉🎉
💀    🚀 setProjectValue      🔗 07:55:32.672 Kwarg  :  has been set
💀 🎉 Successfully running 🔗 'setProjectValue' command
💀 🔎 Job Running...
         Elapsed Time: 229.674294ms
         Current Time: 07:55:32
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 431.018791ms
         Current Time: 07:55:32
zaruba please setProjectValue -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'variableName=pulumiUseLocalBackend' -v 'variableValue=true' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.776µs
         Current Time: 07:55:33
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoFronte... 🏁 07:55:33.233 🚧 Create virtual environment.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:33.233 🚧 Create virtual environment.
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 prepareDemoBacken... 🏁 07:55:33.236 🚧 Create virtual environment.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoFronte... 🏁 07:55:36.682 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:36.682 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:55:36.688 🚧 Install pip packages.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:37.17  Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:37.177 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:37.199 Collecting pulumi<4.0.0,>=3.0.0 (from -r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:38.346   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:38.377 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:38.379   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:38.408   Using cached https://files.pythonhosted.org/packages/bf/1f/0b67ccc0308c37b2823287716f0fca00d6fa3d92cce3f85c100ccdeda4c9/pulumi-3.32.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:38.409 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:38.441 Collecting pulumi-kubernetes<4.0.0,>=3.0.0 (from -r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:38.925   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 07:55:38.936   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoBacken... 🏁 07:55:39.052   Using cached https://files.pythonhosted.org/packages/fd/be/a837fd533218b087360f1f492d15c391a7e68b193abeaedefe07470d9cc4/pulumi_kubernetes-3.19.1.tar.gz
💀    🚀 prepareDemoDbDepl... 🏁 07:55:39.398 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:39.404 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:39.514 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:39.604   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:39.645 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:39.744   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:39.765 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:39.87    Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:39.878 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:39.923   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:39.959 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:40       Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:40.009 Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:40.071   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:40.081 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:40.257   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:40.29  Collecting grpcio>=1.33.2 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:40.696   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:40.81  Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:41.191   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:41.305 Collecting protobuf>=3.6.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:41.34    Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:41.382 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:41.48    Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:41.496 Collecting pyyaml>=5.3.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.529   Using cached https://files.pythonhosted.org/packages/4f/5e/90532773aec77041b65c176a95df5a1c7187e38bd54c6ffd91f9e60dcbd7/grpcio-1.46.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.643 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:41.663   Using cached https://files.pythonhosted.org/packages/d7/42/7ad4b6d67a16229496d4f6e74201bdbebcf4bc1e87d5a70c9297d4961bd2/PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:41.695 Collecting dill>=0.3.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.718   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.736 Collecting semver>=2.8.1 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoBacken... 🏁 07:55:41.769   Using cached https://files.pythonhosted.org/packages/b6/c3/973676ceb86b60835bb3978c6db67a5dc06be6cfdbd14ef0f5a13e3fc9fd/dill-0.3.4-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:41.786 Collecting six>=1.12.0 (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1))
💀    🚀 prepareDemoFronte... 🏁 07:55:41.793   Using cached https://files.pythonhosted.org/packages/c1/4d/1d46234fbdff4ee05cb7ec6cb6ea9282769fa9fefd72d93de4b85fd3d8c4/protobuf-3.20.1-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:41.831 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.847   Using cached https://files.pythonhosted.org/packages/0b/70/b84f9944a03964a88031ef6ac219b6c91e8ba2f373362329d8770ef36f02/semver-2.13.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.858 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:41.884   Using cached https://files.pythonhosted.org/packages/d9/5a/e7c31adbe875f2abbb91bd84cf2dc52d792b5a01506781dbcf25c91daf11/six-1.16.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:41.893 Collecting parver>=0.2.1 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:41.906   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:41.926 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.933   Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:41.95  Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42       Using cached https://files.pythonhosted.org/packages/1a/79/aea13e60a54e453df1a45383e92feda3b280e87ebded788c9c818d93e413/parver-0.3.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.023 Collecting requests<3.0,>=2.21 (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:42.075   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:42.103 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.105   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.138 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.157   Using cached https://files.pythonhosted.org/packages/2d/61/08076519c80041bc0ffa1a8af0cbd3bf3e2b62af10435d269a9d0f40564d/requests-2.27.1-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.186 Collecting attrs>=19.2 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:42.19    Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.229   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:42.254 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.284 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.311   Using cached https://files.pythonhosted.org/packages/be/be/7abce643bfdf8ca01c48afa2ddf8308c2308b0c3b239a44e57d020afa0ef/attrs-21.4.0-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.361 Collecting arpeggio~=1.7 (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.561   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.582 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.661   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.687 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.716   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.725 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:42.761   Using cached https://files.pythonhosted.org/packages/1a/ae/a2dfd99042b8952e86ea6cd6ad5ba8b81c3f9f150e24475cf55e09fbe3e4/Arpeggio-1.10.2-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:42.782 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.816   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:42.831 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.835   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.871 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:42.875   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:42.892 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:42.976   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.981   Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:42.993 Collecting charset-normalizer~=2.0.0; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:43.001   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:43.016 Collecting certifi>=2017.4.17 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoBacken... 🏁 07:55:43.016 Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:43.095   Using cached https://files.pythonhosted.org/packages/06/b3/24afc8868eba069a7f03650ac750a778862dc34941a4bebeb58706715726/charset_normalizer-2.0.12-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:43.108   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:43.11  Collecting idna<4,>=2.5; python_version >= "3" (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoFronte... 🏁 07:55:43.11    Using cached https://files.pythonhosted.org/packages/37/45/946c02767aabb873146011e665728b680884cd8fe70dde973c640e45b775/certifi-2021.10.8-py2.py3-none-any.whl
💀    🚀 prepareDemoFronte... 🏁 07:55:43.122 Collecting urllib3<1.27,>=1.21.1 (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2))
💀    🚀 prepareDemoDbDepl... 🏁 07:55:43.198   Using cached https://files.pythonhosted.org/packages/04/a2/d918dcd22354d8958fe113e1a3630137e0fc8b44859ade3063982eacd2a4/idna-3.3-py3-none-any.whl
💀    🚀 prepareDemoBacken... 🏁 07:55:43.228 Installing collected packages: six, grpcio, protobuf, semver, pyyaml, dill, pulumi, attrs, arpeggio, parver, certifi, charset-normalizer, urllib3, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 07:55:43.272   Using cached https://files.pythonhosted.org/packages/ec/03/062e6444ce4baf1eac17a6a0ebfe36bb1ad05e1df0e20b110de59c278498/urllib3-1.26.9-py2.py3-none-any.whl
💀    🚀 prepareDemoDbDepl... 🏁 07:55:43.352 Installing collected packages: protobuf, six, pyyaml, grpcio, dill, semver, pulumi, attrs, arpeggio, parver, urllib3, certifi, charset-normalizer, idna, requests, pulumi-kubernetes
💀    🚀 prepareDemoFronte... 🏁 07:55:43.462 Installing collected packages: pyyaml, dill, six, semver, grpcio, protobuf, pulumi, attrs, arpeggio, parver, charset-normalizer, idna, certifi, urllib3, requests, pulumi-kubernetes
💀    🚀 prepareDemoBacken... 🏁 07:55:44.313   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoDbDepl... 🏁 07:55:44.513   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoFronte... 🏁 07:55:44.615   Running setup.py install for pulumi-kubernetes: started
💀    🚀 prepareDemoBacken... 🏁 07:55:47.966     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoDbDepl... 🏁 07:55:47.967     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoFronte... 🏁 07:55:47.971     Running setup.py install for pulumi-kubernetes: finished with status 'done'
💀    🚀 prepareDemoBacken... 🏁 07:55:48.037 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.038 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀    🚀 prepareDemoFronte... 🏁 07:55:48.045 Successfully installed arpeggio-1.10.2 attrs-21.4.0 certifi-2021.10.8 charset-normalizer-2.0.12 dill-0.3.4 grpcio-1.46.0 idna-3.3 parver-0.3.1 protobuf-3.20.1 pulumi-3.32.1 pulumi-kubernetes-3.19.1 pyyaml-6.0 requests-2.27.1 semver-2.13.0 six-1.16.0 urllib3-1.26.9
💀 🔥 🚀 prepareDemoBacken... 🏁 07:55:48.078 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:55:48.078 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:55:48.085 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:55:48.085 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:55:48.091 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:55:48.091 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.347 🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 07:55:48.347 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.482 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.482 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.606       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.607 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.608 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.608       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.609 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609       repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609     - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609       version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.609 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61      dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61        repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  for this case.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Usage:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Flags:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61    -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  
💀    🚀 prepareDemoFronte... 🏁 07:55:48.61  Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611 
💀    🚀 prepareDemoFronte... 🏁 07:55:48.611 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:55:48.613 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 07:55:48.614 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'deployDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀 🏁 Run 🏁 'deployDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🔥 🚀 deployDemoFronten... 🏁 07:55:48.928 error: no stack named 'dev' found
💀 🔥 🚀 deployDemoDbDeplo... 🏁 07:55:49.119 error: no stack named 'dev' found
💀    🚀 deployDemoFronten... 🏁 07:55:49.153 Created stack 'dev'
💀    🚀 deployDemoDbDeplo... 🏁 07:55:49.285 Created stack 'dev'
💀    🚀 prepareDemoBacken... 🏁 07:55:49.746 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:55:49.848 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 07:55:49.849 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92      # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92      dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92      - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92        version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92        repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92      - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92        version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.92  
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 
💀    🚀 prepareDemoBacken... 🏁 07:55:49.921 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 07:55:49.924 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'deployDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🔥 🚀 deployDemoBackend... 🏁 07:55:50.262 error: no stack named 'dev' found
💀    🚀 deployDemoBackend... 🏁 07:55:50.431 Created stack 'dev'
💀    🚀 deployDemoFronten... 🏁 07:55:51.822 Previewing update (dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:55:51.956 Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:55:52.61  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:52.689 
💀    🚀 deployDemoBackend... 🏁 07:55:53.15  Previewing update (dev):
💀    🚀 deployDemoFronten... 🏁 07:55:53.395  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:53.488  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:53.785  +  kubernetes:helm.sh/v3:Chart demo-db create 
💀    🚀 deployDemoFronten... 🏁 07:55:53.79   +  kubernetes:helm.sh/v3:Chart demo-frontend create 
💀    🚀 deployDemoBackend... 🏁 07:55:54.14  
💀    🚀 deployDemoFronten... 🏁 07:55:54.378  +  kubernetes:core/v1:ServiceAccount default/demo-frontend create 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.378  +  kubernetes:core/v1:ServiceAccount default/demo-db create 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.385  +  kubernetes:apps/v1:Deployment default/demo-db create 
💀    🚀 deployDemoFronten... 🏁 07:55:54.385  +  kubernetes:apps/v1:Deployment default/demo-frontend create 
💀    🚀 deployDemoFronten... 🏁 07:55:54.643  +  pulumi:pulumi:Stack demoFrontendDeployment-dev create 
💀    🚀 deployDemoFronten... 🏁 07:55:54.643  
💀    🚀 deployDemoFronten... 🏁 07:55:54.643 Resources:
💀    🚀 deployDemoFronten... 🏁 07:55:54.643     + 4 to create
💀    🚀 deployDemoFronten... 🏁 07:55:54.643 
💀    🚀 deployDemoFronten... 🏁 07:55:54.643 Updating (dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.653  +  pulumi:pulumi:Stack demoDbDeployment-dev create 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.654  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.654 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.654     + 4 to create
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.654 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:54.654 Updating (dev):
💀    🚀 deployDemoBackend... 🏁 07:55:54.978  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:55:55.124  +  kubernetes:helm.sh/v3:Chart demo-backend create 
💀    🚀 deployDemoFronten... 🏁 07:55:55.53  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:55.53  
💀    🚀 deployDemoBackend... 🏁 07:55:55.687  +  kubernetes:core/v1:ServiceAccount default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:55:55.689  +  kubernetes:core/v1:Service default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:55:55.714  +  kubernetes:apps/v1:Deployment default/demo-backend create 
💀    🚀 deployDemoBackend... 🏁 07:55:55.974  +  pulumi:pulumi:Stack demoBackendDeployment-dev create 
💀    🚀 deployDemoBackend... 🏁 07:55:55.974  
💀    🚀 deployDemoBackend... 🏁 07:55:55.974 Resources:
💀    🚀 deployDemoBackend... 🏁 07:55:55.974     + 5 to create
💀    🚀 deployDemoBackend... 🏁 07:55:55.974 
💀    🚀 deployDemoBackend... 🏁 07:55:55.974 Updating (dev):
💀    🚀 deployDemoFronten... 🏁 07:55:56.323  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:56.327  +  pulumi:pulumi:Stack demoDbDeployment-dev creating 
💀    🚀 deployDemoFronten... 🏁 07:55:56.462  +  kubernetes:helm.sh/v3:Chart demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:56.467  +  kubernetes:helm.sh/v3:Chart demo-db creating 
💀    🚀 deployDemoBackend... 🏁 07:55:56.856 
💀    🚀 deployDemoFronten... 🏁 07:55:57.022  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.027  +  kubernetes:core/v1:ServiceAccount default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:55:57.03   +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.033  +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.064  +  kubernetes:core/v1:ServiceAccount default/demo-db creating Retry #0; creation failed: serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.064  +  kubernetes:core/v1:ServiceAccount default/demo-db creating error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.064  +  kubernetes:core/v1:ServiceAccount default/demo-db **creating failed** error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoFronten... 🏁 07:55:57.07   +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating Retry #0; creation failed: serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 07:55:57.07   +  kubernetes:apps/v1:Deployment default/demo-frontend creating 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.07   +  kubernetes:apps/v1:Deployment default/demo-db creating 
💀    🚀 deployDemoFronten... 🏁 07:55:57.071  +  kubernetes:core/v1:ServiceAccount default/demo-frontend creating error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 07:55:57.071  +  kubernetes:core/v1:ServiceAccount default/demo-frontend **creating failed** error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 07:55:57.085  +  kubernetes:apps/v1:Deployment default/demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 07:55:57.085  +  pulumi:pulumi:Stack demoFrontendDeployment-dev creating error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.088  +  kubernetes:apps/v1:Deployment default/demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.088  +  pulumi:pulumi:Stack demoDbDeployment-dev creating error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106  +  pulumi:pulumi:Stack demoDbDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106  +  kubernetes:helm.sh/v3:Chart demo-db created 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106 Diagnostics:
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106   pulumi:pulumi:Stack (demoDbDeployment-dev):
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106     error: update failed
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106   kubernetes:core/v1:ServiceAccount (default/demo-db):
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106     error: resource default/demo-db was not successfully created by the Kubernetes API server : serviceaccounts "demo-db" already exists
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106  
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106 Resources:
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106     + 3 created
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106 
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106 Duration: 2s
💀    🚀 deployDemoDbDeplo... 🏁 07:55:57.106 
💀    🚀 deployDemoFronten... 🏁 07:55:57.111  +  pulumi:pulumi:Stack demoFrontendDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoFronten... 🏁 07:55:57.111  +  kubernetes:helm.sh/v3:Chart demo-frontend created 
💀    🚀 deployDemoFronten... 🏁 07:55:57.111  
💀    🚀 deployDemoFronten... 🏁 07:55:57.111 Diagnostics:
💀    🚀 deployDemoFronten... 🏁 07:55:57.111   pulumi:pulumi:Stack (demoFrontendDeployment-dev):
💀    🚀 deployDemoFronten... 🏁 07:55:57.111     error: update failed
💀    🚀 deployDemoFronten... 🏁 07:55:57.111  
💀    🚀 deployDemoFronten... 🏁 07:55:57.111   kubernetes:core/v1:ServiceAccount (default/demo-frontend):
💀    🚀 deployDemoFronten... 🏁 07:55:57.111     error: resource default/demo-frontend was not successfully created by the Kubernetes API server : serviceaccounts "demo-frontend" already exists
💀    🚀 deployDemoFronten... 🏁 07:55:57.111  
💀    🚀 deployDemoFronten... 🏁 07:55:57.111 Resources:
💀    🚀 deployDemoFronten... 🏁 07:55:57.111     + 3 created
💀    🚀 deployDemoFronten... 🏁 07:55:57.111 
💀    🚀 deployDemoFronten... 🏁 07:55:57.111 Duration: 2s
💀    🚀 deployDemoFronten... 🏁 07:55:57.111 
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
💀 🔪 Kill 🏁 'deployDemoBackendDeployment' command (PID=562)
💀    🚀 deployDemoBackend... 🏁 07:55:57.506  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating 
💀    🚀 deployDemoBackend... 🏁 07:55:57.623  +  pulumi:pulumi:Stack demoBackendDeployment-dev creating error: update canceled
💀    🚀 deployDemoBackend... 🏁 07:55:57.632  +  pulumi:pulumi:Stack demoBackendDeployment-dev **creating failed** 1 error
💀    🚀 deployDemoBackend... 🏁 07:55:57.632  
💀    🚀 deployDemoBackend... 🏁 07:55:57.633 Diagnostics:
💀    🚀 deployDemoBackend... 🏁 07:55:57.633   pulumi:pulumi:Stack (demoBackendDeployment-dev):
💀    🚀 deployDemoBackend... 🏁 07:55:57.633     error: update canceled
💀    🚀 deployDemoBackend... 🏁 07:55:57.633  
💀    🚀 deployDemoBackend... 🏁 07:55:57.633 Resources:
💀    🚀 deployDemoBackend... 🏁 07:55:57.633     + 1 created
💀    🚀 deployDemoBackend... 🏁 07:55:57.633 
💀    🚀 deployDemoBackend... 🏁 07:55:57.633 Duration: 1s
💀    🚀 deployDemoBackend... 🏁 07:55:57.633 
💀 🔥 Error running 🏁 'deployDemoBackendDeployment' command: exit status 255
💀 🔎 Job Ended...
         Elapsed Time: 24.906809768s
         Current Time: 07:55:58
zaruba please deploy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["deploy"]
🔥 Stderr    : exit status 255
💀 🔎 Job Starting...
         Elapsed Time: 1.711µs
         Current Time: 07:55:58
💀 🏁 Run 🚢 'zrbSetKubeContext' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run 🏁 'prepareDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀 🏁 Run 🏁 'prepareDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀 🏁 Run 🏁 'prepareDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 07:55:58.726 🚧 Install pip packages.
💀    🚀 prepareDemoFronte... 🏁 07:55:58.743 🚧 Install pip packages.
💀    🚀 prepareDemoBacken... 🏁 07:55:58.745 🚧 Install pip packages.
💀 🎉 Successfully running 🚢 'zrbSetKubeContext' command
💀    🚀 prepareDemoBacken... 🏁 07:55:59.522 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.534 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.544 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.547 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.548 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.559 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.559 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.561 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.569 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.569 Requirement already satisfied: pulumi<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 1)) (3.32.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.572 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.575 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.575 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.578 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.578 Requirement already satisfied: pulumi-kubernetes<4.0.0,>=3.0.0 in ./venv/lib/python3.8/site-packages (from -r requirements.txt (line 2)) (3.19.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.581 Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.588 Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.591 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.592 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.595 Requirement already satisfied: grpcio>=1.33.2 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.46.0)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.6   Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.6   Requirement already satisfied: six>=1.12.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (1.16.0)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.603 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.608 Requirement already satisfied: semver>=2.8.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (2.13.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.61  Requirement already satisfied: dill>=0.3.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (0.3.4)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.617 Requirement already satisfied: protobuf>=3.6.0 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (3.20.1)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.619 Requirement already satisfied: pyyaml>=5.3.1 in ./venv/lib/python3.8/site-packages (from pulumi<4.0.0,>=3.0.0->-r requirements.txt (line 1)) (6.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.623 Requirement already satisfied: parver>=0.2.1 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (0.3.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.629 Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.633 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.65  Requirement already satisfied: requests<3.0,>=2.21 in ./venv/lib/python3.8/site-packages (from pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.27.1)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.664 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.683 Requirement already satisfied: attrs>=19.2 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (21.4.0)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.7   Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.724 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.733 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.736 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.737 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoBacken... 🏁 07:55:59.739 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.747 Requirement already satisfied: arpeggio~=1.7 in ./venv/lib/python3.8/site-packages (from parver>=0.2.1->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.10.2)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.751 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.756 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.761 Requirement already satisfied: charset-normalizer~=2.0.0; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2.0.12)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.763 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoDbDepl... 🏁 07:55:59.766 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.769 Requirement already satisfied: idna<4,>=2.5; python_version >= "3" in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (3.3)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.773 Requirement already satisfied: certifi>=2017.4.17 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (2021.10.8)
💀    🚀 prepareDemoFronte... 🏁 07:55:59.776 Requirement already satisfied: urllib3<1.27,>=1.21.1 in ./venv/lib/python3.8/site-packages (from requests<3.0,>=2.21->pulumi-kubernetes<4.0.0,>=3.0.0->-r requirements.txt (line 2)) (1.26.9)
💀 🔥 🚀 prepareDemoBacken... 🏁 07:55:59.835 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoBacken... 🏁 07:55:59.835 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:55:59.877 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoDbDepl... 🏁 07:55:59.877 You should consider upgrading via the 'pip install --upgrade pip' command.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:55:59.878 WARNING: You are using pip version 19.2.3, however version 22.0.4 is available.
💀 🔥 🚀 prepareDemoFronte... 🏁 07:55:59.878 You should consider upgrading via the 'pip install --upgrade pip' command.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.222 🚧 Deployment config: {"env":[{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-frontend","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoFronte... 🏁 07:56:00.222 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.309 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.309 Manage the dependencies of a chart.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  dependencies.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  'charts/' directory.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31      # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31      dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31      - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31        version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31        repository: "https://example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31      - name: memcached
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31        version: "3.2.1"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31        repository: "https://another.example.com/charts"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  with 'alias:' or '@'.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoFronte... 🏁 07:56:00.31  "file://". For example,
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311     # Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311     dependencies:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311     - name: nginx
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       version: "1.2.3"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 for this case.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 Usage:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   helm dependency [command]
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 Aliases:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   dependency, dep, dependencies
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 Available Commands:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   list        list the dependencies for the given chart
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 Flags:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311   -h, --help   help for dependency
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311 Global Flags:
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --debug                       enable verbose output
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoFronte... 🏁 07:56:00.311       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312 
💀    🚀 prepareDemoFronte... 🏁 07:56:00.312 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoFronte... 🏁 07:56:00.315 🚧 Preparation completed.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.407 🚧 Deployment config: {"env":[{"name":"MYSQL_DATABASE","value":"sample"},{"name":"MYSQL_PASSWORD","value":"mysql"},{"name":"MYSQL_ROOT_PASSWORD","value":"Alch3mist"},{"name":"MYSQL_USER","value":"mysql"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"}],"image.repository":"demo-db","image.tag":"latest","namespace":"default","ports":[],"replicaCount":1,"service.ports":[]}
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.407 🚧 Prepare chart dependencies.
💀 🎉 Successfully running 🏁 'prepareDemoFrontendDeployment' command
💀 🏁 Run 🏁 'destroyDemoFrontendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoFrontendDeployment
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.553 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.553 Manage the dependencies of a chart.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 dependencies.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 'charts/' directory.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554       repository: "https://example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554     - name: memcached
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554       version: "3.2.1"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.554 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 with 'alias:' or '@'.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 "file://". For example,
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555     # Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555     dependencies:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555     - name: nginx
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       version: "1.2.3"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 for this case.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Usage:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   helm dependency [command]
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Aliases:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   dependency, dep, dependencies
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Available Commands:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   list        list the dependencies for the given chart
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   -h, --help   help for dependency
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Global Flags:
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --debug                       enable verbose output
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.555 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoDbDepl... 🏁 07:56:00.558 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoDbDeployment' command
💀 🏁 Run 🏁 'destroyDemoDbDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoDbDeployment
💀    🚀 prepareDemoBacken... 🏁 07:56:02.127 PARTS: ["3000"]
💀    🚀 prepareDemoBacken... 🏁 07:56:02.236 🚧 Deployment config: {"env":[{"name":"APP_ACCESS_TOKEN_ALGORITHM","value":"HS256"},{"name":"APP_ACCESS_TOKEN_EXPIRE_MINUTES","value":"30"},{"name":"APP_ACCESS_TOKEN_SECRET_KEY","value":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"},{"name":"APP_ACCESS_TOKEN_URL","value":"/token/"},{"name":"APP_ENABLE_EVENT_HANDLER","value":"1"},{"name":"APP_ENABLE_ROUTE_HANDLER","value":"1"},{"name":"APP_ENABLE_RPC_HANDLER","value":"1"},{"name":"APP_ERROR_THRESHOLD","value":"10"},{"name":"APP_GUEST_USERNAME","value":"guest"},{"name":"APP_HTTP_PORT","value":"3000"},{"name":"APP_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"APP_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"APP_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"APP_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"APP_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"APP_MESSAGE_BUS_TYPE","value":"local"},{"name":"APP_RABBITMQ_HOST","value":"localhost"},{"name":"APP_RABBITMQ_PASS","value":"Alch3mist"},{"name":"APP_RABBITMQ_USER","value":"root"},{"name":"APP_RABBITMQ_VHOST","value":"/"},{"name":"APP_ROOT_INITIAL_EMAIL","value":"root@innistrad.com"},{"name":"APP_ROOT_INITIAL_FULL_NAME","value":"root"},{"name":"APP_ROOT_INITIAL_PASSWORD","value":"Alch3mist"},{"name":"APP_ROOT_INITIAL_PHONE_NUMBER","value":"+621234567890"},{"name":"APP_ROOT_PERMISSION","value":"root"},{"name":"APP_ROOT_USERNAME","value":"root"},{"name":"APP_RPC_TYPE","value":"local"},{"name":"APP_SQLALCHEMY_DATABASE_URL","value":"sqlite:///database.db"},{"name":"APP_STATIC_DIRECTORY","value":""},{"name":"APP_STATIC_URL","value":"/static"},{"name":"PULUMI_BACKEND_URL","value":""},{"name":"PULUMI_CONFIG_PASSPHRASE","value":"defaultLocalPulumiPassphrase"},{"name":"PYTHONUNBUFFERED","value":"1"},{"name":"TEST_INTEGRATION","value":"0"},{"name":"TEST_KAFKA_BOOTSTRAP_SERVERS","value":"localhost:9092"},{"name":"TEST_KAFKA_SASL_MECHANISM","value":"PLAIN"},{"name":"TEST_KAFKA_SASL_PLAIN_PASSWORD","value":""},{"name":"TEST_KAFKA_SASL_PLAIN_USERNAME","value":""},{"name":"TEST_KAFKA_SCHEMA_REGISTRY","value":"http://localhost:8081"},{"name":"TEST_RABBITMQ_HOST","value":"localhost"},{"name":"TEST_RABBITMQ_PASS","value":"Alch3mist"},{"name":"TEST_RABBITMQ_USER","value":"root"},{"name":"TEST_RABBITMQ_VHOST","value":"/"},{"name":"TEST_SQLALCHEMY_DATABASE_URL","value":"sqlite:///test.db"}],"image.repository":"demo-backend","image.tag":"latest","namespace":"default","ports":[{"containerPort":3000,"name":"port0","protocol":"TCP"}],"replicaCount":1,"service.ports":[{"name":"port0","port":3000,"protocol":"TCP","targetPort":"port0"}]}
💀    🚀 prepareDemoBacken... 🏁 07:56:02.236 🚧 Prepare chart dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.314 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.314 Manage the dependencies of a chart.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.314 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 Helm charts store their dependencies in 'charts/'. For chart developers, it is
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 often easier to manage dependencies in 'Chart.yaml' which declares all
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 dependencies.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 The dependency commands operate on that file, making it easy to synchronize
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 between the desired dependencies and the actual dependencies stored in the
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 'charts/' directory.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 For example, this Chart.yaml declares two dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       repository: "https://example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     - name: memcached
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       version: "3.2.1"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       repository: "https://another.example.com/charts"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 The 'name' should be the name of a chart, where that name must match the name
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 in that chart's 'Chart.yaml' file.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 The 'version' field should contain a semantic version or version range.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 The 'repository' URL should point to a Chart Repository. Helm expects that by
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 appending '/index.yaml' to the URL, it should be able to retrieve the chart
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 repository's index. Note: 'repository' can be an alias. The alias must start
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 with 'alias:' or '@'.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 Starting from 2.2.0, repository can be defined as the path to the directory of
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 the dependency charts stored locally. The path should start with a prefix of
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 "file://". For example,
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     # Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     dependencies:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315     - name: nginx
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       version: "1.2.3"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315       repository: "file://../dependency_chart/nginx"
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 If the dependency chart is retrieved locally, it is not required to have the
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 repository added to helm by "helm add repo". Version matching is also supported
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 for this case.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 Usage:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315   helm dependency [command]
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 Aliases:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315   dependency, dep, dependencies
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315 Available Commands:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.315   build       rebuild the charts/ directory based on the Chart.lock file
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316   list        list the dependencies for the given chart
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316   update      update charts/ based on the contents of Chart.yaml
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 Flags:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316   -h, --help   help for dependency
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 Global Flags:
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --debug                       enable verbose output
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-apiserver string       the address and the port for the Kubernetes API server
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-as-user string         username to impersonate for the operation
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-context string         name of the kubeconfig context to use
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kube-token string           bearer token used for authentication
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --kubeconfig string           path to the kubeconfig file
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316   -n, --namespace string            namespace scope for this request
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --registry-config string      path to the registry config file (default "/home/gofrendi/.config/helm/registry/config.json")
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --repository-cache string     path to the file containing cached repository indexes (default "/home/gofrendi/.cache/helm/repository")
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316       --repository-config string    path to the file containing repository names and URLs (default "/home/gofrendi/.config/helm/repositories.yaml")
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 
💀    🚀 prepareDemoBacken... 🏁 07:56:02.316 Use "helm dependency [command] --help" for more information about a command.
💀    🚀 prepareDemoBacken... 🏁 07:56:02.317 🚧 Preparation completed.
💀 🎉 Successfully running 🏁 'prepareDemoBackendDeployment' command
💀 🏁 Run 🏁 'destroyDemoBackendDeployment' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/demoBackendDeployment
💀    🚀 destroyDemoFronte... 🏁 07:56:03.039 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 07:56:03.208 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.212  -  kubernetes:apps/v1:Deployment default/demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.213  -  kubernetes:helm.sh/v3:Chart demo-frontend delete 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.213  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.214  -  pulumi:pulumi:Stack demoFrontendDeployment-dev delete 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.214  
💀    🚀 destroyDemoFronte... 🏁 07:56:03.215 Resources:
💀    🚀 destroyDemoFronte... 🏁 07:56:03.215     - 3 to delete
💀    🚀 destroyDemoFronte... 🏁 07:56:03.215 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.215 Destroying (dev):
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.224 Previewing destroy (dev):
💀    🚀 destroyDemoFronte... 🏁 07:56:03.348 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.35   -  kubernetes:apps/v1:Deployment default/demo-frontend deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.422 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.429  -  kubernetes:apps/v1:Deployment default/demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432  -  kubernetes:helm.sh/v3:Chart demo-db delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432  -  pulumi:pulumi:Stack demoDbDeployment-dev delete 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432  
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432     - 3 to delete
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.432 Destroying (dev):
💀    🚀 destroyDemoFronte... 🏁 07:56:03.558  -  kubernetes:apps/v1:Deployment default/demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.558  -  kubernetes:helm.sh/v3:Chart demo-frontend deleting 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.558  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleting 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564  -  kubernetes:helm.sh/v3:Chart demo-frontend deleted 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564  -  pulumi:pulumi:Stack demoFrontendDeployment-dev deleted 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564  
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 Resources:
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564     - 3 deleted
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 Duration: 1s
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoFronte... 🏁 07:56:03.564 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoFronte... 🏁 07:56:03.567 hello world
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.637 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.638  -  kubernetes:apps/v1:Deployment default/demo-db deleting 
💀 🎉 Successfully running 🏁 'destroyDemoFrontendDeployment' command
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.876  -  kubernetes:apps/v1:Deployment default/demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.883  -  kubernetes:helm.sh/v3:Chart demo-db deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.884  -  pulumi:pulumi:Stack demoDbDeployment-dev deleting 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.892  -  kubernetes:helm.sh/v3:Chart demo-db deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.892  -  pulumi:pulumi:Stack demoDbDeployment-dev deleted 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893  
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 Resources:
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893     - 3 deleted
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 Duration: 1s
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.893 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoDbDepl... 🏁 07:56:03.896 hello world
💀 🎉 Successfully running 🏁 'destroyDemoDbDeployment' command
💀    🚀 destroyDemoBacken... 🏁 07:56:04.918 Previewing destroy (dev):
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919  -  pulumi:pulumi:Stack demoBackendDeployment-dev delete 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919  
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919 Resources:
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919     - 1 to delete
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.919 Destroying (dev):
💀    🚀 destroyDemoBacken... 🏁 07:56:04.92  
💀    🚀 destroyDemoBacken... 🏁 07:56:04.92   -  pulumi:pulumi:Stack demoBackendDeployment-dev deleting 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921  -  pulumi:pulumi:Stack demoBackendDeployment-dev deleted 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921  
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921 Resources:
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921     - 1 deleted
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921 Duration: 1s
💀    🚀 destroyDemoBacken... 🏁 07:56:04.921 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.922 The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
💀    🚀 destroyDemoBacken... 🏁 07:56:04.922 If you want to remove the stack completely, run 'pulumi stack rm dev'.
💀    🚀 destroyDemoBacken... 🏁 07:56:04.923 hello world
💀 🎉 Successfully running 🏁 'destroyDemoBackendDeployment' command
💀 🏁 Run ❌ 'destroy' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 destroy              ❌ 07:56:05.031 
💀 🎉 Successfully running ❌ 'destroy' command
💀 🔎 Job Running...
         Elapsed Time: 6.470366326s
         Current Time: 07:56:05
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.581193499s
         Current Time: 07:56:05
zaruba please destroy -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

