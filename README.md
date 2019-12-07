# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is agnostic generator and task runner. It sole purpose is to help you create project and maintain dependencies among components.

# Concepts

## Template

Templates are component's blueprint. It can be a node.js package, python program, or even bunch of shell scripts.

A template should contains at least two files:

* `install-template.sh`: A shell script, containing set of commands to be executed after user install the template. You might find some `npm init` or `pip install` here.
* `create-component.sh`: A shell script, containing set of commands to be executed when user creates new component based on current template. `create-component.sh` should at least a single argument containing project directory.

To install a template, you can perform:

```
zaruba install-template https://github.com/someUser/someTemplate
```

## Project

Project is a directory containing set of components. A project might also be a component on it's own.

## Component

Component can be anything from a project, a shared library, or a single service.

A component is usually based on specific template, but user can also create their own components from scratch.

To create a new component, you can perform:

```
zaruba create-component someTemplate
```

# Commmands

## install-template

```
zaruba install-template <template-git-url>
```

This one basically run `git clone <template-gir-url>` and executing `install-template.sh`. While running `install-template.sh`, it is assumed that current working directory is pointing to the newly-cloned template.

## create-component

```
zaruba create-component <template> [project-dir] [--interactive | -i]
```

## link

```
zaruba link <source> <destination>
```

## organize

```
zaruba organize [project-dir]
```

## watch

```
zaruba watch [project-dir]
```