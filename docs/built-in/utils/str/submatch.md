<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [str](README.md)
# submatch
<!--endTocHeader-->

```

Return submatch of a string matching the pattern.

Example:

String   : 'abcddefghi'
Pattern  : 'abc(d+)(.*)'

Submatch : [
"abcdefghi",   # match the whole pattern
"dd",          # match 'd+'
"efghi"        # match '.*'
]

Usage:
  zaruba str submatch <string> <pattern> [flags]

Examples:

> zaruba str submatch 'abcdefghi' 'abc(d+)(.*)'
["abcdefghi","dd","efghi"]


Flags:
  -h, --help   help for submatch

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->