
# ZrbCreateDockerNetwork

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbCreateDockerNetwork.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:

    Create docker network.
    Common configs:
      network : Network name




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


### Configs.includeShellUtil

Value:

    true



### Configs.strictMode

Value:

    true



### Configs.afterStart

Value:





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



### Configs.cmdArg

Value:

    -c



### Configs.beforeStart

Value:





### Configs._start

Value:





### Configs.network

Value:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}



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




### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}




### Configs.finish

Value:





### Configs.setup

Value:





### Configs._finish

Value:





## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1