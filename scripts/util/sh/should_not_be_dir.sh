#!/bin/sh

# USAGE
# sh should_not_be_dir.sh <value> <error-message>

if [ -d "${1}" ]
then
    echo "${2}" 1>&2
    exit 1
fi