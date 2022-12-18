<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Print
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env print --help
```
 
<details>
<summary>Output</summary>
 
```````
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
```````
</details>
<!--endCode-->

# Examples

## Print Environment

<!--startCode-->
```bash
zaruba env print \
    '{"APP_CONTEXT": "aws", "APP_VERSION": "1.0.0"}'
```
 
<details>
<summary>Output</summary>
 
```````
APP_CONTEXT="aws"
APP_VERSION="1.0.0"
```````
</details>
<!--endCode-->

## Print Cascaded Environment

Cascaded environment is very useful if you manage several platforms with similar environments.

For example, you have `dev`, `stag`, and `prod`.

By adding `DEV` prefix (i,e., `zaruba env print <map> DEV`), your `DEV` environment will be assigned to the original ones.

If prefixed environment does not exists, Zaruba will use original values. Thus you can treat original environment as fallback/default values.


<!--startCode-->
```bash
zaruba env print \
    '{"APP_CONTEXT": "aws", "APP_VERSION": "1.0.0", "STAG_APP_CONTEXT": "azure", "STAG_APP_VERSION": "1.1.0", "DEV_APP_CONTEXT": "gcp", "DEV_APP_VERSION": "1.1.1"}' \
    --prefix=DEV
```
 
<details>
<summary>Output</summary>
 
```````
APP_CONTEXT="gcp"
APP_VERSION="1.1.1"
DEV_APP_CONTEXT="gcp"
DEV_APP_VERSION="1.1.1"
STAG_APP_CONTEXT="azure"
STAG_APP_VERSION="1.1.0"
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->