# core.pushDockerImage
```
  TASK NAME     : core.pushDockerImage
  LOCATION      : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
  DESCRIPTION   : Push docker image.
                  Common config:
                    dockerEnv : Docker environment key (default: '{{ .GetValue "docker.env" }}')
                    imageName : Image name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.setupPyUtil, updateProjectLinks ]
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
                                                alias zaruba=${ZARUBA_HOME}/zaruba
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                      : Blank
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  dockerEnv                   : {{ .GetValue "docker.env" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helm.env" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : true
                  includeBootstrapScript      : if [ -f "${HOME}/.profile" ]
                                                then
                                                    . "${HOME}/.profile"
                                                fi
                                                if [ -f "${HOME}/.bashrc" ]
                                                then
                                                    . "${HOME}/.bashrc"
                                                fi
                                                BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                                . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript           : . "${ZARUBA_HOME}/scripts/util.sh"
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
                  kubeContext                 : {{ .GetValue "kube.context" }}
                  playBellScript              : echo $'\a'
                  setup                       : Blank
                  start                       : {{ $d := .Decoration -}}
                                                {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                IMAGE_NAME="{{ if .GetConfig "imageName" }}{{ .GetConfig "imageName" }}{{ else }}$("${ZARUBA_HOME}/zaruba" getDefaultServiceName "$(pwd)"){{ end }}"
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
                                                      docker push "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${TAG}"
                                                    fi
                                                    docker push "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${TAG}-${SHORT_COMMIT}"
                                                  fi
                                                  docker push "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:${SHORT_COMMIT}"
                                                fi
                                                docker push "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}:latest"
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${IMAGE_NAME} pushed{{ $d.Normal }}"
                  useImagePrefix              : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
