#!/bin/sh

# USAGE
# /bin/sh get_current_user.sh

if [ ! -z "$SUDO_USER" ]
then
    echo "$SUDO_USER"
elif [ ! -z "$USER" ]
then
    echo "$USER"
else
    id -u -n
fi