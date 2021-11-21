
# ZrbPushDockerImage

File Location:

    /zaruba-tasks/_base/dockerChore/task.zrbPushDockerImage.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Push docker image.
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


### Configs.imageTag


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.imageName


### Configs._start


### Configs.strictMode

Value:

    true


### Configs.afterStart


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.start

Value:

    {{ $d := .Decoration -}}
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
    echo "${DOCKER_IMAGE_NAME}"
    if [ ! -z "${DOCKER_IMAGE_TAG}" ]
    then
      docker push "${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
    fi
    docker push "${DOCKER_IMAGE_NAME}:latest"
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${DOCKER_IMAGE_NAME} pushed{{ $d.Normal }}"



### Configs._finish


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.includeShellUtil

Value:

    true


### Configs.setup


### Configs.useImagePrefix

Value:

    true


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1