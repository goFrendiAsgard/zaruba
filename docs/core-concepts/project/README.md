<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md)
# 🏗️ Project
<!--endTocHeader-->

Project is a container of your tasks and any other resources. A project is usually also a git repository and a monorepo.

But, any directory containing `index.zaruba.yaml` is a valid project.

There are several ways to create a project.

# Create an Empty Project from Scratch

To create an empty project from scratch, you need to:

* make an empty git repository.
* create a file named `index.zaruba.yaml`.

__Example:__

<!--startCode-->
```bash
mkdir -p examples/playground/myProjectFromScratch
cd examples/playground/myProjectFromScratch
git init
touch index.zaruba.yaml

echo 💀 Project structure
tree
```
 
<details>
<summary>Output</summary>
 
```````
Reinitialized existing Git repository in /home/gofrendi/zaruba/docs/examples/playground/myProjectFromScratch/.git/
💀 Project structure
.
└── index.zaruba.yaml

0 directories, 1 file
```````
</details>
<!--endCode-->

# Generate a New Project

The recommended way to create a new project is by generate it.

To do this, you need to:

* create an empty directory.
* invoke `zaruba please initProject` from inside the directory.

__Example:__

<!--startCode-->
```bash
mkdir -p examples/playground/myGeneratedProject
cd examples/playground/myGeneratedProject
zaruba please initProject

echo 💀 Project structure
tree
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 2.046µs
         Current Time: 13:08:11
💀 🏁 Running 🚧 initProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject is a zaruba project.
💀 🔥 Exit 🚧 initProject runner (Attempt 1 of 3):
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
            14 | echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
            15 | echo "${_BOLD}${_YELLOW}Project created${_NORMAL}"
            16 | 
            17 | 
            18 | 
            19 | 
exit status 1
💀 🏁 Running 🚧 initProject runner (Attempt 2 of 3) on /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject is a zaruba project.
💀 🔥 Exit 🚧 initProject runner (Attempt 2 of 3):
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
            14 | echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
            15 | echo "${_BOLD}${_YELLOW}Project created${_NORMAL}"
            16 | 
            17 | 
            18 | 
            19 | 
exit status 1
💀 🏁 Running 🚧 initProject runner (Attempt 3 of 3) on /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject is a zaruba project.
💀 🔥 Exit 🚧 initProject runner (Attempt 3 of 3):
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
            14 | echo ${_SUCCESS_ICON}${_SUCCESS_ICON}${_SUCCESS_ICON}
            15 | echo "${_BOLD}${_YELLOW}Project created${_NORMAL}"
            16 | 
            17 | 
            18 | 
            19 | 
exit status 1
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 2.32672306s
         Current Time: 13:08:13
zaruba please initProject  -v '/home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["initProject"]
🔥 Stderr    : exit status 1
💀 Project structure
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```````
</details>
<!--endCode-->

# Clone an Existing Project

In some cases, someone has already created a project for you and make it accessible from the internet.

To clone/fork existing projects from GitHub or other git servers do:

```bash
git clone git@github.com:<user>/<repo>.git
```

__Example:__

<!--startCode-->
```bash
cd examples/playground
git clone git@github.com:state-alchemists/zaruba-project myClonedProject
cd myClonedProject

echo 💀 Project structure
tree
```
 
<details>
<summary>Output</summary>
 
```````
fatal: destination path 'myClonedProject' already exists and is not an empty directory.
💀 Project structure
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
# Sub-topics
* [🧬 Project Anatomy](project-anatomy.md)
* [🧳 Includes](includes.md)
* [🔤 Project Inputs](project-inputs.md)
* [⚙️ Project Configs](project-configs.md)
* [🏝️ Project Envs](project-envs.md)
<!--endTocSubTopic-->