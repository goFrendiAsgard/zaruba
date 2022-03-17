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
cd examples/run-tasks
zaruba please printHelloHuman humanName="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.025µs
 Current Time: 09:26:37
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       09:26:37.849 hello Go Frendi
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.581318ms
 Current Time: 09:26:37
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.060293ms
 Current Time: 09:26:38
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
cd examples/run-tasks
zaruba please printHelloHuman -v sample.values.yaml
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.1µs
 Current Time: 09:26:38
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       09:26:38.24  hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.47289ms
 Current Time: 09:26:38
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.713712ms
 Current Time: 09:26:38
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->

If you don't define any value, Zaruba will load `default.values.yaml` as default value.

<!--startTocSubTopic-->
<!--endTocSubTopic-->