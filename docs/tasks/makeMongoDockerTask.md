# makeMongoDockerTask
```
  TASK NAME     : makeMongoDockerTask
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.generator.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.makePresetDockerTask ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : generator.docker.container.name
                    DESCRIPTION : Docker container name (Can be blank)
                    PROMPT      : Docker container name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generator.service.name
                    DESCRIPTION : Service name (Can be blank)
                    PROMPT      : Service name
                    VALIDATION  : ^[a-zA-Z0-9_]*$
                  generator.service.envs
                    DESCRIPTION : Service environments, comma separated.
                                  E.g: HTTP_PORT=3000,MODE=writer
                                  
                                  Many applications rely on environment variables to configure their behavior.
                                  You might need to see service's documentation or open environment files (.env, template.env, etc) to see available options.
                                  If there is no documentation/environment files available, you probably need to run-through the code or ask the developer team.
                    PROMPT      : Service environments
                  generator.task.dependencies
                    DESCRIPTION : Task's dependencies, comma separated.
                                  E.g: runMysql, runRedis
                                  
                                  For example, you want to make sure that MySQL and Redis is already running before starting this task.
                                  
                                  In that case, assuming runMySql and runRedis are tasks to run MySQL and Redis respectively, then you need to set this task's dependencies into:
                                    runMysql,runRedis
                    PROMPT      : Task dependencies
  CONFIG        : _setup                 : set -e
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
                                           TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "templateLocation") }}
                                           IMAGE_NAME={{ .SingleQuoteShellValue (.GetConfig "imageName") }}
                                           CONTAINER_NAME={{ .SingleQuoteShellValue (.GetConfig "containerName") }}
                                           SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                           SERVICE_ENVS={{ .SingleQuoteShellValue (.GetConfig "serviceEnvs") }}
                                           DEPENDENCIES={{ .SingleQuoteShellValue (.GetConfig "dependencies") }}
                                           create_docker_task "template_location=${TEMPLATE_LOCATION}" "image_name=${IMAGE_NAME}" "container_name=${CONTAINER_NAME}" "service_name=${SERVICE_NAME}" "envs=${SERVICE_ENVS}" "dependencies=${DEPENDENCIES}"
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Docker task created{{ $d.Normal }}"
                  template               : mongo
                  templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/docker/{{ .GetConfig "template" }}.zaruba.yaml
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
