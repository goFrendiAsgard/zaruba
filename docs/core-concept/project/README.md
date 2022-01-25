[⬅️ Table of Content](../../README.md)

# 🏗️ Project

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