set -e
echo "Preparing base variables"

_ZRB_TEMPLATE_LOCATIONS="${ZARUBA_CONFIG_TEMPLATE_LOCATIONS}"
_ZRB_APP_CONTAINER_NAME="${ZARUBA_CONFIG_APP_CONTAINER_NAME}"
_ZRB_APP_CONTAINER_VOLUMES="${ZARUBA_CONFIG_APP_CONTAINER_VOLUMES}"
_ZRB_APP_DEPENDENCIES="${ZARUBA_CONFIG_APP_DEPENDENCIES}"
_ZRB_APP_DIRECTORY="${ZARUBA_CONFIG_APP_DIRECTORY}"
_ZRB_APP_ENV_PREFIX="${ZARUBA_CONFIG_APP_ENV_PREFIX}"
_ZRB_APP_ENVS="${ZARUBA_CONFIG_APP_ENVS}"
_ZRB_DEPLOYMENT_DIRECTORY="${ZARUBA_CONFIG_DEPLOYMENT_DIRECTORY}"
_ZRB_DEPLOYMENT_NAME="${ZARUBA_CONFIG_DEPLOYMENT_NAME}"
_ZRB_APP_ICON="${ZARUBA_CONFIG_APP_ICON}"
_ZRB_APP_IMAGE_NAME="${ZARUBA_CONFIG_APP_IMAGE_NAME}"
_ZRB_APP_NAME="${ZARUBA_CONFIG_APP_NAME}"
_ZRB_APP_PORTS="${ZARUBA_CONFIG_APP_PORTS}"
_ZRB_APP_RUNNER_VERSION="${ZARUBA_CONFIG_APP_RUNNER_VERSION}"
_ZRB_APP_CRUD_ENTITY="${ZARUBA_CONFIG_APP_CRUD_ENTITY}"
_ZRB_APP_CRUD_FIELD="${ZARUBA_CONFIG_APP_CRUD_FIELD}"
_ZRB_APP_CRUD_FIELDS="${ZARUBA_CONFIG_APP_CRUD_FIELDS}"
_ZRB_APP_EVENT_NAME="${ZARUBA_CONFIG_APP_EVENT_NAME}"
_ZRB_APP_HTTP_METHOD="${ZARUBA_CONFIG_APP_HTTP_METHOD}"
_ZRB_APP_MODULE_NAME="${ZARUBA_CONFIG_APP_MODULE_NAME}"
_ZRB_APP_RPC_NAME="${ZARUBA_CONFIG_APP_RPC_NAME}"
_ZRB_APP_URL="${ZARUBA_CONFIG_APP_URL}"
_ZRB_TASK_NAME="${ZARUBA_CONFIG_TASK_NAME}"

# app directory
if [ ! -z "${_ZRB_APP_DIRECTORY}" ]
then
    _ZRB_SNAKE_APP_DIRECTORY="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_DIRECTORY}")"
    _ZRB_PASCAL_APP_DIRECTORY="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_DIRECTORY}")"
    _ZRB_KEBAB_APP_DIRECTORY="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_DIRECTORY}")"
fi

# app name
if [ -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_APP_NAME="$("${ZARUBA_BIN}" path getAppName "${_ZRB_APP_DIRECTORY}")"
fi
if [ ! -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_SNAKE_APP_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_NAME}")"
    _ZRB_PASCAL_APP_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_NAME}")"
    _ZRB_KEBAB_APP_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_NAME}")"
fi

# task name
if [ -z "${_ZRB_TASK_NAME}" ]
then
    _ZRB_TASK_NAME="${_ZRB_APP_NAME}"
fi
if [ ! -z "${_ZRB_TASK_NAME}" ]
then
    _ZRB_SNAKE_TASK_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_TASK_NAME}")"
    _ZRB_PASCAL_TASK_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_TASK_NAME}")"
    _ZRB_KEBAB_TASK_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_TASK_NAME}")"
fi

# deployment directory
if [ -z "${_ZRB_DEPLOYMENT_DIRECTORY}" ] && [ ! -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_DEPLOYMENT_DIRECTORY="${_ZRB_APP_NAME}Deployment"
fi
if [ ! -z "${_ZRB_DEPLOYMENT_DIRECTORY}" ]
then
    _ZRB_SNAKE_DEPLOYMENT_DIRECTORY="$("${ZARUBA_BIN}" str toSnake "${_ZRB_DIRECTORY}")"
    _ZRB_PASCAL_DEPLOYMENT_DIRECTORY="$("${ZARUBA_BIN}" str toPascal "${_ZRB_DIRECTORY}")"
    _ZRB_KEBAB_DEPLOYMENT_DIRECTORY="$("${ZARUBA_BIN}" str toKebab "${_ZRB_DIRECTORY}")"
fi

# deployment name
if [ -z "${_ZRB_DEPLOYMENT_NAME}" ]
then
    _ZRB_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" path getAppName "${_ZRB_DEPLOYMENT_DIRECTORY}")"
fi
if [ ! -z "${_ZRB_DEPLOYMENT_NAME}" ]
then
    _ZRB_SNAKE_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_DEPLOYMENT_NAME}")"
    _ZRB_PASCAL_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_DEPLOYMENT_NAME}")"
    _ZRB_KEBAB_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_DEPLOYMENT_NAME}")"
fi

# module name
if [ ! -z "${_ZRB_APP_MODULE_NAME}" ]
then
    _ZRB_UPPER_SNAKE_APP_MODULE_NAME="$("${ZARUBA_BIN}" str toUpperSnake "${_ZRB_APP_MODULE_NAME}")"
    _ZRB_SNAKE_APP_MODULE_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_MODULE_NAME}")"
    _ZRB_PASCAL_APP_MODULE_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_MODULE_NAME}")"
    _ZRB_KEBAB_APP_MODULE_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_MODULE_NAME}")"
fi

# app crud entity
if [ ! -z "${_ZRB_APP_CRUD_ENTITY}" ]
then
    _ZRB_APP_CRUD_ENTITY="$("${ZARUBA_BIN}" str toSingular "${_ZRB_APP_CRUD_ENTITY}")"
    _ZRB_SNAKE_APP_CRUD_ENTITY="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_CRUD_ENTITY}")"
    _ZRB_PASCAL_APP_CRUD_ENTITY="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_CRUD_ENTITY}")"
    _ZRB_KEBAB_APP_CRUD_ENTITY="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_CRUD_ENTITY}")"
fi

# app crud entities
if [ ! -z "${_ZRB_APP_CRUD_ENTITY}" ]
then
    _ZRB_APP_CRUD_ENTITIES="$("${ZARUBA_BIN}" str toPlural "${_ZRB_APP_CRUD_ENTITY}")"
    _ZRB_SNAKE_APP_CRUD_ENTITIES="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_CRUD_ENTITIES}")"
    _ZRB_PASCAL_APP_CRUD_ENTITIES="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_CRUD_ENTITIES}")"
    _ZRB_KEBAB_APP_CRUD_ENTITIES="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_CRUD_ENTITIES}")"
fi

# app crud field
if [ ! -z "${_ZRB_APP_CRUD_FIELD}" ]
then
    _ZRB_SNAKE_APP_CRUD_FIELD="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_CRUD_FIELD}")"
    _ZRB_PASCAL_APP_CRUD_FIELD="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_CRUD_FIELD}")"
    _ZRB_KEBAB_APP_CRUD_FIELD="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_CRUD_FIELD}")"
fi

# app url
if [ ! -z "${_ZRB_APP_URL}" ]
then
    _ZRB_SNAKE_APP_URL="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_URL}")"
    if [ -z "${_ZRB_SNAKE_APP_URL}" ]
    then
        _ZRB_SNAKE_APP_URL="/"
    fi
    _ZRB_PASCAL_APP_URL="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_URL}")"
    if [ -z "${_ZRB_PASCAL_APP_URL}" ]
    then
        _ZRB_PASCAL_APP_URL="/"
    fi
    _ZRB_KEBAB_APP_URL="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_URL}")"
    if [ -z "${_ZRB_KEBAB_APP_URL}" ]
    then
        _ZRB_KEBAB_APP_URL="/"
    fi
fi

# url title
_ZRB_UPPER_SNAKE_APP_URL="$("${ZARUBA_BIN}" str toUpperSnake "${_ZRB_APP_URL}")"
_ZRB_APP_URL_TITLE="$("${ZARUBA_BIN}" str replace "${_ZRB_UPPER_SNAKE_APP_URL}" '{"_": " "}')"
if [ -z "${_ZRB_APP_URL_TITLE}" ]
then
    _ZRB_APP_URL_TITLE="Home"
fi

# app event name
if [ ! -z "${_ZRB_APP_EVENT_NAME}" ]
then
    _ZRB_SNAKE_APP_EVENT_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_EVENT_NAME}")"
    _ZRB_PASCAL_APP_EVENT_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_EVENT_NAME}")"
    _ZRB_KEBAB_APP_EVENT_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_EVENT_NAME}")"
fi

# app rpc name
if [ ! -z "${_ZRB_APP_RPC_NAME}" ]
then
    _ZRB_SNAKE_APP_RPC_NAME="$("${ZARUBA_BIN}" str toSnake "${_ZRB_APP_RPC_NAME}")"
    _ZRB_PASCAL_APP_RPC_NAME="$("${ZARUBA_BIN}" str toPascal "${_ZRB_APP_RPC_NAME}")"
    _ZRB_KEBAB_APP_RPC_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_RPC_NAME}")"
fi

# app task location
if [ ! -z "${_ZRB_APP_DIRECTORY}" ]
then
    _ZRB_APP_TASK_LOCATION="$("${ZARUBA_BIN}" path getRelativePath "./zaruba-tasks/${_ZRB_APP_NAME}" "${_ZRB_APP_DIRECTORY}")"
fi

# deployment task location
if [ ! -z "${_ZRB_DEPLOYMENT_DIRECTORY}" ]
then
    _ZRB_DEPLOYMENT_TASK_LOCATION="$("${ZARUBA_BIN}" path getRelativePath "./zaruba-tasks/${_ZRB_DEPLOYMENT_NAME}" "${_ZRB_DEPLOYMENT_DIRECTORY}")"
fi

# app icon
if [ -z "${_ZRB_APP_ICON}" ]
then
    _ZRB_APP_ICON=🏁
fi

# env prefix
if [ -z "${_ZRB_APP_ENV_PREFIX}" ] && [ ! -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_APP_ENV_PREFIX="$("${ZARUBA_BIN}" str toUpperSnake "${_ZRB_APP_NAME}")"
fi

# image name
if [ -z "${_ZRB_APP_IMAGE_NAME}" ] && [ ! -z "${_ZRB_APP_DIRECTORY}" ]
then
    _ZRB_APP_IMAGE_NAME="${_ZRB_APP_DIRECTORY}"
fi
if [ ! -z "${_ZRB_APP_IMAGE_NAME}" ]
then
    _ZRB_APP_IMAGE_NAME="$("${ZARUBA_BIN}" str toKebab "${_ZRB_APP_IMAGE_NAME}")"
fi

# container name
if [ -z "${_ZRB_APP_CONTAINER_NAME}" ]
then
    _ZRB_APP_CONTAINER_NAME="$("${ZARUBA_BIN}" str toCamel "${_ZRB_APP_NAME}")"
fi

if [ -d "${_ZRB_APP_DIRECTORY}" ]
then
    # envs
    _ZRB_DEFAULT_APP_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}")"
    _ZRB_APP_ENVS="$("${ZARUBA_BIN}" map merge "${_ZRB_APP_ENVS}" "${_ZRB_DEFAULT_APP_ENVS}")"

    # ports
    _ZRB_DEFAULT_APP_PORTS="$("${ZARUBA_BIN}" path getPortConfig "${_ZRB_APP_DIRECTORY}")"
    if [ "$("${ZARUBA_BIN}" list length "${_ZRB_APP_PORTS}")" = 0 ]
    then
        _ZRB_APP_PORTS="${_ZRB_DEFAULT_APP_PORTS}"
    fi
fi

# yaml ports
_ZRB_APP_YAML_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_PORTS}")"

# yaml volumes
_ZRB_APP_YAML_CONTAINER_VOLUMES="$("${ZARUBA_BIN}" list join "${_ZRB_APP_CONTAINER_VOLUMES}")"

# yaml envs
_ZRB_APP_YAML_ENVS="$(_getYamlEnvs "${_ZRB_APP_ENVS}" "${_ZRB_APP_ENV_PREFIX}")"

echo "Base variables prepared"