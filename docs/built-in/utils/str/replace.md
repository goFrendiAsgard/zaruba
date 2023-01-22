<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [str](README.md)
# replace
<!--endTocHeader-->

```
Replace string by jsonMapReplacement

Usage:
  zaruba str replace <string> [{<jsonMapReplacement> | <key> <value>}] [flags]

Examples:

> zaruba str replace 'Capital of country is city' '{"country": "Indonesia", "city": "Jakarta"}'
Capital of Indonesia is Jakarta

> zaruba str replace 'Capital of country is city' country Japan city Tokyo
Capital of Japan is Tokyo


Flags:
  -h, --help   help for replace

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->