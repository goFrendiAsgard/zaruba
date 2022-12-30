<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# submatch
<!--endTocHeader-->

```

Return submatch matching the pattern at a desired pattern index.
Index is started from 0. You can use negative index to count from the end of file.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍌🍓🍈
🍇
🍊
🥑🍓🍎🍏
🍇
First, you want to get a line containing a 🍓 that is located after two 🍊 and before a 🍇.
Then you want to get what characters are preceeding/following the 🍓 at that particular line.

---------------------------------------------------------------------------------------------
Elements   | Element index  | Patterns   | Pattern Index | Note
---------------------------------------------------------------------------------------------
🍊         | 0              | 🍊         | 0/-4          |
🍌🍓🍈     | 1              |            |               |
🍇         | 2              |            |               |
🍊         | 3              | 🍊         | 1/-3          |
🥑🍓🍎🍏   | 4              | (.*)🍓(.*) | 2/-2          | <-- We want "🥑" and "🍎🍏"
🍇         | 5              | 🍇         | 3/-1          |

To do this, you need to invoke the following command:
> zaruba file submatch \
  fruits.txt \
  '["🍊", "🍊", "(.*)🍓(.*)", "🍇"]' \
  --index=2

The result will be:
["🥑🍓🍎🍏","🥑","🍎🍏"]

You can see that there are three elements of the result:
- The whole line : 🥑🍓🍎🍏
- The characters preceding 🍓: 🥑
- The characters following 🍓: 🍎🍏

Usage:
  zaruba file submatch <jsonStrList> <jsonStrListPatterns> [flags]

Examples:

> cat fruits.txt
🍊
🍌🍓🍈
🍇
🍊
🥑🍓🍎🍏
🍇

> zaruba file submatch \
  fruits.txt \
  '["🍊", "🍊", "(.*)🍓(.*)", "🍇"]' \
  --index=2
["🥑🍓🍎🍏","🥑","🍎🍏"]

> zaruba file submatch \
  fruits.txt \
  "(.*)🍓(.*)"
["🍌🍓🍈","🍌","🍈"]


Flags:
  -h, --help        help for submatch
  -i, --index int   index (default -1)

```

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->