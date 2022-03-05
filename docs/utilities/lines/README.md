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
Lines manipulation utilities

Usage:
  zaruba lines [command]

Available Commands:
  fill         Insert suplements to lines if patterns is not found
  getIndex     Return index of matching the pattern
  insertAfter  Insert newLine after lines[index]
  insertBefore Insert newLine before lines[index]
  read         Read lines from file
  replace      Replace lines[index] with replacements
  submatch     Return submatch matching the pattern
  write        Write list to file

Flags:
  -h, --help   help for lines

Use "zaruba lines [command] --help" for more information about a command.
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
# Sub-topics
* [Fill](fill.md)
* [GetIndex](get-index.md)
* [InsertAfter](insert-after.md)
* [InsertBefore](insert-before.md)
* [Read](read.md)
* [Replace](replace.md)
* [Submatch](submatch.md)
* [Write](write.md)
<!--endTocSubTopic-->