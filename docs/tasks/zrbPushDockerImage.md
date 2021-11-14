
# ZrbPushDockerImage

`File Location`:

    /zaruba-tasks/_base/dockerChore/task.zrbPushDockerImage.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

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


## Check




## Inputs


## Configs

`afterStart`:




`imageName`:




`includeShellUtil`:

    true


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`beforeStart`:




`finish`:




`imagePrefix`:

    {{ .GetValue "defaultImagePrefix" }}


`imageTag`:




`strictMode`:

    true


`_finish`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`start`:

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



`useImagePrefix`:

    true


`setup`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1