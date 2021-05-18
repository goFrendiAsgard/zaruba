# Creating a Project

A zaruba project is a git repository containing at least one file: `main.zaruba.yaml`.

In order to initialize a project, you can invoke `zaruba please initProject`

## Example

```sh
mkdir newProject
cd newProject
zaruba please initProject
```

## Involved tasks

* [initProject](tasks/initProject.md)

## What's next


* [Creating Fast API service](creating-fast-api-service.md)
* [Creating docker task](creating-docker-task.md)
* [Creating service task](creating-service-task.md)
* [Understanding task](understanding-task.md)
* [Working with legacy code](working-with-legacy-code.md)
* [Starting a monorepo](starting-a-monorepo.md)
* [Working with multirepo](working-with-multirepo.md)