# core.runInDockerContainer
```
  TASK NAME     : core.runInDockerContainer
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.runInDockerContainer.zaruba.yaml
  DESCRIPTION   : Run command from inside the container
                  Common config:
                    containerName  : Name of the container
                    containerShell : Shell to run script, default to sh
                    commands       : Command to be executed, separated by new line
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ $localTmpFile := printf "tmp/{{ .Name }}.%s.sh" .GetNewUUID -}}
                                      {{ $err := .WriteFile $localTmpFile (.GetConfig "commands") -}}
                                      CONTAINER_SHELL="{{ .GetConfig "containerShell" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      LOCAL_TMP_FILE="{{ $localTmpFile }}"
                                      CONTAINER_TMP_FILE="/{{ .Name }}.{{ .GetNewUUID }}.tmp.sh"
                                      chmod 755 "${LOCAL_TMP_FILE}"
                                      docker cp "${LOCAL_TMP_FILE}" "${CONTAINER_NAME}:${CONTAINER_TMP_FILE}"
                                      rm "${LOCAL_TMP_FILE}"
                                      docker exec "${CONTAINER_NAME}" "${CONTAINER_SHELL}" "${CONTAINER_TMP_FILE}"
                                      docker exec "${CONTAINER_NAME}" rm "${CONTAINER_TMP_FILE}"
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  commands          : Blank
                  containerName     : Blank
                  containerShell    : sh
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```