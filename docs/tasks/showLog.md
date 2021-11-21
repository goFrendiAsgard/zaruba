
# ShowLog

File Location:

    /zaruba-tasks/chore/log/task.showLog.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Show log for all/particular tasks using regex



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


## Inputs


### Inputs.keyword

Description:

    Keyword

Prompt:

    Keyword

Secret:

    false


## Configs


### Configs._finish


### Configs.afterStart


### Configs.setup


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.finish


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.includeShellUtil

Value:

    true


### Configs.keyword

Value:

    {{ if .GetValue "keyword" }}{{ .GetValue "keyword" }}{{ else }}.*{{ end }}


### Configs.strictMode

Value:

    true


### Configs._start


### Configs.beforeStart


### Configs.cmdArg

Value:

    -c


### Configs.start

Value:

    {{ $d := .Decoration -}}
    if [ ! -f "log.zaruba.csv" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}Log is not exist{{ $d.Normal }}"
      exit 1
    fi
    "{{ .ZarubaBin }}" project showLog "{{ .GetWorkPath "log.zaruba.csv" }}" "{{ .GetConfig "keyword"}}"



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1