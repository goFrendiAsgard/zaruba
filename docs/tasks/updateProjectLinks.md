
# UpdateProjectLinks

File Location:

    /zaruba-tasks/chore/link/task.updateProjectLinks.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Update "links" in your project. Very useful if you have multiple apps sharing some parts of code
    USAGE:
      zaruba please updateProjectLinks
      zaruba please updateProjectLinks "link::fibo/css=common-css"
      zaruba please updateProjectLinks "link::app/css=common-css"
    ARGUMENTS
      link::<destination> : Location of the shared code
    TIPS:
      It is recommended to put `link` arguments in `default.values.yaml`.
      In order to do that, you can invoke `zaruba please addProjectLink <linkFrom=source-location> <linkTo=destination-location>`



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


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



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


### Configs.includeShellUtil

Value:

    true


### Configs.setup


### Configs.start

Value:

    {{ $this := . -}}
    {{ $destinations := .GetSubValueKeys "link" -}}
    {{ range $index, $destination := $destinations -}}
      {{ $source := $this.GetValue "link" $destination -}}
      {{ $absSource := $this.GetWorkPath $source -}}
      {{ $absDestination := $this.GetWorkPath $destination -}}
      linkResource "{{ $absSource }}" "{{ $absDestination }}"
    {{ end -}}
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Links updated${_NORMAL}"



### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1