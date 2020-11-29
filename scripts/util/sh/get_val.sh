#!/bin/sh

# USAGE
# /bin/sh get_val.sh <val> <default-val>

if [ -z ${1} ]
then
    echo ${2}
else
    echo ${1}
fi