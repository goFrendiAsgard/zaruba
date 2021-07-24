# core.helmInstall
```
  TASK NAME     : core.helmInstall
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmInstall.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runShellScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : Blank
                  chart                  : Blank
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  finish                 : Blank
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  releaseName            : Blank
                  setup                  : Blank
                  start                  : {{ $fileContent := .ParseFile (.GetConfig "valueTemplateFile") }}
                                           {{ $err := .WriteFile (.GetConfig "valueFile") $fileContent -}}
                                           if [ "$(is_command_error helm status "{{ .GetConfig "releaseName" }}" ]
                                           then
                                             helm install "{{ .GetConfig "releaseName" }}" "{{ .GetConfig "chart" }}" -f "{{ .GetConfig "valueFile" }}"
                                           else
                                             helm upgrade "{{ .GetConfig "releaseName" }}" "{{ .GetConfig "chart" }}" -f "{{ .GetConfig "valueFile" }}"
                                           fi
                  valueFile              : Blank
                  valueTemplateFile      : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```