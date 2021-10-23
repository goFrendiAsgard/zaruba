# app directory
_ZRB_APP_SNAKE_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_DIRECTORY}")"
_ZRB_APP_PASCAL_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_DIRECTORY}")"
_ZRB_APP_KEBAB_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_DIRECTORY}")"

# app name
if [ -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_APP_NAME="$("${ZARUBA_HOME}/zaruba" path getAppName "${_ZRB_APP_DIRECTORY}")"
fi
_ZRB_APP_SNAKE_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_NAME}")"
_ZRB_APP_PASCAL_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_NAME}")"
_ZRB_APP_KEBAB_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_NAME}")"

# helm directory
if [ -z "${_ZRB_APP_HELM_DIRECTORY}" ]
then
    _ZRB_APP_HELM_DIRECTORY="${_ZRB_APP_NAME}Helm"
fi
_ZRB_APP_SNAKE_HELM_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_HELM_DIRECTORY}")"
_ZRB_APP_PASCAL_HELM_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_HELM_DIRECTORY}")"
_ZRB_APP_KEBAB_HELM_DIRECTORY="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_HELM_DIRECTORY}")"

# helm release name
if [ -z "${_ZRB_APP_HELM_RELEASE_NAME}" ]
then
    _ZRB_APP_HELM_RELEASE_NAME="${_ZRB_APP_NAME}"
fi
_ZRB_APP_HELM_RELEASE_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_HELM_RELEASE_NAME}")"

# module name
_ZRB_APP_SNAKE_MODULE_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_MODULE_NAME}")"
_ZRB_APP_PASCAL_MODULE_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_MODULE_NAME}")"

# app crud entity
_ZRB_APP_SNAKE_CRUD_ENTITY="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_CRUD_ENTITY}")"
_ZRB_APP_PASCAL_CRUD_ENTITY="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_CRUD_ENTITY}")"

# app url
_ZRB_APP_SNAKE_URL="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_URL}")"
_ZRB_APP_PASCAL_URL="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_URL}")"

# app event name
_ZRB_APP_SNAKE_EVENT_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_EVENT_NAME}")"
_ZRB_APP_PASCAL_EVENT_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_EVENT_NAME}")"

# app rpc name
_ZRB_APP_SNAKE_RPC_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_RPC_NAME}")"
_ZRB_APP_PASCAL_RPC_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_ZRB_APP_RPC_NAME}")"

# app task location
_ZRB_APP_TASK_LOCATION="$("${ZARUBA_HOME}/zaruba" path getRelativePath "./zaruba-tasks/${_ZRB_APP_NAME}" "${_ZRB_APP_DIRECTORY}")"

# app helm task location
_ZRB_APP_HELM_TASK_LOCATION="$("${ZARUBA_HOME}/zaruba" path getRelativePath "./zaruba-tasks/${_ZRB_APP_NAME}Helm" "${_ZRB_APP_HELM_DIRECTORY}")"

# app icon
if [ -z "${_ZRB_APP_ICON}" ]
then
    _ZRB_APP_ICON=🏁
fi

# env prefix
if [ -z "${_ZRB_APP_ENV_PREFIX}" ]
then
    _ZRB_APP_ENV_PREFIX_LOWER="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_NAME}")"
    _ZRB_APP_ENV_PREFIX="$("${ZARUBA_HOME}/zaruba" str toUpper "${_ZRB_APP_ENV_PREFIX_LOWER}")"
fi

# image name
if [ -z "${_ZRB_APP_IMAGE_NAME}" ]
then
    _ZRB_APP_IMAGE_NAME="${_ZRB_APP_DIRECTORY}"
fi
_ZRB_APP_IMAGE_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_IMAGE_NAME}")"


# container name
if [ -z "${_ZRB_APP_CONTAINER_NAME}" ]
then
    _ZRB_APP_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" str toCamel "${_ZRB_APP_NAME}")"
fi

if [ -d "${_ZRB_APP_DIRECTORY}" ]
then
    # envs
    _ZRB_DEFAULT_APP_ENVS="$("${ZARUBA_HOME}/zaruba" path getEnv "${_ZRB_APP_DIRECTORY}")"
    _ZRB_APP_ENVS="$("${ZARUBA_HOME}/zaruba" map merge "${_ZRB_APP_ENVS}" "${_ZRB_DEFAULT_APP_ENVS}")"

    # ports
    _ZRB_DEFAULT_APP_PORTS="$("${ZARUBA_HOME}/zaruba" path getPortConfig "${_ZRB_APP_DIRECTORY}")"
    if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_APP_PORTS}")" = 0 ]
    then
        _ZRB_APP_PORTS="${_ZRB_DEFAULT_APP_PORTS}"
    fi
fi

# yaml ports
_ZRB_APP_YAML_PORTS="$("${ZARUBA_HOME}/zaruba" list join "${_ZRB_APP_PORTS}")"

# yaml volumes
_ZRB_APP_YAML_CONTAINER_VOLUMES="$("${ZARUBA_HOME}/zaruba" list join "${_ZRB_APP_CONTAINER_VOLUMES}")"

# yaml envs
_ZRB_APP_YAML_ENVS="$(_getYamlEnvs "${_ZRB_APP_ENVS}" "${_ZRB_APP_ENV_PREFIX}")"