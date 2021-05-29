# core.mysql.execSql
```
  TASK NAME     : core.mysql.execSql
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.run.zaruba.yaml
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
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                      : USER="{{ .GetConfig "user" }}"
                                                PASSWORD="{{ .GetConfig "password" }}"
                                                CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                {{ range $index, $query := .GetLConfig "queries" -}}
                                                QUERY="{{ $query }}"
                                                if [ -f "${QUERY}" ]
                                                then
                                                  echo "SQL FILE: ${QUERY}"
                                                  TMP_FILE_NAME="/$(get_uuid4).sql"
                                                  docker cp "${QUERY}" "${CONTAINER_NAME}:${TMP_FILE_NAME}"
                                                  docker exec "${CONTAINER_NAME}" bash -c "mysql --user=\"${USER}\" --password=\"${PASSWORD}\" < \"${TMP_FILE_NAME}\""
                                                  docker exec "${CONTAINER_NAME}" rm "${TMP_FILE_NAME}"
                                                else
                                                  echo "SQL SCRIPT: ${QUERY}"
                                                  docker exec "${CONTAINER_NAME}" mysql --user="${USER}" --password="${PASSWORD}" -e "${QUERY}"
                                                fi
                                                {{ end -}}
                  afterStart                  : Blank
                  beforeStart                 : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  containerName               : Blank
                  database                    : {{ .GetEnv "MYSQL_DATABASE" }}
                  dockerEnv                   : {{ .GetValue "docker.env" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helm.env" }}
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
                                                BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                                . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript           : . "${ZARUBA_HOME}/scripts/util.sh"
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
                  kubeContext                 : {{ .GetValue "kube.context" }}
                  password                    : {{ .GetEnv "MYSQL_ROOT_PASSWORD" }}
                  playBellScript              : echo $'\a'
                  setup                       : Blank
                  start                       : Blank
                  useImagePrefix              : true
                  user                        : root
  LCONFIG       : queries : []
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
