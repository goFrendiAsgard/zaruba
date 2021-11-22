
# ZrbSetKubeContext

File Location:

    /zaruba-tasks/_base/kubeChore/task.zrbSetKubeContext.yaml

Should Sync Env:

    true

Type:

    command


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


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.kubeContext

Value:

    {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}


### Configs.kubeNamespace

Value:

    {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}


### Configs.setup


### Configs.start

Value:

    if [ "$(kubectl config current-context)" != "{{ .GetConfig "kubeContext" }}" ]
    then
      kubectl config use-context "{{ .GetConfig "kubeContext" }}"
    fi


### Configs.afterStart


### Configs.strictMode

Value:

    true


### Configs._finish


### Configs._start


### Configs.cmdArg

Value:

    -c


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.finish


### Configs.includeShellUtil

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1