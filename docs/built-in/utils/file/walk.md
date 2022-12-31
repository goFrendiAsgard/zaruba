<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# walk
<!--endTocHeader-->

```
List files/folder in a path, recursively

Usage:
  zaruba file walk <strDirPath> [flags]

Examples:

> ls myDir
a.txt   b.txt   c
> ls myDir/c
d.txt   e.txt

> zaruba file walk myDir
/a.txt
/b.txt
/c
/c/d.txt
/c/e.txt


Flags:
  -h, --help   help for walk

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->