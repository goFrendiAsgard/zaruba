<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md)
# 🗺️ Map
<!--endTocHeader-->

A `map` is a valid JSON object.

To see list of `map` utilities you can invoke:

<!--startCode-->
```bash
zaruba map
```
 
<details>
<summary>Output</summary>
 
```````
Map manipulation utilities

Usage:
  zaruba map [command]

Available Commands:
  get               Get value from JSON map
  getKeys           Return JSON string list containing keys of JSON map
  merge             Merge JSON maps, in case of duplicate keys, the first ocurrance is going to be used
  rangeKey          Print map keys
  set               Set map[key] to value
  toStringMap       Transform to string map
  toVariedStringMap Transform to string map
  transformKey      Transform map keys
  validate          Check whether jsonMap is valid JSON map or not

Flags:
  -h, --help   help for map

Use "zaruba map [command] --help" for more information about a command.
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
# Sub-topics
* [Get](get.md)
* [GetKeys](get-keys.md)
* [Merge](merge.md)
* [RangeKey](range-key.md)
* [Set](set.md)
* [ToStringMap](to-string-map.md)
* [ToVariedStringMap](to-varied-string-map.md)
* [TransformKey](transform-key.md)
* [Validate](validate.md)
<!--endTocSubTopic-->