
# ZrbBuildDockerImage

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbBuildDockerImage.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Build docker image.
    Common configs:
      imageName : Image name



## Extends

* `zrbRunShellScript`


## Dependencies

* `updateProjectLinks`


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


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.finish


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.setup


### Configs.start

Value:

    DOCKER_FILE="{{ .GetConfig "dockerFilePath" }}"
    if [ ! -f "${DOCKER_FILE}" ]
    then
      echo "${_BOLD}${_RED}${DOCKER_FILE} should be exist${_NORMAL}"
      exit 1
    fi
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
    docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} \
      -t "${DOCKER_IMAGE_NAME}:latest" \
      -t "${DOCKER_IMAGE_NAME}:{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}" \
      -f "${DOCKER_FILE}" .
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Docker image built${_NORMAL}"



### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToEnvironmentVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.imageTag


### Configs.useImagePrefix

Value:

    true


### Configs._start


### Configs.buildArg


### Configs.cmdArg

Value:

    -c


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


### Configs._finish


### Configs.dockerFilePath

Value:

    Dockerfile


### Configs.imageName


### Configs.includeShellUtil

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1