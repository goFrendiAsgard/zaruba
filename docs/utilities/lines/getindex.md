<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🚈 Lines](README.md)
# GetIndex
<!--endTocHeader-->


# Usage

<!--startCode-->
```bash
zaruba lines getIndex --help
```
 
<details>
<summary>Output</summary>
 
```````
Getting line index of a a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of the file.

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1

Usage:
  zaruba lines getIndex <jsonStrList> <jsonStrListPatterns> [flags]

Examples:

> CONTENT='[
"class Num:",
"    def __init__(self, num):",
"        self.num = num",
"    def add(self, addition):",
"        self.num += addition"
]'


> PATTERN='[
"class Num:",
"    def add(self, addition):",
"        self.num += addition"
]'

> zaruba lines getIndex $CONTENT $PATTERN
4

> zaruba lines getIndex $CONTENT $PATTERN --index=-1
4

> zaruba list get $PATTERN 0
class Num:
> zaruba lines getIndex $CONTENT $PATTERN --index=0
0

> zaruba list get $PATTERN 1
    def add(self, addition):
> zaruba lines getIndex $CONTENT $PATTERN --index=1
3


Flags:
  -h, --help        help for getIndex
  -i, --index int   desired pattern index (default -1)
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->