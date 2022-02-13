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

    service

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
      localhost      : Localhost mapping (e.g: host.docker.container)



## Extends

* [zrbStartApp](zrbStartApp.md)


## Dependencies

* [updateProjectLinks](updateProjectLinks.md)
* [zrbCreateDockerNetwork](zrbCreateDockerNetwork.md)


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"
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
    echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"

    ```


## Check

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
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
    echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
    ```


## Configs


### Configs._check

Value:

    {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
      {{ if ne $port "" -}}
        echo "📜 ${_BOLD}${_YELLOW}Waiting for port '{{ $port }}'${_NORMAL}"
        waitPort "localhost" {{ $port }}
        echo "📜 ${_BOLD}${_YELLOW}Port '{{ $port }}' is ready${_NORMAL}"
      {{ end -}}
    {{ end -}}



### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


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


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.localhost

Value:

    localhost


### Configs.ports


### Configs.runInLocal

Value:

    true


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitConfigVariables

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start

Value:

    {{ .GetConfig "_startEnv" }}
    echo "🐳 ${_BOLD}${_YELLOW}Starting docker compose${_NORMAL}"
    docker-compose up -d
    echo "🐳 ${_BOLD}${_YELLOW}Logging docker compose${_NORMAL}"
    docker-compose logs --follow



### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1