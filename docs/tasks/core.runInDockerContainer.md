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
  PARENT TASKS  : [ core.generateAndRun ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _executeScript    : REMOTE_SCRIPT_FILE="{{ .GetConfig "_remoteScriptFile" }}"
                                      CONTAINER_SHELL="{{ .GetConfig "containerShell" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      docker exec "${CONTAINER_NAME}" "${CONTAINER_SHELL}" "${REMOTE_SCRIPT_FILE}"
                                      docker exec -u 0 "${CONTAINER_NAME}" rm "${REMOTE_SCRIPT_FILE}"
                  _generateScript   : {{ $err := .WriteFile (.GetConfig "_localScriptFile") (.GetConfig "_template") -}}
                  _localScriptFile  : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  _remoteScriptFile : /{{ .Name }}.tmp.sh
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ .GetConfig "_generateScript" }}
                                      {{ .GetConfig "_uploadScript" }}
                                      {{ .GetConfig "_executeScript" }}
                  _template         : {{ .GetConfig "commands" }}
                  _uploadScript     : LOCAL_SCRIPT_FILE="{{ .GetConfig "_localScriptFile" }}"
                                      REMOTE_SCRIPT_FILE="{{ .GetConfig "_remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      chmod 755 "${LOCAL_SCRIPT_FILE}"
                                      docker cp "${LOCAL_SCRIPT_FILE}" "${CONTAINER_NAME}:${REMOTE_SCRIPT_FILE}"
                                      rm "${LOCAL_SCRIPT_FILE}"
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