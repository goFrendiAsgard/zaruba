# core.pushDockerImage
```
  TASK NAME     : core.pushDockerImage
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.pushDockerImage.zaruba.yaml
  DESCRIPTION   : Push docker image.
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
                    {{ .Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  imageName         : Blank
                  imagePrefix       : Blank
                  imageTag          : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : {{ $d := .Decoration -}}
                                      DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
                                      DOCKER_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
                                      if [ ! -z "${DOCKER_IMAGE_TAG}" ]
                                      then
                                        docker push "${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                                      fi
                                      docker push "${DOCKER_IMAGE_NAME}:latest"
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${DOCKER_IMAGE_NAME} pushed{{ $d.Normal }}"
                  useImagePrefix    : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```