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
         Elapsed Time: 2.588µs
         Current Time: 00:25:46
💀 🏁 Running 🚧 'initProject' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 initProject          🚧 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.git/
💀    🚀 initProject          🚧 🎉🎉🎉
💀    🚀 initProject          🚧 Project created
💀 🎉 Successfully running 🚧 'initProject' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 16.416864ms
         Current Time: 00:25:46
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 320.267916ms
         Current Time: 00:25:46
zaruba please initProject  
💀 🔎 Job Starting...
         Elapsed Time: 1.525µs
         Current Time: 00:25:46
💀 🏁 Running 🔎 'zrbIsProject' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' runner (Attempt: 1/3)
💀 🏁 Running 🥂 'addSubrepo' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 addSubrepo           🥂 🎉🎉🎉
💀    🚀 addSubrepo           🥂 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 335.00003ms
         Current Time: 00:25:47
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 637.704051ms
         Current Time: 00:25:47
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.845µs
         Current Time: 00:25:47
💀 🏁 Running 🔎 'zrbIsProject' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀 🏁 Running 🔍 'zrbIsValidSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 zrbIsValidSubrepos   🔍 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' runner (Attempt: 1/3)
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' runner (Attempt: 1/3)
💀 🏁 Running 📦 'initSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 initSubrepos         📦 fibo origin does not exist
💀    🚀 initSubrepos         📦 [master (root-commit) 7f2af82] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦  3 files changed, 125 insertions(+)
💀    🚀 initSubrepos         📦  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 warning: no common commits
💀 🔥 🚀 initSubrepos         📦 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 Already up to date.
💀    🚀 initSubrepos         📦 🎉🎉🎉
💀    🚀 initSubrepos         📦 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' runner (Attempt: 1/3)
💀 🏁 Running 🔽 'pullSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 pullSubrepos         🔽 On branch master
💀    🚀 pullSubrepos         🔽 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 Already up to date.
💀    🚀 pullSubrepos         🔽 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 13.739506873s
         Current Time: 00:26:01
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 14.041969947s
         Current Time: 00:26:01
zaruba please pullSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
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
         Elapsed Time: 2.248µs
         Current Time: 00:26:01
💀 🏁 Running 🔍 'zrbIsValidSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀 🏁 Running 🔎 'zrbIsProject' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 zrbIsValidSubrepos   🔍 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' runner (Attempt: 1/3)
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' runner (Attempt: 1/3)
💀 🏁 Running 📦 'initSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 initSubrepos         📦 🎉🎉🎉
💀    🚀 initSubrepos         📦 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' runner (Attempt: 1/3)
💀 🏁 Running 🔽 'pullSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 pullSubrepos         🔽 On branch master
💀    🚀 pullSubrepos         🔽 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 Already up to date.
💀    🚀 pullSubrepos         🔽 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 3.311045931s
         Current Time: 00:26:05
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.613627767s
         Current Time: 00:26:05
zaruba please pullSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
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
         Elapsed Time: 1.515µs
         Current Time: 00:26:05
💀 🏁 Running 🔎 'zrbIsProject' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀 🏁 Running 🔗 'updateProjectLinks' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀 🎉 Successfully running 🔎 'zrbIsProject' runner (Attempt: 1/3)
💀    🚀 updateProjectLinks   🔗 🎉🎉🎉
💀    🚀 zrbIsProject         🔎 Current directory is a valid zaruba project
💀    🚀 updateProjectLinks   🔗 Links updated
💀 🎉 Successfully running 🔗 'updateProjectLinks' runner (Attempt: 1/3)
💀 🏁 Running 🔍 'zrbIsValidSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 zrbIsValidSubrepos   🔍 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' runner (Attempt: 1/3)
💀 🏁 Running 📦 'initSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 initSubrepos         📦 🎉🎉🎉
💀    🚀 initSubrepos         📦 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' runner (Attempt: 1/3)
💀 🏁 Running 🔼 'pushSubrepos' runner on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories (Attempt: 1/3)
💀    🚀 pushSubrepos         🔼 On branch master
💀    🚀 pushSubrepos         🔼 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 3.373048456s
         Current Time: 00:26:08
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.675754739s
         Current Time: 00:26:09
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->