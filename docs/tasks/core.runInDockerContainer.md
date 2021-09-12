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
                    {{ .Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ .GetConfig "generateScript" }}
                                      {{ .GetConfig "uploadScript" }}
                                      {{ .GetConfig "runScript" }}
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  commands          : Blank
                  containerName     : Blank
                  containerShell    : sh
                  finish            : Blank
                  generateScript    : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  localScriptFile   : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  remoteScriptFile  : /{{ .Name }}.tmp.sh
                  runScript         : REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                      CONTAINER_SHELL="{{ .GetConfig "containerShell" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      docker exec "${CONTAINER_NAME}" "${CONTAINER_SHELL}" "${REMOTE_SCRIPT_FILE}"
                                      docker exec -u 0 "${CONTAINER_NAME}" rm "${REMOTE_SCRIPT_FILE}"
                  scriptTemplate    : {{ .GetConfig "commands" }}
                  setup             : Blank
                  start             : Blank
                  uploadScript      : LOCAL_SCRIPT_FILE="{{ .GetConfig "localScriptFile" }}"
                                      REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      chmod 755 "${LOCAL_SCRIPT_FILE}"
                                      docker cp "${LOCAL_SCRIPT_FILE}" "${CONTAINER_NAME}:${REMOTE_SCRIPT_FILE}"
                                      rm "${LOCAL_SCRIPT_FILE}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```