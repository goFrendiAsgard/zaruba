<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [lines](README.md)
# replace
<!--endTocHeader-->

```

Replace a jsonStringList at a particular index with a new lines.
The Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be 0.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to replace 🍓 with two 🍕.

-------------------------------------------------
Elements | Index  | Note
-------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- replace this with two🍕
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba lines replace \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍕","🍕","🍇"]

Usage:
  zaruba lines replace <jsonStrList> <jsonStrListNewLines | strNewLine> [flags]

Examples:

> zaruba lines replace \
  '["🍊", "🍓", "🍇"]' \
  '🍕'
["🍕","🍓","🍇"]

> zaruba lines replace \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1
["🍊","🍕","🍕","🍇"]

> zaruba lines replace \
  '["🍊", "🍓", "🍇"]' \
  '["🍕"]' \
  --index=-1
["🍊","🍓","🍕"]



Flags:
  -h, --help        help for replace
  -i, --index int   desired pattern index

```

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->