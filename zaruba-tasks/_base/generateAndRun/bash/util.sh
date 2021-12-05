_setReplacementMap() {
    __ZRB_KEY="${1}"
    __ZRB_VAL="${2}"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "${1}" "${2}")"
}

_addConfigToReplacementMap() {
    # add config with prefix: 'ztplCfg'
    __ZRB_CONFIG_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_CONFIG_MAP}" -t pascal -p ztplCfg)"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_CONFIG_REPLACEMENT_MAP}")"
}

_addEnvToReplacementMap() {
    # add env with prefix: 'ZTPL_ENV_'
    __ZRB_ENV_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_ENV_MAP}" -p ZTPL_ENV_)"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_ENV_REPLACEMENT_MAP}")"
    # add env with prefix: '$'
    __ZRB_ENV_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_ENV_MAP}" -p '\$')"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_ENV_REPLACEMENT_MAP}")"
    # add env with prefix: '${' and suffix '}'
    __ZRB_ENV_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_ENV_MAP}" -p '\${' -s '}')"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_ENV_REPLACEMENT_MAP}")"
}
