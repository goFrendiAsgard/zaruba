
# Update

`File Location`:

    /zaruba-tasks/chore/task.update.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Update zaruba to the latest version.




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

`start`:

    {{ $d := .Decoration -}}
    cd "{{ .ZarubaHome }}"
    echo "🔽 {{ $d.Bold }}{{ $d.Yellow }}Pull zaruba{{ $d.Normal }}"
    git checkout master
    git pull origin master
    echo "🚧 {{ $d.Bold }}{{ $d.Yellow }}Compile zaruba{{ $d.Normal }}"
    go build
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Zaruba ready!!!{{ $d.Normal }}"
    echo "{{ $d.Bold }}{{ $d.Yellow }}$(getVersion){{ $d.Normal }}"


`strictMode`:

    true


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`afterStart`:




`_start`:




`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`finish`:




`includeShellUtil`:

    true


`_finish`:




`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`setup`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1