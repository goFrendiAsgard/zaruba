#!/bin/sh

# USAGE
# sh should_not_be_file.sh <value> <error-message>

if [ -f "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi