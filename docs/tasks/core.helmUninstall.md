# core.helmUninstall
```
  TASK NAME     : core.helmUninstall
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmUninstall.zaruba.yaml
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
  CONFIG        : _setup      : Blank
                  _start      : Blank
                  afterStart  : Blank
                  beforeStart : Blank
                  chart       : Blank
                  cmd         : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg      : -c
                  finish      : Blank
                  kubeContext : {{ .GetValue "kubeContext" }}
                  releaseName : Blank
                  setup       : Blank
                  start       : if [ "$(kubectl config current-context)" != "{{ .GetConfig "kubeContext" }}" ]
                                then
                                  kubectl config set-context "{{ .GetConfig "kubeContext" }}"
                                fi
                                helm uninstall "{{ .GetConfig "releaseName" }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```