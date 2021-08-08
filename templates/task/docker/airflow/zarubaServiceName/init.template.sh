#!/bin/bash

BANNER='
 _       _ _         _
(_)_ __ (_) |_   ___| |__
| | `_ \| | __| / __| `_ \
| | | | | | |_ _\__ \ | | |
|_|_| |_|_|\__(_)___/_| |_|
      Prepare your airflow.
        Achieve your dream.
===========================
'
echo "${BANNER}"

{{ $prefix := .GetConfig "envPrefix" }}
{{ $d := .Decoration }}

echo 'Your envPrefix is: {{ $prefix }}'
echo 'You can change your envPrefix by set up ENV_PREFIX before running the task'
echo ''
echo 'Your USER is: {{ .GetPrefixedEnv $prefix "USER" }}'
echo 'This value was taken from ${{ $prefix }}_USER'
echo 'In case of ${{ $prefix }}_USER does not exist, $USER will be used instead'
echo ''
echo 'Feel free to modify this script.'
echo 'For example, if you need to set airflow variable, you can do:'
echo '{{ $d.Yellow }}{{ $d.Bold }}airflow variables set key VALUE{{ $d.Normal }}'