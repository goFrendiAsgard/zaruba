
# ZrbStopDockerContainer

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbStopDockerContainer.yaml

Should Sync Env:

    true

Type:

    command

Description:

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


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.containerName


### Configs.finish


### Configs.includeShellUtil

Value:

    true


### Configs.setup


### Configs.start

Value:

    CONTAINER="{{ if .GetConfig "containerName" }}{{ .GetConfig "containerName" }}{{ else }}$("{{ .ZarubaBin }}" path getAppName "$(pwd)"){{ end }}"
    echo "${_BOLD}${_YELLOW}Stop docker container ${CONTAINER}${_NORMAL}"
    stopContainer "${CONTAINER}" 
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Docker container ${CONTAINER} stopped${_NORMAL}"



### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1