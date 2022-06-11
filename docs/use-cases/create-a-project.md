<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# 🏗️ Create a Project
<!--endTocHeader-->

The recommended way to create a project is by invoking `zaruba please initProject`:

__Example:__

<!--startCode-->
```bash
mkdir -p examples/playground/use-cases/newProject
cd examples/playground/use-cases/newProject
zaruba please initProject

tree
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.89µs
         Current Time: 17:10:39
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject
💀    🚀 initProject          🚧 /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject is a zaruba project.
💀 🔥 Error running 🚧 'initProject' command:
        * bash
        * -c
        *    1 | set -e
             2 | . /home/gofrendi/zaruba/zaruba-tasks/_base/run/bash/shellUtil.sh
             3 | _NORMAL='';_BOLD='';_FAINT='';_ITALIC='';_UNDERLINE='';_BLINK_SLOW='';_BLINK_RAPID='';_INVERSE='';_CONCEAL='';_CROSSED_OUT='';_BLACK='';_RED='';_GREEN='';_YELLOW='';_BLUE='';_MAGENTA='';_CYAN='';_WHITE='';_BG_BLACK='';_BG_RED='';_BG_GREEN='';_BG_YELLOW='';_BG_BLUE='';_BG_MAGENTA='';_BG_CYAN='';_BG_WHITE='';_NO_UNDERLINE='';_NO_INVERSE='';_NO_COLOR='';_ZARUBA_ICON='💀';_SUCCESS_ICON='🎉';_ERROR_ICON='🔥';_START_ICON='🏁';_KILL_ICON='🔪';_INSPECT_ICON='🔎';_RUN_ICON='🚀';_WORKER_ICON='👷';_SCRIPT_ICON='📜';_CONSTRUCTION_ICON='🚧';_CONTAINER_ICON='🐳';_EMPTY='  '
             4 | 
             5 | 
             6 | 
             7 | if [ -f "index.zaruba.yaml" ]
             8 | then
             9 |   echo "${_BOLD}${_RED}$(pwd) is a zaruba project.${_NORMAL}"
            10 |   exit 1
            11 | fi
            12 | git init
            13 | "/home/gofrendi/zaruba/zaruba" file copy "/home/gofrendi/zaruba/zaruba-tasks/chore/initProject/template/" .
            14 | touch .env
            15 | echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
            16 | echo "${_BOLD}${_YELLOW}Project created${_NORMAL}"
            17 | 
            18 | 
            19 | 
            20 | 
exit status 1
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 411.472868ms
         Current Time: 17:10:40
zaruba please initProject -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["initProject"]
🔥 Stderr    : exit status 1
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```````
</details>
<!--endCode-->

Aside from generating a project using `zaruba please initProject`, you can also clone/fork existing project from Github or other Git server. Please see [../core-concepts/projects/README.md] for more information.

# Initial Project Structure

Typically, a new project contains of two files:

* `default.values.yaml`: The default project value
* `index.zaruba.yaml`: The entry point of project's zaruba script.

# What's Next

Once you created an empty project, you can start [adding resources to your project](add-resources/README.md), [run some tasks](../run-task/README.md), and [syncrhonize task environments](syncrhonize-task-environments.md)

<!--startTocSubTopic-->
<!--endTocSubTopic-->
