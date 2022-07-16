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
         Elapsed Time: 1.049µs
         Current Time: 14:43:11
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject
💀    🚀 initProject          🚧 Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/myGeneratedProject/.git/
💀    🚀 initProject          🚧 🎉🎉🎉
💀    🚀 initProject          🚧 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 116.230055ms
         Current Time: 14:43:11
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 419.791639ms
         Current Time: 14:43:11
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