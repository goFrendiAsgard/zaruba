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

* [initSubrepos](../../../coreTasks/initSubrepos.md)
* [addSubrepo](../../../coreTasks/addSubrepo.md)
* [pullSubrepos](../../../coreTasks/pullSubrepos.md)
* [pushSubrepos](../../../coreTasks/pushSubrepos.md)


# Add Subrepo

To add subrepo, you can perform:

```
zaruba please addSubrepo subrepoUrl="<subrepo-url>" subrepoPrefix="<subrepo-directory>" 
zaruba please pullSubrepos 

```

__Example:__

Suppose you want to create a zaruba project, and add [git@github.com:state-alchemists/fibonacci-clock.git](https://github.com/state-alchemists/fibonacci-clock) to your project, then you can do:


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
🤖 🔎 Job Starting...
         Elapsed Time: 1.133µs
         Current Time: 09:11:15
🤖 🏁 Running 🚧 initProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🚧 initProject          Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/.git/
🤖    🚀 🚧 initProject          🎉🎉🎉
🤖    🚀 🚧 initProject          Project created
🤖 🎉 Successfully running 🚧 initProject runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 12.839581ms
         Current Time: 09:11:15
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 316.354308ms
         Current Time: 09:11:15
zaruba please initProject  
🤖 🔎 Job Starting...
         Elapsed Time: 1.924µs
         Current Time: 09:11:16
🤖 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
🤖    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
🤖 🏁 Running 🥂 addSubrepo runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🥂 addSubrepo           🎉🎉🎉
🤖    🚀 🥂 addSubrepo           Subrepo fibo has been added
🤖 🎉 Successfully running 🥂 addSubrepo runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 335.556296ms
         Current Time: 09:11:16
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 639.067843ms
         Current Time: 09:11:16
zaruba please addSubrepo  -v 'subrepoUrl=git@github.com:state-alchemists/fibonacci-clock.git' -v 'subrepoPrefix=fibo' -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
🤖 🔎 Job Starting...
         Elapsed Time: 1.063µs
         Current Time: 09:11:16
🤖 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
🤖    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
🤖 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
🤖 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
🤖 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 📦 initSubrepos         fibo origin does not exist
🤖    🚀 📦 initSubrepos         [master (root-commit) ebd97ad] 🤖 Save works before pulling from git@github.com:state-alchemists/fibonacci-clock.git
🤖    🚀 📦 initSubrepos          3 files changed, 131 insertions(+)
🤖    🚀 📦 initSubrepos          create mode 100644 .gitignore
🤖    🚀 📦 initSubrepos          create mode 100644 default.values.yaml
🤖    🚀 📦 initSubrepos          create mode 100644 index.zaruba.yaml
🤖    🚀 📦 initSubrepos         git fetch fibo master
🤖 🔥 🚀 📦 initSubrepos         warning: no common commits
🤖 🔥 🚀 📦 initSubrepos         From github.com:state-alchemists/fibonacci-clock
🤖 🔥 🚀 📦 initSubrepos          * branch            master     -> FETCH_HEAD
🤖 🔥 🚀 📦 initSubrepos          * [new branch]      master     -> fibo/master
🤖 🔥 🚀 📦 initSubrepos         Added dir 'fibo'
🤖 🔥 🚀 📦 initSubrepos         From github.com:state-alchemists/fibonacci-clock
🤖 🔥 🚀 📦 initSubrepos          * branch            master     -> FETCH_HEAD
🤖 🔥 🚀 📦 initSubrepos         From github.com:state-alchemists/fibonacci-clock
🤖 🔥 🚀 📦 initSubrepos          * branch            master     -> FETCH_HEAD
🤖    🚀 📦 initSubrepos         Already up to date.
🤖    🚀 📦 initSubrepos         🎉🎉🎉
🤖    🚀 📦 initSubrepos         Subrepos Initialized
🤖 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
🤖 🏁 Running 🔽 pullSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🔽 pullSubrepos         On branch master
🤖    🚀 🔽 pullSubrepos         nothing to commit, working tree clean
🤖 🔥 🚀 🔽 pullSubrepos         From github.com:state-alchemists/fibonacci-clock
🤖 🔥 🚀 🔽 pullSubrepos          * branch            master     -> FETCH_HEAD
🤖    🚀 🔽 pullSubrepos         Already up to date.
🤖    🚀 🔽 pullSubrepos         🎉🎉🎉
🤖    🚀 🔽 pullSubrepos         Subrepos pulled
🤖 🎉 Successfully running 🔽 pullSubrepos runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 12.083083401s
         Current Time: 09:11:28
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 12.388002568s
         Current Time: 09:11:29
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


After performing the task, you will see `fibo` directory in your project.

# Pull from subrepos

People might contribute to your subrepos. You want any changes in your subrepo is also reflected in your zaruba project. In that case you need to pull from subrepos.

To pull from your subrepos, you can invoke:

```
zaruba please pullSubrepos
```

__Example:__


```bash
cd examples/playground/use-cases/externalRepositories
zaruba please pullSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.721µs
         Current Time: 09:11:29
🤖 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
🤖    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
🤖    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
🤖 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
🤖 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
🤖    🚀 📦 initSubrepos         🎉🎉🎉
🤖    🚀 📦 initSubrepos         Subrepos Initialized
🤖 🏁 Running 🔽 pullSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🔽 pullSubrepos         On branch master
🤖    🚀 🔽 pullSubrepos         nothing to commit, working tree clean
🤖 🔥 🚀 🔽 pullSubrepos         From github.com:state-alchemists/fibonacci-clock
🤖 🔥 🚀 🔽 pullSubrepos          * branch            master     -> FETCH_HEAD
🤖    🚀 🔽 pullSubrepos         Already up to date.
🤖    🚀 🔽 pullSubrepos         🎉🎉🎉
🤖    🚀 🔽 pullSubrepos         Subrepos pulled
🤖 🎉 Successfully running 🔽 pullSubrepos runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 2.995560625s
         Current Time: 09:11:32
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 3.30101991s
         Current Time: 09:11:32
zaruba please pullSubrepos  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
```````
</details>


# Push to subrepos

Sometime you need any changes in your project to be reflected in your subrepos. In that case, you can push to subrepos.

To push to your subrepos, you can invoke:

```
zaruba please pushSubrepos
```

__Example:__


```bash
cd examples/playground/use-cases/externalRepositories
zaruba please pushSubrepos
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.324µs
         Current Time: 09:11:32
🤖 🏁 Running 🔗 updateProjectLinks runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🏁 Running 🔎 zrbIsProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🔎 zrbIsProject         Current directory is a valid zaruba project
🤖 🏁 Running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖 🎉 Successfully running 🔎 zrbIsProject runner (Attempt 1 of 3)
🤖    🚀 🔗 updateProjectLinks   🎉🎉🎉
🤖    🚀 🔗 updateProjectLinks   Links updated
🤖 🎉 Successfully running 🔗 updateProjectLinks runner (Attempt 1 of 3)
🤖    🚀 🔍 zrbIsValidSubrepos   All Subrepos are valid
🤖 🎉 Successfully running 🔍 zrbIsValidSubrepos runner (Attempt 1 of 3)
🤖 🏁 Running 📦 initSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 📦 initSubrepos         🎉🎉🎉
🤖    🚀 📦 initSubrepos         Subrepos Initialized
🤖 🎉 Successfully running 📦 initSubrepos runner (Attempt 1 of 3)
🤖 🏁 Running 🔼 pushSubrepos runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories
🤖    🚀 🔼 pushSubrepos         On branch master
🤖    🚀 🔼 pushSubrepos         nothing to commit, working tree clean
🤖    🚀 🔼 pushSubrepos         git push using:  fibo master
🤖 🔥 🚀 🔼 pushSubrepos         1/3 (0) [0]
2/3 (0) [0]
3/3 (0) [0]
3/3 (1) [1]
3/3 (1) [2]
Everything up-to-date
🤖    🚀 🔼 pushSubrepos         🎉🎉🎉
🤖    🚀 🔼 pushSubrepos         Subrepos pushed
🤖 🎉 Successfully running 🔼 pushSubrepos runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 2.990128106s
         Current Time: 09:11:35
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 3.296257184s
         Current Time: 09:11:36
zaruba please pushSubrepos  -v '/home/gofrendi/zaruba/docs/examples/playground/use-cases/externalRepositories/default.values.yaml'
```````
</details>



<!--startTocSubtopic-->

<!--endTocSubtopic-->