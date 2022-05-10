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
         Elapsed Time: 1.198µs
         Current Time: 07:56:11
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 07:56:11.615 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 07:56:11.621 🎉🎉🎉
💀    🚀 initProject          🚧 07:56:11.621 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 123.517928ms
         Current Time: 07:56:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 324.375746ms
         Current Time: 07:56:11
zaruba please initProject  
zaruba please setProjectValue defaultBranch master -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["setProjectValue","defaultBranch","master"]
🔥 Stderr    : value of input variable 'variableName' does not match '^.+$': 
💀 🔎 Job Starting...
         Elapsed Time: 1.518µs
         Current Time: 07:56:12
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 07:56:12.549 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 07:56:12.673 🎉🎉🎉
💀    🚀 addSubrepo           🥂 07:56:12.673 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 235.580681ms
         Current Time: 07:56:12
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 436.407803ms
         Current Time: 07:56:12
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.538µs
         Current Time: 07:56:13
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 07:56:13.252 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 07:56:13.253 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 07:56:13.515 fibo origin is not exist
💀    🚀 initSubrepos         📦 07:56:13.525 [master (root-commit) ea458f6] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 07:56:13.525  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 07:56:13.525  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 07:56:13.525  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 07:56:13.525  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 07:56:13.551 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 07:56:16.98  warning: no common commits
💀 🔥 🚀 initSubrepos         📦 07:56:17.409 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 07:56:17.409  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 07:56:17.409  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 07:56:17.433 Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 07:56:20.402 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 07:56:20.402  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 07:56:23.591 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 07:56:23.591  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 07:56:23.9   Already up to date.
💀    🚀 initSubrepos         📦 07:56:23.9   🎉🎉🎉
💀    🚀 initSubrepos         📦 07:56:23.9   Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 07:56:24.025 On branch master
💀    🚀 pullSubrepos         🔽 07:56:24.025 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 07:56:26.902 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 07:56:26.902  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 07:56:27.198 Already up to date.
💀    🚀 pullSubrepos         🔽 07:56:27.199 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 07:56:27.199 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 14.057355553s
         Current Time: 07:56:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.258684031s
         Current Time: 07:56:27
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
         Elapsed Time: 1.848µs
         Current Time: 07:56:27
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 07:56:27.825 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 07:56:27.825 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 07:56:28.09  🎉🎉🎉
💀    🚀 initSubrepos         📦 07:56:28.09  Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 07:56:28.217 On branch master
💀    🚀 pullSubrepos         🔽 07:56:28.217 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 07:56:31.243 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 07:56:31.243  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 07:56:31.622 Already up to date.
💀    🚀 pullSubrepos         🔽 07:56:31.623 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 07:56:31.623 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.908725002s
         Current Time: 07:56:31
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 4.109907192s
         Current Time: 07:56:31
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
         Elapsed Time: 2.158µs
         Current Time: 07:56:32
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 07:56:32.235 All Subrepos are valid
💀    🚀 zrbIsProject         🔎 07:56:32.236 Current directory is a valid zaruba project
💀    🚀 updateProjectLinks   🔗 07:56:32.236 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 07:56:32.236 Links updated
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 07:56:32.491 🎉🎉🎉
💀    🚀 initSubrepos         📦 07:56:32.491 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 07:56:32.611 On branch master
💀    🚀 pushSubrepos         🔼 07:56:32.611 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 07:56:32.625 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 07:56:35.836 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 07:56:35.837 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 07:56:35.837 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.717513523s
         Current Time: 07:56:35
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.919817537s
         Current Time: 07:56:36
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->