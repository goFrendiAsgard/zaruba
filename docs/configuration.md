<!--startTocHeader-->
[🏠](README.md)
# ⚙️ Configuration
<!--endTocHeader-->

You can configure Zaruba's behavior by using some environment variables.

To make your changes permanent, you need to declare them in your shell configuration file.

If you are using `bash`, your shell configuration file should be `${HOME}/.bashrc`.

# ZARUBA_HOME

Location of your Zaruba installation directory.

## Default Value

Zaruba's installation directory.

```
${HOME}/.zaruba
```

## Example

```bash
ZARUBA_HOME=/usr/bin/zaruba
```

# ZARUBA_BIN

Location of your Zaruba executable binary.

## Default Value

Zaruba's executable path.

```
${HOME}/.zaruba/zaruba
```

## Example

```bash
ZARUBA_BIN=/usr/bin/zaruba/zaruba
```
# ZARUBA_SHELL

The shell Zaruba uses to run shell scripts (e.g., `bash` or `zsh`). Using `bash` is preferable since most systems support it.

## Default Value

```
bash
```

## Example

```bash
ZARUBA_SHELL=zsh
```

# ZARUBA_SCRIPTS

Location of globally-available scripts. If you have many locations, you can use colons as separators. 

## Default Value

(Not Available)

## Example

```bash
ZARUBA_SCRIPTS=${HOME}/coffee-maker/index.zaruba.yaml:${HOME}/my-organization/index.zaruba.yaml
```

# ZARUBA_DECORATION

How Zaruba output should look like.

There are some possible values for this:
  - `default`: Default decoration. This decoration has several symbols and colors.
  - `colorless`: This decoration has several symbols, but doesn't have any color.
  - `plain`: Plain decoration, has no symbol or color.

## Default Value

```
default
```

## Example

```bash
ZARUBA_DECORATION=colorless
```

# ZARUBA_LOG_TIME

Whether Zaruba output should include time in its output or not.

## Default Value

```
true
```

# Example

```bash
ZARUBA_LOG_TIME=false
```

# ZARUBA_LOG_STATUS_TIME_INTERVAL

Time interval to show status and PID of every processes

## Default Value

```
5m
```

# Example

```bash
ZARUBA_LOG_STATUS_TIME_INTERVAL=10m
```

# ZARUBA_LOG_STATUS_LINE_INTERVAL

Line interval to show status and PID of every processes

## Default Value

```
40
```

# Example

```bash
ZARUBA_LOG_STATUS_LINE_INTERVAL=50
```

# ZARUBA_MAX_LOG_FILE_SIZE

Maximum log file size (in byte)

## Default Value

```
5242880
```

## Example

```bash
# 1 MB = 1024 * 1024 = 1048576
ZARUBA_MAX_LOG_FILE_SIZE=1048576
```

# ZARUBA_ENV 

Zaruba environment. The default value is empty.
Zaruba will load `<ZARUBA_ENV>.env` and `<ZARUBA_ENV>.values.yaml` whenever you run `zaruba please`.

## Default Value

(Not Available)

## Example

```bash
ZARUBA_ENV=dev
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->