_setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}" \
_setReplacementMap "ztplScript" "${_ZRB_SCRIPT}"
_setReplacementMap "ztplSql" "${_ZRB_SQL}"
_setReplacementMap "ztpl-image-name" "${_ZRB_IMAGE_NAME}"
_setReplacementMap "ztplImageTag" "${_ZRB_IMAGE_TAG}"

# add from config and env
_addConfigToReplacementMap
_addEnvToReplacementMap