<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🚢 zrbCopyToKubePod
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/copyToKubePod/task.zrbCopyToKubePod.yaml

Should Sync Env:

    true

Type:

    simple

Description:

    Generate scripts and copy them to kubernetes pod.
    Common configs:
      podLabel         : Label of the pod.
      podName          : Name of the pod.
      podShell         : Shell to run script, default to sh.
      templateLocation : Template script location (source).
      remoteScript     : Remote script location (destination).



## Extends

* [zrbGenerateAndRun](zrb-generate-and-run.md)


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
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Prepare${_NORMAL}"
    {{ .GetConfig "_prepareBaseVariables" }}
    {{ .GetConfig "_prepareVariables" }}
    {{ .GetConfig "_prepareBaseReplacementMap" }}
    {{ .GetConfig "_prepareReplacementMap" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Validate${_NORMAL}"
    {{ .GetConfig "_validateTemplateLocation" }}
    {{ .GetConfig "_validate" }}
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Generate${_NORMAL}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Template Location:${_NORMAL} ${_ZRB_TEMPLATE_LOCATION}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Generated Script Location:${_NORMAL} ${_ZRB_GENERATED_SCRIPT_LOCATION}"
    echo "${_YELLOW}${_CONSTRUCTION_ICON} Replacement Map:${_NORMAL} ${_ZRB_REPLACEMENT_MAP}"
    mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
    "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
    cd "${__ZRB_PWD}"
    echo "${_YELLOW}${_START_ICON} Generated Script${_NORMAL}"
    echo "${ZARUBA_CONFIG_RUN_GENERATED_SCRIPT}"
    echo "${_YELLOW}${_START_ICON} Run Generated Script${_NORMAL}"
    {{ .GetConfig "runGeneratedScript" }}
    cd "${__ZRB_PWD}"



### Configs._validate


### Configs._validateTemplateLocation

Value:

    . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/validateTemplateLocation.sh"


### Configs.afterStart

Value:

    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
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


### Configs.remoteScriptLocation

Value:

    _{{ .Name }}.script


### Configs.runGeneratedScript

Value:

    _ZRB_REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
    _ZRB_KUBE_NAMESPACE="{{ .GetConfig "kubeNamespace" }}"
    _ZRB_KUBE_CONTEXT="{{ .GetConfig "kubeContext" }}"
    echo "${_BOLD}${_YELLOW}Get pod name${_NORMAL}"
    _ZRB_POD_NAME="{{ if .GetConfig "podName" }}{{ .GetConfig "podName" }}{{ else }}$(kubectl get pods -o name --context "${_ZRB_KUBE_CONTEXT}" --namespace "${_ZRB_KUBE_NAMESPACE}" -l "{{ .GetConfig "podLabel" }}" | head -n 1 | cut -d'/' -f 2){{ end }}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Remove ${_ZRB_REMOTE_SCRIPT_LOCATION} at pod ${_ZRB_POD_NAME}${_NORMAL}"
    kubectl exec -n "${_ZRB_KUBE_NAMESPACE}" "${_ZRB_POD_NAME}" -- "{{ .GetConfig "podShell" }}" "-c" "rm -Rf ${_ZRB_REMOTE_SCRIPT_LOCATION}"
    echo "${_BOLD}${_YELLOW}${_WORKER_ICON} Copy from ${_ZRB_GENERATED_SCRIPT_LOCATION} at host to ${_ZRB_REMOTE_SCRIPT_LOCATION} at pod ${_ZRB_POD_NAME}${_NORMAL}"
    kubectl cp "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_KUBE_NAMESPACE}/${_ZRB_POD_NAME}:${_ZRB_REMOTE_SCRIPT_LOCATION}"



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

    {{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/template


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1