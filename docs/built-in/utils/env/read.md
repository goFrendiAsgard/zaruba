<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [env](README.md)
# read
<!--endTocHeader-->

```

Read environment variable from env file and return a jsonMap.

Usage:
  zaruba env read <strFileName> [flags]

Examples:

> cat .env
SERVER=localhost
PORT=3306
> zaruba env read .env
{"SERVER": "localhost", "PORT": "3306"}


Flags:
  -h, --help            help for read
  -p, --prefix string   environment prefix

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->