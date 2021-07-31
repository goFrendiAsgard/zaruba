# core.cassandra.execCql
```
  TASK NAME     : core.cassandra.execCql
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.cassandra.execCql.zaruba.yaml
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
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                      : {{ $localTmpFile := printf "%s.tmp.sql" .GetNewUUID -}}
                                                {{ $err := .WriteFile $localTmpFile (.GetConfig "queries") -}}
                                                USER="{{ .GetConfig "user" }}"
                                                PASSWORD="{{ .GetConfig "password" }}"
                                                CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                LOCAL_TMP_FILE="{{ $localTmpFile }}"
                                                CONTAINER_TMP_FILE="/{{ .GetNewUUID }}.tmp.sql"
                                                docker cp "${LOCAL_TMP_FILE}" "${CONTAINER_NAME}:${CONTAINER_TMP_FILE}"
                                                rm "${LOCAL_TMP_FILE}"
                                                docker exec "${CONTAINER_NAME}" cqlsh -u "${USER}" -p "${PASSWORD}" -f "${CONTAINER_TMP_FILE}"
                                                docker exec "${CONTAINER_NAME}" rm "${CONTAINER_TMP_FILE}"
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  containerName               : Blank
                  dockerEnv                   : {{ .GetValue "dockerEnv" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helmEnv" }}
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : false
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
                  keyspace                    : sample
                  password                    : cassandra
                  queries                     : Blank
                  setup                       : Blank
                  start                       : Blank
                  useImagePrefix              : true
                  user                        : cassandra
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```