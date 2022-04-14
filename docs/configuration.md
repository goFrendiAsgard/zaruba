<!--startTocHeader-->
[🏠](README.md)
# ⚙️ Configuration
<!--endTocHeader-->

# Environment Variables

You can configure Zaruba's behavior by using several environment variables:

* `ZARUBA_HOME`: Location of your Zaruba installation directory. If unset, Zaruba will use the executable's directory path instead (e.g: `${HOME}/.zaruba`).
* `ZARUBA_BIN` Location of your Zaruba executable binary. If unset, Zaruba will use the executable's path instead (e.g: `${HOME}/.zaruba/zaruba`).
* `ZARUBA_SHELL` The shell zaruba used to execute shell scripts (e.g: `bash`, `zsh`, or `sh`). Using `bash` is preferable because it is widely use. If the environment varible is unset, Zaruba will use `bash` by default.
* `ZARUBA_SCRIPTS` List of zaruba scripts that are going to be available from everywhere. Use colon (`:`) to separate the scripts (e.g: `${HOME}/coffee-maker/index.zaruba.yaml:${HOME}/my-organization/index.zaruba.yaml`).
* `ZARUBA_DECORATION`: Zaruba output decoration
    - `default`: Default decoration. This decoration has several symbols and colors.
    - `colorless`: This decoration has several symbols, but doesn't have any color.
    - `plain`: Plain decoration, has no symbol or color.


# Save Configuration

You can put these environment variable in your `~/.bashrc` or `~/.zshrc` depending on your day-to-day terminal.

For example, if you use `zsh` for daily operation, and you want `ZARUBA_DECORATION` to be colorless, you add this at the bottom of your `~/.zshrc`:

```bash
ZARUBA_DECORATION=colorless
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->