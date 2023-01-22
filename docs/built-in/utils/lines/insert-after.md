<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [lines](README.md)
# insertAfter
<!--endTocHeader-->

```

Insert new lines into a jsonStringList after a particular index.
The index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be -1.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to insert two 🍕 after 🍓.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert two🍕 after this
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba lines insertAfter \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍓","🍕","🍕","🍇"]

Usage:
  zaruba lines insertAfter <jsonStrList> <jsonStrListNewLines | strNewLine> [flags]

Aliases:
  insertAfter, append

Examples:

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertAfter "$LINES" 🍕
["🍊","🍓","🍇", "🍕"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertAfter "$LINES" '["🍕", "🍕"]' --index=1
["🍊","🍓","🍕","🍕","🍇"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines insertAfter "$LINES" '["🍕"]' --index=-1
["🍊","🍓","🍇","🍕"]


Flags:
  -h, --help        help for insertAfter
  -i, --index int   desired pattern index (default -1)

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->