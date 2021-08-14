. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generate_fast_api_module.sh

generate_fast_api_crud_handler() {
    _CRUD_TEMPLATE_LOCATION="${1}"
    _MODULE_TEMPLATE_LOCATION="${2}"
    _SERVICE_TEMPLATE_LOCATION="${3}"
    _TASK_TEMPLATE_LOCATION="${4}"
    _SERVICE_NAME="${5}"
    _MODULE_NAME="${6}"
    _ENTITY_NAME="${7}"
    _FIELD_NAMES="${8}"

    generate_fast_api_module \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"

    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${SERVICE_NAME}")
    _SNAKE_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToSnake "${SERVICE_NAME}")
    _UPPER_SNAKE_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToUpper "${SERVICE_NAME}")
    _PASCAL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToPascal "${SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${MODULE_NAME}")
    _PASCAL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" strToPascal "${MODULE_NAME}")
    _SNAKE_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" strToSnake "${MODULE_NAME}")
    _CAMEL_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${ENTITY_NAME}")
    _PASCAL_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" strToPascal "${ENTITY_NAME}")
    _SNAKE_ENTITY_NAME=$("${ZARUBA_HOME}/zaruba" strToSnake "${ENTITY_NAME}")


    _REPLACMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
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


    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/main.py" )"

    # import repo
    _IMPORT_REPO_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/import_repo.py")"
    _IMPORT_REPO_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_REPO_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" insertLineBeforeIndex "${_MAIN_LINES}" 0 "${_IMPORT_REPO_PARTIAL}")"

    # init repo on main.py
    _INIT_REPO_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/init_repo.py")"
    _INIT_REPO_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_INIT_REPO_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _ENGINE_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" "^\s*engine[\s]*=.*$")"
    _ENGINE_DECLARATION_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_MAIN_LINES}" "${_ENGINE_DECLARATION_PATTERN}")"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_MAIN_LINES}" "${_ENGINE_DECLARATION_LINE_INDEX}" "${_INIT_REPO_PARTIAL}")"

    # event controller call
    _EVENT_CONTROLLER_CALL_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" "^(\s*)${_SNAKE_MODULE_NAME}_event_controller\((.*)\)(.*)$")"
    _EVENT_CONTROLLER_CALL_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_MAIN_LINES}" "${_EVENT_CONTROLLER_CALL_PATTERN}")"
    _EVENT_CONTROLLER_CALL_SUBMATCH="$("${ZARUBA_HOME}/zaruba" getLineSubmatch "${_MAIN_LINES}" "${_EVENT_CONTROLLER_CALL_PATTERN}")"
    _INDENTATION="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_CALL_SUBMATCH}" 1)"
    PARAMETERS="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_CALL_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_CALL_SUBMATCH}" 3)"
    _NEW_EVENT_CONTROLLER_CALL="${_INDENTATION}${_SNAKE_MODULE_NAME}_event_controller(${PARAMETERS}, ${_SNAKE_ENTITY_NAME}_repo)${_SUFFIX}"
    _MAIN_LINES="$("${ZARUBA_HOME}/zaruba" replaceLineAtIndex "${_MAIN_LINES}" "${_EVENT_CONTROLLER_CALL_LINE_INDEX}" "${_NEW_EVENT_CONTROLLER_CALL}")"

    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/main.py" "${_MAIN_LINES}"


    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )"

    # import to controller
    _IMPORT_TO_CONTROLLER_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/import_to_controller.py")"
    _IMPORT_TO_CONTROLLER_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_TO_CONTROLLER_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" insertLineBeforeIndex "${_CONTROLLER_LINES}" 0 "${_IMPORT_TO_CONTROLLER_PARTIAL}")"

    # handle route on controller.py
    _CONTROLLER_HANDLE_ROUTE_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/controller_handle_route.py")"
    _CONTROLLER_HANDLE_ROUTE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_CONTROLLER_HANDLE_ROUTE_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_HANDLE_ROUTE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_CONTROLLER_HANDLE_ROUTE_PARTIAL}" "    " )"
    _ROUTE_CONTROLLER_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" "^\s*def route_controller\(.*\):.*$")"
    _ROUTE_CONTROLLER_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_CONTROLLER_LINES}" "${_ROUTE_CONTROLLER_PATTERN}")"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_CONTROLLER_LINES}" "${_ROUTE_CONTROLLER_LINE_INDEX}" "${_CONTROLLER_HANDLE_ROUTE_PARTIAL}")"

    # handle event on controller.py
    _CONTROLLER_HANDLE_EVENT_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/controller_handle_event.py")"
    _CONTROLLER_HANDLE_EVENT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_CONTROLLER_HANDLE_EVENT_PARTIAL}" "${_REPLACMENT_MAP}" )"
    _CONTROLLER_HANDLE_EVENT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_CONTROLLER_HANDLE_EVENT_PARTIAL}" "    " )"
    _EVENT_CONTROLLER_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" "^(\s*)def event_controller\((.*)\):(.*)$")"
    _EVENT_CONTROLLER_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_CONTROLLER_LINES}" "${_EVENT_CONTROLLER_PATTERN}")"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_CONTROLLER_LINES}" "${_EVENT_CONTROLLER_LINE_INDEX}" "${_CONTROLLER_HANDLE_EVENT_PARTIAL}")"

    _EVENT_CONTROLLER_SUBMATCH="$("${ZARUBA_HOME}/zaruba" getLineSubmatch "${_CONTROLLER_LINES}" "${_EVENT_CONTROLLER_PATTERN}")"
    _INDENTATION="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_SUBMATCH}" 1)"
    PARAMETERS="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_HOME}/zaruba" getFromList "${_EVENT_CONTROLLER_SUBMATCH}" 3)"
    _NEW_EVENT_CONTROLLER="${_INDENTATION}def event_controller(${PARAMETERS}, ${_SNAKE_ENTITY_NAME}_repo: ${_PASCAL_ENTITY_NAME}Repo):${_SUFFIX}"
    _CONTROLLER_LINES="$("${ZARUBA_HOME}/zaruba" replaceLineAtIndex "${_CONTROLLER_LINES}" "${_EVENT_CONTROLLER_LINE_INDEX}" "${_NEW_EVENT_CONTROLLER}")"

    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"

    # field count
    _FIELD_COUNT="$("${ZARUBA_HOME}/zaruba" getListLength "${_FIELD_NAMES}")"
    if [ "${_FIELD_COUNT}" -eq 0 ]
    then
        _FIRST_FIELD_NAME="id"
    else
        _FIRST_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" getFromList "${_FIELD_NAMES}" 0)"
    fi
    _SNAKE_FIRST_FIELD_NAME=$("${ZARUBA_HOME}/zaruba" strToSnake "${_FIRST_FIELD_NAME}")
    _REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zaruba_first_field_name" "${_SNAKE_FIRST_FIELD_NAME}" \
    )"

    # per field
    # schema
    _SCHEMA_LINES="$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/schemas/${_CAMEL_ENTITY_NAME}.py")"
    # repo
    _REPO_LINES="$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/repos/db${_PASCAL_ENTITY_NAME}.py")"
    _REPO_LINES="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_LINES}" "${_REPLACEMENT_MAP}")"

    _MAX_FIELD_INDEX="$((${_FIELD_COUNT}-1))"
    for _FIELD_INDEX in $(seq "${_MAX_FIELD_INDEX}" -1 0)
    do
        _FIELD_NAME="$("${ZARUBA_HOME}/zaruba" getFromList "${_FIELD_NAMES}" "${_FIELD_INDEX}")"
        _SNAKE_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" strToSnake "${_FIELD_NAME}")"

        _REPLACMENT_MAP="$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
            "zaruba_entity_name" "${_SNAKE_ENTITY_NAME}" \
            "zaruba_field_name" "${_SNAKE_FIELD_NAME}" \
        )"

        # schema field declaration
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/schema_field_declaration.py")"
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_SCHEMA_FIELD_DECLARATION_PARTIAL}" "${_REPLACMENT_MAP}")"
        _SCHEMA_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_SCHEMA_FIELD_DECLARATION_PARTIAL}" "    ")"

        _SCHEMA_FIELD_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        "^\s*class\s*${_PASCAL_ENTITY_NAME}Data\s*\(.*$"
        )"
        _SCHEMA_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_SCHEMA_LINES}" "${_SCHEMA_FIELD_DECLARATION_PATTERN}")"

        _SCHEMA_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_SCHEMA_LINES}" "${_SCHEMA_FIELD_LINE_INDEX}" "${_SCHEMA_FIELD_DECLARATION_PARTIAL}")"

        # repo field declaration
        _REPO_FIELD_DECLARATION_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_declaration.py")"
        _REPO_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_DECLARATION_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_DECLARATION_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_REPO_FIELD_DECLARATION_PARTIAL}" "    ")"

        _REPO_FIELD_DECLARATION_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Entity\s*\(.*$" \
        "^\s*__tablename__.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_REPO_LINES}" "${_REPO_FIELD_DECLARATION_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_DECLARATION_PARTIAL}")"

        # repo field insert
        _REPO_FIELD_INSERT_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_insert.py")"
        _REPO_FIELD_INSERT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_INSERT_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_INSERT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_REPO_FIELD_INSERT_PARTIAL}" "$("${ZARUBA_HOME}/zaruba" str repeat "    " 4)")"

        _REPO_FIELD_INSERT_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
        "^\s*def\s*insert\s*\(.*$" \
        "^\s*db_entity\s*=.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_REPO_LINES}" "${_REPO_FIELD_INSERT_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_INSERT_PARTIAL}")"

        # repo field update
        _REPO_FIELD_UPDATE_PARTIAL="$(cat "${_CRUD_TEMPLATE_LOCATION}/partials/repo_field_update.py")"
        _REPO_FIELD_UPDATE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str replace "${_REPO_FIELD_UPDATE_PARTIAL}" "${_REPLACMENT_MAP}")"
        _REPO_FIELD_UPDATE_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_REPO_FIELD_UPDATE_PARTIAL}" "$("${ZARUBA_HOME}/zaruba" str repeat "    " 3)")"

        _REPO_FIELD_UPDATE_PATTERN="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        "^\s*class\s*DB${_PASCAL_ENTITY_NAME}Repo\s*\(.*$" \
        "^\s*def\s*update\s*\(.*$" \
        "^\s*db_entity\.updated_at\s*=.*$" \
        )"
        _REPO_FIELD_LINE_INDEX="$("${ZARUBA_HOME}/zaruba" getLineIndex "${_REPO_LINES}" "${_REPO_FIELD_UPDATE_PATTERN}")"

        _REPO_LINES="$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_REPO_LINES}" "${_REPO_FIELD_LINE_INDEX}" "${_REPO_FIELD_UPDATE_PARTIAL}")"
        
    done

    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/schemas/${_CAMEL_ENTITY_NAME}.py" "${_SCHEMA_LINES}"
    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/repos/db${_PASCAL_ENTITY_NAME}.py" "${_REPO_LINES}"

}