
# SyncEnv

`File Location`:

    /zaruba-tasks/chore/task.syncEnv.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Update environment of every task in the current project 




## Extends

* `zrbRunShellScript`


## Dependencies

* `zrbIsProject`


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

`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`includeShellUtil`:

    true


`setup`:




`strictMode`:

    true


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`afterStart`:




`beforeStart`:




`cmdArg`:

    -c


`finish`:




`start`:

    {{ $d := .Decoration -}}
    "{{ .ZarubaBin }}" project syncEnv "./index.zaruba.yaml"
    "{{ .ZarubaBin }}" project syncEnvFiles "./index.zaruba.yaml"
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Environment updated{{ $d.Normal }}"



`_finish`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1