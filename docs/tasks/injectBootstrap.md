# injectBootstrap
```
      TASK NAME    : injectBootstrap
      LOCATION     : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
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
                                                inject_bootstrap "${HOME}/.bashrc"
                                                inject_bootstrap "${HOME}/.zshrc"
                                                inject_bootstrap "${HOME}/.profile"
                                                echo 🎉🎉🎉
                                                echo "{{ $d.Bold }}{{ $d.Yellow }}Close and reopen terminal to apply changes{{ $d.Normal }}"
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
