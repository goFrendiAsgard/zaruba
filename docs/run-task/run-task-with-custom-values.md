<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# Run task with custom values
<!--endTocHeader-->

Some tasks might dependes on [task inputs](./project/task/task-inputs.md). To set input values for your tasks, you can invoke:

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
 Elapsed Time: 1.202µs
 Current Time: 13:00:40
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:00:40.623 hello Go Frendi
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.35293ms
 Current Time: 13:00:40
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.005097ms
 Current Time: 13:00:40
zaruba please printHelloHuman  -v 'humanName=Go Frendi'
```````
</details>
<!--endCode-->


You can also load `value files`.

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
 Elapsed Time: 1.122µs
 Current Time: 13:00:40
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:00:40.987 hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.829482ms
 Current Time: 13:00:41
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.709411ms
 Current Time: 13:00:41
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->



If you don't define any value, Zaruba will load `default.values.yaml` as default value.


> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ one.



<!--startTocSubTopic-->
<!--endTocSubTopic-->