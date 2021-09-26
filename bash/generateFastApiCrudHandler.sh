. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generateFastApiModule.sh

generateFastApiCrudHandler() {
    _CRUD_TEMPLATE_LOCATION="${1}"
    _MODULE_TEMPLATE_LOCATION="${2}"
    _SERVICE_TEMPLATE_LOCATION="${3}"
    _TASK_TEMPLATE_LOCATION="${4}"
    _SERVICE_NAME="${5}"
    _MODULE_NAME="${6}"
    _ENTITY_NAME="${7}"
    _FIELD_NAMES="${8}"

    generateFastApiModule \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"

    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${SERVICE_NAME}")
    _SNAKE_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toSnake "${SERVICE_NAME}")
    _UPPER_SNAKE_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toUpper "${SERVICE_NAME}")
    _PASCAL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toPascal "${SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${MODULE_NAME}")
    _PASCAL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str toPascal "${MODULE_NAME}")
    _SNAKE_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str toSnake "${MODULE_NAME}")
    _CAMEL_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${ENTITY_NAME}")
    _PASCAL_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" str toPascal "${ENTITY_NAME}")
    _SNAKE_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" str toSnake "${ENTITY_NAME}")


    _REPLACMENT_MAP=$("${ZARUBA_HOME}/zaruba" map set "{}" \
        "ZARUBA_SERVICE_NAME" "${_UPPER_SNAKE_SERVICE_NAME}" \
        "zarubaServiceName" "${_CAMEL_SERVICE_NAME}" \
        "ZarubaServiceName" "${_PASCAL_SERVICE_NAME}" \
        "zarubaModuleName" "${_CAMEL_MODULE_NAME}" \
        "ZarubaModuleName" "${_PASCAL_MODULE_NAME}" \
        "zaruba_module_name" "${_SNAKE_MODULE_NAME}" \
        "zarubaEntityName" "${_CAMEL_ENTITY_NAME}" \
        "ZarubaEntityName" "${_PASCAL_ENTITY_NAME}" \
        "zaruba_entity_name" "${_SNAKE_ENTITY_NAME}" \
    )
    "${ZARUBA_HOME}/zaruba" generate "${_CRUD_TEMPLATE_LOCATION}/zarubaServiceName" "${_CAMEL_SERVICE_NAME}" "${_REPLACMENT_MAP}"


    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/main.py" )"

    # import repo
    _IMPORT_REPO_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/import_repo.py")"
    _IMPORT_REPO_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_REPO_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" lines insertBefore "${_MAIN_LINES}" 0 "${_IMPORT_REPO_PARTIAL}")"

    # init repo on main.py
    _INIT_REPO_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/init_repo.py")"
    _INIT_REPO_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_INIT_REPO_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _ENGINE_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" "^\s*engine[\s]*=.*$")"
    _ENGINE_DECLARATION_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_MAIN_LINES}" "${_ENGINE_DECLARATION_PATTERN}")"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_MAIN_LINES}" "${_ENGINE_DECLARATION_LINE_INDEX}" "${_INIT_REPO_PARTIAL}")"

    # rpc controller call
    _RPC_CONTROLLER_CALL_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" "^(\s*)${_SNAKE_MODULE_NAME}_rpc_controller\((.*)\)(.*)$")"
    _RPC_CONTROLLER_CALL_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_MAIN_LINES}" "${_RPC_CONTROLLER_CALL_PATTERN}")"
    _RPC_CONTROLLER_CALL_SUBMATCH="$("${ZARUBA_HOME}/zaruba" lines submatch "${_MAIN_LINES}" "${_RPC_CONTROLLER_CALL_PATTERN}")"
    _INDENTATION="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_CALL_SUBMATCH}" 1)"
    PARAMETERS="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_CALL_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_CALL_SUBMATCH}" 3)"
    _NEW_RPC_CONTROLLER_CALL="${_INDENTATION}${_SNAKE_MODULE_NAME}_rpc_controller(${PARAMETERS}, ${_SNAKE_ENTITY_NAME}_repo)${_SUFFIX}"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" lines replace "${_MAIN_LINES}" "${_RPC_CONTROLLER_CALL_LINE_INDEX}" "${_NEW_RPC_CONTROLLER_CALL}")"

    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/main.py" "${_MAIN_LINES}"


    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )"

    # import to controller
    _IMPORT_TO_CONTROLLER_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/import_to_controller.py")"
    _IMPORT_TO_CONTROLLER_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_TO_CONTROLLER_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" lines insertBefore "${_CONTROLLER_LINES}" 0 "${_IMPORT_TO_CONTROLLER_PARTIAL}")"

    # handle route on controller.py
    _CONTROLLER_HANDLE_HTTP_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/controller_handle_http.py")"
    _CONTROLLER_HANDLE_HTTP_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_CONTROLLER_HANDLE_HTTP_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_HANDLE_HTTP_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_CONTROLLER_HANDLE_HTTP_PARTIAL}" "    " )"
    _HTTP_CONTROLLER_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" "^\s*def http_controller\(.*\):.*$")"
    _HTTP_CONTROLLER_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_CONTROLLER_LINES}" "${_HTTP_CONTROLLER_PATTERN}")"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_CONTROLLER_LINES}" "${_HTTP_CONTROLLER_LINE_INDEX}" "${_CONTROLLER_HANDLE_HTTP_PARTIAL}")"

    # handle rpc on controller.py
    _CONTROLLER_HANDLE_RPC_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/controller_handle_rpc.py")"
    _CONTROLLER_HANDLE_RPC_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_CONTROLLER_HANDLE_RPC_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_HANDLE_RPC_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_CONTROLLER_HANDLE_RPC_PARTIAL}" "    " )"
    _RPC_CONTROLLER_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" "^(\s*)def rpc_controller\((.*)\):(.*)$")"
    _RPC_CONTROLLER_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_CONTROLLER_LINES}" "${_RPC_CONTROLLER_PATTERN}")"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_CONTROLLER_LINES}" "${_RPC_CONTROLLER_LINE_INDEX}" "${_CONTROLLER_HANDLE_RPC_PARTIAL}")"

    _RPC_CONTROLLER_SUBMATCH="$("${ZARUBA_HOME}/zaruba" lines submatch "${_CONTROLLER_LINES}" "${_RPC_CONTROLLER_PATTERN}")"
    _INDENTATION="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_SUBMATCH}" 1)"
    PARAMETERS="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_HOME}/zaruba" list get "${_RPC_CONTROLLER_SUBMATCH}" 3)"
    _NEW_RPC_CONTROLLER="${_INDENTATION}def rpc_controller(${PARAMETERS}, ${_SNAKE_ENTITY_NAME}_repo: ${_PASCAL_ENTITY_NAME}Repo):${_SUFFIX}"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" lines replace "${_CONTROLLER_LINES}" "${_RPC_CONTROLLER_LINE_INDEX}" "${_NEW_RPC_CONTROLLER}")"

    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"

    # field count
    _FIELD_COUNT="$("${ZARUBA_HOME}/zaruba" list length "${_FIELD_NAMES}")"
    if [ "${_FIELD_COUNT}" -eq 0 ]
    then
        _FIRST_FIELD_NAME="id"
    else
        _FIRST_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" list get "${_FIELD_NAMES}" 0)"
    fi
    _SNAKE_FIRST_FIELD_NAME=$("${ZARUBA_HOME}/zaruba" str toSnake "${_FIRST_FIELD_NAME}")
    _REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "{}" \
        "zaruba_first_field_name" "${_SNAKE_FIRST_FIELD_NAME}" \
    )"

    # per field
    # schema
    _SCHEMA_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/schemas/${_CAMEL_ENTITY_NAME}.py")"
    # repo
    _REPO_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/repos/db${_PASCAL_ENTITY_NAME}.py")"
    _REPO_LINES="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_LINES}" "${_REPLACEMENT_MAP}")"

    _MAX_FIELD_INDEX="$((${_FIELD_COUNT}-1))"
    for _FIELD_INDEX in $(seq "${_MAX_FIELD_INDEX}" -1 0)
    do
        _FIELD_NAME="$("${ZARUBA_HOME}/zaruba" list get "${_FIELD_NAMES}" "${_FIELD_INDEX}")"
        _SNAKE_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_FIELD_NAME}")"

        _REPLACMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "{}" \
            "zaruba_entity_name" "${_SNAKE_ENTITY_NAME}" \
            "zaruba_field_name" "${_SNAKE_FIELD_NAME}" \
        )"

        # schema field declaration
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/schema_field_declaration.py")"
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_SCHEMA_FIELD_DECLARATION_PARTIAL}" "${_REPLACMENT_MAP}")"
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_SCHEMA_FIELD_DECLARATION_PARTIAL}" "    ")"

        _SCHEMA_FIELD_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        "^\s*class\s*${_PASCAL_ENTITY_NAME}Data\s*\(.*$"
        )"
        _SCHEMA_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_SCHEMA_LINES}" "${_SCHEMA_FIELD_DECLARATION_PATTERN}")"

        _SCHEMA_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_SCHEMA_LINES}" "${_SCHEMA_FIELD_LINE_INDEX}" "${_SCHEMA_FIELD_DECLARATION_PARTIAL}")"

        # repo field declaration
        _REPO_FIELD_DECLARATION_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_declaration.py")"
        _REPO_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_DECLARATION_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_REPO_FIELD_DECLARATION_PARTIAL}" "    ")"

        _REPO_FIELD_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Entity\s*\(.*$" \
        "^\s*__tablename__.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_REPO_LINES}" "${_REPO_FIELD_DECLARATION_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_DECLARATION_PARTIAL}")"

        # repo field insert
        _REPO_FIELD_INSERT_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_insert.py")"
        _REPO_FIELD_INSERT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_INSERT_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_INSERT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_REPO_FIELD_INSERT_PARTIAL}" "$("${ZARUBA_HOME}/zaruba" str repeat "    " 4)")"

        _REPO_FIELD_INSERT_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
        "^\s*def\s*insert\s*\(.*$" \
        "^\s*db_entity\s*=.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_REPO_LINES}" "${_REPO_FIELD_INSERT_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_INSERT_PARTIAL}")"

        # repo field update
        _REPO_FIELD_UPDATE_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_update.py")"
        _REPO_FIELD_UPDATE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_UPDATE_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_UPDATE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_REPO_FIELD_UPDATE_PARTIAL}" "$("${ZARUBA_HOME}/zaruba" str repeat "    " 3)")"

        _REPO_FIELD_UPDATE_PATTERN="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
        "^\s*def\s*update\s*\(.*$" \
        "^\s*db_entity\.updated_at\s*=.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_REPO_LINES}" "${_REPO_FIELD_UPDATE_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_UPDATE_PARTIAL}")"
        
    done

    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/schemas/${_CAMEL_ENTITY_NAME}.py" "${_SCHEMA_LINES}"
    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/repos/db${_PASCAL_ENTITY_NAME}.py" "${_REPO_LINES}"

}