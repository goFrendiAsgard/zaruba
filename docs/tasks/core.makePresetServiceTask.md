# core.makePresetServiceTask
```
  TASK NAME     : core.makePresetServiceTask
  LOCATION      : /home/gofrendi/zaruba/scripts/core.generator.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makeServiceTask ]
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
                  containerName          : {{ .GetValue "generatorServiceDockerContainerName" }}
                  dependencies           : {{ .GetValue "generatorTaskDependencies" }}
                  finish                 : Blank
                  imageName              : {{ .GetValue "generatorServiceDockerImageName" }}
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
                  runnerVersion          : Blank
                  serviceEnvs            : {{ .GetValue "generatorServiceEnvs" }}
                  serviceLocation        : {{ .GetValue "generatorServiceLocation" }}
                  serviceName            : {{ .GetValue "generatorServiceName" }}
                  servicePorts           : {{ .GetValue "generatorServicePorts" }}
                  serviceStartCommand    : {{ .GetValue "generatorServiceStartCommand" }}
                  setup                  : Blank
                  start                  : {{- $d := .Decoration -}}
                                           TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                           IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                           CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                           SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                           SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                           SERVICE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceLocation") }}
                                           SERVICE_START_COMMAND={{ .EscapeShellArg (.GetConfig "serviceStartCommand") }}
                                           RUNNER_VERSION={{ .EscapeShellArg (.GetConfig "runnerVersion") }}
                                           SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                           DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                           create_service_task "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "image_name=${IMAGE_NAME}" "container_name=${CONTAINER_NAME}" "location=${SERVICE_LOCATION}" "start_command=${SERVICE_START_COMMAND}" "ports=${SERVICE_PORTS}" "envs=${SERVICE_ENVS}" "dependencies=${DEPENDENCIES}" "runner_version=${RUNNER_VERSION}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Service task created{{ $d.Normal }}"
                  template               : Blank
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/{{ .GetConfig "template" }}.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
