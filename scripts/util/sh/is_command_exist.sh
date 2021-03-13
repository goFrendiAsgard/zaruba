# USAGE
# sh is_command_exist.sh <command>
# EXAMPLE
# if [ "$(sh "is_command_exist.sh" java)" = 0 ]
# then
#   echo "Java is not exist"
# fi

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

$@ >> /dev/null
if [ "$?" = 127 ]
then
    echo 0
else
    echo 1
fi

set "${_OLD_STATE}"