# setupSpark
```
  TASK NAME     : setupSpark
  LOCATION      : /home/gofrendi/.zaruba/scripts/core.setup.zaruba.yaml
  DESCRIPTION   : Install spark and hadoop.
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
                    DESCRIPTION : Home directory
                    PROMPT      : Home directory (Can be blank)
                  setup.sparkDownloadUrl
                    DESCRIPTION : Spark download url
                    PROMPT      : Spark download url
                    DEFAULT     : https://downloads.apache.org/spark/spark-3.1.1/spark-3.1.1-bin-hadoop2.7.tgz
                    VALIDATION  : ^.+$
                  setup.sparkVersion
                    DESCRIPTION : Spark version to be installed
                    PROMPT      : Spark version to be installed
                    DEFAULT     : 3.1.1
                    VALIDATION  : ^.+$
                  setup.hadoopVersion
                    DESCRIPTION : Hadoop version to be installed when install spark
                    PROMPT      : Hadoop version to be installed when install spark
                    DEFAULT     : 2.7
                    VALIDATION  : ^.+$
  CONFIG        : _setup                 : set -e
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterStart             : Blank
                  beforeStart            : {{ $d := .Decoration -}}
                                           echo "This command will install spark and hadoop in your home directory. Root privilege is not required"
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
                                           if [ "$(is_command_exist spark-shell --version)" = 1 ]
                                           then
                                             echo "👏 {{ $d.Bold }}{{ $d.Yellow }}Spark was already installed{{ $d.Normal }}"
                                           else
                                             echo "☕ {{ $d.Bold }}{{ $d.Yellow }}Install spark and hadoop{{ $d.Normal }}"
                                             wget -O spark.tgz "{{ .GetValue "setup.sparkDownloadUrl" }}"
                                             mkdir -p "${HOME}/.spark"
                                             tar -xvzf spark.tgz -C "${HOME}/.spark"
                                             TEMPLATE_CONTENT='{{ .ParseFile (.GetRelativePath "./templates/shell/spark.sh") }}'
                                             append_if_exist "${TEMPLATE_CONTENT}" "${BOOTSTRAP_SCRIPT}"
                                             . "${BOOTSTRAP_SCRIPT}"
                                             rm spark.tgz
                                           fi
                                           echo 🎉🎉🎉
                                           echo "{{ $d.Bold }}{{ $d.Yellow }}Complete !!!{{ $d.Normal }}"
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
