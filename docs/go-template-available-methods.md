
# GetConfig

Get configuration value of current task. A configuration might be nested. For example, a `person` might have `name`, `age`, and `job` as subKeys:

```yaml
tasks:
  MY_TASK:
    config:
      person::name: Kouga Seijima
      person::age: Unknown
      person::job: Makai Knight
      madorin: zaruba
    start:
    - bash
    - '-c'
    - |
      echo {{ .GetConfig "madorin" }}
      # ^ will yield "zaruba"
      echo {{ .GetConfig "person" "name" }}
      # ^ will yield "Kouga Seijima"
      echo {{ .GetConfig "person::name" }} 
      # ^ will also yield "Kouga Seijima"
```

# GetConfigs

Get all configs of current task as map.

```yaml
tasks:
  MY_TASK:
    config:
      person::name: Kouga Seijima
      person::age: Unknown
      person::job: Makai Knight
      madorin: zaruba
    start:
    - bash
    - '-c'
    - |
      {{ $map := .GetConfigs }} # getting all configs
      {{ range $key, $val := $map }}
        echo "Key: {{ $key }}, Val: {{ $val }}"
        # ^ will show all config's key and value
      {{ end }}
```

# GetSubConfigKeys

Get Subkeys of current task's config as list.

```yaml
tasks:
  MY_TASK:
    config:
      person::name: Kouga Seijima
      person::age: Unknown
      person::job: Makai Knight
      madorin: zaruba
    start:
    - bash
    - '-c'
    - |
      {{ $keys := .GetSubConfigKeys "person" }}
      # ^ getting all sub keys of "person" config as a list
      {{ range _, $key := $keys }}
        echo "Key: {{ $key }}"
        # ^ will yield "name", "age", and "job"
      {{ end }}
```

# GetLConfig

Get LConfig value of current task.

```yaml
tasks:
  MY_TASK:
    lconfig:
      ports:
      - 5672
      - 15672
    start:
    - bash
    - '-c'
    - |
      {{ $index, $val := range .GetLConfig "ports" }}
        echo "Port {{ $index }}: {{ $val }}"
      {{ end }}
```

# GetValue

Get values of current task. A configuration might be nested. For example, a `person` might have `name`, `age`, and `job` as subKeys:

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .GetValue "madorin" }}
      echo {{ .GetValue "person" "name" }}
      echo {{ .GetValue "person::name" }} 
```

# GetSubValueKeys

Get Subkeys of current task's values as list.

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      {{ $keys := .GetSubValueKeys "person" }}
      {{ range _, $key := $keys }}
        echo "Key: {{ $key }}"
      {{ end }}
```

# GetEnv

Get configuration value of current task. A configuration might be nested. For example, a `person` might have `name`, `age`, and `job` as subKeys:

```yaml
tasks:
  MY_TASK:
    env:
      USER:
        from: TASK_USER
        default: Kouga Seijima
      MADORIN:
        default: zaruba
    start:
    - bash
    - '-c'
    - |
      echo {{ .GetEnv "MADORIN" }}
      # ^ will yield "zaruba"
      echo {{ .GetEnv "USER" }}
      # ^ will yield "Kouga Seijima" or depends on "TASK_USER" env
```

# GetEnvs

Get all configs of current task as map.

```yaml
tasks:
  MY_TASK:
    env:
      USER:
        from: TASK_USER
        default: Kouga Seijima
      MADORIN:
        default: zaruba
    start:
    - bash
    - '-c'
    - |
      {{ $map := .GetEnvs }} # getting all envs
      {{ range $key, $val := $map }}
        echo "Key: {{ $key }}, Val: {{ $val }}"
        # ^ will show all env's key and value
      {{ end }}
```

# GetWorkPath

Get path relative to `WorkPath`

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .GetWorkPath "subDir" }}
```

# GetRelativePath

Get path relative to `DirPath`

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .GetRelativePath "subDir" }}
```

# GetTask

Get other task


```yaml
tasks:

  MY_OTHER_TASK:
    config:
      MY_OTHER_CONFIG: MY_VAL

  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      {{ $task := .GetTask "MY_OTHER_TASK" }}
      echo {{ $task.GetConfig "MY_OTHER_CONFIG" }}
```

# IsTrue


```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .IsTrue "y" }} # 1
      echo {{ .IsTrue "n" }} # 0
```


# IsFalse


```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .IsFalse "y" }} # 0
      echo {{ .IsFalse "n" }} # 1
```

# Trim

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .Trim " lorem ipsum dolor sit amet  " }} # "lorem ipsum dolor sit amet"
```


# ReadFile

Read content of file

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .ReadFile "db.sql" }}
```

# ListDir

Get list of directory as list

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .ListDir "/mnt/c/Users" }}
```


# ParseFile

Like `ReadFile`, but also parse go template.

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .ParseFile "db.sql" }}
```


# ReplaceAllWith

```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      echo {{ .ReplaceAllWith "orange and grape" "orange" "grape" "fruit" }}
      # ^ fruit and fruit
```


# EscapeShellValue


```yaml
tasks:
  MY_TASK:
    start:
    - bash
    - '-c'
    - |
      mysql -u root -e {{ .EscapeShellValue (.ParseFile "db.sql") }}
```