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
  CONFIG        : _finish                    : Blank
                  _prepareBase               : {{ .GetConfig "_prepareBaseVariables" }}
                                               {{ .GetConfig "_prepareVariables" }}
                                               {{ .GetConfig "_prepareBaseReplacementMap" }}
                                               {{ .GetConfig "_prepareReplacementMap" }}
                  _prepareBaseReplacementMap : . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareReplacementMap.sh"
                  _prepareBaseVariables      : . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/prepareVariables.sh"
                  _prepareReplacementMap     : Blank
                  _prepareVariables          : Blank
                  _setup                     : set -e
                                               {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                     : {{ $d := .Decoration -}}
                                               . "{{ .ZarubaHome }}/zaruba-tasks/_base/generateAndRun/bash/util.sh"
                                               _ZRB_TEMPLATE_LOCATION='{{ .GetConfig "templateLocation" }}'
                                               _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
                                               _ZRB_TASK_NAME="{{ .Name }}"
                                               _ZRB_REPLACEMENT_MAP='{}'
                                               _ZRB_SCRIPT='{{ .GetConfig "script" }}'
                                               _ZRB_SQL='{{ .GetConfig "sql" }}'
                                               _ZRB_IMAGE_NAME="{{ .GetDockerImageName }}"
                                               _ZRB_IMAGE_TAG="{{ .GetConfig "imageTag" }}"
                                               _ZRB_ENVS='{{ .ToJSON .GetEnvs }}'
                                               __ZRB_PWD=$(pwd)
                                               echo "{{ $d.Yellow }}🧰 Prepare{{ $d.Normal }}"
                                               {{ .GetConfig "_prepareBase" }} 
                                               cd "${__ZRB_PWD}"
                                               echo "{{ $d.Yellow }}✅ Validate{{ $d.Normal }}"
                                               {{ .GetConfig "_validateTemplateLocation" }}
                                               {{ .GetConfig "_validate" }}
                                               cd "${__ZRB_PWD}"
                                               echo "{{ $d.Yellow }}🚧 Generate{{ $d.Normal }}"
                                               echo "{{ $d.Yellow }}🚧 Template Location:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATION}"
                                               echo "{{ $d.Yellow }}🚧 Generated Script Location:{{ $d.Normal }} ${_ZRB_GENERATED_SCRIPT_LOCATION}"
                                               echo "{{ $d.Yellow }}🚧 Replacement Map:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
                                               mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
                                               "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
                                               cd "${__ZRB_PWD}"
                                               echo "{{ $d.Bold }}{{ $d.Yellow }}🏁 Run Script{{ $d.Normal }}"
                                               {{ .GetConfig "runGeneratedScript" }} 
                                               cd "${__ZRB_PWD}"
                  _validate                  : Blank
                  _validateTemplateLocation  : {{ $d := .Decoration -}}
                                               if [ ! -x "${_ZRB_TEMPLATE_LOCATION}" ]
                                               then
                                                 echo "{{ $d.Red }}{{ $d.Bold }}Template Location doesn't exist: ${_ZRB_TEMPLATE_LOCATION}.{{ $d.Normal }}"
                                                 exit 1
                                               fi
                  afterStart                 : {{ $d := .Decoration -}}
                                               echo 🎉🎉🎉
                                               echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"
                  beforeStart                : Blank
                  cmd                        : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                     : -c
                  containerName              : Blank
                  finish                     : Blank
                  generatedScriptLocation    : {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}
                  includeShellUtil           : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  remoteCommand              : sh "{{ .GetConfig "remoteScriptLocation" }}/run.sh"
                  remoteScriptLocation       : _{{ .Name }}.script.{{ .UUID }}
                  runGeneratedScript         : _ZRB_CONTAINER_NAME="{{ .GetConfig "containerName" }}"
                                               _ZRB_REMOTE_SCRIPT_LOCATION="{{ .GetConfig "remoteScriptLocation" }}"
                                               chmod 755 -R "${_ZRB_GENERATED_SCRIPT_LOCATION}"
                                               docker cp "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_CONTAINER_NAME}:${_ZRB_REMOTE_SCRIPT_LOCATION}"
                                               docker exec "${_ZRB_CONTAINER_NAME}" {{ .GetConfig "remoteCommand" }}
                                               docker exec -u 0 "${_ZRB_CONTAINER_NAME}" rm -Rf "${_ZRB_REMOTE_SCRIPT_LOCATION}"
                                               rm -Rf "${_ZRB_GENERATED_SCRIPT_LOCATION}"
                  script                     : {{ .GetValue "script" }}
                  setup                      : Blank
                  sql                        : {{ .GetValue "sql" }}
                  start                      : Blank
                  templateLocation           : {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```