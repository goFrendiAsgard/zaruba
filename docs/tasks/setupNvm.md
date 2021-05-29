# setupNvm
```
  TASK NAME     : setupNvm
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Install NVM and Node.Js (Including node-gyp and typescript)
  TASK TYPE     : Command Task
  PARENT TASKS  : [ core.runCoreScript ]
  START         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeStart") "\n " }}
                    {{ .Trim (.GetConfig "_start") "\n " }}
                    {{ .Trim (.GetConfig "start") "\n " }}
                    {{ .Trim (.GetConfig "afterStart") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
  INPUTS        : setup.homeDir
                    DESCRIPTION : Home directory (Can be blank)
                    PROMPT      : Home directory
                  setup.nodeVersion
                    DESCRIPTION : Node version to be installed when install nvm
                    PROMPT      : Node version
                    OPTIONS     : [ node, stable, unstable ]
                    DEFAULT     : node
                    VALIDATION  : ^.+$
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will install nvm, typescript, and node-gyp in your home directory. Root privilege is not required"
                                           echo "If this command doesn't run successfully, please open an issue on https://github.com/state-alcemists/zaruba."
                                           echo "Please also specify your OS version."
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
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  playBellScript         : echo $'\a'
                  setup                  : Blank
                  start                  : {{ $d := .Decoration -}}
                                           {{ if .GetValue "setup.homeDir" }}HOME="{{ .GetValue "setup.homeDir" }}"{{ end }}
                                           if [ "$(is_command_exist nvm --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}NVM was already installed{{ $d.Normal }}"
                                           else
                                             rm -Rf "${HOME}/.nvm"
                                             echo "🐸 {{ $d.Bold }}{{ $d.Yellow }}Install NVM{{ $d.Normal }}"
                                             curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | sh
                                             TEMPLATE_CONTENT="$(cat "${ZARUBA_HOME}/scripts/templates/shell/nvm.sh")"
                                             append_if_exist "${TEMPLATE_CONTENT}" "${BOOTSTRAP_SCRIPT}"
                                             . "${BOOTSTRAP_SCRIPT}"
                                             nvm install "{{ .GetValue "setup.nodeVersion" }}"
                                             . "${BOOTSTRAP_SCRIPT}"
                                           fi
                                           if [ "$(is_command_exist node-gyp --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Node-gyp was already installed{{ $d.Normal }}"
                                           else 
                                             echo "🐸 {{ $d.Bold }}{{ $d.Yellow }}Install node-gyp{{ $d.Normal }}"
                                             npm install -g node-gyp
                                           fi
                                           if [ "$(is_command_exist tsc --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Typescript was already installed{{ $d.Normal }}"
                                           else
                                             echo "🐸 {{ $d.Bold }}{{ $d.Yellow }}Install typescript{{ $d.Normal }}"
                                             npm install -g typescript
                                           fi
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Complete !!!{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
