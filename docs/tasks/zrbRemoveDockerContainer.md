
# ZrbRemoveDockerContainer

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbRemoveDockerContainer.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Remove docker container.
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


## Configs


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.afterStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.containerName


### Configs.includeShellUtil

Value:

    true


### Configs.beforeStart


### Configs.setup


### Configs.strictMode

Value:

    true


### Configs._start


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.finish


### Configs.start

Value:

    {{ $d := .Decoration -}}
    CONTAINER="{{ if .GetConfig "containerName" }}{{ .GetConfig "containerName" }}{{ else }}$("{{ .ZarubaBin }}" path getAppName "$(pwd)"){{ end }}"
    echo "{{ $d.Bold }}{{ $d.Yellow }}Stop docker container ${CONTAINER}{{ $d.Normal }}"
    stopContainer "${CONTAINER}" 
    echo "{{ $d.Bold }}{{ $d.Yellow }}Remove docker container ${CONTAINER}{{ $d.Normal }}"
    removeContainer "${CONTAINER}" 
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker container ${CONTAINER} removed{{ $d.Normal }}"



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1