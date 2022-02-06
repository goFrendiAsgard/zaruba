[⬅️ Table of Content](../../README.md)

# Using Go Template

You can use Zaruba's go template in:

* [project configs](./project-configs.md)
* [task configs](./task/task-configs/README.md)
* task's start property
* task's check property

To see complete list of supported functions, you can take a look at [zaruba's source code](../../../core/tpl.go).

# Most Commonly Used Functions

Some of the most commonly used functions are:

* `{{ .GetEnv "SOME_ENV" }}`
* `{{ .GetConfig "someConfig" }}`
* `{{ .GetValue "someValue" }}`

Let's take a look on this example:

```yaml
inputs:

  name:
    prompt: Your name
    description: User name


tasks:

  sayHello:
    extend: zrbRunShellScript
    inputs:
      - name
    envs:
      NAME:
        from: USERNAME
    configs:
      name: |
        {{ if .GetValue "name" }}{{ .GetValue "name" }}{{ else if .GetEnv "NAME" }}{{ .GetEnv "NAME" }}{{ else }}world{{ end }}
      start: |
        echo "Hello {{ .GetConfig "name" }}"

```

Let's break down `tasks.sayHello.configs.name`:

```gotmpl
{{ if .GetValue "name" }}{{ .GetValue "name" }}{{ else if .GetEnv "NAME" }}{{ .GetEnv "NAME" }}{{ else }}world{{ end }}
```

* First of all, Zaruba will check your `name` input value. If the value is not empty, it will be used as `tasks.sayHello.configs.name`'s value.
* Otherwise, if `name` input value is empty, Zaruba will check on `NAME` environment. If it eixsts, it will be used as `tasks.sayHello.configs`.
* Finally if `name` input value and `NAME` environment are empty, `tasks.sayHello.configs` will be set to `world`.

Once `tasks.sayHello.configs.name` value has been determined, you can use it to fill `tasks.sayHello.configs.start`:

```gotmpl
echo "Hello {{ .GetConfig "name" }}"
```

# Other Commonly Used Functions/Values

* `{{ .ZarubaHome }}`: Zaruba's home directory (e.g: `~/zaruba`)
* `{{ .ZarubaBin }}`: Zaruba's bin (e.g: `~/zaruba/zaruba`)
* `{{ .UUID }}`: Generated UUID (e.g: `30fd76b8-8e65-4d20-a124-d46ca1665e1f`)
* `{{ .GeneratedRandomName }}`: Generated name containing two words camel cased (e.g: `corruptedSenate`)
* `{{ .Decoration }}`: Decoration object.
* `{{ .GetWorkPath path }}`: Getting absolute path relative to task location.
* `{{ .GetTaskPath path }}`: Getting absolute path relative to task definition's location.
* `{{ .GetProjectPath path }}`: Getting absolute path relative to project's location.
* `{{ .GetConfig key }}`: Get single config
* `{{ .GetConfigs keyPattern }}`: Get configs as map
* `{{ .GetConfigsAsShellVariables keyPattern variablePrefix}}`: Get configs as shell variables
* `{{ .GetPorts }}`: Get ports/container ports as list of int.
* `{{ .GetSubValueKeys parentKeys... }}`: Get sub keys of sub value (separated by `::`).
* `{{ .GetValue keys... }}`: Get single value.
* `{{ .GetEnv key }}`: Get single env.
* `{{ .GetEnvs }}`: Get envs as map.
* `{{ .ReplaceAll s old new }}`: Replace string.
* `{{ .GetDockerImageName }}`: Get docker image name of current task location.
* `{{ .ParseFile filePath }}`: Parse file. The file might contains Zaruba's go template.
* `{{ .Template content }}`: Escaping go template.


# List of All Functions

Please visit [zaruba's source code](../../../core/tpl.go) to see available functions/values.