<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [lines](README.md)
# getIndex
<!--endTocHeader-->

```

Return the index of a line matching a particular index at a specified patterns.
Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be -1.

For example, you have a jsonStringList ["🍊", "🍓", "🍇","🍊", "🍓","🍇"].
You want to get the index of an 🍓 that is located after two 🍊 and before a 🍇.

---------------------------------------------------------------------------------
Elements | Element index  | Patterns | Pattern Index | Note
---------------------------------------------------------------------------------
🍊       | 0              | 🍊       | 0/-4          |
🍓       | 1              |          |               |
🍇       | 2              |          |               |
🍊       | 3              | 🍊       | 1/-3          |
🍓       | 4              | 🍓       | 2/-2          | <-- We want this 🍓
🍇       | 5              | 🍇       | 3/-1          |


In that case, you need to invoke the following command:
> zaruba lines getIndex \
  '["🍊", "🍓", "🍇","🍊", "🍓","🍇"]' \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=2

The result will be: 4

Usage:
  zaruba lines getIndex <jsonStrList> <jsonStrListPatterns> [flags]

Examples:

> LINES='["🍊", "🍓", "🍇", "🍊", "🍓", "🍇"]'
> zaruba lines getIndex "$LINES" '🍓'
1
> zaruba lines getIndex "$LINES" '["🍊", "🍊", "🍓","🍇"]' --index=1
3
> zaruba lines getIndex "$LINES" '["🍊", "🍊", "🍓","🍇"]' --index=-1
5


Flags:
  -h, --help        help for getIndex
  -i, --index int   desired pattern index (default -1)

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->