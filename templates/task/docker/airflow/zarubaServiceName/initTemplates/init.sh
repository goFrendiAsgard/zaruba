BANNER='
 _       _ _
(_)_ __ (_) |_
| |  _ \| | __|
| | | | | | |_
|_|_| |_|_|\__|
Airflow   Setup
'
echo "${BANNER}"
echo "This script was generated by initiator.gotmpl.sh and will be executed in your airflow container."
echo "Feel free to modify it to accommodate your need."
echo ""
echo "By default, it will load your environment variables and add '_RUNENV_' prefix to each of them."
echo 'Thus, you will be able to load your local $USER: _RUNENV_USER'