# makeStaticServiceTask
```
  TASK NAME     : makeStaticServiceTask
  LOCATION      : ${ZARUBA_HOME}/scripts/task.makeStaticServiceTask.zaruba.yaml
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
  INPUTS        : generatorServiceLocation
                    DESCRIPTION : Service location, relative to this directory
                    PROMPT      : Service location
                    VALIDATION  : ^.+$
                  generatorServiceName
                    DESCRIPTION : Service name (Can be blank)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generatorServiceEnvs
                    DESCRIPTION : Service environments, comma separated.
                                  E.g: HTTP_PORT=3000,MODE=writer
                                  Many applications rely on environment variables to configure their behavior.
                                  You might need to see service's documentation or open environment files (.env, template.env, etc) to see available options.
                                  If there is no documentation/environment files available, you probably need to run-through the code or ask the developer team.
                    PROMPT      : Service environments
                  generatorServicePorts
                    DESCRIPTION : Service ports (number or environment variable), comma separated.
                                  E.g: 3000,HTTP_PORT,PROMETHEUS_PORT
                    PROMPT      : Service ports
                    VALIDATION  : ^[a-zA-Z0-9_,]*$
                  generatorTaskDependencies
                    DESCRIPTION : Task's dependencies, comma separated.
                                  E.g: runMysql, runRedis
                                  For example, you want to make sure that MySQL and Redis is already running before starting this task.
                                  In that case, assuming runMySql and runRedis are tasks to run MySQL and Redis respectively, then you need to set this task's dependencies into:
                                    runMysql,runRedis
                    PROMPT      : Task dependencies
                  generatorServiceDockerImageName
                    DESCRIPTION : Service's docker image name (Can be blank)
                    PROMPT      : Service's docker image name
                    VALIDATION  : ^[a-z0-9_]*$
                  generatorServiceDockerContainerName
                    DESCRIPTION : Service's docker container name (Can be blank)
                    PROMPT      : Service's docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
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
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/bash/util.sh
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
                  template               : static
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/service/{{ .GetConfig "template" }}.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```