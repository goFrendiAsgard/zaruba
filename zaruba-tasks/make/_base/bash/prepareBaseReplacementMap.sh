set -e
echo "Preparing replacement map"

_ZRB_REPLACEMENT_MAP=$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" \
    "[\t ]*ztplAppBuildImageCommand" "${_ZRB_APP_BUILD_IMAGE_COMMAND}" \
    "[\t ]*ztplAppCheckCommand" "${_ZRB_APP_CHECK_COMMAND}" \
    "[\t ]*ztplAppMigrateCommand" "${_ZRB_APP_MIGRATE_COMMAND}" \
    "[\t ]*ztplAppPrepareCommand" "${_ZRB_APP_PREPARE_COMMAND}" \
    "[\t ]*ztplAppPushImageCommand" "${_ZRB_APP_PUSH_IMAGE_COMMAND}" \
    "[\t ]*ztplAppStartCommand" "${_ZRB_APP_START_COMMAND}" \
    "[\t ]*ztplAppStartContainerCommand" "${_ZRB_APP_START_CONTAINER_COMMAND}" \
    "[\t ]*ztplAppTestCommand" "${_ZRB_APP_TEST_COMMAND}" \
    "[\t ]*ztplAppYamlContainerVolumes" "${_ZRB_APP_YAML_CONTAINER_VOLUMES}" \
    "[\t ]*ztplAppYamlEnvs" "${_ZRB_APP_YAML_ENVS}" \
    "[\t ]*ztplAppYamlPorts" "${_ZRB_APP_YAML_PORTS}" \
    "ztplAppContainerName" "${_ZRB_APP_CONTAINER_NAME}" \
    "ztplAppContainerVolumes" "${_ZRB_APP_CONTAINER_VOLUMES}" \
    "ztpl_app_crud_entity" "${_ZRB_SNAKE_APP_CRUD_ENTITY}" \
    "ztplAppCrudEntity" "${_ZRB_APP_CRUD_ENTITY}" \
    "ztpl-app-crud-entity" "${_ZRB_KEBAB_APP_CRUD_ENTITY}" \
    "ZtplAppCrudEntity" "${_ZRB_PASCAL_APP_CRUD_ENTITY}" \
    "ztpl_app_crud_field" "${_ZRB_SNAKE_APP_CRUD_FIELD}" \
    "ztplAppCrudField" "${_ZRB_APP_CRUD_FIELD}" \
    "ztpl-app-crud-field" "${_ZRB_KEBAB_APP_CRUD_FIELD}" \
    "ZtplAppCrudField" "${_ZRB_PASCAL_APP_CRUD_FIELD}" \
    "ztpl_app_crud_entities" "${_ZRB_SNAKE_APP_CRUD_ENTITIES}" \
    "ztplAppCrudEntities" "${_ZRB_APP_CRUD_ENTITIES}" \
    "ztpl-app-crud-entities" "${_ZRB_KEBAB_APP_CRUD_ENTITIES}" \
    "ZtplAppCrudEntities" "${_ZRB_PASCAL_APP_CRUD_ENTITIES}" \
    "ztplAppCrudFields" "${_ZRB_APP_CRUD_FIELDS}" \
    "ztpl_app_directory" "${_ZRB_SNAKE_APP_DIRECTORY}" \
    "ztplAppDirectory" "${_ZRB_APP_DIRECTORY}" \
    "ztpl-app-directory" "${_ZRB_KEBAB_APP_DIRECTORY}" \
    "ZtplAppDirectory" "${_ZRB_PASCAL_APP_DIRECTORY}" \
    "ZTPL_APP_ENV_PREFIX" "${_ZRB_APP_ENV_PREFIX}" \
    "ztplAppEnvs" "${_ZRB_APP_ENVS}" \
    "ztpl_app_event_name" "${_ZRB_SNAKE_APP_EVENT_NAME}" \
    "ztplAppEventName" "${_ZRB_APP_EVENT_NAME}" \
    "ztpl-app-event-name" "${_ZRB_KEBAB_APP_EVENT_NAME}" \
    "ZtplAppEventName" "${_ZRB_PASCAL_APP_EVENT_NAME}" \
    "ztplAppHttpMethod" "${_ZRB_APP_HTTP_METHOD}" \
    "ztplAppIcon" "${_ZRB_APP_ICON}" \
    "ztpl-app-image-name" "${_ZRB_APP_IMAGE_NAME}" \
    "ZTPL_APP_MODULE_NAME" "${_ZRB_UPPER_SNAKE_APP_MODULE_NAME}" \
    "ztpl_app_module_name" "${_ZRB_SNAKE_APP_MODULE_NAME}" \
    "ztplAppModuleName" "${_ZRB_APP_MODULE_NAME}" \
    "ztpl-app-module-name" "${_ZRB_KEBAB_APP_MODULE_NAME}" \
    "ZtplAppModuleName" "${_ZRB_PASCAL_APP_MODULE_NAME}" \
    "ztpl_app_name" "${_ZRB_SNAKE_APP_NAME}" \
    "ztplAppName" "${_ZRB_APP_NAME}" \
    "ztpl-app-name" "${_ZRB_KEBAB_APP_NAME}" \
    "ZtplAppName" "${_ZRB_PASCAL_APP_NAME}" \
    "ztpl_task_name" "${_ZRB_SNAKE_TASK_NAME}" \
    "ztplTaskName" "${_ZRB_TASK_NAME}" \
    "ztpl-task-name" "${_ZRB_KEBAB_TASK_NAME}" \
    "ZtplTaskName" "${_ZRB_PASCAL_TASK_NAME}" \
    "ztplAppPorts" "${_ZRB_APP_PORTS}" \
    "ztpl_app_rpc_name" "${_ZRB_SNAKE_APP_RPC_NAME}" \
    "ztplAppRpcName" "${_ZRB_APP_RPC_NAME}" \
    "ztpl-app-rpc-name" "${_ZRB_KEBAB_APP_RPC_NAME}" \
    "ZtplAppRpcName" "${_ZRB_PASCAL_APP_RPC_NAME}" \
    "ztplAppRunnerVersion" "${_ZRB_APP_RUNNER_VERSION}" \
    "ztplAppTaskLocation" "${_ZRB_APP_TASK_LOCATION}" \
    "ztpl_app_url" "${_ZRB_SNAKE_APP_URL}" \
    "ztplAppUrl" "${_ZRB_APP_URL}" \
    "ztpl-app-url" "${_ZRB_KEBAB_APP_URL}" \
    "ztpl-normalized-app-url" "${_ZRB_NORMALIZED_APP_URL}" \
    "ZtplAppUrl" "${_ZRB_PASCAL_APP_URL}" \
    "ZtplAppUrlTitle" "${_ZRB_APP_URL_TITLE}" \
    "ztpl_deployment_directory" "${_ZRB_SNAKE_DEPLOYMENT_DIRECTORY}" \
    "ztplDeploymentDirectory" "${_ZRB_DEPLOYMENT_DIRECTORY}" \
    "ztpl-deployment-directory" "${_ZRB_KEBAB_DEPLOYMENT_DIRECTORY}" \
    "ZtplDeploymentDirectory" "${_ZRB_PASCAL_DEPLOYMENT_DIRECTORY}" \
    "ztpl_deployment_name" "${_ZRB_SNAKE_DEPLOYMENT_NAME}" \
    "ztplDeploymentName" "${_ZRB_DEPLOYMENT_NAME}" \
    "ztpl-deployment-name" "${_ZRB_KEBAB_DEPLOYMENT_NAME}" \
    "ZtplDeploymentName" "${_ZRB_PASCAL_DEPLOYMENT_NAME}" \
    "ztplDeploymentTaskLocation" "${_ZRB_DEPLOYMENT_TASK_LOCATION}" \
    "ztplUuid" "${_ZRB_UUID}")

# add from config and env
echo "Add config to replacement map"
_addConfigToReplacementMap
echo "Add env to replacement map"
_addEnvToReplacementMap

echo "Replacement map prepared"
