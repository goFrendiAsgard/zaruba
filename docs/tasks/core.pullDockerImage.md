# core.pullDockerImage
```
  TASK NAME     : core.pullDockerImage
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.pullDockerImage.zaruba.yaml
  DESCRIPTION   : Pull docker image.
                  Common config:
                    imageName : Image name
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
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
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
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
                                        pullImage "${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                                      else
                                        pullImage "${DOCKER_IMAGE_NAME}"
                                      fi
                                      echo 🎉🎉🎉
                                      echo "{{ $d.Bold }}{{ $d.Yellow }}Docker image ${DOCKER_IMAGE_NAME} pulled{{ $d.Normal }}"
                  useImagePrefix    : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```