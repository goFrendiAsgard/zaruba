<!--startTocHeader-->
[🏠](../../../README.md) > [👷🏽 Use Cases](../../README.md) > [📦 Add Resources](../README.md) > [🧩 Integration](README.md)
# 📦 External Repository
<!--endTocHeader-->


At some point, you might need to add external repository into your monorepo project.

To do this you need to either use:

* [git submodule](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
* [git subrepo](https://github.com/ingydotnet/git-subrepo), or
* [git subtree](https://www.atlassian.com/git/tutorials/git-subtree)

Under the hood, Zaruba use `git subtree` since it is likely available in every git client.

All external repo will be treated as subrepo.

# Related Task

There are several builtin tasks you can use to manage subrepo:

* [initSubrepos](../../../core-tasks/initSubrepos.md)
* [addSubrepo](../../../core-tasks/addSubrepo.md)
* [pullSubrepos](../../../core-tasks/pullSubrepos.md)
* [pushSubrepos](../../../core-tasks/pushSubrepos.md)


# Add Subrepo

To add subrepo, you can perform:

```
zaruba please addSubrepo subrepoUrl="<subrepo-url>" subrepoPrefix="<subrepo-directory>" 
zaruba please pullSubrepos 

```

__Example:__

Suppose you want to create a zaruba project, and add [git@github.com:state-alchemists/fibonacci-clock.git](https://github.com/state-alchemists/fibonacci-clock) to your project, then you can do:

<!--startCode-->
```bash
# Create a Zaruba project
mkdir -p examples/playground/use-cases/externalRepositories
cd examples/playground/use-cases/externalRepositories
zaruba please initProject

# Set default branch to master
zaruba project setValue defaultBranch master

# Add subrepo and pull
zaruba please addSubrepo subrepoUrl="git@github.com:state-alchemists/fibonacci-clock.git" subrepoPrefix="fibo" 
zaruba please pullSubrepos 

# See the directory structure
tree
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.782µs
         Current Time: 13:08:56
💀 🏁 Running 🚧 initProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories is a zaruba project.
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
💀 🏁 Running 🚧 initProject runner (Attempt 2 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories is a zaruba project.
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
💀 🏁 Running 🚧 initProject runner (Attempt 3 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🚧 initProject          /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories is a zaruba project.
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
         Elapsed Time: 2.321801283s
         Current Time: 13:08:59
zaruba please initProject  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["initProject"]
🔥 Stderr    : exit status 1
💀 🔎 Job Starting...
         Elapsed Time: 1.941µs
         Current Time: 13:08:59
💀 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
💀 🏁 Running 🥂 addSubrepo runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🥂 addSubrepo           🎉🎉🎉
💀    🚀 🥂 addSubrepo           Subrepo fibo has been added
💀 🎉 Successfully running 🥂 addSubrepo runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 312.327762ms
         Current Time: 13:08:59
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 618.278482ms
         Current Time: 13:09:00
zaruba please addSubrepo  -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.721µs
         Current Time: 13:09:00
💀 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
💀    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
💀 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
💀 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 📦 initSubrepos         🎉🎉🎉
💀    🚀 📦 initSubrepos         Subrepos Initialized
💀 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
💀 🏁 Running 🔽 pullSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔽 pullSubrepos         On branch master
💀    🚀 🔽 pullSubrepos         nothing to commit, working tree clean
💀 🔥 🚀 🔽 pullSubrepos         From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 🔽 pullSubrepos          * branch            master     -> FETCH_HEAD
💀    🚀 🔽 pullSubrepos         Already up to date.
💀    🚀 🔽 pullSubrepos         🎉🎉🎉
💀    🚀 🔽 pullSubrepos         Subrepos pulled
💀 🎉 Successfully running 🔽 pullSubrepos runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 4.414127787s
         Current Time: 13:09:04
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.718567663s
         Current Time: 13:09:04
zaruba please pullSubrepos  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
.
├── default.values.yaml
├── fibo
│   ├── Dockerfile
│   ├── README.md
│   ├── bootstrap.unity.css
│   ├── index.css
│   ├── index.html
│   ├── index.js
│   ├── jquery.js
│   ├── sample.env
│   └── start.sh
├── index.zaruba.yaml
└── logs
    └── log.zaruba.csv

2 directories, 12 files
```````
</details>
<!--endCode-->

After performing the task, you will see `fibo` directory in your project.

# Pull from subrepos

People might contribute to your subrepos. You want any changes in your subrepo is also reflected in your zaruba project. In that case you need to pull from subrepos.

To pull from your subrepos, you can invoke:

```
zaruba please pullSubrepos
```

__Example:__

<!--startCode-->
```bash
cd examples/playground/use-cases/externalRepositories
zaruba please pullSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.442µs
         Current Time: 13:09:05
💀 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
💀    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
💀 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
💀 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
💀 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 📦 initSubrepos         🎉🎉🎉
💀    🚀 📦 initSubrepos         Subrepos Initialized
💀 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
💀 🏁 Running 🔽 pullSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔽 pullSubrepos         On branch master
💀    🚀 🔽 pullSubrepos         nothing to commit, working tree clean
💀 🔥 🚀 🔽 pullSubrepos         From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 🔽 pullSubrepos          * branch            master     -> FETCH_HEAD
💀    🚀 🔽 pullSubrepos         Already up to date.
💀    🚀 🔽 pullSubrepos         🎉🎉🎉
💀    🚀 🔽 pullSubrepos         Subrepos pulled
💀 🎉 Successfully running 🔽 pullSubrepos runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 4.572478984s
         Current Time: 13:09:09
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.877515681s
         Current Time: 13:09:10
zaruba please pullSubrepos  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
```````
</details>
<!--endCode-->

# Push to subrepos

Sometime you need any changes in your project to be reflected in your subrepos. In that case, you can push to subrepos.

To push to your subrepos, you can invoke:

```
zaruba please pushSubrepos
```

__Example:__

<!--startCode-->
```bash
cd examples/playground/use-cases/externalRepositories
zaruba please pushSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.74µs
         Current Time: 13:09:10
💀 🏁 Running 🔗 updateProjectLinks runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔗 updateProjectLinks   🎉🎉🎉
💀    🚀 🔗 updateProjectLinks   Links updated
💀 🎉 Successfully running 🔗 updateProjectLinks runner (Attempt 1 of 3)
💀 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
💀 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
💀    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
💀    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
💀 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
💀    🚀 📦 initSubrepos         🎉🎉🎉
💀    🚀 📦 initSubrepos         Subrepos Initialized
💀 🏁 Running 🔼 pushSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
💀    🚀 🔼 pushSubrepos         On branch master
💀    🚀 🔼 pushSubrepos         nothing to commit, working tree clean
💀    🚀 🔼 pushSubrepos         git push using:  fibo master
💀 🔥 🚀 🔼 pushSubrepos         1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 🔼 pushSubrepos         🎉🎉🎉
💀    🚀 🔼 pushSubrepos         Subrepos pushed
💀 🎉 Successfully running 🔼 pushSubrepos runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 4.598054469s
         Current Time: 13:09:14
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.902647472s
         Current Time: 13:09:15
zaruba please pushSubrepos  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->