# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is agnostic generator and service runner. It treat your project as big monorepo application, but still allow you to publish your services as multi-repos

# Installation

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

or

```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

# Why Zaruba

There are bunch of task-runners out there, like `gulp`, `grunts`, `webpack`, and that `good-old-shell-scripts`.

Some things are better written in shell-scripts, while some others (like manipulating JSON) are better written in Javascript. Zaruba solve this problem by introducing a little set of convetions. Belows are problems Zaruba address:

## Language Lock Problem

First, Zaruba allows you to write tasks in any language. In order to achieve this, you script should has `#!` directives:

* Shell script should be started with `#!/bin/sh`
* Node script should be started with `#!/bin/node`
* Python script should be started with `#!/bin/python`

## Monolithic Task Problem

Second, you can spread up your tasks in your project directories in order to make them small, single-responsible, and maintainable.

Let's see the following examples:

```
.
├── LICENSE.md
├── Makefile
├── README.md
├── payment
│   ├── payment.go
│   └── test               # our task-script
├── recommendation
│   ├── recommendation.py
│   └── test               # our task-script
└── authentication
    ├── authentication.js
    └── test               # our task-script
```

When you perform `zaruba test`, all corresponding task script (the ones named `test`) will be executed.

## Dependency Problem

Even in the world full of microservices and isolation, you might encounter some resources has to be shared among different services. For example like proto-buff schema or even a single function. Sometime problems occured when you update a single part in a service, unaware that it will also affect other services as well.

Zaruba tries to solve dependency problem by using `zaruba link`, providing special action named `zaruba organize-project` and even full-fledge project watcher (`zaruba watch-project`).

# Concepts

## Template

Templates are component's blueprint. It can be a node.js package, python program, or even bunch of shell scripts.

A template should contains at least two files:

* `install-template.zaruba`: Any executable script, containing set of commands to be executed after user install the template. You might find some `npm init` or `pip install` here.
* `create-component.zaruba`: Any executable script, containing set of commands to be executed when user creates new component based on current template. `create-component` should at least a single argument containing project directory.

To install a template, you can perform:

```sh
zaruba install-template https://github.com/someUser/someTemplate.git
```

## Project

Project is a directory containing set of components. A project might also be a component on it's own.

## Component

Component can be anything from a project, a shared library, or a single service.

A component is usually based on specific template, but user can also create their own components from scratch. Also, a component should contains at least a single file:

* `link`: Any executable script, containing set of commands to be executed when user perform `zaruba organize-project` or `zaruba watch-project`.

Optionally, a component can also has `organize-project.zaruba` or any other shell script for custom command.

To create a new component, you can perform:

```sh
zaruba create-component someTemplate
```

## pre-action

pre-action is will be executed before an action is executed. You can make pre-action by simply create an executable file with `pre` prefix.

## post-action

post-action is will be executed after an action is executed. You can make post-action by simply create an executable file with `post` prefix.

# Commmands

## install-template

```sh
zaruba install-template <template-git-url> [folder-name]
```

This one basically run `git clone <template-gir-url>` and executing `install-template`.

While running `install-template`, current working directory is set to `[folder-name]`. However, if `[folder-name]` is not specified, zaruba will use `<template-git-url>`'s repository name as `[folder-name]`.

Running `zaruba install-template <template-git-url> [folder-name]` should has the same effect as performing:

```sh
git clone ${template_git_url} ${zaruba_template_dir}/${folder_name}.git
cd ${zaruba_template_dir}/${folder_name}
./install_template.zaruba
```

## create-component

```sh
zaruba create-component <template> [project-dir [...args]]
```

This will run template's `create-component <project-dir> [...args]`. Typically, it should create new component based on `<template>`. It is assumed that current working directory is pointing to `<template>`.

Running `zaruba create-component <template> [project-dir] [...args]` should has the same effect as performing:

```sh
cd ${zaruba_template_dir}/${template}
./create_component.zaruba ${project_dir}
```

## organize

```sh
zaruba organize [project-dir [...args]]
```

This command will do the following actions:

* Copy sources to their respective destinations.
* Recursively look for and run `organize-project.zaruba` in every sub-directory of `<project-dir>` and execute it. This command support pre-action (`pre-organize-project.zaruba`) and post-action (`post-organize-project.zaruba`).

## watch

```sh
zaruba watch [project-dir [...args]]
```

Detect changes in project and `organize-project` automatically.

## custom action

```sh
zaruba do <action> [project-dir [...args]]
```

You can add any custom action by creating a shell script in any directory of the project. The name of the script should match your custom action. Custom action also support pre-action and post-action.

In short, when you perform `zaruba do fight`, zaruba will looks for every `fight.zaruba` script in the current directory, and perform `fight.zaruba <current-directory>`. To make pre-action and post-action, you can simply create `pre-fight.zaruba` and `post-fight.zaruba`.

Note: whenever running the executables, zaruba will automatically add `<project-dir>` as first argument. The value of `<project-dir>` is taken from current working directory.

# Configuration

## Environment Variable

* `ZARUBA_TEMPLATE_DIR`: Zaruba's template directory, default to `<zaruba-parent-dir>/templates`
* `ZARUBA_SHELL`: Default to `/bin/bash`
* `ZARUBA_SHELL_ARG`: Default to `-c`

# Testing

Create `.env` based on `template.env`.

```sh
source .env
make test # or make test-verbose
```