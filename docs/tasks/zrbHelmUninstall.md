# zrbHelmUninstall
```
  TASK NAME     : zrbHelmUninstall
  LOCATION      : /scripts/tasks/zrbHelmUninstall.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunShellScript ]
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
  CONFIG        : _finish     : Blank
                  _setup      : Blank
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
                                  kubectl config use-context "{{ .GetConfig "kubeContext" }}"
                                fi
                                helm uninstall "{{ .GetConfig "releaseName" }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```