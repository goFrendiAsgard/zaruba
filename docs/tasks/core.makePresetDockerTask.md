# core.makePresetDockerTask
```
  TASK NAME     : core.makePresetDockerTask
  LOCATION      : /home/gofrendi/zaruba/scripts/core.generator.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeDockerTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  containerName          : {{ .GetValue "generator.docker.container.name" }}
                  dependencies           : {{ .GetValue "generator.task.dependencies" }}
                  finish                 : Blank
                  imageName              : {{ .GetValue "generator.docker.image.name" }}
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  playBellScript         : echo $'\a'
                  serviceEnvs            : {{ .GetValue "generator.service.envs" }}
                  serviceName            : {{ .GetValue "generator.service.name" }}
                  setup                  : Blank
                  start                  : {{- $d := .Decoration -}}
                                           TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                           IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                           CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                           SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                           SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                           DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                           create_docker_task "template_location=${TEMPLATE_LOCATION}" "image_name=${IMAGE_NAME}" "container_name=${CONTAINER_NAME}" "service_name=${SERVICE_NAME}" "envs=${SERVICE_ENVS}" "dependencies=${DEPENDENCIES}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task created{{ $d.Normal }}"
                  template               : Blank
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/docker/{{ .GetConfig "template" }}.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
