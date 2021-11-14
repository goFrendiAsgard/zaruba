
# ZrbPullDockerImage

`File Location`:

    /zaruba-tasks/_base/dockerChore/task.zrbPullDockerImage.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Pull docker image.
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

`setup`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`imageTag`:




`includeShellUtil`:

    true


`start`:

    {{ $d := .Decoration -}}
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
    if [ ! -z "${DOCKER_IMAGE_TAG}" ]
    then
      pullImage "${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
    else
      pullImage "${DOCKER_IMAGE_NAME}"
    fi
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${DOCKER_IMAGE_NAME} pulled{{ $d.Normal }}"



`cmdArg`:

    -c


`finish`:




`imagePrefix`:

    {{ .GetValue "defaultImagePrefix" }}


`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`strictMode`:

    true


`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`imageName`:




`useImagePrefix`:

    true


`_finish`:




`_start`:




`afterStart`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1