<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [lines](README.md)
# insertBefore
<!--endTocHeader-->

```

Insert new lines into a jsonStringList before a particular index.
The Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be 0.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to insert two 🍕 before 🍓.

--------------------------------------------------
Elements | Index  | Note
--------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert two🍕 before this
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍕","🍕","🍓","🍇"]

Usage:
  zaruba lines insertBefore <jsonStrList> <jsonStrListNewLines | strNewLine> [flags]

Aliases:
  insertBefore, prepend

Examples:

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertBefore "$LINES" 🍕
["🍕","🍊","🍓","🍇"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertBefore "$LINES" '["🍕", "🍕"]' --index=1
["🍊","🍕","🍕","🍓","🍇"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertBefore "$LINES" '["🍕"]' --index=-1
["🍊","🍓","🍕","🍇"]


Flags:
  -h, --help        help for insertBefore
  -i, --index int   desired pattern index

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->