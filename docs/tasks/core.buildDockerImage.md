# core.buildDockerImage
```
  TASK NAME     : core.buildDockerImage
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.buildDockerImage.zaruba.yaml
  DESCRIPTION   : Build docker image.
                  Common config:
                    imageName : Image name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                          : set -e
                                                    {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                          : Blank
                  afterStart                      : Blank
                  beforeStart                     : Blank
                  buildArg                        : Blank
                  cmd                             : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                          : -c
                  finish                          : Blank
                  imagePrefix                     : Blank
                  imagePrefixTrailingSlash        : true
                  imageTag                        : Blank
                  includeUtilScript               : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup                           : Blank
                  start                           : set -e
                                                    {{ $d := .Decoration -}}
                                                    DOCKER_IMAGE_PREFIX="{{ .GetDockerImagePrefix }}"
                                                    if [ ! -f "$(pwd)/Dockerfile" ]
                                                    then
                                                      echo "{{ $d.Bold }}{{ $d.Red }}'Dockerfile' should be exist{{ $d.Normal }}"
                                                      exit 1
                                                    fi
                                                    DOCKER_IMAGE_NAME="{{ if .GetConfig "imageName" }}{{ .GetConfig "imageName" }}{{ else }}$("{{ .ZarubaBin }}" getServiceName "$(pwd)"){{ end }}"
                                                    DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
                                                    if [ ! -z "${DOCKER_IMAGE_TAG}" ]
                                                    then
                                                      docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "${DOCKER_IMAGE_PREFIX}${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}" -t "${DOCKER_IMAGE_PREFIX}${DOCKER_IMAGE_NAME}:latest" .
                                                    else
                                                      docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "${DOCKER_IMAGE_PREFIX}${DOCKER_IMAGE_NAME}:latest" .
                                                    fi
                                                    echo 🎉🎉🎉
                                                    echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image built{{ $d.Normal }}"
                  start.buildDockerImage.buildArg : {{ range $index, $buildArg := .Split (.Trim (.GetConfig "buildArg") "\n" ) "\n" -}}
                                                      {{ if ne $buildArg "" -}}
                                                        --build-arg {{ $buildArg }} {{ "" -}}
                                                      {{ end -}}
                                                    {{ end -}}
                  useImagePrefix                  : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```