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
* 🥝 Core Tasks
  * zrbShowAdv
  * zrbCopyToKubePod
  * zrbBuildDockerImage
  * zrbCreateDockerNetwork
  * zrbPullDockerImage
  * zrbPushDockerImage
  * zrbRemoveDockerContainer
  * zrbStopDockerContainer
  * zrbGenerateAndRun
  * zrbSetKubeContext
  * zrbRunNodeJsScript
  * zrbRunScript
  * zrbRunPythonScript
  * zrbRunShellScript
  * zrbRunDockerContainer
  * zrbRunInDockerContainer
  * zrbRunInKubePod
  * zrbStartApp
  * zrbStartDockerContainer
  * zrbIsProject
  * zrbIsValidSubrepos
  * zrbWrapper
  * showVersion
  * syncEnv
  * update
  * initProject
  * addProjectLink
  * updateProjectLinks
  * clearLog
  * showLog
  * serveHttp
  * addSubrepo
  * initSubrepos
  * pullSubrepos
  * pushSubrepos
  * setProjectValue
  * clearPreviousValues
  * zrbMake
  * makeApp
  * addAppKubeDeployment
  * makeAppKubeDeployment
  * makeAppKubeDeploymentTask
  * addAppKubeDeploymentInPython
  * makeAppKubeDeploymentInPython
  * makeAppKubeDeploymentTaskInPython
  * zrbMakeTask
  * zrbMakeAppRunner
  * makeDockerAppRunner
  * makeNativeAppRunner
  * addAppRunner
  * makeAppRunner
  * addAirflow
  * makeAirflowApp
  * makeAirflowAppRunner
  * addCassandra
  * makeCassandraApp
  * makeCassandraAppRunner
  * addContainerRegistry
  * makeContainerRegistryApp
  * makeContainerRegistryAppRunner
  * addDocker
  * makeDockerApp
  * makeDockerAppRunnerForDockerApp
  * addEksDeployment
  * makeEksDeployment
  * makeEksDeploymentTask
  * addEksDeploymentInPython
  * makeEksDeploymentInPython
  * makeEksDeploymentTaskInPython
  * addElasticsearch
  * makeElasticsearchApp
  * makeElasticsearchAppRunner
  * addFastApi
  * makeFastApiApp
  * makeFastApiAppRunner
  * addFastApiCrud
  * addFastApiEventHandler
  * addFastApiModule
  * addFastApiRouteHandler
  * addFastApiRpcHandler
  * addGitlab
  * makeGitlabApp
  * makeGitlabAppRunner
  * addGoAppRunner
  * makeGoAppRunner
  * addKafka
  * makeKafkaApp
  * makeKafkaAppRunner
  * addMeltano
  * makeMeltanoApp
  * makeMeltanoAppRunner
  * addMongodb
  * makeMongodbApp
  * makeMongodbAppRunner
  * addMysql
  * makeMysqlApp
  * makeMysqlAppRunner
  * addNginx
  * makeNginxApp
  * makeNginxAppRunner
  * addNodeJsAppRunner
  * makeNodeJsAppRunner
  * addPostgresql
  * makePostgresqlApp
  * makePostgresqlAppRunner
  * addPythonAppRunner
  * makePythonAppRunner
  * addRabbitmq
  * makeRabbitmqApp
  * makeRabbitmqAppRunner
  * addRedis
  * makeRedisApp
  * makeRedisAppRunner
  * addSimpleGoApp
  * makeSimpleGoApp
  * makeSimpleGoAppRunner
  * addSimpleNodeJsApp
  * makeSimpleNodeJsApp
  * makeSimpleNodeJsAppRunner
  * addSimplePythonApp
  * makeSimplePythonApp
  * makeSimplePythonAppRunner
  * addSimpleTypeScriptApp
  * makeSimpleTypeScriptApp
  * makeSimpleTypeScriptAppRunner
  * addTrino
  * makeTrinoApp
  * makeTrinoAppRunner
* 📝 Design Document