# zrbBuildDockerImage
```
  TASK NAME     : zrbBuildDockerImage
  LOCATION      : /zaruba-tasks/_base/dockerChore/task.zrbBuildDockerImage.yaml
  DESCRIPTION   : Build docker image.
                  Common config:
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
  CONFIG        : _finish                         : Blank
                  _setup                          : set -e
                                                    {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                          : Blank
                  afterStart                      : Blank
                  beforeStart                     : Blank
                  buildArg                        : Blank
                  cmd                             : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                          : -c
                  dockerFilePath                  : Dockerfile
                  finish                          : Blank
                  imageName                       : Blank
                  imagePrefix                     : {{ .GetValue "defaultImagePrefix" }}
                  imageTag                        : Blank
                  includeShellUtil                : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup                           : Blank
                  start                           : set -e
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
                  start.buildDockerImage.buildArg : {{ range $index, $buildArg := .Util.Str.Split (.Util.Str.Trim (.GetConfig "buildArg") "\n" ) "\n" -}}
                                                      {{ if ne $buildArg "" -}}
                                                        --build-arg {{ $buildArg }} {{ "" -}}
                                                      {{ end -}}
                                                    {{ end -}}
                  useImagePrefix                  : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```