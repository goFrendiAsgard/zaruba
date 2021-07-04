# core.startDockerContainer
```
  TASK NAME     : core.startDockerContainer
  LOCATION      : /home/gofrendi/zaruba/scripts/core.service.zaruba.yaml
  DESCRIPTION   : Start docker container and check it's readiness.
                  If container is already started, it's stdout/stderr will be shown.
                  If container is exist but not started, it will be started.
                  If container is not exist, it will be created and started.
                  Common config:
                    setup          : Script to be executed before start service or check service readiness.
                    beforeStart    : Script to be executed before start service.
                    afterStart     : Script to be executed after start service.
                    beforeCheck    : Script to be executed before check service readiness.
                    afterCheck     : Script to be executed before check service readiness.
                    finish         : Script to be executed after start service or check service readiness.
                    useImagePrefix : Whether image prefix should be used or not
                    imagePrefix    : Image prefix
                    imageName      : Image name
                    imageTag       : Image tag
                    containerName  : Name of the container
                    dockerEnv      : Docker env to be used when useImagePrefix is true,
                                     but imagePrefix is not provided
                    ports          : Port to be checked to confirm service readiness, 
                                     separated by new line.
                    volumes        : Host-container volume mappings,
                                     separated by new line.
                    rebuild        : Should container be rebuild (This will not rebuild the image)
                    command        : Command to be used (Single Line).
                                     Leave blank to use container's CMD.
                                     The command will be executed from inside the container.
                    checkCommand   : Command to check container readiness (Single Line).
                                     The command will be executed from inside the container.
                    localhost      : Localhost mapping (e.g: host.docker.container)
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.startService ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runLocally") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
                      sleep infinity
                    {{ end -}}
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runLocally") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
                      exit 0
                    {{ end -}}
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Trim (.GetConfig "_check") "\n " }}
                    {{ .Trim (.GetConfig "check") "\n " }}
                    {{ .Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  INPUTS        : docker.env
                    DESCRIPTION : Docker env for getting image prefix (Required)
                    PROMPT      : Docker env
                    DEFAULT     : default
                    VALIDATION  : ^.+$
  CONFIG        : _check                      : {{ $d := .Decoration -}}
                                                {{ .GetConfig "_check.containerState" }}
                                                {{ .GetConfig "_check.configPorts" }}
                                                {{ .GetConfig "_check.checkCommand" }}
                                                sleep 1
                  _check.checkCommand         : {{ $d := .Decoration -}}
                                                {{ if .GetConfig "checkCommand" -}}
                                                (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
                                                set +e
                                                sleep 3
                                                docker exec "${CONTAINER_NAME}" {{ .GetConfig "checkCommand" }}
                                                until [ "$?" = "0" ]
                                                do
                                                  sleep 3
                                                  docker exec "${CONTAINER_NAME}" {{ .GetConfig "checkCommand" }}
                                                done
                                                set "${_OLD_STATE}"
                                                {{ end -}}
                  _check.configPorts          : {{ $d := .Decoration -}}
                                                {{ $this := . -}}
                                                {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                                  {{ if ne $port "" -}}
                                                    {{ $portParts := $this.Split ($this.Trim $port  " ") ":" -}}
                                                    {{ $hostPort := index $portParts 0 -}}
                                                    echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Waiting for host port: '{{ $hostPort }}'{{ $d.Normal }}"
                                                    wait_port "localhost" {{ $hostPort }}
                                                    echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Host port '{{ $hostPort }}' is ready{{ $d.Normal }}"
                                                  {{ end -}}
                                                {{ end -}}
                  _check.containerState       : {{ $d := .Decoration -}}
                                                until [ "$(inspect_docker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                do
                                                  sleep 1
                                                done
                                                while [ "$(inspect_docker "container" ".State.Health" "${CONTAINER_NAME}")" = false ]
                                                do
                                                  sleep 1
                                                done
                                                echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is running{{ $d.Normal }}"
                  _setup                      : set -e
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }} 
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }} 
                                                {{ .Trim (.GetConfig "initDockerImagePrefixScript") "\n" }}
                                                {{ .Trim (.GetConfig "_setup.containerName") "\n" }} 
                                                {{ .Trim (.GetConfig "_setup.imageName") "\n" }} 
                  _setup.containerName        : {{ $d := .Decoration -}}
                                                CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                should_not_be_empty "${CONTAINER_NAME}" "{{ $d.Bold }}{{ $d.Red }}containerName is not provided{{ $d.Normal }}"
                  _setup.imageName            : {{ $d := .Decoration -}}
                                                IMAGE_NAME="{{ .GetConfig "imageName" }}"
                                                should_not_be_empty "${IMAGE_NAME}" "{{ $d.Bold }}{{ $d.Red }}imageName is not provided{{ $d.Normal }}"
                  _start                      : {{ $d := .Decoration -}}
                                                {{ $rebuild := .GetConfig "rebuild" -}}
                                                {{ if .IsTrue $rebuild }}{{ .GetConfig "_start.rebuildContainer" }}{{ end }}
                                                if [ "$(inspect_docker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                then
                                                  echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' was already started{{ $d.Normal }}"
                                                  {{ .GetConfig "_start.logContainer" }}
                                                elif [ ! -z $(inspect_docker "container" ".Name" "${CONTAINER_NAME}") ]
                                                then
                                                  echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Retrieve previous log of '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                  sleep 1
                                                  docker logs --tail 20 "${CONTAINER_NAME}"
                                                  echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Starting container '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                  docker start "${CONTAINER_NAME}"
                                                  {{ .GetConfig "_start.logContainer" }}
                                                else
                                                  echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Creating and starting container '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                  {{ .GetConfig "_start.runContainer" }}
                                                  {{ .GetConfig "_start.logContainer" }}
                                                fi
                  _start.logContainer         : {{ $d := .Decoration -}}
                                                echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Logging '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                docker logs --since 0m --follow "${CONTAINER_NAME}"
                  _start.rebuildContainer     : stop_container "${CONTAINER_NAME}"
                                                remove_container "${CONTAINER_NAME}"
                  _start.runContainer         : {{ $d := .Decoration -}}
                                                {{ $imageTag := .GetConfig "imageTag" -}}
                                                {{ $this := . -}}
                                                docker run --name "${CONTAINER_NAME}" {{ "" -}}
                                                {{ .GetConfig "_start.runContainer.env" -}}
                                                {{ .GetConfig "_start.runContainer.ports" -}}
                                                {{ .GetConfig "_start.runContainer.volumes" -}}
                                                {{ if ne (.GetConfig "hostDockerInternal") "host.docker.internal" }}--add-host "{{ .GetConfig "hostDockerInternal" }}:host.docker.internal"{{ end }} {{ "" -}}
                                                -d "${DOCKER_IMAGE_PREFIX}${IMAGE_NAME}{{ if $imageTag }}:{{ $imageTag }}{{ end }}" {{ .GetConfig "command" }}
                  _start.runContainer.env     : {{ $this := . -}}
                                                {{ if eq (.GetConfig "localhost") "localhost" -}}
                                                  {{ range $key, $val := $this.GetEnvs -}}
                                                    -e "{{ $key}}={{ $val }}" {{ "" -}}
                                                  {{ end -}}
                                                {{ else -}}
                                                  {{ range $key, $val := $this.GetEnvs -}}
                                                    {{ $val = $this.ReplaceAll $val "localhost" ($this.GetConfig "localhost") -}}
                                                    {{ $val = $this.ReplaceAll $val "127.0.0.1" ($this.GetConfig "localhost") -}}
                                                    {{ $val = $this.ReplaceAll $val "0.0.0.0" ($this.GetConfig "localhost") -}}
                                                    -e "{{ $key}}={{ $val }}" {{ "" -}}
                                                  {{ end -}}
                                                {{ end -}}
                  _start.runContainer.ports   : {{ $this := . -}}
                                                {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                                  {{ if ne $port "" -}}
                                                    {{ $portParts := $this.Split ($this.Trim $port  " ") ":" -}}
                                                    {{ if eq (len $portParts) 1 -}}
                                                      -p {{ $port }}:{{ $port }} {{ "" -}}
                                                    {{ else -}}
                                                      {{ $hostPort := index $portParts 0 -}}
                                                      {{ $containerPort := index $portParts 1 -}}
                                                      -p {{ $hostPort }}:{{ $containerPort }} {{ "" -}}
                                                    {{ end -}}
                                                  {{ end -}}
                                                {{ end -}}
                  _start.runContainer.volumes : {{ $this := . -}}
                                                {{ range $index, $volume := .Split (.Trim (.GetConfig "volumes") "\n ") "\n" -}}
                                                  {{ if ne $volume "" -}}
                                                    {{ $volumeParts := $this.Split ($this.Trim $volume  " ") ":" -}}
                                                    {{ if eq (len $volumeParts) 2 -}}
                                                      {{ $absHostVolume := $this.GetRelativePath (index $volumeParts 0) -}}
                                                      {{ $containerVolume := index $volumeParts 1 -}}
                                                      -v "{{ $absHostVolume }}:{{ $containerVolume }}" {{ "" -}}
                                                    {{ end -}}
                                                  {{ end -}}
                                                {{ end -}}
                  afterCheck                  : Blank
                  afterStart                  : Blank
                  beforeCheck                 : Blank
                  beforeStart                 : Blank
                  check                       : Blank
                  checkCommand                : Blank
                  cmd                         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                      : -c
                  command                     : Blank
                  containerName               : Blank
                  dockerEnv                   : {{ .GetValue "docker.env" }}
                  finish                      : Blank
                  helmEnv                     : {{ .GetValue "helm.env" }}
                  hostDockerInternal          : {{ if .GetEnv "ZARUBA_HOST_DOCKER_INTERNAL" }}{{ .GetEnv "ZARUBA_HOST_DOCKER_INTERNAL" }}{{ else }}host.docker.internal{{ end }}
                  imageName                   : Blank
                  imagePrefix                 : Blank
                  imagePrefixTrailingSlash    : true
                  imageTag                    : Blank
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
                  localhost                   : localhost
                  playBellScript              : echo $'\a'
                  ports                       : Blank
                  rebuild                     : false
                  runLocally                  : true
                  setup                       : Blank
                  start                       : Blank
                  useImagePrefix              : true
                  volumes                     : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
