# core.postgre.startDockerContainer
```
  TASK NAME     : core.postgre.startDockerContainer
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.postgre.startDockerContainer.zaruba.yaml
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.startDockerContainer ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
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
                    {{ if .IsFalse (.GetConfig "RunInLocal") -}}
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
  CONFIG        : RunInLocal                   : true
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
                                                 {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                                   {{ if ne $port "" -}}
                                                     {{ $portParts := $this.Split ($this.Trim $port  " ") ":" -}}
                                                     {{ $hostPort := index $portParts 0 -}}
                                                     echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Waiting for host port: '{{ $hostPort }}'{{ $d.Normal }}"
                                                     wait_port "localhost" {{ $hostPort }}
                                                     echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Host port '{{ $hostPort }}' is ready{{ $d.Normal }}"
                                                   {{ end -}}
                                                 {{ end -}}
                  _checkContainerState         : {{ $d := .Decoration -}}
                                                 until [ "$(inspect_docker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                 do
                                                   sleep 1
                                                 done
                                                 while [ "$(inspect_docker "container" ".State.Health" "${CONTAINER_NAME}")" = false ]
                                                 do
                                                   sleep 1
                                                 done
                                                 echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is running{{ $d.Normal }}"
                  _setup                       : set -e
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
                                                 if [ "$(inspect_docker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
                                                 then
                                                   echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' was already started{{ $d.Normal }}"
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
                  afterCheck                   : Blank
                  afterStart                   : Blank
                  beforeCheck                  : Blank
                  beforeStart                  : Blank
                  check                        : Blank
                  checkCommand                 : pg_isready -U {{ .GetEnv "POSTGRES_USER" }}
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