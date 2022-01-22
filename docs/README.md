[⬅️](../README.md)

# 🏠 Zaruba Documentation

* 🧠 Core Concept
    * 🏗️ Project
        * Project Anatomy
        * Task
            * Task Anatomy
            * Simple Command
            * Long Running Process
            * Task Inputs
            * Task Configs
                * Shared Configs
            * Task Envs
                * Shared Envs
            * Extend task
            * Define task dependencies
        * Project Configs
        * Project Envs
        * Project Inputs
        * Includes
        * Using Go Template
* 👷🏽 Use Cases
    * Creating a Project
    * Add Subrepo
    * Add Runner for Existing Application
        * Go Application Runner
        * NodeJs Application Runner
        * Python Application Runner
    * Generate New Application
        * Simple Go Application
        * Simple NodeJs Application
        * Simple Python Application
        * Simple Typescript Application
        * FastAPI Application
            * Route
            * Event Handler
            * RPC Handler
            * CRUD
    * Add Application Deployment
    * Add Third Party Service
        * Message Bus
            * Rabbitmq
            * Kafka
        * Storage
            * Cassandra
            * Elasticsearch
            * MongoDB
            * MySQL
            * PostgreSQL
            * Redis
        * Data Engineering
            * Airflow
            * Meltano
            * Trino
        * Gitlab
        * Container Registry
        * Nginx
    * Add EKS Deployment
    * Synchronize environments
    * Run Applications Locally
    * Run Some Applications Locally
    * Deploy Applications
* 🔧 Utilities
    * Env
        * Get
        * Print
        * Read
        * Write
    * Lines
        * Fill
        * GetIndex
        * InsertAfter
        * InsertBefore
        * Read
        * Replace
        * Submatch
        * Write
    * List
        * Append
        * Contain
        * Get
        * Join
        * Length
        * Merge
        * RangeIndex
        * Set
        * Validate
    * Map
        * Get
        * GetKeys
        * Merge
        * RangeKey
        * Set
        * ToStringMap
        * ToVariedStringMap
        * TransformKey
        * Validate
    * Num
        * Range
        * ValidateFloat
        * ValidateInt
    * Path
        * GetAppName
        * GetEnv
        * GetPortConfig
        * GetRelativePath
    * Project
        * AddTaskIfNotExist
        * Include
        * SetValue
        * ShowLog
        * SyncEnv
        * SyncEnvFiles
    * Str
        * AddPrefix
        * DoubleQuote
        * FullIndent
        * GetIndentation
        * Indent
        * NewName
        * NewUUID
        * PadLeft
        * PadRight
        * Repeat
        * Replace
        * SingleQuote
        * Split
        * Submatch
        * ToCamel
        * ToKebab
        * ToLower
        * ToPascal
        * ToPlural
        * ToSingular
        * ToSnake
        * ToUpper
        * ToUpperSnake
    * Task
        * AddDependency
        * AddParent
        * IsExist
        * SetConfig
        * SetEnv
        * SyncEnv
    * YAML
        * Print
        * Read
        * Write
* [🥝 Core Tasks](core-tasks/README.md)
  * [zrbShowAdv](core-tasks/zrbShowAdv.md)
  * [zrbCopyToKubePod](core-tasks/zrbCopyToKubePod.md)
  * [zrbBuildDockerImage](core-tasks/zrbBuildDockerImage.md)
  * [zrbCreateDockerNetwork](core-tasks/zrbCreateDockerNetwork.md)
  * [zrbPullDockerImage](core-tasks/zrbPullDockerImage.md)
  * [zrbPushDockerImage](core-tasks/zrbPushDockerImage.md)
  * [zrbRemoveDockerContainer](core-tasks/zrbRemoveDockerContainer.md)
  * [zrbStopDockerContainer](core-tasks/zrbStopDockerContainer.md)
  * [zrbGenerateAndRun](core-tasks/zrbGenerateAndRun.md)
  * [zrbSetKubeContext](core-tasks/zrbSetKubeContext.md)
  * [zrbRunNodeJsScript](core-tasks/zrbRunNodeJsScript.md)
  * [zrbRunScript](core-tasks/zrbRunScript.md)
  * [zrbRunPythonScript](core-tasks/zrbRunPythonScript.md)
  * [zrbRunShellScript](core-tasks/zrbRunShellScript.md)
  * [zrbRunDockerContainer](core-tasks/zrbRunDockerContainer.md)
  * [zrbRunInDockerContainer](core-tasks/zrbRunInDockerContainer.md)
  * [zrbRunInKubePod](core-tasks/zrbRunInKubePod.md)
  * [zrbStartApp](core-tasks/zrbStartApp.md)
  * [zrbStartDockerContainer](core-tasks/zrbStartDockerContainer.md)
  * [zrbIsProject](core-tasks/zrbIsProject.md)
  * [zrbIsValidSubrepos](core-tasks/zrbIsValidSubrepos.md)
  * [zrbWrapper](core-tasks/zrbWrapper.md)
  * [showVersion](core-tasks/showVersion.md)
  * [syncEnv](core-tasks/syncEnv.md)
  * [update](core-tasks/update.md)
  * [initProject](core-tasks/initProject.md)
  * [addProjectLink](core-tasks/addProjectLink.md)
  * [updateProjectLinks](core-tasks/updateProjectLinks.md)
  * [clearLog](core-tasks/clearLog.md)
  * [showLog](core-tasks/showLog.md)
  * [serveHttp](core-tasks/serveHttp.md)
  * [addSubrepo](core-tasks/addSubrepo.md)
  * [initSubrepos](core-tasks/initSubrepos.md)
  * [pullSubrepos](core-tasks/pullSubrepos.md)
  * [pushSubrepos](core-tasks/pushSubrepos.md)
  * [setProjectValue](core-tasks/setProjectValue.md)
  * [clearPreviousValues](core-tasks/clearPreviousValues.md)
  * [zrbMake](core-tasks/zrbMake.md)
  * [makeApp](core-tasks/makeApp.md)
  * [addAppKubeDeployment](core-tasks/addAppKubeDeployment.md)
  * [makeAppKubeDeployment](core-tasks/makeAppKubeDeployment.md)
  * [makeAppKubeDeploymentTask](core-tasks/makeAppKubeDeploymentTask.md)
  * [addAppKubeDeploymentInPython](core-tasks/addAppKubeDeploymentInPython.md)
  * [makeAppKubeDeploymentInPython](core-tasks/makeAppKubeDeploymentInPython.md)
  * [makeAppKubeDeploymentTaskInPython](core-tasks/makeAppKubeDeploymentTaskInPython.md)
  * [zrbMakeTask](core-tasks/zrbMakeTask.md)
  * [zrbMakeAppRunner](core-tasks/zrbMakeAppRunner.md)
  * [makeDockerAppRunner](core-tasks/makeDockerAppRunner.md)
  * [makeNativeAppRunner](core-tasks/makeNativeAppRunner.md)
  * [addAppRunner](core-tasks/addAppRunner.md)
  * [makeAppRunner](core-tasks/makeAppRunner.md)
  * [addAirflow](core-tasks/addAirflow.md)
  * [makeAirflowApp](core-tasks/makeAirflowApp.md)
  * [makeAirflowAppRunner](core-tasks/makeAirflowAppRunner.md)
  * [addCassandra](core-tasks/addCassandra.md)
  * [makeCassandraApp](core-tasks/makeCassandraApp.md)
  * [makeCassandraAppRunner](core-tasks/makeCassandraAppRunner.md)
  * [addContainerRegistry](core-tasks/addContainerRegistry.md)
  * [makeContainerRegistryApp](core-tasks/makeContainerRegistryApp.md)
  * [makeContainerRegistryAppRunner](core-tasks/makeContainerRegistryAppRunner.md)
  * [addDocker](core-tasks/addDocker.md)
  * [makeDockerApp](core-tasks/makeDockerApp.md)
  * [makeDockerAppRunnerForDockerApp](core-tasks/makeDockerAppRunnerForDockerApp.md)
  * [addEksDeployment](core-tasks/addEksDeployment.md)
  * [makeEksDeployment](core-tasks/makeEksDeployment.md)
  * [makeEksDeploymentTask](core-tasks/makeEksDeploymentTask.md)
  * [addEksDeploymentInPython](core-tasks/addEksDeploymentInPython.md)
  * [makeEksDeploymentInPython](core-tasks/makeEksDeploymentInPython.md)
  * [makeEksDeploymentTaskInPython](core-tasks/makeEksDeploymentTaskInPython.md)
  * [addElasticsearch](core-tasks/addElasticsearch.md)
  * [makeElasticsearchApp](core-tasks/makeElasticsearchApp.md)
  * [makeElasticsearchAppRunner](core-tasks/makeElasticsearchAppRunner.md)
  * [addFastApi](core-tasks/addFastApi.md)
  * [makeFastApiApp](core-tasks/makeFastApiApp.md)
  * [makeFastApiAppRunner](core-tasks/makeFastApiAppRunner.md)
  * [addFastApiCrud](core-tasks/addFastApiCrud.md)
  * [addFastApiEventHandler](core-tasks/addFastApiEventHandler.md)
  * [addFastApiModule](core-tasks/addFastApiModule.md)
  * [addFastApiRouteHandler](core-tasks/addFastApiRouteHandler.md)
  * [addFastApiRpcHandler](core-tasks/addFastApiRpcHandler.md)
  * [addGitlab](core-tasks/addGitlab.md)
  * [makeGitlabApp](core-tasks/makeGitlabApp.md)
  * [makeGitlabAppRunner](core-tasks/makeGitlabAppRunner.md)
  * [addGoAppRunner](core-tasks/addGoAppRunner.md)
  * [makeGoAppRunner](core-tasks/makeGoAppRunner.md)
  * [addKafka](core-tasks/addKafka.md)
  * [makeKafkaApp](core-tasks/makeKafkaApp.md)
  * [makeKafkaAppRunner](core-tasks/makeKafkaAppRunner.md)
  * [addMeltano](core-tasks/addMeltano.md)
  * [makeMeltanoApp](core-tasks/makeMeltanoApp.md)
  * [makeMeltanoAppRunner](core-tasks/makeMeltanoAppRunner.md)
  * [addMongodb](core-tasks/addMongodb.md)
  * [makeMongodbApp](core-tasks/makeMongodbApp.md)
  * [makeMongodbAppRunner](core-tasks/makeMongodbAppRunner.md)
  * [addMysql](core-tasks/addMysql.md)
  * [makeMysqlApp](core-tasks/makeMysqlApp.md)
  * [makeMysqlAppRunner](core-tasks/makeMysqlAppRunner.md)
  * [addNginx](core-tasks/addNginx.md)
  * [makeNginxApp](core-tasks/makeNginxApp.md)
  * [makeNginxAppRunner](core-tasks/makeNginxAppRunner.md)
  * [addNodeJsAppRunner](core-tasks/addNodeJsAppRunner.md)
  * [makeNodeJsAppRunner](core-tasks/makeNodeJsAppRunner.md)
  * [addPostgresql](core-tasks/addPostgresql.md)
  * [makePostgresqlApp](core-tasks/makePostgresqlApp.md)
  * [makePostgresqlAppRunner](core-tasks/makePostgresqlAppRunner.md)
  * [addPythonAppRunner](core-tasks/addPythonAppRunner.md)
  * [makePythonAppRunner](core-tasks/makePythonAppRunner.md)
  * [addRabbitmq](core-tasks/addRabbitmq.md)
  * [makeRabbitmqApp](core-tasks/makeRabbitmqApp.md)
  * [makeRabbitmqAppRunner](core-tasks/makeRabbitmqAppRunner.md)
  * [addRedis](core-tasks/addRedis.md)
  * [makeRedisApp](core-tasks/makeRedisApp.md)
  * [makeRedisAppRunner](core-tasks/makeRedisAppRunner.md)
  * [addSimpleGoApp](core-tasks/addSimpleGoApp.md)
  * [makeSimpleGoApp](core-tasks/makeSimpleGoApp.md)
  * [makeSimpleGoAppRunner](core-tasks/makeSimpleGoAppRunner.md)
  * [addSimpleNodeJsApp](core-tasks/addSimpleNodeJsApp.md)
  * [makeSimpleNodeJsApp](core-tasks/makeSimpleNodeJsApp.md)
  * [makeSimpleNodeJsAppRunner](core-tasks/makeSimpleNodeJsAppRunner.md)
  * [addSimplePythonApp](core-tasks/addSimplePythonApp.md)
  * [makeSimplePythonApp](core-tasks/makeSimplePythonApp.md)
  * [makeSimplePythonAppRunner](core-tasks/makeSimplePythonAppRunner.md)
  * [addSimpleTypeScriptApp](core-tasks/addSimpleTypeScriptApp.md)
  * [makeSimpleTypeScriptApp](core-tasks/makeSimpleTypeScriptApp.md)
  * [makeSimpleTypeScriptAppRunner](core-tasks/makeSimpleTypeScriptAppRunner.md)
  * [addTrino](core-tasks/addTrino.md)
  * [makeTrinoApp](core-tasks/makeTrinoApp.md)
  * [makeTrinoAppRunner](core-tasks/makeTrinoAppRunner.md)
* 📝 Design Document