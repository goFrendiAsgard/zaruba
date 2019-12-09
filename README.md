# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is agnostic generator and task runner. It sole purpose is to help you create project and maintain dependencies among components.

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

* `install-template`: Any executable script, containing set of commands to be executed after user install the template. You might find some `npm init` or `pip install` here.
* `create-component`: Any executable script, containing set of commands to be executed when user creates new component based on current template. `create-component` should at least a single argument containing project directory.

To install a template, you can perform:

```sh
zaruba install-template https://github.com/someUser/someTemplate.git
```

## Project

Project is a directory containing set of components. A project might also be a component on it's own.

## Component

Component can be anything from a project, a shared library, or a single service.

A component is usually based on specific template, but user can also create their own components from scratch. Also, a component should contains at least a single file:

* `link`: Any executable script, containing set of commands to be executed when user perform `zaruba organize` or `zaruba watch`.

Optionally, a component can also has `organize-project` or any other shell script for custome command.

To create a new component, you can perform:

```sh
zaruba create-component someTemplate
```

# Commmands

## install-template

```sh
zaruba install-template <template-git-url> [folder-name]
```

This one basically run `git clone <template-gir-url>` and executing `install-template`.

While running `install-template`, current working directory is set to `[folder-name]`. However, if `[folder-name]` is not specified, zaruba will use `<template-git-url>`'s repository name as `[folder-name]`.

Running `zaruba install-template <template-git-url> [folder-name]` should has the same effect as performing:

```sh
git clone ${template_git_url} ${zaruba_template_dir}/${folder_name}
cd ${zaruba_template_dir}/${folder_name}
./install_template
```

## create-component

```
zaruba create-component <template> [project-dir] [...args]
```

This will run template's `create-component <project-dir> [...args]`. Typically, it should create new component based on `<template>`. It is assumed that current working directory is pointing to `<template>`.

Running `zaruba create-component <template> [project-dir] [...args]` should has the same effect as performing:

```
cd ${zaruba_template_dir}/${template}
./create_component ${project_dir}
```

## link

```sh
zaruba link <project-dir> <source> <destination>
```

This command is usually invoked while performing `organize-project`. Usually, this command is part of `<project-dir>/.../link` and never invoked directly. By invoking this command, user should be able to add dependency to project's `zaruba.dependency.json`.

After running `zaruba-link <project-dir> <source> <destination>`, there should be a json file named `zaruba.dependency.json` in your `<project-dir>`. The file should contains all dependencies in a single project in JSON format:

```json
{
    "<source-1>" : [
        "destination-1", "destination-2", "destination-3"
    ],
    "<source-2>" : [
        "destination-1", "destination-2", "destination-3"
    ]
}
```

## organize-project

```sh
zaruba organize-project [project-dir]
```

This command will do the following actions:

* Re-create `zaruba.dependency.json` by performing `link` action in `<project-dir>` and it's sub-directories.
* Sort dependencies in `zaruba.dependency.json`.
* Copy sources to their respective destinations.
* Recursively look for and run `organize-project` in every sub-directory of `<project-dir>`.

## watch-project

```sh
zaruba watch-project [project-dir]
```

Detect changes in project and `organize-project` automatically.

## custom action

```
zaruba do <action> [...args]
```

You can add any custom action by creating a shell script in any directory of the project. The name of the script should match your custom action.

In short, when you perform `zaruba do fight`, zaruba will looks for every `fight.sh` in the current directory, and perform `fight.sh <current-directory>`.

# Configuration

## Environment Variable

* `ZARUBA_TEMPLATE_DIR`
    - Zaruba's template directory
    - Default to `<zaruba-parent-dir>/templates`