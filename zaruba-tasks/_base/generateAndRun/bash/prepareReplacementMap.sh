_setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}"
_setReplacementMap "ztplScript" "${_ZRB_SCRIPT}"
_setReplacementMap "ztplSql" "${_ZRB_SQL}"
_setReplacementMap "ztpl-image-name" "${_ZRB_IMAGE_NAME}"
_setReplacementMap "ztplImageTag" "${_ZRB_IMAGE_TAG}"

for _ZRB_KEY in $("${ZARUBA_BIN}" map rangeKey "${_ZRB_ENVS}")
do
    _ZRB_VAL="$("${ZARUBA_BIN}" map get "${_ZRB_ENVS}" "${_ZRB_KEY}")"
    _setReplacementMap "\$${_ZRB_KEY}" "${_ZRB_VAL}"
    _setReplacementMap "\${${_ZRB_KEY}}" "${_ZRB_VAL}"
    _setReplacementMap "ztplEnv_${_ZRB_KEY}" "${_ZRB_VAL}"
done