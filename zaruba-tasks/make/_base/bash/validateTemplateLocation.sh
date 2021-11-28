if [ "$("${ZARUBA_BIN}" list validate "${_ZRB_TEMPLATE_LOCATIONS}")" = 0 ]
then
    echo "${_RED}Invalid _ZRB_TEMPLATE_LOCATIONS: ${_ZRB_TEMPLATE_LOCATIONS}${_NORMAL}"
    exit 1
fi
for _ZRB_TEMPLATE_LOCATION_INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_TEMPLATE_LOCATIONS}")
do
    _ZRB_TEMPLATE_LOCATION="$("${ZARUBA_BIN}" list get "${_ZRB_TEMPLATE_LOCATIONS}" "${_ZRB_TEMPLATE_LOCATION_INDEX}")"
    if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
    then
    echo "${_RED}${_BOLD}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.${_NORMAL}"
    exit 1
    fi
done