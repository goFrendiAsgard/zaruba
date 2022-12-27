<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🗺️ Map](README.md)
# ToVariedStringMap
<!--endTocHeader-->

# Usage


```bash
zaruba map toVariedStringMap --help
```
 
<details>
<summary>Output</summary>
 
```````
Transform a jsonMap into a jsonStringMap, every keys and values are transformed into multiple variations

Usage:
  zaruba map toVariedStringMap <jsonMap> [keys...] [flags]

Examples:

zaruba map toVariedStringMap '{"server": "localhost", "port": 3306}'
{"\"port\"":"\"3306\"","\"server\"":"\"localhost\"","'port'":"'3306'","'server'":"'localhost'","PORT":"3306","Port":"3306","SERVER":"LOCALHOST","Server":"Localhost","port":"3306","server":"localhost"}


Flags:
  -h, --help   help for toVariedStringMap
```````
</details>


# Examples



<!--startTocSubtopic-->

<!--endTocSubtopic-->