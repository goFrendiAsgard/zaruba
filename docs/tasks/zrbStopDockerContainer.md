
# ZrbStopDockerContainer

`File Location`:

    /zaruba-tasks/_base/dockerChore/task.zrbStopDockerContainer.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Stop docker container.
    Common configs:
      containerName : Container's name




## Extends

* `zrbRunShellScript`


## Dependencies

* `updateProjectLinks`


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

`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`strictMode`:

    true


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`beforeStart`:




`start`:

    {{ $d := .Decoration -}}
    CONTAINER="{{ if .GetConfig "containerName" }}{{ .GetConfig "containerName" }}{{ else }}$("{{ .ZarubaBin }}" path getAppName "$(pwd)"){{ end }}"
    echo "{{ $d.Bold }}{{ $d.Yellow }}Stop docker container ${CONTAINER}{{ $d.Normal }}"
    stopContainer "${CONTAINER}" 
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker container ${CONTAINER} stopped{{ $d.Normal }}"



`afterStart`:




`includeShellUtil`:

    true


`containerName`:




`_finish`:




`_start`:




`finish`:




`setup`:




`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`cmdArg`:

    -c



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1