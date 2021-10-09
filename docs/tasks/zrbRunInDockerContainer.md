# zrbRunInDockerContainer
```
  TASK NAME     : zrbRunInDockerContainer
  LOCATION      : /zaruba-tasks/_base/run/docker/task.zrbRunInDockerContainer.yaml
  DESCRIPTION   : Run command from inside the container
                  Common config:
                    containerName  : Name of the container
                    containerShell : Shell to run script, default to sh
                    commands       : Command to be executed, separated by new line
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbGenerateAndRun ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish          : Blank
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : {{ .GetConfig "generateScript" }}
                                     {{ .GetConfig "uploadScript" }}
                                     {{ .GetConfig "runScript" }}
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  commands         : Blank
                  containerName    : Blank
                  containerShell   : sh
                  finish           : Blank
                  generateScript   : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeShellUtil : . ${ZARUBA_HOME}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  localScriptFile  : {{ .GetWorkPath (printf "tmp/%s.tmp.sh" .Name) }}
                  remoteScriptFile : /{{ .Name }}.tmp.sh
                  runScript        : REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                     CONTAINER_SHELL="{{ .GetConfig "containerShell" }}"
                                     CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                     docker exec "${CONTAINER_NAME}" "${CONTAINER_SHELL}" "${REMOTE_SCRIPT_FILE}"
                                     docker exec -u 0 "${CONTAINER_NAME}" rm "${REMOTE_SCRIPT_FILE}"
                  scriptTemplate   : {{ .GetConfig "commands" }}
                  setup            : Blank
                  start            : Blank
                  uploadScript     : LOCAL_SCRIPT_FILE="{{ .GetConfig "localScriptFile" }}"
                                     REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                     CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                     chmod 755 "${LOCAL_SCRIPT_FILE}"
                                     docker cp "${LOCAL_SCRIPT_FILE}" "${CONTAINER_NAME}:${REMOTE_SCRIPT_FILE}"
                                     rm "${LOCAL_SCRIPT_FILE}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```