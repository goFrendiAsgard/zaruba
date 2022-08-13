<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🐳 zrbStartDockerCompose
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/start/task.zrbStartDockerCompose.yaml

Should Sync Env:

    true

Type:

    long running

Description:

    Start docker compose and wait until it is ready.
    Common configs:
      setup          : Script to be executed before start app or check app readiness.
      beforeStart    : Script to be executed before start app.
      afterStart     : Script to be executed after start app.
      beforeCheck    : Script to be executed before check app readiness.
      afterCheck     : Script to be executed before check app readiness.
      finish         : Script to be executed after start app or check app readiness.
      runInLocal     : Run app locally or not.
      ports          : Port to be checked to confirm app readiness, 
                       separated by new line.
      localhost      : Localhost mapping (e.g., host.docker.container)



## Extends

* [zrbStartApp](zrb-start-app.md)


## Dependencies

* [updateProjectLinks](update-project-links.md)
* [zrbCreateDockerNetwork](zrb-create-docker-network.md)


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
      echo "${_SCRIPT_ICON} ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"
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
    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
    echo "${_SCRIPT_ICON} ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"

    ```


## Check

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
      echo "${_SCRIPT_ICON} ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
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
    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
    echo "${_SCRIPT_ICON} ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
    ```


## Configs


### Configs._check

Value:

    {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
      {{ if ne $port "" -}}
        echo "${_INSPECT_ICON} ${_BOLD}${_YELLOW}Waiting for port '{{ $port }}'${_NORMAL}"
        waitPort "localhost" {{ $port }}
        echo "${_INSPECT_ICON} ${_BOLD}${_YELLOW}Port '{{ $port }}' is ready${_NORMAL}"
      {{ end -}}
    {{ end -}}
    {{ if .GetConfig "checkCommand" -}}
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    sleep 3
    {{ $checkCommand := .Util.Str.Trim (.GetConfig "checkCommand") "\n" -}}
    echo "${_INSPECT_ICON} ${_BOLD}${_YELLOW}Run check: {{ .Util.Str.EscapeShellValue $checkCommand }}${_NORMAL}"
    {{ $checkCommand }}
    echo "${_INSPECT_ICON} ${_BOLD}${_YELLOW}Successfully run check: {{ .Util.Str.EscapeShellValue $checkCommand }}${_NORMAL}"
    set "${_OLD_STATE}"
    {{ end -}}



### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start

Value:

    if [ -z "$(docker-compose top)" ]
    then
      {{ .GetConfig "_startEnv" }}
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Starting docker compose${_NORMAL}"
      {{ .GetConfig "prepareDockerComposeEnv" }}
      docker-compose up -d
    else
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Docker compose is already running${_NORMAL}"
    fi
    echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Logging docker compose${_NORMAL}"
    docker-compose logs --follow



### Configs._startEnv

Value:

    {{ $this := . -}}
    {{ if eq (.GetConfig "localhost") "localhost" -}}
      {{ range $key, $val := $this.GetEnvs -}}
        export {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }}
      {{ end -}}
    {{ else -}}
      {{ range $key, $val := $this.GetEnvs -}}
        {{ $val = $this.ReplaceAll $val "localhost" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "127.0.0.1" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "0.0.0.0" ($this.GetConfig "localhost") -}}
        export {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }}
      {{ end -}}
    {{ end -}}


### Configs.afterCheck


### Configs.afterStart


### Configs.beforeCheck


### Configs.beforeStart


### Configs.check


### Configs.checkCommand


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.imageName


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.imageTag


### Configs.localhost

Value:

    localhost


### Configs.ports


### Configs.prepareDockerComposeEnv

Value:

    export DOCKER_COMPOSE_IMAGE_NAME="{{ .GetDockerImageName }}"
    export DOCKER_COMPOSE_IMAGE_TAG="{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}"


### Configs.runInLocal

Value:

    true


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.useImagePrefix

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1