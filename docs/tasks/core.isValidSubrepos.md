# core.isValidSubrepos
```
      TASK NAME    : core.isValidSubrepos
      LOCATION     : /home/gofrendi/.zaruba/scripts/core.zaruba.yaml
      TASK TYPE    : Command Task
      PARENT TASKS : [ core.runCoreScript ]
      START        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{ .Trim (.GetConfig "_setup") "\n " }}
                       {{ .Trim (.GetConfig "setup") "\n " }}
                       {{ .Trim (.GetConfig "beforeStart") "\n " }}
                       {{ .Trim (.GetConfig "_start") "\n " }}
                       {{ .Trim (.GetConfig "start") "\n " }}
                       {{ .Trim (.GetConfig "afterStart") "\n " }}
                       {{ .Trim (.GetConfig "finish") "\n " }}
      CONFIG       :   _setup                 : set -e
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                       _start                 : Blank
                       afterStart             : Blank
                       beforeStart            : Blank
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
                                                BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                                . "${BOOTSTRAP_SCRIPT}"
                       includeUtilScript      : . ${ZARUBA_HOME}/scripts/util.sh
                       playBellScript         : echo $'\a'
                       setup                  : Blank
                       start                  : {{ $d := .Decoration -}}
                                                {{ $names := .GetSubValueKeys "subrepo" -}}
                                                {{ $this := . -}}
                                                {{ range $index, $name := $names -}}
                                                  PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
                                                  URL="{{ $this.GetValue "subrepo" $name "url" }}"
                                                  NAME="{{ $name }}"
                                                  should_not_be_empty "${URL}" "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have url{{ $d.Normal }}"
                                                  should_not_be_empty "${PREFIX}" "{{ $d.Bold }}{{ $d.Red }}Subrepo ${NAME} doesn't have prefix{{ $d.Normal }}"
                                                {{ end }}
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}All Subrepos are valid{{ $d.Normal }}"
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
