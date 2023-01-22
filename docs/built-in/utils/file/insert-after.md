<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# insertAfter
<!--endTocHeader-->

```

Insert a new content into a text file after a particular index.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be -1.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to insert a 🍕 after 🍓.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert a 🍕 after this
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba file insertAfter fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍓
🍕
🍇

Usage:
  zaruba file insertAfter <strFileName> <strNewContent> [flags]

Aliases:
  insertAfter, append

Examples:

> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file insertAfter fruits.txt 🍕
> cat fruits.txt
🍊
🍓
🍇
🍕

> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file insertAfter fruits.txt 🍕 --index=1
> cat fruits.txt
🍊
🍓
🍕
🍇


Flags:
  -h, --help        help for insertAfter
  -i, --index int   index (default -1)

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->