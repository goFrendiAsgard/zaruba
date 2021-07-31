# core.pullDockerImage
```
  TASK NAME     : core.pullDockerImage
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.pullDockerImage.zaruba.yaml
  DESCRIPTION   : Pull docker image.
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
  CONFIG        : _setup                      : set -e
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                      : Blank
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  dockerEnv                   : {{ .GetValue "dockerEnv" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helmEnv" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : true
                  includeUtilScript           : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  initDockerImagePrefixScript : {{ if .IsFalse (.GetConfig "useImagePrefix") -}}
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
                  setup                       : Blank
                  start                       : {{ $d := .Decoration -}}
                                                {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                IMAGE_NAME="{{ if .GetConfig "imageName" }}{{ .GetConfig "imageName" }}{{ else }}$({{ .Zaruba }} getServiceName "$(pwd)"){{ end }}"
                                                IMAGE_TAG="{{ .GetConfig "imageTag" }}"
                                                if [ ! -z "${IMAGE_TAG}" ]
                                                then
                                                  pull_image "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${IMAGE_TAG}"
                                                else
                                                  pull_image "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}"
                                                fi
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${DOCKER_IMAGE_PREFIX}${IMAGE_NAME} pulled{{ $d.Normal }}"
                  useImagePrefix              : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```