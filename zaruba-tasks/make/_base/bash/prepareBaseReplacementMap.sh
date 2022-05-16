set -e
echo "Preparing replacement map"

_setReplacementMap "[\t ]*ztplAppBuildImageCommand" "${_ZRB_APP_BUILD_IMAGE_COMMAND}"
_setReplacementMap "[\t ]*ztplAppCheckCommand" "${_ZRB_APP_CHECK_COMMAND}"
_setReplacementMap "[\t ]*ztplAppMigrateCommand" "${_ZRB_APP_MIGRATE_COMMAND}"
_setReplacementMap "[\t ]*ztplAppPrepareCommand" "${_ZRB_APP_PREPARE_COMMAND}"
_setReplacementMap "[\t ]*ztplAppPushImageCommand" "${_ZRB_APP_PUSH_IMAGE_COMMAND}"
_setReplacementMap "[\t ]*ztplAppStartCommand" "${_ZRB_APP_START_ICON_COMMAND}"
_setReplacementMap "[\t ]*ztplAppStartContainerCommand" "${_ZRB_APP_START_ICON_CONTAINER_COMMAND}"
_setReplacementMap "[\t ]*ztplAppTestCommand" "${_ZRB_APP_TEST_COMMAND}"
_setReplacementMap "[\t ]*ztplAppYamlContainerVolumes" "${_ZRB_APP_YAML_CONTAINER_VOLUMES}"
_setReplacementMap "[\t ]*ztplAppYamlEnvs" "${_ZRB_APP_YAML_ENVS}"
_setReplacementMap "[\t ]*ztplAppYamlPorts" "${_ZRB_APP_YAML_PORTS}"
_setReplacementMap "ztplAppContainerName" "${_ZRB_APP_CONTAINER_NAME}"
_setReplacementMap "ztplAppContainerVolumes" "${_ZRB_APP_CONTAINER_VOLUMES}"
_setReplacementMap "ztpl_app_crud_entity" "${_ZRB_SNAKE_APP_CRUD_ENTITY}"
_setReplacementMap "ztplAppCrudEntity" "${_ZRB_APP_CRUD_ENTITY}"
_setReplacementMap "ztpl-app-crud-entity" "${_ZRB_KEBAB_APP_CRUD_ENTITY}"
_setReplacementMap "ZtplAppCrudEntity" "${_ZRB_PASCAL_APP_CRUD_ENTITY}"
_setReplacementMap "ztpl_app_crud_entities" "${_ZRB_SNAKE_APP_CRUD_ENTITIES}"
_setReplacementMap "ztplAppCrudEntities" "${_ZRB_APP_CRUD_ENTITIES}"
_setReplacementMap "ztpl-app-crud-entities" "${_ZRB_KEBAB_APP_CRUD_ENTITIES}"
_setReplacementMap "ZtplAppCrudEntities" "${_ZRB_PASCAL_APP_CRUD_ENTITIES}"
_setReplacementMap "ztplAppCrudFields" "${_ZRB_APP_CRUD_FIELDS}"
_setReplacementMap "ztplAppDependencies" "${_ZRB_APP_DEPENDENCIES}"
_setReplacementMap "ztpl_app_directory" "${_ZRB_SNAKE_APP_DIRECTORY}"
_setReplacementMap "ztplAppDirectory" "${_ZRB_APP_DIRECTORY}"
_setReplacementMap "ztpl-app-directory" "${_ZRB_KEBAB_APP_DIRECTORY}"
_setReplacementMap "ZtplAppDirectory" "${_ZRB_PASCAL_APP_DIRECTORY}"
_setReplacementMap "ZTPL_APP_ENV_PREFIX" "${_ZRB_APP_ENV_PREFIX}"
_setReplacementMap "ztplAppEnvs" "${_ZRB_APP_ENVS}"
_setReplacementMap "ztpl_app_event_name" "${_ZRB_SNAKE_APP_EVENT_NAME}"
_setReplacementMap "ztplAppEventName" "${_ZRB_APP_EVENT_NAME}"
_setReplacementMap "ztpl-app-event-name" "${_ZRB_KEBAB_APP_EVENT_NAME}"
_setReplacementMap "ZtplAppEventName" "${_ZRB_PASCAL_APP_EVENT_NAME}"
_setReplacementMap "ztplAppHttpMethod" "${_ZRB_APP_HTTP_METHOD}"
_setReplacementMap "ztplAppIcon" "${_ZRB_APP_ICON}"
_setReplacementMap "ztpl-app-image-name" "${_ZRB_APP_IMAGE_NAME}"
_setReplacementMap "ztpl_app_module_name" "${_ZRB_SNAKE_APP_MODULE_NAME}"
_setReplacementMap "ztplAppModuleName" "${_ZRB_APP_MODULE_NAME}"
_setReplacementMap "ztpl-app-module-name" "${_ZRB_KEBAB_APP_MODULE_NAME}"
_setReplacementMap "ZtplAppModuleName" "${_ZRB_PASCAL_APP_MODULE_NAME}"
_setReplacementMap "ztpl_app_name" "${_ZRB_SNAKE_APP_NAME}"
_setReplacementMap "ztplAppName" "${_ZRB_APP_NAME}"
_setReplacementMap "ztpl-app-name" "${_ZRB_KEBAB_APP_NAME}"
_setReplacementMap "ZtplAppName" "${_ZRB_PASCAL_APP_NAME}"
_setReplacementMap "ztpl_task_name" "${_ZRB_SNAKE_TASK_NAME}"
_setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}"
_setReplacementMap "ztpl-task-name" "${_ZRB_KEBAB_TASK_NAME}"
_setReplacementMap "ZtplTaskName" "${_ZRB_PASCAL_TASK_NAME}"
_setReplacementMap "ztplAppPorts" "${_ZRB_APP_PORTS}"
_setReplacementMap "ztpl_app_rpc_name" "${_ZRB_SNAKE_APP_RPC_NAME}"
_setReplacementMap "ztplAppRpcName" "${_ZRB_APP_RPC_NAME}"
_setReplacementMap "ztpl-app-rpc-name" "${_ZRB_KEBAB_APP_RPC_NAME}"
_setReplacementMap "ZtplAppRpcName" "${_ZRB_PASCAL_APP_RPC_NAME}"
_setReplacementMap "ztplAppRunnerVersion" "${_ZRB_APP_RUNNER_VERSION}"
_setReplacementMap "ztplAppTaskLocation" "${_ZRB_APP_TASK_LOCATION}"
_setReplacementMap "ztpl_app_url" "${_ZRB_SNAKE_APP_URL}"
_setReplacementMap "ztplAppUrl" "${_ZRB_APP_URL}"
_setReplacementMap "ztpl-app-url" "${_ZRB_KEBAB_APP_URL}"
_setReplacementMap "ZtplAppUrl" "${_ZRB_PASCAL_APP_URL}"
_setReplacementMap "ztpl_deployment_directory" "${_ZRB_SNAKE_DEPLOYMENT_DIRECTORY}"
_setReplacementMap "ztplDeploymentDirectory" "${_ZRB_DEPLOYMENT_DIRECTORY}"
_setReplacementMap "ztpl-deployment-directory" "${_ZRB_KEBAB_DEPLOYMENT_DIRECTORY}"
_setReplacementMap "ZtplDeploymentDirectory" "${_ZRB_PASCAL_DEPLOYMENT_DIRECTORY}"
_setReplacementMap "ztpl_deployment_name" "${_ZRB_SNAKE_DEPLOYMENT_NAME}"
_setReplacementMap "ztplDeploymentName" "${_ZRB_DEPLOYMENT_NAME}"
_setReplacementMap "ztpl-deployment-name" "${_ZRB_KEBAB_DEPLOYMENT_NAME}"
_setReplacementMap "ZtplDeploymentName" "${_ZRB_PASCAL_DEPLOYMENT_NAME}"
_setReplacementMap "ztplDeploymentTaskLocation" "${_ZRB_DEPLOYMENT_TASK_LOCATION}"

# add from config and env
echo "Add config to replacement map"
_addConfigToReplacementMap
echo "Add env to replacement map"
_addEnvToReplacementMap

echo "Replacement map prepared"
