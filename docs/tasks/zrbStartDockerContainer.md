# zrbStartDockerContainer
```
  TASK NAME     : zrbStartDockerContainer
  LOCATION      : /scripts/tasks/zrbStartDockerContainer.zaruba.yaml
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
  PARENT TASKS  : [ zrbStartApp ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
                      sleep infinity
                    {{ end -}}
                    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is started{{ $d.Normal }}"
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runInLocal") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
                      exit 0
                    {{ end -}}
                    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_check") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "check") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  CONFIG        : runInLocal                   : true
                  _check                       : {{ $d := .Decoration -}}
                                                 {{ .GetConfig "_checkContainerState" }}
                                                 {{ .GetConfig "_checkConfigPorts" }}
                                                 {{ .GetConfig "_checkCommand" }}
                                                 sleep 1
                  _checkCommand                : {{ $d := .Decoration -}}
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
                  _checkConfigPorts            : {{ $d := .Decoration -}}
                                                 {{ $this := . -}}
                                                 {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                                   {{ if ne $port "" -}}
                                                     {{ $portParts := $this.Util.Str.Split ($this.Util.Str.Trim $port  " ") ":" -}}
                                                     {{ $hostPort := index $portParts 0 -}}
                                                     echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Waiting for host port: '{{ $hostPort }}'{{ $d.Normal }}"
                                                     waitPort "localhost" {{ $hostPort }}
                                                     echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Host port '{{ $hostPort }}' is ready{{ $d.Normal }}"
                                                   {{ end -}}
                                                 {{ end -}}
                  _checkContainerState         : {{ $d := .Decoration -}}
                                                 until [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                 do
                                                   sleep 1
                                                 done
                                                 while [ "$(inspectDocker "container" ".State.Health" "${CONTAINER_NAME}")" = false ]
                                                 do
                                                   sleep 1
                                                 done
                                                 echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is running{{ $d.Normal }}"
                  _finish                      : Blank
                  _setup                       : set -e
                                                 {{ .Util.Str.Trim (.GetConfig "includeUtilScript") "\n" }} 
                                                 {{ .Util.Str.Trim (.GetConfig "_setupContainerName") "\n" }} 
                                                 {{ .Util.Str.Trim (.GetConfig "_setupImageName") "\n" }} 
                  _setupContainerName          : {{ $d := .Decoration -}}
                                                 CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                 if [ -z "${CONTAINER_NAME}" ]
                                                 then
                                                   echo "{{ $d.Bold }}{{ $d.Red }}containerName is not provided{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _setupImageName              : {{ $d := .Decoration -}}
                                                 DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
                                                 if [ -z "${DOCKER_IMAGE_NAME}" ]
                                                 then
                                                   echo "{{ $d.Bold }}{{ $d.Red }}imageName is not provided{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _start                       : {{ $d := .Decoration -}}
                                                 {{ $rebuild := .GetConfig "rebuild" -}}
                                                 {{ if .IsTrue $rebuild }}{{ .GetConfig "_startRebuildContainer" }}{{ end }}
                                                 if [ "$(inspectDocker network ".Name" "{{ .GetConfig "network" }}")" = "{{ .GetConfig "network" }}" ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Network '{{ .GetConfig "network" }}' is already exist{{ $d.Normal }}"
                                                 else
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Creating network '{{ .GetConfig "network" }}'{{ $d.Normal }}"
                                                   docker network create "{{ .GetConfig "network" }}"
                                                 fi
                                                 if [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is already started{{ $d.Normal }}"
                                                   {{ .GetConfig "_startLogContainer" }}
                                                 elif [ ! -z $(inspectDocker "container" ".Name" "${CONTAINER_NAME}") ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Retrieve previous log of '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                   sleep 1
                                                   docker logs --tail 20 "${CONTAINER_NAME}"
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Starting container '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                   docker start "${CONTAINER_NAME}"
                                                   {{ .GetConfig "_startLogContainer" }}
                                                 else
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Creating and starting container '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                   {{ .GetConfig "_startRunContainer" }}
                                                   {{ .GetConfig "_startLogContainer" }}
                                                 fi
                  _startLogContainer           : {{ $d := .Decoration -}}
                                                 echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Logging '${CONTAINER_NAME}'{{ $d.Normal }}"
                                                 docker logs --since 0m --follow "${CONTAINER_NAME}"
                  _startRebuildContainer       : stopContainer "${CONTAINER_NAME}"
                                                 removeContainer "${CONTAINER_NAME}"
                  _startRunContainer           : {{ $d := .Decoration -}}
                                                 {{ $imageTag := .GetConfig "imageTag" -}}
                                                 {{ $this := . -}}
                                                 docker run --name "${CONTAINER_NAME}" {{ "" -}}
                                                 --hostname "${CONTAINER_NAME}" {{ "" -}}
                                                 --network "{{ .GetConfig "network" }}" {{ "" -}}
                                                 {{ if .GetConfig "user" }}--user "{{ .GetConfig "user" }}" {{ end }} {{ "" -}}
                                                 {{ .GetConfig "_startRunContainerEntryPoint" -}}
                                                 {{ .GetConfig "_startRunContainerEnv" -}}
                                                 {{ .GetConfig "_startRunContainerPorts" -}}
                                                 {{ .GetConfig "_startRunContainerVolumes" -}}
                                                 {{ if ne (.GetConfig "hostDockerInternal") "host.docker.internal" }}--add-host "{{ .GetConfig "hostDockerInternal" }}:host.docker.internal"{{ end }} {{ "" -}}
                                                 -d "${DOCKER_IMAGE_NAME}{{ if $imageTag }}:{{ $imageTag }}{{ end }}" {{ .GetConfig "command" }}
                  _startRunContainerEntryPoint : {{ if .GetConfig "entryPoint" -}}
                                                   --entrypoint "{{ .GetConfig "entryPoint" }}" {{ "" -}}
                                                 {{ end -}}
                  _startRunContainerEnv        : {{ $this := . -}}
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
                  _startRunContainerPorts      : {{ $this := . -}}
                                                 {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                                   {{ if ne $port "" -}}
                                                     {{ $portParts := $this.Util.Str.Split ($this.Util.Str.Trim $port  " ") ":" -}}
                                                     {{ if eq (len $portParts) 1 -}}
                                                       -p {{ $port }}:{{ $port }} {{ "" -}}
                                                     {{ else -}}
                                                       {{ $hostPort := index $portParts 0 -}}
                                                       {{ $containerPort := index $portParts 1 -}}
                                                       -p {{ $hostPort }}:{{ $containerPort }} {{ "" -}}
                                                     {{ end -}}
                                                   {{ end -}}
                                                 {{ end -}}
                  _startRunContainerVolumes    : {{ $this := . -}}
                                                 {{ range $index, $volume := .Util.Str.Split (.Util.Str.Trim (.GetConfig "volumes") "\n ") "\n" -}}
                                                   {{ if ne $volume "" -}}
                                                     {{ $volumeParts := $this.Util.Str.Split ($this.Util.Str.Trim $volume  " ") ":" -}}
                                                     {{ if eq (len $volumeParts) 2 -}}
                                                       {{ $absHostVolume := $this.GetRelativePath (index $volumeParts 0) -}}
                                                       {{ $containerVolume := index $volumeParts 1 -}}
                                                       -v "{{ $absHostVolume }}:{{ $containerVolume }}" {{ "" -}}
                                                     {{ end -}}
                                                   {{ end -}}
                                                 {{ end -}}
                  afterCheck                   : Blank
                  afterStart                   : Blank
                  beforeCheck                  : Blank
                  beforeStart                  : Blank
                  check                        : Blank
                  checkCommand                 : Blank
                  cmd                          : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                       : -c
                  command                      : Blank
                  containerName                : Blank
                  entryPoint                   : Blank
                  finish                       : Blank
                  hostDockerInternal           : {{ if .GetEnv "ZARUBA_HOST_DOCKER_INTERNAL" }}{{ .GetEnv "ZARUBA_HOST_DOCKER_INTERNAL" }}{{ else }}host.docker.internal{{ end }}
                  imageName                    : Blank
                  imagePrefix                  : Blank
                  imageTag                     : Blank
                  includeUtilScript            : . ${ZARUBA_HOME}/bash/util.sh
                  localhost                    : localhost
                  network                      : {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}
                  ports                        : Blank
                  rebuild                      : false
                  setup                        : Blank
                  start                        : Blank
                  useImagePrefix               : true
                  user                         : Blank
                  volumes                      : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```