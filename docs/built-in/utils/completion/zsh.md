<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [completion](README.md)
# zsh
<!--endTocHeader-->

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	zaruba completion zsh > "${fpath[1]}/_zaruba"

#### macOS:

	zaruba completion zsh > /usr/local/share/zsh/site-functions/_zaruba

You will need to start a new shell for this setup to take effect.

Usage:
  zaruba completion zsh [flags]

Flags:
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions

```

<!--startTocSubtopic-->

<!--endTocSubtopic-->