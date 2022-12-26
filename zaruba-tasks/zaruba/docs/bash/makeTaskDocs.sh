rm -Rf ./docs/core-tasks
mkdir -p ./docs/core-tasks

echo '<!--startTocHeader-->' > ./docs/core-tasks/README.md
echo '[🏠](../README.md)' >> ./docs/core-tasks/README.md
echo '# 🥝 Core Tasks' >> ./docs/core-tasks/README.md
echo '<!--endTocHeader-->' >> ./docs/core-tasks/README.md

REPLACEMENT_MAP="$(./zaruba map set "{}" "${ZARUBA_HOME}" '~/.zaruba')"
LINES="$(./zaruba lines read "./core.zaruba.yaml")"
LINE_INDEX=0
MAX_LINE_INDEX=$(($(./zaruba list length "$LINES")-1))

# Create replacement for extends and dependencies
for LINE_INDEX in $(seq 0 "${MAX_LINE_INDEX}")
do
    LINE="$(./zaruba list get "${LINES}" "${LINE_INDEX}")"
    SUBMATCH="$(./zaruba str submatch "'""${LINE}""'" ".*\/task\.(.*)\.yaml.*")"
    if [ "$(./zaruba list length "${SUBMATCH}")" = 2 ]
    then
        echo "Prepare replacement for * ${TASK_NAME} occurrence"
        TASK_NAME="$(./zaruba list get "${SUBMATCH}" 1)"
        KEBAB_TASK_NAME="$(./zaruba str toKebab "${TASK_NAME}")"
        REPLACEMENT_MAP="$(./zaruba map set "${REPLACEMENT_MAP}" '\* `'${TASK_NAME}'`' '* ['${TASK_NAME}']('${KEBAB_TASK_NAME}'.md)')"
    fi
done

export ZARUBA_DECORATION=plain

# Get explanations and write to files
echo '<!--startTocSubtopic-->' >> ./docs/core-tasks/README.md
echo '# Sub-topics' >> ./docs/core-tasks/README.md
TASK_ICON_MAP={}
TASK_KEBAB_NAME_MAP={}
for LINE_INDEX in $(seq 0 "${MAX_LINE_INDEX}")
do
    LINE="$(./zaruba list get "${LINES}" "${LINE_INDEX}")"
    SUBMATCH="$(./zaruba str submatch "'""${LINE}""'" ".*\/task\.(.*)\.yaml.*")"
    if [ "$(./zaruba list length "${SUBMATCH}")" = 2 ]
    then
        TASK_NAME="$(./zaruba list get "${SUBMATCH}" 1)"
        KEBAB_TASK_NAME="$(./zaruba str toKebab "${TASK_NAME}")"
        TASK_ICON="$(./zaruba task getIcon "${TASK_NAME}" "core.zaruba.yaml")"
        echo "Generating documentation for ${TASK_ICON} ${TASK_NAME} at ${KEBAB_TASK_NAME}.md"
        TASK_EXPLANATION=$(./zaruba please "${TASK_NAME}" -x)
        TASK_EXPLANATION=$(./zaruba str replace "${TASK_EXPLANATION}" "${REPLACEMENT_MAP}")
        TASK_EXPLANATION_LINES=$(./zaruba str split "${TASK_EXPLANATION}")
        DOC_LINES="[\"<!--startTocHeader-->\", \"[🏠](../README.md) > [🥝 Core Tasks](README.md)\", \"# ${TASK_ICON} ${TASK_NAME}\", \"<!--endTocHeader-->\"]"
        DOC_LINES=$(./zaruba list merge "${DOC_LINES}" "${TASK_EXPLANATION_LINES}")
        DOC_FILE="./docs/core-tasks/${KEBAB_TASK_NAME}.md"
        ./zaruba lines write "${DOC_LINES}" "${DOC_FILE}"
        echo '* ['${TASK_NAME}'](./'${KEBAB_TASK_NAME}'.md)' >> ./docs/core-tasks/README.md
        TASK_ICON_MAP="$(./zaruba map set "${TASK_ICON_MAP}" "${TASK_NAME}" "${TASK_ICON}")"
        TASK_KEBAB_NAME_MAP="$(./zaruba map set "${TASK_KEBAB_NAME_MAP}" "${TASK_NAME}" "${KEBAB_TASK_NAME}")"
    fi
done
echo '<!--endTocSubtopic-->' >> ./docs/core-tasks/README.md

echo "Updating task TOC"
python ./zaruba-tasks/zaruba/docs/python/update-task-toc.py "${TASK_ICON_MAP}" "${TASK_KEBAB_NAME_MAP}"