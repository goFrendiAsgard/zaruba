# addSubrepo
```
  TASK NAME     : addSubrepo
  LOCATION      : /zaruba-tasks/chore/subrepo/task.addSubrepo.yaml
  DESCRIPTION   : Add subrepository.
                  TIPS: To init added subrepositories, you should perform `zaruba please initSubrepos`
  TASK TYPE     : Command Task
  PARENT TASKS  : [ zrbRunCoreScript ]
  DEPENDENCIES  : [ zrbIsProject ]
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
  INPUTS        : subrepoUrl
                    DESCRIPTION : Subrepo url (Required)
                    PROMPT      : Subrepo url
                    VALIDATION  : ^.+$
                  subrepoPrefix
                    DESCRIPTION : Subrepo directory name (Can be blank)
                    PROMPT      : Subrepo directory name
                  subrepoName
                    DESCRIPTION : Subrepo name (Can be blank)
                    PROMPT      : Subrepo name
  CONFIG        : _finish          : Blank
                  _setup           : set -e
                                     {{ .Util.Str.Trim (.GetConfig "includeShellUtil") "\n" }}
                  _start           : Blank
                  afterStart       : Blank
                  beforeStart      : Blank
                  cmd              : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg           : -c
                  finish           : Blank
                  includeShellUtil : . {{ .ZarubaHome }}/zaruba-tasks/_base/run/coreScript/bash/shellUtil.sh
                  setup            : Blank
                  start            : set -e
                                     {{ $d := .Decoration -}}
                                     URL="{{ .GetValue "subrepoUrl" }}"
                                     if [ -z "${URL}" ]
                                     then
                                       echo "{{ $d.Bold }}{{ $d.Red }}subrepoUrl is not defined{{ $d.Normal }}"
                                       exit 1
                                     fi
                                     {{ if .GetValue "subrepoPrefix" }}
                                       PREFIX="{{ .GetValue "subrepoPrefix" }}"
                                     {{ else }}
                                       {{ $urlSegment := .Util.Str.Split (.GetConfig "subrepoUrl") "/" -}}
                                       {{ $urlSegmentLastIndex := .Subtract (len $urlSegment) 1 -}}
                                       {{ $urlLastSegment := index $urlSegment $urlSegmentLastIndex -}}
                                       {{ $prefix := index (.Util.Str.Split $urlLastSegment ".") 0 -}}
                                       PREFIX="{{ $prefix }}"
                                     {{ end }}
                                     NAME="{{ if .GetValue "subrepoName" }}{{ .GetValue "subrepoName" }}{{ else }}${PREFIX}{{ end }}"
                                     "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::prefix" "${PREFIX}"
                                     "{{ .ZarubaBin }}" project setValue "{{ .GetWorkPath "default.values.yaml" }}" "subrepo::${NAME}::url" "${URL}"
                                     echo 🎉🎉🎉
                                     echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepo ${NAME} has been added{{ $d.Normal }}"
                  subrepoName      : {{ .GetValue "subrepoName" }}
                  subrepoPrefix    : {{ .GetValue "subrepoPrefix" }}
                  subrepoUrl       : {{ .GetValue "subrepoUrl" }}
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```