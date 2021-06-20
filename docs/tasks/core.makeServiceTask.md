# core.makeServiceTask
```
  TASK NAME     : core.makeServiceTask
  LOCATION      : /home/gofrendi/zaruba/scripts/core.generator.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ core.showAdv, core.isProject ]
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
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  containerName          : {{ .GetValue "generator.service.docker.container.name" }}
                  dependencies           : {{ .GetValue "generator.task.dependencies" }}
                  finish                 : Blank
                  imageName              : {{ .GetValue "generator.service.docker.image.name" }}
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
                  serviceEnvs            : {{ .GetValue "generator.service.envs" }}
                  serviceLocation        : {{ .GetValue "generator.service.location" }}
                  serviceName            : {{ .GetValue "generator.service.name" }}
                  servicePorts           : {{ .GetValue "generator.service.ports" }}
                  serviceStartCommand    : {{ .GetValue "generator.service.startCommand" }}
                  setup                  : Blank
                  start                  : {{- $d := .Decoration -}}
                                           TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "templateLocation") }}
                                           IMAGE_NAME={{ .SingleQuoteShellValue (.GetConfig "imageName") }}
                                           CONTAINER_NAME={{ .SingleQuoteShellValue (.GetConfig "containerName") }}
                                           SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                           SERVICE_PORTS={{ .SingleQuoteShellValue (.GetConfig "servicePorts") }}
                                           SERVICE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "serviceLocation") }}
                                           SERVICE_START_COMMAND={{ .SingleQuoteShellValue (.GetConfig "serviceStartCommand") }}
                                           RUNNER_VERSION={{ .SingleQuoteShellValue (.GetConfig "runnerVersion") }}
                                           SERVICE_ENVS={{ .SingleQuoteShellValue (.GetConfig "serviceEnvs") }}
                                           DEPENDENCIES={{ .SingleQuoteShellValue (.GetConfig "dependencies") }}
                                           create_service_task "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}" "image_name=${IMAGE_NAME}" "container_name=${CONTAINER_NAME}" "location=${SERVICE_LOCATION}" "start_command=${SERVICE_START_COMMAND}" "ports=${SERVICE_PORTS}" "envs=${SERVICE_ENVS}" "dependencies=${DEPENDENCIES}" "runner_version=${RUNNER_VERSION}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Service task created{{ $d.Normal }}"
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/default.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
