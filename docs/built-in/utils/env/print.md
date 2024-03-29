<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [env](README.md)
# print
<!--endTocHeader-->

```

Print a jsonMap as environment variable declaration (key=value)

You can cascade the environment variable using --prefix flag.
This is useful if you have multiple environments (e.g., dev, staging, prod)

Usage:
  zaruba env print <jsonMap> [strFileName] [flags]

Aliases:
  print, write

Examples:

> zaruba env print '{"SERVER": "localhost", "PORT": "3306"}'
SERVER="localhost"
PORT="3306"

> zaruba env print '{"SERVER": "localhost", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}' --prefix=STG
SERVER="stg.stalchmst.com"
PORT="3306"
STG_SERVER="stg.stalchmst.com"
PROD_SERVER="stalchmst.com"


Flags:
  -h, --help            help for print
  -p, --prefix string   environment prefix

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->