
# ServeHttp

File Location:

    /zaruba-tasks/chore/serveHttp/task.serveHttp.yaml

Should Sync Env:

    true

Type:

    service

Description:

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

Description:

    HTTP port to serve static files

Prompt:

    HTTP port

Default Value:

    8080

Secret:

    false

Validation:

    ^[0-9]+$

Options:

    8080; 8000; 3000; 5000


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._start


### Configs.afterCheck


### Configs.afterStart


### Configs.beforeCheck


### Configs.start


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.includeShellUtil

Value:

    true


### Configs.ports

Value:

    {{ .GetValue "serverHttpPort" }}


### Configs.runInLocal

Value:

    true


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.check

Value:

    {{- $d := .Decoration -}}
    {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
      {{ if ne $port "" -}}
        echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
        waitPort "localhost" {{ $port }}
        echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
      {{ end -}}
    {{ end -}}



### Configs.finish


### Configs.cmdArg

Value:

    -c


### Configs.setup


### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1