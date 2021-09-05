# USAGE: setProjectValueUnlessExist <key> <val>
setProjectValueUnlessExist() {
    _KEY="${1}"
    _VAL="${2}"
    _VALUE_FILE="./default.values.yaml"
    _VALUE_MAP="$("${ZARUBA_HOME}/zaruba" yaml read "${_VALUE_FILE}")"
    if [ -z "$("${ZARUBA_HOME}/zaruba" map get "${_VALUE_MAP}" "${_KEY}")" ]
    then
        "${ZARUBA_HOME}/zaruba" project setValue "${_VALUE_FILE}" "${_KEY}" "${_VAL}"
    fi
}