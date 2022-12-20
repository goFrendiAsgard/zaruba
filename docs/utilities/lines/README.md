<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md)
# 🚈 Lines
<!--endTocHeader-->

A `lines` is a valid JSON array containing only string elements.

To see list of `lines` utilities you can invoke:

<!--startCode-->
```bash
zaruba lines
```
 
<details>
<summary>Output</summary>
 
```````
JsonStrList manipulation utilities

Usage:
  zaruba lines [command]

Available Commands:
  getIndex     Return the index of a line matching a particular index at a specified patterns
  insertAfter  Insert a new lines into jsonStringList after a particular index
  insertBefore Insert new lines into a jsonStringList before a particular index
  print        Print jsontStrList as a multiline string
  read         Read a text file and return a jsonStrList
  replace      Replace a jsonStringList at a particular index with new lines
  submatch     Return submatch matching the pattern at a desired pattern index

Flags:
  -h, --help   help for lines

Use "zaruba lines [command] --help" for more information about a command.
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
# Sub-topics
* [GetIndex](getindex.md)
* [InsertAfter](insertafter.md)
* [InsertBefore](insertbefore.md)
* [Read](read.md)
* [Replace](replace.md)
* [Submatch](submatch.md)
* [Write](write.md)
<!--endTocSubTopic-->