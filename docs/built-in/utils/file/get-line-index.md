<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# getLineIndex
<!--endTocHeader-->

```

Return the index of a line matching a particular index at a specified patterns.
Index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be -1.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
🍊
🍓
🍇
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


Then, you need to invoke the following command:
> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=2

The result will be: 4

Usage:
  zaruba file getLineIndex <strFileName> <jsonStrListPatterns> [flags]

Examples:

> cat fruits.txt
🍊
🍓
🍇
🍊
🍓
🍇

> zaruba file getLineIndex \
  fruits.txt \
  '🍓'
1

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=1
3

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=-1
5


Flags:
  -h, --help        help for getLineIndex
  -i, --index int   desired pattern index (default -1)

```

<!--startTocSubtopic-->

<!--endTocSubtopic-->