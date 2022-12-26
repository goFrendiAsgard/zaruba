<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🐳 zrbBuildDockerImage
<!--endTocHeader-->

## Information

File Location:

    ~/.zaruba/zaruba-tasks/_base/docker/task.zrbBuildDockerImage.yaml

Should Sync Env:

    true

Type:

    simple

Description:

    Build docker image.
    Common configs:
      imageName : Image name



## Extends

* [zrbRunShellScript](zrb-run-shell-script.md)


## Dependencies

* [updateProjectLinks](update-project-links.md)


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



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.buildArg


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.dockerFilePath

Value:

    Dockerfile


### Configs.finish


### Configs.imageName


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.imageTag


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start

Value:

    DOCKER_FILE="{{ .GetConfig "dockerFilePath" }}"
    if [ ! -f "${DOCKER_FILE}" ]
    then
      echo "${_BOLD}${_RED}${DOCKER_FILE} should be exist${_NORMAL}"
      exit 1
    fi
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    DOCKER_IMAGE_TAG="{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}"
    echo "${_BOLD}${_YELLOW}Build image ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}${_NORMAL}"
    docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} \
      -t "${DOCKER_IMAGE_NAME}:latest" \
      -t "${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}" \
      -f "${DOCKER_FILE}" .
    echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
    echo "${_BOLD}${_YELLOW}Docker image ${DOCKER_IMAGE_NAME} built${_NORMAL}"



### Configs.start.buildDockerImage.buildArg

Value:

    {{ range $index, $buildArg := .Util.Str.Split (.Util.Str.Trim (.GetConfig "buildArg") "\n" ) "\n" -}}
      {{ if ne $buildArg "" -}}
        --build-arg {{ $buildArg }} {{ "" -}}
      {{ end -}}
    {{ end -}}



### Configs.strictMode

Value:

    true


### Configs.useImagePrefix

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1
<!--startTocSubtopic-->

<!--endTocSubtopic-->