# makeGoServiceTask
```
  TASK NAME     : makeGoServiceTask
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.generator.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makePresetServiceTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : generator.service.location
                    DESCRIPTION : Service location, relative to this directory
                    PROMPT      : Service location (Relative to this directory)
                    VALIDATION  : ^.+$
                  generator.service.name
                    DESCRIPTION : Service name
                    PROMPT      : Service name (Can be blank)
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generator.service.envs
                    DESCRIPTION : Service environments, comma separated
                    PROMPT      : Service environments (e.g: 'PORT=3000,WRITE=1')
                  generator.service.ports
                    DESCRIPTION : Service ports, comma separated
                    PROMPT      : Service ports (e.g: '3000,PORT')
                    VALIDATION  : ^[a-zA-Z0-9_,]*$
                  generator.goService.startCommand
                    DESCRIPTION : Service start command
                    PROMPT      : Command to run the service (Required)
                    DEFAULT     : go run .
                    VALIDATION  : ^.+$
                  generator.service.docker.image.name
                    DESCRIPTION : Service's docker image name
                    PROMPT      : Service's Docker image name (Can be blank)
                    VALIDATION  : ^[a-z0-9_]*$
                  generator.service.docker.container.name
                    DESCRIPTION : Service's docker container name
                    PROMPT      : Service's docker container name (Can be blank)
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generator.task.dependencies
                    DESCRIPTION : Task dependencies
                    PROMPT      : Task dependencies (e.g: 'runMySql,runCassandra')
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
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/util.sh
                  playBellScript         : echo $'\a'
                  runnerVersion          : Blank
                  serviceEnvs            : {{ .GetValue "generator.service.envs" }}
                  serviceLocation        : {{ .GetValue "generator.service.location" }}
                  serviceName            : {{ .GetValue "generator.service.name" }}
                  servicePorts           : {{ .GetValue "generator.service.ports" }}
                  serviceStartCommand    : {{ .GetValue "generator.goService.startCommand" }}
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
                  template               : go
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/{{ .GetConfig "template" }}.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
