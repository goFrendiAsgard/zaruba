# core.helmGenerateValue
```
  TASK NAME     : core.helmGenerateValue
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmGenerateValue.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup            : set -e
                                      {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start            : Blank
                  afterStart        : Blank
                  beforeStart       : Blank
                  cmd               : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg            : -c
                  finish            : Blank
                  imagePrefix       : Blank
                  imageTag          : Blank
                  includeUtilScript : . ${ZARUBA_HOME}/bash/util.sh
                  setup             : Blank
                  start             : {{ $templateFile := .GetWorkPath (.GetConfig "valueTemplateFile") -}}
                                      {{ $valueFile := .GetWorkPath (.GetConfig "valueFile") -}}
                                      {{ $fileContent := .ParseFile $templateFile -}}
                                      {{ $err := .WriteFile $valueFile $fileContent -}}
                                      {{ if $err -}}
                                      echo '{{ $err }}'
                                      exit 1
                                      {{ end -}}
                  useImagePrefix    : true
                  valueFile         : Blank
                  valueTemplateFile : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```