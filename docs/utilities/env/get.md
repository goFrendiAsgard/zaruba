<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Get
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env get --help
```
 
<details>
<summary>Output</summary>
 
```````
Get current environment variables as jsonMap.

You can cascade the environment variable using --prefix flag.
This is useful if you have multiple environments (e.g., dev, staging, prod)

Usage:
  zaruba env get [flags]

Examples:

> export SERVER=localhost
> export PORT=3306

> zaruba env get
{"SERVER": "localhost", "PORT": "3306"}

> export SERVER=localhost
> export STG_SERVER=stg.stalchmst.com
> export PROD_SERVER=stalchmst.com
> export PORT=3306

> zaruba env get --prefix=STG
{"SERVER": "stg.stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=PROD
{"SERVER": "stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=DEV
{"SERVER": "localhost", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}


Flags:
  -h, --help            help for get
  -p, --prefix string   environment prefix
```````
</details>
<!--endCode-->



<!--startTocSubtopic-->
<!--endTocSubtopic-->