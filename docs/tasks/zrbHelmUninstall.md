# zrbHelmUninstall
```
  TASK NAME     : zrbHelmUninstall
  LOCATION      : /zaruba-tasks/_base/helmChore/task.zrbHelmUninstall.yaml
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
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
  CONFIG        : _finish               : Blank
                  _setup                : set -e
                                          {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start                : {{ $d := .Decoration -}}
                                          _ZRB_RELEASE_NAME='{{ .GetConfig "releaseName" }}'
                                          _ZRB_KEBAB_RELEASE_NAME="$("{{ .ZarubaBin }}" str toKebab "${_ZRB_RELEASE_NAME}")"
                                          if [ "$(isCommandError helm status "${_ZRB_KEBAB_RELEASE_NAME}")" -eq 1 ]
                                          then
                                            echo "{{ $d.Yellow }}[SKIP] Release doesn't exist: ${_ZRB_KEBAB_RELEASE_NAME}.{{ $d.Normal }}"
                                          else
                                            echo "{{ $d.Yellow }}Uninstall release: ${_ZRB_KEBAB_RELEASE_NAME}.{{ $d.Normal }}"
                                            helm uninstall --namespace "{{ .GetConfig "kubeNamespace" }}" "${_ZRB_KEBAB_RELEASE_NAME}"
                                          fi
                  afterStart            : Blank
                  beforeStart           : Blank
                  chartLocation         : Blank
                  cmd                   : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                : -c
                  finish                : Blank
                  includeShellUtil      : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  kubeContext           : {{ if .GetValue "kubeContext" }}{{ .GetValue "kubeContext" }}{{ else if .GetValue "defaultKubeContext" }}{{ .GetValue "defaultKubeContext" }}docker-desktop{{ end }}
                  kubeNamespace         : {{ if .GetValue "kubeNamespace" }}{{ .GetValue "kubeNamespace" }}{{ else if .GetValue "defaultKubeNamespace" }}{{ .GetValue "defaultKubeNamespace" }}default{{ end }}
                  releaseName           : Blank
                  setup                 : Blank
                  start                 : Blank
                  templateLocation      : {{ .GetConfig "valueTemplateLocation" }}
                  valueTemplateLocation : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```