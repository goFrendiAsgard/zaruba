# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is agnostic generator and task runner. It sole purpose is to help you create project and maintain dependencies among components.

# Concepts

## Template

Templates are component's blueprint. It can be a node.js package, python program, or even bunch of shell scripts.

A template should contains at least two files:

* `install-template`: A shell script, containing set of commands to be executed after user install the template. You might find some `npm init` or `pip install` here.
* `create-component`: A shell script, containing set of commands to be executed when user creates new component based on current template. `create-component` should at least a single argument containing project directory.

To install a template, you can perform:

```sh
zaruba install-template https://github.com/someUser/someTemplate
```

## Project

Project is a directory containing set of components. A project might also be a component on it's own.

## Component

Component can be anything from a project, a shared library, or a single service.

A component is usually based on specific template, but user can also create their own components from scratch. Also, a component should contains at least a single file:

* `link`: A shell script, containing set of commands to be executed when user perform `zaruba orgainze` or `zaruba watch`.

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

This command is usually invoked while performing `organize-project` or `watch`. Usually, this command is part of `<project-dir>/.../link` and never invoked directly. By invoking this command, user should be able to add dependency to project's `zaruba.dependency.json`.

Running this command should has the same effect as performing:

```sh
add_dependency ${project_dir} ${source} ${destination}
```

assuming `add-dependency` is a binary executable with the following source code:

```go
package main

import (
	"encoding/json"
	"io/ioutil"
    "os"
    "path/filepath"
	"syscall"
)

func main() {
    // initiate variables, assuming all parameters are valid
    projectDir, _ := filepath.Abs(os.Args[1])
    source, _ := filepath.Abs(os.Args[2])
    destination, _ := filepath.Abs(os.Args[3])
	depFileName := filepath.Join(projectDir, "zaruba.dependency.json")
    dep := map[string][]string{}

    // create `depFileName` if it is not exists
    if _, err := os.Stat(depFileName); os.IsNotExist(err) {
        os.Create(depFileName)
        ioutil.WriteFile(depFileName, []byte("{}"), 0644)
    }

    // open `depFile`
	depFile, err := os.Open(depFileName)
	if err != nil {
		defer depFile.Close()
    }

	// lock `depFile`
    syscall.Flock(int(depFile.Fd()), syscall.LOCK_EX)

	// read `dep` from `defFileName`, assuming it is a valid json file
    b, err := ioutil.ReadFile(depFileName)
    json.Unmarshal(b, &dep)
    
    // add `source` and `destination` to `dep`
    if _, sourceExists := dep[source]; !sourceExists {
        dep[source] = []string{}
    }
    dep[source] = append(dep[source], destination)

    // write `dep` to `depFileName`, assuming operation should always success
	b, _ = json.Marshal(dep)
    ioutil.WriteFile(depFileName, b, 0644)

	// unloack `depFile`
	syscall.Flock(int(depFile.Fd()), syscall.LOCK_UN)
}
```

As you can see, after running `zaruba-link <source> <destination>`, there should be a json file name `<preject-dir>/zaruba.dependency.json`. The file should contains all dependencies in a single project in JSON format:

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

* Delete `zaruba.dependency.json`.
* Recursively look for and run `organize-project` in every sub-directory of `<project-dir>`

Invoking `zaruba organize-project [project-dir]` should have the same effect as:

```sh
cd ${project_dir}

# remove `zaruba.dependency.json`
rm zaruba.dependency.json

# collect all dependencies into `zaruba.dependency.json`
for subdir in $(find "./") 
do
    filename = "./{$subdir}/link"
    if [ -f ${filename} ]
    then
        ./${filename} ${project_dir}
    fi
done     

# perform organize-project
./organize-project 
```

assuming `organize-project` is a binary executable with the following source code:

```go
package main

import (
	"encoding/json"
	"io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "sort"
	"syscall"
)

func main() {
    // initiate variables, assuming all parameters are valid
    projectDir, _ := filepath.Abs(os.Args[1])
	depFileName := filepath.Join(projectDir, "zaruba.dependency.json")
    dep := map[string][]string{}

    // read `dep` from `defFileName`, assuming it is a valid json file
    b, err := ioutil.ReadFile(depFileName)
    json.Unmarshal(b, &dep)

    // get all keys of dep (i.e: list of sources)
    sources := []string{}
    for source := range dep {
        sources = append(sources, source)
    }

    // sort keys
    sort.SliceStable(sources, func(i int, j int) bool {
		firstSource, secondSource := sources[i], sources[j]
        // get destination
        firstDestinations = dep[firstSource]
        // compare
        for _, destination := range firstDestinations {
            if strings.HasPrefix(secondSource, destination) {
                return true
            }
        }
        return false
    }) 

    for source, destinations := range(sources) {
        for _, destination := range(destinations) {
            exec.Command("rm", "-Rf", destination)
            exec.Command("cp", "-r", source, destination)
        }
    }
    

}
```

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