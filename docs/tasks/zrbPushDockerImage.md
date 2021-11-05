# zrbPushDockerImage
```
  TASK NAME     : zrbPushDockerImage
  LOCATION      : /zaruba-tasks/_base/dockerChore/task.zrbPushDockerImage.yaml
  DESCRIPTION   : Push docker image.
                  Common configs:
                    imageName : Image name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish          : Blank
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  imageName        : Blank
                  imagePrefix      : {{ .GetValue "defaultImagePrefix" }}
                  imageTag         : Blank
                  includeShellUtil : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
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
                  useImagePrefix   : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```