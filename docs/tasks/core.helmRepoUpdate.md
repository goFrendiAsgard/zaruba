# core.helmRepoUpdate
```
  TASK NAME     : core.helmRepoUpdate
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/core.helmRepoUpdate.zaruba.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runShellScript ]
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
  CONFIG        : _finish             : Blank
                  _setup              : Blank
                  _start              : Blank
                  afterStart          : Blank
                  beforeStart         : Blank
                  cmd                 : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg              : -c
                  defaultHelmRepoName : bitnami
                  defaultHelmRepoUrl  : https://charts.bitnami.com/bitnami
                  finish              : Blank
                  setup               : Blank
                  start               : helm repo add {{ .GetConfig "defaultHelmRepoName" }} {{ .GetConfig "defaultHelmRepoUrl" }}
                                        helm repo update
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```