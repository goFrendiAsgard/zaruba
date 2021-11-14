
# ZrbGenerateAndRun

`File Location`:

    /zaruba-tasks/_base/generateAndRun/task.zrbGenerateAndRun.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Generate script and run it



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

`_prepareReplacementMap`:




`_validate`:




`runGeneratedScript`:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}/run.sh


`start`:




`_prepareBaseVariables`:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"


`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`_prepareBaseReplacementMap`:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"


`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`script`:

    {{ .GetValue "script" }}


`setup`:




`sql`:

    {{ .GetValue "sql" }}


`_finish`:




`strictMode`:

    true


`cmdArg`:

    -c


`finish`:




`includeShellUtil`:

    true


`afterStart`:

    {{ $d := .Decoration -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"



`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:

    {{ $d := .Decoration -}}
    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
    _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
    _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
    _ZRB_TASK_NAME="{{ .Name }}"
    _ZRB_REPLACEMENT_MAP='{}'
    _ZRB_SCRIPT='{{ .GetConfig "script" }}'
    _ZRB_SQL='{{ .GetConfig "sql" }}'
    _ZRB_IMAGE_NAME="{{ .GetDockerImageName }}"
    _ZRB_IMAGE_TAG="{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}"
    _ZRB_ENVS='{{ .ToJSON .GetEnvs }}'
    __ZRB_PWD=$(pwd)
    echo "{{ $d.Yellow }}🧰 Prepare{{ $d.Normal }}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "{{ $d.Yellow }}✅ Validate{{ $d.Normal }}"
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "{{ $d.Yellow }}🚧 Generate{{ $d.Normal }}"
    echo "{{ $d.Yellow }}🚧 Template Location:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATION}"
    echo "{{ $d.Yellow }}🚧 Generated Script Location:{{ $d.Normal }} ${_ZRB_GENERATED_SCRIPT_LOCATION}"
    echo "{{ $d.Yellow }}🚧 Replacement Map:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
    mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
    cd "${__ZRB_PWD}"
    echo "{{ $d.Yellow }}🏁 Run Script{{ $d.Normal }}"
    echo '{{ .GetConfig "runGeneratedScript" }}'
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



`_validateTemplateLocation`:

    {{ $d := .Decoration -}}
    if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
    then
      echo "{{ $d.Red }}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.{{ $d.Normal }}"
      exit 1
    fi



`generatedScriptLocation`:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


`templateLocation`:

    {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template


`_prepareVariables`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1