<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# replaceAt
<!--endTocHeader-->

```

Replace a file at a particular index with a new content.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be 0.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to replace 🍓 with a 🍕.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- replace this with 🍕
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba file replaceAt fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍕
🍇

Usage:
  zaruba file replaceAt <strFileName> <strNewContent> [flags]

Examples:

> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file replaceAt fruits.txt 🍕
> cat fruits.txt
🍕
🍓
🍇

> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file replaceAt fruits.txt 🍕 --index=1
> cat fruits.txt
🍊
🍕
🍇


Flags:
  -h, --help        help for replaceAt
  -i, --index int   index

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->