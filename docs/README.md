[⬅️](../README.md)

# 🏠 Zaruba Documentation

* [Core Concept](core-concepts/README.md)
    * Project
        * Project Anatomy
        * [Tasks](core-concepts/tasks.md)
            * Task Anatomy
            * Simple Command
            * Long Running Process
            * Task inputs
            * Task Configs
                * Shared Configs
            * Task Envs
                * Shared Envs
            * Extend task
            * Define task dependencies
        * [Project Configs](core-concepts/configs.md)
        * [Project Envs](core-concepts/envs.md)
        * [Project Inputs](core-concepts/inputs.md)
        * [Includes](includes.md)
        * [Using Go Template](core-concepts/using-go-template.md)
    * Values
    * Environments
* [Use Cases](use-cases/README.md)
    * [Creating a Project](use-cases/creating-a-project.md)
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
        * MessageBus
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
* Utilities
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
* [Core Tasks](core-tasks/README.md)
* [Design document](design-document/README.md)