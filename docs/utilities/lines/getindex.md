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
Return index of lines matching the patterns at desiredPatternIndex

Usage:
  zaruba lines getIndex <jsonStrList> <jsonStrListPatterns> [flags]

Examples:

Getting line index that match the last element of the pattern
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]'
    4
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=-1
    4

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                                    ^
                                    line index that match the last index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                                    ^
                                    last index of the pattern

Getting line index that match the desired index of the pattern
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=1
    2

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                          ^
                          line index that match the desired index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                          ^
                          desired index of the pattern


Flags:
  -h, --help        help for getIndex
  -i, --index int   desired pattern index (default -1)
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->