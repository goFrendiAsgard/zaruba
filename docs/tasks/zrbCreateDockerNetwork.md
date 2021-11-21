
# ZrbCreateDockerNetwork

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbCreateDockerNetwork.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Create docker network.
    Common configs:
      network : Network name



## Extends

* `zrbRunShellScript`


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


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._start


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.network

Value:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.setup


### Configs.afterStart


### Configs.beforeStart


### Configs.start

Value:

    {{ $d := .Decoration -}}
    set -e
    if [ "$(inspectDocker network ".Name" "{{ .GetConfig "network" }}")" = "{{ .GetConfig "network" }}" ]
    then
      echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Network '{{ .GetConfig "network" }}' is already exist{{ $d.Normal }}"
    else
      echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Creating network '{{ .GetConfig "network" }}'{{ $d.Normal }}"
      docker network create "{{ .GetConfig "network" }}"
    fi



### Configs.strictMode

Value:

    true


### Configs._finish


### Configs.includeShellUtil

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1