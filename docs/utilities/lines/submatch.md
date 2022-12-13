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
Return submatch matching the pattern

Usage:
  zaruba lines submatch <lines> <patterns> [flags]

Examples:

> zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]'
["d"]
> zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=-1
["d"]

    Getting line index that match the last element of the pattern
    lines:        ["a", "a", "b", "c", "d", "e"]
                    0    1    2    3    4    5
                                        ^
                                        line index that match the last index of the pattern
    patterns:     ["a",    , "b",      "d"]
                    0         1         2
                                        ^
                                        last index of the pattern


> zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=1
["b"]

    Getting line index that match the desired index of the pattern
    lines:        ["a", "a", "b", "c", "d", "e"]
                    0    1    2    3    4    5
                              ^
                              line index that match the desired index of the pattern
    patterns:     ["a",    , "b",      "d"]
                    0         1         2
                              ^
                              desired index of the pattern


Flags:
  -h, --help        help for submatch
  -i, --index int   desired pattern index (default -1)
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->