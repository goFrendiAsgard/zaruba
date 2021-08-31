# core.runDockerContainer
```
  TASK NAME     : core.runDockerContainer
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.runDockerContainer.zaruba.yaml
  DESCRIPTION   : Run docker container.
                  If container is already started, it's stdout/stderr will be shown.
                  If container is exist but not started, it will be started.
                  If container is not exist, it will be created and started.
                  Common config:
                    setup          : Script to be executed before start service or check service readiness.
                    beforeStart    : Script to be executed before start service.
                    afterStart     : Script to be executed after start service.
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
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ updateProjectLinks ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                       : set -e
                                                 {{ .Trim (.GetConfig "includeUtilScript") "\n" }} 
                                                 {{ .Trim (.GetConfig "_setupContainerName") "\n" }} 
                                                 {{ .Trim (.GetConfig "_setupImageName") "\n" }} 
                  _setupContainerName          : {{ $d := .Decoration -}}
                                                 CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                                 if [ -z "${CONTAINER_NAME}" ]
                                                 then
                                                   echo "{{ $d.Bold }}{{ $d.Red }}containerName is not provided{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _setupImageName              : {{ $d := .Decoration -}}
                                                 DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
                                                 if [ -z "${CONTAINER_NAME}" ]
                                                 then
                                                   echo "{{ $d.Bold }}{{ $d.Red }}imageName is not provided{{ $d.Normal }}"
                                                   exit 1
                                                 fi
                  _start                       : {{ $d := .Decoration -}}
                                                 {{ $rebuild := .GetConfig "rebuild" -}}
                                                 {{ if .IsTrue $rebuild }}{{ .GetConfig "_startRebuildContainer" }}{{ end }}
                                                 if [ "$(is_docker_network_exist "{{ .GetConfig "network" }}")" ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Network '{{ .GetConfig "network" }}' is already exist{{ $d.Normal }}"
                                                 else
                                                   docker network create "{{ .GetConfig "network" }}"
                                                 fi
                                                 if [ "$(inspect_docker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is already started{{ $d.Normal }}"
                                                   {{ .GetConfig "_startLogContainer" }}
                                                 elif [ ! -z $(inspect_docker "container" ".Name" "${CONTAINER_NAME}") ]
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
                  _startRebuildContainer       : stop_container "${CONTAINER_NAME}"
                                                 remove_container "${CONTAINER_NAME}"
                  _startRunContainer           : {{ $d := .Decoration -}}
                                                 {{ $imageTag := .GetConfig "imageTag" -}}
                                                 {{ $this := . -}}
                                                 docker run --name "${CONTAINER_NAME}" {{ "" -}}
                                                 --hostname "${CONTAINER_NAME}" {{ "" -}}
                                                 --network "{{ .GetConfig "network" }}" {{ "" -}}
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
                  _startRunContainerVolumes    : {{ $this := . -}}
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
                  afterStart                   : Blank
                  beforeStart                  : Blank
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
                  volumes                      : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```