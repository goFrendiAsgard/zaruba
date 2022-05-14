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
         Elapsed Time: 1.197µs
         Current Time: 08:48:10
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 08:48:10.041 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 08:48:10.047 🎉🎉🎉
💀    🚀 initProject          🚧 08:48:10.047 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 115.678406ms
         Current Time: 08:48:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 316.088976ms
         Current Time: 08:48:10
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.324µs
         Current Time: 08:48:10
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:48:10.538 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 08:48:10.659 🎉🎉🎉
💀    🚀 addSubrepo           🥂 08:48:10.659 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 230.313788ms
         Current Time: 08:48:10
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 431.113702ms
         Current Time: 08:48:10
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.18µs
         Current Time: 08:48:11
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:48:11.122 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:48:11.123 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:48:11.392 fibo origin is not exist
💀    🚀 initSubrepos         📦 08:48:11.403 [master (root-commit) 317a0f9] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 08:48:11.403  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 08:48:11.403  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 08:48:11.403  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 08:48:11.403  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 08:48:11.436 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 08:48:14.62  warning: no common commits
💀 🔥 🚀 initSubrepos         📦 08:48:15.101 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:48:15.101  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:48:15.102  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 08:48:15.12  Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 08:48:17.805 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:48:17.805  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:48:20.974 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:48:20.974  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 08:48:21.252 Already up to date.
💀    🚀 initSubrepos         📦 08:48:21.252 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:48:21.252 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:48:21.372 On branch master
💀    🚀 pullSubrepos         🔽 08:48:21.372 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:48:24.136 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:48:24.136  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:48:24.453 Already up to date.
💀    🚀 pullSubrepos         🔽 08:48:24.454 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:48:24.454 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 13.437550414s
         Current Time: 08:48:24
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 13.637925233s
         Current Time: 08:48:24
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
         Elapsed Time: 1.119µs
         Current Time: 08:48:24
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:48:24.926 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:48:24.926 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:48:25.186 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:48:25.186 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:48:25.303 On branch master
💀    🚀 pullSubrepos         🔽 08:48:25.303 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:48:28.023 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:48:28.023  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:48:28.296 Already up to date.
💀    🚀 pullSubrepos         🔽 08:48:28.297 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:48:28.297 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.477957579s
         Current Time: 08:48:28
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.679497121s
         Current Time: 08:48:28
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
         Elapsed Time: 1.095µs
         Current Time: 08:48:28
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 08:48:28.78  All Subrepos are valid
💀    🚀 zrbIsProject         🔎 08:48:28.78  Current directory is a valid zaruba project
💀    🚀 updateProjectLinks   🔗 08:48:28.78  🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:48:28.78  Links updated
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:48:29.039 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:48:29.04  Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 08:48:29.159 On branch master
💀    🚀 pushSubrepos         🔼 08:48:29.159 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 08:48:29.17  git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 08:48:32.185 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 08:48:32.185 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 08:48:32.186 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.513210053s
         Current Time: 08:48:32
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.714402459s
         Current Time: 08:48:32
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->