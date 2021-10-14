# zrbRunInDockerContainer
```
  TASK NAME     : zrbRunInDockerContainer
  LOCATION      : /zaruba-tasks/_base/run/inDockerContainer/task.zrbRunInDockerContainer.yaml
  DESCRIPTION   : Run command from inside the container
                  Common config:
                    containerName  : Name of the container
                    containerShell : Shell to run script, default to sh
                    command       : Command to be executed, can be multi line
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbGenerateAndRun ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
                    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}
  CONFIG        : _finish                   : Blank
                  _prepareEnvReplacementMap : {{ range $key, $val := .GetEnvs -}}
                                              _setReplacementMap '${{ $key }}' '{{ $val }}'
                                              _setReplacementMap '${{ (print "{" $key "}") }}' '{{ $val }}'
                                              _setReplacementMap 'ztplEnv_{{ $key }}' '{{ $val }}'
                                              {{ end -}}
                  _setup                    : set -e
                                              {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                    : {{ $d := .Decoration -}}
                                              . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
                                              _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
                                              _ZRB_TASK_NAME="{{ .Name }}"
                                              _ZRB_REPLACEMENT_MAP='{}'
                                              _ZRB_SCRIPT='{{ .GetConfig "script" }}'
                                              _setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}"
                                              _setReplacementMap "ztplScript" "${_ZRB_SCRIPT}"
                                              __ZRB_PWD=$(pwd)
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Prepare{{ $d.Normal }}"
                                              {{ .GetConfig "prepare" }} 
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Generate{{ $d.Normal }}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}_ZRB_TEMPLATE_LOCATION:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATION}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}_ZRB_REPLACEMENT_MAP:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
                                              "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "{{ .GetConfig "generatedScriptLocation" }}" "${_ZRB_REPLACEMENT_MAP}"
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Run Generated Script{{ $d.Normal }}"
                                              {{ .GetConfig "runGeneratedScript" }} 
                                              cd "${__ZRB_PWD}"
                  afterStart                : Blank
                  beforeStart               : Blank
                  cmd                       : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                    : -c
                  containerName             : Blank
                  finish                    : Blank
                  generatedScriptLocation   : {{ .GetProjectPath "tmp" }}/${_ZRB_TASK_NAME}
                  includeShellUtil          : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  prepare                   : {{ .GetConfig "_prepareEnvReplacementMap" }}
                  remoteCommand             : sh "{{ .GetConfig "remoteScriptLocation" }}/run.sh"
                  remoteScriptLocation      : _{{ .Name }}Script
                  runGeneratedScript        : CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                              GENERATED_SCRIPT_LOCATION="{{ .GetConfig "generatedScriptLocation" }}"
                                              REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
                                              chmod 755 -R "${GENERATED_SCRIPT_LOCATION}"
                                              docker cp "${GENERATED_SCRIPT_LOCATION}" "${CONTAINER_NAME}:${REMOTE_SCRIPT_LOCATION}"
                                              docker exec "${CONTAINER_NAME}" {{ .GetConfig "remoteCommand" }}
                                              docker exec -u 0 "${CONTAINER_NAME}" rm -Rf "${_REMOTE_SCRIPT_LOCATION}"
                                              rm "${GENERATED_SCRIPT_LOCATION}"
                  script                    : Blank
                  setup                     : Blank
                  start                     : Blank
                  templateLocation          : {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```