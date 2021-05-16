# setProjectValue
```
      TASK NAME    : setProjectValue
      LOCATION     : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
      DESCRIPTION  : Set project value.
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
      INPUTS       : variable.name
                       DESCRIPTION : Variable name
                       PROMPT      : Variable name (Required)
                       VALIDATION  : ^.+$
                     variable.value
                       DESCRIPTION : Variable value
                       PROMPT      : Variable value (Required)
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
                                                KEY="{{ .GetValue "variable.name" }}"
                                                should_not_be_empty "${KEY}" "{{ $d.Bold }}{{ $d.Red }}'variable.name' argument is not defined{{ $d.Normal }}"
                                                VALUE="{{ .GetValue "variable.value" }}"
                                                should_not_be_empty "${VALUE}" "{{ $d.Bold }}{{ $d.Red }}'variable.value' argument is not defined{{ $d.Normal }}"
                                                set_project_value "${KEY}" "${VALUE}"
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Kwarg ${KEY} : ${VALUE} has been set{{ $d.Normal }}"
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
