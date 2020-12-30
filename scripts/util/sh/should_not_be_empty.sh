#!/bin/sh

# USAGE
# sh should_not_be_empty.sh <value> <error-message>

if [ -z "${1}" ] ]
then
    echo "${2}" 1>&2
    exit 1
fi