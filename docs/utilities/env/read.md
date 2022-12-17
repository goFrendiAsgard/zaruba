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
Read environment variable declarations from environment file as a jsonMap

Usage:
  zaruba env read <strFileName> [flags]

Flags:
  -h, --help            help for read
  -p, --prefix string   environment prefix
```````
</details>
<!--endCode-->

# Examples

## Print Environment

<!--startCode-->
```bash
cd examples/run-tasks
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