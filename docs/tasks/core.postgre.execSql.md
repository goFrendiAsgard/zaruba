# core.postgre.execSql
```
  TASK NAME     : core.postgre.execSql
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.postgre.execSql.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runInDockerContainer ]
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
  CONFIG        : _executeScript    : REMOTE_SCRIPT_FILE="{{ .GetConfig "_remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      USER="{{ .GetConfig "user" }}"
                                      PASSWORD="{{ .GetConfig "password" }}"
                                      docker exec "${CONTAINER_NAME}" bash -c "psql --user=${USER} -w --file=${REMOTE_SCRIPT_FILE}"
                                      docker exec -u 0 "${CONTAINER_NAME}" rm "${REMOTE_SCRIPT_FILE}"
                  _finish           : Blank
                  _generateScript   : {{ $err := .WriteFile (.GetConfig "_localScriptFile") (.GetConfig "_template") -}}
                  _localScriptFile  : {{ .GetWorkPath (printf "tmp/%s.tmp.sql" .Name) }}
                  _remoteScriptFile : /{{ .Name }}.tmp.sql
                  _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : {{ .GetConfig "_generateScript" }}
                                      {{ .GetConfig "_uploadScript" }}
                                      {{ .GetConfig "_executeScript" }}
                  _template         : {{ .GetConfig "queries" }}
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
                  database          : {{ .GetEnv "POSTGRES_DB" }}
                  finish            : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  password          : {{ .GetEnv "POSTGRES_PASSWORD" }}
                  queries           : Blank
                  setup             : Blank
                  start             : Blank
                  user              : {{ .GetEnv "POSTGRES_USER" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```