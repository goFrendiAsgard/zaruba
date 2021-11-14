
# ZrbIsValidSubrepos

File Location:

    /zaruba-tasks/_base/validation/task.zrbIsValidSubrepos.yaml


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


### Configs.beforeStart

Value:





### Configs.cmdArg

Value:

    -c



### Configs.finish

Value:





### Configs.setup

Value:





### Configs.start

Value:

    {{ $d := .Decoration -}}
    {{ $names := .GetSubValueKeys "subrepo" -}}
    {{ $this := . -}}
    {{ range $index, $name := $names -}}
      PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
      URL="{{ $this.GetValue "subrepo" $name "url" }}"
      NAME="{{ $name }}"
      if [ -z "${URL}" ]
      then
        echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have url{{ $d.Normal }}"
        exit 1
      fi
      if [ -z "${PREFIX}" ]
      then
        echo "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have prefix{{ $d.Normal }}"
        exit 1
      fi
    {{ end }}
    echo "{{ $d.Bold }}{{ $d.Yellow }}All Subrepos are valid{{ $d.Normal }}"



### Configs.strictMode

Value:

    true



### Configs.includeShellUtil

Value:

    true



### Configs._finish

Value:





### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}




### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}



### Configs._start

Value:





### Configs.afterStart

Value:





### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED


Default:

    1