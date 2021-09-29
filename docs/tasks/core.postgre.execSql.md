# core.postgre.execSql
```
  TASK NAME     : core.postgre.execSql
  LOCATION      : /scripts/tasks/core.postgre.execSql.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runInDockerContainer ]
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
  CONFIG        : _finish           : Blank
                  _setup            : set -e
                                      {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }}
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
                  database          : {{ .GetEnv "POSTGRES_DB" }}
                  finish            : Blank
                  generateScript    : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  localScriptFile   : {{ .GetWorkPath (printf "tmp/%s.tmp.sql" .Name) }}
                  password          : {{ .GetEnv "POSTGRES_PASSWORD" }}
                  queries           : Blank
                  remoteScriptFile  : /{{ .Name }}.tmp.sql
                  runScript         : REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      USER="{{ .GetConfig "user" }}"
                                      PASSWORD="{{ .GetConfig "password" }}"
                                      docker exec "${CONTAINER_NAME}" bash -c "psql --user=${USER} -w --file=${REMOTE_SCRIPT_FILE}"
                                      docker exec -u 0 "${CONTAINER_NAME}" rm "${REMOTE_SCRIPT_FILE}"
                  scriptTemplate    : {{ .GetConfig "queries" }}
                  setup             : Blank
                  start             : Blank
                  uploadScript      : LOCAL_SCRIPT_FILE="{{ .GetConfig "localScriptFile" }}"
                                      REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      chmod 755 "${LOCAL_SCRIPT_FILE}"
                                      docker cp "${LOCAL_SCRIPT_FILE}" "${CONTAINER_NAME}:${REMOTE_SCRIPT_FILE}"
                                      rm "${LOCAL_SCRIPT_FILE}"
                  user              : {{ .GetEnv "POSTGRES_USER" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```