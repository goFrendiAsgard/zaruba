<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [completion](README.md)
# bash
<!--endTocHeader-->

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(zaruba completion bash)

To load completions for every new session, execute once:

#### Linux:

	zaruba completion bash > /etc/bash_completion.d/zaruba

#### macOS:

	zaruba completion bash > /usr/local/etc/bash_completion.d/zaruba

You will need to start a new shell for this setup to take effect.

Usage:
  zaruba completion bash

Flags:
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->