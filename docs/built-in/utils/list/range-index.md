<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [list](README.md)
# rangeIndex
<!--endTocHeader-->

```
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

```

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->