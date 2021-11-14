
# ZrbSetKubeContext

File Location:

    /zaruba-tasks/_base/kubeChore/task.zrbSetKubeContext.yaml


Location:




Should Sync Env:

    true


Sync Env Location:




Type:

    command


Description:





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


### Configs._finish

Value:





### Configs._start

Value:





### Configs.beforeStart

Value:





### Configs.finish

Value:





### Configs.strictMode

Value:

    true



### Configs.cmdArg

Value:

    -c



### Configs.includeShellUtil

Value:

    true



### Configs.setup

Value:





### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}




### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs.afterStart

Value:





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



### Configs.kubeContext

Value:

    {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}



### Configs.kubeNamespace

Value:

    {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}



### Configs.start

Value:

    if [ "$(kubectl config current-context)" != "{{ .GetConfig "kubeContext" }}" ]
    then
      kubectl config use-context "{{ .GetConfig "kubeContext" }}"
    fi



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1