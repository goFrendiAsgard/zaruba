
# SetProjectValue

`File Location`:

    /zaruba-tasks/chore/value/task.setProjectValue.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Set project value.



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


### Inputs.variableName

`Default Value`:




`Description`:

    Variable name (Required)


`Prompt`:

    Name


`Secret`:

    false


`Validation`:

    ^.+$


`Options`:





### Inputs.variableValue

`Default Value`:




`Description`:

    Variable value (Required)


`Prompt`:

    Value


`Secret`:

    false


`Validation`:

    ^.+$


`Options`:





## Configs

`variableValue`:

    {{ .GetValue "variableValue" }}


`strictMode`:

    true


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`beforeStart`:




`setup`:




`cmdArg`:

    -c


`_finish`:




`afterStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`includeShellUtil`:

    true


`start`:

    {{ $d := .Decoration -}}
    "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "{{ .GetConfig "variableName" }}" "{{ .GetConfig "variableValue" }}"
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Kwarg ${KEY} : ${VALUE} has been set{{ $d.Normal }}"



`variableName`:

    {{ .GetValue "variableName" }}


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`finish`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1