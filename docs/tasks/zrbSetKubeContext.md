
# ZrbSetKubeContext

`File Location`:

    /zaruba-tasks/_base/kubeChore/task.zrbSetKubeContext.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:





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

`_start`:




`beforeStart`:




`start`:

    if [ "$(kubectl config current-context)" != "{{ .GetConfig "kubeContext" }}" ]
    then
      kubectl config use-context "{{ .GetConfig "kubeContext" }}"
    fi


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`finish`:




`strictMode`:

    true


`includeShellUtil`:

    true


`kubeContext`:

    {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}


`setup`:




`_finish`:




`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`afterStart`:




`kubeNamespace`:

    {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1