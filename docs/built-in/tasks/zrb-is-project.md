<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# zrbIsProject
<!--endTocHeader-->


## Information

File Location:

    ${ZARUBA_HOME}zaruba-tasks/_base/validation/task.zrbIsProject.yaml

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


### Configs.finish


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

    echo "Project dir: {{ .ProjectDir }}"
    echo "Task dir: {{ .TaskDir }}"
    if [ ! -f "{{ .GetProjectPath "index.zaruba.yaml" }}" ] && [ ! -f "{{ .GetProjectPath "index.zaruba.yml" }}" ]
    then
      echo "${_BOLD}${_RED}${ZARUBA_PROJECT_DIR} is not a zaruba project.${_NORMAL}"
      exit 1
    fi
    echo "${_BOLD}${_YELLOW}${ZARUBA_PROJECT_DIR} is a valid zaruba project${_NORMAL}"



### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1



<!--startTocSubtopic-->
<!--endTocSubtopic-->