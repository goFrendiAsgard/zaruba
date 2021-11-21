
# ZrbStartDockerContainer

File Location:

    /zaruba-tasks/_base/start/task.zrbStartDockerContainer.yaml

Should Sync Env:

    true

Type:

    service

Description:

    Start docker container and check it's readiness.
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
    {{- $d := .Decoration -}}
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
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

    ```


## Check

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{- $d := .Decoration -}}
    {{ if .Util.Bool.IsFalse (.GetConfig "runInLocal") -}}
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
    ```


## Configs


### Configs.imageTag


### Configs.setup


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



### Configs.afterStart


### Configs.beforeCheck


### Configs.hostDockerInternal

Value:

    {{ if .GetValue "hostDockerInternal" }}{{ .GetValue "hostDockerInternal" }}{{ else }}host.docker.internal{{ end }}


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



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



### Configs.checkCommand


### Configs.network

Value:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}


### Configs._setup

Value:

    set -e
    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupContainerName") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupImageName") "\n" }} 



### Configs._startRunContainerEntryPoint

Value:

    {{ if .GetConfig "entryPoint" -}}
      --entrypoint "{{ .GetConfig "entryPoint" }}" {{ "" -}}
    {{ end -}}



### Configs.includeShellUtil

Value:

    true


### Configs.cmdArg

Value:

    -c


### Configs.containerName


### Configs.imageName


### Configs.localhost

Value:

    localhost


### Configs._checkConfigPorts

Value:

    {{ $d := .Decoration -}}
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



### Configs._start

Value:

    {{ $d := .Decoration -}}
    {{ $rebuild := .GetConfig "rebuild" -}}
    {{ if .Util.Bool.IsTrue $rebuild }}{{ .GetConfig "_startRebuildContainer" }}{{ end }}
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



### Configs.beforeStart


### Configs._checkCommand

Value:

    {{ $d := .Decoration -}}
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



### Configs._finish


### Configs._startRebuildContainer

Value:

    stopContainer "${CONTAINER_NAME}"
    removeContainer "${CONTAINER_NAME}"



### Configs.afterCheck


### Configs.restartPolicy

Value:

    on-failure


### Configs.start


### Configs.useImagePrefix

Value:

    true


### Configs.volumes


### Configs._check

Value:

    {{ $d := .Decoration -}}
    {{ .GetConfig "_checkContainerState" }}
    {{ .GetConfig "_checkConfigPorts" }}
    {{ .GetConfig "_checkCommand" }}
    sleep 1



### Configs._startRunContainer

Value:

    {{ $d := .Decoration -}}
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



### Configs.imagePrefix

Value:

    {{ .GetValue "defaultImagePrefix" }}


### Configs.ports


### Configs.finish


### Configs.runInLocal

Value:

    true


### Configs.strictMode

Value:

    true


### Configs.user


### Configs._checkContainerState

Value:

    {{ $d := .Decoration -}}
    until [ "$(inspectDocker "container" ".State.Running" "${CONTAINER_NAME}")" = true ]
    do
      sleep 1
    done
    while [ "$(inspectDocker "container" ".State.Health" "${CONTAINER_NAME}")" = false ]
    do
      sleep 1
    done
    echo "🔎 {{ $d.Bold }}{{ $d.Yellow }}Container '${CONTAINER_NAME}' is running{{ $d.Normal }}"



### Configs.check


### Configs.command


### Configs.entryPoint


### Configs.rebuild

Value:

    false


### Configs._setupContainerName

Value:

    {{ $d := .Decoration -}}
    CONTAINER_NAME="{{ .GetConfig "containerName" }}"
    if [ -z "${CONTAINER_NAME}" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}containerName is not provided{{ $d.Normal }}"
      exit 1
    fi



### Configs._setupImageName

Value:

    {{ $d := .Decoration -}}
    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    if [ -z "${DOCKER_IMAGE_NAME}" ]
    then
      echo "{{ $d.Bold }}{{ $d.Red }}imageName is not provided{{ $d.Normal }}"
      exit 1
    fi



### Configs._startLogContainer

Value:

    {{ $d := .Decoration -}}
    echo "🐳 {{ $d.Bold }}{{ $d.Yellow }}Logging '${CONTAINER_NAME}'{{ $d.Normal }}"
    docker logs --since 0m --follow "${CONTAINER_NAME}"



### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1