<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [json](README.md)
# set
<!--endTocHeader-->

```
Set value of nested JsonMap or JsonList

Usage:
  zaruba json set <jsonAny> <key> <value> [flags]

Aliases:
  set, write

Examples:

> zaruba json set '{"characters": [{"name": "doraemon"}, {"name": "nobita"}]}' 'characters.[1].name' '"dorami"'
{"characters": [{"name": "doraemon"}, {"name": "dorami"}]}


Flags:
  -h, --help   help for set

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->