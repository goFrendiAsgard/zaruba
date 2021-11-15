
# ZrbRunInDockerContainer

File Location:

    /zaruba-tasks/_base/run/inDockerContainer/task.zrbRunInDockerContainer.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Run command from inside the container
    Common configs:
      containerName  : Name of the container
      containerShell : Shell to run script, default to sh
      command       : Command to be executed, can be multi line



## Extends

* `zrbGenerateAndRun`


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


### Configs._prepareReplacementMap


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.cmdArg

Value:

    -c


### Configs.containerName


### Configs.start


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._prepareVariables


### Configs.finish


### Configs._finish


### Configs._validateTemplateLocation

Value:

    {{ $d := .Decoration -}}
    if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
    then
      echo "{{ $d.Red }}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.{{ $d.Normal }}"
      exit 1
    fi



### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.includeShellUtil

Value:

    true


### Configs.script

Value:

    {{ .GetValue "script" }}


### Configs.templateLocation

Value:

    {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template


### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"


### Configs.remoteCommand

Value:

    sh "{{ .GetConfig "remoteScriptLocation" }}/run.sh"


### Configs.runGeneratedScript

Value:

    _ZRB_CONTAINER_NAME="{{ .GetConfig "containerName" }}"
    _ZRB_REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
    chmod 755 -R "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    docker cp "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_CONTAINER_NAME}:${_ZRB_REMOTE_SCRIPT_LOCATION}"
    docker exec "${_ZRB_CONTAINER_NAME}" {{ .GetConfig "remoteCommand" }}
    docker exec -u 0 "${_ZRB_CONTAINER_NAME}" rm -Rf "${_ZRB_REMOTE_SCRIPT_LOCATION}"
    rm -Rf "${_ZRB_GENERATED_SCRIPT_LOCATION}"


### Configs.sql

Value:

    {{ .GetValue "sql" }}


### Configs.strictMode

Value:

    true


### Configs._start

Value:

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



### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"


### Configs._validate


### Configs.setup


### Configs.generatedScriptLocation

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


### Configs.remoteScriptLocation

Value:

    _{{ .Name }}.script.{{ .UUID }}


### Configs.afterStart

Value:

    {{ $d := .Decoration -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"



## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1