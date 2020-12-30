#!/bin/sh

# USAGE
# sh should_exist.sh <value> <error-message>

if [ ! -e "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi