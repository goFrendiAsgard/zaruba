# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is technology agnostic artefact generator. Zaruba take care about file dependencies and will do necessary action to maintain consistency in your project.

# Concept

Zaruba's main purpose is to help you generate, test, and probably deploy things faster.

Let's say you do a bit modification in your repository. Zaruba will automatically tell you whether your modification pass the unit-test or not. Then, it will copy the repository to dependant-services. Finally, it will do some service-test and integration test. Whenever everything is okay, it will generate docker images and deploy them to your kubernetes-development-cluster.

Using Zaruba, you can get that level of automation.

In order to achieve the purpose, zaruba needs `template` and `dependency tree`.

## Template

Template is basically bunch of text files. A `template` might contains `zaruba.template.yaml`.

Below is an example of valid template config:

```yaml
# base mode, invoked by performing `zaruba create <template> <target>` or `zaruba create <template> <target> interactively`
base:
  # list of actions to be performed before `copy` and `copy-and-substitute`
  pre-triggers: []
  # every file in this list will be copy-pasted from `template` to `target` without any modification.
  copy:
    readme.txt: readme.txt
    zaruba.ignore: zaruba.ignore
  # pairs of key-value that will be use for substitution in `copy-and-substitute` action. You can override the values by using environment variable, or on runtime by adding `interactively` as the last argument of the invoked command.
  substitutions:
    sender: default-sender@gmail.com
    receiver: default-receiver@gmail.com
  # every file in this list will be copy-pasted from `template` to `target`. However, it will also perform substitution as needed.
  copy-and-substitute:
    email/email.txt: email/email.txt
  # list of actions to be performed after `copy` and `copy-and-substitute`
  post-triggers:
    - 'echo "hello world" > hello.txt'
    - git init
# special mode, inherited from base, invoked by performing or `zaruba create <template>:special <target>` or `zaruba create <template>:special <target> interactively`
special:
  copy:
    special.txt: special/special.txt
```

## Project

Project is zaruba's workspace. Zaruba will watch every changes you made in a project and perform several actions. You can provide the action by editing `dependency tree`.

A project might contains `zaruba.ignore` containing list of directory that should be ignored by `zaruba`.

## Dependency Tree (Hook File)

At root path of your project, you can have `zaruba.hook.yaml`.

Below is an example of valid hook-file:

```yaml
repos/ml-classifier:
    pre-triggers:
        - python -m pytest repos/ml-classifier
    copy-to:
        - services/ner/repo/model
        - services/sentiment-analysis/repo/model
    post-triggers: []
services/ner:
    post-triggers:
        - python -m pytest services/ner
        - docker build -t gofrendi/ner-service services/ner
services/sentiment-analysis/:
    post-triggers:
        - python -m pytest services/sentiment-analysis
        - docker build -t gofrendi/sentiment-analysis-service services/sentiment-analysis
```

You have two services `services/ner` and `services/sentiment-analysis`. These services need machine-learning model from `repos/ml-classifier`.

Everytime `repos/ml-classifier` edited, you want  both `services/ner` and `services/sentiment-analysis` are updated as well.

Finally, you want to run test and create docker image whenever those services updated.

# Command

## Create

```
zaruba create <template> <target>
```

Zaruba will copy a `template` into `target`. Depends on template's configuration, every `{{keyword}}` will be replaced by `keyword` envvar value.

## Watch

```
zaruba watch [project]
```

Zaruba will watch over your project. Detect any changes in your files, and perform necessary actions to maintain consistency.

For example, if you change a file in a repository, any services depend on the repository will be updated as well.

# Configuration

## Environment Variable

* `ZARUBA_TEMPLATE_DIR`
    - Zaruba's template directory
    - Default to `<zaruba-parent-dir>/templates`
* `ZARUBA_SHELL`
    - Shell to perform commands
    - Default to `/bin/sh`
* `ZARUBA_SHELL_ARGS`
    - Shell arguments of `ZARUBA_SHELL`
    - Default to `-c`