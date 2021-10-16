# zrbGenerateAndRun
```
  TASK NAME     : zrbGenerateAndRun
  LOCATION      : /zaruba-tasks/_base/generateAndRun/task.zrbGenerateAndRun.yaml
  DESCRIPTION   : Generate script and run it
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
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
                                              _ZRB_SQL='{{ .GetConfig "sql" }}'
                                              _ZRB_GENERATED_SCRIPT_LOCATION='{{ .GetConfig "generatedScriptLocation" }}'
                                              _setReplacementMap "ztplTaskName" "${_ZRB_TASK_NAME}"
                                              _setReplacementMap "ztplScript" "${_ZRB_SCRIPT}"
                                              _setReplacementMap "ztplSql" "${_ZRB_SQL}"
                                              __ZRB_PWD=$(pwd)
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Prepare{{ $d.Normal }}"
                                              {{ .GetConfig "prepare" }} 
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Generate{{ $d.Normal }}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}_ZRB_TEMPLATE_LOCATION:{{ $d.Normal }} ${_ZRB_TEMPLATE_LOCATION}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}_ZRB_REPLACEMENT_MAP:{{ $d.Normal }} ${_ZRB_REPLACEMENT_MAP}"
                                              mkdir -p "${_ZRB_GENERATED_SCRIPT_LOCATION}"
                                              "{{ .ZarubaBin }}" generate "${_ZRB_TEMPLATE_LOCATION}" "${_ZRB_GENERATED_SCRIPT_LOCATION}" "${_ZRB_REPLACEMENT_MAP}"
                                              cd "${__ZRB_PWD}"
                                              echo "{{ $d.Yellow }}{{ $d.Bold }}Run Generated Script{{ $d.Normal }}"
                                              {{ .GetConfig "runGeneratedScript" }} 
                                              cd "${__ZRB_PWD}"
                  afterStart                : Blank
                  beforeStart               : Blank
                  cmd                       : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                    : -c
                  finish                    : Blank
                  generatedScriptLocation   : {{ .GetProjectPath "tmp" }}/{{ .Name }}
                  includeShellUtil          : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  prepare                   : {{ .GetConfig "_prepareEnvReplacementMap" }}
                  runGeneratedScript        : {{ .GetProjectPath "tmp" }}/{{ .Name }}/run.sh
                  script                    : {{ .GetValue "script" }}
                  setup                     : Blank
                  sql                       : {{ .GetValue "sql" }}
                  start                     : Blank
                  templateLocation          : {{ .ZarubaHome }}/zaruba-tasks/generateAndRun/template
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```