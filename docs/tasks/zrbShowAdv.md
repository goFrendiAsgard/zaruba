
# ZrbShowAdv

`File Location`:

    /zaruba-tasks/_base/advertisement/task.zrbShowAdv.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:





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

`_finish`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`finish`:




`strictMode`:

    true


`afterStart`:




`beforeStart`:




`includeShellUtil`:

    true


`setup`:




`start`:

    {{ $showAdvertisement := .GetValue "showAdvertisement" -}}
    {{ if .Util.Bool.IsTrue $showAdvertisement -}}
      "{{ .ZarubaBin }}" advertisement show "{{ printf "%s/advertisement.yaml" .ZarubaHome }}"
    {{ end -}}




## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1