# addLink
```
      TASK NAME    : addLink
      LOCATION     : /home/gofrendi/.zaruba/scripts/core.zaruba.yaml
      DESCRIPTION  : Add link.
                     TIPS: To update links, you should perform `zaruba please updateLinks`
      TASK TYPE    : Command Task
      PARENT TASKS : [ core.runCoreScript ]
      DEPENDENCIES : [ core.isProject, core.setupPyUtil ]
      START        : - {{ .GetConfig "cmd" }}
                     - {{ .GetConfig "cmdArg" }}
                     - {{ .Trim (.GetConfig "_setup") "\n " }}
                       {{ .Trim (.GetConfig "setup") "\n " }}
                       {{ .Trim (.GetConfig "beforeStart") "\n " }}
                       {{ .Trim (.GetConfig "_start") "\n " }}
                       {{ .Trim (.GetConfig "start") "\n " }}
                       {{ .Trim (.GetConfig "afterStart") "\n " }}
                       {{ .Trim (.GetConfig "finish") "\n " }}
      INPUTS       : link.from
                       DESCRIPTION : Link source
                       PROMPT      : Link source (Required)
                       VALIDATION  : ^.+$
                     link.to
                       DESCRIPTION : Link source
                       PROMPT      : Link destination (Required)
                       VALIDATION  : ^.+$
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
                                                SOURCE="{{ .GetValue "link.from" }}"
                                                should_not_be_empty "${SOURCE}" "{{ $d.Bold }}{{ $d.Red }}'link.from' argument is not defined{{ $d.Normal }}"
                                                DESTINATION="{{ .GetValue "link.to" }}"
                                                should_not_be_empty "${DESTINATION}" "{{ $d.Bold }}{{ $d.Red }}'link.to' argument is not defined{{ $d.Normal }}"
                                                set_project_value "link::${DESTINATION)" "${SOURCE}"
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Link ${SOURCE} -> ${DESTINATION} has been added{{ $d.Normal }}"
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
