# core.mysql.execSql
```
  TASK NAME     : core.mysql.execSql
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.mysql.execSql.zaruba.yaml
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
                  database          : {{ .GetEnv "MYSQL_DATABASE" }}
                  finish            : Blank
                  generateScript    : {{ $err := .WriteFile (.GetConfig "localScriptFile") (.GetConfig "scriptTemplate") -}}
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  localScriptFile   : {{ .GetWorkPath (printf "tmp/%s.tmp.sql" .Name) }}
                  password          : {{ .GetEnv "MYSQL_ROOT_PASSWORD" }}
                  queries           : Blank
                  remoteScriptFile  : /{{ .Name }}.tmp.sql
                  runScript         : REMOTE_SCRIPT_FILE="{{ .GetConfig "remoteScriptFile" }}"
                                      CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                      USER="{{ .GetConfig "user" }}"
                                      PASSWORD="{{ .GetConfig "password" }}"
                                      docker exec "${CONTAINER_NAME}" bash -c "mysql --user=\"${USER}\" --password=\"${PASSWORD}\" < \"${REMOTE_SCRIPT_FILE}\""
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
                  user              : root
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```