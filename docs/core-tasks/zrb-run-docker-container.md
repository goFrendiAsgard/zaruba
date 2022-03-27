<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 🐳 zrbRunDockerContainer
<!--endTocHeader-->

[1m[33m## Information[0m

[1m[34mFile Location[0m:

    ~/.zaruba/zaruba-tasks/_base/run/dockerContainer/task.zrbRunDockerContainer.yaml

[1m[34mShould Sync Env[0m:

    true

[1m[34mType[0m:

    command

[1m[34mDescription[0m:

    Run docker container.
    If container is already started, its stdout/stderr will be shown.
    If container is exist but not started, it will be started.
    If container is not exist, it will be created and started.
    Common configs:
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



[1m[33m## Extends[0m

* [zrbRunShellScript](zrb-run-shell-script.md)


[1m[33m## Dependencies[0m

* [updateProjectLinks](update-project-links.md)
* [zrbCreateDockerNetwork](zrb-create-docker-network.md)


[1m[33m## Start[0m

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}

    ```


[1m[33m## Configs[0m


[1m[33m### Configs._finish[0m


[1m[33m### Configs._initShell[0m

[1m[34mValue[0m:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



[1m[33m### Configs._setup[0m

[1m[34mValue[0m:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupContainerName") "\n" }} 
    {{ .Util.Str.Trim (.GetConfig "_setupImageName") "\n" }} 



[1m[33m### Configs._setupContainerName[0m

[1m[34mValue[0m:

    CONTAINER_NAME="{{ .GetConfig "containerName" }}"
    if [ -z "${CONTAINER_NAME}" ]
    then
      echo "${_BOLD}${_RED}containerName is not provided${_NORMAL}"
      exit 1
    fi



[1m[33m### Configs._setupImageName[0m

[1m[34mValue[0m:

    DOCKER_IMAGE_NAME="{{ .GetDockerImageName }}"
    if [ -z "${DOCKER_IMAGE_NAME}" ]
    then
      echo "${_BOLD}${_RED}imageName is not provided${_NORMAL}"
      exit 1
    fi



[1m[33m### Configs._start[0m

[1m[34mValue[0m:

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



[1m[33m### Configs._startLogContainer[0m

[1m[34mValue[0m:

    echo "🐳 ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
    docker logs --since 0m --follow "${CONTAINER_NAME}"



[1m[33m### Configs._startRebuildContainer[0m

[1m[34mValue[0m:

    if [ "$(isContainerExist "${CONTAINER}")" = 1 ]
    then
      stopContainer "${CONTAINER_NAME}"
      removeContainer "${CONTAINER_NAME}"
    fi



[1m[33m### Configs._startRunContainer[0m

[1m[34mValue[0m:

    {{ $imageTag := .GetConfig "imageTag" -}}
    {{ $this := . -}}
    docker run --name "${CONTAINER_NAME}" {{ "" -}}
    --hostname "${CONTAINER_NAME}" {{ "" -}}
    --network "{{ .GetConfig "network" }}" {{ "" -}}
    {{ if .GetConfig "user" }}--user "{{ .GetConfig "user" }}" {{ end }} {{ "" -}}
    {{ if .GetConfig "shmSize" }}--shm-size "{{ .GetConfig "shmSize" }}" {{ end }}{{ "" -}}
    {{ .GetConfig "_startRunContainerEntryPoint" -}}
    {{ .GetConfig "_startRunContainerEnv" -}}
    {{ .GetConfig "_startRunContainerPorts" -}}
    {{ .GetConfig "_startRunContainerVolumes" -}}
    {{ if ne (.GetConfig "hostDockerInternal") "host.docker.internal" }}--add-host "{{ .GetConfig "hostDockerInternal" }}:host.docker.internal"{{ end }} {{ "" -}}
    --restart {{ .GetConfig "restartPolicy" }} -d "${DOCKER_IMAGE_NAME}{{ if $imageTag }}:{{ $imageTag }}{{ end }}" {{ .GetConfig "command" }}



[1m[33m### Configs._startRunContainerEntryPoint[0m

[1m[34mValue[0m:

    {{ if .GetConfig "entryPoint" -}}
      --entrypoint "{{ .GetConfig "entryPoint" }}" {{ "" -}}
    {{ end -}}



[1m[33m### Configs._startRunContainerEnv[0m

[1m[34mValue[0m:

    {{ $this := . -}}
    {{ if eq (.GetConfig "localhost") "localhost" -}}
      {{ range $key, $val := $this.GetEnvs -}}
        -e {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }} {{ "" -}}
      {{ end -}}
    {{ else -}}
      {{ range $key, $val := $this.GetEnvs -}}
        {{ $val = $this.ReplaceAll $val "localhost" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "127.0.0.1" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "0.0.0.0" ($this.GetConfig "localhost") -}}
        -e {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }} {{ "" -}}
      {{ end -}}
    {{ end -}}



[1m[33m### Configs._startRunContainerPorts[0m

[1m[34mValue[0m:

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



[1m[33m### Configs._startRunContainerVolumes[0m

[1m[34mValue[0m:

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



[1m[33m### Configs.afterStart[0m


[1m[33m### Configs.beforeStart[0m


[1m[33m### Configs.checkCommand[0m


[1m[33m### Configs.cmd[0m

[1m[34mValue[0m:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


[1m[33m### Configs.cmdArg[0m

[1m[34mValue[0m:

    -c


[1m[33m### Configs.command[0m


[1m[33m### Configs.containerName[0m


[1m[33m### Configs.entryPoint[0m


[1m[33m### Configs.finish[0m


[1m[33m### Configs.hostDockerInternal[0m

[1m[34mValue[0m:

    {{ if .GetValue "hostDockerInternal" }}{{ .GetValue "hostDockerInternal" }}{{ else }}host.docker.internal{{ end }}


[1m[33m### Configs.imageName[0m


[1m[33m### Configs.imagePrefix[0m

[1m[34mValue[0m:

    {{ .GetValue "defaultImagePrefix" }}


[1m[33m### Configs.imageTag[0m


[1m[33m### Configs.localhost[0m

[1m[34mValue[0m:

    localhost


[1m[33m### Configs.network[0m

[1m[34mValue[0m:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}


[1m[33m### Configs.ports[0m


[1m[33m### Configs.rebuild[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.restartPolicy[0m

[1m[34mValue[0m:

    no


[1m[33m### Configs.setup[0m


[1m[33m### Configs.shmSize[0m

[1m[34mValue[0m:

    100m


[1m[33m### Configs.shouldInitConfigMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitEnvMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitUtil[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.start[0m


[1m[33m### Configs.strictMode[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.useImagePrefix[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.user[0m


[1m[33m### Configs.volumes[0m


[1m[33m## Envs[0m


[1m[33m### Envs.PYTHONUNBUFFERED[0m

[1m[34mFrom[0m:

    PYTHONUNBUFFERED

[1m[34mDefault[0m:

    1