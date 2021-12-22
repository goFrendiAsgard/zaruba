
# ZrbRunInKubePod

File Location:

    /zaruba-tasks/_base/run/inKubePod/task.zrbRunInKubePod.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Run command in a kubernetes pod.
    Common configs:
      podLabel       : Label of the pod.
      podName        : Name of the pod.
      podShell       : Shell to run script, default to sh.
      remoteCommand  : Command to be executed.
      script         : Script to be executed (Can be multi line).



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


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._prepareBaseReplacementMap

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"


### Configs._prepareBaseVariables

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"


### Configs._prepareReplacementMap


### Configs._prepareVariables


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
    _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
    _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
    _ZRB_REPLACEMENT_MAP='{}'
    __ZRB_PWD=$(pwd)
    echo "${_YELLOW}🧰 Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}✅ Validate${_NORMAL}"
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🚧 Generate${_NORMAL}"
    echo "${_YELLOW}🚧 Template Location:${_NORMAL} ${_ZRB_TEMPLATE_LOCATION}"
    echo "${_YELLOW}🚧 Generated Script Location:${_NORMAL} ${_ZRB_GENERATED_SCRIPT_LOCATION}"
    echo "${_YELLOW}🚧 Replacement Map:${_NORMAL} ${_ZRB_REPLACEMENT_MAP}"
    mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}🏁 Generated Script${_NORMAL}"
    echo "${_ZRB_CFG_RUN_GENERATED_SCRIPT}"
    echo "${_YELLOW}🏁 Run Generated Script${_NORMAL}"
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



### Configs._validate


### Configs._validateTemplateLocation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/validateTemplateLocation.sh"


### Configs.afterStart

Value:

    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Done${_NORMAL}"



### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.generatedScriptLocation

Value:

    {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}


### Configs.kubeContext

Value:

    {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}


### Configs.kubeNamespace

Value:

    {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}


### Configs.podLabel


### Configs.podName


### Configs.podShell

Value:

    bash


### Configs.remoteCommand

Value:

    {{ .GetConfig "podShell" }} "{{ .GetConfig "remoteScriptLocation" }}/run.sh"


### Configs.remoteScriptLocation

Value:

    _{{ .Name }}.script.{{ .UUID }}


### Configs.runGeneratedScript

Value:

    _ZRB_REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
    _ZRB_KUBE_NAMESPACE="{{ .GetConfig "kubeNamespace" }}"
    _ZRB_KUBE_CONTEXT="{{ .GetConfig "kubeContext" }}"
    _ZRB_POD_NAME="{{ if .GetConfig "podName" }}{{ .GetConfig "podName" }}{{ else }}$(kubectl get pods -o name --context "${_ZRB_KUBE_CONTEXT}" --namespace "${_ZRB_KUBE_NAMESPACE}" -l "{{ .GetConfig "podLabel" }}" | head -n 1 | cut -d'/' -f 2){{ end }}"
    kubectl cp --context "${_ZRB_KUBE_CONTEXT}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_KUBE_NAMESPACE}/${_ZRB_POD_NAME}:${_ZRB_REMOTE_SCRIPT_LOCATION}"
    kubectl exec --context "${_ZRB_KUBE_CONTEXT}" --namespace "${_ZRB_KUBE_NAMESPACE}" "${_ZRB_POD_NAME}" -- "bash" "${_ZRB_REMOTE_SCRIPT_LOCATION}"
    kubectl exec --context "${_ZRB_KUBE_CONTEXT}" --namespace "${_ZRB_KUBE_NAMESPACE}" "${_ZRB_POD_NAME}" -- "{{ .GetConfig "podShell" }}" "-c" "rm -Rf ${_ZRB_REMOTE_SCRIPT_LOCATION}"
    rm -Rf "${_ZRB_GENERATED_SCRIPT_LOCATION}"


### Configs.script

Value:

    {{ .GetValue "script" }}


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    true


### Configs.shouldInitConfigVariables

Value:

    true


### Configs.shouldInitEnvMapVariable

Value:

    true


### Configs.shouldInitUtil

Value:

    true


### Configs.sql

Value:

    {{ .GetValue "sql" }}


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.templateLocation

Value:

    {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1