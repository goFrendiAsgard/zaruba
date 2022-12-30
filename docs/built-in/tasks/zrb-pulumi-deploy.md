<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# zrbPulumiDeploy
<!--endTocHeader-->


## Information

File Location:

    ${ZARUBA_HOME}zaruba-tasks/_base/pulumi/task.zrbPulumiDeploy.yaml

Should Sync Env:

    true

Type:

    simple


## Extends

- [zrbRunShellScript](zrb-run-shell-script.md)


## Start

- `{{ .GetConfig "cmd" }}`
- `{{ .GetConfig "cmdArg" }}`
-
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
    if [ ! -z "${ZARUBA_CONFIG_PULUMI_BACKEND_URL}" ]
    then
      export PULUMI_BACKEND_URL=${ZARUBA_CONFIG_PULUMI_BACKEND_URL}
    else
      mkdir -p ./pulumiLock
      export PULUMI_BACKEND_URL="file://./pulumiLock"
    fi
    echo "${_YELLOW}Pulumi backend URL: ${PULUMI_BACKEND_URL}${_NORMAL}"
    {{ end -}}
    pulumi stack select "${ZARUBA_CONFIG_PULUMI_STACK}" || pulumi stack init "${ZARUBA_CONFIG_PULUMI_STACK}" 
    echo "${_YELLOW}Pulumi stack: ${ZARUBA_CONFIG_PULUMI_STACK}${_NORMAL}"



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start

Value:

    {{ .GetConfig "_preparePulumi" }}
    pulumi up -y



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


### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1



# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->