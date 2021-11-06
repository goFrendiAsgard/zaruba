# zrbPullDockerImage
```
  TASK NAME     : zrbPullDockerImage
  LOCATION      : /zaruba-tasks/_base/dockerChore/task.zrbPullDockerImage.yaml
  DESCRIPTION   : Pull docker image.
                  Common configs:
                    imageName : Image name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunShellScript ]
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
                  _initShell       : {{ if .IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
                                     {{ if .IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}
                  _setup           : {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  imageName        : Blank
                  imagePrefix      : {{ .GetValue "defaultImagePrefix" }}
                  imageTag         : Blank
                  includeShellUtil : true
                  setup            : Blank
                  start            : {{ $d := .Decoration -}}
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
                  strictMode       : true
                  useImagePrefix   : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```