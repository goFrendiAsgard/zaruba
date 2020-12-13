#!/bin/sh

# USAGE
# sh wait_port.sh <host> <port>

until nc -z ${1} ${2}
do
    sleep 1
done