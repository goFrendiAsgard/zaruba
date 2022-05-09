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
         Elapsed Time: 1.705µs
         Current Time: 09:32:56
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initProject          🚧 09:32:56.153 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.git/
💀    🚀 initProject          🚧 09:32:56.158 🎉🎉🎉
💀    🚀 initProject          🚧 09:32:56.158 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 115.411847ms
         Current Time: 09:32:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 316.038848ms
         Current Time: 09:32:56
zaruba please initProject  
zaruba please setProjectValue defaultBranch master -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
🔥 Command   : zaruba please
🔥 Arguments : ["setProjectValue","defaultBranch","master"]
🔥 Stderr    : value of input variable 'variableName' does not match '^.+$': 
💀 🔎 Job Starting...
         Elapsed Time: 1.946µs
         Current Time: 09:32:56
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 09:32:56.841 Current directory is a valid zaruba project
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 🥂 'addSubrepo' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 addSubrepo           🥂 09:32:56.96  🎉🎉🎉
💀    🚀 addSubrepo           🥂 09:32:56.96  Subrepo fibo has been added
💀 🎉 Successfully running 🥂 'addSubrepo' command
💀 🔎 Job Running...
         Elapsed Time: 226.810848ms
         Current Time: 09:32:57
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 428.006965ms
         Current Time: 09:32:57
zaruba please addSubrepo -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
💀 🔎 Job Starting...
         Elapsed Time: 1.265µs
         Current Time: 09:32:57
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 09:32:57.434 Current directory is a valid zaruba project
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsValidSubrepos   🔍 09:32:57.437 All Subrepos are valid
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 09:32:57.696 fibo origin is not exist
💀    🚀 initSubrepos         📦 09:32:57.714 [master (root-commit) b673ac5] 💀 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
💀    🚀 initSubrepos         📦 09:32:57.714  3 files changed, 92 insertions(+)
💀    🚀 initSubrepos         📦 09:32:57.715  create mode 100644 .gitignore
💀    🚀 initSubrepos         📦 09:32:57.715  create mode 100644 default.values.yaml
💀    🚀 initSubrepos         📦 09:32:57.715  create mode 100644 index.zaruba.yaml
💀    🚀 initSubrepos         📦 09:32:57.741 git fetch fibo master
💀 🔥 🚀 initSubrepos         📦 09:33:01.966 warning: no common commits
💀 🔥 🚀 initSubrepos         📦 09:33:02.433 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 09:33:02.433  * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 09:33:02.433  * [new branch]      master     -> fibo/master
💀 🔥 🚀 initSubrepos         📦 09:33:02.451 Added dir 'fibo'
💀 🔥 🚀 initSubrepos         📦 09:33:06.11  From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 09:33:06.11   * branch            master     -> FETCH_HEAD
💀 🔥 🚀 initSubrepos         📦 09:33:10.37  From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 initSubrepos         📦 09:33:10.37   * branch            master     -> FETCH_HEAD
💀    🚀 initSubrepos         📦 09:33:10.763 Already up to date.
💀    🚀 initSubrepos         📦 09:33:10.763 🎉🎉🎉
💀    🚀 initSubrepos         📦 09:33:10.763 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 09:33:10.882 On branch master
💀    🚀 pullSubrepos         🔽 09:33:10.882 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 09:33:13.567 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 09:33:13.567  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 09:33:13.821 Already up to date.
💀    🚀 pullSubrepos         🔽 09:33:13.821 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 09:33:13.821 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 16.495287762s
         Current Time: 09:33:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 16.696917375s
         Current Time: 09:33:14
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
         Elapsed Time: 1.173µs
         Current Time: 09:33:14
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 09:33:14.295 Current directory is a valid zaruba project
💀    🚀 zrbIsValidSubrepos   🔍 09:33:14.295 All Subrepos are valid
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 09:33:14.557 🎉🎉🎉
💀    🚀 initSubrepos         📦 09:33:14.557 Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔽 'pullSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pullSubrepos         🔽 09:33:14.682 On branch master
💀    🚀 pullSubrepos         🔽 09:33:14.682 nothing to commit, working tree clean
💀 🔥 🚀 pullSubrepos         🔽 09:33:19.956 From github.com:state-alchemists/fibonacci-clock
💀 🔥 🚀 pullSubrepos         🔽 09:33:19.956  * branch            master     -> FETCH_HEAD
💀    🚀 pullSubrepos         🔽 09:33:20.488 Already up to date.
💀    🚀 pullSubrepos         🔽 09:33:20.488 🎉🎉🎉
💀    🚀 pullSubrepos         🔽 09:33:20.488 Subrepos pulled
💀 🎉 Successfully running 🔽 'pullSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 6.299372084s
         Current Time: 09:33:20
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.500796144s
         Current Time: 09:33:20
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
         Elapsed Time: 1.345µs
         Current Time: 09:33:20
💀 🏁 Run 🔎 'zrbIsProject' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔍 'zrbIsValidSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀 🏁 Run 🔗 'updateProjectLinks' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 zrbIsProject         🔎 09:33:20.961 Current directory is a valid zaruba project
💀    🚀 updateProjectLinks   🔗 09:33:20.961 🎉🎉🎉
💀    🚀 updateProjectLinks   🔗 09:33:20.961 Links updated
💀    🚀 zrbIsValidSubrepos   🔍 09:33:20.962 All Subrepos are valid
💀 🎉 Successfully running 🔗 'updateProjectLinks' command
💀 🎉 Successfully running 🔎 'zrbIsProject' command
💀 🎉 Successfully running 🔍 'zrbIsValidSubrepos' command
💀 🏁 Run 📦 'initSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 initSubrepos         📦 09:33:21.22  🎉🎉🎉
💀    🚀 initSubrepos         📦 09:33:21.22  Subrepos Initialized
💀 🎉 Successfully running 📦 'initSubrepos' command
💀 🏁 Run 🔼 'pushSubrepos' command on /home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories
💀    🚀 pushSubrepos         🔼 09:33:21.336 On branch master
💀    🚀 pushSubrepos         🔼 09:33:21.336 nothing to commit, working tree clean
💀    🚀 pushSubrepos         🔼 09:33:21.347 git push using:  fibo master
💀 🔥 🚀 pushSubrepos         🔼 09:33:27.527 1/3 (0) [0]2/3 (0) [0]3/3 (0) [0]3/3 (1) [1]3/3 (1) [2]Everything up-to-date
💀    🚀 pushSubrepos         🔼 09:33:27.528 🎉🎉🎉
💀    🚀 pushSubrepos         🔼 09:33:27.528 Subrepos pushed
💀 🎉 Successfully running 🔼 'pushSubrepos' command
💀 🔎 Job Running...
         Elapsed Time: 6.675039961s
         Current Time: 09:33:27
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 6.875505066s
         Current Time: 09:33:27
zaruba please pushSubrepos -e '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/.env' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/external-repositories/default.values.yaml'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->