<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 📜 zrbStartApp
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/start/task.zrbStartApp.yaml

Should Sync Env:

    true

Type:

    long running

Description:

    Start app and check its readiness.
    Common configs:
      setup       : Script to be executed before start app or check app readiness.
      start       : Script to start the app (e.g., python -m http.server 9090).
      beforeStart : Script to be executed before start app.
      afterStart  : Script to be executed after start app.
      beforeCheck : Script to be executed before check app readiness.
      afterCheck  : Script to be executed before check app readiness.
      finish      : Script to be executed after start app or check app readiness.
      runInLocal  : Run app locally or not.
      ports       : Port to be checked to confirm app readiness, separated by new line.



## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


## Dependencies

* [updateProjectLinks](update-project-links.md)


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
    until [ "$?" = "0" ]
    do
      sleep 3
      {{ $checkCommand }}
    done
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


### Configs.ports


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


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1