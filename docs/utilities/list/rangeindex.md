<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🧺 List](README.md)
# RangeIndex
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba list rangeIndex --help
```
 
<details>
<summary>Output</summary>
 
```````
Print jsonList indexes

Usage:
  zaruba list rangeIndex <jsonList> [flags]

Examples:

> zaruba list rangeIndex '["🍊","🍓","🍇"]'
0
1
2

> LIST='["🍊","🍓","🍇"]'
> for INDEX in $(zaruba list rangeIndex "$LIST")
  do
	VALUE=$(zaruba list get "$LIST" $INDEX)
	echo "$INDEX $VALUE"
  done

0 🍊
1 🍓
2 🍇


Flags:
  -h, --help   help for rangeIndex
```````
</details>
<!--endCode-->

# Examples


<!--startTocSubTopic-->
<!--endTocSubTopic-->