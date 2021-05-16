# setupKubeClient
```
      TASK NAME    : setupKubeClient
      LOCATION     : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
      DESCRIPTION  : Install kubectl, helm, and helmfile
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
      INPUTS       : setup.homeDir
                       DESCRIPTION : Home directory
                       PROMPT      : Home directory (Can be blank)
                     setup.helmfileVersion
                       DESCRIPTION : Helmfile version to be installed
                       PROMPT      : Helmfile version to be installed
                       DEFAULT     : v0.138.2
                       VALIDATION  : ^.+$
      CONFIG       :   _setup                 : set -e
                                                {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                                {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                       _start                 : Blank
                       afterStart             : Blank
                       beforeStart            : {{ $d := .Decoration -}}
                                                echo "This command will install Kubectl and helm in your home directory. Root privilege is not required"
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
                       includeUtilScript      : . ${ZARUBA_HOME}/scripts/util.sh
                       playBellScript         : echo $'\a'
                       setup                  : Blank
                       start                  : {{ $d := .Decoration -}}
                                                {{ if .GetValue "setup.homeDir" }}HOME="{{ .GetValue "setup.homeDir" }}"{{ end }}
                                                if [ "$(is_command_exist kubectl version)" = 1 ]
                                                then
                                                  echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Kubectl was already installed{{ $d.Normal }}"
                                                else
                                                  echo "🎡 {{ $d.Bold }}{{ $d.Yellow }}Install Kubectl{{ $d.Normal }}"
                                                  wget "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
                                                  chmod 755 kubectl
                                                  mkdir -p "${HOME}/.local/bin"
                                                  mv kubectl "${HOME}/.local/bin"
                                                fi
                                                if [ "$(is_command_exist helm version)" = 1 ]
                                                then
                                                  echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Helm was already installed{{ $d.Normal }}"
                                                else
                                                  echo "🎡 {{ $d.Bold }}{{ $d.Yellow }}Install helm{{ $d.Normal }}"
                                                  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
                                                  chmod 700 get_helm.sh
                                                  export HELM_INSTALL_DIR="${HOME}/.local/bin"
                                                  ./get_helm.sh --no-sudo
                                                  rm ./get_helm.sh
                                                fi
                                                if [ "$(is_command_exist helmfile --version)" = 1 ]
                                                then
                                                  echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Helmfile was already installed{{ $d.Normal }}"
                                                else
                                                  echo "🎡 {{ $d.Bold }}{{ $d.Yellow }}Install helmfile{{ $d.Normal }}"
                                                  wget https://github.com/roboll/helmfile/releases/download/{{ .GetValue "setup.helmfileVersion" }}/helmfile_linux_amd64
                                                  chmod 755 ./helmfile_linux_amd64
                                                  mkdir -p "${HOME}/.local/bin"
                                                  mv ./helmfile_linux_amd64 "${HOME}/.local/bin/helmfile"
                                                fi
      ENVIRONMENTS : PYTHONUNBUFFERED
                       FROM    : PYTHONUNBUFFERED
                       DEFAULT : 1
```
