
# ServeHttp

`File Location`:

    /zaruba-tasks/chore/serveHttp/task.serveHttp.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    service


`Description`:

    Run static web server from your working directory.




## Extends

* `zrbStartApp`


## Dependencies

* `updateProjectLinks`


## Start

* `{{ .GetEnv "ZARUBA_HOME" }}/zaruba`
* `serve`
* `.`
* `{{ index (.Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n") 0 }}`


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


### Inputs.serverHttpPort

`Default Value`:

    8080


`Description`:

    HTTP port to serve static files


`Prompt`:

    HTTP port


`Secret`:

    false


`Validation`:

    ^[0-9]+$


`Options`:

    8080; 8000; 3000; 5000



## Configs

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



`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`start`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`cmdArg`:

    -c


`runInLocal`:

    true


`strictMode`:

    true


`_finish`:




`afterCheck`:




`afterStart`:




`beforeCheck`:




`includeShellUtil`:

    true


`ports`:

    {{ .GetValue "serverHttpPort" }}


`setup`:




`_start`:




`finish`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1