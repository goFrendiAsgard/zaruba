# makeDockerTask
```
      TASK NAME    : makeDockerTask
      LOCATION     : /home/gofrendi/zaruba/scripts/core.generator.zaruba.yaml
      TASK TYPE    : Command Task
      PARENT TASKS : [ core.makeDockerTask ]
      START        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{ .Trim (.GetConfig "_setup") "\n " }}
                       {{ .Trim (.GetConfig "setup") "\n " }}
                       {{ .Trim (.GetConfig "beforeStart") "\n " }}
                       {{ .Trim (.GetConfig "_start") "\n " }}
                       {{ .Trim (.GetConfig "start") "\n " }}
                       {{ .Trim (.GetConfig "afterStart") "\n " }}
                       {{ .Trim (.GetConfig "finish") "\n " }}
      INPUTS       : generator.docker.image.name
                       DESCRIPTION : Docker image name
                       PROMPT      : Docker image name (Required)
                       VALIDATION  : ^[a-z0-9_]+$
                     generator.docker.container.name
                       DESCRIPTION : Docker container name
                       PROMPT      : Docker container name (Can be blank)
                       VALIDATION  : ^[a-zA-Z0-9_]*$
                     generator.service.name
                       DESCRIPTION : Service name
                       PROMPT      : Service name (Can be blank)
                       VALIDATION  : ^[a-zA-Z0-9_]*$
                     generator.service.envs
                       DESCRIPTION : Service environments, comma separated
                       PROMPT      : Service environments (e.g: 'PORT=3000,WRITE=1')
                     generator.task.dependencies
                       DESCRIPTION : Task dependencies
                       PROMPT      : Task dependencies (e.g: 'runMySql,runCassandra')
      CONFIG       :   _setup                 : set -e
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
                       includeUtilScript      : . ${ZARUBA_HOME}/scripts/util.sh
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
                       templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/task/docker/default.zaruba.yaml
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
