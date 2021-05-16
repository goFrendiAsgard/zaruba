# makeFastApiService
```
      TASK NAME    : makeFastApiService
      LOCATION     : /home/gofrendi/.zaruba/scripts/core.generator.zaruba.yaml
      DESCRIPTION  : Make Fast API service
      TASK TYPE    : Command Task
      PARENT TASKS : [ core.runCoreScript ]
      DEPENDENCIES : [ core.showAdv ]
      START        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{ .Trim (.GetConfig "_setup") "\n " }}
                       {{ .Trim (.GetConfig "setup") "\n " }}
                       {{ .Trim (.GetConfig "beforeStart") "\n " }}
                       {{ .Trim (.GetConfig "_start") "\n " }}
                       {{ .Trim (.GetConfig "start") "\n " }}
                       {{ .Trim (.GetConfig "afterStart") "\n " }}
                       {{ .Trim (.GetConfig "finish") "\n " }}
      INPUTS       : generator.fastApi.service.name
                       DESCRIPTION : Service name
                       PROMPT      : Service name (Required)
                       VALIDATION  : ^[a-zA-Z0-9_]+$
      CONFIG       :   _setup                 : set -e
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                       _start                 : Blank
                       afterStart             : Blank
                       beforeStart            : Blank
                       cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                       cmdArg                 : -c
                       finish                 : Blank
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
                       serviceName            : {{ .GetValue "generator.fastApi.service.name" }}
                       setup                  : Blank
                       start                  : {{- $d := .Decoration -}}
                                                create_link()
                                                {
                                                  add_link "${1}" "${2}"
                                                  link_resource "${1}" "${2}"
                                                }
                                                TEMPLATE_LOCATION={{ .SingleQuoteShellValue (.GetConfig "templateLocation") }}
                                                SERVICE_NAME={{ .SingleQuoteShellValue (.GetConfig "serviceName") }}
                                                create_fast_service "template_location=${TEMPLATE_LOCATION}" "service_name=${SERVICE_NAME}"
                                                if [ -f "./main.zaruba.yaml" ]
                                                then
                                                  if [ ! -d "./shared-libs/python/helpers" ]
                                                  then
                                                    mkdir -p "./shared-libs/python/helpers"
                                                    cp -rnT "./${SERVICE_NAME}/helpers" "./shared-libs/python/helpers"
                                                  fi
                                                  create_link "shared-libs/python/helpers" "${SERVICE_NAME}/helpers"
                                                fi
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Fast API service created: ${SERVICE_NAME}{{ $d.Normal }}"
                       templateLocation       : {{ .GetEnv "ZARUBA_HOME" }}/scripts/templates/fastApiService
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
