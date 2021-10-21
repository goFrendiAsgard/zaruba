# zrbHelmInstall
```
  TASK NAME     : zrbHelmInstall
  LOCATION      : /zaruba-tasks/_base/helmChore/task.zrbHelmInstall.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbGenerateAndRun ]
  DEPENDENCIES  : [ zrbSetKubeContext ]
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
                  _prepare                  : Blank
                  _prepareBase              : {{ .GetConfig "_prepareEnvReplacementMap" }}
                                              {{ .GetConfig "_prepare" }}
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
                                              _ZRB_SQL='{{ .GetConfig "sql" }}'
                                              _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
                                              _setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}"
                                              _setReplacementMap "ztplScript" "${_ZRB_SCRIPT}"
                                              _setReplacementMap "ztplSql" "${_ZRB_SQL}"
                                              __ZRB_PWD=$(pwd)
                                              echo "{{ $d.Yellow }}🧰 Prepare{{ $d.Normal }}"
                                              {{ .GetConfig "_prepareBase" }} 
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Yellow }}🚧 Generate{{ $d.Normal }}"
                                              echo "{{ $d.Yellow }}🚧 _ZRB_TEMPLATE_LOCATION:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATION}"
                                              echo "{{ $d.Yellow }}🚧 _ZRB_REPLACEMENT_MAP:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
                                              mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
                                              "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}🏁 Run Script{{ $d.Normal }}"
                                              {{ .GetConfig "runGeneratedScript" }} 
                                              cd "${__ZRB_PWD}"
                  afterStart                : {{ $d := .Decoration -}}
                                              echo 🎉🎉🎉
                                              echo "{{ $d.Bold }}{{ $d.Yellow }}Done{{ $d.Normal }}"
                  beforeStart               : Blank
                  chart                     : {{ .ZarubaHome }}/zaruba-tasks/helm/chart
                  cmd                       : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                    : -c
                  finish                    : Blank
                  generatedScriptLocation   : {{ .GetProjectPath "tmp" }}/{{ .Name }}.script.{{ .UUID }}
                  includeShellUtil          : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  kubeContext               : {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}
                  kubeNmespace              : {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}
                  name                      : Blank
                  runGeneratedScript        : helm install -f "{{ .GetProjectPath "tmp" }}/{{ .Name }}/values.yaml" "{{ .GetConfig "name" }}" "{{ .GetConfig "chart" }}" 
                  script                    : {{ .GetValue "script" }}
                  setup                     : Blank
                  sql                       : {{ .GetValue "sql" }}
                  start                     : Blank
                  templateLocation          : {{ .ZarubaHome }}/zaruba-tasks/helm/template
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```