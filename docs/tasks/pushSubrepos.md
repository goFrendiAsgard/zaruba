# pushSubrepos
```
  TASK NAME     : pushSubrepos
  LOCATION      : ${ZARUBA_HOME}/scripts/tasks/pushSubrepos.zaruba.yaml
  DESCRIPTION   : Publish subrepositories.
                  ARGUMENTS:
                    subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
                    subrepo::<name>::url      : Remote url of the subrepo
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  DEPENDENCIES  : [ initSubrepos, updateProjectLinks, core.setupPyUtil ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
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
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bash/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . ${ZARUBA_HOME}/scripts/bash/util.sh
                  setup                  : Blank
                  start                  : set -e
                                           {{ $d := .Decoration -}}
                                           {{ $names := .GetSubValueKeys "subrepo" -}}
                                           {{ $this := . -}}
                                           BRANCH="{{ if .GetValue "defaultBranch" }}{{ .GetValue "defaultBranch" }}{{ else }}main{{ end }}"
                                           ORIGINS=$({{ .Zaruba }} split "$(git remote)")
                                           {{ range $index, $name := $names -}}
                                             PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
                                             URL="{{ $this.GetValue "subrepo" $name "url" }}"
                                             NAME="{{ $name }}"
                                             ORIGIN_EXISTS=$({{ .Zaruba }} listContains "${ORIGINS}" "${NAME}")
                                             if [ $ORIGIN_EXISTS = 1 ]
                                             then
                                               git_save.sh" "Save works before p
                                               git subtree push --prefix="${PREFIX}" "${NAME}" "${BRANCH}"
                                             fi
                                           {{ end -}}
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepos pushed{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```