_setReplacementMap "ztplAppBaseImageName" "${_ZRB_APP_BASE_IMAGE_NAME}" \
    "[\t ]*ztplAppBuildImageCommand" "${_ZRB_APP_BUILD_IMAGE_COMMAND}" \
    "[\t ]*ztplAppCheckCommand" "${_ZRB_APP_CHECK_COMMAND}" \
    "ztplAppContainerName" "${_ZRB_APP_CONTAINER_NAME}" \
    "ztplAppContainerVolumes" "${_ZRB_APP_CONTAINER_VOLUMES}" \
    "ztplAppDependencies" "${_ZRB_APP_DEPENDENCIES}" \
    "ztplAppDirectory" "${_ZRB_APP_DIRECTORY}" \
    "ZtplAppDirectory" "${_ZRB_PASCAL_APP_DIRECTORY}" \
    "ztpl_app_directory" "${_ZRB_SNAKE_APP_DIRECTORY}" \
    "ztpl-app-directory" "${_ZRB_KEBAB_APP_DIRECTORY}" \
    "ztplAppTaskLocation" "${_ZRB_APP_TASK_LOCATION}" \
    "ztplAppEnvPrefix" "${_ZRB_APP_ENV_PREFIX}" \
    "ztplAppEnvs" "${_ZRB_APP_ENVS}" \
    "ztplDeploymentDirectory" "${_ZRB_DEPLOYMENT_DIRECTORY}" \
    "ZtplDeploymentDirectory" "${_ZRB_PASCAL_DEPLOYMENT_DIRECTORY}" \
    "ztpl_deployment_directory" "${_ZRB_SNAKE_DEPLOYMENT_DIRECTORY}" \
    "ztpl-deployment-directory" "${_ZRB_KEBAB_DEPLOYMENT_DIRECTORY}" \
    "ztplDeploymentName" "${_ZRB_DEPLOYMENT_NAME}" \
    "ZtplDeploymentName" "${_ZRB_PASCAL_DEPLOYMENT_NAME}" \
    "ztpl_deployment_name" "${_ZRB_SNAKE_DEPLOYMENT_NAME}" \
    "ztpl-deployment-name" "${_ZRB_KEBAB_DEPLOYMENT_NAME}" \
    "ztplDeploymentTaskLocation" "${_ZRB_DEPLOYMENT_TASK_LOCATION}" \
    "ztplAppIcon" "${_ZRB_APP_ICON}" \
    "ztplAppImageName" "${_ZRB_APP_IMAGE_NAME}" \
    "ztpl-app-image-name" "${_ZRB_APP_IMAGE_NAME}" \
    "ztplAppName" "${_ZRB_APP_NAME}" \
    "ZtplAppName" "${_ZRB_PASCAL_APP_NAME}" \
    "ztpl_app_name" "${_ZRB_SNAKE_APP_NAME}" \
    "ztpl-app-name" "${_ZRB_KEBAB_APP_NAME}" \
    "[\t ]*ztplAppPrepareCommand" "${_ZRB_APP_PREPARE_COMMAND}" \
    "ztplAppPorts" "${_ZRB_APP_PORTS}" \
    "[\t ]*ztplAppPushImageCommand" "${_ZRB_APP_PUSH_IMAGE_COMMAND}" \
    "ztplAppRunnerVersion" "${_ZRB_APP_RUNNER_VERSION}" \
    "[\t ]*ztplAppStartCommand" "${_ZRB_APP_START_COMMAND}" \
    "[\t ]*ztplAppStartContainerCommand" "${_ZRB_APP_START_CONTAINER_COMMAND}" \
    "[\t ]*ztplAppTestCommand" "${_ZRB_APP_TEST_COMMAND}" \
    "[\t ]*ztplAppMigrateCommand" "${_ZRB_APP_MIGRATE_COMMAND}" \
    "ztplAppCrudEntity" "${_ZRB_APP_CRUD_ENTITY}" \
    "ZtplAppCrudEntity" "${_ZRB_PASCAL_APP_CRUD_ENTITY}" \
    "ztpl_app_crud_entity" "${_ZRB_SNAKE_APP_CRUD_ENTITY}" \
    "ztplAppCrudFields" "${_ZRB_APP_CRUD_FIELDS}" \
    "ztplAppEventName" "${_ZRB_APP_EVENT_NAME}" \
    "ZtplAppEventName" "${_ZRB_PASCAL_APP_EVENT_NAME}" \
    "ztpl_app_event_name" "${_ZRB_SNAKE_APP_EVENT_NAME}" \
    "ztplAppHttpMethod" "${_ZRB_APP_HTTP_METHOD}" \
    "ztplAppModuleName" "${_ZRB_APP_MODULE_NAME}" \
    "ZtplAppModuleName" "${_ZRB_PASCAL_APP_MODULE_NAME}" \
    "ztpl_app_module_name" "${_ZRB_SNAKE_APP_MODULE_NAME}" \
    "ztplAppRpcName" "${_ZRB_APP_RPC_NAME}" \
    "ZtplAppRpcName" "${_ZRB_PASCAL_APP_RPC_NAME}" \
    "ztpl_app_rpc_name" "${_ZRB_SNAKE_APP_RPC_NAME}" \
    "ztplAppUrl" "${_ZRB_APP_URL}" \
    "ZtplAppUrl" "${_ZRB_PASCAL_APP_URL}" \
    "ztpl_app_url" "${_ZRB_SNAKE_APP_URL}" \
    "[\t ]*ztplAppYamlPorts" "${_ZRB_APP_YAML_PORTS}" \
    "[\t ]*ztplAppYamlEnvs" "${_ZRB_APP_YAML_ENVS}" \
    "[\t ]*ztplAppYamlContainerVolumes" "${_ZRB_APP_YAML_CONTAINER_VOLUMES}"

# add from config and env
_addConfigToReplacementMap
_addEnvToReplacementMap