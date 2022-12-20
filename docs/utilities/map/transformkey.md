<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🗺️ Map](README.md)
# TransformKey
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba map transformKey --help
```
 
<details>
<summary>Output</summary>
 
```````
Transform map keys

Usage:
  zaruba map transformKey <jsonMap> [flags]

Examples:

> zaruba map transformKey '{"server": "localhost", "port": 3306}' -p=DB_ -s=_DEV -t=upperSnake
{"DB_PORT_DEV":3306,"DB_SERVER_DEV":"localhost"}


Flags:
  -h, --help                         help for transformKey
  -p, --prefix string                key prefix
  -s, --suffix string                key suffix
  -t, --transformation stringArray   transformation (e.g., '-t upper', '-t lower', '-t upperSnake', -t 'camel', '-t kebab', '-t pascal', '-t snake')
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->