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
JsonMap manipulation utilities

Usage:
  zaruba map [command]

Available Commands:
  get               Get value from jsonMap at a particular key
  getKeys           Return a jsonStringList containing all keys in a jsonMap
  merge             Merge multiple jsonMaps
  rangeKey          Print jsonMap keys
  set               Set a value of a jsonMap on a particular key
  toStringMap       Transform a jsonMap into a jsonStringMap
  toVariedStringMap Transform a jsonMap into a jsonStringMap, every keys and values are transformed into multiple variations
  transformKey      Transform map keys
  validate          Check whether jsonMap is valid JSON map or not

Flags:
  -h, --help   help for map

Use "zaruba map [command] --help" for more information about a command.
```````
</details>
<!--endCode-->

<!--startTocSubtopic-->
# Sub-topics
* [Get](get.md)
* [GetKeys](getkeys.md)
* [Merge](merge.md)
* [RangeKey](rangekey.md)
* [Set](set.md)
* [ToStringMap](tostringmap.md)
* [ToVariedStringMap](tovariedstringmap.md)
* [TransformKey](transformkey.md)
* [Validate](validate.md)
<!--endTocSubtopic-->