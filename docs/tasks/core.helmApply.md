# core.helmApply
```
  TASK NAME     : core.helmApply
  LOCATION      : /home/gofrendi/zaruba/scripts/task.core.helmApply.zaruba.yaml
  DESCRIPTION   : Apply helm charts by using helmfile.
                  Common config:
                    dockerEnv   : docker environment key (default: '{{ .GetValue "dockerEnv" }}')
                    helmEnv     : helm environment key (default: '{{ .GetValue "helmEnv" }}')
                    kubeContext : kubernetes context (default: '{{ .GetValue "kube.content" }}')
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
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
                  dockerEnv                   : {{ .GetValue "dockerEnv" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helmEnv" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : false
                  includeBootstrapScript      : if [ -f "${HOME}/.profile" ]
                                                then
                                                    . "${HOME}/.profile"
                                                fi
                                                if [ -f "${HOME}/.bashrc" ]
                                                then
                                                    . "${HOME}/.bashrc"
                                                fi
                                                BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                                . "${BOOTSTRAP_SCRIPT}"
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
                  kubeContext                 : {{ .GetValue "kubeContext" }}
                  setup                       : Blank
                  start                       : {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                KUBE_CONTEXT="{{ .GetConfig "kubeContext" }}"
                                                kubectl config use-context "${KUBE_CONTEXT}"
                                                export IMAGE_PREFIX="${DOCKER_IMAGE_PREFIX}"
                                                helmfile --environment "{{ .GetConfig "helmEnv" }}" sync
                  useImagePrefix              : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
