_setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}" \
    "ztplScript" "${_ZRB_SCRIPT}" \
    "ztplSql" "${_ZRB_SQL}" \
    "ztpl-image-name" "${_ZRB_IMAGE_NAME}" \
    "ztplImageTag" "${_ZRB_IMAGE_TAG}"

# add from config and env
_addConfigToReplacementMap
_addEnvToReplacementMap