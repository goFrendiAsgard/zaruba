
# ShowLog

`File Location`:

    /zaruba-tasks/chore/log/task.showLog.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Show log for all/particular tasks using regex




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


### Inputs.keyword

`Default Value`:




`Description`:

    Keyword


`Prompt`:

    Keyword


`Secret`:

    false


`Validation`:




`Options`:





## Configs

`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`beforeStart`:




`includeShellUtil`:

    true


`strictMode`:

    true


`_finish`:




`afterStart`:




`cmdArg`:

    -c


`finish`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`keyword`:

    {{ if .GetValue "keyword" }}{{ .GetValue "keyword" }}{{ else }}.*{{ end }}


`setup`:




`start`:

    {{ $d := .Decoration -}}
    if [ ! -f "log.zaruba.csv" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}Log is not exist{{ $d.Normal }}"
      exit 1
    fi
    "{{ .ZarubaBin }}" project showLog "{{ .GetWorkPath "log.zaruba.csv" }}" "{{ .GetConfig "keyword"}}"




## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1