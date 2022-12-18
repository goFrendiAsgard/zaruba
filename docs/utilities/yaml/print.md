<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🍠 Yaml](README.md)
# Print
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba yaml print --help
```
 
<details>
<summary>Output</summary>
 
```````
Print JSON map or list as YAML

Usage:
  zaruba yaml print <jsonAny> [yamlFileName] [flags]

Aliases:
  print, write

Examples:

> zaruba yaml print '{"id": 1, "title": "Doraemon"}'
id: 1
title: Doraemon

> zaruba yaml print '{"id": 1, "title": "Doraemon"}' book.yaml
> cat book.yaml
id: 1
title: Doraemon



Flags:
  -h, --help   help for print
```````
</details>
<!--endCode-->

# Examples



<!--startTocSubTopic-->
<!--endTocSubTopic-->