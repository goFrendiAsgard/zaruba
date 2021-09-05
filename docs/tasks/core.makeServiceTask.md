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
  CONFIG        : _generatorFunctionName   : generateServiceTask
                  _generatorScriptLocation : ${ZARUBA_HOME}/bash/generateServiceTask.sh
                  _setup                   : TEMPLATE_LOCATION={{ .EscapeShellArg (.GetConfig "templateLocation") }}
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
                  _start                   : {{- $d := .Decoration -}}
                                             . "{{ .GetConfig "_generatorScriptLocation" }}"
                                             {{ .GetConfig "_generatorFunctionName" }} \
                                               "${TEMPLATE_LOCATION}" \
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
                                             echo 🎉🎉🎉
                                             echo "{{ $d.Bold }}{{ $d.Yellow }}Service task created{{ $d.Normal }}"
                  afterStart               : Blank
                  beforeStart              : Blank
                  cmd                      : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                   : -c
                  containerName            : {{ .GetValue "serviceContainerName" }}
                  dependencies             : {{ .GetValue "taskDependencies" }}
                  finish                   : Blank
                  imageName                : {{ .GetValue "serviceImageName" }}
                  includeUtilScript        : . ${ZARUBA_HOME}/bash/util.sh
                  registerRunner           : true
                  replacementMap           : {}
                  serviceEnvs              : {{ .GetValue "serviceEnvs" }}
                  serviceLocation          : {{ .GetValue "serviceLocation" }}
                  serviceName              : {{ .GetValue "serviceName" }}
                  servicePorts             : {{ .GetValue "servicePorts" }}
                  serviceRunnerVersion     : Blank
                  serviceStartCommand      : {{ .GetValue "startCommand" }}
                  setup                    : Blank
                  start                    : Blank
                  templateLocation         : {{ .GetEnv "ZARUBA_HOME" }}/templates/task/service/default
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```