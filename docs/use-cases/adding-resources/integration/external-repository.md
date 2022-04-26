<!--startTocHeader-->
[🏠](../../../README.md) > [👷🏽 Use Cases](../../README.md) > [Adding Resources](../README.md) > [🧩 Integration](README.md)
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
mkdir -p examples/playground/use-cases/external-repositories
cd examples/playground/use-cases/external-repositories
zaruba please initProject

# Set default branch to master
zaruba please setProjectValue defaultBranch master

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
         Elapsed Time: 1.148µs
         Current Time: 08:31:37
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 08:31:37.564 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 08:31:37.568 🎉🎉🎉
💀    🚀 initProject          🚧 08:31:37.568 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 112.409127ms
         Current Time: 08:31:37
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 313.207718ms
         Current Time: 08:31:37
zaruba please initProject  
zaruba please setProjectValue defaultBranch master -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["setProjectValue","defaultBranch","master"]
🔥 Stderr    : value of input variable 'variableName' does not match '^.+$': 
💀 🔎 Job Starting...
         Elapsed Time: 2.224µs
         Current Time: 08:31:38
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:31:38.152 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 08:31:38.263 🎉🎉🎉
💀    🚀 addSubrepo           🥂 08:31:38.263 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 214.078882ms
         Current Time: 08:31:38
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 415.434263ms
         Current Time: 08:31:38
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.179µs
         Current Time: 08:31:38
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:31:38.706 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:31:38.706 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:31:38.968 fibo origin is not exist
💀    🚀 initSubrepos         📦 08:31:38.974 [master (root-commit) c5757de] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 08:31:38.974  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 08:31:38.974  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 08:31:38.974  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 08:31:38.974  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 08:31:38.991 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 08:31:42.22  warning: no common commits
💀 🔥 🚀 initSubrepos         📦 08:31:42.729 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:31:42.729  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:31:42.73   * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 08:31:42.742 Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 08:31:46.34  From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:31:46.34   * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:31:49.355 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:31:49.355  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 08:31:49.633 Already up to date.
💀    🚀 initSubrepos         📦 08:31:49.633 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:31:49.633 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:31:49.749 On branch master
💀    🚀 pullSubrepos         🔽 08:31:49.749 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:31:53.043 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:31:53.043  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:31:53.362 Already up to date.
💀    🚀 pullSubrepos         🔽 08:31:53.362 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:31:53.362 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 14.759814222s
         Current Time: 08:31:53
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.961888864s
         Current Time: 08:31:53
zaruba please pullSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
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
└── log.zaruba.csv

1 directory, 12 files
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
cd examples/playground/use-cases/external-repositories
zaruba please pullSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.149µs
         Current Time: 08:31:53
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 08:31:53.834 All Subrepos are valid
💀    🚀 zrbIsProject         🔎 08:31:53.834 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:31:54.094 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:31:54.094 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:31:54.207 On branch master
💀    🚀 pullSubrepos         🔽 08:31:54.208 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:31:57.186 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:31:57.186  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:31:57.456 Already up to date.
💀    🚀 pullSubrepos         🔽 08:31:57.456 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:31:57.456 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.72661613s
         Current Time: 08:31:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.927787664s
         Current Time: 08:31:57
zaruba please pullSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
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
cd examples/playground/use-cases/external-repositories
zaruba please pushSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.151µs
         Current Time: 08:31:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 08:31:57.944 All Subrepos are valid
💀    🚀 zrbIsProject         🔎 08:31:57.944 Current directory is a valid zaruba project
💀    🚀 updateProjectLinks   🔗 08:31:57.944 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:31:57.944 Links updated
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:31:58.206 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:31:58.206 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 08:31:58.321 On branch master
💀    🚀 pushSubrepos         🔼 08:31:58.321 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 08:31:58.333 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 08:32:01.41  1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 08:32:01.41  🎉🎉🎉
💀    🚀 pushSubrepos         🔼 08:32:01.41  Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.570797723s
         Current Time: 08:32:01
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.772492692s
         Current Time: 08:32:01
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->