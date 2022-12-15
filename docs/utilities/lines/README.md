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
  getIndex     Return index of lines matching the patterns at desiredPatternIndex
  insertAfter  Insert newLine after lines[index]
  insertBefore Insert newLine before lines[index]
  print        Print lines as multiline string
  read         Read lines from a file, return a jsonStrList
  replace      Replace lines[index] with replacements
  submatch     Return submatch matching the pattern

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