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
Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myProjectFromScratch/.git/
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
         Elapsed Time: 1.89µs
         Current Time: 23:53:08
💀 🏁 Running 🚧 initProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject
💀    🚀 🚧 initProject          Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject/.git/
💀    🚀 🚧 initProject          🎉🎉🎉
💀    🚀 🚧 initProject          Project created
💀 🎉 Successfully running 🚧 initProject runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 16.087903ms
         Current Time: 23:53:08
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 320.253997ms
         Current Time: 23:53:08
zaruba please initProject  
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
Cloning into 'myClonedProject'...
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