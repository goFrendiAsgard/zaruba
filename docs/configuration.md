<!--startTocHeader-->
[🏠](README.md)
# ⚙️ Configuration
<!--endTocHeader-->

You can configure Zaruba's behavior by using some environment variables.

## ZARUBA_HOME

Location of your Zaruba installation directory. If unset, Zaruba will use its executable's directory location (e.g., `${HOME}/.zaruba`). 

Example:

```bash
ZARUBA_HOME=${HOME}/zaruba
```

## ZARUBA_BIN

Location of your Zaruba executable binary. If unset, Zaruba will use its executable's path (e.g., `${HOME}/.zaruba/zaruba`).

Example:

```bash
ZARUBA_HOME=${HOME}/zaruba/zaruba
```

## ZARUBA_SHELL

The shell Zaruba uses to run shell scripts (e.g., `bash`, `zsh`, or `sh`). Using `bash` is preferable since most systems support it. If unset, Zaruba will use `bash` by default.

Example:

```bash
ZARUBA_SHELL=bash
```

## ZARUBA_SCRIPTS

Location of globally-available scripts. If you have many locations, you can use colons as separators. 

Example:

```bash
ZARUBA_SCRIPTS=${HOME}/coffee-maker/index.zaruba.yaml:${HOME}/my-organization/index.zaruba.yaml`)
```

## ZARUBA_DECORATION

How Zaruba output should look like. The default value is `default`.

There are some possible values for this:
  - `default`: Default decoration. This decoration has several symbols and colors.
  - `colorless`: This decoration has several symbols, but doesn't have any color.
  - `plain`: Plain decoration, has no symbol or color.

Example:

```bash
ZARUBA_DECORATION=default
```

## ZARUBA_LOG_TIME

Whether Zaruba output should include time or not. The default value is `true`. 

Example:

```bash
ZARUBA_LOG_TIME=false
```

## ZARUBA_ENV 

Zaruba environment. The default value is empty.
Zaruba will load `<ZARUBA_ENV>.env` and `<ZARUBA_ENV>.values.yaml` whenever you run `zaruba please`.

Example:

```bash
ZARUBA_ENV=dev
```

# Save Configuration

You can set Zaruba configuration anytime you want. But, to make your changes permanent, you need to declare them in your shell configuration file.

If you are using `bash`, your shell configuration file should be `${HOME}/.bashrc`.

<!--startTocSubTopic-->
<!--endTocSubTopic-->