<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🚈 Lines](README.md)
# Submatch
<!--endTocHeader-->


# Usage

<!--startCode-->
```bash
zaruba lines submatch --help
```
 
<details>
<summary>Output</summary>
 
```````
Getting line index of a a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of line

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1

Usage:
  zaruba lines submatch <jsonStrList> <jsonStrListPatterns> [flags]

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
"( *)def add\\(self, (.*)\\):",
"( *)self\\.num \\+= (.*)"
]'

> zaruba lines submatch $CONTENT $PATTERN
["        self.num += addition","        ","addition"]

> zaruba lines submatch $CONTENT $PATTERN --index=-1
["        self.num += addition","        ","addition"]

> zaruba list get $PATTERN 0
class Num:
> zaruba lines submatch $CONTENT $PATTERN --index=0
["class Num:"]

> zaruba list get $PATTERN 1
( *)def add\(self, (.*)\):
> zaruba lines submatch $CONTENT $PATTERN --index=1
["    def add(self, addition):","    ","addition"]


Flags:
  -h, --help        help for submatch
  -i, --index int   desired pattern index (default -1)
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->