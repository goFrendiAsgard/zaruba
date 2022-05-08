<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# ❇️ From Zero to Cloud
<!--endTocHeader-->

# A Use Case

Suppose you want to build a simple book catalogue system.

In your first iteration, you want to deploy your book catalogue as a web application. But in the future, you also want to build a mobile app version as well.

Furthermore, you also want to some relevant information in your website. For example, you want to show company profile, office location, etc.

Thus, you decide to split up your system into three components:

* 🐍 `Book Catalogue API`
* 🐸 `Static web server`
* 🐬 `MySQL server`.

![Application components](images/from-zero-to-cloud-architecture.png)

# Discover Dependencies

Your 🐸 `Static web server` might not only serve book catalogue. It also show company profile and other information. Thus, you want your 🐸 `Static web server` to be independent from other components.

In the other hand, your 🐍 `Book Catalogue API` is pretty unusable once the 🐬 `MySQL server` is down. In this case, you can say that your `Book Catalogue API` __depends on__ `MySQL Server`.

![Component dependencies](images/from-zero-to-cloud-dependencies.png)

# Create a Project

<!--startCode-->
```bash
mkdir -p examples/playground/myEndToEndDemo
cd examples/playground/myEndToEndDemo
zaruba please initProject
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.143µs
         Current Time: 15:19:04
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 initProject          🚧 15:19:04.333 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.git/
💀    🚀 initProject          🚧 15:19:04.34  🎉🎉🎉
💀    🚀 initProject          🚧 15:19:04.34  Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 115.594764ms
         Current Time: 15:19:04
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 316.419297ms
         Current Time: 15:19:04
zaruba please initProject
```````
</details>
<!--endCode-->

# Add MySQL

<!--startCode-->
```bash
cd examples/playground/myEndToEndDemo
zaruba please addMysql appDirectory=myDb
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.699µs
         Current Time: 15:19:04
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 15:19:04.817 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  Hello Human, 
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 15:19:04.82          '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 15:19:04.82      
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 15:19:04.82    * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 15:19:04.82    * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 15:19:04.82    * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 15:19:04.82  
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 🐬 'makeMysqlApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlApp         🐬 15:19:05.258 🧰 Prepare
💀    🚀 makeMysqlApp         🐬 15:19:05.258 Preparing base variables
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Base variables prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Preparing start command
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Start command prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Preparing prepare command
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Prepare command prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Preparing test command
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Test command prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Preparing migrate command
💀    🚀 makeMysqlApp         🐬 15:19:05.334 Migrate command prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.335 Preparing check command
💀    🚀 makeMysqlApp         🐬 15:19:05.335 Check command prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.335 Preparing replacement map
💀    🚀 makeMysqlApp         🐬 15:19:05.542 Add config to replacement map
💀    🚀 makeMysqlApp         🐬 15:19:05.549 Add env to replacement map
💀    🚀 makeMysqlApp         🐬 15:19:05.556 Replacement map prepared
💀    🚀 makeMysqlApp         🐬 15:19:05.556 ✅ Validate
💀    🚀 makeMysqlApp         🐬 15:19:05.556 Validate app directory
💀    🚀 makeMysqlApp         🐬 15:19:05.556 Done validating app directory
💀    🚀 makeMysqlApp         🐬 15:19:05.556 Validate app container volumes
💀    🚀 makeMysqlApp         🐬 15:19:05.56  Done validating app container volumes
💀    🚀 makeMysqlApp         🐬 15:19:05.56  Validate template locations
💀    🚀 makeMysqlApp         🐬 15:19:05.57  Done validating template locations
💀    🚀 makeMysqlApp         🐬 15:19:05.57  Validate app ports
💀    🚀 makeMysqlApp         🐬 15:19:05.573 Done validating app ports
💀    🚀 makeMysqlApp         🐬 15:19:05.573 Validate app crud fields
💀    🚀 makeMysqlApp         🐬 15:19:05.577 Done validating app crud fields
💀    🚀 makeMysqlApp         🐬 15:19:05.577 🚧 Generate
💀    🚀 makeMysqlApp         🐬 15:19:05.577 🚧 Template Location: [
💀    🚀 makeMysqlApp         🐬 15:19:05.577   "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate"
💀    🚀 makeMysqlApp         🐬 15:19:05.577 ]
💀    🚀 makeMysqlApp         🐬 15:19:05.577 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"MyDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"MyDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyDbDeployment","ZtplTaskName":"MyDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myDb\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start myDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myDb\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":{},"[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"my-db","ztpl-app-event-name":"","ztpl-app-image-name":"my-db","ztpl-app-module-name":"","ztpl-app-name":"my-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-db-deployment","ztpl-task-name":"my-db","ztplAppContainerName":"myDb","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"myDb","ztplAppEnvs":{},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"myDb","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"myDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"myDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appTemplate\"\n]","ztplDeploymentDirectory":"myDbDeployment","ztplDeploymentName":"myDbDeployment","ztplDeploymentTaskLocation":"../../myDbDeployment","ztplTaskName":"myDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"my_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"my_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_db_deployment","ztpl_task_name":"my_db"}
💀    🚀 makeMysqlApp         🐬 15:19:05.592 🔩 Integrate
💀    🚀 makeMysqlApp         🐬 15:19:05.592 🎉🎉🎉
💀    🚀 makeMysqlApp         🐬 15:19:05.592 Done
💀 🎉 Successfully running 🐬 'makeMysqlApp' command
💀 🏁 Run 🐬 'makeMysqlAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeMysqlAppRunner   🐬 15:19:05.99  🧰 Prepare
💀    🚀 makeMysqlAppRunner   🐬 15:19:05.99  Preparing base variables
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Base variables prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing start command
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Start command prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing prepare command
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Prepare command prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing test command
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Test command prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing migrate command
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Migrate command prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing check command
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Check command prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.168 Preparing replacement map
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.393 Add config to replacement map
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.399 Add env to replacement map
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.405 Replacement map prepared
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.405 ✅ Validate
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.405 Validate app directory
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.405 Done validating app directory
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.405 Validate app container volumes
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.408 Done validating app container volumes
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.408 Validate template locations
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.423 Done validating template locations
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.423 Validate app ports
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.426 Done validating app ports
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.426 Validate app crud fields
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.429 Done validating app crud fields
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43  🚧 Generate
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43  🚧 Template Location: [
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template",
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43    "/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate"
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43  ]
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.43  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_DB","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"MyDb","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"MyDb","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyDbDeployment","ZtplTaskName":"MyDb","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myDb\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myDb\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start myDb\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myDb\"","[\\t ]*ztplAppYamlContainerVolumes":"initdb.d:/docker-entrypoint-initdb.d","[\\t ]*ztplAppYamlEnvs":"MYSQL_DATABASE:\n  default: sample\n  from: MY_DB_MYSQL_DATABASE\nMYSQL_PASSWORD:\n  default: mysql\n  from: MY_DB_MYSQL_PASSWORD\nMYSQL_ROOT_PASSWORD:\n  default: Alch3mist\n  from: MY_DB_MYSQL_ROOT_PASSWORD\nMYSQL_USER:\n  default: mysql\n  from: MY_DB_MYSQL_USER","[\\t ]*ztplAppYamlPorts":3306,"ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"my-db","ztpl-app-event-name":"","ztpl-app-image-name":"my-db","ztpl-app-module-name":"","ztpl-app-name":"my-db","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-db-deployment","ztpl-task-name":"my-db","ztplAppContainerName":"myDb","ztplAppContainerVolumes":["initdb.d:/docker-entrypoint-initdb.d"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"myDb","ztplAppEnvs":{"MYSQL_DATABASE":"sample","MYSQL_PASSWORD":"mysql","MYSQL_ROOT_PASSWORD":"Alch3mist","MYSQL_USER":"mysql"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🐬","ztplAppModuleName":"","ztplAppName":"myDb","ztplAppPorts":["3306"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myDb","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"mysql -u \"root\" -p{{ .GetEnv \"MYSQL_ROOT_PASSWORD\" }} -e \"SHOW SCHEMAS\"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"myDb","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"🐬","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\n  \"3306\"\n]\n","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"initdb.d:/docker-entrypoint-initdb.d\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoMysql","ztplCfgDefaultAppPorts":"[\n  \"3306\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoMysqlDeployment","ztplCfgDeploymentDirectory":"myDbDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/mysql/appRunnerTemplate\"\n]","ztplDeploymentDirectory":"myDbDeployment","ztplDeploymentName":"myDbDeployment","ztplDeploymentTaskLocation":"../../myDbDeployment","ztplTaskName":"myDb","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"my_db","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"my_db","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_db_deployment","ztpl_task_name":"my_db"}
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.468 🔩 Integrate
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.471 Registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.475 Checking prepareMyDb
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.637 Checking testMyDb
💀    🚀 makeMysqlAppRunner   🐬 15:19:06.828 Checking migrateMyDb
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.01  Checking startMyDb
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.153 Checking start
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.156 Adding startMyDb as dependency of start
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.321 Checking startMyDbContainer
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.481 Checking startContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.484 Adding startMyDbContainer as dependency of startContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.628 Checking runMyDb
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.805 Checking runMyDbContainer
💀    🚀 makeMysqlAppRunner   🐬 15:19:07.96  Checking stopMyDbContainer
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.113 Checking stopContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.116 Adding stopMyDbContainer as dependency of stopContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.26  Checking removeMyDbContainer
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.497 Checking removeContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.502 Adding removeMyDbContainer as dependency of removeContainers
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.72  Checking buildMyDbImage
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.866 Checking buildImages
💀    🚀 makeMysqlAppRunner   🐬 15:19:08.869 Adding buildMyDbImage as dependency of buildImages
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.055 Checking pushMyDbImage
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.228 Checking pushImages
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.231 Adding pushMyDbImage as dependency of pushImages
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.373 Checking pullMyDbImage
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.52  Done registering app runner tasks
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.523 🎉🎉🎉
💀    🚀 makeMysqlAppRunner   🐬 15:19:09.523 Done
💀 🎉 Successfully running 🐬 'makeMysqlAppRunner' command
💀 🎉 Reach 🐬 'addMysql' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.812283894s
         Current Time: 15:19:09
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.923304681s
         Current Time: 15:19:09
zaruba please addMysql -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=myDb' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

# Add Book Catalogue API

<!--startCode-->
```bash
cd examples/playground/myEndToEndDemo
zaruba please addFastApiCrud \
  appDirectory=myBackend \
  appModuleName=library \
  appCrudEntity=books \
  appCrudFields='["title","author","synopsis"]' \
  appDependencies='["myDb"]' \
  appEnvs='{"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}'
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.535µs
         Current Time: 15:19:10
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 15:19:10.074 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 15:19:10.076         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 15:19:10.076     
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 15:19:10.076   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 15:19:10.076   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 15:19:10.076   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 15:19:10.076 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run ⚡ 'makeFastApiApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiApp       ⚡ 15:19:10.518 🧰 Prepare
💀    🚀 makeFastApiApp       ⚡ 15:19:10.518 Preparing base variables
💀    🚀 makeFastApiApp       ⚡ 15:19:10.671 Base variables prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.671 Preparing start command
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Start command prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Preparing prepare command
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Prepare command prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Preparing test command
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Test command prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Preparing migrate command
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Migrate command prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Preparing check command
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Check command prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.672 Preparing replacement map
💀    🚀 makeFastApiApp       ⚡ 15:19:10.877 Add config to replacement map
💀    🚀 makeFastApiApp       ⚡ 15:19:10.883 Add env to replacement map
💀    🚀 makeFastApiApp       ⚡ 15:19:10.889 Replacement map prepared
💀    🚀 makeFastApiApp       ⚡ 15:19:10.889 ✅ Validate
💀    🚀 makeFastApiApp       ⚡ 15:19:10.889 Validate app directory
💀    🚀 makeFastApiApp       ⚡ 15:19:10.889 Done validating app directory
💀    🚀 makeFastApiApp       ⚡ 15:19:10.889 Validate app container volumes
💀    🚀 makeFastApiApp       ⚡ 15:19:10.892 Done validating app container volumes
💀    🚀 makeFastApiApp       ⚡ 15:19:10.892 Validate template locations
💀    🚀 makeFastApiApp       ⚡ 15:19:10.901 Done validating template locations
💀    🚀 makeFastApiApp       ⚡ 15:19:10.901 Validate app ports
💀    🚀 makeFastApiApp       ⚡ 15:19:10.904 Done validating app ports
💀    🚀 makeFastApiApp       ⚡ 15:19:10.904 Validate app crud fields
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907 Done validating app crud fields
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907 🚧 Generate
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907 🚧 Template Location: [
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate"
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907 ]
💀    🚀 makeFastApiApp       ⚡ 15:19:10.907 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"MyBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"MyBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyBackendDeployment","ZtplTaskName":"MyBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start myBackend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_HTTP_PORT:\n  default: \"3000\"\n  from: MY_BACKEND_APP_HTTP_PORT\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: MY_BACKEND_APP_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"my-backend","ztpl-app-event-name":"","ztpl-app-image-name":"my-backend","ztpl-app-module-name":"library","ztpl-app-name":"my-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-backend-deployment","ztpl-task-name":"my-backend","ztplAppContainerName":"myBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["myDb"],"ztplAppDirectory":"myBackend","ztplAppEnvs":{"APP_HTTP_PORT":"3000","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"myBackend","ztplAppPorts":[],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"myDb\"]","ztplCfgAppDirectory":"myBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"myBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApi/appTemplate\"\n]","ztplDeploymentDirectory":"myBackendDeployment","ztplDeploymentName":"myBackendDeployment","ztplDeploymentTaskLocation":"../../myBackendDeployment","ztplTaskName":"myBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"my_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"my_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_backend_deployment","ztpl_task_name":"my_backend"}
💀    🚀 makeFastApiApp       ⚡ 15:19:11.475 🔩 Integrate
💀    🚀 makeFastApiApp       ⚡ 15:19:11.476 🎉🎉🎉
💀    🚀 makeFastApiApp       ⚡ 15:19:11.476 Done
💀 🎉 Successfully running ⚡ 'makeFastApiApp' command
💀 🏁 Run ⚡ 'makeFastApiAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeFastApiAppRunner ⚡ 15:19:11.975 🧰 Prepare
💀    🚀 makeFastApiAppRunner ⚡ 15:19:11.975 Preparing base variables
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Base variables prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Preparing start command
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Start command prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Preparing prepare command
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Prepare command prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Preparing test command
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Test command prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Preparing migrate command
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Migrate command prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Preparing check command
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.079 Check command prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.08  Preparing replacement map
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.297 Add config to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.305 Add env to replacement map
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.311 Replacement map prepared
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.311 ✅ Validate
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.311 Validate app directory
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.311 Done validating app directory
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.311 Validate app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.314 Done validating app container volumes
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.315 Validate template locations
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.329 Done validating template locations
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.329 Validate app ports
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.332 Done validating app ports
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.332 Validate app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.335 Done validating app crud fields
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.335 🚧 Generate
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.335 🚧 Template Location: [
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336   "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template",
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336   "/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate"
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336 ]
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336 
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.336 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"MyBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"MyBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyBackendDeployment","ZtplTaskName":"MyBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare command\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"pytest -rP -v --cov=\"$(pwd)\" --cov-report html","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: MY_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: MY_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: MY_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: MY_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: MY_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: MY_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: MY_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: MY_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: MY_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: MY_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: MY_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: MY_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: MY_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: MY_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: MY_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: MY_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: MY_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: MY_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"my-backend","ztpl-app-event-name":"","ztpl-app-image-name":"my-backend","ztpl-app-module-name":"library","ztpl-app-name":"my-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-backend-deployment","ztpl-task-name":"my-backend","ztplAppContainerName":"myBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["myDb"],"ztplAppDirectory":"myBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"⚡","ztplAppModuleName":"library","ztplAppName":"myBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"myDb\"]","ztplCfgAppDirectory":"myBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"⚡","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"","ztplCfgDefaultAppPorts":"[\"{{ .GetEnv \\\"APP_HTTP_PORT\\\" }}\"]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"","ztplCfgDefaultPythonAppPorts":"[\n  \"3000\"\n]\n","ztplCfgDeploymentDirectory":"myBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgPythonStartCommand":"./start.sh","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/native/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/pythonAppRunner/appRunnerTemplate\"\n]\n","ztplDeploymentDirectory":"myBackendDeployment","ztplDeploymentName":"myBackendDeployment","ztplDeploymentTaskLocation":"../../myBackendDeployment","ztplTaskName":"myBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"my_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"my_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_backend_deployment","ztpl_task_name":"my_backend"}
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.38  🔩 Integrate
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.383 Registering app runner tasks
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.386 Checking prepareMyBackend
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.54  Checking prepare
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.543 Adding prepareMyBackend as dependency of prepare
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.696 Checking testMyBackend
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.859 Checking test
💀    🚀 makeFastApiAppRunner ⚡ 15:19:13.863 Adding testMyBackend as dependency of test
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.064 Checking migrateMyBackend
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.256 Checking migrate
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.26  Adding migrateMyBackend as dependency of migrate
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.425 Checking startMyBackend
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.576 Checking start
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.579 Adding startMyBackend as dependency of start
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.728 Checking startMyBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.875 Checking startContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:14.879 Adding startMyBackendContainer as dependency of startContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.03  Checking runMyBackend
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.19  Checking runMyBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.339 Checking stopMyBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.487 Checking stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.49  Adding stopMyBackendContainer as dependency of stopContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.655 Checking removeMyBackendContainer
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.822 Checking removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:15.826 Adding removeMyBackendContainer as dependency of removeContainers
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.016 Checking buildMyBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.223 Checking buildImages
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.228 Adding buildMyBackendImage as dependency of buildImages
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.42  Checking pushMyBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.574 Checking pushImages
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.577 Adding pushMyBackendImage as dependency of pushImages
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.724 Checking pullMyBackendImage
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.87  Done registering app runner tasks
💀 🔥 🚀 makeFastApiAppRunner ⚡ 15:19:16.873 🔥 Command   : zaruba list rangeIndex
💀 🔥 🚀 makeFastApiAppRunner ⚡ 15:19:16.873 🔥 Arguments : ["[myDb]"]
💀 🔥 🚀 makeFastApiAppRunner ⚡ 15:19:16.873 🔥 Stderr    : invalid character 'm' looking for beginning of value
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.874 🎉🎉🎉
💀    🚀 makeFastApiAppRunner ⚡ 15:19:16.874 Done
💀 🎉 Successfully running ⚡ 'makeFastApiAppRunner' command
💀 🎉 Reach ⚡ 'addFastApi' wrapper
💀 🏁 Run ⚡ 'addFastApiModule' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiModule     ⚡ 15:19:17.27  🧰 Prepare
💀    🚀 addFastApiModule     ⚡ 15:19:17.27  Preparing base variables
💀    🚀 addFastApiModule     ⚡ 15:19:18.1   Base variables prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.1   Preparing start command
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Start command prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Preparing prepare command
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Prepare command prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Preparing test command
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Test command prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Preparing migrate command
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Migrate command prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Preparing check command
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Check command prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.101 Preparing replacement map
💀    🚀 addFastApiModule     ⚡ 15:19:18.32  Add config to replacement map
💀    🚀 addFastApiModule     ⚡ 15:19:18.326 Add env to replacement map
💀    🚀 addFastApiModule     ⚡ 15:19:18.333 Replacement map prepared
💀    🚀 addFastApiModule     ⚡ 15:19:18.333 ✅ Validate
💀    🚀 addFastApiModule     ⚡ 15:19:18.333 Validate app directory
💀    🚀 addFastApiModule     ⚡ 15:19:18.333 Done validating app directory
💀    🚀 addFastApiModule     ⚡ 15:19:18.333 Validate app container volumes
💀    🚀 addFastApiModule     ⚡ 15:19:18.336 Done validating app container volumes
💀    🚀 addFastApiModule     ⚡ 15:19:18.336 Validate template locations
💀    🚀 addFastApiModule     ⚡ 15:19:18.344 Done validating template locations
💀    🚀 addFastApiModule     ⚡ 15:19:18.344 Validate app ports
💀    🚀 addFastApiModule     ⚡ 15:19:18.347 Done validating app ports
💀    🚀 addFastApiModule     ⚡ 15:19:18.347 Validate app crud fields
💀    🚀 addFastApiModule     ⚡ 15:19:18.35  Done validating app crud fields
💀    🚀 addFastApiModule     ⚡ 15:19:18.35  🚧 Generate
💀    🚀 addFastApiModule     ⚡ 15:19:18.35  🚧 Template Location: [
💀    🚀 addFastApiModule     ⚡ 15:19:18.35    "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template"
💀    🚀 addFastApiModule     ⚡ 15:19:18.351 ]
💀    🚀 addFastApiModule     ⚡ 15:19:18.351 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"MyBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"MyBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyBackendDeployment","ZtplTaskName":"MyBackend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: MY_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: MY_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: MY_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: MY_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: MY_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: MY_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: MY_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: MY_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: MY_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: MY_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: MY_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: MY_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: MY_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: MY_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: MY_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: MY_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: MY_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: MY_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"my-backend","ztpl-app-event-name":"","ztpl-app-image-name":"my-backend","ztpl-app-module-name":"library","ztpl-app-name":"my-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-backend-deployment","ztpl-task-name":"my-backend","ztplAppContainerName":"myBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppDependencies":["myDb"],"ztplAppDirectory":"myBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"myBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"myDb\"]","ztplCfgAppDirectory":"myBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"myBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiModule/template\"\n]","ztplDeploymentDirectory":"myBackendDeployment","ztplDeploymentName":"myBackendDeployment","ztplDeploymentTaskLocation":"../../myBackendDeployment","ztplTaskName":"myBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"my_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"my_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_backend_deployment","ztpl_task_name":"my_backend"}
💀    🚀 addFastApiModule     ⚡ 15:19:18.366 🔩 Integrate
💀    🚀 addFastApiModule     ⚡ 15:19:18.366 Registering module
💀    🚀 addFastApiModule     ⚡ 15:19:18.389 Done registering module
💀    🚀 addFastApiModule     ⚡ 15:19:18.389 🎉🎉🎉
💀    🚀 addFastApiModule     ⚡ 15:19:18.389 Done
💀 🎉 Successfully running ⚡ 'addFastApiModule' command
💀 🏁 Run ⚡ 'addFastApiCrud' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 addFastApiCrud       ⚡ 15:19:18.719 🧰 Prepare
💀    🚀 addFastApiCrud       ⚡ 15:19:18.719 Preparing base variables
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Base variables prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.474 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 15:19:19.475 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.475 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 15:19:19.475 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.475 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:19.699 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:19.705 Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:19.712 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:19.712 Set app's crud first field
💀    🚀 addFastApiCrud       ⚡ 15:19:19.721 Done setting app's crud first field
💀    🚀 addFastApiCrud       ⚡ 15:19:19.721 Set repo field declaration
💀    🚀 addFastApiCrud       ⚡ 15:19:19.784 Done setting repo field declaration
💀    🚀 addFastApiCrud       ⚡ 15:19:19.784 Set repo field insert
💀    🚀 addFastApiCrud       ⚡ 15:19:19.856 Done setting repo field insert
💀    🚀 addFastApiCrud       ⚡ 15:19:19.856 Set repo field update
💀    🚀 addFastApiCrud       ⚡ 15:19:19.955 Done setting repo field update
💀    🚀 addFastApiCrud       ⚡ 15:19:19.955 Set schema field declaration
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Done setting schema field declaration
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Preparing start command
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Start command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Preparing prepare command
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Prepare command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Preparing test command
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Test command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Preparing migrate command
💀    🚀 addFastApiCrud       ⚡ 15:19:20.012 Migrate command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.013 Preparing check command
💀    🚀 addFastApiCrud       ⚡ 15:19:20.013 Check command prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.013 Preparing replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:20.253 Add config to replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:20.26  Add env to replacement map
💀    🚀 addFastApiCrud       ⚡ 15:19:20.266 Replacement map prepared
💀    🚀 addFastApiCrud       ⚡ 15:19:20.266 ✅ Validate
💀    🚀 addFastApiCrud       ⚡ 15:19:20.266 Validate app directory
💀    🚀 addFastApiCrud       ⚡ 15:19:20.266 Done validating app directory
💀    🚀 addFastApiCrud       ⚡ 15:19:20.266 Validate app container volumes
💀    🚀 addFastApiCrud       ⚡ 15:19:20.269 Done validating app container volumes
💀    🚀 addFastApiCrud       ⚡ 15:19:20.269 Validate template locations
💀    🚀 addFastApiCrud       ⚡ 15:19:20.277 Done validating template locations
💀    🚀 addFastApiCrud       ⚡ 15:19:20.277 Validate app ports
💀    🚀 addFastApiCrud       ⚡ 15:19:20.28  Done validating app ports
💀    🚀 addFastApiCrud       ⚡ 15:19:20.28  Validate app crud fields
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 Done validating app crud fields
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 🚧 Generate
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 🚧 Template Location: [
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283   "/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template"
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 ]
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 
💀    🚀 addFastApiCrud       ⚡ 15:19:20.283 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_BACKEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"Books","ZtplAppCrudEntity":"Book","ZtplAppDirectory":"MyBackend","ZtplAppEventName":"","ZtplAppModuleName":"Library","ZtplAppName":"MyBackend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyBackendDeployment","ZtplTaskName":"MyBackend","[\\t ]*(class[\\t ]*ZtplAppCrudEntityData.*)":"$1\n    title: str\n    author: str\n    synopsis: str","[\\t ]*(db_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"db_book.title = book_data.title\ndb_book.author = book_data.author\ndb_book.synopsis = book_data.synopsis\n$1","[\\t ]*(id[\\t ]*=[\\t ]*new_ztpl_app_crud_entity_id[\\t ]*,[\\t ]*)":"$1\ntitle=book_data.title,\nauthor=book_data.author,\nsynopsis=book_data.synopsis,","[\\t ]*(id[\\t ]*=[\\t ]Column\\(.*)":"$1\ntitle = Column(String(255), index=True)\nauthor = Column(String(255), index=True)\nsynopsis = Column(String(255), index=True)","[\\t ]*(mem_ztpl_app_crud_entity.updated_at[\\t ]*=[\\t ]datetime.datetime.now\\(.*)":"mem_book.title = book_data.title\nmem_book.author = book_data.author\nmem_book.synopsis = book_data.synopsis\n$1","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myBackend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myBackend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myBackend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"./start.sh","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myBackend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"APP_ACCESS_TOKEN_ALGORITHM:\n  default: HS256\n  from: MY_BACKEND_APP_ACCESS_TOKEN_ALGORITHM\nAPP_ACCESS_TOKEN_EXPIRE_MINUTES:\n  default: \"30\"\n  from: MY_BACKEND_APP_ACCESS_TOKEN_EXPIRE_MINUTES\nAPP_ACCESS_TOKEN_SECRET_KEY:\n  default: 09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7\n  from: MY_BACKEND_APP_ACCESS_TOKEN_SECRET_KEY\nAPP_ACCESS_TOKEN_URL:\n  default: /token/\n  from: MY_BACKEND_APP_ACCESS_TOKEN_URL\nAPP_ENABLE_EVENT_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_EVENT_HANDLER\nAPP_ENABLE_ROUTE_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_ROUTE_HANDLER\nAPP_ENABLE_RPC_HANDLER:\n  default: \"1\"\n  from: MY_BACKEND_APP_ENABLE_RPC_HANDLER\nAPP_ERROR_THRESHOLD:\n  default: \"10\"\n  from: MY_BACKEND_APP_ERROR_THRESHOLD\nAPP_GUEST_USERNAME:\n  default: guest\n  from: MY_BACKEND_APP_GUEST_USERNAME\nAPP_HTTP_PORT:\n  default: \"3000\"\n  from: MY_BACKEND_APP_HTTP_PORT\nAPP_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_APP_KAFKA_BOOTSTRAP_SERVERS\nAPP_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_APP_KAFKA_SASL_MECHANISM\nAPP_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_PASSWORD\nAPP_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_APP_KAFKA_SASL_PLAIN_USERNAME\nAPP_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_APP_KAFKA_SCHEMA_REGISTRY\nAPP_MESSAGE_BUS_TYPE:\n  default: local\n  from: MY_BACKEND_APP_MESSAGE_BUS_TYPE\nAPP_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_APP_RABBITMQ_HOST\nAPP_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_APP_RABBITMQ_PASS\nAPP_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_APP_RABBITMQ_USER\nAPP_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_APP_RABBITMQ_VHOST\nAPP_ROOT_INITIAL_EMAIL:\n  default: root@innistrad.com\n  from: MY_BACKEND_APP_ROOT_INITIAL_EMAIL\nAPP_ROOT_INITIAL_FULL_NAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_INITIAL_FULL_NAME\nAPP_ROOT_INITIAL_PASSWORD:\n  default: Alch3mist\n  from: MY_BACKEND_APP_ROOT_INITIAL_PASSWORD\nAPP_ROOT_INITIAL_PHONE_NUMBER:\n  default: \"+621234567890\"\n  from: MY_BACKEND_APP_ROOT_INITIAL_PHONE_NUMBER\nAPP_ROOT_PERMISSION:\n  default: root\n  from: MY_BACKEND_APP_ROOT_PERMISSION\nAPP_ROOT_USERNAME:\n  default: root\n  from: MY_BACKEND_APP_ROOT_USERNAME\nAPP_RPC_TYPE:\n  default: local\n  from: MY_BACKEND_APP_RPC_TYPE\nAPP_SQLALCHEMY_DATABASE_URL:\n  default: mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\n  from: MY_BACKEND_APP_SQLALCHEMY_DATABASE_URL\nAPP_STATIC_DIRECTORY:\n  default: \"\"\n  from: MY_BACKEND_APP_STATIC_DIRECTORY\nAPP_STATIC_URL:\n  default: /static\n  from: MY_BACKEND_APP_STATIC_URL\nTEST_INTEGRATION:\n  default: \"0\"\n  from: MY_BACKEND_TEST_INTEGRATION\nTEST_KAFKA_BOOTSTRAP_SERVERS:\n  default: localhost:9092\n  from: MY_BACKEND_TEST_KAFKA_BOOTSTRAP_SERVERS\nTEST_KAFKA_SASL_MECHANISM:\n  default: PLAIN\n  from: MY_BACKEND_TEST_KAFKA_SASL_MECHANISM\nTEST_KAFKA_SASL_PLAIN_PASSWORD:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_PASSWORD\nTEST_KAFKA_SASL_PLAIN_USERNAME:\n  default: \"\"\n  from: MY_BACKEND_TEST_KAFKA_SASL_PLAIN_USERNAME\nTEST_KAFKA_SCHEMA_REGISTRY:\n  default: http://localhost:8081\n  from: MY_BACKEND_TEST_KAFKA_SCHEMA_REGISTRY\nTEST_RABBITMQ_HOST:\n  default: localhost\n  from: MY_BACKEND_TEST_RABBITMQ_HOST\nTEST_RABBITMQ_PASS:\n  default: Alch3mist\n  from: MY_BACKEND_TEST_RABBITMQ_PASS\nTEST_RABBITMQ_USER:\n  default: root\n  from: MY_BACKEND_TEST_RABBITMQ_USER\nTEST_RABBITMQ_VHOST:\n  default: /\n  from: MY_BACKEND_TEST_RABBITMQ_VHOST\nTEST_SQLALCHEMY_DATABASE_URL:\n  default: sqlite:///test.db\n  from: MY_BACKEND_TEST_SQLALCHEMY_DATABASE_URL","[\\t ]*ztplAppYamlPorts":"{{ .GetEnv \"APP_HTTP_PORT\" }}","ztpl-app-crud-entities":"books","ztpl-app-crud-entity":"book","ztpl-app-directory":"my-backend","ztpl-app-event-name":"","ztpl-app-image-name":"my-backend","ztpl-app-module-name":"library","ztpl-app-name":"my-backend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-backend-deployment","ztpl-task-name":"my-backend","ztplAppContainerName":"myBackend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"books","ztplAppCrudEntity":"book","ztplAppCrudFields":["title","author","synopsis"],"ztplAppCrudFirstField":"title","ztplAppDependencies":["myDb"],"ztplAppDirectory":"myBackend","ztplAppEnvs":{"APP_ACCESS_TOKEN_ALGORITHM":"HS256","APP_ACCESS_TOKEN_EXPIRE_MINUTES":"30","APP_ACCESS_TOKEN_SECRET_KEY":"09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7","APP_ACCESS_TOKEN_URL":"/token/","APP_ENABLE_EVENT_HANDLER":"1","APP_ENABLE_ROUTE_HANDLER":"1","APP_ENABLE_RPC_HANDLER":"1","APP_ERROR_THRESHOLD":"10","APP_GUEST_USERNAME":"guest","APP_HTTP_PORT":"3000","APP_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","APP_KAFKA_SASL_MECHANISM":"PLAIN","APP_KAFKA_SASL_PLAIN_PASSWORD":"","APP_KAFKA_SASL_PLAIN_USERNAME":"","APP_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","APP_MESSAGE_BUS_TYPE":"local","APP_RABBITMQ_HOST":"localhost","APP_RABBITMQ_PASS":"Alch3mist","APP_RABBITMQ_USER":"root","APP_RABBITMQ_VHOST":"/","APP_ROOT_INITIAL_EMAIL":"root@innistrad.com","APP_ROOT_INITIAL_FULL_NAME":"root","APP_ROOT_INITIAL_PASSWORD":"Alch3mist","APP_ROOT_INITIAL_PHONE_NUMBER":"+621234567890","APP_ROOT_PERMISSION":"root","APP_ROOT_USERNAME":"root","APP_RPC_TYPE":"local","APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4","APP_STATIC_DIRECTORY":"","APP_STATIC_URL":"/static","TEST_INTEGRATION":"0","TEST_KAFKA_BOOTSTRAP_SERVERS":"localhost:9092","TEST_KAFKA_SASL_MECHANISM":"PLAIN","TEST_KAFKA_SASL_PLAIN_PASSWORD":"","TEST_KAFKA_SASL_PLAIN_USERNAME":"","TEST_KAFKA_SCHEMA_REGISTRY":"http://localhost:8081","TEST_RABBITMQ_HOST":"localhost","TEST_RABBITMQ_PASS":"Alch3mist","TEST_RABBITMQ_USER":"root","TEST_RABBITMQ_VHOST":"/","TEST_SQLALCHEMY_DATABASE_URL":"sqlite:///test.db"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"library","ztplAppName":"myBackend","ztplAppPorts":["{{ .GetEnv \"APP_HTTP_PORT\" }}"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myBackend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"books","ztplCfgAppCrudFields":"[\"title\",\"author\",\"synopsis\"]","ztplCfgAppDependencies":"[\"myDb\"]","ztplCfgAppDirectory":"myBackend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"APP_HTTP_PORT\": \"3000\", \"APP_SQLALCHEMY_DATABASE_URL\":\"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"library","ztplCfgAppName":"","ztplCfgAppPorts":"[]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoFastApi","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoFastApiDeployment","ztplCfgDeploymentDirectory":"myBackendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/fastApiCrud/template\"\n]\n","ztplDeploymentDirectory":"myBackendDeployment","ztplDeploymentName":"myBackendDeployment","ztplDeploymentTaskLocation":"../../myBackendDeployment","ztplTaskName":"myBackend","ztpl_app_crud_entities":"books","ztpl_app_crud_entity":"book","ztpl_app_directory":"my_backend","ztpl_app_event_name":"","ztpl_app_module_name":"library","ztpl_app_name":"my_backend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_backend_deployment","ztpl_task_name":"my_backend"}
💀    🚀 addFastApiCrud       ⚡ 15:19:20.309 🔩 Integrate
💀    🚀 addFastApiCrud       ⚡ 15:19:20.309 Registering route handler
💀    🚀 addFastApiCrud       ⚡ 15:19:20.344 Done registering route handler
💀    🚀 addFastApiCrud       ⚡ 15:19:20.344 Registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 15:19:20.402 Done registering rpc handler
💀    🚀 addFastApiCrud       ⚡ 15:19:20.403 Registering repo
💀    🚀 addFastApiCrud       ⚡ 15:19:20.464 Done registering repo
💀    🚀 addFastApiCrud       ⚡ 15:19:20.465 🎉🎉🎉
💀    🚀 addFastApiCrud       ⚡ 15:19:20.465 Done
💀 🎉 Successfully running ⚡ 'addFastApiCrud' command
💀 🔎 Job Running...
         Elapsed Time: 10.49902546s
         Current Time: 15:19:20
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 10.700725059s
         Current Time: 15:19:20
zaruba please addFastApiCrud -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=myBackend' -v 'appModuleName=library' -v 'appCrudEntity=books' -v 'appCrudFields=["title","author","synopsis"]' -v 'appDependencies=["myDb"]' -v 'appEnvs={"APP_HTTP_PORT": "3000", "APP_SQLALCHEMY_DATABASE_URL":"mysql+pymysql://root:Alch3mist@localhost/sample?charset=utf8mb4"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->

# Add Static Web Server

<!--startCode-->
```bash
cd examples/playground/myEndToEndDemo
zaruba please addNginx \
  appDirectory=myFrontend \
  appPorts='["80:80"]' \
  appEnvs='{"API_HOST":"localhost:3000"}'
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 3.141µs
         Current Time: 15:19:22
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀 🏁 Run ☕ 'zrbShowAdv' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 zrbIsProject         🔎 15:19:22.553 Current directory is a valid zaruba project
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 Hello Human, 
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 Did 💀 Zaruba help you saving your keystrokes?
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         ,---,---,---,---,---,---,---,---,---,---,---,---,---,-------,
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |1/2| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0 | + | ' | <-    |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |---'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-----|
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         | ->| | Q | W | E | R | T | Y | U | I | O | P | ] | ^ |  💀 |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |-----',--',--',--',--',--',--',--',--',--',--',--',--'|    |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         | Caps | A | S | D | F | G | H | J | K | L | \ | [ | * |    |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |----,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'-,-'---'----|
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |    | < | Z | X | C | V | B | N | M | , | . | - |          |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         |----'-,-',--'--,'---'---'---'---'---'---'-,-'---',--,------|
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         | ctrl |  | alt |                          |altgr |  | ctrl |
💀    🚀 zrbShowAdv           ☕ 15:19:22.556         '------'  '-----'--------------------------'------'  '------'
💀    🚀 zrbShowAdv           ☕ 15:19:22.556     
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 Zaruba is a free and open source project.
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 You can support Zaruba's development by:
💀    🚀 zrbShowAdv           ☕ 15:19:22.556   * Follow Zaruba's twitter account at: 🐤 @zarubastalchmst 
💀    🚀 zrbShowAdv           ☕ 15:19:22.556   * Openning pull request/issue at: https://github.com/state-alchemists/zaruba
💀    🚀 zrbShowAdv           ☕ 15:19:22.556   * Or donating ☕ to: https://paypal.me/gofrendi
💀    🚀 zrbShowAdv           ☕ 15:19:22.556 
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running ☕ 'zrbShowAdv' command
💀 🏁 Run 📗 'makeNginxApp' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxApp         📗 15:19:22.986 🧰 Prepare
💀    🚀 makeNginxApp         📗 15:19:22.987 Preparing base variables
💀    🚀 makeNginxApp         📗 15:19:23.079 Base variables prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing start command
💀    🚀 makeNginxApp         📗 15:19:23.079 Start command prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing prepare command
💀    🚀 makeNginxApp         📗 15:19:23.079 Prepare command prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing test command
💀    🚀 makeNginxApp         📗 15:19:23.079 Test command prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing migrate command
💀    🚀 makeNginxApp         📗 15:19:23.079 Migrate command prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing check command
💀    🚀 makeNginxApp         📗 15:19:23.079 Check command prepared
💀    🚀 makeNginxApp         📗 15:19:23.079 Preparing replacement map
💀    🚀 makeNginxApp         📗 15:19:23.291 Add config to replacement map
💀    🚀 makeNginxApp         📗 15:19:23.297 Add env to replacement map
💀    🚀 makeNginxApp         📗 15:19:23.303 Replacement map prepared
💀    🚀 makeNginxApp         📗 15:19:23.304 ✅ Validate
💀    🚀 makeNginxApp         📗 15:19:23.304 Validate app directory
💀    🚀 makeNginxApp         📗 15:19:23.304 Done validating app directory
💀    🚀 makeNginxApp         📗 15:19:23.304 Validate app container volumes
💀    🚀 makeNginxApp         📗 15:19:23.307 Done validating app container volumes
💀    🚀 makeNginxApp         📗 15:19:23.307 Validate template locations
💀    🚀 makeNginxApp         📗 15:19:23.316 Done validating template locations
💀    🚀 makeNginxApp         📗 15:19:23.316 Validate app ports
💀    🚀 makeNginxApp         📗 15:19:23.319 Done validating app ports
💀    🚀 makeNginxApp         📗 15:19:23.319 Validate app crud fields
💀    🚀 makeNginxApp         📗 15:19:23.322 Done validating app crud fields
💀    🚀 makeNginxApp         📗 15:19:23.322 🚧 Generate
💀    🚀 makeNginxApp         📗 15:19:23.322 🚧 Template Location: [
💀    🚀 makeNginxApp         📗 15:19:23.322   "/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate"
💀    🚀 makeNginxApp         📗 15:19:23.322 ]
💀    🚀 makeNginxApp         📗 15:19:23.322 🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"MyFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"MyFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyFrontendDeployment","ZtplTaskName":"MyFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start myFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: MY_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"my-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"my-frontend","ztpl-app-module-name":"","ztpl-app-name":"my-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-frontend-deployment","ztpl-task-name":"my-frontend","ztplAppContainerName":"myFrontend","ztplAppContainerVolumes":[],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"myFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"🏁","ztplAppModuleName":"","ztplAppName":"myFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[]","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"myFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[]","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[]","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"myFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/nginx/appTemplate\"\n]","ztplDeploymentDirectory":"myFrontendDeployment","ztplDeploymentName":"myFrontendDeployment","ztplDeploymentTaskLocation":"../../myFrontendDeployment","ztplTaskName":"myFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"my_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"my_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_frontend_deployment","ztpl_task_name":"my_frontend"}
💀    🚀 makeNginxApp         📗 15:19:23.341 🔩 Integrate
💀    🚀 makeNginxApp         📗 15:19:23.341 🎉🎉🎉
💀    🚀 makeNginxApp         📗 15:19:23.341 Done
💀 🎉 Successfully running 📗 'makeNginxApp' command
💀 🏁 Run 📗 'makeNginxAppRunner' command on /home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo
💀    🚀 makeNginxAppRunner   📗 15:19:23.72  🧰 Prepare
💀    🚀 makeNginxAppRunner   📗 15:19:23.72  Preparing base variables
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Base variables prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing start command
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Start command prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing prepare command
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Prepare command prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing test command
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Test command prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing migrate command
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Migrate command prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing check command
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Check command prepared
💀    🚀 makeNginxAppRunner   📗 15:19:23.82  Preparing replacement map
💀    🚀 makeNginxAppRunner   📗 15:19:24.033 Add config to replacement map
💀    🚀 makeNginxAppRunner   📗 15:19:24.04  Add env to replacement map
💀    🚀 makeNginxAppRunner   📗 15:19:24.046 Replacement map prepared
💀    🚀 makeNginxAppRunner   📗 15:19:24.046 ✅ Validate
💀    🚀 makeNginxAppRunner   📗 15:19:24.046 Validate app directory
💀    🚀 makeNginxAppRunner   📗 15:19:24.046 Done validating app directory
💀    🚀 makeNginxAppRunner   📗 15:19:24.046 Validate app container volumes
💀    🚀 makeNginxAppRunner   📗 15:19:24.05  Done validating app container volumes
💀    🚀 makeNginxAppRunner   📗 15:19:24.05  Validate template locations
💀    🚀 makeNginxAppRunner   📗 15:19:24.064 Done validating template locations
💀    🚀 makeNginxAppRunner   📗 15:19:24.064 Validate app ports
💀    🚀 makeNginxAppRunner   📗 15:19:24.066 Done validating app ports
💀    🚀 makeNginxAppRunner   📗 15:19:24.066 Validate app crud fields
💀    🚀 makeNginxAppRunner   📗 15:19:24.07  Done validating app crud fields
💀    🚀 makeNginxAppRunner   📗 15:19:24.07  🚧 Generate
💀    🚀 makeNginxAppRunner   📗 15:19:24.07  🚧 Template Location: [
💀    🚀 makeNginxAppRunner   📗 15:19:24.07    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template",
💀    🚀 makeNginxAppRunner   📗 15:19:24.07    "/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template"
💀    🚀 makeNginxAppRunner   📗 15:19:24.07  ]
💀    🚀 makeNginxAppRunner   📗 15:19:24.07  🚧 Replacement Map: {"ZTPL_APP_ENV_PREFIX":"MY_FRONTEND","ZTPL_ENV_PYTHONUNBUFFERED":"1","ZtplAppCrudEntities":"","ZtplAppCrudEntity":"","ZtplAppDirectory":"MyFrontend","ZtplAppEventName":"","ZtplAppModuleName":"","ZtplAppName":"MyFrontend","ZtplAppRpcName":"","ZtplAppUrl":"","ZtplDeploymentDirectory":"","ZtplDeploymentName":"MyFrontendDeployment","ZtplTaskName":"MyFrontend","[\\t ]*ztplAppBuildImageCommand":"","[\\t ]*ztplAppCheckCommand":"echo \"check myFrontend\"","[\\t ]*ztplAppMigrateCommand":"echo \"migrate myFrontend\"","[\\t ]*ztplAppPrepareCommand":"echo \"prepare myFrontend\"","[\\t ]*ztplAppPushImageCommand":"","[\\t ]*ztplAppStartCommand":"echo \"Replace this with command to start myFrontend\" \u0026\u0026 exit 1","[\\t ]*ztplAppStartContainerCommand":"","[\\t ]*ztplAppTestCommand":"echo \"test myFrontend\"","[\\t ]*ztplAppYamlContainerVolumes":"letsencrypt:/etc/letsencrypt\nhtml:/opt/bitnami/nginx/html\nserver_blocks:/opt/bitnami/nginx/conf/server_blocks","[\\t ]*ztplAppYamlEnvs":"API_HOST:\n  default: localhost:3000\n  from: MY_FRONTEND_API_HOST","[\\t ]*ztplAppYamlPorts":"80:80","ztpl-app-crud-entities":"","ztpl-app-crud-entity":"","ztpl-app-directory":"my-frontend","ztpl-app-event-name":"","ztpl-app-image-name":"my-frontend","ztpl-app-module-name":"","ztpl-app-name":"my-frontend","ztpl-app-rpc-name":"","ztpl-app-url":"","ztpl-deployment-directory":"","ztpl-deployment-name":"my-frontend-deployment","ztpl-task-name":"my-frontend","ztplAppContainerName":"myFrontend","ztplAppContainerVolumes":["letsencrypt:/etc/letsencrypt","html:/opt/bitnami/nginx/html","server_blocks:/opt/bitnami/nginx/conf/server_blocks"],"ztplAppCrudEntities":"","ztplAppCrudEntity":"","ztplAppCrudFields":[],"ztplAppDependencies":[],"ztplAppDirectory":"myFrontend","ztplAppEnvs":{"API_HOST":"localhost:3000"},"ztplAppEventName":"","ztplAppHttpMethod":"get","ztplAppIcon":"📗","ztplAppModuleName":"","ztplAppName":"myFrontend","ztplAppPorts":["80:80"],"ztplAppRpcName":"","ztplAppRunnerVersion":"","ztplAppTaskLocation":"../../myFrontend","ztplAppUrl":"","ztplCfgAfterStart":"echo 🎉🎉🎉\necho \"${_BOLD}${_YELLOW}Done${_NORMAL}\"","ztplCfgAppBaseImageName":"","ztplCfgAppBuildImageCommand":"","ztplCfgAppCheckCommand":"","ztplCfgAppContainerName":"","ztplCfgAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgAppCrudEntity":"","ztplCfgAppCrudFields":"[]","ztplCfgAppDependencies":"[]","ztplCfgAppDirectory":"myFrontend","ztplCfgAppEnvPrefix":"","ztplCfgAppEnvs":"{\"API_HOST\":\"localhost:3000\"}","ztplCfgAppEventName":"","ztplCfgAppHttpMethod":"get","ztplCfgAppIcon":"📗","ztplCfgAppImageName":"","ztplCfgAppMigrateCommand":"","ztplCfgAppModuleName":"","ztplCfgAppName":"","ztplCfgAppPorts":"[\"80:80\"]","ztplCfgAppPrepareCommand":"","ztplCfgAppPushImageCommand":"","ztplCfgAppRpcName":"","ztplCfgAppRunnerVersion":"","ztplCfgAppStartCommand":"","ztplCfgAppStartContainerCommand":"","ztplCfgAppTestCommand":"","ztplCfgAppUrl":"","ztplCfgBeforeStart":"","ztplCfgCmd":"bash","ztplCfgCmdArg":"-c","ztplCfgDefaultAppBaseImageName":"","ztplCfgDefaultAppCheckCommand":"","ztplCfgDefaultAppContainerVolumes":"[\n  \"letsencrypt:/etc/letsencrypt\",\n  \"html:/opt/bitnami/nginx/html\",\n  \"server_blocks:/opt/bitnami/nginx/conf/server_blocks\"\n]\n","ztplCfgDefaultAppDirectory":"myEndToEndDemoNginx","ztplCfgDefaultAppPorts":"[\n  \"80\",\n  \"443\"\n]\n","ztplCfgDefaultAppStartCommand":"","ztplCfgDefaultAppStartContainerCommand":"","ztplCfgDefaultDeploymentDirectory":"myEndToEndDemoNginxDeployment","ztplCfgDeploymentDirectory":"myFrontendDeployment","ztplCfgDeploymentName":"","ztplCfgFinish":"","ztplCfgSetup":"","ztplCfgShouldInitConfigMapVariable":"true","ztplCfgShouldInitConfigVariables":"true","ztplCfgShouldInitEnvMapVariable":"true","ztplCfgShouldInitUtil":"true","ztplCfgStart":"","ztplCfgStrictMode":"true","ztplCfgTaskName":"","ztplCfgTemplateLocations":"[\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/_base/template\",\n  \"/home/gofrendi/zaruba/zaruba-tasks/make/appRunner/dockerContainer/template\"\n]","ztplDeploymentDirectory":"myFrontendDeployment","ztplDeploymentName":"myFrontendDeployment","ztplDeploymentTaskLocation":"../../myFrontendDeployment","ztplTaskName":"myFrontend","ztpl_app_crud_entities":"","ztpl_app_crud_entity":"","ztpl_app_directory":"my_frontend","ztpl_app_event_name":"","ztpl_app_module_name":"","ztpl_app_name":"my_frontend","ztpl_app_rpc_name":"","ztpl_app_url":"","ztpl_deployment_directory":"","ztpl_deployment_name":"my_frontend_deployment","ztpl_task_name":"my_frontend"}
💀    🚀 makeNginxAppRunner   📗 15:19:24.095 🔩 Integrate
💀    🚀 makeNginxAppRunner   📗 15:19:24.098 Registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 15:19:24.101 Checking prepareMyFrontend
💀    🚀 makeNginxAppRunner   📗 15:19:24.258 Checking testMyFrontend
💀    🚀 makeNginxAppRunner   📗 15:19:24.418 Checking migrateMyFrontend
💀    🚀 makeNginxAppRunner   📗 15:19:24.579 Checking startMyFrontend
💀    🚀 makeNginxAppRunner   📗 15:19:24.728 Checking start
💀    🚀 makeNginxAppRunner   📗 15:19:24.731 Adding startMyFrontend as dependency of start
💀    🚀 makeNginxAppRunner   📗 15:19:24.888 Checking startMyFrontendContainer
💀    🚀 makeNginxAppRunner   📗 15:19:25.054 Checking startContainers
💀    🚀 makeNginxAppRunner   📗 15:19:25.057 Adding startMyFrontendContainer as dependency of startContainers
💀    🚀 makeNginxAppRunner   📗 15:19:25.227 Checking runMyFrontend
💀    🚀 makeNginxAppRunner   📗 15:19:25.4   Checking runMyFrontendContainer
💀    🚀 makeNginxAppRunner   📗 15:19:25.559 Checking stopMyFrontendContainer
💀    🚀 makeNginxAppRunner   📗 15:19:25.725 Checking stopContainers
💀    🚀 makeNginxAppRunner   📗 15:19:25.729 Adding stopMyFrontendContainer as dependency of stopContainers
💀    🚀 makeNginxAppRunner   📗 15:19:25.913 Checking removeMyFrontendContainer
💀    🚀 makeNginxAppRunner   📗 15:19:26.09  Checking removeContainers
💀    🚀 makeNginxAppRunner   📗 15:19:26.094 Adding removeMyFrontendContainer as dependency of removeContainers
💀    🚀 makeNginxAppRunner   📗 15:19:26.274 Checking buildMyFrontendImage
💀    🚀 makeNginxAppRunner   📗 15:19:26.426 Checking buildImages
💀    🚀 makeNginxAppRunner   📗 15:19:26.429 Adding buildMyFrontendImage as dependency of buildImages
💀    🚀 makeNginxAppRunner   📗 15:19:26.596 Checking pushMyFrontendImage
💀    🚀 makeNginxAppRunner   📗 15:19:26.765 Checking pushImages
💀    🚀 makeNginxAppRunner   📗 15:19:26.769 Adding pushMyFrontendImage as dependency of pushImages
💀    🚀 makeNginxAppRunner   📗 15:19:26.946 Checking pullMyFrontendImage
💀    🚀 makeNginxAppRunner   📗 15:19:27.1   Done registering app runner tasks
💀    🚀 makeNginxAppRunner   📗 15:19:27.103 🎉🎉🎉
💀    🚀 makeNginxAppRunner   📗 15:19:27.103 Done
💀 🎉 Successfully running 📗 'makeNginxAppRunner' command
💀 🎉 Reach 📗 'addNginx' wrapper
💀 🔎 Job Running...
         Elapsed Time: 4.657245386s
         Current Time: 15:19:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.767703098s
         Current Time: 15:19:27
zaruba please addNginx -e '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/.env' -v 'appDirectory=myFrontend' -v 'appPorts=["80:80"]' -v 'appEnvs={"API_HOST":"localhost:3000"}' -v '/home/gofrendi/zaruba/docs/examples/playground/myEndToEndDemo/default.values.yaml'
```````
</details>
<!--endCode-->


# Create Front Page



# Run Project

# Run Project as Containers

# Push Images

# Deploy to Kubernetes