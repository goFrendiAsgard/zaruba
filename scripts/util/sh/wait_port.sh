# USAGE
# sh wait_port.sh <host> <port>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

until nc -z ${1} ${2}
do
    sleep 1
done