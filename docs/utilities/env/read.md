<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Read
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env read --help
```
 
<details>
<summary>Output</summary>
 
```````
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
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
<!--endTocSubTopic-->