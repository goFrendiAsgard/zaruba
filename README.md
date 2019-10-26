# Zaruba

> "My name is Zaruba. I came to be when Garo came to be and I am forever with him.”

Zaruba is technology agnostic artefact generator. Zaruba take care about file dependencies and will do necessary action to maintain consistency in your project.

# Concept

Zaruba's main purpose is to help you generate, test, and probably deploy things faster.

Let's say you do a bit modification in your repository. Zaruba will automatically tell you whether your modification pass the unit-test or not. Then, it will copy the repository to dependant-services. Finally, it will do some service-test and integration test. Whenever everything is okay, it will generate docker images and deploy them to your kubernetes-development-cluster.

Using Zaruba, you can get that level of automation.

In order to achieve the purpose, zaruba needs `template` and `dependency tree`.

## Template

Template is basically bunch of text files. A `template` might contains `zaruba.template.yaml`. Zaruba will exclude this config file whenever it needs to copy the template into `target`.

Template config's file might contains:

* `Substitute`
    - List of files (in regex format) containing keywords (i.e: `{{keyword}}`). Zaruba will replace those keywords based on the values in envvar.
    - Default to: `[]`
* `Ignore`
    - List of files (in regex format) that should not be watched by `zaruba`. For example, zaruba should not watch over `.git` directory
    - Default to: `["\.git"]`

## Project

Project is zaruba's workspace. Zaruba will watch every changes you made in a project and perform several actions. You can provide the action by editing `dependency tree`.

## Dependency Tree

In anywhere of your project, you can have `zaruba.dependency.yaml` containing a map. The key of the map is file's regex pattern, while it values are list of commands.

The dependency trees will be cascaded.

Below is a simple local-deployment-example:

```yaml
repos/ml-classifier/.*:
    - [python, "-m", "pytest", "repos/ml-classifier"]
    - [cp, "repos/ml-classifier", "services/ner/repo/model"]
    - [cp, "respos/ml-classifier", "services/sentiment-analysis/repo/model"]
services/ner/.*:
    - [python, "-m", "pytest", "services/ner"]
    - [docker, build, "-t", "gofrendi/ner-service", "services/ner/"]
services/sentiment-analysis/.*:
    - [python, "-m", "pytest", "services/sentiment-analysis"]
    - [docker, build, "-t", "gofrendi/sentiment-analysis-service", "services/sentiment-analysis-service/"]
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
* `ZARUBA_GITIGNORE_SEPARATOR`
    - Separator for user's gitignore items and zaruba's generated gitignore items. 
    - Default to `# zaruba's:`