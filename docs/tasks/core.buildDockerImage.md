# core.buildDockerImage
```
  TASK NAME     : core.buildDockerImage
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.buildDockerImage.zaruba.yaml
  DESCRIPTION   : Build docker image.
                  Common config:
                    dockerEnv : Docker environment key (default: '{{ .GetValue "dockerEnv" }}')
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
                                                    alias zaruba=${ZARUBA_HOME}/zaruba
                                                    {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                    {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                          : Blank
                  afterStart                      : Blank
                  beforeStart                     : Blank
                  buildArg                        : Blank
                  cmd                             : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                          : -c
                  dockerEnv                       : {{ .GetValue "dockerEnv" }}
                  finish                          : Blank
                  helmEnv                         : {{ .GetValue "helmEnv" }}
                  imagePrefix                     : Blank
                  imagePrefixTrailingSlash        : true
                  includeBootstrapScript          : if [ -f "${HOME}/.profile" ]
                                                    then
                                                        . "${HOME}/.profile"
                                                    fi
                                                    if [ -f "${HOME}/.bashrc" ]
                                                    then
                                                        . "${HOME}/.bashrc"
                                                    fi
                                                    BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                                    . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript               : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  initDockerImagePrefixScript     : {{ if .IsFalse (.GetConfig "useImagePrefix") -}}
                                                      DOCKER_IMAGE_PREFIX=""
                                                    {{ else if .GetConfig "imagePrefix" -}}
                                                      DOCKER_IMAGE_PREFIX="{{ .GetConfig "imagePrefix" }}"
                                                    {{ else if and (.GetConfig "dockerEnv") (.GetValue "dockerImagePrefix" (.GetConfig "dockerEnv")) -}}
                                                      DOCKER_IMAGE_PREFIX="{{ .GetValue "dockerImagePrefix" (.GetConfig "dockerEnv") }}"
                                                    {{ else if .GetValue "dockerImagePrefix" "default" -}}
                                                      DOCKER_IMAGE_PREFIX="{{ .GetValue "dockerImagePrefix" "default" }}"
                                                    {{ else -}}
                                                      DOCKER_IMAGE_PREFIX="local"
                                                    {{ end -}}
                                                    {{ if .IsTrue (.GetConfig "imagePrefixTrailingSlash" ) -}}
                                                      if [ ! -z "${DOCKER_IMAGE_PREFIX}" ]
                                                      then
                                                        DOCKER_IMAGE_PREFIX="${DOCKER_IMAGE_PREFIX}/"
                                                      fi
                                                    {{ end -}}
                  setup                           : Blank
                  start                           : set -e
                                                    {{ $d := .Decoration -}}
                                                    {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                    should_be_file "$(pwd)/Dockerfile" "{{ $d.Bold }}{{ $d.Red }}'Dockerfile' should be exist{{ $d.Normal }}"
                                                    IMAGE_NAME="{{ if .GetConfig "imageName" }}{{ .GetConfig "imageName" }}{{ else }}$({{ .Zaruba }} getServiceName "$(pwd)"){{ end }}"
                                                    COMMIT="$(get_latest_git_commit)"
                                                    if [ ! -z "${COMMIT}" ]
                                                    then
                                                      SHORT_COMMIT="$(echo "${COMMIT}" | cut -c1-12)"
                                                      TAG="$(get_latest_git_tag)"
                                                      if [ ! -z "${TAG}" ]
                                                      then
                                                        TAG_COMMIT="$(get_latest_git_tag_commit)"
                                                        if [ "${TAG_COMMIT}" = "${COMMIT}" ]
                                                        then
                                                          docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "local/${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${TAG}" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${TAG}-${SHORT_COMMIT}" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${SHORT_COMMIT}" .
                                                        else
                                                          docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "local/${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${TAG}-${SHORT_COMMIT}" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${SHORT_COMMIT}" .
                                                        fi
                                                      else
                                                        docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "local/${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${SHORT_COMMIT}" .
                                                      fi
                                                    else
                                                      docker build {{ .GetConfig "start.buildDockerImage.buildArg" }} -t "local/${IMAGE_NAME}:latest" -t "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:latest" .
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