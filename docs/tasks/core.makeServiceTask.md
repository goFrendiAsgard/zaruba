# core.makeServiceTask
```
  TASK NAME     : core.makeServiceTask
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.makeServiceTask.zaruba.yaml
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
                    {{ .Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish                     : {{- $d := .Decoration -}}
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Service task for ${SERVICE_NAME} created{{ $d.Normal }}"
                  _setup                      : set -e
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                                                . "${ZARUBA_HOME}/bash/generatorUtil.sh"
                                                {{ if .IsTrue (.GetConfig "allowInexistServiceLocation") -}}
                                                mkdir -p "{{ .GetConfig "serviceLocation" }}"
                                                {{ end -}}
                                                TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
                                                SERVICE_LOCATION={{ .EscapeShellArg (.GetConfig "serviceLocation") }}
                                                SERVICE_NAME={{ .EscapeShellArg (.GetConfig "serviceName") }}
                                                IMAGE_NAME={{ .EscapeShellArg (.GetConfig "imageName") }}
                                                CONTAINER_NAME={{ .EscapeShellArg (.GetConfig "containerName") }}
                                                SERVICE_START_COMMAND={{ .EscapeShellArg (.GetConfig "serviceStartCommand") }}
                                                SERVICE_RUNNER_VERSION={{ .EscapeShellArg (.GetConfig "serviceRunnerVersion") }}
                                                SERVICE_PORTS={{ .EscapeShellArg (.GetConfig "servicePorts") }}
                                                SERVICE_ENVS={{ .EscapeShellArg (.GetConfig "serviceEnvs") }}
                                                DEPENDENCIES={{ .EscapeShellArg (.GetConfig "dependencies") }}
                                                REPLACEMENT_MAP={{ .EscapeShellArg (.GetConfig "replacementMap") }}
                                                # ensure SERVICE_NAME is not empty
                                                SERVICE_NAME="$(getServiceName "${SERVICE_NAME}" "${SERVICE_LOCATION}")"
                                                # ensure IMAGE_NAME is not empty
                                                IMAGE_NAME="$(getServiceImageName "${IMAGE_NAME}" "${SERVICE_NAME}")"
                                                # ensure CONTAINER_NAME is not empty
                                                CONTAINER_NAME="$(getServiceContainerName "${CONTAINER_NAME}" "${SERVICE_NAME}")"
                  _start                      : . "{{ .GetConfig "generatorScriptLocation" }}"
                                                {{ .GetConfig "generatorFunctionName" }} \
                                                {{ .GetConfig "generatorFunctionArgs" }}
                  afterStart                  : Blank
                  allowInexistServiceLocation : false
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  containerName               : {{ .GetValue "serviceContainerName" }}
                  dependencies                : {{ .GetValue "taskDependencies" }}
                  finish                      : Blank
                  generatorFunctionArgs       : "${TEMPLATE_LOCATION}" \
                                                "${SERVICE_LOCATION}" \
                                                "${SERVICE_NAME}" \
                                                "${IMAGE_NAME}" \
                                                "${CONTAINER_NAME}" \
                                                "${SERVICE_START_COMMAND}" \
                                                "${SERVICE_RUNENR_VERSION}" \
                                                "${SERVICE_PORTS}" \
                                                "${SERVICE_ENVS}" \
                                                "${DEPENDENCIES}" \
                                                "${REPLACEMENT_MAP}" \
                                                "{{ if .IsFalse (.GetConfig "registerRunner") }}0{{ else }}1{{ end }}"
                  generatorFunctionName       : generateServiceTask
                  generatorScriptLocation     : ${ZARUBA_HOME}/bash/generateServiceTask.sh
                  imageName                   : {{ .GetValue "serviceImageName" }}
                  includeUtilScript           : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner              : true
                  replacementMap              : {}
                  serviceEnvs                 : {{ .GetValue "serviceEnvs" }}
                  serviceLocation             : {{ .GetValue "serviceLocation" }}
                  serviceName                 : {{ .GetValue "serviceName" }}
                  servicePorts                : {{ .GetValue "servicePorts" }}
                  serviceRunnerVersion        : Blank
                  serviceStartCommand         : {{ .GetValue "startCommand" }}
                  setup                       : Blank
                  start                       : Blank
                  templateLocation            : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```