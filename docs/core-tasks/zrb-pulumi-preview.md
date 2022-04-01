<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🦆 zrbPulumiPreview
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/pulumi/task.zrbPulumiPreview.yaml

Should Sync Env:

    true

Type:

    command


## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


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
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._preparePulumi

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "pulumiUseLocalBackend") -}}
    mkdir -p ./pulumiLock
    PULUMI_BACKEND_URL="file://./pulumiLock"
    {{ else -}}
    PULUMI_BACKEND_URL=${ZARUBA_CONFIG_PLUMI_BACKEND_URL}
    {{ end -}}
    pulumi stack select "${ZARUBA_CONFIG_PULUMI_STACK}" || pulumi stack init "${ZARUBA_CONFIG_PULUMI_STACK}" 


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start

Value:

    {{ .GetConfig "_preparePulumi" }}
    pulumi preview


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.pulumiBackendUrl


### Configs.pulumiStack

Value:

    {{ if .GetValue "pulumiStack" }}{{ .GetValue "pulumiStack" }}{{ else }}dev{{ end }}


### Configs.pulumiUseLocalBackend

Value:

    {{ .GetValue "pulumiUseLocalBackend" }}


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

Value:

    echo hello world


### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1