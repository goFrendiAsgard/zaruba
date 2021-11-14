
# InitProject

`File Location`:

    /zaruba-tasks/chore/initProject/task.initProject.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Initiate empty zaruba project.




## Extends

* `zrbRunShellScript`


## Dependencies




## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}

    ```


## Check




## Inputs


## Configs

`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_start`:




`beforeStart`:




`finish`:




`setup`:




`start`:

    {{ $d := .Decoration -}}
    if [ -f "index.zaruba.yaml" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}$(pwd) is a zaruba project.{{ $d.Normal }}"
      exit 1
    fi
    git init
    cp -rT "{{ .ZarubaHome }}/zaruba-tasks/chore/initProject/template/" .
    touch .env
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Project created{{ $d.Normal }}"



`_finish`:




`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`afterStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`includeShellUtil`:

    true


`strictMode`:

    true



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1