<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🍠 Json](README.md)
# Print
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba json print --help
```
 
<details>
<summary>Output</summary>
 
```````
Print json.
You can print the output into stdout or a file.
You can also define whether you want to "pretty print" or not (It is pretty print by default).

Usage:
  zaruba json print <jsonAny> [jsonFileName] [flags]

Aliases:
  print, write

Examples:

> zaruba json print '{"id": 1, "title": "Doraemon"}'
{
  "id": 1,
  "title": "Doraemon"
}

> zaruba json print '{"id": 1, "title": "Doraemon"}' --pretty=false
{"id":1,"title":"Doraemon"}

> zaruba json print '{"id": 1, "title": "Doraemon"}' book.json
> cat book.json
{
  "id": 1,
  "title": "Doraemon"
}


Flags:
  -h, --help     help for print
  -p, --pretty   pretty print (default true)
```````
</details>
<!--endCode-->


<!--startTocSubtopic-->
<!--endTocSubtopic-->