<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md)
# 🏗️ Project
<!--endTocHeader-->

A project is a directory containing `index.zaruba.yaml`. Usually, a project is also a git repository.

# Create an Empty Project

To create an empty project from scratch, you can do:

```bash
mkdir myproject
cd myproject
git init
touch index.zaruba.yaml
```

# Generate a Project

To create a project with sane boilerplate you can invoke `zaruba please initProject`:

```bash
mkdir myproject
cd myproject
zaruba please initProject
```

# Clone a Project

To clone/fork existing projects from GitHub or other git servers.

```bash
git clone git@github.com:<user>/<repo>.git
```

# Project Anatomy

Please look on [project anatomy](./project-anatomy.md) for more information. 

<!--startTocSubTopic-->
# Sub-topics
* [🧬 Project Anatomy](project-anatomy.md)
* [🧳 Includes](includes.md)
* [🔤 Project Inputs](project-inputs.md)
* [⚙️ Project Configs](project-configs.md)
* [🏝️ Project Envs](project-envs.md)
<!--endTocSubTopic-->