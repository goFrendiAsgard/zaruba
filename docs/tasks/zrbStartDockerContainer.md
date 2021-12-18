
# ZrbStartDockerContainer

File Location:

    /zaruba-tasks/_base/start/task.zrbStartDockerContainer.yaml

Should Sync Env:

    true

Type:

    service

Description:

    Start docker container and wait until it is ready.
    If container is already started, it's stdout/stderr will be shown.
    If container is exist but not started, it will be started.
    If container is not exist, it will be created and started.
    Common configs:
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



## Extends

* `zrbStartApp`


## Dependencies

* `updateProjectLinks`
* `zrbCreateDockerNetwork`


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"
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
    echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is started${_NORMAL}"

    ```


## Check

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
      echo 🎉🎉🎉
      echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
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
    echo "📜 ${_BOLD}${_YELLOW}Task '{{ .Name }}' is ready${_NORMAL}"
    ```


## Configs


### Configs._check

Value:

    {{ .GetConfig "_checkContainerState" }}
    {{ .GetConfig "_checkConfigPorts" }}
    {{ .GetConfig "_checkCommand" }}
    sleep 1



### Configs._checkCommand

Value:

    {{ if .GetConfig "checkCommand" -}}
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    sleep 3
    {{ $checkCommand := .Util.Str.Trim (.GetConfig "checkCommand") "\n" -}}
    echo "🔎 ${_BOLD}${_YELLOW}Run check in '${CONTAINER_NAME}': {{ .Util.Str.EscapeShellValue $checkCommand }}${_NORMAL}"
    docker exec "${CONTAINER_NAME}" {{ $checkCommand }}
    until [ "$?" = "0" ]
    do
      sleep 3
      docker exec "${CONTAINER_NAME}" {{ $checkCommand }}
    done
    set "${_OLD_STATE}"
    {{ end -}}



### Configs._checkConfigPorts

Value:

    {{ $this := . -}}
    {{ range $index, $port := .Util.Str.Split (.Util.Str.Trim (.GetConfig "ports") "\n ") "\n" -}}
      {{ if ne $port "" -}}
        {{ $portParts := $this.Util.Str.Split ($this.Util.Str.Trim $port  " ") ":" -}}
        {{ $hostPort := index $portParts 0 -}}
        echo "🔎 ${_BOLD}${_YELLOW}Waiting for host port: '{{ $hostPort }}'${_NORMAL}"
        waitPort "localhost" {{ $hostPort }}
        echo "🔎 ${_BOLD}${_YELLOW}Host port '{{ $hostPort }}' is ready${_NORMAL}"
      {{ end -}}
    {{ end -}}



### Configs._checkContainerState

Value:

    echo "🔎 ${_BOLD}${_YELLOW}Waiting docker container '${CONTAINER_NAME}' running status${_NORMAL}"
    until [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
    do
      sleep 1
    done
    echo "🔎 ${_BOLD}${_YELLOW}Waiting docker container '${CONTAINER_NAME}' healthcheck${_NORMAL}"
    while [ "$(inspectDocker "container" ".State.Health" "${CONTAINER_NAME}")" = false ]
    do
      sleep 1
    done
    echo "🔎 ${_BOLD}${_YELLOW}Docker container '${CONTAINER_NAME}' is running${_NORMAL}"



### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupContainerName") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupImageName") "\n" }} 



### Configs._setupContainerName

Value:

    CONTAINER_NAME="{{ .GetConfig "containerName" }}"
    if [ -z "${CONTAINER_NAME}" ]
    then
      echo "${_BOLD}${_RED}containerName is not provided${_NORMAL}"
      exit 1
    fi



### Configs._setupImageName

Value:

    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    if [ -z "${DOCKER_IMAGE_NAME}" ]
    then
      echo "${_BOLD}${_RED}imageName is not provided${_NORMAL}"
      exit 1
    fi



### Configs._start

Value:

    {{ $rebuild := .GetConfig "rebuild" -}}
    {{ if .Util.Bool.IsTrue $rebuild }}{{ .GetConfig "_startRebuildContainer" }}{{ end }}
    if [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
    then
      echo "🐳 ${_BOLD}${_YELLOW}Container '${CONTAINER_NAME}' is already started${_NORMAL}"
      {{ .GetConfig "_startLogContainer" }}
    elif [ ! -z $(inspectDocker "container" ".Name" "${CONTAINER_NAME}") ]
    then
      echo "🐳 ${_BOLD}${_YELLOW}Retrieve previous log of '${CONTAINER_NAME}'${_NORMAL}"
      sleep 1
      docker logs --tail 20 "${CONTAINER_NAME}"
      echo "🐳 ${_BOLD}${_YELLOW}Starting container '${CONTAINER_NAME}'${_NORMAL}"
      docker start "${CONTAINER_NAME}"
      {{ .GetConfig "_startLogContainer" }}
    else
      echo "🐳 ${_BOLD}${_YELLOW}Creating and starting container '${CONTAINER_NAME}'${_NORMAL}"
      {{ .GetConfig "_startRunContainer" }}
      {{ .GetConfig "_startLogContainer" }}
    fi



### Configs._startLogContainer

Value:

    echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
    docker logs --since 0m --follow "${CONTAINER_NAME}"



### Configs._startRebuildContainer

Value:

    if [ "$(isContainerExist "${CONTAINER}")" = 1 ]
    then
      stopContainer "${CONTAINER_NAME}"
      removeContainer "${CONTAINER_NAME}"
    fi



### Configs._startRunContainer

Value:

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
    --restart {{ .GetConfig "restartPolicy" }} -d "${DOCKER_IMAGE_NAME}{{ if $imageTag }}:{{ $imageTag }}{{ end }}" {{ .GetConfig "command" }}



### Configs._startRunContainerEntryPoint

Value:

    {{ if .GetConfig "entryPoint" -}}
      --entrypoint "{{ .GetConfig "entryPoint" }}" {{ "" -}}
    {{ end -}}



### Configs._startRunContainerEnv

Value:

    {{ $this := . -}}
    {{ if eq (.GetConfig "localhost") "localhost" -}}
      {{ range $key, $val := $this.GetEnvs -}}
        -e "{{ $key }}={{ $val }}" {{ "" -}}
      {{ end -}}
    {{ else -}}
      {{ range $key, $val := $this.GetEnvs -}}
        {{ $val = $this.ReplaceAll $val "localhost" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "127.0.0.1" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "0.0.0.0" ($this.GetConfig "localhost") -}}
        -e "{{ $key }}={{ $val }}" {{ "" -}}
      {{ end -}}
    {{ end -}}



### Configs._startRunContainerPorts

Value:

    {{ $this := . -}}
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



### Configs._startRunContainerVolumes

Value:

    {{ $this := . -}}
    {{ range $index, $volume := .Util.Str.Split (.Util.Str.Trim (.GetConfig "volumes") "\n ") "\n" -}}
      {{ if ne $volume "" -}}
        {{ $volumeParts := $this.Util.Str.Split ($this.Util.Str.Trim $volume  " ") ":" -}}
        {{ if eq (len $volumeParts) 2 -}}
          {{ $absHostVolume := $this.GetWorkPath (index $volumeParts 0) -}}
          {{ $containerVolume := index $volumeParts 1 -}}
          -v "{{ $absHostVolume }}:{{ $containerVolume }}" {{ "" -}}
        {{ end -}}
      {{ end -}}
    {{ end -}}



### Configs.afterCheck


### Configs.afterStart


### Configs.beforeCheck


### Configs.beforeStart


### Configs.check


### Configs.checkCommand


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.command


### Configs.containerName


### Configs.entryPoint


### Configs.finish


### Configs.hostDockerInternal

Value:

    {{ if .GetValue "hostDockerInternal" }}{{ .GetValue "hostDockerInternal" }}{{ else }}host.docker.internal{{ end }}


### Configs.imageName


### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.imageTag


### Configs.localhost

Value:

    localhost


### Configs.network

Value:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}


### Configs.ports


### Configs.rebuild

Value:

    false


### Configs.restartPolicy

Value:

    on-failure


### Configs.runInLocal

Value:

    true


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitConfigVariables

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start


### Configs.strictMode

Value:

    true


### Configs.useImagePrefix

Value:

    true


### Configs.user


### Configs.volumes


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1