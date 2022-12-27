<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🗺️ Map](README.md)
# RangeKey
<!--endTocHeader-->

# Usage


```bash
zaruba map rangeKey --help
```
 
<details>
<summary>Output</summary>
 
```````
Print jsonMap keys

Usage:
  zaruba map rangeKey <jsonMap> [flags]

Examples:

> zaruba map rangeKey '{"server": "localhost", "port": 3306}'
server
port

> MAP={"server": "localhost", "port": 3306}
> for KEY in $(zaruba map rangeKey "$MAP")
  do
	VALUE=$(zaruba map get "$MAP" $KEY)
	echo "$KEY $VALUE"
  done

server localhost
port 3306


Flags:
  -h, --help   help for rangeKey
```````
</details>


# Examples



<!--startTocSubtopic-->

<!--endTocSubtopic-->