_setReplacementMap() {
    __ZRB_KEY="${1}"
    __ZRB_VAL="${2}"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_KEY}" "${__ZRB_VAL}")"
}