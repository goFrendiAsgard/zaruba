<!--startTocHeader-->
[🏠](README.md)
# ⚙️ Configuration
<!--endTocHeader-->

You can configure Zaruba's behavior by using some environment variables.

## ZARUBA_HOME

Location of your Zaruba installation directory. If unset, Zaruba will use its executable's directory location (e.g: `${HOME}/.zaruba`).

## ZARUBA_BIN

Location of your Zaruba executable binary. If unset, Zaruba will use its executable's path (e.g: `${HOME}/.zaruba/zaruba`).

## ZARUBA_SHELL

The shell used to run shell scripts (e.g: `bash`, `zsh`, or `sh`). Using `bash` is preferable since most systems support `bash`. If unset, Zaruba will use `bash` by default.
## ZARUBA_SCRIPTS

Location of globally-available tasks. You can use colon (`:`) as separator (e.g: `${HOME}/coffee-maker/index.zaruba.yaml:${HOME}/my-organization/index.zaruba.yaml`).

## ZARUBA_DECORATION

Zaruba output decoration
    - `default`: Default decoration. This decoration has several symbols and colors.
    - `colorless`: This decoration has several symbols, but doesn't have any color.
    - `plain`: Plain decoration, has no symbol or color.


# Save Configuration

To make your changes permanent, you need to save your configuration into `~/.bashrc` or `~/.zshrc` (depending on your terminal).

For example, if you are using `zsh` for your daily operation, and you want to set `ZARUBA_DECORATION` into `colorless`, you can add this at the bottom of your `~/.zshrc`:

```bash
ZARUBA_DECORATION=colorless
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->