# core.postgre.execSql
```
  TASK NAME     : core.postgre.execSql
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.postgre.execSql.zaruba.yaml
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
                  _start            : {{ $localTmpFile := printf "tmp/{{ .Name }}.%s.sql" .GetNewUUID -}}
                                      {{ $err := .WriteFile $localTmpFile (.GetConfig "queries") -}}
                                      USER="{{ .GetConfig "user" }}"
                                      PASSWORD="{{ .GetConfig "password" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      LOCAL_TMP_FILE="{{ $localTmpFile }}"
                                      CONTAINER_TMP_FILE="/{{ .Name }}.{{ .GetNewUUID }}.tmp.sql"
                                      docker cp "${LOCAL_TMP_FILE}" "${CONTAINER_NAME}:${CONTAINER_TMP_FILE}"
                                      rm "${LOCAL_TMP_FILE}"
                                      docker exec "${CONTAINER_NAME}" bash -c "psql --user=${USER} -w --file=${CONTAINER_TMP_FILE}"
                                      docker exec "${CONTAINER_NAME}" rm "${CONTAINER_TMP_FILE}"
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  containerName     : Blank
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