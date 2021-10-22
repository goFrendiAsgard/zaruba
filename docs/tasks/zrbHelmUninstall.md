# zrbHelmUninstall
```
  TASK NAME     : zrbHelmUninstall
  LOCATION      : /zaruba-tasks/_base/helmChore/task.zrbHelmUninstall.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunShellScript ]
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
  CONFIG        : _finish          : Blank
                  _setup           : Blank
                  _start           : helm uninstall --namespace "{{ .GetConfig "kubeNamespace" }}"  "{{ .GetConfig "releaseName" }}"
                  afterStart       : Blank
                  beforeStart      : Blank
                  chartLocation    : {{ .ZarubaHome }}/zaruba-tasks/helm/chart
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  kubeContext      : {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}
                  kubeNmespace     : {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}
                  releaseName      : Blank
                  setup            : Blank
                  start            : Blank
                  templateLocation : {{ .ZarubaHome }}/zaruba-tasks/helm/template
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```