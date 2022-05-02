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
         Elapsed Time: 4.2µs
         Current Time: 22:31:24
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 22:31:25.198 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 22:31:25.212 🎉🎉🎉
💀    🚀 initProject          🚧 22:31:25.212 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 763.3709ms
         Current Time: 22:31:25
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 965.6158ms
         Current Time: 22:31:25
zaruba please initProject  
zaruba please setProjectValue defaultBranch master -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["setProjectValue","defaultBranch","master"]
🔥 Stderr    : value of input variable 'variableName' does not match '^.+$': 
💀 🔎 Job Starting...
         Elapsed Time: 6.2µs
         Current Time: 22:31:26
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 22:31:27.556 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 22:31:28.386 🎉🎉🎉
💀    🚀 addSubrepo           🥂 22:31:28.386 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 1.6736292s
         Current Time: 22:31:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 1.8754584s
         Current Time: 22:31:28
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 3.3µs
         Current Time: 22:31:29
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 22:31:29.93  All Subrepos are valid
💀    🚀 zrbIsProject         🔎 22:31:29.931 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 22:31:30.496 fibo origin is not exist
💀    🚀 initSubrepos         📦 22:31:30.517 [master (root-commit) cba801e] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 22:31:30.517  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 22:31:30.517  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 22:31:30.517  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 22:31:30.517  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 22:31:30.588 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 22:31:33.412 warning: no common commits
💀 🔥 🚀 initSubrepos         📦 22:31:33.914 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 22:31:33.915  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 22:31:33.916  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 22:31:33.97  Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 22:31:36.482 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 22:31:36.482  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 22:31:39.268 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 22:31:39.268  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 22:31:39.533 Already up to date.
💀    🚀 initSubrepos         📦 22:31:39.537 🎉🎉🎉
💀    🚀 initSubrepos         📦 22:31:39.537 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 22:31:40.364 On branch master
💀    🚀 pullSubrepos         🔽 22:31:40.364 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 22:31:42.837 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 22:31:42.837  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 22:31:43.106 Already up to date.
💀    🚀 pullSubrepos         🔽 22:31:43.107 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 22:31:43.107 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 13.860632s
         Current Time: 22:31:43
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.0615724s
         Current Time: 22:31:43
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
         Elapsed Time: 2.2µs
         Current Time: 22:31:43
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 22:31:44.272 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 22:31:44.273 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 22:31:44.814 🎉🎉🎉
💀    🚀 initSubrepos         📦 22:31:44.814 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 22:31:45.285 On branch master
💀    🚀 pullSubrepos         🔽 22:31:45.285 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 22:31:47.803 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 22:31:47.803  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 22:31:48.077 Already up to date.
💀    🚀 pullSubrepos         🔽 22:31:48.078 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 22:31:48.078 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 4.2712457s
         Current Time: 22:31:48
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.4733561s
         Current Time: 22:31:48
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
         Elapsed Time: 2.3µs
         Current Time: 22:31:48
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 updateProjectLinks   🔗 22:31:49.205 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 22:31:49.205 Links updated
💀    🚀 zrbIsProject         🔎 22:31:49.211 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 22:31:49.268 All Subrepos are valid
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 22:31:49.948 🎉🎉🎉
💀    🚀 initSubrepos         📦 22:31:49.948 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 22:31:50.424 On branch master
💀    🚀 pushSubrepos         🔼 22:31:50.424 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 22:31:50.452 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 22:31:53.424 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 22:31:53.425 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 22:31:53.425 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 4.6916129s
         Current Time: 22:31:53
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.8934313s
         Current Time: 22:31:53
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->