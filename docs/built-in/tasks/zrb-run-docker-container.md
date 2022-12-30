<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Tasks](README.md)
# zrbRunDockerContainer
<!--endTocHeader-->


## Information

File Location:

    ${ZARUBA_HOME}zaruba-tasks/_base/run/dockerContainer/task.zrbRunDockerContainer.yaml

Should Sync Env:

    true

Type:

    simple

Description:

    Run docker container.
    If container is already started, its stdout/stderr will be shown.
    If container is exist but not started, it will be started.
    If container does not exist, it will be created and started.
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
      escapedEnvs    : Escaped envs would not be altered/parsed into host.docker.internal,
                       separated by new line.
      ports          : Port to be checked to confirm service readiness, 
                       separated by new line.
      volumes        : Host-container volume mappings,
                       separated by new line.
      user           : docker user (e.g., 0 for root)
      shmSize        : Size of /dev/shm. The format is <number><unit>. number must be greater than 0. 
                       Unit is optional and can be b (bytes), k (kilobytes), m (megabytes), or g (gigabytes).
                       If you omit the unit, the system uses bytes. If you omit the size entirely, the system uses 64m.
      memory         : Memory limit, default: 512m
      cpus           : CPU limit, default: 1
      gpus           : GPU config, default: unset. Possible value: 'all,capabilities=utlity'
      rebuild        : Should container be rebuild (This will not rebuild the image)
      command        : Command to be used (Single Line).
                       Leave blank to use container's CMD.
                       The command will be executed from inside the container.
      checkCommand   : Command to check container readiness (Single Line).
                       The command will be executed from inside the container.
      localhost      : Localhost mapping (e.g., host.docker.internal)



## Extends

- [zrbRunShellScript](zrb-run-shell-script.md)


## Dependencies

- [updateProjectLinks](update-project-links.md)
- [zrbCreateDockerNetwork](zrb-create-docker-network.md)


## Start

- `{{ .GetConfig "cmd" }}`
- `{{ .GetConfig "cmdArg" }}`
-
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


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
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
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Container '${CONTAINER_NAME}' is already started${_NORMAL}"
      {{ .GetConfig "_startLogContainer" }}
    elif [ ! -z $(inspectDocker "container" ".Name" "${CONTAINER_NAME}") ]
    then
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Retrieve previous log of '${CONTAINER_NAME}'${_NORMAL}"
      sleep 1
      docker logs --tail 20 "${CONTAINER_NAME}"
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Starting container '${CONTAINER_NAME}'${_NORMAL}"
      docker start "${CONTAINER_NAME}"
      {{ .GetConfig "_startLogContainer" }}
    else
      echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Creating and starting container '${CONTAINER_NAME}'${_NORMAL}"
      {{ .GetConfig "_startRunContainer" }}
      {{ .GetConfig "_startLogContainer" }}
    fi



### Configs._startLogContainer

Value:

    echo "${_CONTAINER_ICON} ${_BOLD}${_YELLOW}Logging '${CONTAINER_NAME}'${_NORMAL}"
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
    {{ .GetConfig "dockerOptions" }} {{ "" -}}
    {{ if .GetConfig "user" }}--user "{{ .GetConfig "user" }}" {{ end }} {{ "" -}}
    {{ if .GetConfig "memory" }}--memory "{{ .GetConfig "memory" }}" {{ end }}{{ "" -}}
    {{ if .GetConfig "cpus" }}--cpus "{{ .GetConfig "cpus" }}" {{ end }}{{ "" -}}
    {{ if .GetConfig "gpus" }}--gpus "{{ .GetConfig "gpus" }}" {{ end }}{{ "" -}}
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
    {{ $escapedEnvs := .Util.Str.Split (.Util.Str.Trim (.GetConfig "escapedEnvs") "\n ") "\n " -}}
    {{ range $key, $val := $this.GetEnvs -}}
      {{ if or ($this.Util.List.Contains $escapedEnvs $key) (eq ($this.GetConfig "localhost") "localhost") -}}
        -e {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }} {{ "" -}}
      {{ else -}}
        {{ $val = $this.ReplaceAll $val "localhost" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "127.0.0.1" ($this.GetConfig "localhost") -}}
        {{ $val = $this.ReplaceAll $val "0.0.0.0" ($this.GetConfig "localhost") -}}
        -e {{ $this.Util.Str.EscapeShellValue (printf "%s=%s" $key $val) }} {{ "" -}}
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



### Configs.afterStart


### Configs.beforeStart


### Configs.checkCommand


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.command


### Configs.containerName


### Configs.cpus

Value:

    1


### Configs.dockerOptions


### Configs.entryPoint


### Configs.escapedEnvs


### Configs.finish


### Configs.gpus


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


### Configs.memory

Value:

    512m


### Configs.network

Value:

    {{ if .GetValue "defaultNetwork" }}{{ .GetValue "defaultNetwork" }}{{ else }}zaruba{{ end }}


### Configs.ports


### Configs.rebuild

Value:

    false


### Configs.restartPolicy

Value:

    no


### Configs.setup


### Configs.shmSize

Value:

    100m


### Configs.shouldInitConfigMapVariable

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



# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->