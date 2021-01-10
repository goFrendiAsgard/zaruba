# USAGE
# sh is_command_exist.sh <command>
# EXAMPLE
# if [ -z "$(sh "is_command_exist.sh" java)" ]
# then
#   echo "Java is not exist"
# fi

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

command -v "${1}"
if [ "$?" = 0 ]
then
    echo "ok"
fi

set "${_OLD_STATE}"