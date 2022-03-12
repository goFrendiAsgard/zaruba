<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Read
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env read --help
```
 
<details>
<summary>Output</summary>
 
```````
Read envmap from file

Usage:
  zaruba env read <fileName> [prefix] [flags]

Flags:
  -h, --help   help for read
```````
</details>
<!--endCode-->

# Examples

## Print Environment

<!--startCode-->
```bash
cd examples/tasks
zaruba env read sample.env
```
 
<details>
<summary>Output</summary>
 
```````
{"GREETINGS":"Hola"}
```````
</details>
<!--endCode-->

<!--startTocSubTopic-->
<!--endTocSubTopic-->