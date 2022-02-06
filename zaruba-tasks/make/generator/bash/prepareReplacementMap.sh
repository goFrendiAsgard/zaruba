set -e
echo "Preparing replacement map for generator"

__OLD_IFS="${IFS}"
IFS=$'\n'
for _REPLACEMENT_KEY in $(${ZARUBA_BIN} map rangeKey "${_ZRB_REPLACEMENT_MAP}")
do
    _GEN_REPLACEMENT_KEY="$("${ZARUBA_BIN}" str replace "${_REPLACEMENT_KEY}" '{"ztpl":"gen", "Ztpl":"Gen", "ZTPL": "GEN"}')"
    _setReplacementMap "${_GEN_REPLACEMENT_KEY}" "${_REPLACEMENT_KEY}"
done
IFS="${__OLD_IFS}"