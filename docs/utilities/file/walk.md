<!--startTocheader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [File](README.md)
# Walk
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba file walk --help
```
 
<details>
<summary>Output</summary>
 
```````
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
```````
</details>
<!--endCode-->

# Examples


<!--startTocSubTopic-->
<!--endTocSubTopic-->