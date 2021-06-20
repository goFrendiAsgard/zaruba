# core.helmDestroy
```
  TASK NAME     : core.helmDestroy
  LOCATION      : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
  DESCRIPTION   : Destroy helm deployments by using helmfile....
                  Common config:
                    helmEnv     : helm environment key (default: '{{ .GetValue "helm.env" }}')
                    dockerEnv   : docker environment key (default: '{{ .GetValue "docker.env" }}')
                    kubeContext : kubernetes context (default: '{{ .GetValue "kube.content" }}')
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                      : Blank
                  _start                      : Blank
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  dockerEnv                   : {{ .GetValue "docker.env" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helm.env" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : false
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
                  setup                       : Blank
                  start                       : {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                KUBE_CONTEXT="{{ .GetConfig "kubeContext" }}"
                                                kubectl config use-context "${KUBE_CONTEXT}"
                                                export IMAGE_PREFIX="${DOCKER_IMAGE_PREFIX}"
                                                helmfile --environment "{{ .GetConfig "helmDestroy" }}" destroy
                  useImagePrefix              : true
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
