# Zaruba Project

A typical zaruba project contains:

```
main.zaruba.yaml        # task declaration
default.kwargs.yaml     # kwargs declaration
```

# Zaruba Command: Please

Run some tasks declared in `task declaration file`.

```sh
zaruba please [tasks...] [kwarg=val] [-f task-declaration.yaml] [-k kwargs-declaration.yaml] [-e environment.env] [-e key=val]
zaruba please showTasks
zaruba please showPublishedTasks
zaruba please showUnpublishedTasks
```

## Tasks

Zaruba consider any argument without `=` character as a `task`.
You can provide as many `task` as you want.

If there is no `task` provided, Zaruba will show you all available tasks instead.

## Kwargs

Zaruba consider any argument containing `=` character as `keyword argument`.
If you have multiple `keyword arguments`, you can use `kwargs declaration file` instead.

## Task Declaration File

To provide custom `task declaration file`, you can use `-f` flag. Otherwise:

* Zaruba will try to use `main.zaruba.yaml` in the current directory.
* If `main.zaruba.yaml` is not found, Zaruba will use `${ZARUBA_HOME}/scripts/core.zaruba.yaml`. If `${ZARUBA_HOME}` is not defined, Zaruba will used it's parent directory as `${ZARUBA_HOME}`. In most cases, you don't need to set the environment variable.

### Kwargs Declaration File

To provide custom `kwargs declaration file`, you can use `-k` flag. Otherwise, Zaruba will try to use `default.kwargs.yaml` in the current directory.


# Task Declaration

```yaml
includes:
  - "${ZARUBA_HOME}/scripts/core.zaruba.yaml"
  - ./third-party/script.zaruba.yaml
  - http://example.com/remote-script.yaml # not implemented

tasks:

  # This is a very simple "command" task to open twitter.
  # This task only have `start` command.
  # So, the task is considered finished once the `start` command executed.
  openTwitter:
    description: Just open twitter
    start:
      - firefox
      - 'https://twitter.com'

  # Unlike the `openTwitter` task, 
  # this "process" task also has `check` command.
  #
  # Once `start` command executed successfully, 
  # Zaruba will run `check` command.
  # 
  # The task is considered finished once:
  # * The `check` command executed.
  # * Or the `start` command failed.
  receivePayment:
    description: Ask Jarvis to check client payment
    start:
      - jarvis
      - send-email
      - content="Dear client, please send payment"
    check:
      - /bin/sh
      - "-c"
      - until [ ! -z $(jarvis receive-payment) ]; do sleep 1; done

  # Similar to `receivePayment`, this task is a "process" task.
  # It has both `start` and `check` command.
  #
  # This task also has dependencies.
  # When you run a task with dependencies,
  # Zaruba will try to execute all of it's dependencies first.
  #
  # The task won't be executed if any of it's dependencies failed.
  #
  # You can also add environment variable for a task by using `env` key.
  # In this task, JARVIS_MODE value will be taken from CLIENT_MODE.
  # However, if CLIENT_MODE is not set, 
  # JARVIS_MODE will be set to `violently`.
  #
  # Both, env.<ENVVAR>.from and env.<ENVVAR>.default are optional.
  makeWebsite:
    description: Ask Jarvis to make website for client
    env:
      JARVIS_TOKEN:
        from: CLIENT_TOKEN
      JARVIS_MODE:
        from: CLIENT_MODE
        default: violently
    start:
      - jarvis
      - make-website'
      - host={{ .Kwargs.host}}
      - port={{ .Kwargs.port }}
    check:
      - /bin/sh
      - "-c"
      - until nc -z {{ .Kwargs.host }} {{ .Kwargs.port }}; do sleep 1; done
    dependencies:
      - receivePayment
  

  # Unlike the other tasks, this task extend `core.startDockerContainer`
  # Open `./scripts/core.zaruba.yaml` to read it's definition.
  #
  # A task migh have several `config`.
  # By look at the config, it is obvious that:
  # * This task will run `myNginx` container.
  # * The container should be based on `nginx` image
  # * Port `3030` is binded to port `80` of the container.
  # * Similarly, `./conf.d` and `./html` is binded to `/etc/nginx/conf.d`
  #   and `/usr/share/nginx/html` respectively.
  #
  # You don't need to define `start` and `end`
  # because they are already defined in `core.startDockerContainer`
  #
  # Finally, this task also has `location`.
  # Whenever you run a task with `location`, Zaruba will assume the `location`
  # as working directory.
  startNginx:
    extend: core.startDockerContainer
    location: ./nginx
    config:
      containerName: myNginx
      imageName: nginx
      port::3030: 80
      volume::./conf.d: /etc/nginx/conf.d
      volume::./html: /usr/share/nginx/html
  
  # Another useful base task is `core.startService`.
  #
  # Since `serveStatic` extend `core.startService`,
  # you can override `start` command to execute your service.
  #
  # This task also has `lconfig` since a service might expose many ports.
  serveStatic:
    extend: core.startService
    location: ./static
    lconfig:
      ports: [3031]
    start: ["python", "-m", "http.server", "{{ index .LConfig.ports 0 }}"]

  
  # This task has no `start` or `end`, but has some `dependencies`.
  # Task like this is called `"wrapper".
  startNginxAndServeStatic:
    dependencies:
      - startNginx
      - serveStatic

```

# Kwargs Declaration

You can think Kwargs declaration as set of `key value` pairs.

If you need nested keys, you can use `::` notation:

```yaml
defaultBranch: main
subrepo::talk::prefix: git-talk
subrepo::talk::url: https://github.com/jamietanna/gittalk15
subrepo::fancy::prefix: fancy
subrepo::fancy::url: https://github.com/jamietanna/gittalk15
```

The above configuration is mentally equivalent to:

```json
{
  "defaultBranch": "main",
  "subrepo": {
    "talk": {
      "prefix": "git-talk",
      "url": "https://github.com/jamietanna/gittalk15"
    },
    "fancy": {
      "prefix": "fancy",
      "url": "https://github.com/jamietanna/gittalk15"
  }
}
```

# Template

You can use [Go Template](https://golang.org/pkg/text/template/) to define the following configurations:

* task.config
* task.lconfig
* task.start
* task.check

The template expose `TaskData` as `{{ . }}`.

Please check `config/templatedata.go` to see `TaskData` definition. Here is a glimpse of the declaratioon:

```go
// TaskData is struct sent to template
type TaskData struct {
	task         *Task
	Name         string
	ProjectName  string
	BasePath     string
	WorkPath     string
	DirPath      string
	FileLocation string
	Kwargs       Dictionary
	Env          Dictionary
	Config       Dictionary
	LConfig      map[string][]string
	Decoration   logger.Decoration
}

// GetEnv of TaskData
func (td *TaskData) GetEnv(key string) (val string) {
  // ...
	return val
}

// GetAbsPath of any string
func (td *TaskData) GetAbsPath(parentPath, path string) (absPath string) {
  // ...
  return absPath
}

// Dictionary is advance map
type Dictionary map[string]string

// GetSubKeys get subkeys
func (d Dictionary) GetSubKeys(parentKeys ...string) (subKeys []string) {
  // ...
  return subKeys
}

// GetSubKeysBySeparator get subkeys by separator
func (d Dictionary) GetSubKeysBySeparator(separator string, parentKeys ...string) (subKeys []string) {
  // ...
  return subKeys
}

// GetValue of dictionary
func (d Dictionary) GetValue(keys ...string) (val string) {
  // ...
	return val
}

// GetValueBySeparator of dictionary
func (d Dictionary) GetValueBySeparator(separator string, keys ...string) (val string) {
  // ...
	return val
}
```

## Template Usage Example

Suppose you have the following kwargs:

```yaml
defaultBranch: main
subrepo::talk::prefix: git-talk
subrepo::talk::url: https://github.com/jamietanna/gittalk15
subrepo::fancy::prefix: fancy
subrepo::fancy::url: https://github.com/jamietanna/gittalk15
```

### Accessing Property

#### Template

```
{{ .Kwargs.defaultBranch}}
```

#### Output

```
main
```

### Condition

#### Template

```
{{ if .Kwargs.defaultBranch }}{{ .Kwargs.defaultBranch }}{{ else }}master{{ end }}
```

#### Output

```
main
```

### Subkeys

#### Template

```
{{ $names := .Kwargs.GetSubKeys "subrepo" -}}
{{ $kwargs := .Kwargs -}}
{{ range $index, $name := $names -}}
  {{ index }} {{ name }} {{ $kwargs.GetValue "subrepo" $name "prefix" }}
{{ end }}
```

#### Output
```
  0 talk git-talk
  0 fancy fancy
```
