
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


### Configs.imageTag


### Configs.strictMode

Value:

    true


### Configs.buildArg


### Configs.dockerFilePath

Value:

    Dockerfile


### Configs._start


### Configs.beforeStart


### Configs.cmdArg

Value:

    -c


### Configs.imageName


### Configs._finish


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs.setup


### Configs.useImagePrefix

Value:

    true


### Configs.afterStart


### Configs.includeShellUtil

Value:

    true


### Configs.finish


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.start

Value:

    set -e
    {{ $d := .Decoration -}}
    DOCKER_FILE="{{ .GetConfig "dockerFilePath" }}"
    if [ ! -f "${DOCKER_FILE}" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}${DOCKER_FILE} should be exist{{ $d.Normal }}"
      exit 1
    fi
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
    docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} \
      -t "${DOCKER_IMAGE_NAME}:latest" \
      -t "${DOCKER_IMAGE_NAME}:{{ if .GetConfig "imageTag" }}{{ .GetConfig "imageTag" }}{{ else }}latest{{ end }}" \
      -f "${DOCKER_FILE}" .
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image built{{ $d.Normal }}"



### Configs.start.buildDockerImage.buildArg

Value:

    {{ range $index, $buildArg := .Util.Str.Split (.Util.Str.Trim (.GetConfig "buildArg") "\n" ) "\n" -}}
      {{ if ne $buildArg "" -}}
        --build-arg {{ $buildArg }} {{ "" -}}
      {{ end -}}
    {{ end -}}



### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1