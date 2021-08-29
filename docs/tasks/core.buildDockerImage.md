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
                  dockerFilePath                  : Dockerfile
                  finish                          : Blank
                  imagePrefix                     : Blank
                  imageTag                        : Blank
                  includeUtilScript               : . ${ZARUBA_HOME}/bash/util.sh
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