<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# getLine
<!--endTocHeader-->

```

Return the index of a line matching a particular index at a specified patterns.
Index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be -1.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓A
🍇
🍊
🍓B
🍇
You want to get the index of a line containing a 🍓 that is located after two 🍊 and before a 🍇.

---------------------------------------------------------------------------------
Elements | Element index  | Patterns | Pattern Index | Note
---------------------------------------------------------------------------------
🍊       | 0              | 🍊       | 0/-4          |
🍓A      | 1              |          |               |
🍇       | 2              |          |               |
🍊       | 3              | 🍊       | 1/-3          |
🍓B      | 4              | 🍓.*     | 2/-2          | <-- We want this 🍓
🍇       | 5              | 🍇       | 3/-1          |


Then, you need to invoke the following command:
> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓.*","🍇"]' \
  --index=2

The result will be:
🍓B

Usage:
  zaruba file getLine <strFileName> <index> [flags]

Examples:

> cat fruits.txt
🍊A
🍓B
🍇C
🍊D
🍓E
🍇F

> zaruba file getLineIndex \
  fruits.txt \
  '🍓.*'
🍓B

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊.*", "🍊.*", "🍓.*","🍇.*"]' \
  --index=1
🍊D

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊.*", "🍊.*", "🍓.*","🍇.*"]' \
  --index=-1
🍇F


Flags:
  -h, --help   help for getLine

```

<!--startTocSubtopic-->

<!--endTocSubtopic-->