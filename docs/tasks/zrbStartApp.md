
# ZrbStartApp

`File Location`:

    /zaruba-tasks/_base/start/task.zrbStartApp.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    service


`Description`:

    Start service and check it's readiness.
    Common configs:
      setup       : Script to be executed before start service or check service readiness.
      start       : Script to start the service (e.g: python -m http.server 9000).
      beforeStart : Script to be executed before start service.
      afterStart  : Script to be executed after start service.
      beforeCheck : Script to be executed before check service readiness.
      afterCheck  : Script to be executed before check service readiness.
      finish      : Script to be executed after start service or check service readiness.
      runInLocal  : Run service locally or not.
      ports       : Port to be checked to confirm service readiness, separated by new line.




## Extends

* `zrbRunShellScript`


## Dependencies

* `updateProjectLinks`


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{- $d := .Decoration -}}
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
      sleep infinity
    {{ end -}}
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
    echo 🎉🎉🎉
    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"

    ```


## Check

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{- $d := .Decoration -}}
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
      exit 0
    {{ end -}}
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeCheck") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_check") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "check") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterCheck") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
    echo 🎉🎉🎉
    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
    ```


## Inputs


## Configs

`afterCheck`:




`cmdArg`:

    -c


`ports`:




`strictMode`:

    true


`_finish`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`afterStart`:




`finish`:




`includeShellUtil`:

    true


`runInLocal`:

    true


`setup`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`beforeStart`:




`check`:

    {{- $d := .Decoration -}}
    {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
      {{ if ne $port "" -}}
        echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
        waitPort "localhost" {{ $port }}
        echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
      {{ end -}}
    {{ end -}}



`start`:




`beforeCheck`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1