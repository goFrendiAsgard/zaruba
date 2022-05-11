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
         Elapsed Time: 1.122µs
         Current Time: 08:17:26
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 08:17:26.66  Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 08:17:26.664 🎉🎉🎉
💀    🚀 initProject          🚧 08:17:26.664 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 111.158465ms
         Current Time: 08:17:26
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 313.225246ms
         Current Time: 08:17:26
zaruba please initProject  
zaruba please setProjectValue defaultBranch master -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["setProjectValue","defaultBranch","master"]
🔥 Stderr    : value of input variable 'variableName' does not match '^.+$': 
💀 🔎 Job Starting...
         Elapsed Time: 1.169µs
         Current Time: 08:17:27
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:17:27.261 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 08:17:27.376 🎉🎉🎉
💀    🚀 addSubrepo           🥂 08:17:27.376 Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 220.963693ms
         Current Time: 08:17:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 423.113208ms
         Current Time: 08:17:27
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.31µs
         Current Time: 08:17:27
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:17:27.834 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:17:27.835 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:17:28.096 fibo origin is not exist
💀    🚀 initSubrepos         📦 08:17:28.102 [master (root-commit) c4f827c] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 08:17:28.102  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 08:17:28.102  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 08:17:28.102  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 08:17:28.102  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 08:17:28.117 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 08:17:31.243 warning: no common commits
💀 🔥 🚀 initSubrepos         📦 08:17:31.745 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:17:31.745  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:17:31.746  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 08:17:31.759 Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 08:17:34.526 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:17:34.526  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 08:17:37.625 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 08:17:37.625  * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 08:17:37.943 Already up to date.
💀    🚀 initSubrepos         📦 08:17:37.943 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:17:37.943 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:17:38.061 On branch master
💀    🚀 pullSubrepos         🔽 08:17:38.061 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:17:40.912 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:17:40.912  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:17:41.214 Already up to date.
💀    🚀 pullSubrepos         🔽 08:17:41.215 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:17:41.215 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 13.486837687s
         Current Time: 08:17:41
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 13.689403105s
         Current Time: 08:17:41
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
         Elapsed Time: 1.139µs
         Current Time: 08:17:41
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 08:17:41.683 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:17:41.684 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:17:41.945 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:17:41.945 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 08:17:42.061 On branch master
💀    🚀 pullSubrepos         🔽 08:17:42.061 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 08:17:44.843 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 08:17:44.843  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 08:17:45.133 Already up to date.
💀    🚀 pullSubrepos         🔽 08:17:45.133 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 08:17:45.133 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.556665264s
         Current Time: 08:17:45
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.757217124s
         Current Time: 08:17:45
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
         Elapsed Time: 1.439µs
         Current Time: 08:17:45
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 updateProjectLinks   🔗 08:17:45.605 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 08:17:45.605 Links updated
💀    🚀 zrbIsProject         🔎 08:17:45.605 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 08:17:45.605 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 08:17:45.865 🎉🎉🎉
💀    🚀 initSubrepos         📦 08:17:45.865 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 08:17:45.981 On branch master
💀    🚀 pushSubrepos         🔼 08:17:45.981 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 08:17:45.993 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 08:17:49.069 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 08:17:49.069 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 08:17:49.069 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 3.571690654s
         Current Time: 08:17:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 3.773215299s
         Current Time: 08:17:49
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->