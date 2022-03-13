<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🔤 Run task with custom values
<!--endTocHeader-->

Some tasks might dependes on [task inputs](../core-concepts/task/task-inputs.md).

To set input values for your tasks, you can invoke:

```bash
zaruba please <task-name> -v <first-key=first-value> -v <second-key=second-value>
```

or shorter:


```bash
zaruba please <task-name> <first-key=first-value> <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman humanName="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.294µs
 Current Time: 12:34:05
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       12:34:05.554 hello Go Frendi
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.901319ms
 Current Time: 12:34:05
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.588799ms
 Current Time: 12:34:05
zaruba please printHelloHuman  -v 'humanName=Go Frendi'
```````
</details>
<!--endCode-->

> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ one.

# Load Value File

You can load `value files` by performing.

```bash
zaruba please <task-name> -v <first-file.value.yaml> -v <second-file.value.yaml>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman -v sample.values.yaml
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.246µs
 Current Time: 12:34:05
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       12:34:05.907 hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.247124ms
 Current Time: 12:34:06
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.35468ms
 Current Time: 12:34:06
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->

If you don't define any value, Zaruba will load `default.values.yaml` as default value.

<!--startTocSubTopic-->
<!--endTocSubTopic-->