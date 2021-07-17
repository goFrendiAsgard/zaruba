# core.mysql.execSql
```
  TASK NAME     : core.mysql.execSql
  LOCATION      : ${ZARUBA_HOME}/scripts/task.core.mysql.execSql.zaruba.yaml
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
  CONFIG        : _setup                      : set -e
                                                alias zaruba=${ZARUBA_HOME}/zaruba
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                      : {{ $localTmpFile := printf "%s.tmp.sql" .NewUUIDString -}}
                                                {{ $err := .WriteFile $localTmpFile (.GetConfig "queries") -}}
                                                USER="{{ .GetConfig "user" }}"
                                                PASSWORD="{{ .GetConfig "password" }}"
                                                CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                LOCAL_TMP_FILE="{{ $localTmpFile }}"
                                                CONTAINER_TMP_FILE="/{{ .NewUUIDString }}.tmp.sql"
                                                docker cp "${LOCAL_TMP_FILE}" "${CONTAINER_NAME}:${CONTAINER_TMP_FILE}"
                                                rm "${LOCAL_TMP_FILE}"
                                                docker exec "${CONTAINER_NAME}" bash -c "mysql --user=\"${USER}\" --password=\"${PASSWORD}\" < \"${CONTAINER_TMP_FILE}\""
                                                docker exec "${CONTAINER_NAME}" rm "${CONTAINER_TMP_FILE}"
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  containerName               : Blank
                  database                    : {{ .GetEnv "MYSQL_DATABASE" }}
                  dockerEnv                   : {{ .GetValue "dockerEnv" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helmEnv" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : false
                  includeBootstrapScript      : if [ -f "${HOME}/.profile" ]
                                                then
                                                    . "${HOME}/.profile"
                                                fi
                                                if [ -f "${HOME}/.bashrc" ]
                                                then
                                                    . "${HOME}/.bashrc"
                                                fi
                                                BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                                . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript           : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  initDockerImagePrefixScript : {{ if .IsFalse (.GetConfig "useImagePrefix") -}}
                                                  DOCKER_IMAGE_PREFIX=""
                                                {{ else if .GetConfig "imagePrefix" -}}
                                                  DOCKER_IMAGE_PREFIX="{{ .GetConfig "imagePrefix" }}"
                                                {{ else if and (.GetConfig "dockerEnv") (.GetValue "dockerImagePrefix" (.GetConfig "dockerEnv")) -}}
                                                  DOCKER_IMAGE_PREFIX="{{ .GetValue "dockerImagePrefix" (.GetConfig "dockerEnv") }}"
                                                {{ else if .GetValue "dockerImagePrefix" "default" -}}
                                                  DOCKER_IMAGE_PREFIX="{{ .GetValue "dockerImagePrefix" "default" }}"
                                                {{ else -}}
                                                  DOCKER_IMAGE_PREFIX="local"
                                                {{ end -}}
                                                {{ if .IsTrue (.GetConfig "imagePrefixTrailingSlash" ) -}}
                                                  if [ ! -z "${DOCKER_IMAGE_PREFIX}" ]
                                                  then
                                                    DOCKER_IMAGE_PREFIX="${DOCKER_IMAGE_PREFIX}/"
                                                  fi
                                                {{ end -}}
                  password                    : {{ .GetEnv "MYSQL_ROOT_PASSWORD" }}
                  queries                     : Blank
                  setup                       : Blank
                  start                       : Blank
                  useImagePrefix              : true
                  user                        : root
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```