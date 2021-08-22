set -e
BANNER='
 _____                _       _
|__  /__ _ _ __ _   _| |__   / \
  / // _` | `__| | | | `_ \ / _ \
 / /| (_| | |  | |_| | |_) / ___ \
/____\__,_|_|   \__,_|_.__/_/   \_\
        3rd Party Package Installer
'

SETUP_PATH="$(dirname $(realpath "${0}"))"
ZARUBA_BIN="${SETUP_PATH}/../zaruba"
SCRIPT_LIST="$("${ZARUBA_BIN}" yaml read "${SETUP_PATH}/initScript.yaml")"
EXIT_CHOICE="$("${ZARUBA_BIN}" list length "${SCRIPT_LIST}")"
PROMPT=""
for INDEX in $("${ZARUBA_BIN}" list rangeIndex "${SCRIPT_LIST}")
do
    CONFIG="$("${ZARUBA_BIN}" list get "${SCRIPT_LIST}" "${INDEX}")"
    PADDED_INDEX="$("${ZARUBA_BIN}" str padLeft "${INDEX}" 3)"
    CAPTION="$("${ZARUBA_BIN}" map get "${CONFIG}" "caption")"
    PROMPT="${PROMPT}${PADDED_INDEX}. ${CAPTION}\n"
done
PADDED_EXIT_CHOICE="$("${ZARUBA_BIN}" str padLeft "${EXIT_CHOICE}" 3)"
PROMPT="${PROMPT}${PADDED_EXIT_CHOICE}. Exit\n"
PROMPT="${PROMPT}\n💀 Type your choice:"

echo "${BANNER}"

until [ "${CHOICE}" = "${EXIT_CHOICE}" ]
do
    echo "${PROMPT}"
    read CHOICE
    if [ "${CHOICE}" != "${EXIT_CHOICE}" ]
    then
        CONFIG="$("${ZARUBA_BIN}" list get "${SCRIPT_LIST}" "${CHOICE}")"
        SCRIPT_PATH="$("${ZARUBA_BIN}" map get "${CONFIG}" "script")"
        SCRIPT_PATH="$("${ZARUBA_BIN}" str replace "${SCRIPT_PATH}" "{\"SETUP_PATH\":\"${SETUP_PATH}\"}")"
        . "${SCRIPT_PATH}"
    fi
done
