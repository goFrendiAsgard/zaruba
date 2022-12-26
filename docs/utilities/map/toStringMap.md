<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🗺️ Map](README.md)
# ToStringMap
<!--endTocHeader-->

# Usage


```bash
zaruba map toStringMap --help
```
 
<details>
<summary>Output</summary>
 
```````
Transform a jsonMap into a jsonStringMap

Usage:
  zaruba map toStringMap <jsonMap> [flags]

Examples:

> zaruba map toStringMap '{"server": "localhost", "port": 3306, "env": {"enable_ui": 0}}'
{"env":"{\"enable_ui\":0}","port":"3306","server":"localhost"}


Flags:
  -h, --help   help for toStringMap
```````
</details>


# Examples



<!--startTocSubtopic-->
<!--endTocSubtopic-->