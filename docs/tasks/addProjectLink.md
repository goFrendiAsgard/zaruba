
# AddProjectLink

`File Location`:

    /zaruba-tasks/chore/link/task.addProjectLink.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Add link.
    TIPS: To update links, you should perform `zaruba please updateProjectLinks`




## Extends

* `zrbRunShellScript`


## Dependencies

* `zrbIsProject`


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


### Inputs.linkFrom

`Default Value`:




`Description`:

    Link source (Required)


`Prompt`:

    Source


`Secret`:

    false


`Validation`:

    ^.+$


`Options`:





### Inputs.linkTo

`Default Value`:




`Description`:

    Link destination (Required)


`Prompt`:

    Destination


`Secret`:

    false


`Validation`:

    ^.+$


`Options`:





## Configs

`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_start`:




`cmdArg`:

    -c


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`afterStart`:




`beforeStart`:




`start`:

    {{ $d := .Decoration -}}
    "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "link::{{ .GetConfig "linkTo" }}" "{{ .GetConfig "linkFrom" }}"
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Link ${SOURCE} -> ${DESTINATION} has been added{{ $d.Normal }}"



`finish`:




`linkFrom`:

    {{ .GetValue "linkFrom" }}


`strictMode`:

    true


`_finish`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`includeShellUtil`:

    true


`linkTo`:

    {{ .GetValue "linkTo" }}


`setup`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1